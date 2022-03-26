package validation

import (
	"fmt"
	"strings"

	"github.com/yndd/ndd-runtime/pkg/logging"
	srlv1alpha1 "github.com/yndd/nddp-srl3/apis/srl3/v1alpha1"
)

// nameValidator is a container for validating the name of pods
type nameValidator struct {
	Logger logging.Logger
}

// nameValidator implements the podValidator interface
var _ itfceValidator = (*nameValidator)(nil)

// Name returns the name of nameValidator
func (n nameValidator) Name() string {
	return "name_validator"
}

// Validate inspects the name of a given pod and returns validation.
// The returned validation is only valid if the pod name does not contain some
// bad string.
func (n nameValidator) Validate(deviceConfig *srlv1alpha1.Srl3Device) (validation, error) {
	badString := "offensive"

	if strings.Contains(deviceConfig.Name, badString) {
		v := validation{
			Valid:  false,
			Reason: fmt.Sprintf("deviceConfig name contains %q", badString),
		}
		return v, nil
	}

	return validation{Valid: true, Reason: "valid name"}, nil
}
