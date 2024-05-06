package errors

import "net/http"

const (
	MySQLDuplicateEntryErr = 1062
)

func getStatusCodeForMySQLErr(code uint16) int {
	switch code {
	case MySQLDuplicateEntryErr:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
