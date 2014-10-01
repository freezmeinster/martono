package main

import (
    "fmt"
    "github.com/go-martini/martini"
)

func HelloAction(params martini.Params) string {
    return fmt.Sprintf("Hallo %s", params["name"])
}

func main(){
    m := martini.Classic()
    m.Get("/",func() string{
        return "Nguknguk"
    })
    
    m.Group("/tools", func(r martini.Router){
        r.Get("/hello/:name", HelloAction)
    })
    m.Run()
}