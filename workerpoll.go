package main
import (
	"fmt"
	"time"
)
func worker(id int, jobs <- chan int, results chan<-int){
	for j := range jobs {
		fmt.Println("worker-", id, " started job ", j)
		time.Sleep(time.Second * 2)
		fmt.Println("worker-", id, " finished job ", j)
		results <- j * 2
	}
}

func main(){
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3;w++{
		go worker(w, jobs, results)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	for a:= 1; a <=10; a++ {
		<-results
	}
}
