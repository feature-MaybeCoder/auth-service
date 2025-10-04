package httphandlers

import (
	usercase "backend/internal/application/user"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RegistrationHandler struct {
	registerUC *usercase.RegisterUserUseCase
}

func NewRegistrationHandler() *RegistrationHandler {
	return &RegistrationHandler{registerUC: usercase.NewRegisterUserUseCase()}
}

func (h *RegistrationHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email    string `json:"email" validate:"min=3,max=40"`
		Name     string `json:"name" validate:"min=3,max=40"`
		Password string `json:"password" validate:"min=3,max=40"`
	}

	v := validator.New()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	if err = v.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.registerUC.Execute(
		req.Email,
		req.Name,
		req.Password,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
