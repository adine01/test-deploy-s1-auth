since **pod is crashing (CrashLoopBackOff)**, the safest way is to go through a **systematic checklist in your code** to make sure it runs cleanly in Choreo.

Here’s the full list of what to check:

---

## 🔎 1. Environment Variables

* Your code should **not** depend on `.env` files in production.
* Instead, it should read env vars directly:

  ```go
  dbURL := os.Getenv("DATABASE_URL")
  if dbURL == "" {
      log.Fatal("DATABASE_URL environment variable is required")
  }
  ```
* Make sure **every required config** (`DATABASE_URL`, `PORT`, secrets, etc.) is read this way.
* ✅ In Choreo → set them under **Configurations → Environment Variables**.

---

## 🔎 2. Port Binding

* Choreo expects your container to listen on **port 8080**.
* In Go, make sure you use:

  ```go
  http.ListenAndServe(":8080", router)
  ```
* ❌ Don’t hardcode `3000`, `5000`, etc. — otherwise the health check fails → pod restarts.

---

## 🔎 3. Health Check Endpoint

* Choreo probes `/health` (or your defined health path).
* It must return **200 OK quickly** with JSON or plain text.
* Example:

  ```go
  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(`{"status":"ok"}`))
  })
  ```
* If it depends on DB/Redis and those are down, it’s fine to still return `ok` (otherwise service crashes repeatedly).

---

## 🔎 4. Database Initialization

* At startup, don’t block forever waiting for DB connection.
* Example safe pattern:

  ```go
  db, err := sql.Open("postgres", dbURL)
  if err != nil {
      log.Fatalf("failed to open db: %v", err)
  }
  if err := db.Ping(); err != nil {
      log.Fatalf("failed to connect to db: %v", err)
  }
  ```
* Log the error instead of panicking without explanation.
* ✅ Test the same `DATABASE_URL` locally before deploying.

---

## 🔎 5. Graceful Error Handling

* Don’t use bare `panic()` on startup. Instead use `log.Fatal` or `fmt.Println` so logs show the issue.
* Choreo restarts the container automatically if the process exits.

---

## 🔎 6. Startup Logs

* Add logs at each step (`Loaded env vars`, `Connecting DB`, `Listening on :8080`, etc.).
* This way, when you check Choreo logs, you’ll know where it failed.

---

## 🔎 7. Optional: Config for Multiple Envs

* In **dev**, use `.env`.
* In **prod (Choreo)**, read from `os.Getenv`.
* Example:

  ```go
  if os.Getenv("APP_ENV") != "production" {
      _ = godotenv.Load()
  }
  ```

