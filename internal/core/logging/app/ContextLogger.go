package app

import "context"

func ContextLogger(ctx context.Context, factory LoggerFactory) Logger {
	existingFields := ctx.Value(contextKey)
	if existingFields == nil {
		return factory([]Fields{})
	}
	return factory(existingFields.([]Fields))
}
