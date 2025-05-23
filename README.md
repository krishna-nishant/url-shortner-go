# URL Shortener

A complete URL shortener application built with Go, featuring a modern UI.

## Features

- Create shortened URLs from long URLs
- Track click statistics for each shortened URL
- Modern, responsive UI with smooth animations
- Copy shortened URLs to clipboard with one click
- View list of all shortened URLs with statistics
- Persistent storage using JSON files


## Technology Stack

- **Backend**: Go with Gin web framework
- **Storage**: In-memory with JSON file persistence
- **Frontend**: HTML, CSS, JavaScript with Tailwind CSS

## Project Structure

- `main.go` - Main application entry point
- `api/handlers/url_handlers.go` - Handler functions for HTTP requests
- `db/memory_store.go` - In-memory URL storage with persistence
- `templates/` - HTML templates
  - `index.html` - Main page template
  - `404.html` - Not found error page
- `static/` - Static assets
  - `css/style.css` - Custom styles and animations
  - `js/app.js` - Frontend JavaScript

## How to Launch

1. Ensure you have Go installed (version 1.16+ recommended)
2. Clone the repository:
   ```
   git clone https://github.com/krishna-nishant/url-shortener-go.git
   cd url-shortener-go
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Run the application:
   ```
   go run main.go
   ```
5. Visit http://localhost:9000 in your browser

## API Endpoints

- `GET /` - Home page
- `POST /shorten` - Create a shortened URL
- `GET /:shortURL` - Redirect to original URL
- `GET /api/urls` - Get all URLs

## Planned Enhancements

- [ ] URL validation
- [ ] QR code generation for shortened URLs
- [ ] Rate limiting to prevent abuse
- [ ] Custom short URLs (allow users to choose their own short code)
- [ ] Advanced analytics dashboard for URL statistics
- [ ] Expiration date for short URLs
- [ ] User authentication for managing personal URLs