package main

import (
	"errors"
	"fmt"
	"net/http"
)

// define port number
const portNumber = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2,3)
	_,_ = fmt.Fprintf(w, fmt.Sprintf("This is about page and the value of two numbers sum is %d", sum))
}

/** if we define a function name with simple letter
 *  it means this function only scoped for this package.
 *  For example this addValue function only scoped for main package.
**/
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil{
		_, _ = fmt.Fprintf(w, "cannot divide by 0")
		return
	}
	// we use _, _ for skip the return values already comes with Fprintf function
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0{
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	http.HandleFunc("/headers", headers)

	fmt.Println(fmt.Sprintf("Start application using port %s", portNumber))
	_ = http.ListenAndServe(":3000", nil)
}