package main

import (
	"fmt"
	"briangriffey.com/injection"
	"github.com/gorilla/mux"
	"net/http"
)

type DataFetcher interface {
	FetchData() string
}

type MockDataFetcher struct {

}

func (m MockDataFetcher) FetchData() string {
	return "This is all mock data"
}

type RealDataFetcher struct {

}

func (r RealDataFetcher) FetchData() string {
	return "I'm actually going to fetch data here"
}

type AppConfiguration struct {
	Fetcher DataFetcher
}


type RandomHandler struct {
	DataFetcher DataFetcher `inject:"*"`
}

func (r RandomHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "text/plain")

	fetchedData := r.DataFetcher.FetchData()
	writer.Write([]byte(fetchedData))
}

func main() {
	//Create an app configuration that uses mock data
	appConfiguration := AppConfiguration{MockDataFetcher{}}
	handler := RandomHandler{}

	injector, error := injection.NewInjector(appConfiguration)

	if error != nil {
		fmt.Println(error)
		return
	}

	injector.Inject(&handler)

	router := mux.NewRouter()
	router.Handle("/", handler)
}

