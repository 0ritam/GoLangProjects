package main

import (
	"encoding/json" //it uis used to handle encoding and the decoding of the json data

	"fmt"      // provides the formatting the like prntln
	"net/http" // provide sthe http server and clinet functionality , used to create and handle web request http://localhost:8000/
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello , Binayak Sir!") //Prints formatted output to a specified writer, followed by a newline.
}

func greetGuest(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Developer"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name) //Prints formatted output to a specified writer.

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := map[string]string{"message": "Wecome to te Go Web Server!🚀"} //map[keyType]ValueType

	json.NewEncoder(w).Encode(res)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/greet", greetGuest)
	http.HandleFunc("/json", jsonHandler)

	fmt.Println("Server is runnig...")
	http.ListenAndServe(":8000", nil)
}
