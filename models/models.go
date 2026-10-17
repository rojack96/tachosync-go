package models

import "time"

type DriverFiltersType string

const (
	DriverFilterStatuses         DriverFiltersType = "Statuses"
	DriverFilterSchedules        DriverFiltersType = "Schedules"
	DriverFilterDownloadStatuses DriverFiltersType = "DownloadStatuses"
)

type DriverFileFiltersType string

const (
	DriverFileFilterDriverName DriverFileFiltersType = "DriverName"
	DriverFileFilterSchedule   DriverFileFiltersType = "Schedule"
	DriverFileFilterIMEI       DriverFileFiltersType = "Imei"
)

type VehicleFileFiltersType string

const (
	VehicleFileFilterVehicleNumber VehicleFileFiltersType = "VehicleNumber"
	VehicleFileFilterSchedule      VehicleFileFiltersType = "Schedule"
)

type VehicleFiltersType string

const (
	VehicleFilterStatuses         VehicleFiltersType = "Statuses"
	VehicleFilterSchedules        VehicleFiltersType = "Schedules"
	VehicleFilterDownloadStatuses VehicleFiltersType = "DownloadStatuses"
	VehicleFilterCreditStatuses   VehicleFiltersType = "CreditStatuses"
)

type ScheduleScopeType int32

const (
	ScheduleScopeType0 ScheduleScopeType = iota
	ScheduleScopeType1
	ScheduleScopeType2
)

type DriverScope int32

const (
	DriverScope0 DriverScope = iota
	DriverScope1
	DriverScope2
	DriverScope3
)

type CompanyHeader struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type CompanyResponse struct {
	ID                                 string            `json:"id"`
	Name                               string            `json:"name"`
	Comment                            *string           `json:"comment"`
	ParentCompanyID                    *string           `json:"parentCompanyId"`
	ParentCompanyName                  *string           `json:"parentCompanyName"`
	AutomaticDriverRegistrationEnabled bool              `json:"automaticDriverRegistrationEnabled"`
	AutoAssignGlobalScheduleToVehicle  bool              `json:"autoAssignGlobalScheduleToVehicle"`
	AutoAssignGlobalScheduleToDriver   bool              `json:"autoAssignGlobalScheduleToDriver"`
	AutoAssignVRN                      bool              `json:"autoAssignVrn"`
	AutoInitialVehicleDownload         bool              `json:"autoInitialVehicleDownload"`
	AutoDriverFileDownload             bool              `json:"autoDriverFileDownload"`
	ChildCompanies                     []CompanyResponse `json:"childCompanies"`
}

type CompanyDetailsResponse struct {
	ID                                 string                `json:"id"`
	Name                               *string               `json:"name"`
	Comment                            *string               `json:"comment"`
	ParentCompanyID                    *string               `json:"parentCompanyId"`
	ParentCompanyName                  *string               `json:"parentCompanyName"`
	EnabledVehiclesCount               int32                 `json:"enabledVehiclesCount"`
	DisabledVehiclesCount              int32                 `json:"disabledVehiclesCount"`
	AllVehiclesCount                   int32                 `json:"allVehiclesCount"`
	DriversCount                       int32                 `json:"driversCount"`
	UsersCount                         int32                 `json:"usersCount"`
	CreditsCount                       *int32                `json:"creditsCount"`
	ReservedCreditsCount               *int32                `json:"reservedCreditsCount"`
	CreditsExpirationAt                *time.Time            `json:"creditsExpirationAt"`
	AutomaticDriverRegistrationEnabled bool                  `json:"automaticDriverRegistrationEnabled"`
	AutoAssignGlobalScheduleToVehicle  bool                  `json:"autoAssignGlobalScheduleToVehicle"`
	AutoAssignGlobalScheduleToDriver   bool                  `json:"autoAssignGlobalScheduleToDriver"`
	AutoDriverFileDownload             bool                  `json:"autoDriverFileDownload"`
	AutoInitialVehicleDownload         bool                  `json:"autoInitialVehicleDownload"`
	AutoAssignVRN                      bool                  `json:"autoAssignVrn"`
	CompanyCards                       []CompanyCardResponse `json:"companyCards"`
	OverDraftAmount                    *int32                `json:"overDraftAmount"`
	OverDraftEnabled                   *bool                 `json:"overDraftEnabled"`
	ClientCode                         *string               `json:"clientCode"`
}

