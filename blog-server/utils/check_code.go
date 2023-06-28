package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Code(length int) string {
	//指定随机数种子
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(10000))
}
