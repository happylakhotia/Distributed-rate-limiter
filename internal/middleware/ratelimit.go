package middleware

import (
	"net"
	"net/http"

	"github.com/amitcode/distributed-rate-limiter/internal/limiter"
)

func RateLimit(l limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract IP only (without port)
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr
			}
			key := "ratelimit:" + ip

			allowed, err := l.Allow(key)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			if !allowed {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}


