package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MdSadiqMd/mail.send/internal/models"
	"github.com/MdSadiqMd/mail.send/internal/services"
	dto "github.com/MdSadiqMd/mail.send/internal/types"
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
	"github.com/MdSadiqMd/mail.send/pkg/utils"
)

type UserHandler struct {
	service services.UserService
}

var handleLogger = logger.New("UserHandler")

func NewUserHandler(service services.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user dto.UserSignup
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleLogger.Error("Failed at Backend DTO while register: %v", err)
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed at Backend DTO", err)
		return
	}

	token, err := h.service.Signup(models.User{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: user.IsVerified,
	})
	if err != nil {
		handleLogger.Error("Error at Signup Auth: %v", err)
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error at Signup Auth", err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Signup successful", token)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleLogger.Error("Failed at Backend DTO while login: %v", err)
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed at Backend DTO", err)
		return
	}

	token, err := h.service.Login(user.Email, user.Password)
	if err != nil {
		handleLogger.Error("Error at Login Auth: %v", err)
		utils.ErrorResponse(w, http.StatusUnauthorized, "Error at Login Auth", err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Login successful", token)
}

func (h *UserHandler) GetVerificationCode(w http.ResponseWriter, r *http.Request) {
	_, ok := h.service.Auth.GetCurrentUser(r)
	if !ok {
		handleLogger.Error("user not found in context")
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", errors.New("user not found in context"))
		return
	}

	code, err := h.service.Auth.GenerateCode()
	if err != nil {
		handleLogger.Error("Failed to get verification code: %v", err)
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get verification code", err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "User got verification code successfully", code)
}

func (h *UserHandler) Verify(w http.ResponseWriter, r *http.Request) {
	user, _ := h.service.Auth.GetCurrentUser(r)
	if user == nil {
		handleLogger.Error("user not found in context")
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req dto.VerificationCodeInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleLogger.Error("Failed at Backend DTO while verify: %v", err)
		utils.ErrorResponse(w, http.StatusBadRequest, "Please provide a valid verification code", err)
		return
	}

	_, err = h.service.Auth.GenerateCode()
	if err != nil {
		handleLogger.Error("Failed to get verification code: %v", err)
		utils.ErrorResponse(w, http.StatusInternalServerError, "Verification failed", err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "User verified successfully", true)
}
