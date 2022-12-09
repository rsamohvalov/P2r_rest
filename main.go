/*
 * Сервис интерфейса P2r
 *
 * Сервис для реализации процедур интерфейса P2r. 
 *
 * API version: 1.0.
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
	sw "github.com/rsamohvalov/P2r_rest/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
