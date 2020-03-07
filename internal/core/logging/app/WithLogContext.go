package app

import "context"

// Creates a copy of its first argument, with a new logging context
// All the logging fields of the previous logging context are shallow-copied into the new one
func WithLogContext(ctx context.Context, fields Fields) context.Context {
	existingFields := ctx.Value(contextKey)
	var actualFields = []Fields{fields}
	if existingFields != nil {
		actualFields = append(actualFields, existingFields.([]Fields)...)
	}
	return context.WithValue(ctx, contextKey, actualFields)
}
