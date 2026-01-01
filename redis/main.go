package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type SessionData struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	opt, err := redis.ParseURL("")

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	defer rdb.Close()

	ctx := context.Background()

	sessionID, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	sessionData := SessionData{
		UserID:   123,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	jsonData, err := json.Marshal(sessionData)
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, sessionID.String(), jsonData, 1*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Session created with 1 minute TTL")

	val, err := rdb.Get(ctx, sessionID.String()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Session value (JSON):", val)
	var retrievedSession SessionData
	err = json.Unmarshal([]byte(val), &retrievedSession)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized: UserID=%d, Username=%s, Email=%s\n",
		retrievedSession.UserID, retrievedSession.Username, retrievedSession.Email)

	ttl, err := rdb.TTL(ctx, sessionID.String()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Remaining TTL: %v\n", ttl)

	time.Sleep(15 * time.Second)
	rdb.Expire(ctx, sessionID.String(), 5*time.Minute)
	fmt.Println("Session refreshed")
}
