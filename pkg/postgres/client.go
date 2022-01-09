package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lffranca/queryngo/pkg/util"
	"github.com/lib/pq"
	"log"
)

type Option func(*Client) error

func WithConnectionString(conn string) Option {
	return func(client *Client) error {
		db, err := sql.Open("postgres", conn)
		if err != nil {
			return err
		}

		client.db = db
		return nil
	}
}

func WithDB(db *sql.DB) Option {
	return func(client *Client) error {
		if db == nil {
			return errors.New("*sql.DB is required")
		}

		client.db = db
		return nil
	}
}

func New(options ...Option) (*Client, error) {
	client := new(Client)

	for _, op := range options {
		if err := op(client); err != nil {
			return nil, err
		}
	}

	if client.db == nil {
		return nil, errors.New("db param is required")
	}

	return client, nil
}

type Client struct {
	db *sql.DB
}

func (pkg *Client) query(ctx context.Context, query string, variables []interface{}) (*sql.Rows, error) {
	if variables == nil || len(variables) <= 0 {
		return pkg.db.QueryContext(ctx, query)
	}

	var args []interface{}
	for _, varItem := range variables {
		args = append(args, pq.Array(varItem))
	}

	return pkg.db.QueryContext(ctx, query, args...)
}

func (pkg *Client) Query(ctx context.Context, query string, variables []interface{}) ([]string, []string, [][]interface{}, error) {
	rows, err := pkg.query(ctx, query, variables)
	if err != nil {
		return nil, nil, nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("ERROR CLOSE ROWS: ", err)
		}
	}()

	return util.SQLRowModel(rows)
}
