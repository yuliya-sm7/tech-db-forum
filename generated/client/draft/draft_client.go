package draft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new draft API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for draft API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BatchForumCreate добавлениеs новых форумов

Массовое добавление Форумов в базу.

*/
func (a *Client) BatchForumCreate(params *BatchForumCreateParams) (*BatchForumCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchForumCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batchForumCreate",
		Method:             "POST",
		PathPattern:        "/batch/forum/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BatchForumCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BatchForumCreateOK), nil

}

/*
BatchUserCreate добавлениеs новых пользователей

Массовое добавление пользователей в базу.

*/
func (a *Client) BatchUserCreate(params *BatchUserCreateParams) (*BatchUserCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchUserCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batchUserCreate",
		Method:             "POST",
		PathPattern:        "/batch/user/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BatchUserCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BatchUserCreateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}