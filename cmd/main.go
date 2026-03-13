package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	userService "github.com/ArtoIi/To-Do-List-API/internal/application/user_service"
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
	service := userService.NewToDoService(repo)
	handler := interfaces.NewToDoHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user", handler.Register)
	mux.HandleFunc("GET /getEmail/{email}", handler.GetEmail)
	mux.HandleFunc("GET /getId/{id}", handler.GetId)
	mux.HandleFunc("PUT /user/{id}", handler.Update)
	mux.Handle("DELETE /user/{id}", interfaces.AuthMiddleware(http.HandlerFunc(handler.Delete)))
	mux.HandleFunc("POST /login", handler.Login)
	mux.Handle("GET /identify", interfaces.AuthMiddleware(http.HandlerFunc(handler.Identify)))

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
