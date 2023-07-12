package userhandler

import (
	"go-do/internal/database"
	"go-do/internal/user"
	"go-do/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserGet(
	s database.Storer[user.UserModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		t := "users"

		if !s.Check(user.UserModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		q := c.Query("id")
		i, err := strconv.Atoi(q)
		if err != nil {
			i = 0
		}

		u := user.UserModel{Id: i}
		if s.Check(u, t) {
			s.Scan(&u, t)
		} else {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, u)
	}
}

func UserPost(
	s database.Storer[user.UserModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := "users"

		if !s.Check(user.UserModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		b := user.UserPost{}
		c.Bind(&b)
		q := b.ToUserModel()
		s.Add(q, t)

		c.Status(http.StatusNoContent)
	}
}

func UserPatch(
	s database.Storer[user.UserModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := "users"

		if !s.Check(user.UserModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		b := user.UserPost{}
		c.Bind(&b)
		q := b.ToUserModel()

		if !s.Check(q, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		s.Update(q, t)

		c.Status(http.StatusNoContent)
	}
}

func UserDelete(
	s database.Storer[user.UserModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Query("table")

		if !s.Check(user.UserModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		q := c.Query("id")
		i, err := strconv.Atoi(q)
		util.ErrLog(err)

		u := user.UserModel{Id: i}

		if i == 0 || s.Check(u, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		s.Delete(u, t)

		c.Status(http.StatusNoContent)
	}
}
