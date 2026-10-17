package tachosync

import (
	"context"
	"net/http"

	httpclient "github.com/rojakc96/tachosync-go/client"
	"github.com/rojakc96/tachosync-go/models"
)

func (c *Client) ListDrivers(ctx context.Context, params models.ListDriversParams) (*models.DriversPage, error) {
	return requestJSON[models.DriversPage](ctx, c, http.MethodGet, DriversEndpoint, params, nil)
}

func (c *Client) CreateDriver(ctx context.Context, command models.CreateDriverCommand) (string, error) {
	response, err := c.request(ctx, http.MethodPost, DriversEndpoint, nil, command, nil)
	return locationID(response), err
}

func (c *Client) EditDriver(ctx context.Context, command models.EditDriverCommand) error {
	_, err := c.request(ctx, http.MethodPut, DriversEndpoint, nil, command, nil)
	return err
}

func (c *Client) GetDriver(ctx context.Context, id string) (*models.DriverDetailedResponse, error) {
	return requestJSON[models.DriverDetailedResponse](ctx, c, http.MethodGet, DriverByIDEndpoint, nil, nil, httpclient.PathParam("id", id))
}

func (c *Client) DeleteDriver(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, DriverByIDEndpoint, nil, nil, nil, httpclient.PathParam("id", id))
	return err
}

func (c *Client) GetDriverFilters(ctx context.Context, params models.DriverFilterParams) ([]models.FilterItem, error) {
	result, err := requestJSON[[]models.FilterItem](ctx, c, http.MethodGet, DriverFiltersEndpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}
