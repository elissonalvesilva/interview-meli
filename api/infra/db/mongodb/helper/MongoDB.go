package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoDB interface {
	Collection(name string) CollectionHelper
	Client() ClientHelper
}

type CollectionHelper interface {
	FindOne(context.Context, interface{}) SingleResultHelper
	InsertOne(context.Context, interface{}) (interface{}, error)
	Find(context.Context, interface{}) (*mongo.Cursor,error)
	Aggregate(ctx context.Context, pipeline interface{}) (*mongo.Cursor, error)
	CountDocuments(ctx context.Context, filter interface{}) (int64, error)
}

type SingleResultHelper interface {
	Decode(v interface{}) error
}

type ClientHelper interface {
	Database(string) MongoDB
	Connect() error
	StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

func NewClient() (ClientHelper, error) {
	c, err := mongo.NewClient(options.Client().SetAuth(
		options.Credential{
			Username:   os.Getenv("DB_USERNAME"),
			Password:   os.Getenv("DB_PASSWORD"),
		}).ApplyURI(os.Getenv("DB_URI")))

	return &mongoClient{cl: c}, err

}

func NewDatabase(client ClientHelper) MongoDB {
	return client.Database(os.Getenv("DB_NAME"))
}

func (mc *mongoClient) Database(dbName string) MongoDB {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect() error {
	return mc.cl.Connect(nil)
}

func (md *mongoDatabase) Collection(colName string) CollectionHelper {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	if id != nil {
		return id.InsertedID, err
	}
	return id, err
}

func (mc *mongoCollection) Find(ctx context.Context, filters interface{}) (*mongo.Cursor, error) {
	c, err := mc.coll.Find(ctx, filters)
	return c, err
}

func (mc *mongoCollection) 	Aggregate(ctx context.Context, pipeline interface{}) (*mongo.Cursor, error) {
	return mc.coll.Aggregate(ctx, pipeline)
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	return mc.coll.CountDocuments(ctx, filter)
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
