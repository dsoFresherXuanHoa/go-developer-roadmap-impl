package services

import (
	"gin-restful-mongo/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllContact(ctx *gin.Context, client *mongo.Client) (models.Contacts, error) {
	collection := client.Database("gin_restful_mongo_dev").Collection("contacts")
	var filter = bson.D{{}}
	if cur, err := collection.Find(ctx, filter); err != nil {
		log.Println("Something went wrong: " + err.Error())
		return nil, err
	} else {
		var contacts models.Contacts
		if err := cur.All(ctx, &contacts); err != nil {
			log.Println("Can't parse: " + err.Error())
			return nil, err
		} else {
			return contacts, nil
		}
	}
}

func FindContactByName(ctx *gin.Context, client *mongo.Client, name string) (models.Contact, error) {
	collection := client.Database("gin_restful_mongo_dev").Collection("contacts")
	var filter = bson.D{{"name", name}}
	var contact models.Contact
	if err := collection.FindOne(ctx, filter).Decode(&contact); err != nil {
		log.Println("Something went wrong: " + err.Error())
		return models.Contact{}, err
	} else {
		return contact, nil
	}
}

func SaveContact(ctx *gin.Context, client *mongo.Client, contact models.Contact) (interface{}, error) {
	collection := client.Database("gin_restful_mongo_dev").Collection("contacts")
	if result, err := collection.InsertOne(ctx, contact); err != nil {
		log.Println("Something went wrong: " + err.Error())
		return nil, err
	} else {
		return result.InsertedID, nil
	}
}

func UpdateContact(ctx *gin.Context, client *mongo.Client, contact models.Contact, name string) (interface{}, error) {
	collection := client.Database("gin_restful_mongo_dev").Collection("contacts")
	filter := bson.D{{"name", name}}
	if result, err := collection.ReplaceOne(ctx, filter, contact); err != nil {
		log.Println("Something went wrong: " + err.Error())
		return nil, err
	} else {
		return result.ModifiedCount, nil
	}
}

func DeleteContact(ctx *gin.Context, client *mongo.Client, name string) (interface{}, error) {
	collection := client.Database("gin_restful_mongo_dev").Collection("contacts")
	var filter = bson.D{{"name", name}}
	if result, err := collection.DeleteOne(ctx, filter); err != nil {
		log.Println("Something went wrong: " + err.Error())
		return nil, err
	} else {
		return result.DeletedCount, nil
	}
}
