# PicShare GoFiber

A simple image upload API built with GoFiber.

## Features

- Upload multiple images under `cover_image` and `body_image` keys
- Automatically organizes files into directories by `public_id`
- Supports versioned API (`api/v1`)
- **Directories**: Files are stored under `/var/www/uploads`:
  - `contents/` for content-related images
  - `businesses/` for business-related images
- **Dynamic Directory Structure**: Each upload creates a new subdirectory identified by a unique `public_id`
- **Multipart Support**: Supports multiple files per key in a single request

## Prerequisites

- Go 1.19+, Fiber v2
- A directory `/var/www/uploads` with proper write permissions

## Installation

1. Install dependencies:

   ```bash
   go mod tidy
   ```

2. Run the server:

   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Base URL

```
http://localhost:8081/api/v1
```

### Upload Images (Contents)

**Endpoint**:  
`POST /upload/contents`

**Headers**:  
`Content-Type: multipart/form-data`

**Body**:  
- Key: `cover_image`  
  - Value: (Upload one or more image files for cover images)
- Key: `body_image`  
  - Value: (Upload one or more image files for body images)

**Response**:
```json
{
  "message": "Images uploaded successfully",
  "public_id": "12a3b4cd",
  "cover_image": [
    "/images/contents/12a3b4cd/cover1.jpg",
    "/images/contents/12a3b4cd/cover2.png"
  ],
  "body_image": [
    "/images/contents/12a3b4cd/body1.jpg",
    "/images/contents/12a3b4cd/body2.webp"
  ]
}
```

---

### Upload Images (Businesses)

**Endpoint**:  
`POST /upload/businesses`

**Headers**:  
`Content-Type: multipart/form-data`

**Body**:  
- Key: `cover_image`  
  - Value: (Upload one or more image files for cover images)
- Key: `body_image`  
  - Value: (Upload one or more image files for body images)

**Response**:
```json
{
  "message": "Images uploaded successfully",
  "public_id": "56c7d8ef",
  "cover_image": [
    "/images/businesses/56c7d8ef/cover1.jpg",
    "/images/businesses/56c7d8ef/cover2.png"
  ],
  "body_image": [
    "/images/businesses/56c7d8ef/body1.jpg",
    "/images/businesses/56c7d8ef/body2.webp"
  ]
}
```

---

### Serve Uploaded Images

**Endpoint**:  
`GET /images/<directory>/<public_id>/<filename>`

**Examples**:  
- Contents:
  ```
  http://localhost:8081/images/contents/12a3b4cd/cover1.jpg
  http://localhost:8081/images/contents/34b5a6de/body1.webp
  ```
- Businesses:
  ```
  http://localhost:8081/images/businesses/78c9d1ef/cover2.png
  http://localhost:8081/images/businesses/45e2f3gh/body2.jpg
  ```
