package models

import (
	"time"

	"github.com/omongaco/visitindonesia/models/db"
	"gopkg.in/mgo.v2/bson"
)

//Destination object
type Destination struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Intro      string        `json:"intro" bson:"intro"`
	Content    string        `json:"content" bson:"content"`
	CoverImage string        `json:"cover_image" bson:"cover_image"`
	CreatedOn  time.Time     `json:"created_on" bson:"created_on"`
	UpdatedOn  time.Time     `json:"updated_on" bson:"updated_on"`
}

func newDestinationCollection() *db.Collection {
	return db.NewCollectionSession("destination")
}

//CreateDestination API to create Destination
func CreateDestination(destination Destination) (Destination, error) {
	var (
		err error
	)

	c := newDestinationCollection()
	defer c.Close()

	destination.ID = bson.NewObjectId()
	destination.CreatedOn = time.Now()

	err = c.Session.Insert(&destination)
	if err != nil {
		return destination, err
	}

	return destination, err
}

//UpdateDestination API to update Destination
func (destination Destination) UpdateDestination(destinationParam Destination) (Destination, error) {
	var (
		err error
	)

	//Get destination collection
	c := newDestinationCollection()
	defer c.Close()

	//update destination
	err = c.Session.Update(bson.M{
		"_id": destination.ID,
	}, bson.M{
		"$set": bson.M{
			"name":        destinationParam.Name,
			"intro":       destinationParam.Intro,
			"content":     destinationParam.Content,
			"cover_image": destinationParam.CoverImage,
			"created_on":  destinationParam.CreatedOn,
			"updated_on":  time.Now(),
		},
	})

	if err != nil {
		return destination, err
	}

	return destination, err
}

//GetDestinations API to get all destinations
func GetDestinations() ([]Destination, error) {
	var (
		err          error
		destinations []Destination
	)

	c := newDestinationCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-updated_on").All(&destinations)
	if err != nil {
		return destinations, err
	}

	return destinations, err
}

//GetDestination by ID API return one document
func GetDestination(id bson.ObjectId) (Destination, error) {
	var (
		err         error
		destination Destination
	)

	c := newDestinationCollection()
	defer c.Close()

	err = c.Session.FindId(id).One(&destination)
	if err != nil {
		return destination, err
	}

	return destination, err
}

func DeleteDestination(destination Destination) (Destination, error) {
	var (
		err error
	)

	c := newDestinationCollection()
	defer c.Close()

	err = c.Session.Remove(bson.M{"_id": destination.ID})
	if err != nil {
		return err
	}

	return err
}
