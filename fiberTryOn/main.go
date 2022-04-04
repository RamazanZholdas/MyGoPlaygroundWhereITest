package main

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todos struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var todos = []Todos{
	{
		Id:   1,
		Name: "Suck",
		Done: true,
	},
	{
		Id:   2,
		Name: "Kartoshka",
		Done: false,
	},
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Am froma hel")
	})
	app.Get("/getTodos", GetTodos)
	app.Post("/createTodo", CreateTodo)
	app.Get("/getTodoById/:id", GetTodoById)
	app.Delete("deleteTodo/:id", DeleteTodoById)
	app.Put("updateTodo/:id", UpdateTodo)

	app.Listen(":3000")
}

func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}

func GetTodoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	for i := range todos {
		if strconv.Itoa(todos[i].Id) == id {
			return ctx.Status(fiber.StatusFound).JSON(todos[i])
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "id not found",
	})
}

func UpdateTodo(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "passed parameters is not valid",
		})
	}

	type request struct {
		Name *string `json:"name"`
		Done *bool   `json:"done"`
	}

	var body request
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for i := range todos {
		if todos[i].Id == id {
			todos[i].Name = *body.Name
			todos[i].Done = *body.Done
			return ctx.Status(fiber.StatusAccepted).JSON(todos)
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "this todo does not exist",
	})
}

func DeleteTodoById(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for i := range todos {
		if todos[i].Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return ctx.Status(fiber.StatusOK).JSON(todos)
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "id not found",
	})
}

func CreateTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.New("wrong json body fix it"))
	}

	newTodo := Todos{
		Id:   len(todos) + 1,
		Name: body.Name,
		Done: false,
	}

	todos = append(todos, newTodo)

	return ctx.Status(fiber.StatusCreated).JSON(newTodo)
}
