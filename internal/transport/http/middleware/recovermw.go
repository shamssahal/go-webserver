package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/shamssahal/go-server/pkg/errors"
	"github.com/shamssahal/go-server/pkg/utils"
)

// Recover recovers from panics and returns a 500 Internal Server Error.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// Capture stack trace
				stack := debug.Stack()

				// Format panic value based on its type
				var panicErr string
				if err, ok := rec.(error); ok {
					// It's an error type, use Error() method
					panicErr = err.Error()
				} else {
					// It's something else (string, int, struct, etc.), format it
					panicErr = fmt.Sprintf("%v", rec)
				}

				slog.Error("panic recovered",
					"error", panicErr,
					"stack", string(stack),
					"request_id", utils.RequestIDFromContext(r.Context()),
					"method", r.Method,
					"path", r.URL.Path,
					"remote", r.RemoteAddr,
					"user_agent", r.UserAgent(),
				)
				errors.WriteError(w, errors.NewError(500, "internal server error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
