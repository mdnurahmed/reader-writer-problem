package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var collection []int
var mutex sync.RWMutex

func writer(id int) {
	defer mutex.Unlock()
	defer func() {
		log.Printf("writer %d released write lock", id)
	}()
	mutex.Lock()
	log.Printf("writer %d acquired write lock", id)
	value := rand.Intn(1000)
	index := rand.Intn(len(collection))
	log.Printf("writing %d at index-%d", collection[index], value)
	collection[index] = value
	time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)
}

func reader(id int) {
	defer func() {
		log.Printf("reader %d released read lock", id)
	}()
	defer mutex.RUnlock()
	mutex.RLock()
	log.Printf("reader %d acquired read lock\n", id)
	log.Printf("printing array %+v\n", collection)
	time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)
}

func main() {
	for i := 1; i <= 10; i++ {
		val := rand.Intn(10)
		collection = append(collection, val)
	}
	log.Printf("original array %+v\n", collection)
	go func() {

		for id := 1; ; id++ {
			time.Sleep(time.Duration(rand.Intn(4)+2) * time.Second)
			go writer(id)
		}
	}()
	for id := 1; ; id++ {
		time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)
		go reader(id)
	}
}
