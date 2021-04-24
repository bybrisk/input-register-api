package data

import (
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var resultID string

func AddUserToDatabase(d *RegisterUserStructure) string{
	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	result, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, d)
	if insertErr != nil {
		log.Error("AddUserToDatabase() ERROR:")
		log.Error(insertErr)
	} else {
		fmt.Println("AddUserToDatabase() API result:", result)

		newID := result.InsertedID
		fmt.Println("AddUserToDatabase() newID:", newID)
		resultID = newID.(primitive.ObjectID).Hex()
	}
	return resultID
}