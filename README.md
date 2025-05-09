# Snippetbox

## About This Project
Snippetbox is a simple web application for saving and sharing text snippets (like small pieces of code or notes). This project follows the "Let's Go" book by Alex Edwards as a learning exercise.

## What It Does
- Shows a homepage with a list of snippets
- Lets you view individual snippets
- Allows creating new snippets through a form
- Serves static files (CSS, JavaScript, images)

## Project Structure
- `cmd/web`: Contains the main application code
  - `main.go`: Sets up the web server and routes
  - `handlers.go`: Handles web requests
- `ui`: Contains the user interface files
  - `html`: HTML templates
  - `static`: Static files (CSS, JavaScript, images)

## Implemented Features
- Basic web application structure with clean organization
- Page routing (home page, view snippet, create snippet)
- HTML templating for displaying pages
- Static file serving for CSS, JavaScript, and images
- Form handling for creating new snippets
- Custom logging for better debugging
- HTTP header customization
- Error handling and user-friendly error messages

## How to Run
To start the server:
```
go run ./cmd/web
```

The application will be available at http://localhost:4000

You can change the port by using the -addr flag:
```
go run ./cmd/web -addr=:8080
```
