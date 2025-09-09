package main

import (
	"go-logging/internal/handler"
	"net/http"
	"net/http/httptest"

	"go.uber.org/zap"
)


func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Debug("Đây là log DEBUG")
	logger.Info("Đây là log INFO")
	logger.Warn("Đây là log WARN")
	logger.Error("Đây là log ERROR")
	// logger.Fatal("Đây là log FATAL") // Dừng chương trình nếu bật lên
	// logger.Panic("Đây là log PANIC") // Panic nếu bật lên
	logger.Info("Demo logging với zap - không khởi động server.")

	// Demo gọi trực tiếp handler (log của handler vẫn dùng zerolog)
	req := httptest.NewRequest(http.MethodGet, "/user?id=1", nil)
	w := httptest.NewRecorder()
	handler.GetUserHandler(w, req)
	logger.Info("Kết quả trả về từ handler.GetUserHandler", zap.String("handler_response", w.Body.String()))

}