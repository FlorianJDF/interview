package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

func getPaths(url string) (basepath string, nextpath string) {
	cleanPath := path.Clean(url)[1:]
	basePath := strings.Split(cleanPath, "/")[0]
	return basePath, cleanPath[len(basePath):]
}

type FoodHandler struct {
	mangoHandler     *MangoHandler
	pineappleHandler *PineappleHandler
}

func (handler *FoodHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	basePath, nextPath := getPaths(request.URL.Path)

	if basePath == "" {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(writer).Encode(map[string]interface{}{"message": "wrong http method"})
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(map[string]interface{}{"message": "home page"})
		return
	}

	request.URL.Path = nextPath
	switch basePath {
	case "getMango":
		handler.mangoHandler.ServeHTTP(writer, request)
	case "postPineapple":
		handler.pineappleHandler.ServeHTTP(writer, request)
	default:
		writer.WriteHeader(http.StatusNotFound)
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(map[string]interface{}{"message": "route matches no host"})
	}
}

type PineappleHandler struct{}

func (handler *PineappleHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, nextPath := getPaths(request.URL.Path)
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(map[string]interface{}{"message": "welcome to pineapple page"})
}

type MangoHandler struct{}

func (handler *MangoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, _ = getPaths(request.URL.Path)
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(map[string]interface{}{"message": "welcome to mango page"})
}

func main() {
	foodHandler := FoodHandler{}
	server := http.Server{
		Handler:      foodHandler,
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
