package mongo

import "errors"

func Insert(databaseName, collectionName string, data interface{}) (err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		return
	}
	defer closeDB(db)

	err = db.C(collectionName).Insert(&data)
	if err != nil {
		return
	}

	return
}

func Remove(databaseName, collectionName string, filter map[string]interface{}) (err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		return
	}
	defer closeDB(db)

	err = db.C(collectionName).Remove(filter)
	if err != nil {
		return
	}

	return
}

func Update(databaseName, collectionName string, filter map[string]interface{}, data interface{}) (err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		return
	}
	defer closeDB(db)

	err = db.C(collectionName).Update(filter, &data)
	if err != nil {
		return
	}

	return
}

//计算某个集合的数量
func FindCount (databaseName, collectionName string, filter map[string]interface{}) (total int, err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		return
	}
	defer closeDB(db)

	total,err = db.C(collectionName).Find(filter).Count()
	if err != nil {
		return
	}

	return
}

func FindGiven (databaseName, collectionName string, filter map[string]interface{}, offset, limit int, sort string) (data []interface{}, err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		err = errors.New("database connection fail")
		return
	}
	defer closeDB(db)

	if sort != ""{
		err = db.C(collectionName).Find(filter).Sort(sort).Skip(offset).Limit(limit).All(&data)
	} else {
		err = db.C(collectionName).Find(filter).Skip(offset).Limit(limit).All(&data)
	}

	if err != nil {
		return
	}
	return
}

func FindAll (databaseName, collectionName string, filter map[string]interface{}) (data []interface{}, err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		err = errors.New("database connection fail")
		return
	}
	defer closeDB(db)

	err = db.C(collectionName).Find(filter).All(&data)
	if err != nil {
		return
	}
	return
}

func FindOne (databaseName, collectionName string, filter map[string]interface{}) (data interface{}, err error) {
	db, err :=  getDB(databaseName)
	if err != nil {
		err = errors.New("database connection fail")
		return
	}
	defer closeDB(db)

	err = db.C(collectionName).Find(filter).One(&data)
	if err != nil {
		return
	}
	return
}

func AggregateOne(databaseName, collectionName string, selector interface{}) (data interface{}, err error){
	db, err :=  getDB(databaseName)
	if err != nil {
		err = errors.New("database connection fail")
		return
	}
	defer closeDB(db)

	err = db.C(collectionName).Pipe(selector).AllowDiskUse().One(&data)
	if err != nil {
		return
	}
	return
}