package mw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/tasker-back/internal/config"
)

func CORSMiddleware(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Origins},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Set-Cookie", "Cookie"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}

func ProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authServiceURL := config.AuthUrl
		authToken := c.GetHeader("Authorization")
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, authServiceURL, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req.Header.Set("Authorization", authToken)

		userCookie, _ := c.Cookie("jwt")
		cookie := &http.Cookie{
			Name:  "jwt",
			Value: userCookie,
		}
		req.AddCookie(cookie)

		res, err := client.Do(req)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if res.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(res.StatusCode, gin.H{
				"message": fmt.Sprintf("Authorization failed: %s", string(body)),
			})
			return
		}

		// Получение заголовка Authorization из ответа и установка его в ответ сервера
		authHeader := res.Header.Get("Authorization")
		if authHeader != "" {
			c.Header("Authorization", authHeader)
		}

		// Получение cookie из ответа и установка ее в ответ сервера
		cookieHeader := res.Header.Get("Set-Cookie")
		if cookieHeader != "" {
			c.Header("Set-Cookie", cookieHeader)
		}
	}
}
