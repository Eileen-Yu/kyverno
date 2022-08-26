package engine

import (
	"github.com/go-logr/logr"
	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	"github.com/kyverno/kyverno/pkg/engine/response"
)

// IsResponseSuccessful return true if all responses are successful
func IsResponseSuccessful(engineReponses []*response.EngineResponse) bool {
	for _, er := range engineReponses {
		if !er.IsSuccessful() {
			return false
		}
	}
	return true
}

// BlockRequest returns true when:
// 1. a policy fails (i.e. creates a violation) and validationFailureAction is set to 'enforce'
// 2. a policy has a processing error and failurePolicy is set to 'Fail`
func BlockRequest(er *response.EngineResponse, failurePolicy kyvernov1.FailurePolicyType, log logr.Logger) bool {
	log.V(0).Info("Block request is called")
	action := er.GetValidationFailureAction(log)
	log.V(0).Info("Action has been got")
	if er.IsFailed() && action.Enforce() {
		return true
	}
	if er.IsError() && failurePolicy == kyvernov1.Fail {
		return true
	}
	return false
}
