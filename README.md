# 🎵 Groupie Tracker

A full-stack web application built with Go (Golang) that consumes a RESTful API to present detailed insights into musical bands and artists. The project processes complex data structures—including members, activity start dates, first album release dates, concert locations, and event dates—and renders them into a smooth, interactive user interface.

---

## 🚀 Key Features

*  **Artist Catalogue:** Dynamic presentation of artists and bands featuring images, creation dates, first album releases, and member rosters.
*  **Filter & Search System:** Custom filtering logic allowing users to search and discover artists based on dynamic criteria.
*  **Geographic & Concert Mapping:** Unified data fetching that links artists directly to their concert dates and worldwide locations.
*  **Client-Server Interactivity:** Clean HTTP request and response handling connecting the frontend interface to backend handlers.
*  **Robust Error Handling:** Built-in HTTP status handling (400, 404, 500) ensuring smooth operation and crash prevention under all conditions.
*  **Zero External Dependencies:** Built purely using the Go Standard Library without any third-party packages.

---

## 💡 Key Learnings & Takeaways

Building this project was a hands-on learning experience in full-stack backend development. Key concepts mastered include:

*  **Mastering Go Standard Library:** Deep dive into net/http, html/template, and encoding/json without relying on external frameworks.
*  **RESTful API Consumption & JSON Parsing:** Structuring Go models (structs) to unmarshal complex nested API responses efficiently.
*  **Clean Architecture:** Designing a maintainable project layout with dedicated packages for handlers, models, and services.
*  **Error Management:** Implementing crash-proof web server practices and graceful error responses for optimal user experience.

---

## 🛠️ Tech Stack

* ⚙️ **Backend:** Go (Golang)
* 🎨 **Frontend:** HTML5, CSS3
* 📡 **Data Interchange:** RESTful API (JSON)

---

## 📁 Project Structure

```text
group-tracker/
├── handlers/            # HTTP request handlers (artist, home)
├── models/              # Go structs and data definitions
├── services/            # API integration & filtering logic
├── static/              # CSS styling files
├── templates/           # HTML templates
├── go.mod               # Go module configuration
├── main.go              # Application entry point
└── README.md            # Documentation
```
---

## 📦 Getting Started

### 📋 Prerequisites

* Go (version 1.18 or higher) installed on your machine.

### ⚙️ Installation & Execution

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/jabutair/group-tracker.git](https://github.com/jabutair/group-tracker.git)
   cd group-tracker
   ```
 2. **Run the application:**
   ```bash
   go run main.go
   ```
 3. **Open in browser:**
   ```bash
   Navigate to http://localhost:8080.
   ```
