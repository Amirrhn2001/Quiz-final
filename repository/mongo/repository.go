package mongo

import (
	"context"
	"os"
	"service-template/domain"
	"time"

	"github.com/arzesh-co/arzesh-common/API"
	"github.com/arzesh-co/arzesh-common/constant"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	var cred options.Credential
	cred.Username = os.Getenv("dbUserName")
	cred.Password = os.Getenv("dbPassword")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (domain.Repository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}

func (r mongoRepository) FindAccountRepository(uuid string) (*domain.Account, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.AccountsCollection)
	result := collection.FindOne(ctx, bson.D{{Key: "acnt_uuid", Value: uuid}})
	var account *domain.Account
	err := result.Decode(&account)
	// if result != nil || account == (&domain.Account{}) {
	// 	return nil, constant.NotFoundErr
	// }
	if err != nil {
		return nil, constant.UnknownErr
	}
	return account, ""
}

func (r mongoRepository) InsertQuestionsRepository(questions []interface{}) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuestionsCollection)
	_, err := collection.InsertMany(ctx, questions)
	if err != nil {
		return constant.UnknownErr
	}
	return ""
}

func (r mongoRepository) InsertQuizRepository(quiz *domain.Quiz) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizzesCollection)
	_, err := collection.InsertOne(ctx, quiz)
	if err != nil {
		return constant.UnknownErr
	}

	return ""
}

func (r mongoRepository) InsertQuizQuestionsRepository(questions []interface{}) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizQuestionsCollection)
	_, err := collection.InsertMany(ctx, questions)
	if err != nil {
		return constant.UnknownErr
	}

	return ""
}

func (r mongoRepository) FindQuestionRepository(request *API.InfoRequest) (*domain.Question, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuestionsCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "GetOneQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return nil, constant.UnknownErr
	}
	question := &domain.Question{}
	err = collection.FindOne(ctx, filter).Decode(question)
	if err != nil {
		return nil, constant.NotFoundErr
	}

	return question, ""
}

func (r mongoRepository) FindQuestionsRepository(request *API.InfoRequest, filter any) (*domain.Questions, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuestionsCollection)
	cursor, err := collection.Find(ctx, filter) 
	questions := &domain.Questions{}
	for cursor.Next(ctx) {
		question := &domain.Question{}
		err := cursor.Decode(question)
		if err != nil {
			return nil, constant.UnknownErr
		}
		questions.Questions = append(questions.Questions, *question)
	}
	if err != nil {
		return nil, constant.NotFoundErr
	}

	return questions, ""
}

func (r mongoRepository) FindQuizRepository(request *API.InfoRequest) (*domain.Quiz, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizzesCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "GetOneQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return nil, constant.UnknownErr
	}
	quiz := &domain.Quiz{}
	err = collection.FindOne(ctx, filter).Decode(quiz)
	if err != nil {
		return nil, constant.NotFoundErr
	}
	

	return quiz, ""
}

func (r mongoRepository) FindQuizzesRepository(request *API.InfoRequest, filter any) (*domain.Quizzes, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizzesCollection)
	cursor, err := collection.Find(ctx, filter) 
	quizzes := &domain.Quizzes{}
	for cursor.Next(ctx) {
		quiz := &domain.Quiz{}
		err := cursor.Decode(quiz)
		if err != nil {
			return nil, constant.UnknownErr
		}
		quizzes.Quizzes = append(quizzes.Quizzes, *quiz)
	}
	if err != nil {
		return nil, constant.NotFoundErr
	}

	return quizzes, ""
}

func (r mongoRepository) FindQuizQuestionsRepository(request *API.InfoRequest) (*domain.QuizQuestions, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizQuestionsCollection)
	filter, err := request.MongoDbFilter(false)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "GetOneQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return nil, constant.UnknownErr
	}

	cursor, err := collection.Find(ctx, filter)
	quizQuestions := &domain.QuizQuestions{}
	for cursor.Next(ctx) {
		quizQuestion := &domain.QuizQuestion{}
		err := cursor.Decode(quizQuestion)
		if err != nil {
			return nil, constant.UnknownErr
		}
		quizQuestions.Questions = append(quizQuestions.Questions, *quizQuestion)
	}
	if err != nil {
		return nil, constant.NotFoundErr
	}

	return quizQuestions, ""
}

