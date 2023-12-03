package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var tfs embed.FS

func main() {
	var day int
	flag.IntVar(&day, "day", 1, "day number")
	flag.Parse()

	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	tmpl, err := template.ParseFS(tfs, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("parsing templates directory: %s", err)
	}

	dstPrefix := fmt.Sprintf("day%02d", day)
	dstPath := filepath.Join(dirname(), "..", "days", dstPrefix)
	if err := os.MkdirAll(dstPath, 0755); err != nil {
		log.Fatalf("failed to create dir: %s", err)
	}

	if err := processTemplates(tmpl, dstPrefix); err != nil {
		log.Fatalf("failed to processTemplates: %s", err)
	}
}

func processTemplates(tmpl *template.Template, dstPrefix string) error {
	return fs.WalkDir(tfs, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking directory: %w", err)
		}
		if path == "templates" || d.IsDir() {
			return nil
		}
		return processFile(tmpl, path, dstPrefix)
	})
}

func processFile(tmpl *template.Template, path, dstPrefix string) error {
	suffix := strings.TrimPrefix(path, "templates/")
	suffix = strings.TrimSuffix(suffix, ".tmpl")

	dstPath := filepath.Join(dirname(), "..", "days", dstPrefix, suffix)
	if err := ensureNotOverwriting(dstPath); err != nil {
		return fmt.Errorf("ensureNotOverwriting: %w", err)
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create file %s: %w", dstPath, err)
	}
	defer dst.Close()

	if err := tmpl.ExecuteTemplate(dst, filepath.Base(path), nil); err != nil {
		return fmt.Errorf("executing template for %s: %w", path, err)
	}

	return nil
}

func dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("getting calling function")
	}
	return filepath.Dir(filename)
}

func ensureNotOverwriting(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		return fmt.Errorf("file already exists: %s", filename)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("stat file: %s, %w", filename, err)
	}
	return nil
}
