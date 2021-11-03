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
	mongo.UpdateResult
}

type SingleResult struct {
	err
	mongo.SingleResult
}

type InsertOneResult struct {
	err
	mongo.InsertOneResult
}

type InsertManyResult struct {
	err
	mongo.InsertManyResult
}

type DeleteResult struct {
	err
	mongo.DeleteResult
}