type CreateCompanyCommand struct {
	ParentCompanyID                    string  `json:"parentCompanyId"`
	Name                               string  `json:"name"`
	Comment                            *string `json:"comment"`
	AutomaticDriverRegistrationEnabled bool    `json:"automaticDriverRegistrationEnabled"`
	AutoAssignGlobalScheduleToVehicle  bool    `json:"autoAssignGlobalScheduleToVehicle"`
	AutoAssignGlobalScheduleToDriver   bool    `json:"autoAssignGlobalScheduleToDriver"`
	AutoDriverFileDownload             bool    `json:"autoDriverFileDownload"`
	AutoInitialVehicleDownload         bool    `json:"autoInitialVehicleDownload"`
	AutoAssignVRN                      bool    `json:"autoAssignVrn"`
}

type EditCompanyCommand struct {
	ID                                 string  `json:"id"`
	Name                               string  `json:"name"`
	Comment                            *string `json:"comment"`
	AutomaticDriverRegistrationEnabled bool    `json:"automaticDriverRegistrationEnabled"`
	AutoAssignGlobalScheduleToVehicle  bool    `json:"autoAssignGlobalScheduleToVehicle"`
	AutoAssignGlobalScheduleToDriver   bool    `json:"autoAssignGlobalScheduleToDriver"`
	AutoDriverFileDownload             bool    `json:"autoDriverFileDownload"`
	AutoInitialVehicleDownload         bool    `json:"autoInitialVehicleDownload"`
	AutoAssignVRN                      bool    `json:"autoAssignVrn"`
}

type DeleteCompaniesCommand struct {
	CompanyIDs []string `json:"companyIds"`
}

type CompanyCardResponse struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	ValidFrom      *time.Time `json:"validFrom"`
	ValidUntil     *time.Time `json:"validUntil"`
	CompanyCardID  string     `json:"companyCardId"`
	Status         string     `json:"status"`
	DisconnectedAt *time.Time `json:"disconnectedAt"`
	CardNumber     *string    `json:"cardNumber"`
	CCAVersion     *string    `json:"ccaVersion"`
	LatestVersion  *string    `json:"latestVersion"`
	CompanyID      string     `json:"companyId"`
}

type CreateCompanyCardCommand struct {
	CompanyID     string `json:"companyId"`
	Name          string `json:"name"`
	CompanyCardID string `json:"companyCardId"`
}

type EditCompanyCardCommand struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CompanyCardID string `json:"companyCardId"`
}

type FilterItem struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type Page[T any] struct {
	Page        int32 `json:"page"`
	TotalPages  int32 `json:"totalPages"`
	PageSize    int32 `json:"pageSize"`
	TotalCount  int32 `json:"totalCount"`
	HasPrevious bool  `json:"hasPrevious"`
	HasNext     bool  `json:"hasNext"`
	Items       []T   `json:"items"`
}

type CompanyCardsPage = Page[CompanyCardResponse]

type CreateQuickScheduleResult struct {
	ScheduleID  string `json:"scheduleId"`
	Overwritten bool   `json:"overwritten"`
}

type CreateDriverCommand struct {
	CompanyID  *string `json:"companyId"`
	CardNumber string  `json:"cardNumber"`
	CardName   string  `json:"cardName"`
	Comment    *string `json:"comment"`
}

type EditDriverCommand struct {
	ID         string  `json:"id"`
	CardNumber string  `json:"cardNumber"`
	CardName   string  `json:"cardName"`
	Comment    *string `json:"comment"`
}

type DriverResponse struct {
	ID                         string                `json:"id"`
	CardNumber                 *string               `json:"cardNumber"`
	CardName                   *string               `json:"cardName"`
	Comment                    *string               `json:"comment"`
	Company                    CompanyHeader         `json:"company"`
	CreatedAt                  time.Time             `json:"createdAt"`
	CreatedBy                  *string               `json:"createdBy"`
	UpdatedAt                  time.Time             `json:"updatedAt"`
	UpdatedBy                  *string               `json:"updatedBy"`
	Disabled                   bool                  `json:"disabled"`
	LastConnected              *time.Time            `json:"lastConnected"`
	LastActivity               *string               `json:"lastActivity"`
	SchedulesCount             int32                 `json:"schedulesCount"`
	LastVehicleNumber          *string               `json:"lastVehicleNumber"`
	LastVehicleID              *string               `json:"lastVehicleId"`
	LastSuccessfulDownloadTime *time.Time            `json:"lastSuccessfulDownloadTime"`
	CardExpiryDate             *time.Time            `json:"cardExpiryDate"`
	Schedules                  []ScheduleDriverItem  `json:"schedules"`
	ScheduleSummaries          []ScheduleSummaryItem `json:"scheduleSummaries"`
}

