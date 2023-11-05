package controller

import (
	"context"
	"github.com/AminN77/upera_test/history_service/api/dto"
	"github.com/AminN77/upera_test/history_service/internal"
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

func (con *Controller) FetchRevisionsOfOneProduct(c *fiber.Ctx) error {
	var response dto.Response[[]*internal.Revision]
	productID, err := strconv.ParseInt(c.Params("productId"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}
	pageIndex, err := strconv.ParseInt(c.Query("pageIndex"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	res, err := con.srv.FetchRevisionsOfOneProduct(pageSize, pageIndex, productID, context.Background())
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusOK
	response.Result = &res
	return c.JSON(response)
}

func (con *Controller) FetchRevision(c *fiber.Ctx) error {
	var response dto.Response[*internal.Product]
	revisionNumber := c.Params("revisionNumber")
	if revisionNumber == "" {
		response.Status = http.StatusBadRequest
		response.Message = "revisionNumber is empty"
		return c.Status(response.Status).JSON(response)
	}

	res, err := con.srv.FetchRevision(revisionNumber, context.Background())
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusOK
	response.Result = &res
	return c.JSON(response)
}
