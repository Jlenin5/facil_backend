package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	SignIn(email, password string) (*domain.LoginResponse, error)
	SignUp(username, email, password string) (*domain.LoginResponse, error)
	ValidateRefreshToken(refreshToken string) (*domain.Users, error)
}

type DauthUseCase struct {
	authRepo      repository.AuthRepository
	jwtKey        []byte
	accessExp     time.Duration
}

func NewAuthUseCase(authRepo repository.AuthRepository, jwtKey []byte) *DauthUseCase {
	return &DauthUseCase{
		authRepo:  authRepo,
		jwtKey:   jwtKey,
		accessExp: 24 * time.Hour,
	}
}

func (uc *DauthUseCase) SignIn(email, password string) (*domain.LoginResponse, error) {
	user, err := uc.authRepo.FindUserByTypeChar("email", email)
	if err != nil || user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	// Transformar usuario antes de generar el token
	transformedUser, err := transformUser(user)
	if err != nil {
		return nil, err
	}

	// Generar el token con todos los datos del usuario
	token, err := uc.generateToken(transformedUser)
	if err != nil {
		return nil, errors.New("no se pudo generar el token")
	}

	return &domain.LoginResponse{
		User:  transformedUser,
		Token: token,
	}, nil
}

func (uc *DauthUseCase) SignUp(username, email, password string) (*domain.LoginResponse, error) {
	// Verificar si ya existe un usuario con el mismo correo electrónico
	existingUser, err := uc.authRepo.FindUserByTypeChar("email", email)
	if err == nil && existingUser != nil {
		return nil, errors.New("el correo electrónico ya está registrado")
	}

	// Verificar si ya existe un usuario con el mismo username
	existingUser, err = uc.authRepo.FindUserByTypeChar("username", username)
	if err == nil && existingUser != nil {
		return nil, errors.New("el nombre de usuario ya está en uso")
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("no se pudo hashear la contraseña")
	}

	// Crear el nuevo usuario
	newUser := &domain.Users{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Guardar el nuevo usuario en la base de datos
	createdUser, err := uc.authRepo.CreateUser(newUser)
	if err != nil {
		return nil, errors.New("no se pudo crear el usuario: " + err.Error())
	}

	// Transformar usuario para token
	transformedUser, err := transformUser(createdUser)
	if err != nil {
		return nil, err
	}

	// Generar token
	token, err := uc.generateToken(transformedUser)
	if err != nil {
		return nil, errors.New("no se pudo generar el token")
	}

	return &domain.LoginResponse{
		User:  transformedUser,
		Token: token,
	}, nil
}

func (uc *DauthUseCase) ValidateRefreshToken(refreshToken string) (map[string]interface{}, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return uc.jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("refresh token inválido o expirado")
	}

	userId, ok := (*claims)["sub"].(float64)
	if !ok {
		return nil, errors.New("el token no contiene un Id de usuario válido")
	}

	user, err := uc.authRepo.FindUserByTypeChar("id", fmt.Sprintf("%d", int(userId)))
	if err != nil || user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	// Transformar usuario
	transformedUser, err := transformUser(user)
	if err != nil {
		return nil, err
	}

	return transformedUser, nil
}

func (uc *DauthUseCase) generateToken(userData map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userData["id"], // Incluir el user_id en el claim "sub"
		"user": userData,       // Incluir todos los datos del usuario en el claim "user"
		"exp":  time.Now().Add(uc.accessExp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(uc.jwtKey)
}

// Función auxiliar para transformar un usuario en la estructura deseada
func transformUser(user *domain.Users) (map[string]interface{}, error) {
	settings, err := parseJSON(user.Settings)
	if err != nil {
		return nil, err
	}

	shortcuts, err := parseJSON(user.Shortcuts)
	if err != nil {
		return nil, err
	}

	// Transformar company
	var company map[string]interface{}
	if user.Company != nil {
		company = map[string]interface{}{
			"id":              user.Company.Id,
			"name":            user.Company.Name,
		}
	}
	
	// Transformar employee
	var employee map[string]interface{}
	if user.Employee != nil {
		employee = map[string]interface{}{
			"id":              user.Employee.Id,
			"names":      user.Employee.Names,
			"surname":         user.Employee.Surname,
			"second_surname":  user.Employee.Second_Surname,
			"document_number": user.Employee.Document_Number,
			"warehouse_id":    user.Employee.Warehouse_Id,
			"status":          user.Employee.Status,
		}
	}

	// Transformar role
	var role map[string]interface{}
	if user.Role != nil {
		role = map[string]interface{}{
			"id":          user.Role.Id,
			"name":        user.Role.Name,
			"description": user.Role.Description,
		}
	}

	// Estructura transformada
	return map[string]interface{}{
		"id":          user.Id,
		"company_id":  user.Company_Id,
		"company":     company,
		"role_id":     user.Role_Id,
		"role":        role,
		"username":    user.Username,
		"employee_id": user.Employee_Id,
		"employee":    employee,
		"avatar":      user.Avatar,
		"email":       user.Email,
		"subscription":user.Subscription,
		"settings":    settings,
		"shortcuts":   shortcuts,
		"status":      user.Status,
	}, nil
}

func parseJSON(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case nil:
		return nil, nil
	case []byte:
		var result interface{}
		if err := json.Unmarshal(v, &result); err != nil {
			return nil, fmt.Errorf("error al decodificar JSON: %v", err)
		}
		return result, nil
	case string:
		var result interface{}
		if err := json.Unmarshal([]byte(v), &result); err != nil {
			return nil, fmt.Errorf("error al decodificar string JSON: %v", err)
		}
		return result, nil
	default:
		// Ya es un objeto Go (por ejemplo, un map o slice)
		return v, nil
	}
}