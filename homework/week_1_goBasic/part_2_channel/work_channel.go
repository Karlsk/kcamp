package main

import (
	"context"
	"fmt"
	"time"
)

func Product(ctx context.Context,message chan <- int) {
	go func() {
		i := 0
		for {
			message <- i
			i ++
			time.Sleep(time.Second)
		}
	}()
	select {
	case <- ctx.Done():
		close(message)
	}
}
func Consume(id int,message <- chan int,done chan <- bool){
	for {
		fmt.Printf("goroutines = %d,get message\n",id)
		data,notClose := <- message
		if notClose{
			fmt.Printf("goroutines = %d,message = %v\n",id,data)
			time.Sleep(time.Second)
		}else {
			if len(done) <1{
				done <- true
				break
			}else {
				break
			}
		}
	}
}




func main() {

	var message = make(chan int,10)
	var done = make(chan bool)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*20)
	defer cancel()
	go Product(ctx,message)

	for i := 0; i < 4; i++ {
		go Consume(i,message,done)
	}
	for {
		<- done
		break
	}

	

}
