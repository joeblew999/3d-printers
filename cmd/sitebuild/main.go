package main

import (
	"bytes"
	"flag"
	"html/template"
	"io"
	"log"
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
  <meta name="viewport" content="width=device-width, initial-scale=1">
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
	src := flag.String("src", "docs", "source directory")
	out := flag.String("out", ".site", "output directory for static HTML")
	flag.Parse()

	if err := os.RemoveAll(*out); err != nil {
		log.Fatalf("clean output: %v", err)
	}
	if err := copyAndRender(*src, *out); err != nil {
		log.Fatalf("build site: %v", err)
	}
	log.Printf("built site -> %s", *out)
}

func copyAndRender(src, out string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(out, rel)

		if strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			// Render markdown to HTML
			destPath = strings.TrimSuffix(destPath, ".md") + ".html"
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var buf bytes.Buffer
			if err := goldmark.Convert(data, &buf); err != nil {
				return err
			}
			if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
				return err
			}
			f, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer f.Close()

			return pageTpl.Execute(f, map[string]any{
				"Title":   strings.TrimSuffix(info.Name(), ".md"),
				"Content": template.HTML(buf.String()),
			})
		}

		// Copy other assets as-is
		if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
			return err
		}
		return copyFile(path, destPath)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
