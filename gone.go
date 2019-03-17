package main

import (
    "fmt"
    "net/http"
)

var response = `<!doctype html>
<html>
    <head>
        <meta charset="UTF-8">
        <meta name="robots" content="noindex">
        <title></title>
    </head>
    <body></body>
</html>`

func main() {
    http.HandleFunc("/", goneHandler)
    http.ListenAndServe(":80", nil)
}

func goneHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, response)
}
