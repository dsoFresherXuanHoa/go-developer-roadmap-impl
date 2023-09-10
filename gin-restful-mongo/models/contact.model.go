package models

type Contact struct {
	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
	Address string `bson:"address"`
}

type Contacts []Contact
