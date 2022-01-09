package querying

import "context"

type QueryService service

func (pkg *QueryService) Execute(ctx context.Context, queryID, formatID *string, value interface{}) ([]byte, error) {
	return nil, nil
}
