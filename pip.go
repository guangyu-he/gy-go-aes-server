package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 包目录
const PackageDir = "./packages"
const ValidToken = "my-secret-token"

// 模板生成索引页面
var indexTemplate = `<html>
<head><title>Python Package Index</title></head>
<body>
<h1>Available Packages</h1>
<ul>
{{range .}}
    <li><a href="/pip/{{.}}">{{.}}</a></li>
{{end}}
</ul>
</body>
</html>`

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(PackageDir)
	if err != nil {
		http.Error(w, "Failed to list packages", http.StatusInternalServerError)
		return
	}

	var packages []string
	for _, file := range files {
		if file.IsDir() {
			packages = append(packages, file.Name())
		}
	}

	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	tmpl.Execute(w, packages)
}

func PackageHandler(w http.ResponseWriter, r *http.Request) {
	packageName := filepath.Base(r.URL.Path)
	packagePath := filepath.Join(PackageDir, packageName)

	// 检查 token 参数
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token missing", http.StatusForbidden)
		return
	}

	files, err := os.ReadDir(packagePath)
	if err != nil {
		http.Error(w, "Package not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "<html><body>")
	for _, file := range files {
		link := fmt.Sprintf("/packages/%s/%s?token=%s", packageName, file.Name(), token)
		fmt.Fprintf(w, `<a href="%s">%s</a><br>`, link, file.Name())
	}
	fmt.Fprintln(w, "</body></html>")
}

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token != ValidToken {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	cleanPath := path.Clean(r.URL.Path)
	rootDir := "./packages/"
	fullPath := rootDir + cleanPath

	if !strings.HasPrefix(fullPath, rootDir) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, fullPath)
}
