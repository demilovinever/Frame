package main

//поменять local.yml на .env
import (
	"Frame/internal/config"
	"fmt"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println("Config loaded:", cfg)
}

func setupLogger(env string) {

}
