package main

import (
    "fmt"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func HelloAction(params martini.Params) string {
    return fmt.Sprintf("Hallo %s", params["name"])
}

func RenderAction(r render.Render){
   r.HTML(200, "index", "Bram")
}


func main(){
    m := martini.Classic()
    m.Use(render.Renderer())
    
    m.Get("/",func() string{
        return "Nguknguk"
    })
    
    m.Group("/tools", func(r martini.Router){
        r.Get("/hello/:name", HelloAction)
        r.Get("/render", RenderAction)
    })
    m.Run()
}