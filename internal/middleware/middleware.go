package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/MCPutro/E-commerce/pkg/constant"
	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "XRequestID"

func NewMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := uuid.New().String()
		ctx = context.WithValue(ctx, requestIDKey, requestID)

		w.Header().Set(constant.HeaderXRequestID, requestID)

		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Println(time.Since(start), r.Method, r.URL.Path)
	})
}
