package domain

import (
	"github.com/arzesh-co/arzesh-common/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateFilterQuestion(filter *FilterQuestion) bson.D {
	var query bson.D
	if filter.Question == "" && len(filter.Tags) == 0 && filter.OwnerId == "" && filter.Difficulty == 0 && len(filter.Tags) == 0{
		return query
	}
	if filter.Question != "" {
		query = append(query, bson.E{Key: "text", Value: filter.Question})
	}
	if len(filter.Tags) != 0 {
		query = append(query, bson.E{Key: "tags", Value: bson.D{
			{Key: "$elemMatch",
				Value: bson.D{
					{Key: "uuid",
						Value: bson.D{
							{Key: "$in",
								Value: filter.Tags,
							},
						},
					},
				},
			},
		}})
	}
	if filter.OwnerId != "" {
		query = append(query, bson.E{Key: "owner_uuid", Value: filter.OwnerId})
	}
	if filter.Difficulty != 0 {
		query = append(query, bson.E{Key: "difficulty", Value: filter.Difficulty})
	}
	return query
}

func CreateFilterQuiz(filter *FilterQuiz) bson.D {
	var query bson.D
	if filter.Owner == "" && len(filter.Tags) == 0 && filter.Title == "" && len(filter.Tags) == 0 && len(filter.Refs) != 0 {
		return query
	}
	if filter.Owner != "" {
		query = append(query, bson.E{Key: "owner_uuid", Value: filter.Owner})
	}
	if len(filter.Tags) != 0 {
		query = append(query, bson.E{Key: "tags", Value: bson.D{
			{Key: "$elemMatch",
				Value: bson.D{
					{Key: "uuid",
						Value: bson.D{
							{Key: "$in",
								Value: filter.Tags,
							},
						},
					},
				},
			},
		}})
	}
	if len(filter.Refs) != 0 {
		query = append(query, bson.E{Key: "refs", Value: bson.D{
			{Key: "$elemMatch",
				Value: bson.D{
					{Key: "uuid",
						Value: bson.D{
							{Key: "$in",
								Value: bson.A{
									filter.Tags,
								},
							},
						},
					},
				},
			},
		}})
	}
	if filter.Title != "" {
		query = append(query, bson.E{Key: "title", Value: filter.Title})
	}
	return query
}

func GetStringValueFromToken(token string, key string) (value string) {
	TokenAccount := jwt.FindValue(token, key)
	if TokenAccount == nil {
		return
	}
	value = TokenAccount.(string)
	return
}
