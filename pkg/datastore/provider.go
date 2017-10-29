package datastore

import (
	"messagebird/pkg/model"

	"gopkg.in/mgo.v2/bson"
)

// Provider defines the interface for a data provider
type Provider interface {
	Create(sms *model.SMS) error
	GetByRefID(refID string) []model.SMS
	UpdateStatus(refID string, status string) error
}

type mongo struct {
	sessionFactory *SessionFactory
}

// NewMongo return a pointer to the mongo implementation of Provider
func NewMongo(s *SessionFactory) Provider {
	return &mongo{sessionFactory: s}
}

// Create - create an SMS document
func (m *mongo) Create(sms *model.SMS) error {
	session := m.sessionFactory.Get()
	defer session.Close()

	return session.DB("messagebird").C("sms").Insert(sms)
}

// GetByRefID - get sms data by reference ID
func (m *mongo) GetByRefID(refID string) []model.SMS {
	session := m.sessionFactory.Get()
	defer session.Close()

	var s []model.SMS
	session.DB("messagebird").C("sms").Find(bson.M{"reference": refID}).All(&s)

	return s
}

// UpdateStatus updates the status of the SMS's with the given ref ID
func (m *mongo) UpdateStatus(refID string, status string) error {
	session := m.sessionFactory.Get()
	defer session.Close()

	query := bson.M{"reference": refID}
	update := bson.M{
		"$set": bson.M{"status": status},
	}

	_, err := session.DB("messagebird").C("sms").UpdateAll(query, update)

	return err
}
