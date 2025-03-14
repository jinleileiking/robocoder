# Device API Documentation

## Device Structure

```go
package main

import (
	"gorm.io/gorm"
	"time"
)

type Device struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"name"`
	Type          string    `gorm:"type:varchar(50);not null;comment:设备种类" json:"type"`
	HealthStatus  string    `gorm:"type:varchar(50);not null;comment:设备健康状态" json:"health_status"`
	ModelType     string    `gorm:"type:varchar(100);not null;comment:设备型号" json:"model_type"`
	Uptime        int64     `gorm:"not null;comment:开机时长（秒）" json:"uptime"`
	CPUUsage      float64   `gorm:"not null;comment:CPU占用百分比" json:"cpu_usage"`
	MemoryUsage   float64   `gorm:"not null;comment:内存占用百分比" json:"memory_usage"`
	RoomID        uint      `gorm:"not null;comment:所属包房ID" json:"room_id"`
	LastCheckedAt time.Time `gorm:"not null;comment:最后检查时间" json:"last_checked_at"`
}
```

## API Endpoints

### Get Device List

**GET** `/v1/devices`

#### Response

- **200 OK**: Returns a list of devices.

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-01-01T00:00:00Z",
    "UpdatedAt": "2023-01-01T00:00:00Z",
    "DeletedAt": null,
    "name": "Device1",
    "type": "Type1",
    "health_status": "Good",
    "model_type": "Model1",
    "uptime": 3600,
    "cpu_usage": 50.5,
    "memory_usage": 70.5,
    "room_id": 1,
    "last_checked_at": "2023-01-01T00:00:00Z"
  }
]
```

### Get Device by ID

**GET** `/v1/devices/:id`

#### Parameters

- `id` (path): The ID of the device.

#### Response

- **200 OK**: Returns the device with the specified ID.
- **500 Internal Server Error**: If the ID is not a valid integer or if there is an error retrieving the device.

```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:00:00Z",
  "DeletedAt": null,
  "name": "Device1",
  "type": "Type1",
  "health_status": "Good",
  "model_type": "Model1",
  "uptime": 3600,
  "cpu_usage": 50.5,
  "memory_usage": 70.5,
  "room_id": 1,
  "last_checked_at": "2023-01-01T00:00:00Z"
}
```

### Create Device

**POST** `/v1/devices`

#### Request Body

- The device object to create.

```json
{
  "name": "Device1",
  "type": "Type1",
  "health_status": "Good",
  "model_type": "Model1",
  "uptime": 3600,
  "cpu_usage": 50.5,
  "memory_usage": 70.5,
  "room_id": 1,
  "last_checked_at": "2023-01-01T00:00:00Z"
}
```

#### Response

- **201 Created**: Returns the created device.
- **400 Bad Request**: If the request body is invalid.
- **500 Internal Server Error**: If there is an error creating the device.

```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:00:00Z",
  "DeletedAt": null,
  "name": "Device1",
  "type": "Type1",
  "health_status": "Good",
  "model_type": "Model1",
  "uptime": 3600,
  "cpu_usage": 50.5,
  "memory_usage": 70.5,
  "room_id": 1,
  "last_checked_at": "2023-01-01T00:00:00Z"
}
```

### Update Device

**PUT** `/v1/devices/:id`

#### Parameters

- `id` (path): The ID of the device.

#### Request Body

- The updated device object.

```json
{
  "name": "UpdatedDevice",
  "type": "UpdatedType",
  "health_status": "UpdatedStatus",
  "model_type": "UpdatedModel",
  "uptime": 7200,
  "cpu_usage": 60.5,
  "memory_usage": 80.5,
  "room_id": 2,
  "last_checked_at": "2023-01-02T00:00:00Z"
}
```

#### Response

- **200 OK**: Returns the updated device.
- **400 Bad Request**: If the request body is invalid.
- **404 Not Found**: If the device with the specified ID is not found.
- **500 Internal Server Error**: If there is an error updating the device.

```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:00:00Z",
  "DeletedAt": null,
  "name": "UpdatedDevice",
  "type": "UpdatedType",
  "health_status": "UpdatedStatus",
  "model_type": "UpdatedModel",
  "uptime": 7200,
  "cpu_usage": 60.5,
  "memory_usage": 80.5,
  "room_id": 2,
  "last_checked_at": "2023-01-02T00:00:00Z"
}
```

### Delete Device

**DELETE** `/v1/devices/:id`

#### Parameters

- `id` (path): The ID of the device.

#### Response

- **204 No Content**: If the device is successfully deleted.
- **500 Internal Server Error**: If there is an error deleting the device.

```json
// No content
```