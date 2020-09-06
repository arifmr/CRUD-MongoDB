package main

import (
	"context"
	"fmt"

	"latihan-mongo/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	r := gin.Default()

	p := r.Group("/user")
	{
		p.GET("/", Get)
		p.GET("/:id", GetByID)
		p.POST("/create", Create)
		p.PUT("/:id", Update)
		p.DELETE("/:id", Delete)
	}
	r.Run(":8080")
}

func Get(c *gin.Context) {
	client, err := db.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := client.Database("user").Collection("user").Find(context.Background(), bson.M{})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	client, err := db.Mongodb()

	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := client.Database("user").Collection("user").Find(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}
	fmt.Println(result)
	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func Create(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	client, err := db.Mongodb()
	if err != nil {
		c.String(200, err.Error())
		return
	}

	_, err = client.Database("user").Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Update(c *gin.Context) {
	username := c.Param("id")

	var user User
	c.BindJSON(&user)

	client, err := db.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("user").Collection("user").UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": user})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	client, err := db.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("user").Collection("user").DeleteOne(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}
