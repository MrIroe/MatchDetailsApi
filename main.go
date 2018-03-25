package main

func main() {

	forever := make(chan bool)
	go StartWeb()
	<-forever
}
