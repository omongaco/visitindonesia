package models

import (
	"time"

	"github.com/omongaco/visitindonesia/models/db"
	"gopkg.in/mgo.v2/bson"
)

//Province object
type Category struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Intro      string        `json:"intro" bson:"intro"`
	Content    string        `json:"content" bson:"content"`
	CoverImage string        `json:"cover_image" bson:"cover_image"`
	Type       Type          `json:"type" bson:"type"`
	CreatedOn  time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time     `json:"updated_on" bson:"updated_on"`
}

type Type struct {
	Name       string    `json:"name" bson:"name"`
	Content    string    `json:"content" bson:"content"`
	CoverImage string    `json:"cover_image" bson:"cover_image"`
	CreatedOn  time.Time `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time `json:"updated_on" bson:"updated_on"`
}

func newCategoryCollection() *db.Collection {
	return db.NewCollectionSession("category")
}

//CreateProvince API to create Destination
func CreateCategory(category Category) (Province, error) {
	var (
		err error
	)

	c := newCategoryCollection()
	defer c.Close()

	category.ID = bson.NewObjectId()
	categoty.CreatedOn = time.Now()

	err = c.Session.Insert(&category)
	if err != nil {
		return category, err
	}

	return category, err
}

//UpdateProvince API to update Provice
func (category Category) UpdateCategory(categoryParam Category) (Category, error) {
	var (
		err error
	)

	//Get province collection
	c := newCategoryCollection()
	defer c.Close()

	//update province
	err = c.Session.Update(bson.M{
		"_id": category.ID,
	}, bson.M{
		"$set": bson.M{
			"name":        categoryParam.Name,
			"intro":       categoryParam.Intro,
			"content":     categoryParam.Content,
			"cover_image": categoryParam.CoverImage,
			"type":        categoryParam.Type,
			"created_on":  categoryParam.CreatedOn,
			"updated_on":  time.Now(),
		},
	})

	if err != nil {
		return category, err
	}

	return category, err
}

//GetDestinations API to get all destinations
func GetCategories() ([]Category, error) {
	var (
		err        error
		categories []Category
	)

	c := newCategoryCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-updated_on").All(&categories)
	if err != nil {
		return categories, err
	}

	return categories, err
}

//GetDestination by ID API return one document
func GetCategory(id bson.ObjectId) (Category, error) {
	var (
		err      error
		category Category
	)

	c := newCategoryCollection()
	defer c.Close()

	err = c.Session.FindId(id).One(&category)
	if err != nil {
		return category, err
	}

	return category, err
}

func DeleteCategory(category Category) (Category, error) {
	var (
		err error
	)

	c := newCategoryCollection()
	defer c.Close()

	err = c.Session.Remove(bson.M{"_id": category.ID})
	if err != nil {
		return err
	}

	return err
}
