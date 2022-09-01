package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	user struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

func CreateUser(c echo.Context) error {

	log.Println("Masuk ke create user")
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++
	log.Println("User berhasil dibuat.", "ID:", u.ID, "Nama:", u.Name, "Email:", u.Email)
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	log.Println("Masuk ke Get User")
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	log.Println("Masuk ke Update User")

	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	log.Println("User berhasil diupdate.", "ID:", u.ID, "Nama:", u.Name, "Email:", u.Email)

	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	log.Println("Masuk ke delete user")

	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	log.Println("User berhasil didelete")
	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	log.Println("Masuk ke cek semua user")
	return c.JSON(http.StatusOK, users)
}
