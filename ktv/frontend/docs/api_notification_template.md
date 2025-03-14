# NotificationTemplate API Documentation

## Structure Definition

```go
package main

import (
	"gorm.io/gorm"
)

type NotificationTemplate struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null;comment:通知模板名称" json:"name"`
	Content    string `gorm:"type:text;not null;comment:通知模板内容" json:"content"`
	ReceiverID uint   `gorm:"not null;comment:接收人ID" json:"receiver_id"`
}
```

## API Endpoints

### List Notification Templates

**Endpoint:** `GET /v1/notification_templates`

**Description:** Retrieve a list of all notification templates.

**Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2023-10-01T12:34:56Z",
      "UpdatedAt": "2023-10-01T12:34:56Z",
      "DeletedAt": null,
      "name": "Template Name",
      "content": "Template Content",
      "receiver_id": 123
    },
    ...
  ]
  ```

### Get Notification Template by ID

**Endpoint:** `GET /v1/notification_templates/:id`

**Description:** Retrieve a specific notification template by its ID.

**Path Parameters:**

- `id` (int, required): The ID of the notification template.

**Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "ID": 1,
    "CreatedAt": "2023-10-01T12:34:56Z",
    "UpdatedAt": "2023-10-01T12:34:56Z",
    "DeletedAt": null,
    "name": "Template Name",
    "content": "Template Content",
    "receiver_id": 123
  }
  ```

### Create Notification Template

**Endpoint:** `POST /v1/notification_templates`

**Description:** Create a new notification template.

**Request Body:**

- **Content-Type:** `application/json`
- **Body:**
  ```json
  {
    "name": "Template Name",
    "content": "Template Content",
    "receiver_id": 123
  }
  ```

**Response:**

- **Status Code:** `201 Created`
- **Body:**
  ```json
  {
    "ID": 1,
    "CreatedAt": "2023-10-01T12:34:56Z",
    "UpdatedAt": "2023-10-01T12:34:56Z",
    "DeletedAt": null,
    "name": "Template Name",
    "content": "Template Content",
    "receiver_id": 123
  }
  ```

### Update Notification Template

**Endpoint:** `PUT /v1/notification_templates/:id`

**Description:** Update an existing notification template.

**Path Parameters:**

- `id` (int, required): The ID of the notification template.

**Request Body:**

- **Content-Type:** `application/json`
- **Body:**
  ```json
  {
    "name": "Updated Template Name",
    "content": "Updated Template Content",
    "receiver_id": 456
  }
  ```

**Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "ID": 1,
    "CreatedAt": "2023-10-01T12:34:56Z",
    "UpdatedAt": "2023-10-01T12:34:56Z",
    "DeletedAt": null,
    "name": "Updated Template Name",
    "content": "Updated Template Content",
    "receiver_id": 456
  }
  ```

### Delete Notification Template

**Endpoint:** `DELETE /v1/notification_templates/:id`

**Description:** Delete a specific notification template by its ID.

**Path Parameters:**

- `id` (int, required): The ID of the notification template.

**Response:**

- **Status Code:** `204 No Content`