package middlewares

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

func CaptureError(err error) {
	if err != nil {
		wrappedErr := errors.Wrap(err, "captured error")
		sentry.CaptureException(wrappedErr)
		sentry.Flush(2 * time.Second)
	}
}
