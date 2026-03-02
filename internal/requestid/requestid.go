package requestid

import "context"

type contextKey string

const RequestIDKey contextKey = "requestID"

func FromContext(ctx context.Context) string {
	if v := ctx.Value(RequestIDKey); v != nil {
		if id, ok := v.(string); ok {
			return id
		}
	}
	return ""
}
