package querying

import "context"

type QueryService service

func (pkg *QueryService) Execute(ctx context.Context, queryID, formatID *string, value interface{}) ([]byte, error) {
	queryTemplate, err := pkg.Client.templateData.ByID(ctx, queryID)
	if err != nil {
		return nil, err
	}

	formatTemplate, err := pkg.Client.templateData.ByID(ctx, formatID)
	if err != nil {
		return nil, err
	}

	query, err := pkg.Client.formatterData.Transform(ctx, queryTemplate, value)
	if err != nil {
		return nil, err
	}

	columns, columnTypes, values, err := pkg.Client.queryData.Query(ctx, string(query), nil)
	if err != nil {
		return nil, err
	}

	return pkg.Client.formatterData.Transform(ctx, formatTemplate, map[string]interface{}{
		"Columns":     columns,
		"ColumnTypes": columnTypes,
		"Values":      values,
	})
}
