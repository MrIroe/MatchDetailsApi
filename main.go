package main

import (
	mongo "matchDetailsApi/mongo"
)

func main() {
	mongo.InitMongoRepo()

	forever := make(chan bool)
	go StartWeb()
	<-forever
}
