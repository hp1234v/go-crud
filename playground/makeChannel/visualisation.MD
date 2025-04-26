# 🛠 GoRoutines + Channels Visualisation in Go

---

## 📜 Code Example

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello from Goroutine"
    }()

    message := <-ch
    fmt.Println(message)
}

     |
     V
+----------------------+
| Channel (type: string)|
|-----------------------|
| Buffer (empty)        |
| Waiting List (empty)  |
+----------------------+

+----------------+
|   Main Thread  |
|                |
| ch := make(chan string)  <-- Creates this Channel ------+
+----------------+                                             |
                                                               |
                                                               V
                                                +---------------------------+
                                                |  Channel ch (type: string) |
                                                |---------------------------|
                                                | Data Buffer (empty)        |
                                                | Waiting List (empty)       |
                                                +---------------------------+
                                                               |
     +-------------------------------------------+             |
     |                                           |             |
+------------------+                     +------------------+ |
|  Goroutine 1     |                     | Main Goroutine   | |
|                  |                     |                  | |
| ch <- "Hello"    |                     | <- ch            | |
| (send)           |                     | (receive)        | |
+------------------+                     +------------------+ |
     |                                              |          |
     +----------------------------------------------+----------+