package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
)

func LoadPage(response http.ResponseWriter, request *http.Request) {
	path, _ := strings.CutPrefix(request.URL.Path, "/");

	if (path == "" || path == "index") {
		path = "index.html";
	}

	data, err := os.ReadFile(path);
	if (err != nil) {
		fmt.Printf("[ERR] :: Page %s not found\n", path);

		http.NotFound(response, request);
	} else {
		fmt.Printf("[INFO] :: Page %s loaded\n", path);

		response.Header().Set("Content-Type", "text/html");
		response.Write(data);
	}
}

func main() {
	http.HandleFunc("/", LoadPage);
	http.ListenAndServe(":8080", nil);
}
