package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

//go:embed template.md
var templateContent string

// Config holds the project configuration
type Config struct {
	User      string
	Repo      string
	Binaries  []string
	Platforms []string
}

// TemplateData holds data for template rendering
type TemplateData struct {
	Title         string
	User          string
	Repo          string
	RepoURL       string
	ReleasesURL   string
	PagesURL      string
	Binaries      []string
	DownloadTable string
}

// Taskfile represents the structure we need from Taskfile.yml
type Taskfile struct {
	Vars map[string]any `yaml:"vars"`
}

func main() {
	user := flag.String("user", "", "GitHub username")
	repo := flag.String("repo", "", "GitHub repo name")
	binaries := flag.String("binaries", "", "Space-separated list of binaries")
	output := flag.String("output", "docs/index.md", "Output file path")
	flag.Parse()

	cfg := Config{
		Platforms: []string{
			"linux/amd64", "linux/arm64",
			"darwin/amd64", "darwin/arm64",
			"windows/amd64", "windows/arm64",
		},
	}

	// If flags provided, use them
	if *user != "" && *repo != "" && *binaries != "" {
		cfg.User = *user
		cfg.Repo = *repo
		cfg.Binaries = strings.Fields(*binaries)
	} else {
		// Otherwise parse Taskfile.yml
		var err error
		cfg, err = parseTaskfile("Taskfile.yml")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing Taskfile.yml: %v\n", err)
			fmt.Fprintf(os.Stderr, "You can also use flags: -user USER -repo REPO -binaries \"bin1 bin2\"\n")
			os.Exit(1)
		}
	}

	// Generate the docs
	content, err := generate(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating docs: %v\n", err)
		os.Exit(1)
	}

	// Write output
	if err := os.WriteFile(*output, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", *output, err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", *output)
}

func parseTaskfile(path string) (Config, error) {
	cfg := Config{
		Platforms: []string{
			"linux/amd64", "linux/arm64",
			"darwin/amd64", "darwin/arm64",
			"windows/amd64", "windows/arm64",
		},
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	var tf Taskfile
	if err := yaml.Unmarshal(data, &tf); err != nil {
		return cfg, err
	}

	if v, ok := tf.Vars["GITHUB_USER"].(string); ok {
		cfg.User = v
	}
	if v, ok := tf.Vars["GITHUB_REPO"].(string); ok {
		cfg.Repo = v
	}
	if v, ok := tf.Vars["BINARIES"].(string); ok {
		cfg.Binaries = strings.Fields(v)
	}
	if v, ok := tf.Vars["PLATFORMS"].(string); ok {
		cfg.Platforms = strings.Fields(v)
	}

	if cfg.User == "" || cfg.Repo == "" || len(cfg.Binaries) == 0 {
		return cfg, fmt.Errorf("missing required vars: GITHUB_USER, GITHUB_REPO, BINARIES")
	}

	return cfg, nil
}

func generate(cfg Config) (string, error) {
	data := TemplateData{
		Title:         cfg.Repo,
		User:          cfg.User,
		Repo:          cfg.Repo,
		RepoURL:       fmt.Sprintf("https://github.com/%s/%s", cfg.User, cfg.Repo),
		ReleasesURL:   fmt.Sprintf("https://github.com/%s/%s/releases", cfg.User, cfg.Repo),
		PagesURL:      fmt.Sprintf("https://%s.github.io/%s", cfg.User, cfg.Repo),
		Binaries:      cfg.Binaries,
		DownloadTable: generateDownloadTable(cfg),
	}

	tmpl, err := template.New("docs").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func generateDownloadTable(cfg Config) string {
	var sb strings.Builder

	// Header
	sb.WriteString("| | Linux | macOS | Windows |\n")
	sb.WriteString("|---|---|---|---|\n")

	baseURL := fmt.Sprintf("https://github.com/%s/%s/releases/latest/download", cfg.User, cfg.Repo)

	for _, bin := range cfg.Binaries {
		sb.WriteString(fmt.Sprintf("| **%s** ", bin))

		// Linux
		sb.WriteString(fmt.Sprintf("| [amd64](%s/%s_linux_amd64) / [arm64](%s/%s_linux_arm64) ",
			baseURL, bin, baseURL, bin))

		// macOS
		sb.WriteString(fmt.Sprintf("| [Intel](%s/%s_darwin_amd64) / [Apple Silicon](%s/%s_darwin_arm64) ",
			baseURL, bin, baseURL, bin))

		// Windows
		sb.WriteString(fmt.Sprintf("| [amd64](%s/%s_windows_amd64.exe) / [arm64](%s/%s_windows_arm64.exe) |\n",
			baseURL, bin, baseURL, bin))
	}

	return sb.String()
}
