package main

import (
	"AS-UNIVERSITY/config"
	"AS-UNIVERSITY/route"
)

//MAIN FUNCTION
func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
