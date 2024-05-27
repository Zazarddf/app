package main

import "net/http"

func createrSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}
	w.Write([]bbyte("Создание новой заметки"))
}
func main() {
	http.HandleFunc("/snippet", createSnippet)
	http.ListenAndServe(":8080", nil)
}
