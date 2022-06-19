package middleware

import (
	// "fmt"
	"net/http"
	"strconv"
	"strings"

	"BackendGo/src/helpers"
)

func CheckAuth(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")
		var response helpers.Response

		if !strings.Contains(headerToken, "Bearer") {
			response.ResponseJSON(401, "Invalid Header Type").Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := helpers.CheckToken(token, role)
		if err != nil {
			response.ResponseJSON(401, "Token Salah, Silahkan login kembali").Send(w)
			return
		}

		if !checkToken {
			response.ResponseJSON(401, "Anda Bukan Admin").Send(w)
			return
		}

		eksToken, err := helpers.EksToken(token)
		if err != nil {
			return
		}

		idUser := strconv.Itoa(int(eksToken.User_Id))

		r.Header.Set("id", idUser)
		r.Header.Set("role", eksToken.Role)

		next.ServeHTTP(w, r)
	}
}
