package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/yndd/ndd-runtime/pkg/logging"
	"github.com/yndd/nddp-srl3/internal/webhook/admission"
	admissionv1 "k8s.io/api/admission/v1"
)

// Option can be used to manipulate Options.
type Option func(Server)

// WithLogger specifies how the Reconciler should log messages.
func WithLogger(log logging.Logger) Option {
	return func(s Server) {
		s.WithLogger(log)
	}
}

type Server interface {
	WithLogger(log logging.Logger)
	Start() error
}

func New(addr string, opts ...Option) Server {

	mux := http.NewServeMux()

	s := &server{
		server: &http.Server{
			Addr:           addr,
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20, // 1048576
		},
	}

	mux.HandleFunc("/", s.handleRoot)
	mux.HandleFunc("/mutate", s.handleMutate)
	mux.HandleFunc("/validate", s.handleValidate)

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type server struct {
	server *http.Server
	// logging and parsing
	log logging.Logger
}

func (s *server) WithLogger(log logging.Logger) {
	s.log = log
}

func (s *server) Start() error {
	return s.server.ListenAndServeTLS("/cert/cert.crt", "/cert/cert.key")
}

func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
}

// ServeValidatePods validates an admission request and then writes an admission
// review to `w`
func (s *server) handleValidate(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithValues("uri", r.RequestURI)
	log.Debug("received validation request")

	in, err := parseRequest(*r)
	if err != nil {
		log.Debug("error", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adm := admission.Admitter{
		Logger:  log,
		Request: in.Request,
	}

	out, err := adm.Validate()
	if err != nil {
		e := fmt.Sprintf("could not generate admission response: %v", err)
		log.Debug("error", "error", e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jout, err := json.Marshal(out)
	if err != nil {
		e := fmt.Sprintf("could not parse admission response: %v", err)
		log.Debug("error", "error", e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	log.Debug("sending response", "response", jout)
	fmt.Fprintf(w, "%s", jout)
}

func (s *server) handleMutate(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithValues("uri", r.RequestURI)
	log.Debug("received mutation request")

	in, err := parseRequest(*r)
	if err != nil {
		log.Debug("error", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adm := admission.Admitter{
		Logger:  log,
		Request: in.Request,
	}

	out, err := adm.Validate()
	if err != nil {
		e := fmt.Sprintf("could not generate admission response: %v", err)
		log.Debug("error", "error", e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jout, err := json.Marshal(out)
	if err != nil {
		e := fmt.Sprintf("could not parse admission response: %v", err)
		log.Debug("error", "error", e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	log.Debug("sending response", "response", jout)
	fmt.Fprintf(w, "%s", jout)
}

// parseRequest extracts an AdmissionReview from an http.Request if possible
func parseRequest(r http.Request) (*admissionv1.AdmissionReview, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("Content-Type: %q should be %q",
			r.Header.Get("Content-Type"), "application/json")
	}

	bodybuf := new(bytes.Buffer)
	bodybuf.ReadFrom(r.Body)
	body := bodybuf.Bytes()

	if len(body) == 0 {
		return nil, fmt.Errorf("admission request body is empty")
	}

	var a admissionv1.AdmissionReview

	if err := json.Unmarshal(body, &a); err != nil {
		return nil, fmt.Errorf("could not parse admission review request: %v", err)
	}

	if a.Request == nil {
		return nil, fmt.Errorf("admission review can't be used: Request field is nil")
	}

	return &a, nil
}
