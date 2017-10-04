package models

import (
	"time"

	"github.com/omongaco/visitindonesia/models/db"
	"gopkg.in/mgo.v2/bson"
)

//Province object
type Province struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Intro      string        `json:"intro" bson:"intro"`
	Content    string        `json:"content" bson:"content"`
	CoverImage string        `json:"cover_image" bson:"cover_image"`
	CreatedOn  time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time     `json:"updated_on" bson:"updated_on"`
}

func newProvinceCollection() *db.Collection {
	return db.NewCollectionSession("province")
}

//CreateProvince API to create Destination
func CreateProvince(province Province) (Province, error) {
	var (
		err error
	)

	c := newProvinceCollection()
	defer c.Close()

	province.ID = bson.NewObjectId()
	province.CreatedOn = time.Now()

	err = c.Session.Insert(&province)
	if err != nil {
		return province, err
	}

	return provice, err
}

//UpdateProvince API to update Provice
func (province Province) UpdateProvince(provinceParam Province) (Province, error) {
	var (
		err error
	)

	//Get province collection
	c := newProvinceCollection()
	defer c.Close()

	//update province
	err = c.Session.Update(bson.M{
		"_id": province.ID,
	}, bson.M{
		"$set": bson.M{
			"name":        provinceParam.Name,
			"intro":       provinceParam.Intro,
			"content":     provinceParam.Content,
			"cover_image": provinceParam.CoverImage,
			"created_on":  provinceParam.CreatedOn,
			"updated_on":  time.Now(),
		},
	})

	if err != nil {
		return province, err
	}

	return province, err
}

//GetDestinations API to get all destinations
func GetProvinces() ([]Province, error) {
	var (
		err       error
		provinces []Province
	)

	c := newProvinceCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-updated_on").All(&provinces)
	if err != nil {
		return provinces, err
	}

	return provinces, err
}

//GetDestination by ID API return one document
func GetProvince(id bson.ObjectId) (Province, error) {
	var (
		err      error
		province Province
	)

	c := newProvinceCollection()
	defer c.Close()

	err = c.Session.FindId(id).One(&province)
	if err != nil {
		return province, err
	}

	return province, err
}

func DeleteProvince(province Province) (Province, error) {
	var (
		err error
	)

	c := newProvinceCollection()
	defer c.Close()

	err = c.Session.Remove(bson.M{"_id": province.ID})
	if err != nil {
		return err
	}

	return err
}
