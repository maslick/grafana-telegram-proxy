package main

func main() {
	server := RestController{Service: &TelegramSender{}}
	server.Start()
}
