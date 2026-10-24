package tachosync

import (
	"context"
	"net/http"

	httpclient "github.com/rojack96/tachosync-go/internal/client"
	"github.com/rojack96/tachosync-go/models"
)

func (c *Client) GetCompanyCard(ctx context.Context, id string) (*models.CompanyCardResponse, error) {
	return requestJSON[models.CompanyCardResponse](ctx, c, http.MethodGet, CompanyCardByIDEndpoint, nil, nil, httpclient.PathParam("id", id))
}

func (c *Client) DeleteCompanyCard(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, CompanyCardByIDEndpoint, nil, nil, nil, httpclient.PathParam("id", id))
	return err
}

func (c *Client) ListCompanyCards(ctx context.Context, params models.ListCompanyCardsParams) (*models.CompanyCardsPage, error) {
	return requestJSON[models.CompanyCardsPage](ctx, c, http.MethodGet, CompanyCardsEndpoint, params, nil)
}

func (c *Client) CreateCompanyCard(ctx context.Context, command models.CreateCompanyCardCommand) (string, error) {
	response, err := c.request(ctx, http.MethodPost, CompanyCardsEndpoint, nil, command, nil)
	return locationID(response), err
}

func (c *Client) EditCompanyCard(ctx context.Context, command models.EditCompanyCardCommand) error {
	_, err := c.request(ctx, http.MethodPut, CompanyCardsEndpoint, nil, command, nil)
	return err
}
