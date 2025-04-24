package routes

import (
	"github.com/Jlenin5/facil_backend/internal/repository"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func PublicRoutes(router *mux.Router, db *sqlx.DB) {
	saleUC := usecase.NewSaleUseCase(repository.NewSaleRepository(db))
	saleHandler := handlers.NewSaleHandler(saleUC)
	router.HandleFunc("/sales/pdf/document-{bill}", saleHandler.OpenPDF).Methods("GET")
	router.HandleFunc("/sales/export-excel", saleHandler.ExportExcel).Methods("POST")

	productUC := usecase.NewProductUseCase(repository.NewProductRepository(db))
	productHandler := handlers.NewProductHandler(productUC)
	router.HandleFunc("/products/import-excel", productHandler.ImportExcel).Methods("POST")
	router.HandleFunc("/products/export-excel", productHandler.ExportExcel).Methods("POST")
}