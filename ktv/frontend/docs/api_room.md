# Room API Documentation

## Room Struct Definition

```go
package main

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:包房名称" json:"name"`
	Description string `gorm:"type:varchar(255);comment:包房描述" json:"description"`
	Status      string `gorm:"type:varchar(50);comment:包房状态" json:"status"`
	DeviceCount int    `gorm:"type:int;comment:包房内设备数量" json:"device_count"`
}
```

## API Endpoints

### List Rooms

**GET /v1/rooms**

Retrieve a list of all rooms.

#### Response

- **200 OK**

```json
[
    {
        "ID": 1,
        "CreatedAt": "2023-10-01T00:00:00Z",
        "UpdatedAt": "2023-10-01T00:00:00Z",
        "DeletedAt": null,
        "name": "Room 1",
        "description": "Description 1",
        "status": "Available",
        "device_count": 10
    },
    ...
]
```

### Get Room by ID

**GET /v1/rooms/:id**

Retrieve a room by its ID.

#### Parameters

- **id** (required): The ID of the room.

#### Response

- **200 OK**

```json
{
    "ID": 1,
    "CreatedAt": "2023-10-01T00:00:00Z",
    "UpdatedAt": "2023-10-01T00:00:00Z",
    "DeletedAt": null,
    "name": "Room 1",
    "description": "Description 1",
    "status": "Available",
    "device_count": 10
}
```

- **500 Internal Server Error**

```json
{
    "error": "Error message"
}
```

### Create Room

**POST /v1/rooms**

Create a new room.

#### Request Body

```json
{
    "name": "Room Name",
    "description": "Room Description",
    "status": "Room Status",
    "device_count": 5
}
```

#### Response

- **201 Created**

```json
{
    "ID": 1,
    "CreatedAt": "2023-10-01T00:00:00Z",
    "UpdatedAt": "2023-10-01T00:00:00Z",
    "DeletedAt": null,
    "name": "Room Name",
    "description": "Room Description",
    "status": "Room Status",
    "device_count": 5
}
```

- **400 Bad Request**

```json
{
    "error": "Error message"
}
```

### Update Room

**PUT /v1/rooms/:id**

Update an existing room by its ID.

#### Parameters

- **id** (required): The ID of the room.

#### Request Body

```json
{
    "name": "Updated Room Name",
    "description": "Updated Room Description",
    "status": "Updated Room Status",
    "device_count": 10
}
```

#### Response

- **200 OK**

```json
{
    "ID": 1,
    "CreatedAt": "2023-10-01T00:00:00Z",
    "UpdatedAt": "2023-10-01T00:00:00Z",
    "DeletedAt": null,
    "name": "Updated Room Name",
    "description": "Updated Room Description",
    "status": "Updated Room Status",
    "device_count": 10
}
```

- **400 Bad Request**

```json
{
    "error": "Error message"
}
```

- **404 Not Found**

```json
{
    "error": "Room not found"
}
```

### Delete Room

**DELETE /v1/rooms/:id**

Delete a room by its ID.

#### Parameters

- **id** (required): The ID of the room.

#### Response

- **204 No Content**

- **500 Internal Server Error**

```json
{
    "error": "Room not found"
}
```