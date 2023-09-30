package router

import (
	"github.com/gorilla/mux"
	"github.com/techswarn/host/middleware"
)

func Router() *mux.Router{
    router := mux.NewRouter()
    // router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.CreateStocks).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllstock).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}