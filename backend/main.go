package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Date    string             `bson:"date" json:"date"`
	Content string             `bson:"content" json:"content"`
}

func main() {

	// 创建Gin路由
	r := gin.Default()
	// 配置CORS中间件
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// 使用CORS中间件
	r.Use(cors.New(corsConfig))
	// 连接到本地MongoDB
	client, err := connectToMongoDB()
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	fmt.Println("Connected to MongoDB")

	// 关闭MongoDB客户端
	defer func() {
		err = client.Disconnect(context.Background())
		if err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// 设置路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/events", func(c *gin.Context) {
		getEvents(c, client)
	})
	r.POST("/events", func(c *gin.Context) {
		createEvent(c, client)
	})

	r.GET("/events/:id", func(c *gin.Context) {
		getEvent(c, client)
	})

	r.PUT("/events/:id", func(c *gin.Context) {
		updateEvent(c, client)
	})

	r.DELETE("/events/:id", func(c *gin.Context) {
		deleteEvent(c, client)
	})

	// 启动Gin服务
	r.Run(":3000")
}

func connectToMongoDB() (*mongo.Client, error) {
	// 设置本地MongoDB连接字符串
	mongoURI := "mongodb://localhost:27017/?directConnection=true"

	// 设置客户端选项
	clientOptions := options.Client().ApplyURI(mongoURI)

	// 连接到MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getEvents(c *gin.Context, client *mongo.Client) {
	collection := client.Database("local").Collection("EventDetail")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching events"})
		return
	}

	var events []Event
	err = cursor.All(ctx, &events)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context, client *mongo.Client) {
	collection := client.Database("local").Collection("EventDetail")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var event Event
	c.BindJSON(&event)

	result, err := collection.InsertOne(ctx, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating event"})
		return
	}

	event.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, event)
}

func getEvent(c *gin.Context, client *mongo.Client) {
	collection := client.Database("local").Collection("EventDetail")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var event Event
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)

}

func updateEvent(c *gin.Context, client *mongo.Client) {
	collection := client.Database("local").Collection("EventDetail")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var updatedEvent Event
	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"date":    updatedEvent.Date,
			"content": updatedEvent.Content,
		},
	}

	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating event"})
		}
		return
	}

	var event Event
	err = result.Decode(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding updated event"})
		return
	}

	event.Date = updatedEvent.Date
	event.Content = updatedEvent.Content
	c.JSON(http.StatusOK, event)
}

func deleteEvent(c *gin.Context, client *mongo.Client) {
	collection := client.Database("local").Collection("EventDetail")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting event"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
