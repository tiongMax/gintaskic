# TaskMaster API (Gin Edition) 🚀

**TaskMaster** is a high-performance RESTful API built with the **Gin Gonic** framework.

After mastering the Go standard library in [GoServe], this project marks the transition to professional-grade web frameworks. The goal is to build a robust **Task Management Backend** while exploring the "Gin way" of handling routing, JSON validation, middleware, and database integration.

---

## 📅 Project Roadmap (7-Day Sprint)

I am building this backend-only service over one week to master the Gin lifecycle and idiomatic Go API patterns.

### ☐ **Day 1: The Gin Engine**
- [ ] Project initialization and Gin installation
- [ ] Basic `gin.Default()` setup and the "Ping" endpoint

### ☐ **Day 2: Advanced Routing & Grouping**
- [ ] Route grouping (`/api/v1`)
- [ ] Path parameters (`/tasks/:id`) and Query parameters

### ☐ **Day 3: Request Binding & Validation**
- [ ] Concepts: `c.ShouldBindJSON`, Struct Tags (`binding:"required"`)
- [ ] Implementation: Creating tasks with automatic validation

### ☐ **Day 4: Persistence with GORM**
- [ ] Concepts: ORM vs Raw SQL, SQLite setup
- [ ] Implementation: Moving from in-memory slices to a real database

### ☐ **Day 5: Gin Middleware Mastery**
- [ ] Concepts: `gin.HandlerFunc`, `c.Next()`, `c.Abort()`
- [ ] Implementation: Custom Auth-Key middleware and Request ID tracking

### ☐ **Day 6: Error Handling & Responses**
- [ ] Concepts: Standardizing JSON error envelopes
- [ ] Implementation: Centralized error handling logic

### ☐ **Day 7: Performance & Documentation**
- [ ] Concepts: Benchmarking Gin handlers
- [ ] Implementation: Generating Swagger/OpenAPI docs for the API



---

## 🛠️ Installation & Usage

### Prerequisites

* [Go 1.22+](https://go.dev/dl/) installed.
* [Postman](https://www.postman.com/) or `curl` for testing (No frontend included).

### Quick Start

1. **Initialize and Install Gin:**
```bash
mkdir taskmaster && cd taskmaster
go mod init github.com/yourusername/taskmaster
go get -u github.com/gin-gonic/gin

```


2. **Run the server:**
```bash
go run main.go

```


3. **Test the API:**
```bash
# Create a task
$ curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Gin", "description": "Finish Day 1 of Winter Arc"}'

```



---

## 🧠 Technical Learnings (In Progress)

*Documenting the transition from `net/http` to Gin.*

### **1. The Gin Context (`*gin.Context`)**

In `net/http`, I had to manage `http.ResponseWriter` and `*http.Request` separately. Gin collapses these into a single **Context** object.

**Key Insight:** The Context is the "Swiss Army Knife" of Gin. It handles everything from metadata and flow control (`c.Next()`) to the final JSON output (`c.JSON()`).

### **2. Declarative Validation**

One of the biggest upgrades from the standard library is **Binding**.

**Comparison:**

* **Old way:** Manually unmarshal JSON, check if strings are empty, check if dates are valid, return 400.
* **Gin way:** Define tags in the struct: `json:"title" binding:"required,min=3"`. Gin’s validator (using `go-playground/validator`) does the work automatically.

### **3. Grouped Routing**

Gin allows for logical grouping of routes, which makes API versioning (`/v1`, `/v2`) incredibly simple compared to a flat `ServeMux`.

```go
v1 := router.Group("/api/v1")
{
    v1.GET("/tasks", GetTasks)
    v1.POST("/tasks", CreateTask)
}

```

---

## 📂 Project Structure

```text
.
├── main.go           # Entry point & Router setup
├── handlers/         # Controller logic (The "C" in MVC)
│   └── task.go
├── models/           # Data structures & Validation tags
│   └── task.go
├── middleware/       # Custom Gin middleware
├── repository/       # Database logic (Day 4+)
├── go.mod
└── README.md

```

