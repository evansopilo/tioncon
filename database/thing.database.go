package database

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/evansopilo/tioncon/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrInsert = errors.New("insert thing failed")
	ErrRead   = errors.New("read thing failed")
	ErrUpdate = errors.New("update thing failed")
	ErrDelete = errors.New("delete thing failed")
)

type FetchOptions struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	DeviceID string `json:"device_id,omitempty" bson:"device_id,omitempty"`
}

type IThing interface {
	Insert(thing models.IThing) error
	Read(id string) ([]models.IThing, error)
	Fetch(opts FetchOptions, page, limit int64) ([]models.Thing, error)
	Remove(id string) error
}

type Thing struct {
	Col    *mongo.Collection
	Logger *log.Entry
}

func NewThing(col *mongo.Collection, logger *log.Entry) *Thing {
	return &Thing{Col: col, Logger: logger}
}

func (t Thing) Insert(thing models.IThing) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := t.Col.InsertOne(ctx, thing); err != nil {
		t.Logger.Error(err)
		return ErrInsert
	}
	t.Logger.Infof("insert thing with id: %v", thing.GetThing().ID)
	return nil
}

func (t Thing) Read(id string) ([]models.IThing, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filterCursor, err := t.Col.Find(ctx, bson.M{"_id": id})

	if err != nil {
		t.Logger.Error(err)
		return nil, err
	}

	var things []models.IThing

	for filterCursor.Next(ctx) {
		var thing models.IThing = models.NewThing()
		if err := filterCursor.Decode(thing); err != nil {
			t.Logger.Error(err)
			return nil, err
		}
		things = append(things, thing)
	}
	t.Logger.Infof("read thing with id: %v", id)
	return things, nil
}

func (t Thing) Fetch(opts FetchOptions, page, limit int64) ([]models.Thing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := new(options.FindOptions)

	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	filterCursor, err := t.Col.Find(ctx, opts, options)
	if err != nil {
		t.Logger.Error(err)
		return nil, ErrRead
	}

	var things []models.Thing

	for filterCursor.Next(ctx) {
		var thing models.Thing
		if err := filterCursor.Decode(&thing); err != nil {
			t.Logger.Error(err)
			return nil, ErrRead
		}
		things = append(things, thing)
	}
	t.Logger.Info("read thing at page: %v limit: %v", page, limit)
	return things, nil
}

func (t Thing) Remove(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := t.Col.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		t.Logger.Error(err)
		return ErrDelete
	}
	t.Logger.Infof("delete thing with id: %v", id)
	return nil
}
