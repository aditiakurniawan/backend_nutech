package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/mysql"
	"nutech/repositories"

	"github.com/gorilla/mux"
)

func BarangRoutes(r *mux.Router) {
	barangRepository := repositories.RepositoryBarang(mysql.DB)
	h := handlers.HandlerBarang(barangRepository)

	r.HandleFunc("/barangs", h.FindBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", h.GetBarang).Methods("GET")
	r.HandleFunc("/barang", middleware.Auth(middleware.UploadFile(h.CreateBarang))).Methods("POST")
	r.HandleFunc("/barang/{id}", h.UpdateBarang).Methods("PATCH")
	r.HandleFunc("/barang/{id}", h.DeleteBarang).Methods("DELETE")
}
