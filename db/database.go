package db

import (
  "github.com/go-martini/martini"
  "labix.org/v2/mgo"
)

type User struct {
  Name string `form:"name"`
  Email string `form:"email"`
  Password string `form:"password"`
}

type Wish struct {
  Name string `form:"name"`
  Description string `form:"description"`
}

func DB() martini.Handler {
  session, err := mgo.Dial("mongodb://localhost")
  if err != nil {
    panic(err)
  }

  return func(c martini.Context) {
    s := session.Clone()
    c.Map(s.DB("advent"))
    defer s.Close()
    c.Next()
  }
}

func GetAll(db *mgo.Database) []Wish {
  var wishlist []Wish
  db.C("wishes").Find(nil).All(&wishlist)
  return wishlist
}

func GetUsers(db *mgo.Database) []User {
  var users []User
  db.C("users").Find(nil).All(&users)
  return users
}
