package AccountCreation

import (
	"AdvertiserCentral/MySQLDB"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Advertiser struct{
	//AdvId string `json:"advid"`
	AdvCompanyName string `json:"advcompanyname"`
	//AdvAccountStatus string `json:"advaccountstatus"`
	AdvEmail string `json:"advemail"`
	AdvContactNumber string `json:advcontactnumber`
	//AdvSignupDate string `json:advsignupdate`
	//AdvAccountBal float64 `json:advaccountbal`
}

func CreateAccount(w http.ResponseWriter, r *http.Request){
	AdvertiserX := Advertiser{}
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"Error reading the data from browser!")
	}
	fmt.Fprintf(w,"Please enter values in this format, {\"advcompanyname\":\",\"advemail\":\"\",\"AdvContactNumber\":\"\"}")
	json.Unmarshal(reqbody, &AdvertiserX)

	w.WriteHeader(http.StatusCreated)



	insert, err := MySQLDB.ConnectToMySQLDB().Query("insert into advertiser_signup(adv_companyname, adv_email, adv_contactnum) values(?,?,?)",strings.TrimSpace(AdvertiserX.AdvCompanyName) ,strings.TrimSpace(AdvertiserX.AdvEmail),strings.TrimSpace(AdvertiserX.AdvContactNumber))
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("Advertiser info stored in MySQL db!")
	defer insert.Close()

	////inserting the adv info in mongodb
	//collection := MongoDB.MongoConn().Database("bluetree").Collection("AdvertiserInfoMgo")
	//insertInMgo, err := collection.InsertOne(context.TODO(), AdvertiserX)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("Inserted a single document:", insertInMgo.InsertedID)
}