func (r mongoRepository) UpdateQuestionRepository(request *API.InfoRequest, question *domain.Question) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuestionsCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "GetOneQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return constant.UnknownErr
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": question})
	if err != nil {
		return constant.UpdateErr
	}

	return ""
}

func (r mongoRepository) UpdateQuizRepository(request *API.InfoRequest, quiz *domain.Quiz) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizzesCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "GetOneQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return constant.UnknownErr
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": quiz})
	if err != nil {
		return constant.UpdateErr
	}

	return ""
}

func (r mongoRepository) GetDomainsRepository(request *API.InfoRequest) ([]*domain.Question, int64, string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizQuestionsCollection)
	pipe, filterCount := request.PipeLineMongoDbAggregate(true)
	count, err := collection.CountDocuments(ctx, filterCount)
	if err != nil {
		return nil, 0, constant.NotFoundErr
	}
	res, err := collection.Aggregate(ctx, pipe)
	if err != nil {
		return nil, 0, constant.UnknownErr
	}
	var results []*domain.Question
	if err = res.All(ctx, &results); err != nil {
		return nil, 0, constant.UnknownErr
	}
	if err := res.Close(ctx); err != nil {
		return nil, 0, constant.UnknownErr
	}
	return results, count, ""
}

func (r mongoRepository) DeleteQuestionRepository(request *API.InfoRequest) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuestionsCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "DeleteQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return constant.UnknownErr
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{ "Status": -1 }})
	if err != nil {
		return constant.RemoveErr
	}

	return ""
}

func (r mongoRepository) DeleteQuizRepository(request *API.InfoRequest) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizzesCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "DeleteQuiz", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return constant.UnknownErr
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{ "Status": -1 }})
	if err != nil {
		return constant.RemoveErr
	}

	return ""
}

func (r mongoRepository) DeleteQuizQuestionRepository(request *API.InfoRequest) string {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.QuizQuestionsCollection)
	filter, err := request.MongoDbFilter(true)
	if err != nil {
		commonAttrs := []attribute.KeyValue{
			attribute.String("filter_error", err.Error()),
		}
		c, span := request.Tracer.Start(request.Ctx, "DeleteQuizQuestion", trace.WithAttributes(commonAttrs...))
		request.Ctx = c
		defer span.End()
		return constant.UnknownErr
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{ "Status": -1 }})
	if err != nil {
		return constant.RemoveErr
	}

	return ""
}

// func (r mongoRepository) FindQuestionsRepository(filter any) (*domain.Questions, string) {
// 	var questions *domain.Questions
// 	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
// 	defer cancel()
// 	collection := r.client.Database(r.database).Collection(domain.QuestionsCollection)
// 	cursor, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for cursor.Next(ctx) {
// 		var question *domain.Question
// 		err := cursor.Decode(&question)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Println(question)
// 		questions.Questions = append(questions.Questions, *question)
// 	}

// 	return questions, nil
// }

// func (r mongoRepository) FindQuizRepository(ownerId string) (*domain.Quiz, string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
// 	defer cancel()
// 	collection := r.client.Database(r.database).Collection(domain.QuizzesCollection)
// 	result := collection.FindOne(ctx, bson.M{"owner_id": ownerId})
// 	if result.Err() == mongo.ErrNoDocuments {
// 		return nil, result.Err()
// 	}

// 	quiz := &domain.Quiz{}
// 	err := result.Decode(quiz)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return quiz, nil
// }

// func (r mongoRepository) FindQuizQuestionsRepository(quizMCQId string) ([]domain.QuizQuestion, string) {
// 	var questions []domain.QuizQuestion
// 	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
// 	defer cancel()
// 	collection := r.client.Database(r.database).Collection(domain.QuizQuestionsCollection)
// 	cursor, err := collection.Find(ctx, bson.M{"quiz_mcq_id": quizMCQId})
// 	if err != nil {
// 		return nil, err
// 	}
// 	for cursor.Next(ctx) {
// 		var question domain.QuizQuestion
// 		err := cursor.Decode(&question)
// 		if err != nil {
// 			return nil, err
// 		}
// 		questions = append(questions, question)
// 	}
// 	// fmt.Println(questions)
// 	return questions, nil
// }

// func (r mongoRepository) UpdateQuestionRepository() {
	
// }