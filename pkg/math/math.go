package math

import (
	"fmt"
	"math/rand"
)

// 生成40位uuid
func GenerateUUid() string {

	return fmt.Sprintf("%d", rand.Intn(40))
}
