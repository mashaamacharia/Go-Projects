# 🧠 AI Text Summarizer API (Go + AI)

## 📌 Overview

This project is a minimal AI-powered REST API built using **Go** and the **Gin framework**. It integrates with the Anthropic API to summarize text and extract key points.

### 🎯 Purpose

* Demonstrate use of a new programming language (**Go**)
* Integrate AI into a backend service
* Build a simple, working API within a short timeframe

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin
* **AI Provider:** Anthropic (Claude models)
* **Environment Config:** `godotenv`

---

## ⚙️ Setup Instructions

### 1. Clone or create project

```bash
mkdir Summarizer
cd Summarizer
```

### 2. Initialize Go module

```bash
go mod init Summarizer
```

### 3. Install dependencies

```bash
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
```

### 4. Create `.env` file

```env
ANTHROPIC_API_KEY=your_api_key_here
```

### 5. Run the application

```bash
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 🚀 Minimal Working Example

### Endpoint:

```
POST /summarize
```

### Request:

```json
{
  "text": "Artificial Intelligence is transforming industries..."
}
```

### Response:

```json
{
  "summary": "AI is transforming industries by improving efficiency.",
  "key_points": [
    "Used in healthcare and finance",
    "Improves automation",
    "Raises ethical concerns"
  ]
}
```

---

## 🤖 AI Prompt Used

```text
Return ONLY valid JSON. No explanation.

Format:
{
  "summary": "...",
  "key_points": ["...", "..."]
}

Text:
<user input>
```

### 🧠 Why this prompt?

* Forces structured JSON output
* Avoids messy text formatting
* Makes parsing easier in Go

---

## 📘 Learning Reflections

* Learned how to build APIs using Go and Gin
* Understood how to integrate external AI APIs
* Learned importance of **prompt engineering**
* Practiced structuring backend code into layers (handler vs service)
* Learned how to handle JSON parsing and HTTP requests in Go

---

## ⚠️ Common Errors & Fixes

### 1. ❌ `not_found_error` (model issue)

**Cause:** Invalid or unavailable model
**Fix:** Use:

```
claude-sonnet-4-20250514
```

---

### 2. ❌ 401 Unauthorized

**Cause:** Missing or incorrect API key
**Fix:**

* Check `.env` file
* Ensure key is valid
* Restart server after changes

---

### 3. ❌ JSON parsing fails

**Cause:** AI returns text instead of JSON
**Fix:**

* Improve prompt:

  ```
  Return ONLY valid JSON
  ```
* Add fallback error handling

---

### 4. ❌ Empty or invalid request

**Cause:** Missing `"text"` field
**Fix:**

* Validate request input in handler

---

## 📚 Reference Resources

* Go Programming Language docs: https://go.dev/doc/
* Gin docs: https://gin-gonic.com/docs/
* Anthropic API docs: https://docs.anthropic.com/

---

## 🎯 Conclusion

This project demonstrates how to:

* Build a REST API using Go
* Integrate AI into backend services
* Design clean and minimal architecture

It focuses on simplicity, functionality, and practical AI usage within a short development timeframe.
