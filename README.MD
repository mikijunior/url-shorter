# URL Shortener Service

This is a simple URL shortening service written in Go. It allows you to shorten long URLs and redirect to the original URLs using the shortened version.

## Features

- Shorten URLs via a REST API
- Redirect to original URLs using shortened URLs
- In-memory storage for simplicity (can be extended to use a database)

## Setup and Running the Server

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/url_shortener.git
    cd url_shortener
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Run the server:

    ```sh
    go run cmd/main.go
    ```

The server will start on port 8080.

## API Endpoints

### Shorten a URL

- **URL**: `/shorten`
- **Method**: `POST`
- **Request Body**: JSON object with the original URL
    ```json
    {
        "url": "https://example.com"
    }
    ```
- **Response**: JSON object with the shortened URL
    ```json
    {
        "short_url": "shortURL"
    }
    ```

### Redirect to Original URL

- **URL**: `/{shortURL}`
- **Method**: `GET`
- **Response**: Redirects to the original URL

