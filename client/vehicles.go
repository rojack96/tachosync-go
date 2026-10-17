package tachosync

import (
	"context"
	"net/http"

	httpclient "github.com/rojakc96/tachosync-go/internal/client"
	"github.com/rojakc96/tachosync-go/internal/models"
)

func (c *Client) ListVehicles(ctx context.Context, params models.ListVehiclesParams) (*models.VehiclesPage, error) {
	return requestJSON[models.VehiclesPage](ctx, c, http.MethodGet, VehiclesEndpoint, params, nil)
}

func (c *Client) CreateVehicle(ctx context.Context, command models.CreateVehicleCommand) (string, error) {
	response, err := c.request(ctx, http.MethodPost, VehiclesEndpoint, nil, command, nil)
	return locationID(response), err
}

func (c *Client) EditVehicle(ctx context.Context, command models.EditVehicleCommand) error {
	_, err := c.request(ctx, http.MethodPut, VehiclesEndpoint, nil, command, nil)
	return err
}

func (c *Client) GetVehicle(ctx context.Context, id string) (*models.VehicleDetailedResponse, error) {
	return requestJSON[models.VehicleDetailedResponse](ctx, c, http.MethodGet, VehicleByIDEndpoint, nil, nil, httpclient.PathParam("id", id))
}

func (c *Client) DeleteVehicle(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, VehicleByIDEndpoint, nil, nil, nil, httpclient.PathParam("id", id))
	return err
}

func (c *Client) GetVehicleDownloadsHistory(ctx context.Context, id string) ([]models.DownloadHistoryResponse, error) {
	result, err := requestJSON[[]models.DownloadHistoryResponse](ctx, c, http.MethodGet, VehicleDownloadsHistoryEndpoint, nil, nil, httpclient.PathParam("id", id))
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) GetVehicleFilters(ctx context.Context, params models.VehicleFilterParams) ([]models.FilterItem, error) {
	result, err := requestJSON[[]models.FilterItem](ctx, c, http.MethodGet, VehicleFiltersEndpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}
