/*package consisting of all the monog db operations*/
package mongo_package

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	ProductId     string        `bson:"productId" json:"productId"`
	ProductName   string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	Current_price float64       `bson:"current_price,omitempty" json:"current_price,omitempty"`
	Currency      string        `bson:"currency,omitempty" json:"currency,omitempty"`
}

const (
	host       = "localhost"
	port       = 27017
	db_name    = "test"
	collection = "products"
)

var Coll_mongo *mgo.Collection

func init() {
	//var product Product
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	//defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	Coll_mongo = session.DB(db_name).C(collection)
}

/*creating data*/
func Createproduct(prod_details Product) error {
	// Insert Data
	err := Coll_mongo.Insert(&prod_details)

	if err != nil {
		panic(err)
		return err
	}
	return nil
}
func FindProduct(productid int) Product {
	result := Product{}
	err := Coll_mongo.Find(bson.M{"productId": productid}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result.Description)
	return result
}
func Updateproduct(prod_details Product) error {
	// Update
	colQuerier := bson.M{"productId": prod_details.ProductId}
	change := bson.M{"$set": bson.M{"productId": prod_details.ProductId, "title": prod_details.ProductName, "description": prod_details.Description}}
	err := Coll_mongo.Update(colQuerier, change)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
func Deleteproduct(product string) error {
	fmt.Println("productid: ", product)
	err := Coll_mongo.Remove(bson.D{{"productId", product}})
	if err != nil {
		return err
	}
	return nil
}
