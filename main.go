package main

import (
	"fmt"
	"log"
	"net/http"
	"math/big"
)

func check_prime(x string) bool {
	var x1 big.Int
	x1.SetString(x, 10)
	return x1.ProbablyPrime(x1.BitLen())
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodGet {
		// Отображаем форму
		fmt.Fprintf(w, `
        <!DOCTYPE html>
       <html>
       <head>
       <title>Page Title</title>
       </head>
       <body>

       <form method="POST">
           <input type="text" name="int" placeholder="Enter your integer">
           <button type="submit">Submit</button>
       </form>

       </body>
       </html>
        `)
	} else if r.Method == http.MethodPost {
		// Обрабатываем данные формы
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}
		
		val := r.FormValue("int")
		
		if check_prime(val) {
			fmt.Fprintf(w, "%v is prime", val)
		} else {
			fmt.Fprintf(w, "%v is not prime", val)
		}
	}
}

func main() {
	// Регистрируем обработчики для разных маршрутов
	http.HandleFunc("/", formHandler)

	// Запускаем сервер
	fmt.Println("Starting server at port 8080")
	log.Print(http.ListenAndServe(":8080", nil))
}
