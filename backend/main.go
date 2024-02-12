// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(1)

// 	d := time.Now().Add(3 * time.Second)
// 	ctx, cancel := context.WithDeadline(context.Background(), d)

// 	go PrintTick(ctx)

// 	time.Sleep(time.Second * 5)
// 	cancel()

// 	wg.Wait()
// }

// func PrintTick(ctx context.Context) {
// 	tick := time.Tick(time.Second)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Done:", ctx.Err())
// 			wg.Done()
// 			return
// 		case <-tick:
// 			fmt.Println("tick")
// 		}
// 	}
// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("init")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// time.Sleep(time.Second * 15)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
