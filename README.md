# PicShare GoFiber

A robust image upload API built with GoFiber.

## Features

1. **Multi-image Uploads**  
   Upload multiple images under the keys `cover_image` and `body_image` in a single request.

2. **Dynamic Directory Organization**  
   - Files are automatically organized into directories based on their purpose:
     - `contents/`: For content-related images
     - `businesses/`: For business-related images
   - Each upload creates a unique subdirectory identified by a `public_id`.

3. **Versioned API**  
   The API follows a versioned structure (`api/v1`) for easier future upgrades.

4. **Environment-based Configuration**  
   The upload directory is dynamically configured using the `UPLOADS_DIR` variable from `.env`.

5. **Advanced Multipart Support**  
   Handle multiple files per key (`cover_image`, `body_image`) seamlessly in a single API call.

## Prerequisites & Installation

1. **Install Dependencies**  
  Use `go mod tidy` to install dependencies:
  ```bash
  go mod tidy
  ```

2. **Set Up Environment**  
  Create a `.env` file and configure the upload directory:
  ```bash
  echo "UPLOADS_DIR=/path/to/upload/directory" > .env

  example

  echo "UPLOADS_DIR=/var/www/uploads" > .env
  ```

3. **Run the Server**  
  Start the application:
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
