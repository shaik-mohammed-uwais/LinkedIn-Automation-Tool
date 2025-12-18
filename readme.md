# LinkedIn Automation Tool (Go)

## 🎥 Demonstration Video

Project walkthrough video (setup, configuration, execution, and features):

👉 **Video Link:** [https://your-video-link-here](https://your-video-link-here)

## 📌 Project Overview

A LinkedIn automation bot written in Go that uses browser automation to log in, maintain sessions, search profiles/content, and send messages while behaving like a real human.

The project is structured to be modular, configurable, and stealthy, focusing on avoiding bot detection by mimicking real user actions.

The tool simulates a LinkedIn automation workflow such as logging in, searching profiles, sending connection requests, and managing limits. The focus of this project is **code structure and design**, not misuse of LinkedIn services.

---

## 🛠 Tools & Technologies Used

- **Go (Golang)** – Core programming language
- **Go Modules** – Dependency management
- **Environment Variables (.env)** – Secure configuration handling
- **Chromium / Browser Automation** – For simulating LinkedIn actions (headless or non-headless)
- **Logging Package** – For structured logs

---

## 📂 Project Structure

```
cmd/            → Application entry point (main.go)
config/         → Loads and manages environment variables
auth/           → Handles login & authentication logic
search/         → Profile search functionality
connect/        → Sends connection requests
message/        → Sends messages to profiles
stealth/        → Delay & safety logic to avoid detection
storage/        → Saves local state (limits, progress)
logger/         → Centralized logging system
.env.example    → Environment variables template
```

---

## ⚙️ Environment Setup

This project uses environment variables for configuration.

### 1️⃣ Create `.env` file

Copy the example file:

```bash
cp .env.example .env
```

### 2️⃣ `.env.example`

```env
LINKEDIN_EMAIL=your_email_here
LINKEDIN_PASSWORD=your_password_here
HEADLESS=false
DAILY_CONNECTION_LIMIT=10
```

### 3️⃣ Environment Variables Explanation

- **LINKEDIN_EMAIL** – LinkedIn account email
- **LINKEDIN_PASSWORD** – LinkedIn account password
- **HEADLESS** – Run browser in headless mode (`true` or `false`)
- **DAILY_CONNECTION_LIMIT** – Maximum connection requests per day

> ⚠️ Do not commit the `.env` file to GitHub.

---

## ▶️ How to Run the Project

### 1️⃣ Install dependencies

```bash
go mod tidy
```

### 2️⃣ Run the application

```bash
go run cmd/main.go
```

---

## ✨ Key Features

- Modular and scalable Go architecture
- Secure configuration using environment variables
- Clear separation of responsibilities
- Safety limits and delays to simulate real usage
- Persistent storage to track daily limits

---
