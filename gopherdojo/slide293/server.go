package slide293

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func StartServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
