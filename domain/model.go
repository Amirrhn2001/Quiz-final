package domain

// `json:"" bson:""`
type Account struct {
	Uuid      string `json:"acnt_uuid" bson:"acnt_uuid"`
	Title     string `json:"acnt_title" bson:"acnt_title"`
	OwnerUuid string `json:"owner_uuid" bson:"owner_uuid"`
	Config    Config `json:"config" bson:"config"`
	Status    int    `json:"status" bson:"status"` // 1 -1
}

type Question struct {
	ID              string   `json:"id" bson:"_id"`
	Uuid            string   `json:"acnt_uuid" bson:"acnt_uuid"`
	Text            string   `json:"text" bson:"text"`
	Image           string   `json:"image" bson:"image"`
	Choices         []Choice `json:"choices" bson:"choices"`
	Answer          int      `json:"answer" bson:"answer"`
	DifficultyLevel int      `json:"difficulty_level" bson:"difficulty_level"` //1.easy 2.intermediate 3.hard
	Tags            []Tag    `json:"tags" bson:"tags"`
	Refs            []Ref    `json:"refs" bson:"refs"`
	Owner           string   `json:"owner_uuid" bson:"owner_uuid"`
	CreatedAt       int64    `json:"created_at" bson:"created_at"`
	CreatedBy       string   `json:"created_by" bson:"created_by"`
	Status          int      `json:"status" bson:"status"` // 0.NonAudit 1.Audit -1.removed
}

type Quiz struct {
	ID                string     `json:"id" bson:"_id"`
	Uuid              string     `json:"acnt_uuid" bson:"acnt_uuid"`
	QuizUuid          string     `json:"quiz_uuid" bson:"quiz_uuid"`
	Title             string     `json:"title" bson:"title"`
	Description       string     `json:"description" bson:"description"`
	DurationType      int        `json:"duration_type" bson:"duration_type"` //1:notLimited, 2:totalDuration 3:perQuestionDuration
	Duration          int        `json:"duration" bson:"duration"`
	ParticipationType int        `json:"participation_type" bson:"participation_type"` //1:invitation  2:volunteering
	Audiences         []Audience `json:"audiences" bson:"audiences"`
	MinScore          int        `json:"min_score" bson:"min_score"`
	Tags              []Tag      `json:"tags" bson:"tags"`
	Refs              []Ref      `json:"refs" bson:"refs"`
	Owner             string     `json:"owner_uuid" bson:"owner_uuid"` // Party uuid
	CreatedAt         int64      `json:"created_at" bson:"created_at"`
	CreatedBy         string     `json:"created_by" bson:"created_by"`
	Status            int        `json:"status" bson:"status"` //1.active -1.remove
}

type QuizQuestion struct {
	ID           string `json:"_id" bson:"_id"`
	Uuid         string `json:"acnt_uuid" bson:"acnt_uuid"`
	QuizUuid     string `json:"quiz_uuid" bson:"quiz_uuid"`
	QuestionUuid string `json:"question_uuid" bson:"question_uuid"`
	Score        int    `json:"score" bson:"score"`
	Duration     int    `json:"duration" bson:"duration"`
	CreatedAt    int64  `json:"created_at" bson:"created_at"`
	CreatedBy    string `json:"created_by" bson:"created_by"`
	Status       int    `json:"status" bson:"status"` //1.active -1.remove
}

type Questions struct {
	Questions []Question `json:"questions" bson:"questions"`
}

type QuizQuestions struct {
	Questions []QuizQuestion `json:"questions" bson:"questions"`
}

type Quizzes struct {
	Quizzes []Quiz `json:"questions" bson:"questions"`
}

type Config struct {
	HasAudit bool `json:"has_audit" bson:"has_audit"`
	MinTag   int  `json:"min_tag" bson:"min_tag"`
}

type Choice struct {
	Index int    `json:"index" bson:"index"`
	Text  string `json:"text" bson:"text"`
}

type Tag struct {
	Uuid  string `json:"uuid" bson:"uuid"`
	Title Title  `json:"title" bson:"title"`
}

type Ref struct {
	RefType string `json:"ref_type" bson:"ref_type"`
	RegId   string `json:"ref_id" bson:"ref_id"`
}

type Audience struct {
	Owner string `json:"owner_uuid" bson:"owner_uuid"` // Party uuid
	Title Title  `json:"title" bson:"title"`
}

type Title struct {
	EN string `json:"en" bson:"en"`
	FA string `json:"fa" bson:"fa"`
}

type FilterQuestion struct {
	Question   string   `json:"text"`
	Tags       []string `json:"tags"`
	OwnerId    string   `json:"owner_uuid"`
	Difficulty int      `json:"difficulty"`
}

type FilterQuiz struct {
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Refs  []string `json:"refs"`
	Owner string   `json:"owner"`
}

type QuizResponse struct {
	Quiz      Quiz             `json:"quiz"`
	Questions QuestionResponse `json:"questions"`
}

type QuestionResponse struct {
	Text     string   `json:"text" bson:"text"`
	Choices  []Choice `json:"choices" bson:"choices"`
	Score    int      `json:"score" bson:"score"`
	Duration int      `json:"duration" bson:"duration"`
}

type QuestionsResponse struct {
	Questions []QuestionResponse `json:"questions" bson:"questions"`
}
