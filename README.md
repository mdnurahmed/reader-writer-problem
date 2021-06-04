# reader-writer-problem
This is a concurrency problem from the book [The Little Book of Semaphores](https://greenteapress.com/wp/semaphores/)

## Problem Description 

Any situation where a data structure, database, or file system is read and modified by concurrent threads. While the data structure is being written or modified it is often necessary to bar other threads from reading, in order to prevent a reader from interrupting a modification in progress and reading inconsistent or invalid data.”

In other words, any number of readers can read the shared resource and only one writer can write to it at a time.

The implementation has to be such that if a number of readers are holding the lock, and a writer tries to acquire the lock, it will wait until all the readers release it and it prevents more readers from acquiring it in the meantime. This also have to solve the writer“Starvation” problem : a writer can try to acquire the lock and fails to do so because readers keep coming in and out without ever reaching the situation where all of them have released it at the same time.

So the writers have to have priority over the readers.

## My Approach 
We can use the built in RWMutex from sync package to solve this. What sync.RWMutex does is it allows for multiple readers to acquire the read lock (RWMutex.RLock()), but once there’s a writer trying to acquire the write lock (RWMutex.Lock()), then no reader will get the lock until the writer acquires it and releases it again. If writers queue up trying to acquire the lock, they will be prioritized over the waiting readers.

## How to run 
using go 
```
go run reader-writer.go
```
or using docker 
```
docker-compose -f reader-writer.yaml up --build
docker run -it reader-writer-problem_reader-writer /bin/sh
go run reader-writer.go
exit
docker-compose -f reader-writer.yaml down
```