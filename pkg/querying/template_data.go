package querying

import "context"

type TemplateData interface {
	ByID(ctx context.Context, id *string) ([]byte, error)
}
