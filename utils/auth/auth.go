package auth

import (
	"context"
	"encoding/json"
	"event_backend/models"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception models.Exception

//JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("Authorization")
		token := strings.TrimSpace(header)

		fmt.Println(token)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}
		//user := &models.User{}

		_, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		//fmt.Println(tk.Email)

		json.NewEncoder(w).Encode(tk)

		ctx := context.WithValue(r.Context(), "user", tk)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
