package tachosync

import (
	"context"
	"net/http"

	httpclient "github.com/rojack96/tachosync-go/internal/client"
	"github.com/rojack96/tachosync-go/models"
)

func (c *Client) GetCompanies(ctx context.Context) (*models.CompanyResponse, error) {
	return requestJSON[models.CompanyResponse](ctx, c, http.MethodGet, CompaniesEndpoint, nil, nil)
}

func (c *Client) CreateCompany(ctx context.Context, command models.CreateCompanyCommand) (string, error) {
	response, err := c.request(ctx, http.MethodPost, CompaniesEndpoint, nil, command, nil)
	return locationID(response), err
}

func (c *Client) EditCompany(ctx context.Context, command models.EditCompanyCommand) (*models.CompanyResponse, error) {
	return requestJSON[models.CompanyResponse](ctx, c, http.MethodPut, CompaniesEndpoint, nil, command)
}

func (c *Client) GetCompany(ctx context.Context, id string) (*models.CompanyDetailsResponse, error) {
	return requestJSON[models.CompanyDetailsResponse](ctx, c, http.MethodGet, CompanyByIDEndpoint, nil, nil, httpclient.PathParam("id", id))
}

func (c *Client) DeleteCompanies(ctx context.Context, command models.DeleteCompaniesCommand) error {
	_, err := c.request(ctx, http.MethodPost, CompaniesBulkDeleteEndpoint, nil, command, nil)
	return err
}
