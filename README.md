# ?? Groupie Tracker

A responsive web application built with **Go (Golang)** that consumes a RESTful API to present detailed insights into musical bands and artists. The project parses complex data structures—including members, activity years, first album release dates, concert locations, and event dates—and renders them into an interactive user interface.

---

## ?? Features

* **Artist Catalogue:** Display of artists and bands featuring high-quality images, creation dates, first album releases, and member rosters.
* **Filter & Search:** Custom filter logic allowing users to search and filter artists dynamically.
* **Geographic & Schedule Relations:** Unified data fetching connecting artists directly to their concert dates and locations.
* **Client-Server Interactivity:** Real-time HTTP requests and backend responses handled with clean separation of models, handlers, and services.
* **Standard Library Pureness:** Developed purely using the Go Standard Library without any external dependencies.

---

## ??? Tech Stack

* **Backend:** Go (\
et/http\, \html/template\, \encoding/json\)
* **Frontend:** HTML5, CSS3
* **Data Interchange:** RESTful API (JSON)

---

## ?? Project Structure

\\\	ext
group-tracker/
+-- handlers/            # HTTP request handlers (artist, home)
+-- models/              # Go structs and data definitions
+-- services/            # API integration & filtering logic
+-- static/              # CSS styling files
+-- templates/           # HTML templates
+-- go.mod               # Go module configuration
+-- main.go              # Application entry point
+-- README.md            # Documentation
\\\

---

## ?? Getting Started

### Prerequisites

* [Go](https://golang.org/doc/install) (version 1.18 or higher) installed on your system.

### Installation & Execution

1. **Clone the repository:**
   \\\ash
   git clone https://github.com/YOUR_USERNAME/group-tracker.git
   cd group-tracker
   \\\

2. **Run the server:**
   \\\ash
   go run main.go
   \\\

3. **Access the web application:**
   Open your browser and navigate to \http://localhost:8080\.

---

## ?? Running Unit Tests

To run automated tests across all packages:

\\\ash
go test ./... -v
\\\

---

## ?? License

This project was developed for educational purposes as part of the Software Engineering curriculum.
