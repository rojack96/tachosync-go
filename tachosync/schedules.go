package tachosync

import (
	"context"
	"net/http"

	httpclient "github.com/rojakc96/tachosync-go/client"
	"github.com/rojakc96/tachosync-go/models"
)

func (c *Client) ListDriverSchedules(ctx context.Context, params models.ListSchedulesParams) ([]models.DriverScheduleItem, error) {
	result, err := requestJSON[[]models.DriverScheduleItem](ctx, c, http.MethodGet, DriverSchedulesEndpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) CreateDriverSchedule(ctx context.Context, command models.CreateDriverScheduleCommand) (string, error) {
	response, err := c.request(ctx, http.MethodPost, DriverSchedulesEndpoint, nil, command, nil)
	return locationID(response), err
}

func (c *Client) EditDriverSchedule(ctx context.Context, command models.EditDriverScheduleCommand) (*models.DriverScheduleResponse, error) {
	return requestJSON[models.DriverScheduleResponse](ctx, c, http.MethodPut, DriverSchedulesEndpoint, nil, command)
}

func (c *Client) GetDriverSchedule(ctx context.Context, id string) (*models.DriverScheduleResponse, error) {
	return requestJSON[models.DriverScheduleResponse](ctx, c, http.MethodGet, DriverScheduleByIDEndpoint, nil, nil, httpclient.PathParam("id", id))
}

func (c *Client) DeleteDriverSchedule(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, DriverScheduleByIDEndpoint, nil, nil, nil, httpclient.PathParam("id", id))
	return err
}

func (c *Client) ApplyDriverSchedule(ctx context.Context, command models.ApplyDriverScheduleCommand) error {
	_, err := c.request(ctx, http.MethodPut, DriverScheduleApplyEndpoint, nil, command, nil)
	return err
}

func (c *Client) ApplyDriverScheduleToAll(ctx context.Context, command models.ApplyDriverScheduleToAllCommand) error {
	_, err := c.request(ctx, http.MethodPut, DriverScheduleApplyToAllEndpoint, nil, command, nil)
	return err
}

func (c *Client) CreateQuickDriverSchedule(ctx context.Context, command models.CreateQuickDriverScheduleCommand) (*models.CreateQuickScheduleResult, error) {
	return requestJSON[models.CreateQuickScheduleResult](ctx, c, http.MethodPost, DriverQuickScheduleEndpoint, nil, command)
}

func (c *Client) CreateBulkQuickDriverSchedules(ctx context.Context, command models.CreateBulkQuickDriverScheduleCommand) (*models.CreateBulkQuickDriverScheduleResult, error) {
	return requestJSON[models.CreateBulkQuickDriverScheduleResult](ctx, c, http.MethodPost, DriverBulkQuickScheduleEndpoint, nil, command)
}

func (c *Client) RemoveDriverSchedule(ctx context.Context, driverID, scheduleID string) error {
	_, err := c.request(ctx, http.MethodDelete, DriverScheduleRemoveEndpoint, nil, nil, nil,
		httpclient.PathParam("driverId", driverID),
		httpclient.PathParam("scheduleId", scheduleID),
	)
	return err
}

func (c *Client) ListVehicleSchedules(ctx context.Context, params models.ListSchedulesParams) ([]models.VehicleScheduleItem, error) {
	result, err := requestJSON[[]models.VehicleScheduleItem](ctx, c, http.MethodGet, VehicleSchedulesEndpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) CreateVehicleSchedule(ctx context.Context, command models.CreateVehicleScheduleCommand) (string, error) {
	response, err := c.request(ctx, http.MethodPost, VehicleSchedulesEndpoint, nil, command, nil)
	return locationID(response), err
}

func (c *Client) EditVehicleSchedule(ctx context.Context, command models.EditVehicleScheduleCommand) (*models.VehicleScheduleResponse, error) {
	return requestJSON[models.VehicleScheduleResponse](ctx, c, http.MethodPut, VehicleSchedulesEndpoint, nil, command)
}

func (c *Client) GetVehicleSchedule(ctx context.Context, id string) (*models.VehicleScheduleResponse, error) {
	return requestJSON[models.VehicleScheduleResponse](ctx, c, http.MethodGet, VehicleScheduleByIDEndpoint, nil, nil, httpclient.PathParam("id", id))
}

func (c *Client) DeleteVehicleSchedule(ctx context.Context, id string) error {
	_, err := c.request(ctx, http.MethodDelete, VehicleScheduleByIDEndpoint, nil, nil, nil, httpclient.PathParam("id", id))
	return err
}

func (c *Client) CreateQuickVehicleSchedule(ctx context.Context, command models.CreateQuickVehicleScheduleCommand) (*models.CreateQuickScheduleResult, error) {
	return requestJSON[models.CreateQuickScheduleResult](ctx, c, http.MethodPost, VehicleQuickScheduleEndpoint, nil, command)
}

func (c *Client) CreateBulkQuickVehicleSchedules(ctx context.Context, command models.CreateBulkQuickVehicleScheduleCommand) (*models.CreateBulkQuickScheduleResult, error) {
	return requestJSON[models.CreateBulkQuickScheduleResult](ctx, c, http.MethodPost, VehicleBulkQuickScheduleEndpoint, nil, command)
}

func (c *Client) ApplyVehicleSchedule(ctx context.Context, command models.ApplyVehicleScheduleCommand) (*models.VehicleScheduleResponse, error) {
	return requestJSON[models.VehicleScheduleResponse](ctx, c, http.MethodPut, VehicleScheduleApplyEndpoint, nil, command)
}

func (c *Client) ApplyVehicleScheduleToAll(ctx context.Context, command models.ApplyVehicleScheduleToAllCommand) (*models.VehicleScheduleResponse, error) {
	return requestJSON[models.VehicleScheduleResponse](ctx, c, http.MethodPut, VehicleScheduleApplyToAllEndpoint, nil, command)
}

func (c *Client) RemoveVehicleSchedule(ctx context.Context, vehicleID, scheduleID string) error {
	_, err := c.request(ctx, http.MethodDelete, VehicleScheduleRemoveEndpoint, nil, nil, nil,
		httpclient.PathParam("vehicleId", vehicleID),
		httpclient.PathParam("scheduleId", scheduleID),
	)
	return err
}
