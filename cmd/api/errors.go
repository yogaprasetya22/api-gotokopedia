package main

import (
	"fmt"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_INTERNAL]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	// log.Printf("\033[33m[ERROR_FORBIDDEN]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[0m", r.Method, r.URL.Path)
	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")
	writeJSONError(w, http.StatusForbidden, "forbidden")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_BAD_REQUEST]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_CONFLICT]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Errorf("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_NOT_FOUND]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_UNAUTHORIZED]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Warnf("unauthorized response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusUnauthorized, fmt.Sprintf("unauthorized: %s", err))
}
func (app *application) unauthorizedActiveErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_UNAUTHORIZED]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Warnf("unauthorized response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusForbidden, "unauthorized active")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("\033[33m[ERROR_UNAUTHORIZED]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m errors: \033[90m%s\033[0m", r.Method, r.URL.Path, err)
	app.logger.Warnf("unauthorized response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	// log.Printf("\033[33m[ERROR_RATE_LIMIT_EXCEEDED]: \033[35m\033[1m%s\033[33m:\033[34m%s\033[33m retry-after: \033[90m%s\033[0m", r.Method, r.URL.Path, retryAfter)
	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}
