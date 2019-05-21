# Restaurant_Backend
Install MongoDB in your system.

https://tecadmin.net/install-mongodb-on-ubuntu/



Install GO and mgo package in your system.

https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79

https://gopkg.in/mgo.v2


Clone the Repository
and type "go get" in your terminal. It will install all the packages for that.

# Command to Add Sample Restaurant in MongoDB
```bash
~ go run json_insert.go
```

# Command to Run
```bash
~cd /restapplication
~go build main.go
~./main
```

Open the Browser and type

To get All the Restaurant Data in JSON Format
http://localhost:8080/restaurant/       

---------------------------------------------------------------
To get the Restaurant Data of given id  in JSON Format
http://localhost:8080/restaurant/id_of_restaurant       

Example 
http://localhost:8080/restaurant/5c9d06dc2077e45c3322198f

----------------------------------------------------------------
To get the Restaurant Data of given name in JSON Format
http://localhost:8080/restaurant/?name=Restaurant_name

Example
http://localhost:8080/restaurant/?name=prce123

----------------------------------------------------------------
To get the Restaurant Data of given Type of Food in JSON Format
http://localhost:8080/restaurant/?type_Of_food=Food_Type

Example
http://localhost:8080/restaurant/?type_Of_food=Chinese

----------------------------------------------------------------
To get the Restaurant Data of given Postcode in JSON Format
http://localhost:8080/restaurant/?postcode=Restaurant_Postcode

Example
http://localhost:8080/restaurant/?postcode=123

You can use Postman for hitting the API's
