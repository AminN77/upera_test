package controller

import (
	"github.com/AminN77/upera_test/product_service/api/dto"
	"github.com/AminN77/upera_test/product_service/internal"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type Controller struct {
	srv internal.Service
}

func NewController(srv internal.Service) *Controller {
	return &Controller{
		srv: srv,
	}
}

func (con *Controller) AddProduct(c *fiber.Ctx) error {
	var request *dto.UpsertProductRequest
	var response dto.BaseResponse

	if err := c.BodyParser(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	p := &internal.Product{
		Name:        request.Name,
		Description: request.Description,
		Color:       request.Color,
		Price:       request.Price,
		ImageUrl:    request.ImageUrl,
	}

	if err := con.srv.Add(p); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusCreated
	return c.JSON(response)
}

func (con *Controller) UpdateProduct(c *fiber.Ctx) error {
	var request *dto.UpsertProductRequest
	var response dto.BaseResponse

	if err := c.BodyParser(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	p := &internal.Product{
		Name:        request.Name,
		Description: request.Description,
		Color:       request.Color,
		Price:       request.Price,
		ImageUrl:    request.ImageUrl,
	}

	if err := con.srv.Update(p, productID); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusOK
	return c.JSON(response)
}

func (con *Controller) FetchProduct(c *fiber.Ctx) error {
	var response dto.Response[*internal.Product]
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	res, err := con.srv.Fetch(productID)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusOK
	response.Result = &res
	return c.JSON(response)
}