package mongo

import (
	"gopkg.in/mgo.v2"
)

func getDB(databaseName string) (*mgo.Database, error) {
	session, err := mgo.Dial("ip:port")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB(databaseName), err
}

func closeDB(db *mgo.Database) {
	db.Session.Close()
}