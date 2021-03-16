package http

import (
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/cnjack/throttle"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/goe2eds/services/http/controllers"
	"github.com/nugrohosam/goe2eds/services/http/exceptions"
	"github.com/spf13/viper"

	"github.com/nugrohosam/goe2eds/services/http/middlewares"
)

// Routes ...
var Routes *gin.Engine

// Serve using for listen to specific port
func Serve() error {
	Prepare()

	port := viper.GetString("app.port")
	if err := Routes.Run(":" + port); err != nil {
		return err
	}

	return nil
}

// Prepare ...
func Prepare() {

	Routes = gin.New()

	isDebug := viper.GetBool("debug")
	if !isDebug {
		Routes.Use(exceptions.Recovery500())
	}

	rateLimiterCount := viper.GetUint64("rate-limiter.count")
	rateLimiterTime := viper.GetInt("rate-limiter.time-in-minutes")
	Routes.Use(throttle.Policy(&throttle.Quota{
		Limit:  rateLimiterCount,
		Within: time.Duration(rateLimiterTime) * time.Minute,
	}))

	Routes.Static("assets", "./assets")
	Routes.Static("web", "./web")

	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	Routes.GET("test-throttle", func(c *gin.Context) {
		c.Writer.Write([]byte("hello world"))
	})

	// test-sentry
	Routes.GET("test-sentry", func(ctx *gin.Context) {
		panic("make panic test")
	})

	// v1
	v1 := Routes.Group("v1")
	v1.Use(middlewares.AuthJwt())

	// v1/auth
	key := v1.Group("key")
	key.Use(gzip.Gzip(gzip.DefaultCompression))
	{
		key.POST("", controllers.KeyHandlerCreate())
	}

	// v1/auth
	message := v1.Group("message")
	message.Use(gzip.Gzip(gzip.DefaultCompression))
	{
		message.POST("", controllers.MessageHandlerCreate())
		message.POST("verify", controllers.MessageHandlerVerify())
	}

	// v1/auth
	file := v1.Group("file")
	file.Use(gzip.Gzip(gzip.DefaultCompression))
	{
		file.POST("", controllers.MessageHandlerCreate())
		file.POST("verify", controllers.MessageHandlerVerify())
	}
}
