package Services

import (
	"time"

	"../Api"
	"../Utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IMongoService interface {
	MongoSession() *mgo.Session
	AllRestaurants(s *mgo.Session) ([]Api.Restaurant, error)
	FindRestaurant(s *mgo.Session, id int) (Api.Restaurant, error)
	InsertRestaurant(s *mgo.Session, restaurant Api.Restaurant) (bool, error)
	UpdateRestaurant(s *mgo.Session, restaurant Api.Restaurant) (bool, error)
	DeleteRestaurant(s *mgo.Session, id int) (bool, error)
	inject()
}

type MongoService struct {
	settings   Utils.Setting
	logger     Utils.ILogger
	mgoSession *mgo.Session
}

func (m *MongoService) MongoSession() *mgo.Session {
	m.inject()

	if m.mgoSession == nil {
		info := &mgo.DialInfo{
			Addrs:    []string{m.settings.Mongo.Hosts},
			Timeout:  60 * time.Second,
			Database: m.settings.Mongo.Db,
			Username: m.settings.Mongo.User,
			Password: m.settings.Mongo.Pass,
		}

		var err error
		m.mgoSession, err = mgo.DialWithInfo(info)

		if err != nil {
			m.logger.WriteLog(err.Error(), Utils.Fatal)
		}

		m.mgoSession.SetMode(mgo.Monotonic, true)
	}
	return m.mgoSession.Clone()
}

func (m *MongoService) AllRestaurants(s *mgo.Session) ([]Api.Restaurant, error) {
	m.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(m.settings.Mongo.Db).C("restaurants")

	var restaurants []Api.Restaurant
	err := c.Find(bson.M{}).All(&restaurants)

	if err != nil {
		m.logger.WriteLog(err.Error(), Utils.Info)
	}

	return restaurants, nil
}

func (m *MongoService) FindRestaurant(s *mgo.Session, id int) (Api.Restaurant, error) {
	m.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(m.settings.Mongo.Db).C("restaurants")

	var restaurant Api.Restaurant
	err := c.Find(bson.M{"id": id}).One(&restaurant)
	if err != nil {
		m.logger.WriteLog(err.Error(), Utils.Info)
		return restaurant, err
	}

	return restaurant, nil
}

func (m *MongoService) InsertRestaurant(s *mgo.Session, restaurant Api.Restaurant) (bool, error) {
	m.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(m.settings.Mongo.Db).C("restaurants")

	_, errf := m.FindRestaurant(s, restaurant.ID)

	if errf == nil {
		return false, nil
	}

	err := c.Insert(restaurant)
	if err != nil {
		m.logger.WriteLog(err.Error(), Utils.Info)
		return false, err
	}

	return true, nil
}

func (m *MongoService) UpdateRestaurant(s *mgo.Session, restaurant Api.Restaurant) (bool, error) {
	m.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(m.settings.Mongo.Db).C("restaurants")

	err := c.Update(bson.M{"id": restaurant.ID}, &restaurant)
	if err != nil {
		m.logger.WriteLog(err.Error(), Utils.Info)
		return false, err
	}

	return true, nil
}

func (m *MongoService) DeleteRestaurant(s *mgo.Session, id int) (bool, error) {
	m.inject()

	session := s.Copy()
	defer session.Close()

	c := session.DB(m.settings.Mongo.Db).C("restaurants")

	err := c.Remove(bson.M{"id": id})
	if err != nil {
		m.logger.WriteLog(err.Error(), Utils.Info)
		return false, err
	}

	return true, nil
}

//Injections
func (m *MongoService) inject() {
	var injLogger Utils.Logger

	m.logger = &injLogger
	m.settings = Utils.GetSettings()
}
