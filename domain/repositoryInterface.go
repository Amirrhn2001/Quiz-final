package domain

import "github.com/arzesh-co/arzesh-common/API"

type Repository interface {
	FindAccountRepository(uuid string) (*Account, string)

	InsertQuestionsRepository(questions []interface{}) string
	InsertQuizRepository(quiz *Quiz) string
	InsertQuizQuestionsRepository(questions []interface{}) string

	FindQuestionRepository(request *API.InfoRequest) (*Question, string)
	FindQuestionsRepository(request *API.InfoRequest, filter any) (*Questions, string)
	FindQuizRepository(request *API.InfoRequest) (*Quiz, string)
	FindQuizzesRepository(request *API.InfoRequest, filter any) (*Quizzes, string)
	FindQuizQuestionsRepository(request *API.InfoRequest) (*QuizQuestions, string)

	UpdateQuestionRepository(request *API.InfoRequest, question *Question) string
	UpdateQuizRepository(request *API.InfoRequest, quiz *Quiz) string

	DeleteQuestionRepository(request *API.InfoRequest) string
	DeleteQuizRepository(request *API.InfoRequest) string
	DeleteQuizQuestionRepository(request *API.InfoRequest) string
	// FindQuestionsRepository(filter any) (*Questions, string)
	// FindQuizRepository(id string) (*Quiz, string)
	// FindQuizQuestionsRepository(quizUuid string) ([]QuizQuestion, string)

	// UpdateQuestionRepository()
	// UpdateQuizRepository()
	// UpdateQuizQuestionRepository()
}
