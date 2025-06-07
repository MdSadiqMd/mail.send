package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/MdSadiqMd/mail.send/internal/models"
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
	"github.com/MdSadiqMd/mail.send/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func NewAuth(secret string) Auth {
	return Auth{
		Secret: secret,
	}
}

var authLogger = logger.New("AuthMiddleware")

func (auth Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) == 0 {
		authLogger.Error("password is required")
		return "", errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		authLogger.Error("failed to hash password: %v", err)
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}

func (auth Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || len(email) == 0 {
		authLogger.Error("failed to generate token")
		return "", errors.New("failed to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"expiry":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(auth.Secret))
	if err != nil {
		authLogger.Error("unable to sign the token: %v", err)
		return "", errors.New("unable to sign the token")
	}
	return tokenString, nil
}

func (auth Auth) VerifyPassword(password string, hashedPassword string) error {
	if len(password) == 0 || len(hashedPassword) == 0 {
		authLogger.Error("password is required")
		return errors.New("password is required")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		authLogger.Error("password does not match: %v", err)
		return errors.New("password does not match")
	}
	return nil
}

func (auth Auth) VerifyToken(token string) (models.User, error) {
	tokenHeader := strings.Split(token, " ")
	if len(tokenHeader) != 2 {
		authLogger.Error("invalid token")
		return models.User{}, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenHeader[1], claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(auth.Secret), nil
	})
	if err != nil {
		authLogger.Error("invalid token: %v", err)
		return models.User{}, errors.New("invalid token")
	}

	if claims["expiry"].(float64) < float64(time.Now().Unix()) {
		return models.User{}, errors.New("token expired")
	}

	user := models.User{
		Id:    uint(claims["user_id"].(float64)),
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}
	return user, nil
}

func (auth Auth) CurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			authLogger.Error("missing authorization header")
			utils.ErrorResponse(w, http.StatusUnauthorized, "Authorization header required", errors.New("missing authorization header"))
			return
		}

		user, err := auth.VerifyToken(authHeader)
		if err != nil || user.Id == 0 {
			authLogger.Error("unauthorized: %v", err)
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}

		ctx := utils.SetUserInContext(r.Context(), &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (auth Auth) GetCurrentUser(r *http.Request) (*models.User, bool) {
	user, err := utils.GetUserFromContext(r.Context())
	if err {
		authLogger.Error("user not found in context")
		return nil, false
	}
	return user, true
}

func (auth Auth) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			authLogger.Error("missing authorization header")
			utils.ErrorResponse(w, http.StatusUnauthorized, "Authorization header required", errors.New("missing authorization header"))
			return
		}

		user, err := auth.VerifyToken(authHeader)
		if err != nil || user.Id == 0 {
			authLogger.Error("unauthorized: %v", err)
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}

		ctx := utils.SetUserInContext(r.Context(), &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (auth Auth) AuthorizeSeller(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			authLogger.Error("missing authorization header")
			utils.ErrorResponse(w, http.StatusUnauthorized, "Authorization header required", errors.New("missing authorization header"))
			return
		}

		user, err := auth.VerifyToken(authHeader)
		if err != nil {
			authLogger.Error("unauthorized: %v", err)
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}

		if user.Id == 0 {
			authLogger.Error("invalid user")
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid user", errors.New("user ID is invalid"))
			return
		}

		if user.Role != "seller" {
			authLogger.Error("access denied")
			utils.ErrorResponse(w, http.StatusForbidden, "Access denied", errors.New("please login as a seller"))
			return
		}

		ctx := utils.SetUserInContext(r.Context(), &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (auth Auth) AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			authLogger.Error("missing authorization header")
			utils.ErrorResponse(w, http.StatusUnauthorized, "Authorization header required", errors.New("missing authorization header"))
			return
		}

		user, err := auth.VerifyToken(authHeader)
		if err != nil {
			authLogger.Error("unauthorized: %v", err)
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}

		if user.Id == 0 {
			authLogger.Error("invalid user")
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid user", errors.New("user ID is invalid"))
			return
		}

		if user.Role != "admin" {
			authLogger.Error("access denied")
			utils.ErrorResponse(w, http.StatusForbidden, "Access denied", errors.New("admin access required"))
			return
		}

		ctx := utils.SetUserInContext(r.Context(), &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (auth Auth) GenerateCode() (string, error) {
	code, err := utils.GenerateRandomString(6)
	if err != nil {
		authLogger.Error("failed to generate code: %v", err)
		return "", err
	}
	return code, nil
}
