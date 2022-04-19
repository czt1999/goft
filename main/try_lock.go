package main

import (
    "fmt"
    "time"
)

// The following code demonstates how to implement a mutex
// using channel, select branch ang timer

type Mutex struct {
    ch chan int
}

func NewMutex() *Mutex {
    m := &Mutex{ch: make(chan int, 1)}
    return m
}

func (m *Mutex) Lock() {
    m.ch <- 0
}

func (m *Mutex) Unlock() {
    select {
    case <-m.ch:
    default:
        panic("Mutex is already unlocked")
    }
}

func (m *Mutex) TryLock() bool {
    select {
    case m.ch <- 0:
        return true
    default:
        return false
    }
}

func (m *Mutex) TryLockWithTimeout(timeout time.Duration) bool {
    timer := time.NewTimer(timeout)
    select {
    case m.ch <- 0:
        timer.Stop()
        return true
    case <-timer.C:
        return false
    }
}

func (m *Mutex) unlocked() bool {
    return len(m.ch) == 1
}

func main() {
    done := make(chan int)

    m := NewMutex()
    go func() {
        m.Lock()
        fmt.Println("goroutine 1 has got Mutex")
        time.Sleep(3 * time.Second)
        m.Unlock()
        done <- 0
    }()

    time.Sleep(1 * time.Second)

    go func() {
        if m.TryLock() {
            fmt.Println("goroutine 2 TryLock succeed")
        } else {
            fmt.Println("goroutine 2 TryLock failed")
        }
        done <- 0
    }()

    go func() {
        if m.TryLockWithTimeout(3 * time.Second) {
            fmt.Println("goroutine 3 TryLockWithTimeout succeed")
        } else {
            fmt.Println("goroutine 3 TryLockWithTimeout failed")
        }
        done <- 0
    }()

    for i := 0; i < 3; i++ {
        <-done
    }
}