type DriverDetailedResponse struct {
	ID                         string               `json:"id"`
	CardNumber                 *string              `json:"cardNumber"`
	CardName                   *string              `json:"cardName"`
	Comment                    *string              `json:"comment"`
	Company                    CompanyHeader        `json:"company"`
	CreatedAt                  time.Time            `json:"createdAt"`
	CreatedBy                  *string              `json:"createdBy"`
	UpdatedAt                  time.Time            `json:"updatedAt"`
	UpdatedBy                  *string              `json:"updatedBy"`
	Disabled                   bool                 `json:"disabled"`
	LastConnected              *time.Time           `json:"lastConnected"`
	LastActivity               *string              `json:"lastActivity"`
	SchedulesCount             int32                `json:"schedulesCount"`
	LastVehicleNumber          *string              `json:"lastVehicleNumber"`
	LastVehicleID              *string              `json:"lastVehicleId"`
	LastSuccessfulDownloadTime *time.Time           `json:"lastSuccessfulDownloadTime"`
	CardExpiryDate             *time.Time           `json:"cardExpiryDate"`
	HasFiles                   bool                 `json:"hasFiles"`
	Schedules                  []ScheduleDriverItem `json:"schedules"`
}

type ScheduleDriverItem struct {
	ScheduleID                 string     `json:"scheduleId"`
	Name                       *string    `json:"name"`
	WasLastAttemptSuccessful   *bool      `json:"wasLastAttemptSuccessful"`
	LastSuccessfulDownloadTime *time.Time `json:"lastSuccessfulDownloadTime"`
	CurrentStatus              *string    `json:"currentStatus"`
	ScheduleStartTime          time.Time  `json:"scheduleStartTime"`
	Period                     *string    `json:"period"`
	CustomRepeat               *int32     `json:"customRepeat"`
	NextDownloadAttempt        *string    `json:"nextDownloadAttempt"`
	Quick                      bool       `json:"quick"`
	Global                     bool       `json:"global"`
	NextDueDate                time.Time  `json:"nextDueDate"`
}

type DriversPage = Page[DriverResponse]

type CreateQuickDriverScheduleCommand struct {
	CompanyID *string `json:"companyId"`
	DriverID  string  `json:"driverId"`
}

type CreateBulkQuickDriverScheduleCommand struct {
	CompanyID *string  `json:"companyId"`
	DriverIDs []string `json:"driverIds"`
}

type CreateBulkQuickDriverScheduleItemResult struct {
	DriverID        string  `json:"driverId"`
	Status          *string `json:"status"`
	FailureReason   *string `json:"failureReason"`
	QuickScheduleID *string `json:"quickScheduleId"`
}

type CreateBulkQuickDriverScheduleResult struct {
	Items []CreateBulkQuickDriverScheduleItemResult `json:"items"`
}

type FileDriver struct {
	DriverID   string  `json:"driverId"`
	DriverName *string `json:"driverName"`
	CardNumber *string `json:"cardNumber"`
}

type DriverFileSchedule struct {
	ID        string    `json:"id"`
	Name      *string   `json:"name"`
	Created   time.Time `json:"created"`
	Repeat    *string   `json:"repeat"`
	FileTypes []string  `json:"fileTypes"`
}

type DriverFileResponse struct {
	ID           string             `json:"id"`
	DownloadTime time.Time          `json:"downloadTime"`
	FileName     *string            `json:"fileName"`
	VehicleIMEI  *string            `json:"vehicleImei"`
	Driver       FileDriver         `json:"driver"`
	Company      CompanyHeader      `json:"company"`
	Schedule     DriverFileSchedule `json:"schedule"`
	SourceName   *string            `json:"sourceName"`
}

type DriverFilesPage = Page[DriverFileResponse]

type DriverHeader struct {
	ID         string  `json:"id"`
	DriverName *string `json:"driverName"`
}

type DriverScheduleItem struct {
	ID                 string    `json:"id"`
	Name               *string   `json:"name"`
	FileTypes          []string  `json:"fileTypes"`
	StartDate          time.Time `json:"startDate"`
	Repeat             *string   `json:"repeat"`
	CustomRepeat       *int32    `json:"customRepeat"`
	Global             bool      `json:"global"`
	Quick              bool      `json:"quick"`
	NextDate           time.Time `json:"nextDate"`
	AttachedToRoot     bool      `json:"attachedToRoot"`
	InheritanceEnabled bool      `json:"inheritanceEnabled"`
}

