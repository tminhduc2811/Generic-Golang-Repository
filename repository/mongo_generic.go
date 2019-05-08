package repository

import (
	"gopkg.in/mgo.v2"
)

type GenericRepository struct {
	session      *mgo.Session
	dbName       string
	dbCollection string
}

func NewGenericRepository(session *mgo.Session, dbName string, dbCollection string) *GenericRepository {
	return &GenericRepository{
		session:      session,
		dbName:       dbName,
		dbCollection: dbCollection,
	}
}

func (r *GenericRepository) Fetch(query interface{}, model interface{}) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Find(query).All(model)
	return err
}

func (r *GenericRepository) Create(model interface{}) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Insert(model)
	return err
}

func (r *GenericRepository) Update(selector interface{}, model interface{}) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Update(selector, model)
	return err
}

func (r *GenericRepository) Delete(query interface{}) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Remove(query)
	return err
}

func (r *GenericRepository) FetchOne(query interface{}, model interface{}) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Find(query).One(model)
	return err
}

func (r *GenericRepository) Exists(query interface{}) (bool, error) {
	session := r.getSession()
	defer session.Close()
	count, err := session.DB(r.dbName).C(r.dbCollection).Find(query).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, err
	}
	return false, err
}

func (r *GenericRepository) getSession() *mgo.Session {
	return r.session.Copy()
}
