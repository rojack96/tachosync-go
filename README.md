# TachoSync Go

A typed Go client for the TachoSync API, based on OpenAPI specification `2026.10.15`.

The library automatically handles:

- authentication through the `X-Api-Key` header;
- JSON request and response encoding;
- path and query parameters;
- pagination and filters;
- binary tachograph file downloads;
- request cancellation through `context.Context`;
- empty HTTP responses and resource IDs returned in the `Location` header.

## Requirements

- Go `1.24.0` or later;
- a TachoSync API base URL;
- a valid TachoSync API key.

The module requires Go `1.24.0` or later. Go 1.24 is the minimum version because the public models use generic type aliases. The base URL must point to the API origin without the `/v1` suffix because endpoint paths are appended automatically.

## Installation

```bash
go get github.com/rojack96/tachosync-go@latest
```

The client and its public models are provided by separate packages:

```go
import (
    tachosync "github.com/rojack96/tachosync-go/client"
    "github.com/rojack96/tachosync-go/models"
)
```

The package in the `client` directory is named `tachosync`; the explicit import alias makes its purpose clear.

## Configuration

Reading the base URL and API key from environment variables is recommended:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    tachosync "github.com/rojack96/tachosync-go/client"
)

func main() {
    baseURL := os.Getenv("TACHOSYNC_BASE_URL")
    apiKey := os.Getenv("TACHOSYNC_API_KEY")
    if baseURL == "" || apiKey == "" {
        log.Fatal("TACHOSYNC_BASE_URL and TACHOSYNC_API_KEY are required")
    }

    api := tachosync.NewClient(baseURL, apiKey)

    company, err := api.GetCompanies(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Company: %s (%s)\n", company.Name, company.ID)
}
```

The API key is added to every request automatically. You do not need to build headers or endpoint URLs manually.

## Listing drivers

Search parameters and request payloads are defined in the `models` package:

```go
ctx := context.Background()

page, err := api.ListDrivers(ctx, models.ListDriversParams{
    PageNumber: 1,
    PageSize:   50,
    Statuses:   []string{"Active"},
    Contains:   "Mario",
})
if err != nil {
    log.Fatal(err)
}

for _, driver := range page.Items {
    fmt.Printf("%s - %s\n", driver.ID, valueOrEmpty(driver.CardName))
}

fmt.Printf("Page %d of %d, total: %d\n",
    page.Page,
    page.TotalPages,
    page.TotalCount,
)
```

Nullable fields and parameters that must distinguish between a zero value and an absent value are represented by pointers. A small helper can be useful when reading them:

```go
func valueOrEmpty(value *string) string {
    if value == nil {
        return ""
    }
    return *value
}
```

## Creating a driver

```go
comment := "Created by the Go integration"

driverID, err := api.CreateDriver(ctx, models.CreateDriverCommand{
    CardNumber: "I12345678901234",
    CardName:   "Mario Rossi",
    Comment:    &comment,
})
if err != nil {
    log.Fatal(err)
}

fmt.Println("New driver ID:", driverID)
```

Create methods return the resource ID from the HTTP `Location` header. If an endpoint does not provide that header, the returned string is empty.

## Retrieving and updating a vehicle

```go
vehicle, err := api.GetVehicle(ctx, "vehicle-id")
if err != nil {
    log.Fatal(err)
}

newName := "AB123CD"
if err := api.EditVehicle(ctx, models.EditVehicleCommand{
    ID:       vehicle.ID,
    Name:     &newName,
    IMEI:     vehicle.IMEI,
    Comment:  vehicle.Comment,
    Disabled: vehicle.Disabled,
}); err != nil {
    log.Fatal(err)
}
```

## Available filters

Filter endpoints use typed enums from the `models` package:

```go
filters, err := api.GetDriverFilters(ctx, models.DriverFilterParams{
    Filter:    models.DriverFilterStatuses,
    CompanyID: "company-id",
})
if err != nil {
    log.Fatal(err)
}

for _, filter := range filters {
    fmt.Println(valueOrEmpty(filter.ID), valueOrEmpty(filter.Name))
}
```

Dedicated enum types are also available for:

- `DriverFileFiltersType`;
- `VehicleFileFiltersType`;
- `VehicleFiltersType`;
- `ScheduleScopeType`;
- `DriverScope`.

## Downloading files

Download methods return the original response body as `[]byte`:

```go
data, err := api.DownloadDriverFiles(ctx, models.DownloadFilesParams{
    IDs: []string{"file-id-1", "file-id-2"},
})
if err != nil {
    log.Fatal(err)
}

if err := os.WriteFile("driver-files.bin", data, 0o600); err != nil {
    log.Fatal(err)
}
```

`DownloadVehicleFiles` works in the same way. Set `FileFormat` to one of the formats supported by the API when a specific output format is required.

## Quick schedules

The following example creates a one-time schedule for a driver:

```go
result, err := api.CreateQuickDriverSchedule(ctx, models.CreateQuickDriverScheduleCommand{
    DriverID: "driver-id",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Schedule: %s, overwritten: %t\n",
    result.ScheduleID,
    result.Overwritten,
)
```

Regular, global, bulk, apply, and remove operations are available for both driver and vehicle schedules.

## Custom HTTP client

Use `WithHTTPClient` to configure timeouts, proxies, or a custom transport:

```go
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

api := tachosync.NewClient(
    os.Getenv("TACHOSYNC_BASE_URL"),
    os.Getenv("TACHOSYNC_API_KEY"),
    tachosync.WithHTTPClient(httpClient),
)
```

The same client can be reused for multiple requests.

## Context and timeouts

Every operation accepts a `context.Context` for cancellation and deadlines:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

driver, err := api.GetDriver(ctx, "driver-id")
if err != nil {
    log.Fatal(err)
}

fmt.Println(driver.ID)
```

## Error handling

An HTTP response outside the `2xx` range produces an error containing the status code and response body:

```go
vehicle, err := api.GetVehicle(ctx, "vehicle-id")
if err != nil {
    log.Printf("TachoSync request failed: %v", err)
    return
}

fmt.Println(vehicle.ID)
```

Serialization, request construction, transport, and JSON decoding errors are returned to the caller in the same way.

## Available API methods

### Companies

- `GetCompanies`
- `GetCompany`
- `CreateCompany`
- `EditCompany`
- `DeleteCompanies`

### Company cards

- `ListCompanyCards`
- `GetCompanyCard`
- `CreateCompanyCard`
- `EditCompanyCard`
- `DeleteCompanyCard`

### Drivers and driver files

- `ListDrivers`
- `GetDriver`
- `CreateDriver`
- `EditDriver`
- `DeleteDriver`
- `GetDriverFilters`
- `ListDriverFiles`
- `GetDriverFileFilters`
- `DownloadDriverFiles`

### Vehicles and vehicle files

- `ListVehicles`
- `GetVehicle`
- `CreateVehicle`
- `EditVehicle`
- `DeleteVehicle`
- `GetVehicleDownloadsHistory`
- `GetVehicleFilters`
- `ListVehicleFiles`
- `GetVehicleFileFilters`
- `DownloadVehicleFiles`

### Driver schedules

- `ListDriverSchedules`
- `GetDriverSchedule`
- `CreateDriverSchedule`
- `EditDriverSchedule`
- `DeleteDriverSchedule`
- `ApplyDriverSchedule`
- `ApplyDriverScheduleToAll`
- `CreateQuickDriverSchedule`
- `CreateBulkQuickDriverSchedules`
- `RemoveDriverSchedule`

### Vehicle schedules

- `ListVehicleSchedules`
- `GetVehicleSchedule`
- `CreateVehicleSchedule`
- `EditVehicleSchedule`
- `DeleteVehicleSchedule`
- `ApplyVehicleSchedule`
- `ApplyVehicleScheduleToAll`
- `CreateQuickVehicleSchedule`
- `CreateBulkQuickVehicleSchedules`
- `RemoveVehicleSchedule`

## Project structure

```text
.
├── client/           # Public client, endpoints, and API methods
├── models/           # Public models, commands, responses, and parameters
└── internal/client/  # Internal HTTP transport
```

## Tests

```bash
go test ./...
```

## License

Distributed under the MIT License. See [LICENSE](LICENSE).
