package querying

import "context"

type FormatterData interface {
	Transform(ctx context.Context, template []byte, input interface{}) (expected []byte, err error)
}
