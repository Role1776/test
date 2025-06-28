package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "иди нахуй", http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
