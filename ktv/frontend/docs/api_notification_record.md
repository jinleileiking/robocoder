# NotificationRecord API 文档

## 结构体定义

```go
package main

import (
	"gorm.io/gorm"
	"time"
)

type NotificationRecord struct {
	gorm.Model
	Recipient       string    `gorm:"type:varchar(100);not null;comment:接收人" json:"recipient"`                        // 接收人
	ContentTemplate string    `gorm:"type:text;not null;comment:通知内容模板" json:"content_template"`                     // 通知内容模板
	RoomName        string    `gorm:"type:varchar(100);not null;comment:包房名称" json:"room_name"`                        // 包房名称
	DeviceName      string    `gorm:"type:varchar(100);not null;comment:设备名称" json:"device_name"`                      // 设备名称
	Link            string    `gorm:"type:varchar(255);comment:跳转链接" json:"link"`                                     // 跳转链接
	SentAt          time.Time `gorm:"type:datetime;not null;comment:发送时间" json:"sent_at"`                              // 发送时间
	Status          string    `gorm:"type:varchar(50);not null;comment:通知状态" json:"status"`                            // 通知状态
}
```

## 接口列表

### 获取通知记录列表

**URL:** `/v1/notification_records`

**方法:** `GET`

**响应:**

- `200 OK` - 成功获取通知记录列表
- `500 Internal Server Error` - 服务器内部错误

**响应示例:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-10-01T12:00:00Z",
    "UpdatedAt": "2023-10-01T12:00:00Z",
    "DeletedAt": null,
    "recipient": "John Doe",
    "content_template": "Template 1",
    "room_name": "Room 1",
    "device_name": "Device 1",
    "link": "http://example.com",
    "sent_at": "2023-10-01T12:00:00Z",
    "status": "Sent"
  }
]
```

### 获取单个通知记录

**URL:** `/v1/notification_records/:id`

**方法:** `GET`

**参数:**

- `id` - 通知记录的ID

**响应:**

- `200 OK` - 成功获取通知记录
- `500 Internal Server Error` - 服务器内部错误
- `404 Not Found` - 通知记录未找到

**响应示例:**

```json
{
  "ID": 1,
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z",
  "DeletedAt": null,
  "recipient": "John Doe",
  "content_template": "Template 1",
  "room_name": "Room 1",
  "device_name": "Device 1",
  "link": "http://example.com",
  "sent_at": "2023-10-01T12:00:00Z",
  "status": "Sent"
}
```

### 创建通知记录

**URL:** `/v1/notification_records`

**方法:** `POST`

**请求体:**

```json
{
  "recipient": "John Doe",
  "content_template": "Template 1",
  "room_name": "Room 1",
  "device_name": "Device 1",
  "link": "http://example.com",
  "sent_at": "2023-10-01T12:00:00Z",
  "status": "Sent"
}
```

**响应:**

- `201 Created` - 成功创建通知记录
- `400 Bad Request` - 请求体格式错误
- `500 Internal Server Error` - 服务器内部错误

**响应示例:**

```json
{
  "ID": 1,
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z",
  "DeletedAt": null,
  "recipient": "John Doe",
  "content_template": "Template 1",
  "room_name": "Room 1",
  "device_name": "Device 1",
  "link": "http://example.com",
  "sent_at": "2023-10-01T12:00:00Z",
  "status": "Sent"
}
```

### 更新通知记录

**URL:** `/v1/notification_records/:id`

**方法:** `PUT`

**参数:**

- `id` - 通知记录的ID

**请求体:**

```json
{
  "recipient": "John Doe",
  "content_template": "Template 1",
  "room_name": "Room 1",
  "device_name": "Device 1",
  "link": "http://example.com",
  "sent_at": "2023-10-01T12:00:00Z",
  "status": "Sent"
}
```

**响应:**

- `200 OK` - 成功更新通知记录
- `400 Bad Request` - 请求体格式错误
- `404 Not Found` - 通知记录未找到
- `500 Internal Server Error` - 服务器内部错误

**响应示例:**

```json
{
  "ID": 1,
  "CreatedAt": "2023-10-01T12:00:00Z",
  "UpdatedAt": "2023-10-01T12:00:00Z",
  "DeletedAt": null,
  "recipient": "John Doe",
  "content_template": "Template 1",
  "room_name": "Room 1",
  "device_name": "Device 1",
  "link": "http://example.com",
  "sent_at": "2023-10-01T12:00:00Z",
  "status": "Sent"
}
```

### 删除通知记录

**URL:** `/v1/notification_records/:id`

**方法:** `DELETE`

**参数:**

- `id` - 通知记录的ID

**响应:**

- `204 No Content` - 成功删除通知记录
- `404 Not Found` - 通知记录未找到
- `500 Internal Server Error` - 服务器内部错误