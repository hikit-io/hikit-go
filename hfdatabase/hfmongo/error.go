package hfmongo

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Error interface {
	IsNoDocuments() bool
	ExceptNoDocuments() error
}

type err struct {
	error
}

func (e err) IsNoDocuments() bool {
	if IsErrNoDocuments(e) {
		return true
	}
	return false
}

func (e err) ExceptNoDocuments() error {
	if e.IsNoDocuments() {
		return nil
	}
	return e
}

func (e err) NotNil() bool {
	return e.error != nil
}

func (e err) Err() error {
	return e.error
}

func IsErrNoDocuments(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}

func ExceptNoDocuments(err error) error {
	if IsErrNoDocuments(err) {
		return nil
	}
	return err
}

type UpdateResult struct {
	err
	*mongo.UpdateResult
}

type SingleResult struct {
	err
}

type InsertOneResult struct {
	err
	*mongo.InsertOneResult
}

type InsertManyResult struct {
	err
	*mongo.InsertManyResult
}

type DeleteResult struct {
	err
	*mongo.DeleteResult
}

type FindResult struct {
	err
}

type CountResult struct {
	err
	Count int64
}