type DriverScheduleResponse struct {
	ID                 string         `json:"id"`
	Name               *string        `json:"name"`
	FileTypes          []string       `json:"fileTypes"`
	StartDate          time.Time      `json:"startDate"`
	Repeat             *string        `json:"repeat"`
	CustomRepeat       *int32         `json:"customRepeat"`
	Company            CompanyHeader  `json:"company"`
	CreatedAt          time.Time      `json:"createdAt"`
	CreatedBy          *string        `json:"createdBy"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	UpdatedBy          *string        `json:"updatedBy"`
	Deleted            bool           `json:"deleted"`
	NextDate           time.Time      `json:"nextDate"`
	AttachedToRoot     bool           `json:"attachedToRoot"`
	InheritanceEnabled bool           `json:"inheritanceEnabled"`
	Drivers            []DriverHeader `json:"drivers"`
}

type CreateDriverScheduleCommand struct {
	CompanyID          *string   `json:"companyId"`
	Name               string    `json:"name"`
	StartDate          time.Time `json:"startDate"`
	CustomRepeat       *int32    `json:"customRepeat"`
	Repeat             string    `json:"repeat"`
	AttachedToRoot     bool      `json:"attachedToRoot"`
	InheritanceEnabled bool      `json:"inheritanceEnabled"`
	Global             bool      `json:"global"`
}

type EditDriverScheduleCommand struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	StartDate          time.Time `json:"startDate"`
	Repeat             string    `json:"repeat"`
	CustomRepeat       *int32    `json:"customRepeat"`
	AttachedToRoot     bool      `json:"attachedToRoot"`
	InheritanceEnabled bool      `json:"inheritanceEnabled"`
}

type ApplyDriverScheduleCommand struct {
	CompanyID  *string  `json:"companyId"`
	ScheduleID string   `json:"scheduleId"`
	DriverIDs  []string `json:"driverIds"`
}

type ApplyDriverScheduleToAllCommand struct {
	CompanyID  *string `json:"companyId"`
	ScheduleID string  `json:"scheduleId"`
}

type ScheduleSummaryItem struct {
	Name  *string           `json:"name"`
	Scope ScheduleScopeType `json:"scope"`
}

type VehicleHeader struct {
	ID            string  `json:"id"`
	VehicleNumber *string `json:"vehicleNumber"`
}

type VehicleScheduleActivities struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type VehicleScheduleItem struct {
	ID                 string    `json:"id"`
	Name               *string   `json:"name"`
	FileTypes          []string  `json:"fileTypes"`
	StartDate          time.Time `json:"startDate"`
	Repeat             *string   `json:"repeat"`
	CustomRepeat       *int32    `json:"customRepeat"`
	Global             bool      `json:"global"`
	Quick              bool      `json:"quick"`
	NextDate           time.Time `json:"nextDate"`
	AttachedToRoot     bool      `json:"attachedToRoot"`
	InheritanceEnabled bool      `json:"inheritanceEnabled"`
}

type VehicleScheduleResponse struct {
	ID                 string                    `json:"id"`
	Name               *string                   `json:"name"`
	FileTypes          []string                  `json:"fileTypes"`
	StartDate          time.Time                 `json:"startDate"`
	Repeat             *string                   `json:"repeat"`
	CustomRepeat       *int32                    `json:"customRepeat"`
	Company            CompanyHeader             `json:"company"`
	CreatedAt          time.Time                 `json:"createdAt"`
	CreatedBy          *string                   `json:"createdBy"`
	UpdatedAt          time.Time                 `json:"updatedAt"`
	UpdatedBy          *string                   `json:"updatedBy"`
	Deleted            bool                      `json:"deleted"`
	NextDate           time.Time                 `json:"nextDate"`
	AttachedToRoot     bool                      `json:"attachedToRoot"`
	InheritanceEnabled bool                      `json:"inheritanceEnabled"`
	ActivitiesDetails  VehicleScheduleActivities `json:"activitiesDetails"`
	Vehicles           []VehicleHeader           `json:"vehicles"`
}

type CreateVehicleScheduleCommand struct {
	CompanyID          *string                   `json:"companyId"`
	Name               string                    `json:"name"`
	StartDate          time.Time                 `json:"startDate"`
	CustomRepeat       *int32                    `json:"customRepeat"`
	Repeat             string                    `json:"repeat"`
	AttachedToRoot     bool                      `json:"attachedToRoot"`
	InheritanceEnabled bool                      `json:"inheritanceEnabled"`
	Global             bool                      `json:"global"`
	Activities         VehicleScheduleActivities `json:"activities"`
	FileTypes          []string                  `json:"fileTypes"`
}

type EditVehicleScheduleCommand struct {
	ID                 string                    `json:"id"`
	Name               string                    `json:"name"`
	StartDate          time.Time                 `json:"startDate"`
	Repeat             string                    `json:"repeat"`
	CustomRepeat       *int32                    `json:"customRepeat"`
	AttachedToRoot     bool                      `json:"attachedToRoot"`
	InheritanceEnabled bool                      `json:"inheritanceEnabled"`
	Activities         VehicleScheduleActivities `json:"activities"`
	FileTypes          []string                  `json:"fileTypes"`
}

type ApplyVehicleScheduleCommand struct {
	CompanyID  *string  `json:"companyId"`
	ScheduleID string   `json:"scheduleId"`
	VehicleIDs []string `json:"vehicleIds"`
}

type ApplyVehicleScheduleToAllCommand struct {
	CompanyID  *string `json:"companyId"`
	ScheduleID string  `json:"scheduleId"`
}

type CreateQuickVehicleScheduleCommand struct {
	CompanyID   *string                   `json:"companyId"`
	VehicleID   string                    `json:"vehicleId"`
	FileTypes   []string                  `json:"fileTypes"`
	Activities  VehicleScheduleActivities `json:"activities"`
	DriverScope DriverScope               `json:"driverScope"`
}

type CreateBulkQuickVehicleScheduleCommand struct {
	CompanyID   *string                   `json:"companyId"`
	VehicleIDs  []string                  `json:"vehicleIds"`
	FileTypes   []string                  `json:"fileTypes"`
	Activities  VehicleScheduleActivities `json:"activities"`
	DriverScope DriverScope               `json:"driverScope"`
}

type CreateBulkQuickScheduleItemResult struct {
	VehicleID       string  `json:"vehicleId"`
	Status          *string `json:"status"`
	FailureReason   *string `json:"failureReason"`
	QuickScheduleID *string `json:"quickScheduleId"`
}

type CreateBulkQuickScheduleResult struct {
	Items []CreateBulkQuickScheduleItemResult `json:"items"`
}

type CreateVehicleCommand struct {
	CompanyID         *string `json:"companyId"`
	IMEI              *string `json:"imei"`
	Name              *string `json:"name"`
	Comment           *string `json:"comment"`
	Disable           *bool   `json:"disable"`
	ForceReassignIMEI bool    `json:"forceReassignImei"`
}

type EditVehicleCommand struct {
	ID                            string  `json:"id"`
	IMEI                          *string `json:"imei"`
	Name                          *string `json:"name"`
	Comment                       *string `json:"comment"`
	ForceReassignIMEI             bool    `json:"forceReassignImei"`
	ConfirmIMEIConflictResolution bool    `json:"confirmImeiConflictResolution"`
	Disabled                      bool    `json:"disabled"`
}

type ScheduleVehicleItem struct {
	ScheduleID                 string      `json:"scheduleId"`
	Name                       *string     `json:"name"`
	WasLastAttemptSuccessful   *bool       `json:"wasLastAttemptSuccessful"`
	LastSuccessfulDownloadTime *time.Time  `json:"lastSuccessfulDownloadTime"`
	CurrentStatus              *string     `json:"currentStatus"`
	ScheduleStartTime          time.Time   `json:"scheduleStartTime"`
	ActivitiesNextStartDate    *time.Time  `json:"activitiesNextStartDate"`
	ActivitiesNextEndDate      *time.Time  `json:"activitiesNextEndDate"`
	Repeat                     *string     `json:"repeat"`
	CustomRepeat               *int32      `json:"customRepeat"`
	NextDownloadAttempt        *string     `json:"nextDownloadAttempt"`
	FileTypes                  []string    `json:"fileTypes"`
	AttemptDownloadTime        *time.Time  `json:"attemptDownloadTime"`
	Quick                      bool        `json:"quick"`
	Global                     bool        `json:"global"`
	NextDueDate                time.Time   `json:"nextDueDate"`
	DriverCardNumber1          *string     `json:"driverCardNumber1"`
	DriverCardNumber2          *string     `json:"driverCardNumber2"`
	DriverScope                DriverScope `json:"driverScope"`
}

type VehicleResponse struct {
	ID                         string                `json:"id"`
	IMEI                       *string               `json:"imei"`
	Number                     *string               `json:"number"`
	Comment                    *string               `json:"comment"`
	Company                    CompanyHeader         `json:"company"`
	LastConnected              *time.Time            `json:"lastConnected"`
	Status                     *string               `json:"status"`
	CreatedAt                  time.Time             `json:"createdAt"`
	CreatedBy                  *string               `json:"createdBy"`
	UpdatedAt                  time.Time             `json:"updatedAt"`
	UpdatedBy                  *string               `json:"updatedBy"`
	Disabled                   bool                  `json:"disabled"`
	SchedulesCount             int32                 `json:"schedulesCount"`
	LastSuccessfulDownloadTime *time.Time            `json:"lastSuccessfulDownloadTime"`
	ScheduleSummaries          []ScheduleSummaryItem `json:"scheduleSummaries"`
	Schedules                  []ScheduleVehicleItem `json:"schedules"`
	CreditStatus               *string               `json:"creditStatus"`
}

type VehicleDetailedResponse struct {
	ID                         string                `json:"id"`
	IMEI                       *string               `json:"imei"`
	Number                     *string               `json:"number"`
	Comment                    *string               `json:"comment"`
	Company                    CompanyHeader         `json:"company"`
	LastConnected              *time.Time            `json:"lastConnected"`
	Status                     *string               `json:"status"`
	CreatedAt                  time.Time             `json:"createdAt"`
	CreatedBy                  *string               `json:"createdBy"`
	UpdatedAt                  time.Time             `json:"updatedAt"`
	UpdatedBy                  *string               `json:"updatedBy"`
	Disabled                   bool                  `json:"disabled"`
	SchedulesCount             int32                 `json:"schedulesCount"`
	LastSuccessfulDownloadTime *time.Time            `json:"lastSuccessfulDownloadTime"`
	HasFiles                   bool                  `json:"hasFiles"`
	IsAlreadyCharged           bool                  `json:"isAlreadyCharged"`
	Schedules                  []ScheduleVehicleItem `json:"schedules"`
}

type VehiclesPage = Page[VehicleResponse]

type FileVehicle struct {
	ID     string  `json:"id"`
	Number *string `json:"number"`
	IMEI   *string `json:"imei"`
}

type VehicleFileSchedule struct {
	ID             string     `json:"id"`
	Name           *string    `json:"name"`
	Created        time.Time  `json:"created"`
	Repeat         *string    `json:"repeat"`
	FileTypes      []string   `json:"fileTypes"`
	ActivitiesFrom *time.Time `json:"activitiesFrom"`
	ActivitiesTo   *time.Time `json:"activitiesTo"`
}

type VehicleFileResponse struct {
	ID           string              `json:"id"`
	FileName     *string             `json:"fileName"`
	Vehicle      FileVehicle         `json:"vehicle"`
	Schedule     VehicleFileSchedule `json:"schedule"`
	S3FileKey    *string             `json:"s3FileKey"`
	DownloadTime time.Time           `json:"downloadTime"`
	Company      CompanyHeader       `json:"company"`
	SourceName   *string             `json:"sourceName"`
}

type VehicleFilesPage = Page[VehicleFileResponse]

type DownloadHistoryResponse struct {
	Date            time.Time `json:"date"`
	Stage           *string   `json:"stage"`
	Status          *string   `json:"status"`
	Details         *string   `json:"details"`
	Driver          *string   `json:"driver"`
	VehicleSchedule *string   `json:"vehicleSchedule"`
	DriverSchedule  *string   `json:"driverSchedule"`
	ScheduleID      *string   `json:"scheduleId"`
	FailedStep      bool      `json:"failedStep"`
}

type APIError struct {
	StatusCode HTTPStatusCode `json:"statusCode"`
	Message    *string        `json:"message"`
	Entries    []ErrorEntry   `json:"entries"`
}

type ErrorEntry struct {
	Message *string `json:"message"`
	Field   *string `json:"field"`
	Rows    *string `json:"rows"`
}

type HTTPStatusCode int32

type ProblemDetails struct {
	Type     *string `json:"type"`
	Title    *string `json:"title"`
	Status   *int32  `json:"status"`
	Detail   *string `json:"detail"`
	Instance *string `json:"instance"`
}
