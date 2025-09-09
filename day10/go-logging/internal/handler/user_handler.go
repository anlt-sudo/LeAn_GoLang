package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	hlog := log.With().
		Str("handler", "GetUserHandler").
		Logger()

	hlog.Info().Msg("Bắt đầu xử lý request lấy thông tin người dùng...")

	time.Sleep(50 * time.Millisecond)
	userID := r.URL.Query().Get("id")
	if userID == "" {
		hlog.Warn().Msg("Request thiếu 'id' query parameter.")
		http.Error(w, "Missing user id", http.StatusBadRequest)
		return
	}

	if userID == "999" {
		dbErr := fmt.Errorf("user with id %s not found in database", userID)
		hlog.Error().Err(dbErr).Str("userID", userID).Msg("Lỗi khi truy vấn database.")
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	user := User{ID: 123, Name: "John Doe"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	hlog.Info().
		Str("userID", userID).
		Interface("response", user).
		Msg("Xử lý request thành công.")
}