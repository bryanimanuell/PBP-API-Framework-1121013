package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	name := c.Request.URL.Query()["name"]
	age := c.Request.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"
	}
	if age != nil {
		if name != nil {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age='" + age[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"Error": err.Error()})
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, user)
		}
	}

	if len(users) < 5 {
		var response Response
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		c.IndentedJSON(response.Status, response)
	} else {
		c.IndentedJSON(400, gin.H{"Error, Incorrect Array Size": err.Error()})
	}
}

func InsertUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		c.IndentedJSON(400, gin.H{"Failed": err.Error()})
	}
	name := c.Request.Form.Get("name")
	age, _ := strconv.Atoi(c.Request.Form.Get("age"))
	address := c.Request.Form.Get("address")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")

	_, errQuery := db.Exec("INSERT INTO users (name, age, address, email, password) values (?,?,?,?,?)",
		name,
		age,
		address,
		email,
		password,
	)

	if errQuery == nil {
		var response Response
		response.Status = 200
		response.Message = "Success"
		c.IndentedJSON(response.Status, response)
	} else {
		fmt.Println(errQuery)
		c.IndentedJSON(400, gin.H{"Insert Failed": err.Error()})
	}
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		return
	}

	name := c.Request.Form.Get("name")
	age, _ := strconv.Atoi(c.Request.Form.Get("age"))
	address := c.Request.Form.Get("address")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	userId := c.Param("user_id")

	var user User
	user.ID, _ = strconv.Atoi(userId)
	user.Name = name
	user.Age = age
	user.Address = address
	user.Email = email
	user.Password = password

	_, errQuery := db.Exec("UPDATE users SET name = ?, age = ?, address = ?, email = ?, password = ? WHERE id = ?",
		name,
		age,
		address,
		email,
		password,
		userId,
	)

	if errQuery == nil {
		var response Response
		response.Status = 200
		response.Message = "Success"
		c.IndentedJSON(response.Status, response)
	} else {
		fmt.Println(errQuery)
		c.IndentedJSON(400, gin.H{"Insert Failed": err.Error()})
	}
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		return
	}
	userId := c.Param("user_id")

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?", userId)

	if errQuery == nil {
		c.IndentedJSON(200, "Delete Success")
	} else {
		fmt.Println(errQuery)
		c.IndentedJSON(400, gin.H{"Delete Failed": err.Error()})
	}
}
