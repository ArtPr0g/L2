package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import (
	"fmt"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Здесь можно добавить проверку аутентификации пользователя
		fmt.Println("AuthMiddleware checking authentication")
		next(w, r)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Здесь можно добавить логгирование запроса
		fmt.Println("LoggingMiddleware logging the request")
		next(w, r)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func main() {
	indexHandler := http.HandlerFunc(IndexHandler)

	// Собираем цепочку обработчиков
	handler := AuthMiddleware(LoggingMiddleware(indexHandler))

	http.Handle("/", handler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
