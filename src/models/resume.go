package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	resultado *Resume
)

func GetOneResume(uid string) (u *Resume, err error) {

	mongodbconf := beego.AppConfig.String("mongodb")
	mcollectionconf := beego.AppConfig.String("mcollection")
	mdocumentconf := beego.AppConfig.String("mdocument")

	session, err := mgo.Dial(mongodbconf)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mcollectionconf).C(mdocumentconf)
	result := Resume{}
	err = c.Find(bson.M{"resumeid": uid}).One(&result)
	if err != nil {
		return nil, errors.New("User not exists")
	}
	resultado = &result

	return resultado, nil
}

type Resume struct {
	Resumeid       string				`json:"resumeid"`
	Name           string               `json:"name"`
	Profile        string               `json:"profile"`
	Description    string               `json:"description"`
	Urlimgbanner   string               `json:"urlimgbanner"`
	Urlimgaboutme  string               `json:"urlimgaboutme"`
	Emailcontact   string               `json:"emailcontact"`
	Skills         []Skill              `json:"skills"`
	Languages      []Language           `json:"languages"`
	Portfolios     []Portfolio          `json:"portfolios"`
	Somestats      []Somestat           `json:"somestats"`
	Experiences    []Experience         `json:"experiences"`
	Educations     []Education          `json:"educations"`
	Socialnetworks []Socialnetwork      `json:"socialnetworks"`
}

type Skill struct {
	Nameskill   string		`json:"nameskill"`			
	Levelskill  string      `json:"levelskill"`
	Growup      string      `json:"growup"`
	Ordernumber string      `json:"ordernumber"`
}

type Language struct {
	Namelanguages  string	`json:"namelanguages"`
	Levellanguages string   `json:"levellanguages"`
	Growup         string   `json:"growup"`
	Ordernumber    string   `json:"ordernumber"`
}

type Portfolio struct {
	Portfoliocategorieid   string          `json:"portfoliocategorieid"`
	Portfoliocategoriename string          `json:"portfoliocategoriename"`
	Portfolioitems         []Portfolioitem `json:"portfolioitems"`
}

type Somestat struct {
	Name  string 	`json:"name"`
	Value string    `json:"value"`
	Delay string    `json:"delay"`
}

type Experience struct {
	Employer    string  `json:"employer"`
	Ocupation   string  `json:"ocupation"`
	Description string  `json:"description"`
	Ordernumber string  `json:"ordernumber"`
	Period      Periodo	`json:"period"`
}

type Education struct {
	Institution string  `json:"institution"`
	Course      string  `json:"course"`
	Description string  `json:"description"`
	Ordernumber string  `json:"ordernumber"`
	Period      Periodo	`json:"period"`
}

type Periodo struct {
	Datein     string	`json:"datein"`
	Dateout    string   `json:"dateout"`
	Presentday bool     `json:"presentday"`
}

type Socialnetwork struct {
	Name string	`json:"name"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
}

type Portfolioitem struct {
	Nameportfolio string	`json:"nameportfolio"`
	Descportfolio string    `json:"descportfolio"`
	Linkportfolio string    `json:"linkportfolio"`
	Date          string    `json:"date"`
}
