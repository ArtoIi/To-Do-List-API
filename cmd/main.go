package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	todoservice "github.com/ArtoIi/To-Do-List-API/internal/application/todo_service"
	userService "github.com/ArtoIi/To-Do-List-API/internal/application/user_service"
	todo_repo "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/repository/todo"
	userRepo "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/repository/user"
	"github.com/ArtoIi/To-Do-List-API/internal/interfaces"
	todoHandler "github.com/ArtoIi/To-Do-List-API/internal/interfaces/todo_handler"
	userHandler "github.com/ArtoIi/To-Do-List-API/internal/interfaces/user_handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis do sistema")
	}

	dns := os.Getenv("MYSQL_DSN")
	dbConn, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatalf("Erro ao abrir o servidor: %v", err)
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Erro ao se conectar ao servidor: %v", err)
	}
	defer dbConn.Close()
	fmt.Println("Conectado ao MySQL com sucesso!")

	mux := http.NewServeMux()

	//User
	userRepo := userRepo.NewUserRepository(dbConn)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	mux.HandleFunc("POST /user", userHandler.Register)
	mux.HandleFunc("GET /getEmail/{email}", userHandler.GetEmail)
	mux.HandleFunc("GET /getId/{id}", userHandler.GetId)
	mux.HandleFunc("PUT /user/{id}", userHandler.Update)
	mux.Handle("DELETE /user/{id}", interfaces.AuthMiddleware(http.HandlerFunc(userHandler.Delete)))
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.Handle("GET /identify", interfaces.AuthMiddleware(http.HandlerFunc(userHandler.Identify)))

	//Todo
	toDoRepo := todo_repo.NewUserRepository(dbConn)
	toDoService := todoservice.NewToDoService(toDoRepo)
	toDoHandler := todoHandler.NewToDoHandler(toDoService)
	mux.Handle("POST /todo", interfaces.AuthMiddleware(http.HandlerFunc(toDoHandler.Post)))

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
