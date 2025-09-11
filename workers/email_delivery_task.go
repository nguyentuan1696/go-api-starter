package workers

import (
	"context"
	"fmt"
)

func SendEmail(ctx context.Context, data []byte) error {
	fmt.Println(string(data) + " hehe")
	return nil
}
