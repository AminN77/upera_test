package internal

import (
	"context"
	"errors"
	mongoPkg "github.com/AminN77/upera_test/history_service/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	ErrDatabase = errors.New("some error occurred on the database side")
)

type Repository interface {
	Insert(r *Revision) error

	// InsertBatch is more performant with bulk data
	InsertBatch(r []*Revision) error

	GetRevisionsOfOneProduct(pageSize, pageIndex, productID int64, ctx context.Context) ([]*Revision, error)

	GetRevisionByRevisionNumber(revisionNumber string, ctx context.Context) (*Revision, error)
}

type mongoRepository struct {
	cli        *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository() Repository {
	opts := options.Client()

	opts.ApplyURI(os.Getenv("MONGO_URL"))
	cli := mongoPkg.NewMongoClient(context.Background(), opts)

	collection := cli.Database(os.Getenv("MONGO_DB_NAME")).
		Collection(os.Getenv("MONGO_COLLECTION_NAME"))

	unique := true
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"revisionNumber": 1},
		Options: &options.IndexOptions{Unique: &unique},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	return &mongoRepository{
		cli:        cli,
		collection: collection,
	}
}

func (mr *mongoRepository) Insert(r *Revision) error {
	_, err := mr.collection.InsertOne(context.Background(), r)
	if err != nil {
		log.Println(err)
		return ErrDatabase
	}

	return nil
}

func (mr *mongoRepository) InsertBatch(r []*Revision) error {
	docs := make([]interface{}, len(r))
	for i := 0; i < len(r); i++ {
		docs[i] = r[i]
	}

	opts := options.InsertMany()
	opts.SetOrdered(false)
	_, err := mr.collection.InsertMany(context.Background(), docs, opts)
	if err != nil {
		log.Println(err)
		return ErrDatabase
	}

	return nil
}

func (mr *mongoRepository) GetRevisionsOfOneProduct(pageSize, pageIndex, productID int64, ctx context.Context) ([]*Revision, error) {
	var res []*Revision

	filter := bson.M{"productID": productID}
	opts := options.Find()
	opts.SetSort(bson.M{"createdAt": -1})
	skip := pageSize * (pageIndex - 1)
	opts.Skip = &skip
	opts.Limit = &pageSize

	cursor, err := mr.collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println(err)
		return nil, ErrDatabase
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("could not close cursor, err", err.Error())
		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var temp *Revision
		err := cursor.Decode(&temp)
		if err != nil {
			log.Println(err)
			return nil, ErrDatabase
		}

		res = append(res, temp)
	}

	// Check for errors from cursor.Err()
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, ErrDatabase
	}

	return res, nil
}

func (mr *mongoRepository) GetRevisionByRevisionNumber(revisionNumber string, ctx context.Context) (*Revision, error) {
	var result Revision

	filter := bson.M{"revisionNumber": revisionNumber}

	if err := mr.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("document with revisionNumber %s not found.\n", revisionNumber)
			return nil, ErrDatabase
		} else {
			log.Fatal(err)
			return nil, ErrDatabase
		}
	}

	return &result, nil
}
