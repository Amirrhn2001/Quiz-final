package domain

import (
	"fmt"

	"github.com/arzesh-co/arzesh-common/API"
	"github.com/arzesh-co/arzesh-common/constant"
	"github.com/arzesh-co/arzesh-common/errors"
	"github.com/google/uuid"
)

type domainService struct {
	Repo Repository
}

func NewService(
	repo Repository,
) *domainService {
	return &domainService{
		repo,
	}
}

func (s domainService) InsertQuestionsService(request *API.InfoRequest, questions *Questions) (*Questions, *errors.ResponseErrors) {
	var validQuestions []interface{}
	var invalidQuestions Questions
	var responseValidQuestions Questions

	accountUUID := GetStringValueFromToken(request.ClientToken, "quiz")
	account, err := s.Repo.FindAccountRepository(accountUUID)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	for _, question := range questions.Questions {
		question.ID = uuid.NewString()
		question.Uuid = accountUUID
		question.CreatedAt, _ = uuid.NameSpaceDNS.Time().UnixTime()
		question.CreatedBy = GetStringValueFromToken(request.UserToken, "user_uuid")

		if account.Config.HasAudit {
			question.Status = 1
		}
		if len(question.Tags) < account.Config.MinTag || len(question.Choices) != 4 {
			invalidQuestions.Questions = append(invalidQuestions.Questions, question)
			fmt.Println("push to invalid array")
		} else {
			validQuestions = append(validQuestions, question)
			responseValidQuestions.Questions = append(responseValidQuestions.Questions, question)
		}
	}

	if len(validQuestions) == 0 {
		return &invalidQuestions, errors.New(constant.InvalidateInputsErr, request.Lang, "quiz", "", nil)
	}

	err = s.Repo.InsertQuestionsRepository(validQuestions)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	if len(invalidQuestions.Questions) > 0 {
		return &invalidQuestions, errors.New(constant.InvalidateInputsErr, request.Lang, "quiz", "", nil)
	}

	return &responseValidQuestions, nil
}

func (s domainService) InsertQuizService(request *API.InfoRequest, quiz *Quiz) (*Quiz, *errors.ResponseErrors) {
	quiz.ID = uuid.NewString()
	quiz.Uuid = GetStringValueFromToken(request.ClientToken, "quiz")
	quiz.QuizUuid = uuid.NewString()
	quiz.CreatedAt, _ = uuid.NameSpaceDNS.Time().UnixTime()
	quiz.CreatedBy = GetStringValueFromToken(request.UserToken, "user_uuid")
	quiz.Status = 1

	err := s.Repo.InsertQuizRepository(quiz)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return quiz, nil
}

func (s domainService) InsertQuizQuestionsService(request *API.InfoRequest, quizQuestions *QuizQuestions) (*QuizQuestions, *errors.ResponseErrors) {
	var questions []interface{}
	var responseQuizQuestions QuizQuestions

	for _, quizQuestion := range quizQuestions.Questions {
		quizQuestion.ID = uuid.NewString()
		quizQuestion.Uuid = GetStringValueFromToken(request.ClientToken, "quiz")
		quizQuestion.CreatedAt, _ = uuid.NameSpaceDNS.Time().UnixTime()
		quizQuestion.CreatedBy = GetStringValueFromToken(request.UserToken, "user_uuid")
		quizQuestion.Status = 1
		questions = append(questions, quizQuestion)
		responseQuizQuestions.Questions = append(responseQuizQuestions.Questions, quizQuestion)
	}

	err := s.Repo.InsertQuizQuestionsRepository(questions)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return &responseQuizQuestions, nil
}

