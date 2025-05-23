package main

import (
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
)

func main() {
	main := logger.New("main")
	main.Info("Hello World")
	main.Debug("Hello World")
	main.Error("Hello World")
	main.Warn("Hello World")
	main.Fatal("Hello World")
}
