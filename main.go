package main

import (
	"beginner_crudapi/route"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 참고 : https://codesk.tistory.com/4 <- encoding/json 설명
// 참고 : https://jeonghwan-kim.github.io/dev/2019/01/18/go-encoding-json.html <- encoding과 marshal에 대한 설명

var movies []route.Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, route.Movie{ID: "1", Isbn: "123456", Title: "Golang", Director: &route.Director{Firstname: "Lee", Lastname: "Sungwoo"}})
	movies = append(movies, route.Movie{ID: "2", Isbn: "789101", Title: "Gopher", Director: &route.Director{Firstname: "park", Lastname: "Misun"}})
	r.HandleFunc("/movies", route.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", route.GetMovie).Methods("GET")
	r.HandleFunc("/movies", route.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", route.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", route.DeleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
