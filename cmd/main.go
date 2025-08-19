package main

import (
	"net/http"

	"github.com/Jlenin5/facil_backend/internal/repository"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/middleware"
	"github.com/Jlenin5/facil_backend/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Conectar a la base de datos
	db := repository.ConnectDB()

	jwtKey := []byte("mi_llave_secreta")
	// refreshJwtKey := []byte("mi_llave_secreta_refresh")

	// Repositorios
	authRepo := repository.NewAuthRepository(db)
	// subscriptionRepo := repository.NewSubscriptionRepository(db)
	// userRepo := repository.NewUserRepository(db)
	
	// UseCases
	authUC := usecase.NewAuthUseCase(authRepo, jwtKey)
	// subscriptionUC := usecase.NewSubscriptionUseCase(subscriptionRepo)
	// userUC := usecase.NewUserUseCase(userRepo)

	// Crear routers
	publicRouter := mux.NewRouter()
	routes.AuthRoutes(publicRouter, authUC)

	// Middleware de autenticación para rutas protegidas
	protectedRouter := publicRouter.PathPrefix("/api").Subrouter()

	// Rutas públicas (PDF y Export-Excel)
	routes.PublicRoutes(publicRouter, db)

	// Middleware para rutas protegidas
	protectedRouter.Use(middleware.JWTAuthMiddleware(jwtKey))
	// protectedRouter.Use(middleware.NewSubscriptionMiddleware(subscriptionUC, userUC).ValidateSubscription)
	routes.SetupProtectedRoutes(protectedRouter, db)

	// Configurar CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-User-ID"},
		AllowCredentials: true,
	})

	// Servir archivos estáticos
	publicRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploads"))))

	// Combinar los routers
	mainRouter := mux.NewRouter()
	mainRouter.Use(corsMiddleware.Handler)
	mainRouter.PathPrefix("/api").Handler(protectedRouter)
	mainRouter.PathPrefix("/").Handler(publicRouter)

	// Iniciar servidor
	http.ListenAndServe(":8081", mainRouter)
}