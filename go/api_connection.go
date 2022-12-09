/*
 * Сервис интерфейса P2r
 *
 * Сервис для реализации процедур интерфейса P2r. 
 *
 * API version: 1.0.
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
)

func ReleaseConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
