package admission

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yndd/ndd-runtime/pkg/logging"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
	"github.com/yndd/nddp-srl3/internal/webhook/validation"
	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// Admitter is a container for admission business
type Admitter struct {
	Logger  logging.Logger
	Request *admissionv1.AdmissionRequest
}

// MutatePodReview takes an admission request and validates the pod within
// it returns an admission review
func (a Admitter) Validate() (*admissionv1.AdmissionReview, error) {
	d, err := a.Srl3Device()
	if err != nil {
		e := fmt.Sprintf("could not parse pod in admission review request: %v", err)
		return reviewResponse(a.Request.UID, false, http.StatusBadRequest, e), err
	}

	v := validation.NewValidator(a.Logger)
	val, err := v.ValidateDeviceConfig(d)
	if err != nil {
		e := fmt.Sprintf("could not validate device config: %v", err)
		return reviewResponse(a.Request.UID, false, http.StatusBadRequest, e), err
	}

	if !val.Valid {
		return reviewResponse(a.Request.UID, false, http.StatusForbidden, val.Reason), nil
	}

	return reviewResponse(a.Request.UID, true, http.StatusAccepted, "valid device config"), nil
}

// Pod extracts a pod from an admission request
func (a Admitter) Srl3Device() (*srlv1alpha1.Srl3Device, error) {
	if a.Request.Kind.Kind != "Srl3Device" {
		return nil, fmt.Errorf("only srl3devices are supported here")
	}

	i := srlv1alpha1.Srl3Device{}
	if err := json.Unmarshal(a.Request.Object.Raw, &i); err != nil {
		return nil, err
	}

	return &i, nil
}

func reviewResponse(uid types.UID, allowed bool, httpCode int32,
	reason string) *admissionv1.AdmissionReview {
	return &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1",
		},
		Response: &admissionv1.AdmissionResponse{
			UID:     uid,
			Allowed: allowed,
			Result: &metav1.Status{
				Code:    httpCode,
				Message: reason,
			},
		},
	}
}

// patchReviewResponse builds an admission review with given json patch
func patchReviewResponse(uid types.UID, patch []byte) (*admissionv1.AdmissionReview, error) {
	patchType := admissionv1.PatchTypeJSONPatch

	return &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1",
		},
		Response: &admissionv1.AdmissionResponse{
			UID:       uid,
			Allowed:   true,
			PatchType: &patchType,
			Patch:     patch,
		},
	}, nil
}
