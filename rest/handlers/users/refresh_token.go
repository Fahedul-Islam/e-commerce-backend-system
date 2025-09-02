package users

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *UserHandler) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	newToken, err := h.GenerateRefreshToken(w, r)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]string{"refresh_token": newToken}, http.StatusOK)
}
