package problem

import (
	"OnlineJudge_Backend/util"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var problems *mongo.Collection

type testCase struct {
	Input  string `json:"input" bson:"input"`
	Output string `json:"output" bson:"output"`
}

type problemListItem struct {
	Id       string `json:"problem_id,omitempty" bson:"_id,omitempty"`
	No       string `json:"problem_no,omitempty" bson:"problem_no"`
	Title    string `json:"title,omitempty" bson:"title,omitempty"`
	AuthorId string `json:"author_id,omitempty" bson:"author_id,omitempty"`
}

type problemDetail struct {
	problemListItem
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	InputDesc   string     `json:"input_desc,omitempty" bson:"input_desc,omitempty"`
	OutputDesc  string     `json:"output_desc,omitempty" bson:"output_desc,omitempty"`
	SampleCases []testCase `json:"sample_cases,omitempty" bson:"sample_cases,omitempty"`
}

type TestCases struct {
	Cases []testCase `json:"test_cases,omitempty" bson:"test_cases,omitempty"`
}

type problem struct {
	Id          string     `json:"problem_id,omitempty" bson:"_id,omitempty"`
	No          string     `json:"problem_no,omitempty" bson:"problem_no,omitempty"`
	Title       string     `json:"title,omitempty" bson:"title,omitempty"`
	AuthorId    string     `json:"author_id,omitempty" bson:"author_id,omitempty"`
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	InputDesc   string     `json:"input_desc,omitempty" bson:"input_desc,omitempty"`
	OutputDesc  string     `json:"output_desc,omitempty" bson:"output_desc,omitempty"`
	SampleCases []testCase `json:"sample_cases,omitempty" bson:"sample_cases,omitempty"`
	Cases       []testCase `json:"test_cases,omitempty" bson:"test_cases,omitempty"`
}

type QueryProblemRequest struct {
	ProblemId string `json:"problem_id" bson:"_id"`
}

type deleteProblemRequest QueryProblemRequest

func addProblem(ctx context.Context, p *problem) error {
	log.Print(p.No)

	_, err := problems.InsertOne(ctx, p)

	if e, ok := err.(mongo.WriteException); ok && e.HasErrorCode(11000) {
		return err
	}

	return err
}

func deleteProblem(ctx context.Context, req deleteProblemRequest) error {
	id, err := primitive.ObjectIDFromHex(req.ProblemId)
	if err != nil {
		return err
	}

	r, err := problems.DeleteOne(ctx, bson.D{{"_id", id}})

	if r.DeletedCount == 0 {
		return errors.New("problem does not exist")
	}

	return err
}

func modifyProblem(ctx context.Context, p *problem) error {
	id, err := primitive.ObjectIDFromHex(p.Id)

	p.Id = ""
	data, err := problems.UpdateMany(
		ctx,
		bson.D{{"_id", id}},
		bson.D{{"$set", p}})

	if err != nil {
		return err
	}

	if data.ModifiedCount == 0 {
		return util.ErrNothingModified
	}

	return nil
}

func doListProblem(ctx context.Context, data *mongo.Cursor) (ls []problemListItem, err error) {
	for data.Next(ctx) {
		item := problemListItem{}
		err = data.Decode(&item)
		if err != nil {
			return nil, err
		}
		ls = append(ls, item)
	}

	return
}

type listProblemRequest struct {
	Index int `json:"index"`
}

func listProblem(ctx context.Context, r listProblemRequest) ([]problemListItem, error) {
	pageLimit := 50

	data, err := problems.Aggregate(ctx, bson.A{
		bson.D{{"$sort", bson.D{
			{"problem_no", 1},
		}}},
		bson.D{{"$skip", (r.Index - 1) * pageLimit}},
		bson.D{{"$limit", pageLimit}},
		bson.D{{"$project", bson.D{
			{"_id", 1},
			{"problem_no", 1},
			{"title", 1},
			{"author_id", 1},
		}}},
	})

	if err != nil {
		return nil, err
	}

	return doListProblem(ctx, data)
}

type searchProblemRequest struct {
	Title string `json:"title"`
}

func searchProblem(ctx context.Context, r searchProblemRequest) ([]problemListItem, error) {
	data, err := problems.Aggregate(ctx, bson.A{
		bson.D{
			{"$match", bson.D{
				{"$text", bson.D{
					{"$search", r.Title},
				}},
			}}},
		bson.D{
			{"$project", bson.D{
				{"_id", 1},
				{"problem_no", 1},
				{"title", 1},
				{"author_id", 1},
			}}},
		bson.D{{"$sort", bson.D{
			{"problem_no", 1},
		}}},
	})

	if err != nil {
		return nil, err
	}

	return doListProblem(ctx, data)
}

func queryProblemDetail(ctx context.Context, r QueryProblemRequest) (*problemDetail, error) {
	id, err := primitive.ObjectIDFromHex(r.ProblemId)
	if err != nil {
		return nil, err
	}

	detail := problemDetail{}
	err = problems.FindOne(ctx, bson.D{{"_id", id}},
		&options.FindOneOptions{
			Projection: bson.D{
				{"_id", 1},
				{"problem_no", 1},
				{"title", 1},
				{"author_id", 1},
				{"description", 1},
				{"input_desc", 1},
				{"output_desc", 1},
				{"sample_cases", 1},
			}}).Decode(&detail)

	if err != nil {
		return nil, err
	}

	return &detail, nil
}

func queryProblem(ctx context.Context, r QueryProblemRequest) (*problem, error) {
	id, err := primitive.ObjectIDFromHex(r.ProblemId)
	if err != nil {
		return nil, err
	}

	problem := problem{}
	err = problems.FindOne(ctx, bson.D{{"_id", id}}).Decode(&problem)
	if err != nil {
		return nil, err
	}

	return &problem, nil
}

func getProblemOwner(ctx context.Context, r QueryProblemRequest) (string, error) {
	id, err := primitive.ObjectIDFromHex(r.ProblemId)
	if err != nil {
		return "", err
	}

	var query bson.M
	err = problems.FindOne(ctx,
		bson.D{{"_id", id}}).Decode(&query)

	log.Print(query["author_id"])

	if err != nil {
		return "", err
	}

	author, ok := query["author_id"].(string)
	if !ok {
		return "", errors.New("author_id does not exist")
	}

	return author, nil
}

func QueryTestCase(ctx context.Context, id primitive.ObjectID) (*TestCases, error) {

	tcs := TestCases{}
	err := problems.FindOne(ctx,
		bson.D{{"_id", id}},
		&options.FindOneOptions{
			Projection: bson.D{
				{"_id", 1},
				{"test_cases", 1},
			}}).Decode(&tcs)

	if err != nil {
		return nil, err
	}

	return &tcs, nil
}
