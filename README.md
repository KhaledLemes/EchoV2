# Go Echo Server 

A simple, beginner-friendly web application built with Go that demonstrates basic web server setup, routing, and handling POST requests using the **`net/http`** package and the **`gorilla/mux`** router.

## How It Works

This application serves an HTML page with a single form. When a user submits the form with a message, the server:

1.  Receives the POST request on the `/echo` route.
2.  Parses the form data.
3.  Reads the value of the `message` field.
4.  Sends the exact message back to the client as a response.

Essentially, whatever you type, the server "echoes" back to you!

## Prerequisites

Before you begin, ensure you have the following installed:

* **Go (1.18 or higher):** For running the application.
* **Git:** For cloning the repository.
