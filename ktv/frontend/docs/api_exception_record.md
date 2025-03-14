# ExceptionRecord API Documentation

## ExceptionRecord Structure

```go
package main

import (
    "time"
    "gorm.io/gorm"
)

// ExceptionRecord represents a record of an exception that occurred in a device.
type ExceptionRecord struct {
    gorm.Model
    RoomName       string    `gorm:"type:varchar(100);not null;comment:包房名称" json:"room_name"`
    DeviceName     string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"device_name"`
    ExceptionTime  time.Time `gorm:"type:datetime;not null;comment:异常时间" json:"exception_time"`
    Description    string    `gorm:"type:text;comment:异常描述" json:"description"`
    Handler        string    `gorm:"type:varchar(100);comment:处理人" json:"handler"`
    HandledTime    time.Time `gorm:"type:datetime;comment:处理时间" json:"handled_time"`
    HandleResult   string    `gorm:"type:text;comment:处理结果" json:"handle_result"`
}
```

## API Endpoints

### List Exception Records

**GET** `/v1/exception-records`

- **Description**: Retrieve a list of all exception records.
- **Response**:
  - `200 OK`: Returns a list of `ExceptionRecord` objects.

### Get Exception Record by ID

**GET** `/v1/exception-records/:id`

- **Description**: Retrieve a specific exception record by its ID.
- **Parameters**:
  - `id` (path): The ID of the exception record.
- **Response**:
  - `200 OK`: Returns the `ExceptionRecord` object.
  - `500 Internal Server Error`: If the ID is not a valid integer or if there's an error retrieving the record.

### Create Exception Record

**POST** `/v1/exception-records`

- **Description**: Create a new exception record.
- **Request Body**:
  - `ExceptionRecord` object.
- **Response**:
  - `201 Created`: Returns the created `ExceptionRecord` object.
  - `400 Bad Request`: If the request body is not valid.
  - `500 Internal Server Error`: If there's an error creating the record.

### Update Exception Record

**PUT** `/v1/exception-records/:id`

- **Description**: Update an existing exception record by its ID.
- **Parameters**:
  - `id` (path): The ID of the exception record.
- **Request Body**:
  - `ExceptionRecord` object.
- **Response**:
  - `200 OK`: Returns the updated `ExceptionRecord` object.
  - `400 Bad Request`: If the request body is not valid.
  - `404 Not Found`: If the exception record with the given ID does not exist.
  - `500 Internal Server Error`: If there's an error updating the record.

### Delete Exception Record

**DELETE** `/v1/exception-records/:id`

- **Description**: Delete an existing exception record by its ID.
- **Parameters**:
  - `id` (path): The ID of the exception record.
- **Response**:
  - `204 No Content`: If the deletion is successful.
  - `500 Internal Server Error`: If the ID is not a valid integer or if there's an error deleting the record.