package handler

import "net/http"

// HTTPInterceptor functions as http request interceptor for token validation
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			if len(username) < 3 || !IsTokenValid(token) {
				http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
			}
			h(w, r)
		})
}

// IsTokenValid check token
func IsTokenValid(token string) bool {
	// 1 check token len == 40
	if len(token) != 40 {
		return false
	}
	// 2 check expiretime

	// 3 checkout

	return false
}
