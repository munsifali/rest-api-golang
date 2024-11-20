package student

import "net/http"

func CreateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student Api"))
	}
}
