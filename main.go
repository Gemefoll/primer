package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
)

func check_prime(x string) bool {
	var x1 big.Int
	x1.SetString(x, 10)
	return x1.ProbablyPrime(x1.BitLen())
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, `
        <!DOCTYPE html>
       <html>
       <head>
       <title>PRIMER</title>
       </head>
       <body>

       <form method="POST">
           <input type="text" name="int" placeholder="Enter your integer">
           <button type="submit">Submit</button>
       </form>

       </body>
       </html>
        `)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}

		val := r.FormValue("int")

		var ans string
		if check_prime(val) {
			ans = val + " is prime"
		} else {
			ans = val + " is not prime"
		}

		fmt.Fprintf(w, `
        <!DOCTYPE html>
       <html>
       <head>
       <title>PRIMER</title>
       </head>
       <body>

       <p> `+ans+` </p>
       <form method="GET">
           <button type="retry">Retry</button>
       </form>

       </body>
       </html>
        `)

	}
}

func main() {
	http.HandleFunc("/", formHandler)

	fmt.Println("Starting server at port 8080")
	log.Print(http.ListenAndServe(":8080", nil))
}
