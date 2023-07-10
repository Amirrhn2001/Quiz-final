package domain

import (
	"github.com/arzesh-co/arzesh-common/API"
	"github.com/arzesh-co/arzesh-common/errors"
)

type Service interface {
	InsertQuestionsService(request *API.InfoRequest, questions *Questions) (*Questions, *errors.ResponseErrors)
	InsertQuizService(request *API.InfoRequest, quiz *Quiz) (*Quiz, *errors.ResponseErrors)
	InsertQuizQuestionsService(request *API.InfoRequest, quizQuestions *QuizQuestions) (*QuizQuestions, *errors.ResponseErrors)

	FindQuestionService(request *API.InfoRequest, questionId string) (*Question, *errors.ResponseErrors)
	FindQuestionsService(request *API.InfoRequest, filter *FilterQuestion) (*Questions, *errors.ResponseErrors)
	FindQuizService(request *API.InfoRequest, quizId string) (*Quiz, *errors.ResponseErrors)
	FindQuizzesService(request *API.InfoRequest, filter *FilterQuiz) (*Quizzes, *errors.ResponseErrors)
	FindQuizQuestionsService(request *API.InfoRequest, quizUuid string) ([]*QuestionResponse, *errors.ResponseErrors)

	UpdateQuestionService(request *API.InfoRequest, question *Question) (*Question, *errors.ResponseErrors)
	UpdateQuizService(request *API.InfoRequest, quiz *Quiz) (*Quiz, *errors.ResponseErrors)

	DeleteQuestionService(request *API.InfoRequest, id string) *errors.ResponseErrors
	DeleteQuizService(request *API.InfoRequest, id string) *errors.ResponseErrors
	DeleteQuizQuestionService(request *API.InfoRequest, id string) *errors.ResponseErrors
	// UpdateQuestionService()
	// UpdateQuizService()
	// UpdateQuizQuestionService()
}
