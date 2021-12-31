package main

import (
	"fmt"
	// "log"
	"todo_app/app/models"
	"todo_app/app/controllers"
)

func main() {
	fmt.Println(models.Db)
	
	controllers.StertMainServer()
	
	
}
