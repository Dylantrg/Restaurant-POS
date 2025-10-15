package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler shows a friendly landing page.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!doctype html>
<html>
  <head><meta charset="utf-8"><title>POS Backend</title></head>
  <body style="font-family: system-ui, -apple-system, Segoe UI, Roboto, sans-serif; padding: 24px;">
    <h1>POS Backend</h1>
    <p>Service is running.</p>
    <p>Health: <a href="/api/health">/api/health</a></p>
  </body>
</html>`)
}
