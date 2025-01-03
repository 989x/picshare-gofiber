# PicShare GoFiber

A simple image upload API built with GoFiber.

## Features

- Upload images to organized directories
- Automatically generates `public_id` for each upload
- Supports versioned API (`api/v1`)

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

### Upload Image (Contents)

**Endpoint**:  
`POST /upload/contents`

**Headers**:  
`Content-Type: multipart/form-data`

**Body**:  
- Key: `image`  
- Value: (Upload your image file)

**Response**:
```json
{
  "message": "Image uploaded successfully",
  "public_id": "12a3b4cd",
  "path": "/images/contents/12a3b4cd/your-image.webp"
}
```

---

### Upload Image (Businesses)

**Endpoint**:  
`POST /upload/businesses`

**Headers**:  
`Content-Type: multipart/form-data`

**Body**:  
- Key: `image`  
- Value: (Upload your image file)

**Response**:
```json
{
  "message": "Image uploaded successfully",
  "public_id": "56c7d8ef",
  "path": "/images/businesses/56c7d8ef/your-image.webp"
}
```

---

### Serve Uploaded Images

**Endpoint**:  
`GET /images/<directory>/<public_id>/<filename>`

**Examples**:  
- Contents:
  ```
  http://localhost:8081/images/contents/12a3b4cd/your-image.webp
  http://localhost:8081/images/contents/34b5a6de/sample-image.jpg
  ```
- Businesses:
  ```
  http://localhost:8081/images/businesses/78c9d1ef/business-logo.png
  http://localhost:8081/images/businesses/45e2f3gh/company-banner.webp
  ```
