package middlewares

import (
	"net/http"

	"github.com/getsentry/sentry-go"
)

func SentryScopeMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sentry.GetHubFromContext(r.Context())

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetRequest(r)

				scope.SetTag("method", r.Method)
				scope.SetTag("url", r.URL.Path)
				scope.SetTag("Type", "Request")

				scope.SetExtra("headers", r.Header)
				scope.SetExtra("query", r.URL.Query())
				scope.SetExtra("body", r.Body)
				scope.SetExtra("remote_addr", r.RemoteAddr)

			})

			next.ServeHTTP(w, r)
		})
	}
}
