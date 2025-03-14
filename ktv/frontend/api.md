```markdown
# RESTful API Documentation

## Room

### Get All Rooms
- **URL:** `/rooms`
- **Method:** `GET`
- **Description:** Retrieve a list of all rooms.
- **Response:**
  - **200 OK:** List of rooms.

### Get Room by ID
- **URL:** `/rooms/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific room by ID.
- **Response:**
  - **200 OK:** Room details.
  - **404 Not Found:** Room not found.

### Create Room
- **URL:** `/rooms`
- **Method:** `POST`
- **Description:** Create a new room.
- **Request Body:**
  ```json
  {
    "name": "string",
    "description": "string"
  }
  ```
- **Response:**
  - **201 Created:** Room created successfully.
  - **400 Bad Request:** Invalid input.

### Update Room
- **URL:** `/rooms/{id}`
- **Method:** `PUT`
- **Description:** Update an existing room by ID.
- **Request Body:**
  ```json
  {
    "name": "string",
    "description": "string"
  }
  ```
- **Response:**
  - **200 OK:** Room updated successfully.
  - **404 Not Found:** Room not found.

### Delete Room
- **URL:** `/rooms/{id}`
- **Method:** `DELETE`
- **Description:** Delete a room by ID.
- **Response:**
  - **204 No Content:** Room deleted successfully.
  - **404 Not Found:** Room not found.

## Device

### Get All Devices
- **URL:** `/devices`
- **Method:** `GET`
- **Description:** Retrieve a list of all devices.
- **Response:**
  - **200 OK:** List of devices.

### Get Device by ID
- **URL:** `/devices/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific device by ID.
- **Response:**
  - **200 OK:** Device details.
  - **404 Not Found:** Device not found.

### Create Device
- **URL:** `/devices`
- **Method:** `POST`
- **Description:** Create a new device.
- **Request Body:**
  ```json
  {
    "name": "string",
    "type": "string",
    "roomId": "string"
  }
  ```
- **Response:**
  - **201 Created:** Device created successfully.
  - **400 Bad Request:** Invalid input.

### Update Device
- **URL:** `/devices/{id}`
- **Method:** `PUT`
- **Description:** Update an existing device by ID.
- **Request Body:**
  ```json
  {
    "name": "string",
    "type": "string",
    "roomId": "string"
  }
  ```
- **Response:**
  - **200 OK:** Device updated successfully.
  - **404 Not Found:** Device not found.

### Delete Device
- **URL:** `/devices/{id}`
- **Method:** `DELETE`
- **Description:** Delete a device by ID.
- **Response:**
  - **204 No Content:** Device deleted successfully.
  - **404 Not Found:** Device not found.

## Exception

### Get All Exceptions
- **URL:** `/exceptions`
- **Method:** `GET`
- **Description:** Retrieve a list of all exceptions.
- **Response:**
  - **200 OK:** List of exceptions.

### Get Exception by ID
- **URL:** `/exceptions/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific exception by ID.
- **Response:**
  - **200 OK:** Exception details.
  - **404 Not Found:** Exception not found.

### Create Exception
- **URL:** `/exceptions`
- **Method:** `POST`
- **Description:** Create a new exception.
- **Request Body:**
  ```json
  {
    "message": "string",
    "deviceId": "string"
  }
  ```
- **Response:**
  - **201 Created:** Exception created successfully.
  - **400 Bad Request:** Invalid input.

### Update Exception
- **URL:** `/exceptions/{id}`
- **Method:** `PUT`
- **Description:** Update an existing exception by ID.
- **Request Body:**
  ```json
  {
    "message": "string",
    "deviceId": "string"
  }
  ```
- **Response:**
  - **200 OK:** Exception updated successfully.
  - **404 Not Found:** Exception not found.

### Delete Exception
- **URL:** `/exceptions/{id}`
- **Method:** `DELETE`
- **Description:** Delete an exception by ID.
- **Response:**
  - **204 No Content:** Exception deleted successfully.
  - **404 Not Found:** Exception not found.

## ExceptionRecord

### Get All Exception Records
- **URL:** `/exception-records`
- **Method:** `GET`
- **Description:** Retrieve a list of all exception records.
- **Response:**
  - **200 OK:** List of exception records.

