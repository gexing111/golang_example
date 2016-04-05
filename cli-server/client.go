package main

import "log"
import "net/rpc"
import "fmt"
import "time"

func main() {
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var args = "hello rpc"
    var reply string
    err = client.Call("Echo.Hi", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("reply\n",reply)
    time.Sleep(10000*time.Second)
}
