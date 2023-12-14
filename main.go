package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	models.InitMySQL()
	utils.InitRedis()

	r := router.Router() // router.Router()
	r.Run(":8081")       // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
