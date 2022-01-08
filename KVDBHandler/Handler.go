package KVDBHandler

import "os"

const path = "./Database/"

const dbName = "MyDB.kvdb"

func init() {
	os.MkdirAll(path, os.ModePerm)
	if _, err := os.Stat(path + dbName); os.IsNotExist(err) {
		os.Create(path + dbName)
	}
}

func create() {

}

func set() {

}

func get() {

}

func remove() {

}

func purge() {

}
