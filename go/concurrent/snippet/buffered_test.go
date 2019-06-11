package snippet

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// buffered channel usecase
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		log.Println("wrote ", i, "to ch")
	}
	// close writing
	close(ch)
}

func TestWrite(t *testing.T) {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2e9)
	// read from channel, need to close the channel after the writing operation
	for v := range ch {
		log.Println("read ", v, "from ch")
		time.Sleep(2e9)
	}
}

// waitgroup usecase
func process(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("start ", i)
	time.Sleep(1e9)
	log.Println("end ", i)
}

func TestProcess(t *testing.T) {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	log.Println("all finished.")
}

// worker pool usecase
// a worker pool listens to a job channel
// a result channel to contain the result of job done by the worker from the pool
type Job struct {
	id       int
	randomno int
}

type Result struct {
	job Job
	sum int
}

// add noOfJobs to the jobs chan
func allocate(jobs chan Job, noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		jobs <- Job{i, rand.Intn(999)}
	}
	close(jobs)
}

// print the result from results chan
func consumeResults(results chan Result, done chan struct{}) {
	for result := range results {
		log.Printf("id=%d,num=%d,sum=%d", result.job.id, result.job.randomno, result.sum)
	}
	done <- struct{}{}
}

// worker pool
func createWorkerpool(results chan Result, jobs chan Job, noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(results, jobs, &wg)
	}
	wg.Wait()
	close(results)
}

// do the job from jobs chan, add the result to the results chan
func worker(results chan<- Result, jobs <-chan Job, wg *sync.WaitGroup) {
	for job := range jobs {
		results <- Result{job, digits(job.randomno)}
	}
	wg.Done()
}

// digits is the logic of the job
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		sum += no % 10
		no /= 10
	}
	time.Sleep(1e9)
	return sum
}

// time related utility
var startTime = time.Now()

func now() time.Duration {
	return time.Since(startTime)
}

// go test -v -run=TestPool
func TestPool(t *testing.T) {
	// use the buffered chan to create worker pool
	var jobs = make(chan Job, 10)
	var results = make(chan Result, 10)
	noOfJobs := 100
	noOfWorkers := 10
	done := make(chan struct{})
	// start
	s := now()
	// during
	go allocate(jobs, noOfJobs)
	go consumeResults(results, done)
	// worker pool running
	createWorkerpool(results, jobs, noOfWorkers)
	<-done
	// end
	e := now()
	log.Println("total=", (e - s).Seconds(), "s")
}
