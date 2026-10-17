package models

import "time"

type ListCompanyCardsParams struct {
	Status        string     `url:"Status"`
	ValidFrom     *time.Time `url:"ValidFrom"`
	ExpiresBefore *time.Time `url:"ExpiresBefore"`
	PageNumber    int32      `url:"PageNumber"`
	PageSize      int32      `url:"PageSize"`
	Contains      string     `url:"Contains"`
	CompanyID     string     `url:"CompanyId"`
}

type ListDriverFilesParams struct {
	DriverIDs              []string   `url:"DriverIds"`
	VehicleIMEIs           []string   `url:"VehicleImeis"`
	NoVehicleIMEIAssigned  *bool      `url:"NoVehicleImeiAssigned"`
	OrderBy                string     `url:"OrderBy"`
	Descending             *bool      `url:"Descending"`
	From                   *time.Time `url:"From"`
	To                     *time.Time `url:"To"`
	AllCompanies           *bool      `url:"AllCompanies"`
	CompanyID              string     `url:"CompanyId"`
	ScheduleNames          []string   `url:"ScheduleNames"`
	NoSchedulesAssigned    *bool      `url:"NoSchedulesAssigned"`
	QuickScheduleAssigned  *bool      `url:"QuickScheduleAssigned"`
	GlobalScheduleAssigned *bool      `url:"GlobalScheduleAssigned"`
	FileFormat             string     `url:"FileFormat"`
	PageNumber             int32      `url:"PageNumber"`
	PageSize               int32      `url:"PageSize"`
	Contains               string     `url:"Contains"`
}

type ListVehicleFilesParams struct {
	VehicleIDs             []string   `url:"VehicleIds"`
	OrderBy                string     `url:"OrderBy"`
	Descending             *bool      `url:"Descending"`
	From                   *time.Time `url:"From"`
	To                     *time.Time `url:"To"`
	AllCompanies           *bool      `url:"AllCompanies"`
	CompanyID              string     `url:"CompanyId"`
	ScheduleNames          []string   `url:"ScheduleNames"`
	NoSchedulesAssigned    *bool      `url:"NoSchedulesAssigned"`
	QuickScheduleAssigned  *bool      `url:"QuickScheduleAssigned"`
	GlobalScheduleAssigned *bool      `url:"GlobalScheduleAssigned"`
	FileFormat             string     `url:"FileFormat"`
	PageNumber             int32      `url:"PageNumber"`
	PageSize               int32      `url:"PageSize"`
	Contains               string     `url:"Contains"`
}

type FileFilterParams[T ~string] struct {
	Filter    T      `url:"Filter"`
	CompanyID string `url:"CompanyId"`
}

type DownloadFilesParams struct {
	FileFormat string   `url:"FileFormat"`
	IDs        []string `url:"Ids"`
}

type ListDriversParams struct {
	ScheduleIDs         []string `url:"ScheduleIds"`
	NoSchedulesAssigned *bool    `url:"NoSchedulesAssigned"`
	Statuses            []string `url:"Statuses"`
	OrderBy             string   `url:"OrderBy"`
	DownloadStatuses    []string `url:"DownloadStatuses"`
	Descending          *bool    `url:"Descending"`
	ScheduleDetails     *bool    `url:"ScheduleDetails"`
	PageNumber          int32    `url:"PageNumber"`
	PageSize            int32    `url:"PageSize"`
	Contains            string   `url:"Contains"`
	CompanyID           string   `url:"CompanyId"`
}

type DriverFilterParams struct {
	Filter    DriverFiltersType `url:"Filter"`
	CompanyID string            `url:"CompanyId"`
}

type ListVehiclesParams struct {
	ScheduleIDs         []string `url:"ScheduleIds"`
	NoSchedulesAssigned *bool    `url:"NoSchedulesAssigned"`
	Statuses            []string `url:"Statuses"`
	OrderBy             string   `url:"OrderBy"`
	CreditStatuses      []string `url:"CreditStatuses"`
	NoDeviceIMEI        *bool    `url:"NoDeviceImei"`
	DeactivatedVehicles *bool    `url:"DeactivatedVehicles"`
	DownloadStatuses    []string `url:"DownloadStatuses"`
	Descending          *bool    `url:"Descending"`
	ScheduleDetails     *bool    `url:"ScheduleDetails"`
	PageNumber          int32    `url:"PageNumber"`
	PageSize            int32    `url:"PageSize"`
	Contains            string   `url:"Contains"`
	CompanyID           string   `url:"CompanyId"`
}

type VehicleFilterParams struct {
	Filter    VehicleFiltersType `url:"Filter"`
	CompanyID string             `url:"CompanyId"`
}

type ListSchedulesParams struct {
	CompanyID string `url:"CompanyId"`
}
