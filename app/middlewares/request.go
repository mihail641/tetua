package middlewares

import (
	"fmt"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ngocphuongnb/tetua/app/logger"
	"github.com/ngocphuongnb/tetua/app/server"
)

const HEADER_REQUEST_ID = "X-Request-Id"

func RequestID(c server.Context) error {
	requestId := c.Header(HEADER_REQUEST_ID)

	if requestId == "" {
		requestId = uuid.NewString()
	}

	c.Locals("request_id", requestId)
	c.Header(HEADER_REQUEST_ID, requestId)

	return c.Next()
}

func RequestLog(c server.Context) error {
	start := time.Now()
	err := c.Next()
	latency := time.Since(start).Round(time.Millisecond)
	logContext := logger.Context{
		"latency": latency.String(),
		"status":  c.Response().StatusCode(),
		"method":  c.Method(),
		"path":    c.Path(),
		"ip":      c.IP(),
	}

	if err != nil {
		logContext["error"] = err.Error()
	}

	c.Logger().Info("Request completed", logContext)
	return nil
}

// RequestSize проверка является ли тело запроса от клиента больше разрешенного значения
func RequestSize(c server.Context) error {
	bodyRequest := c.Body()
	requestSize := len(bodyRequest)

	url := c.Path()
	if url == "/files/upload/image" {
		return c.Next()
	}
	err := fmt.Sprintf("Error saving file, file larger than the allowed value of %d Mb", config.REQUEST_LIMIT)
	if requestSize > config.REQUEST_LIMIT*1024*1024 {
		// Размер тела запроса превышает 4 Мбайта, отправляем пользователю ошибку.
		return c.Status(http.StatusRequestEntityTooLarge).Json(entities.Map{
			"error": err})

	}
	return c.Next()
}
