package auth

import (
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/arthur-rl/applyx/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

func SetupRoutes(app *fiber.App) {
	fmt.Println("Loading Auth Routes")
	auth := app.Group("/auth")
	auth.Post("/register", Register)
	auth.Post("/login", Login)
	auth.Get("/state", State)
}

type RegisterUserRequest struct {
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Permissions datatypes.JSON `json:"permissions"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Register(c *fiber.Ctx) error {
	body := new(RegisterUserRequest)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	if database.DoesUserExistByUsername(body.Username) {
		c.SendStatus(400)
		err := ErrorResponse{
			Error: "A user with that username already exists",
		}
		b, _ := json.Marshal(err)
		c.Send(b)
		return nil
	}

	user := database.CreateNewUser(body.Username, body.Password, body.Email, body.Name, body.Permissions)
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	c.Send(b)
	return nil
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var store *session.Store = session.New()

func Login(c *fiber.Ctx) error {
	body := new(LoginUserRequest)

	if err := c.BodyParser(body); err != nil {
		return err
	}
	if database.DoesUserExistByUsername(body.Username) {
		user := database.GetUserByUsername(body.Username)
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err == nil {
			//create user session

			if err != nil {
				c.SendStatus(500)
				return err
			} else {
				sess, _ := store.Get(c)
				gob.Register(user)
				sess.Set("user", user)
				if err := sess.Save(); err != nil {
					panic(err)
				}

			}
			b, _ := json.Marshal(user)
			c.Send(b)
		} else {
			err := ErrorResponse{
				Error: "Incorrect username or password",
			}
			b, _ := json.Marshal(err)
			c.Send(b)
			return c.SendStatus(401)
		}
	} else {
		err := ErrorResponse{
			Error: "Incorrect username or password",
		}
		b, _ := json.Marshal(err)
		c.Send(b)
		return c.SendStatus(401)
	}
	return nil
}

func State(c *fiber.Ctx) error {
	sess, _ := store.Get(c)
	fmt.Println(sess.ID())
	user := sess.Get("user")
	b, _ := json.Marshal(user)
	if user == nil {
		err := ErrorResponse{
			Error: "Session not found",
		}
		b, _ := json.Marshal(err)
		c.Send(b)
		return c.SendStatus(401)
	}
	c.Send(b)
	return nil
}
