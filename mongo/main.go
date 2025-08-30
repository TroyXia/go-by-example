package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User 示例文档结构
type User struct {
	ID       string    `bson:"_id,omitempty"`
	Name     string    `bson:"name"`
	Email    string    `bson:"email"`
	Age      int       `bson:"age"`
	CreateAt time.Time `bson:"createAt"`
}

func main() {
	// MongoDB连接URI
	uri := "mongodb://iam:iam59!z$@127.0.0.1:27017/iam_analytics?authSource=iam_analytics"

	// 连接到MongoDB
	client, err := connectToMongoDB(uri)
	if err != nil {
		log.Fatalf("无法连接到MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("断开连接时出错: %v", err)
		}
	}()

	fmt.Println("成功连接到MongoDB!")

	// 获取数据库和集合
	db := client.Database("iam_analytics")
	coll := db.Collection("users")

	// 创建一个示例用户
	sampleUser := User{
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Age:      30,
		CreateAt: time.Now(),
	}

	// 测试增删改查操作
	fmt.Println("\n===== 测试增删改查操作 =====")

	// 1. 插入文档
	insertedID, err := insertDocument(coll, sampleUser)
	if err != nil {
		log.Printf("插入文档失败: %v", err)
	} else {
		fmt.Printf("插入文档成功，ID: %v\n", insertedID)
	}

	// 2. 查询文档
	var result User
	filter := bson.M{"name": "张三"}
	err = findOneDocument(coll, filter, &result)
	if err != nil {
		log.Printf("查询文档失败: %v", err)
	} else {
		fmt.Printf("查询文档成功: %+v\n", result)
	}

	// 3. 更新文档
	update := bson.M{
		"$set": bson.M{"age": 31, "email": "zhangsan_new@google.com"},
	}
	updateCount, err := updateDocument(coll, filter, update)
	if err != nil {
		log.Printf("更新文档失败: %v", err)
	} else {
		fmt.Printf("更新文档成功，影响行数: %d\n", updateCount)
	}

	// 4. 查询所有文档
	var users []User
	err = findAllDocuments(coll, bson.M{}, &users)
	if err != nil {
		log.Printf("查询所有文档失败: %v", err)
	} else {
		fmt.Printf("查询所有文档成功，共找到 %d 个文档\n", len(users))
	}

	// 5. 删除文档
	deleteCount, err := deleteDocument(coll, filter)
	if err != nil {
		log.Printf("删除文档失败: %v", err)
	} else {
		fmt.Printf("删除文档成功，影响行数: %d\n", deleteCount)
	}
}

// connectToMongoDB 连接到MongoDB
func connectToMongoDB(uri string) (*mongo.Client, error) {
	// 设置连接选项
	clientOptions := options.Client().ApplyURI(uri)

	// 连接超时设置为10秒
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// insertDocument 插入单个文档
func insertDocument(coll *mongo.Collection, document interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// findOneDocument 查询单个文档
func findOneDocument(coll *mongo.Collection, filter interface{}, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return coll.FindOne(ctx, filter).Decode(result)
}

// findAllDocuments 查询所有匹配的文档
func findAllDocuments(coll *mongo.Collection, filter interface{}, results interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

// updateDocument 更新文档
func updateDocument(coll *mongo.Collection, filter interface{}, update interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

// deleteDocument 删除文档
func deleteDocument(coll *mongo.Collection, filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
