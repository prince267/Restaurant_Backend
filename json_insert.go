package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	//"io"
	//"io/ioutil"
	//"gopkg.in/mgo.v2"
	dbrepo "MAD_Assignment/dbrepository"
	domain "MAD_Assignment/domain"
	mongoutils "MAD_Assignment/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)

	file, err := os.Open("./restaurant.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var res domain.Restaurant
	count := 0
	scanner := bufio.NewScanner(file) //read file line by line
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		jsn := []byte(scanner.Text())
		json.Unmarshal(jsn, &res)
		//fmt.Println(res)
		res.DBID = domain.NewID()
		//fmt.Println(res)
		repoaccess.Store(&res) //store the jason into database
		count = count + 1
	}
	fmt.Println(count)
}
