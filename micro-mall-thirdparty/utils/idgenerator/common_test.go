package idgenerator

import (
	"fmt"
	"testing"
)

func TestGetSnowflakeId(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetSnowflakeId())
	}
}
