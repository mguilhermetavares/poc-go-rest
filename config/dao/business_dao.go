package dao

import (
	"log"

	. "github.com/mguilhermetavares/poc-go-rest/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BusinessDAO  -  estrutura de comunicação com o banco
type BusinessDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Collection
const (
	COLLECTION = "business"
)

// Connect - funcao de  conecao com o banco
func (m *BusinessDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// GetAll - busca todos os business do banco
func (m *BusinessDAO) GetAll() ([]Business, error) {
	var business []Business
	err := db.C(COLLECTION).Find(bson.M{}).All(&business)
	return business, err
}

// GetByID -busca pelo ID
func (m *BusinessDAO) GetByID(id string) (Business, error) {
	var business Business
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&business)
	return business, err
}

// Create - insere um objeto no banco
func (m *BusinessDAO) Create(business Business) error {
	err := db.C(COLLECTION).Insert(&business)
	return err
}

// Delete - remove um objeto no banco
func (m *BusinessDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update - atualiza um objeto no banco
func (m *BusinessDAO) Update(id string, business Business) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &business)
	return err
}
