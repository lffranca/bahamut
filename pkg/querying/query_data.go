package querying

import "context"

type QueryData interface {
	Query(ctx context.Context, query string, variables []interface{}) (columns []string, columnTypes []string, values [][]interface{}, err error)
}
