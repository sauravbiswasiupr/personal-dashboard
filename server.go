package main

import (
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  "labix.org/v2/mgo"
  "./database"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  // Database middleware
  m.Use(database.DB())
  m.Get("/", func(r render.Render) {
    r.Redirect("/signup")
  })

  m.Get("/wishes", func(r render.Render, db *mgo.Database) {
    r.HTML(200, "list", database.GetAll(db))
  })

  m.Post("/wishes", binding.Form(database.Wish{}), func(wish database.Wish, r render.Render, db *mgo.Database) {
    db.C("wishes").Insert(wish)
    r.HTML(200, "list", database.GetAll(db))
  })

  m.Post("/signup", binding.Form(database.User{}), func(user database.User, r render.Render, db *mgo.Database) {
    db.C("users").Insert(user)
    r.HTML(200, "user-profile", user)
  })

  m.Get("/signup", func(r render.Render) {
    r.HTML(200, "signup", "")
  })

  m.Get("/users", func(r render.Render, db *mgo.Database) {
    r.HTML(200, "users", database.GetUsers(db))
  })

  m.Run()
}
