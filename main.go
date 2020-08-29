package main

import (
	"abhinav/ticket-service/routes"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var envVariables = map[string]interface{}{
	"port":   9000,
	"mode":   "test",
	"config": "config/config.json",
}

func main() {
	parseCLIArgs()
	initMain()
	router := gin.Default()
	routes.InitRoutes(router)
	err := router.Run(":" + strconv.Itoa(envVariables["port"].(int)))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func parseCLIArgs() {
	argList := os.Args[1:]
	var argMap = make(map[string]string)
	for i := 0; i < len(argList); i += 2 {
		argMap[argList[i]] = argList[i+1]
	}
	for k, v := range envVariables {
		if v1, ok := argMap[k]; ok {
			if reflect.TypeOf(v).Kind() == reflect.Int {
				x, err := strconv.Atoi(v1)
				if err != nil {
					panic(fmt.Sprintf("Incorrect Value :%s for ENV Variable: %s", v1, k))
				}
				envVariables[k] = x
			} else {
				envVariables[k] = v1
			}
		}
	}
}
