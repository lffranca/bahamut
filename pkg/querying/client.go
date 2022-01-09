package querying

import "errors"

type Option func(*Client) error

func WithFormatter(formatter FormatterData) Option {
	return func(client *Client) error {
		if formatter == nil {
			return errors.New("FormatterData is required")
		}

		client.formatterData = formatter
		return nil
	}
}

func WithQuery(query QueryData) Option {
	return func(client *Client) error {
		if query == nil {
			return errors.New("QueryData is required")
		}

		client.queryData = query
		return nil
	}
}

func WithTemplate(temp TemplateData) Option {
	return func(client *Client) error {
		if temp == nil {
			return errors.New("TemplateData is required")
		}

		client.templateData = temp
		return nil
	}
}

func New(options ...Option) (*Client, error) {
	client := new(Client)
	client.common.Client = client

	for _, op := range options {
		if err := op(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

type Client struct {
	common        service
	formatterData FormatterData
	queryData     QueryData
	templateData  TemplateData
}

type service struct {
	Client *Client
}
