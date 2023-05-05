//go:build e2e

package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

// step 1. curl http://localhost:8081/user
// step 2. curl http://localhost:8081/users/1
// step 3. curl http://localhost:8081/users/2
// step 4. curl http://localhost:8082/metrics
func TestMiddlewareBuilder_Build(t *testing.T) {
	builder := MiddlewareBuilder{
		Namespace: "server",
		Subsystem: "web",
		Name:      "http_response",
	}

	r := gin.Default()
	r.Use(builder.Build())

	r.GET("/user", func(ctx *gin.Context) {
		val := rand.Intn(1000) + 1
		time.Sleep(time.Duration(val) * time.Millisecond)
		ctx.JSON(http.StatusOK, User{
			Name: "Tom",
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		val := rand.Intn(1000) + 1
		time.Sleep(time.Duration(val) * time.Millisecond)
		ctx.JSON(http.StatusOK, User{
			Name: "Jack",
		})
	})

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		t.Error(http.ListenAndServe(":8082", nil))
	}()

	t.Error(r.Run(":8081"))
}

type User struct {
	Name string
}
