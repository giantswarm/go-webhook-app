package aws

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HAUpdater knows how to mark Kubernetes resources.
type HAUpdater interface {
	HAUpdate(ctx context.Context, obj metav1.Object) error
}

// NewLabelHAUpdater returns a new marker that will mark with labels.
func NewHAUpdater(marks map[string]string) HAUpdater {
	return labelmarker{marks: marks}
}

type labelmarker struct {
	marks map[string]string
}

func (l labelmarker) HAUpdate(_ context.Context, obj metav1.Object) error {
	labels := obj.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}

	for k, v := range l.marks {
		labels[k] = v
	}

	obj.SetLabels(labels)
	return nil
}

// DummyHAUpdater is a marker that doesn't do anything.
var DummyHAUpdater HAUpdater = dummyMaker(0)

type dummyMaker int

func (dummyMaker) HAUpdate(_ context.Context, _ metav1.Object) error { return nil }