func (s domainService) FindQuestionService(request *API.InfoRequest, questionId string) (*Question, *errors.ResponseErrors) {
	request.ServiceFilter = []API.Filter{
		{Condition: questionId, Label: "_id", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}

	question, err := s.Repo.FindQuestionRepository(request)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return question, nil
}

func (s domainService) FindQuestionsService(request *API.InfoRequest, filter *FilterQuestion) (*Questions, *errors.ResponseErrors) {
	f := CreateFilterQuestion(filter)
	questions, err := s.Repo.FindQuestionsRepository(request, f)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return questions, nil
}

func (s domainService) FindQuizService(request *API.InfoRequest, quizId string) (*Quiz, *errors.ResponseErrors) {
	request.ServiceFilter = []API.Filter{
		{Condition: quizId, Label: "quiz_uuid", Operation: "="},
		{Condition: 2, Label: "participation_type", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}

	quiz, err := s.Repo.FindQuizRepository(request) 
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return quiz, nil
}

func (s domainService) FindQuizzesService(request *API.InfoRequest, filter *FilterQuiz) (*Quizzes, *errors.ResponseErrors) {
	f := CreateFilterQuiz(filter)
	quizzes, err := s.Repo.FindQuizzesRepository(request, f) 
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return quizzes, nil
}

func (s domainService) FindQuizQuestionsService(request *API.InfoRequest, quizUuid string) ([]*QuestionResponse, *errors.ResponseErrors) {
	var questionsResponse []*QuestionResponse
	request.ServiceFilter = []API.Filter{
		{Condition: quizUuid, Label: "quiz_uuid", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}
	
	quizQuestions, err := s.Repo.FindQuizQuestionsRepository(request)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	fmt.Println("service", quizQuestions)
	for _, quizQuestion := range quizQuestions.Questions {
		questionResponse := &QuestionResponse{}
		request.ServiceFilter = []API.Filter{
			{Condition: quizQuestion.QuestionUuid, Label: "_id", Operation: "="},
			{Condition: -1, Label: "status", Operation: "!="},
		}
		question, err := s.Repo.FindQuestionRepository(request)
		if err != "" {
			return nil, errors.New(err, request.Lang, "quiz", "", nil)
		}
		questionResponse.Text = question.Text
		questionResponse.Choices = question.Choices
		questionResponse.Score = quizQuestion.Score
		questionResponse.Duration = quizQuestion.Duration

		questionsResponse = append(questionsResponse, questionResponse)
	}

	return questionsResponse, nil
}

func (s domainService) UpdateQuestionService(request *API.InfoRequest, question *Question) (*Question, *errors.ResponseErrors) {
	request.ServiceFilter = []API.Filter{
		{Condition: question.ID, Label: "_id", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}

	err := s.Repo.UpdateQuestionRepository(request, question)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return question, nil
}

func (s domainService) UpdateQuizService(request *API.InfoRequest, quiz *Quiz) (*Quiz, *errors.ResponseErrors) {
	request.ServiceFilter = []API.Filter{
		{Condition: quiz.ID, Label: "_id", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}

	err := s.Repo.UpdateQuizRepository(request, quiz)
	if err != "" {
		return nil, errors.New(err, request.Lang, "quiz", "", nil)
	}

	return quiz, nil
}

func (s domainService) DeleteQuestionService(request *API.InfoRequest, id string) *errors.ResponseErrors {
	request.ServiceFilter = []API.Filter{
		{Condition: id, Label: "question_uuid", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}
	
	quizQuestions, _ := s.Repo.FindQuizQuestionsRepository(request)
	if quizQuestions != nil  {
		return errors.New(constant.RemoveErr, request.Lang, "quiz", "", nil)
	}

	request.ServiceFilter = []API.Filter{
		{Condition: id, Label: "_id", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}

	err := s.Repo.DeleteQuestionRepository(request)
	if err != "" {
		return errors.New(err, request.Lang, "quiz", "", nil)
	}

	return nil
}

func (s domainService) DeleteQuizService(request *API.InfoRequest, id string) *errors.ResponseErrors {
	request.ServiceFilter = []API.Filter{
		{Condition: id, Label: "quiz_uuid", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}

	err := s.Repo.DeleteQuizRepository(request)
	if err != "" {
		return errors.New(err, request.Lang, "quiz", "", nil)
	}

	return nil
}

func (s domainService) DeleteQuizQuestionService(request *API.InfoRequest, id string) *errors.ResponseErrors {
	request.ServiceFilter = []API.Filter{
		{Condition: id, Label: "_id", Operation: "="},
		{Condition: -1, Label: "status", Operation: "!="},
	}
	
	err := s.Repo.DeleteQuizQuestionRepository(request)
	if err != "" {
		return errors.New(err, request.Lang, "quiz", "", nil)
	}
	
	return nil
}