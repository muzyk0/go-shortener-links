package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger(level string) (*Logger, error) {
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	//устанавливаем синглтон
	//Log = zl

	return &Logger{
		log: zl,
	}, nil
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) NewRequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logFn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			responseData := &responseData{
				status: 0,
				size:   0,
			}
			lw := loggingResponseWriter{
				ResponseWriter: w, // встраиваем оригинальный http.ResponseWriter
				responseData:   responseData,
			}
			next.ServeHTTP(&lw, r) // внедряем реализацию http.ResponseWriter

			duration := time.Since(start)

			l.log.Info("HTTP request",
				zap.String("uri", r.RequestURI),
				zap.String("method", r.Method),
				zap.Int("status", responseData.status), // получаем перехваченный код статуса ответа
				zap.Duration("duration", duration),
				zap.Int("size", responseData.size), // получаем перехваченный размер ответа
			)
		}
		return http.HandlerFunc(logFn)
	}
}

type (
	// берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// добавляем реализацию http.ResponseWriter
	loggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}
