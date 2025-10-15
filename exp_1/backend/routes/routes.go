package routes

import (
	"MES/exp1/controller"
	"MES/exp1/settings"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	g := r.Group("/api")
	hardwareGroup := g.Group("/hardware")
	{
		hardwareGroup.POST("/create", controller.CreateHardwareView)
		hardwareGroup.GET("/query", controller.ListHardwareView)
		hardwareGroup.PUT("/update", controller.UpdateHardwareView)
		hardwareGroup.DELETE("/delete", controller.DeleteHardwareView)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})

	_ = r.Run(fmt.Sprintf(":%d", settings.Conf.SystemConf.Port))
}
