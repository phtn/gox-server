package utils

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// HTML template for syntax highlighting JSON
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Gox</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/themes/prism.min.css" rel="stylesheet" />
    <script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/prism.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.25.0/components/prism-json.min.js"></script>
</head>
<body>
    <h1>G0X</h1>
    <pre><code class="language-json">{{ . }}</code></pre>
</body>
</html>
`

func Indent(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	jsonData, err := json.MarshalIndent(data, "", "  ")
	ErrHandler(w, err)

	tmpl, err := template.New("jsonHighlight").Parse(htmlTemplate)
	ErrHandler(w, err)

	err = tmpl.Execute(w, string(jsonData))
	ErrHandler(w, err)

}

func ErrHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
