package db

import (
	"fmt"
	"testing"
)

func TestGetRedisInstance(t *testing.T) {
	client := GetRedisInstance(DefaultRedisOption)
	if err := client.Set("k1", "v1", 0).Err(); err != nil {
		fmt.Println(err)
	}
	val, _ := client.Get("k1").Result()
	fmt.Println(val)
}
