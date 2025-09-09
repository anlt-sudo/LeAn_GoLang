package logdemo

import (
	"go-logging/internal/handler"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

func LogWithZap() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Debug("Đây là log DEBUG (zap)")
	logger.Info("Đây là log INFO (zap)")
	logger.Warn("Đây là log WARN (zap)")
	logger.Error("Đây là log ERROR (zap)")
	logger.Info("Demo logging với zap - không khởi động server.")
	req := httptest.NewRequest(http.MethodGet, "/user?id=1", nil)
	w := httptest.NewRecorder()
	handler.GetUserHandler(w, req)
	logger.Info("Kết quả trả về từ handler.GetUserHandler", zap.String("handler_response", w.Body.String()))
}

func LogWithZerolog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Msg("Đây là log DEBUG (zerolog) - sẽ không hiển thị do level mặc định là Info.")
	log.Info().Msg("Đây là log INFO (zerolog)")
	log.Warn().Msg("Đây là log WARN (zerolog)")
	log.Error().Msg("Đây là log ERROR (zerolog)")
	log.Info().Msg("Demo logging với zerolog - không khởi động server.")
	req := httptest.NewRequest(http.MethodGet, "/user?id=1", nil)
	w := httptest.NewRecorder()
	handler.GetUserHandler(w, req)
	log.Info().Str("handler_response", w.Body.String()).Msg("Kết quả trả về từ handler.GetUserHandler")
}
