package main

import "fmt"

func Hello() string {
    return "Hello, world"
}

func not_main() {
//func main() {
    fmt.Println(Hello())
}