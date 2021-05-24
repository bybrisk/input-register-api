package data

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var resultID string

func AddUserToDatabase(d *RegisterUserStructure) (string,error) {
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
	return resultID,insertErr
}

func GetBusinessName(docID string) (string,string) {

	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	
	type BusinessInfo struct {
		BusinessName string `json:"businessname"`
		BusinessCategory string `json:"businesscategory"`
	}

	var document BusinessInfo

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetBusinessName ERROR:")
		log.Error(err)
		return document.BusinessName,document.BusinessCategory
	}
	return document.BusinessName,document.BusinessCategory
}

func RegisterToBusinessMongo(d *RegisterUserToBusinessStruct,nameOfBusiness string, businessCategory string) int64 {

	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(d.UserID)
	filter := bson.M{"_id": id}

	type SubscriptionInfo struct {
		BusinessID string `json:"businessID"`
		BusinessName string `json:"businessname"`
		BusinessCategory string `json:"businesscategory"`
	}

	document := SubscriptionInfo{
		BusinessID:d.BusinessID,
		BusinessName:nameOfBusiness,
		BusinessCategory:businessCategory,
	}

	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.D{{Key: "$push", Value: bson.M{"subscription": document}}})
	if err != nil {
		log.Error("RegisterToBusinessMongo ERROR:")
		log.Error(err)
		return 0
	}

	return updateResult.ModifiedCount
}

func IsSubscribedAlready(d *RegisterUserToBusinessStruct) (int64,error) {
	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(d.UserID)
	filter := bson.M{"_id": id}

	type BusinessIDArray struct {
		Subscription []struct {
			BusinessID string `json:"businessID"`
		} `json:"subscription"`
	}

	var document BusinessIDArray

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("IsSubscribedAlready ERROR:")
		log.Error(err)
	}

	var isSubscribed int64
	for _,val:=range document.Subscription{
		if val.BusinessID==d.BusinessID {
			isSubscribed = 1
		}
	}
	
	return isSubscribed,err
}

func GetUserIDByPhoneMongo(phone string) (IdOfDoc, error) {
	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	filter := bson.M{"phonenumber": phone}

	var document IdOfDoc

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetUserIDByPhoneMongo ERROR:")
		log.Error(err)
	}

	return document,err
}

func UpdateUserToDatabase(d *RegisterUserStructure) (int64,error) {
	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(d.UserID)
	update := bson.M{"$set": bson.M{"userid": d.UserID, "username":d.UserName, "address":d.Address, "latitude":d.Latitude, "longitude":d.Longitude}}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("UpdateUserToDatabase ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount,err
}