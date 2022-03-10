package controller

import (
	"fmt"
	"github.com/MickStanciu/SC/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JediController interface {
	GetJediById(c *gin.Context)
}

type jediController struct {
	service service.JediService
}

func (ct *jediController) GetJediById(c *gin.Context) {
	jediID := strings.TrimSpace(c.Param("id"))

	jedi, err := ct.service.GetJediById(jediID)
	if err != nil {
		fmt.Sprintf("we have error %v", err.Message)
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, jedi)
}

func NewJediController(s service.JediService) JediController {
	return &jediController{
		service: s,
	}
}
