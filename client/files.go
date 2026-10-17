package tachosync

import (
	"context"
	"net/http"

	"github.com/rojakc96/tachosync-go/models"
)

func (c *Client) ListDriverFiles(ctx context.Context, params models.ListDriverFilesParams) (*models.DriverFilesPage, error) {
	return requestJSON[models.DriverFilesPage](ctx, c, http.MethodGet, DriverFilesEndpoint, params, nil)
}

func (c *Client) GetDriverFileFilters(ctx context.Context, params models.FileFilterParams[models.DriverFileFiltersType]) ([]models.FilterItem, error) {
	result, err := requestJSON[[]models.FilterItem](ctx, c, http.MethodGet, DriverFileFiltersEndpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) DownloadDriverFiles(ctx context.Context, params models.DownloadFilesParams) ([]byte, error) {
	var result []byte
	_, err := c.request(ctx, http.MethodGet, DriverFilesDownloadEndpoint, params, nil, &result)
	return result, err
}

func (c *Client) ListVehicleFiles(ctx context.Context, params models.ListVehicleFilesParams) (*models.VehicleFilesPage, error) {
	return requestJSON[models.VehicleFilesPage](ctx, c, http.MethodGet, VehicleFilesEndpoint, params, nil)
}

func (c *Client) GetVehicleFileFilters(ctx context.Context, params models.FileFilterParams[models.VehicleFileFiltersType]) ([]models.FilterItem, error) {
	result, err := requestJSON[[]models.FilterItem](ctx, c, http.MethodGet, VehicleFileFiltersEndpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) DownloadVehicleFiles(ctx context.Context, params models.DownloadFilesParams) ([]byte, error) {
	var result []byte
	_, err := c.request(ctx, http.MethodGet, VehicleFilesDownloadEndpoint, params, nil, &result)
	return result, err
}
