# Web Page Analyzer – Submission Document

## 🧩 Overview

This application analyzes a given web page URL and provides information about:

* HTML version
* Page title
* Heading counts
* Internal and external links
* Presence of a login form
* Error handling for unreachable URLs

The solution is implemented using:

* **Backend:** Golang
* **Frontend:** React (Vite)
* **Deployment:** Docker & Docker Compose

---

## 🚀 Build & Run Instructions

### Prerequisites

* Docker
* Docker Compose

---

### Steps to Run

1. Clone the repository:

```
https://github.com/hasharts/lucytech
```

2. Build and run the application:

```
docker compose up --build
```

3. Access the application:

* Frontend: http://localhost:3000
* Backend API: http://localhost:8080/api/analyze

---

## ⚙️ Application Architecture

```
Browser (React UI)
        ↓
Frontend (Nginx container)
        ↓
Backend API (Go service)
        ↓
Web page fetch + HTML parsing (goquery)
```

---

## 🧠 Design Decisions & Assumptions

### 1. HTML Version Detection

* HTML version is determined by parsing the **DOCTYPE** from raw HTML.
* Since `goquery` does not expose DOCTYPE, the raw HTML string is inspected.

---

### 2. Internal vs External Links

* Links starting with `http://` or `https://` are treated as external.
* All other links are considered internal.

> Note: Domain-based comparison is not implemented for simplicity.

---

### 3. Login Form Detection

* Presence of `<input type="password">` is used as an indicator of a login form.

---

### 4. Error Handling

* If the URL is unreachable or returns a non-200 status:

  * Backend returns HTTP error
  * Frontend displays the error message

---

### 5. CORS Handling

* CORS is enabled in the backend to allow communication between frontend and backend during development.

---

### 6. Dockerized Setup

* Multi-stage Docker builds are used:

  * Go backend compiled in Alpine
  * React app built and served via Nginx

---

## ⚠️ Limitations

* HTML version detection is based on string matching and may not cover all edge cases.
* Internal/external link classification is simplified.
* Broken link detection is not implemented.
* No authentication or rate limiting.

---

## 🚀 Possible Improvements

### 🔹 Functional Enhancements

* Detect and count headings (H1–H6)
* Validate links and detect broken links (with concurrency)
* Domain-aware internal vs external link detection

---

### 🔹 Performance Improvements

* Use goroutines to parallelize link validation
* Add request timeouts and retries

---

### 🔹 Architecture Improvements

* Add middleware for logging and error handling
* Introduce structured JSON error responses
* Use environment-based configuration

---

### 🔹 Frontend Improvements

* Improve UI/UX with Tailwind CSS
* Add loading spinner and better error display
* Form validation before submission

---

### 🔹 DevOps Improvements

* Add CI/CD pipeline (GitHub Actions)
* Add Kubernetes deployment manifests
* Add health checks and monitoring

---

## ✅ Summary

This solution demonstrates:

* Clean separation of frontend and backend
* REST API design in Go
* HTML parsing using goquery
* Containerized deployment using Docker
* Practical handling of real-world issues (CORS, networking, builds)

---

