package submission

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var submissions *mongo.Collection

type submissionListItem struct {
	Id           string `json:"submission_id,omitempty" bson:"_id,omitempty"`
	AuthorId     string `json:"author_id,omitempty" bson:"author_id,omitempty"`
	ProblemId    string `json:"problem_id,omitempty" bson:"problem_id,omitempty"`
	CodeLanguage string `json:"code_language,omitempty" bson:"code_language,omitempty"`
	Status       string `json:"status,omitempty" bson:"status,omitempty"`
}

type submission struct {
	Id           string `json:"submission_id,omitempty" bson:"_id,omitempty"`
	AuthorId     string `json:"author_id,omitempty" bson:"author_id,omitempty"`
	ProblemId    string `json:"problem_id,omitempty" bson:"problem_id,omitempty"`
	CodeLanguage string `json:"code_language,omitempty" bson:"code_language,omitempty"`
	Status       string `json:"status,omitempty" bson:"status,omitempty"`
	CodeText     string `json:"code_text,omitempty" bson:"code_text,omitempty"`
}

func submit(ctx context.Context, s *submission) error {
	ret, err := submissions.InsertOne(ctx, s)
	s.Id = ret.InsertedID.(primitive.ObjectID).Hex()

	return err
}

type searchSubmissionRequest struct {
	AuthorId  string `json:"author_id,omitempty" bson:"author_id,omitempty"`
	ProblemId string `json:"problem_id,omitempty" bson:"problem_id,omitempty"`
}

func searchSubmissions(ctx context.Context, req *searchSubmissionRequest) ([]submissionListItem, error) {
	r := bson.D{}

	if req.AuthorId != "" {
		id, err := primitive.ObjectIDFromHex(req.AuthorId)
		if err != nil {
			return nil, err
		}
		r = append(r, bson.E{"author_id", id})
	}

	if req.ProblemId != "" {
		id, err := primitive.ObjectIDFromHex(req.ProblemId)
		if err != nil {
			return nil, err
		}
		r = append(r, bson.E{"problem_id", id})
	}

	data, err := submissions.Find(ctx, r,
		&options.FindOptions{Projection: bson.D{
			{"_id", 1},
			{"author_id", 1},
			{"problem_id", 1},
			{"code_language", 1},
			{"status", 1},
		}})

	if err != nil {
		return nil, err
	}

	var ls []submissionListItem
	for data.Next(ctx) {
		var item submissionListItem
		err = data.Decode(&item)
		if err != nil {
			return nil, err
		}
		ls = append(ls, item)
	}

	return ls, nil
}

type querySubmissionRequest struct {
	Id string `json:"submission_id,omitempty" bson:"_id,omitempty"`
}

func querySubmission(ctx context.Context, q *querySubmissionRequest) (*submission, error) {
	var resp submission
	err := submissions.FindOne(ctx, q).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
