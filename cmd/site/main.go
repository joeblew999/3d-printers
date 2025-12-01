package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

var pageTpl = template.Must(template.New("page").Parse(`
<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>{{.Title}}</title>
  <style>
    body { font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif; max-width: 900px; margin: 40px auto; padding: 0 16px; line-height: 1.6; color: #111; }
    pre { background: #f6f8fa; padding: 12px; overflow: auto; border-radius: 6px; }
    code { background: #f6f8fa; padding: 2px 4px; border-radius: 4px; }
    a { color: #0366d6; }
  </style>
</head>
<body>
{{.Content}}
</body>
</html>
`))

func main() {
	dir := flag.String("dir", "docs", "directory to serve")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		servePath := r.URL.Path
		if strings.HasSuffix(servePath, "/") {
			servePath = filepath.Join(servePath, "index.md")
		}
		if strings.HasSuffix(servePath, ".md") {
			serveMarkdown(w, filepath.Join(*dir, strings.TrimPrefix(servePath, "/")))
			return
		}
		http.FileServer(http.Dir(*dir)).ServeHTTP(w, r)
	})

	log.Printf("serving %s on %s", *dir, *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func serveMarkdown(w http.ResponseWriter, path string) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		http.NotFound(w, nil)
		return
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(data, &buf); err != nil {
		http.Error(w, "markdown render error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = pageTpl.Execute(w, map[string]any{
		"Title":   filepath.Base(path),
		"Content": template.HTML(buf.String()),
	})
}
