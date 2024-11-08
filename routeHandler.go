package main

import "net/http"

func errorRouteHandler(w http.ResponseWriter, r *http.Request){
	errorHandler(w, 400, "Bad request")
}

func routeHandler(w http.ResponseWriter, r *http.Request){
	responseHandler(w, 200, "OK")
}