### Get Exception Record by ID
- **URL:** `/exception-records/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific exception record by ID.
- **Response:**
  - **200 OK:** Exception record details.
  - **404 Not Found:** Exception record not found.

### Create Exception Record
- **URL:** `/exception-records`
- **Method:** `POST`
- **Description:** Create a new exception record.
- **Request Body:**
  ```json
  {
    "exceptionId": "string",
    "timestamp": "string",
    "details": "string"
  }
  ```
- **Response:**
  - **201 Created:** Exception record created successfully.
  - **400 Bad Request:** Invalid input.

### Update Exception Record
- **URL:** `/exception-records/{id}`
- **Method:** `PUT`
- **Description:** Update an existing exception record by ID.
- **Request Body:**
  ```json
  {
    "exceptionId": "string",
    "timestamp": "string",
    "details": "string"
  }
  ```
- **Response:**
  - **200 OK:** Exception record updated successfully.
  - **404 Not Found:** Exception record not found.

### Delete Exception Record
- **URL:** `/exception-records/{id}`
- **Method:** `DELETE`
- **Description:** Delete an exception record by ID.
- **Response:**
  - **204 No Content:** Exception record deleted successfully.
  - **404 Not Found:** Exception record not found.

## NotificationTemplate

### Get All Notification Templates
- **URL:** `/notification-templates`
- **Method:** `GET`
- **Description:** Retrieve a list of all notification templates.
- **Response:**
  - **200 OK:** List of notification templates.

### Get Notification Template by ID
- **URL:** `/notification-templates/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific notification template by ID.
- **Response:**
  - **200 OK:** Notification template details.
  - **404 Not Found:** Notification template not found.

### Create Notification Template
- **URL:** `/notification-templates`
- **Method:** `POST`
- **Description:** Create a new notification template.
- **Request Body:**
  ```json
  {
    "name": "string",
    "content": "string"
  }
  ```
- **Response:**
  - **201 Created:** Notification template created successfully.
  - **400 Bad Request:** Invalid input.

### Update Notification Template
- **URL:** `/notification-templates/{id}`
- **Method:** `PUT`
- **Description:** Update an existing notification template by ID.
- **Request Body:**
  ```json
  {
    "name": "string",
    "content": "string"
  }
  ```
- **Response:**
  - **200 OK:** Notification template updated successfully.
  - **404 Not Found:** Notification template not found.

### Delete Notification Template
- **URL:** `/notification-templates/{id}`
- **Method:** `DELETE`
- **Description:** Delete a notification template by ID.
- **Response:**
  - **204 No Content:** Notification template deleted successfully.
  - **404 Not Found:** Notification template not found.

## NotificationRecord

### Get All Notification Records
- **URL:** `/notification-records`
- **Method:** `GET`
- **Description:** Retrieve a list of all notification records.
- **Response:**
  - **200 OK:** List of notification records.

### Get Notification Record by ID
- **URL:** `/notification-records/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific notification record by ID.
- **Response:**
  - **200 OK:** Notification record details.
  - **404 Not Found:** Notification record not found.

### Create Notification Record
- **URL:** `/notification-records`
- **Method:** `POST`
- **Description:** Create a new notification record.
- **Request Body:**
  ```json
  {
    "templateId": "string",
    "timestamp": "string",
    "status": "string",
    "recipient": "string"
  }
  ```
- **Response:**
  - **201 Created:** Notification record created successfully.
  - **400 Bad Request:** Invalid input.

### Update Notification Record
- **URL:** `/notification-records/{id}`
- **Method:** `PUT`
- **Description:** Update an existing notification record by ID.
- **Request Body:**
  ```json
  {
    "templateId": "string",
    "timestamp": "string",
    "status": "string",
    "recipient": "string"
  }
  ```
- **Response:**
  - **200 OK:** Notification record updated successfully.
  - **404 Not Found:** Notification record not found.

### Delete Notification Record
- **URL:** `/notification-records/{id}`
- **Method:** `DELETE`
- **Description:** Delete a notification record by ID.
- **Response:**
  - **204 No Content:** Notification record deleted successfully.
  - **404 Not Found:** Notification record not found.
```