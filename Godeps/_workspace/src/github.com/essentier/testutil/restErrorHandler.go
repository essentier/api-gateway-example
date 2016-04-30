package testutil

import "testing"

type RestErrorHandler interface {
	HandleError(err error, message string)
}

type failTestRestErrHanlder struct {
	t *testing.T
}

func (h *failTestRestErrHanlder) HandleError(err error, message string) {
	if err != nil {
		h.t.Fatalf(message+" Error is: %v.", err)
	}
}
