package middleware

import (
	"net/http"
	"strconv"
	"github.com/Jlenin5/facil_backend/internal/usecase"
)

type SubscriptionMiddleware struct {
	SubscriptionUC *usecase.SubscriptionUseCase
	UserUC         *usecase.UserUseCase
}

func NewSubscriptionMiddleware(subscriptionUC *usecase.SubscriptionUseCase, userUC *usecase.UserUseCase) *SubscriptionMiddleware {
	return &SubscriptionMiddleware{
		SubscriptionUC: subscriptionUC,
		UserUC:         userUC,
	}
}

func (sm *SubscriptionMiddleware) ValidateSubscription(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del usuario desde el header
		userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		// Obtener el usuario para verificar su company_id
		user, err := sm.UserUC.GetUserById(userID)
		if err != nil {
			http.Error(w, "Error retrieving user information", http.StatusInternalServerError)
			return
		}

		// Si el usuario no está vinculado a una empresa (es super_admin), permitir el acceso
		if !user.Company_Id.Valid {
			next.ServeHTTP(w, r)
			return
		}

		// Si el usuario está vinculado a una empresa, validar la suscripción
		valid, err := sm.SubscriptionUC.IsSubscriptionValid(int(*user.Company_Id.Int))
		if err != nil {
			http.Error(w, "Error validating subscription", http.StatusInternalServerError)
			return
		}

		if !valid {
			http.Error(w, "Subscription is not valid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}