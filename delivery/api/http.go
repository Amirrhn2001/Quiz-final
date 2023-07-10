package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/arzesh-co/arzesh-common/API"
	"github.com/arzesh-co/arzesh-common/constant"
	"github.com/arzesh-co/arzesh-common/errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"service-template/domain"
)

type Handler interface {
	AddQuestions(c *gin.Context)
	AddQuiz(c *gin.Context)
	AddQuizQuestions(c *gin.Context)
	GetQuestion(c *gin.Context)
	GetQuestions(c *gin.Context)
	GetQuiz(c *gin.Context)
	GetQuizzes(c *gin.Context)
	GetQuizQuestions(c *gin.Context) 
	UpdateQuestion(c *gin.Context)
	UpdateQuiz(c *gin.Context)
	DeleteQuestion(c *gin.Context)
	DeleteQuiz(c *gin.Context)
	DeleteQuizQuestion(c *gin.Context)
}

type handler struct {
	Service domain.Service
}

func NewHandler(Service domain.Service) Handler {
	return &handler{Service: Service}
}

func (h handler) AddQuestions(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	questions := &domain.Questions{}
	err = json.Unmarshal(body, questions)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.InsertQuestionsService(request, questions)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) AddQuiz(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	quiz := &domain.Quiz{}
	err = json.Unmarshal(body, quiz)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.InsertQuizService(request, quiz)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) AddQuizQuestions(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	quizQuestions := &domain.QuizQuestions{}
	err = json.Unmarshal(body, quizQuestions)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.InsertQuizQuestionsService(request, quizQuestions)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) GetQuestion(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	questionId := c.Param("id")
	data, Err := h.Service.FindQuestionService(request, questionId)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetQuestions(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"errors": err})
		return
	}

	filter := &domain.FilterQuestion{}
	err = json.Unmarshal(body, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"errors": err})
		return
	}

	data, Err := h.Service.FindQuestionsService(request, filter)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetQuiz(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	quizId := c.Param("id")
	data, Err := h.Service.FindQuizService(request, quizId)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetQuizzes(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		c.JSON(http.StatusUnauthorized, bson.M{"error": "Unauthorized"})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"errors": err})
		return
	}

	filter := &domain.FilterQuiz{} 
	err = json.Unmarshal(body, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"errors": err})
		return
	}

	data, Err := h.Service.FindQuizzesService(request, filter)
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetQuizQuestions(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	quizUuid := c.Param("id")

	data, Err := h.Service.FindQuizQuestionsService(request, quizUuid) 
	response := request.NewResponse(data, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) UpdateQuestion(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	question := &domain.Question{}
	err = json.Unmarshal(body, question)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.UpdateQuestionService(request, question) 
	response := request.NewResponse(data, domain.Updated, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) UpdateQuiz(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	quiz := &domain.Quiz{}
	err = json.Unmarshal(body, quiz)
	if err != nil {
		Err := errors.New(constant.InvalidateInputsErr, request.Lang, "action", "", nil)
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.UpdateQuizService(request, quiz) 
	response := request.NewResponse(data, domain.Updated, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) DeleteQuestion(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	quizUuid := c.Param("id")

	Err := h.Service.DeleteQuestionService(request, quizUuid) 
	response := request.NewResponse(nil, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) DeleteQuiz(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	quizUuid := c.Param("id")

	Err := h.Service.DeleteQuizService(request, quizUuid) 
	response := request.NewResponse(nil, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) DeleteQuizQuestion(c *gin.Context) {
	request := API.New(c.Request, "quiz", "1.0.0")
	isValid, _ := request.UserValidationRequest()
	if !isValid {
		Err := errors.New(constant.AccessErr, request.Lang, "domain", "", nil)
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	quizUuid := c.Param("id")

	Err := h.Service.DeleteQuizQuestionService(request, quizUuid) //
	response := request.NewResponse(nil, domain.Inserted, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

// func (h handler) GetQuiz(c *gin.Context) {
// 	request := API.New(c.Request, "quiz", "1.0.0")
// 	isValid, _ := request.UserValidationRequest()
// 	if !isValid {
// 		c.JSON(http.StatusUnauthorized, bson.M{"error": "Unauthorized"})
// 		return
// 	}
// 	quiz, err := h.Service.FindQuizService(request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, quiz)
// }

// func (h handler) UpdateQuestion(c *gin.Context) {
// 	request := API.New(c.Request, "quiz", "1.0.0")
// 	isValid, _ := request.UserValidationRequest()
// 	if !isValid {
// 		c.JSON(http.StatusUnauthorized, bson.M{"error": "Unauthorized"})
// 		return
// 	}

// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, bson.M{"error": err.Error()})
// 		return
// 	}

// 	question := &domain.Question{}
// 	err = json.Unmarshal(body, question)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, bson.M{"error": err.Error()})
// 		return
// 	}

// }
// func (h handler) UpdateQuiz(c *gin.Context) {
// 	request := API.New(c.Request, "quiz", "1.0.0")
// 	isValid, _ := request.UserValidationRequest()
// 	if !isValid {
// 		c.JSON(http.StatusUnauthorized, bson.M{"error": "Unauthorized"})
// 		return
// 	}

// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, bson.M{"error": err.Error()})
// 		return
// 	}

// 	quiz := &domain.Quiz{}
// 	err = json.Unmarshal(body, quiz)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, bson.M{"error": err.Error()})
// 		return
// 	}
// }
// func (h handler) UpdateQuizQuestion(c *gin.Context) {
// 	request := API.New(c.Request, "quiz", "1.0.0")
// 	isValid, _ := request.UserValidationRequest()
// 	if !isValid {
// 		c.JSON(http.StatusUnauthorized, bson.M{"error": "Unauthorized"})
// 		return
// 	}

// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, bson.M{"error": err.Error()})
// 		return
// 	}

// 	quizQuestion := &domain.QuizQuestion{}
// 	err = json.Unmarshal(body, quizQuestion)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, bson.M{"error": err.Error()})
// 		return
// 	}
// }
