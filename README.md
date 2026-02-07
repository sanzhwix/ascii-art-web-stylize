ascii-art-web
Description

ascii-art-web is a web application written in Go that converts user-provided text into a graphical ASCII-art representation.
The application runs an HTTP server and provides a web-based graphical interface where users can enter text, select a visual banner style, and receive the generated ASCII-art output directly in the browser.

The project demonstrates server-side rendering, client–server communication, and string processing, using only the Go standard library.

Features

Web GUI for ASCII-art generation

Converts text into graphical output using ASCII characters

Supports multiple banner styles:

standard

shadow

thinkertoy

Server-side processing and rendering

Clean error handling and stable server behavior

How It Works

The user enters text into a web form

The user selects a banner style

The browser sends the data to the Go server

The server generates ASCII-art using banner definition files

The result is displayed on the web page

HTTP Endpoints
GET /

Serves the main HTML page

Displays:

Text input field

Banner selection options

Submit button

POST /ascii-art

Receives text and selected banner

Generates ASCII-art on the server

Returns the formatted result to the client

HTTP Status Codes

The server returns appropriate HTTP responses:

200 OK — Request processed successfully

400 Bad Request — Invalid input

404 Not Found — Missing banners or templates

500 Internal Server Error — Internal server error

Usage
Run the server
go run .


Open your browser and visit:

http://localhost:8080

Implementation Details
Algorithm Overview

Start an HTTP server using net/http

Parse and validate user input from HTML forms

Load banner files corresponding to the selected style

Convert each character of the input text into its ASCII-art representation

Render the result using Go HTML templates

Handle all errors safely and return proper HTTP status codes

Project Structure
.
├── main.go
├── templates/
│   └── index.html
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── go.mod
└── README.md

Constraints

Written entirely in Go

Uses only standard Go packages

HTML templates located in the templates directory

Server must not crash under any circumstances
