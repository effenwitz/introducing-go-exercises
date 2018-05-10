package main

import (
	"fmt"
	"io"
	"net/http"
)

//function name is the rest endpoint
func hello(res http.ResponseWriter, re *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	io.WriteString(res,
		`<DOCTYPE html>
            <html>
              <head>
                  <title>Hello, World</title>
              </head>
              <body>
                  Hello, World!
              </body>
</html>`,
	)
}

func main() {
	fmt.Println("HTTP Server")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
	//handle static files
	// http.Handle("/assets", http.StripPrefix("/assets/",
	// 	http.FileServer(http.Dir("assets")),
	// 	),
	// )
}
