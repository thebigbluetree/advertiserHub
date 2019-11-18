package dbconnection

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//creating a mongo session
func GetSession() *mgo.Session{
	s, err:= mgo.Dial("mongodb://localhost")
	if err!= nil{
		log.Fatalln("%s", err)
	}
	return s
}

//UserController
type UserController struct{
	session *mgo.Session
}

func NewUserContoller(s *mgo.Session) *UserController{
	return &UserController{s}
}


type input struct{
	Id bson.ObjectId `json:"id" bson:"_id"`
	Date string `json:"date" bson:"date"`
	Note string `json:"note" bson:"note"`
}

func (uc UserController) InsertNote(x string, y string){
	input2 := input{Id:bson.NewObjectId(), Date:y, Note:x}

	if err := uc.session.DB("PersonalDB").C("notepad").Insert(input2); err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Note Saved!")
}