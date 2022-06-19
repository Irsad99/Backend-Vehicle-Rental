package middleware

import "net/http"

type Middleware func(string, http.HandlerFunc) http.HandlerFunc

func Do(hf http.HandlerFunc, role string , middle ...Middleware) http.HandlerFunc {
	for _, m := range middle {
		hf = m(role, hf)
	}

	return hf
}
