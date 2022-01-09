package main

import (
	"fmt"
	"log"
)



func main() {
	var Apps = AppInfo{}
	Apps.Add(serveApp)
	Apps.Add(serverDebug)
	err := Apps.Run()
	if err != nil{
		fmt.Println(err)
		log.Fatal(err)
	}


}
