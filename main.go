package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	h "service-template/delivery/api"
	"service-template/domain"
	repo "service-template/repository/mongo"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	repo := ChoseRepository()
	service := domain.NewService(repo)
	handler := h.NewHandler(service)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/api/quiz/questions", handler.AddQuestions)
	r.POST("/api/quiz/quiz", handler.AddQuiz)
	r.POST("/api/quiz/quiz-questions", handler.AddQuizQuestions)
	r.GET("/api/quiz/question/:id", handler.GetQuestion)
	r.GET("/api/quiz/questions", handler.GetQuestions)
	r.GET("/api/quiz/quiz/:id", handler.GetQuiz)
	r.GET("/api/quiz/quizzes", handler.GetQuizzes)
	r.GET("/api/quiz/quiz-questions/:id", handler.GetQuizQuestions)
	r.PUT("/api/quiz/question", handler.UpdateQuestion)
	r.PUT("/api/quiz/quiz", handler.UpdateQuiz)
	r.DELETE("/api/quiz/question/:id", handler.DeleteQuestion)
	r.DELETE("/api/quiz/quiz/:id", handler.DeleteQuiz)
	r.DELETE("/api/quiz/quiz-question/:id", handler.DeleteQuizQuestion)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port", GoDotEnvVariable("port"))
		errs <- http.ListenAndServe(httpPort(), r)

	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(uuid.NewString())

	//////////////////////
	fmt.Printf("Terminated %s", <-errs)
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func httpPort() string {
	port := "8000"
	if GoDotEnvVariable("port") != "" {
		port = GoDotEnvVariable("port")
	}
	return fmt.Sprintf(":%s", port)
}

func ChoseRepository() domain.Repository {
	mongoURL := GoDotEnvVariable("mongoURL")
	mongodb := GoDotEnvVariable("mongodb")
	mongoTimeout, err := strconv.Atoi(GoDotEnvVariable("mongoTimeout"))
	if err != nil {
		log.Fatal(err)
	}
	repository, err := repo.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}
	return repository
}
