package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ArtoIi/To-Do-List-API/internal/application"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/db"
	"github.com/ArtoIi/To-Do-List-API/internal/interfaces"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis do sistema")
	}

	dns := os.Getenv("MYSQL_DSN")
	repo, err := db.NewMySQLRepository(dns)
	if err != nil {
		log.Fatalf("Erro ao conectar no MySQL: %v", err)
	}
	fmt.Println("Conectado ao MySQL com sucesso!")
	service := application.NewToDoService(repo)
	handler := interfaces.NewToDoHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler.Register)

	port := ":8080"

	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Printf("Servidor rodando na porta %s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
