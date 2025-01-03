# Picshare GoFiber

A simple image upload server using Go and Fiber.

## Installation

### Prerequisites
- Install [Golang](https://golang.org/dl/).
- Ensure Nginx is installed (optional for proxy).

### Setup

```bash
# Clone the repository
mkdir picshare-gofiber && cd picshare-gofiber

# Initialize Go module
go mod init picshare-gofiber

# Install dependencies
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/fiber/v2/middleware/cors
go get github.com/gofiber/fiber/v2/middleware/logger

# Create necessary directories
sudo mkdir -p /var/www/uploads/contents
sudo mkdir -p /var/www/uploads/businesses
sudo chown -R $USER:$USER /var/www/uploads
sudo chmod -R 755 /var/www/uploads

# Run the server
go run cmd/main.go
```

### Testing

#### Upload an Image
```bash
curl -X POST -F "image=@path/to/image.jpg" http://127.0.0.1:8081/upload/contents
```

#### Serve Static Files
Visit: `http://127.0.0.1:8081/images/contents/`
