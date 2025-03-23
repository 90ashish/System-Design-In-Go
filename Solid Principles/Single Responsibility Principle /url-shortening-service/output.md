# Steps to Test the URL Shortener System in Golang

After running the command:

```bash
go run main.go
```

You'll see the output like this:

```
Original URL: https://www.google.com/search?q=golang+system+design
Shortened URL: http://short.url/X3U2rV
Server is running on :8080
```

---

## Step 1: Test URL Redirection

1. **Copy the generated short URL** (e.g., `http://short.url/X3U2rV`).
2. **Open your browser** or use `curl` to test the redirection.

### In Browser
➡️ Visit: [http://localhost:8080/X3U2rV](http://localhost:8080/X3U2rV)

**Expected Result:** It should redirect you to:  
`https://www.google.com/search?q=golang+system+design`

### Using `curl` (CLI)
```bash
curl -i http://localhost:8080/X3U2rV
```

**Expected Response:**
```
HTTP/1.1 302 Found
Location: https://www.google.com/search?q=golang+system+design
```

---

## Step 2: Test Analytics Service

To check the visit count for a URL:

**URL Format:**  
```
http://localhost:8080/analytics?shortCode=X3U2rV
```

### Example Command Using `curl`
```bash
curl "http://localhost:8080/analytics?shortCode=X3U2rV"
```

**Expected Output:**
```
URL X3U2rV has been visited 1 times
```

- If you refresh the short URL multiple times, the visit count should increment accordingly.
