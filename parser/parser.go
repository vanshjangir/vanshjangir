package main

import(
    "fmt"
    "os"
)

func main(){
    fmt.Println("parser it baby!")
    dat, err := os.ReadFile("./blog.md")
    if(err != nil){
        println(err)
    }
    println(dat)
}
