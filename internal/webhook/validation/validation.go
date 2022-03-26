package validation

import (
	"github.com/yndd/ndd-runtime/pkg/logging"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
)

// Validator is a container for mutation
type Validator struct {
	Logger logging.Logger
}

// NewValidator returns an initialised instance of Validator
func NewValidator(logger logging.Logger) *Validator {
	return &Validator{Logger: logger}
}

// itfceValidator is an interface used to group functions validating pods
type itfceValidator interface {
	Validate(*srlv1alpha1.Srl3Device) (validation, error)
	Name() string
}

type validation struct {
	Valid  bool
	Reason string
}

// ValidatePod returns true if a pod is valid
func (v *Validator) ValidateDeviceConfig(devcieConfig *srlv1alpha1.Srl3Device) (validation, error) {
	var deviceConfigName string
	if devcieConfig.Name != "" {
		deviceConfigName = devcieConfig.Name
	} else {
		if devcieConfig.ObjectMeta.GenerateName != "" {
			deviceConfigName = devcieConfig.ObjectMeta.GenerateName
		}
	}
	log := v.Logger.WithValues("device_config_name", deviceConfigName)
	log.Debug("delete me")

	// list of all validations to be applied to the pod
	validations := []itfceValidator{
		nameValidator{v.Logger},
	}

	// apply all validations
	for _, v := range validations {
		var err error
		vp, err := v.Validate(devcieConfig)
		if err != nil {
			return validation{Valid: false, Reason: err.Error()}, err
		}
		if !vp.Valid {
			return validation{Valid: false, Reason: vp.Reason}, err
		}
	}

	return validation{Valid: true, Reason: "valid device config"}, nil
}
