package app

import (
	"context"
	logging "github.com/helstern/kommol/internal/core/logging/app"
	"testing"
)

func TestWithLogContextCopiesFields(t *testing.T) {
	ctxParent := context.WithValue(context.Background(), "name", "parent")

	ctxOne := logging.WithLogContext(ctxParent, logging.Fields{
		"context": "one",
	})
	ctxTwo := logging.WithLogContext(ctxOne, logging.Fields{
		"context": "two",
	})

	var (
		ctxOneFields []logging.Fields
		ctxTwoFields []logging.Fields
	)

	logging.ContextLogger(ctxOne, func(fields []logging.Fields) logging.Logger {
		ctxOneFields = fields
		return nil
	})
	logging.ContextLogger(ctxTwo, func(fields []logging.Fields) logging.Logger {
		ctxTwoFields = fields
		return nil
	})

	if len(ctxOneFields) != 1 {
		t.Logf("ctxOneFields size := %d", len(ctxOneFields))
		t.Errorf("WithLogContext should create the logging context if it does not exist")
	}

	if len(ctxTwoFields) != 2 {
		t.Logf("ctxTwoFields size := %d", len(ctxTwoFields))
		t.Errorf("WithLogContext should copy the existing logging context")
	}
}
