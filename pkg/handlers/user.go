package handlers

import (
	"bytes"
	"encoding/json"
	"text/template"
	"time"

	"github.com/arrrden/cass-docker/pkg/db"
	common "github.com/arrrden/cass-docker/pkg/handlers/common"
	"github.com/arrrden/cass-docker/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Keyspace string
}

func (h *UserHandler) Post(ctx *fiber.Ctx) error {
	user := models.User{}

	body := ctx.Body()
	dec := json.NewDecoder(bytes.NewReader(body))

	if err := dec.Decode(&user); err != nil {
		ctx.Status(500)
		return err
	}

	t := time.Now()

	data, err := addUser(&user, h.Keyspace, &t)
	if err != nil {
		ctx.Status(500)
		return err
	}

	if err := ctx.JSON(data); err != nil {
		ctx.Status(500)
		return err
	}

	ctx.Status(201)
	return nil
}

func (h *UserHandler) Get(ctx *fiber.Ctx) error {
	var user models.User

	id := ctx.Params("id")
	user.Id = id

	t := time.Now()

	data, err := getUser(&user, h.Keyspace, &t)
	if err != nil {
		ctx.Status(500)
		return err
	}

	if err := ctx.JSON(data); err != nil {
		ctx.Status(500)
		return err
	}

	ctx.Status(200)
	return nil
}

func addUser(user *models.User, keyspace string, t *time.Time) (*models.User, error) {
	query := `
		INSERT INTO {{.Keyspace}}.user(name, email, created_at, updated_at, id)
		VALUES (?, ?, ?, ?, ?);
	`

	tmpl, err := template.New("Query").Parse(query)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	uuid, err := common.UUID()
	if err != nil {
		return nil, err
	}
	data := common.Data{
		Id:        uuid,
		CreatedAt: t,
		UpdatedAt: t,
		Keyspace:  keyspace,
	}

	if err = tmpl.Execute(buffer, data); err != nil {
		return nil, err
	}

	err = db.Mutate(buffer.String(), user.Name, user.Email, data.CreatedAt, data.UpdatedAt, data.Id)
	if err != nil {
		return nil, err
	}

	var u models.User
	u.Id = data.Id
	u.CreatedAt = data.CreatedAt
	u.UpdatedAt = data.UpdatedAt
	u.Name = user.Name
	u.Email = user.Email

	return &u, err
}

func getUser(user *models.User, keyspace string, time *time.Time) (*models.User, error) {
	query := `
		SELECT * FROM {{.Keyspace}}.user 
		WHERE id = {{.Id}}
	`

	tmpl, err := template.New("Query").Parse(query)
	if err != nil {
		return nil, err
	}

	data := common.Data{
		Keyspace: keyspace,
		Id:       user.Id,
	}

	buffer := new(bytes.Buffer)

	if err := tmpl.Execute(buffer, data); err != nil {
		return nil, err
	}

	qry := db.FindOne(buffer.String())

	var u models.User
	qry.Scan(&u.Id, &u.CreatedAt, &u.Email, &u.EmailVerified, &u.Name, &u.UpdatedAt)

	return &u, nil
}
