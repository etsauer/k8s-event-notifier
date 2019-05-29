package apis

import (
	eventv1 "github.com/redhat-cop/k8s-notify/pkg/apis/event/v1"
	notifyv1 "github.com/redhat-cop/k8s-notify/pkg/apis/notify/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

type SharedResources struct {
	Subscriptions map[string]eventv1.EventSubscription
	Notifiers     map[string]notifyv1.Notifier
}

func NewSharedResources() SharedResources {
	es := make(map[string]eventv1.EventSubscription)
	n := make(map[string]notifyv1.Notifier)
	return SharedResources{
		Subscriptions: es,
		Notifiers:     n,
	}
}
