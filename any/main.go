package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num < 3; num++ {
		printCh <- num
	}

	defer cancelCtx()

	// doSomething() is done
	time.Sleep(100 * time.Second)
	fmt.Printf("doSomething: finished\n")

}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if er := ctx.Err(); er != nil {
				fmt.Printf("doAnother: %v\n", er)
			}

			fmt.Printf("doAnother: finished\n")
			return

		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
			return
		}
	}
}

func buildAnotherContext(ctx context.Context) string {
	ctx, cancelCtx := context.WithCancel(ctx)

	go func() {
		cancelCtx()
	}()

	return ctx.Value("myKey").(string)
}
