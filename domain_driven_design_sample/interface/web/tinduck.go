package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/2019/PPLB4/back-end/application"
)

func getDuck(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	duck, err := application.GetTinduck().GetDuck(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, duck)
}

func getDucks(c *gin.Context) {
	ducks, err := application.GetTinduck().GetDucks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, ducks)
}

func addDuck(c *gin.Context) {
	rawData, _ := c.GetRawData()
	data := struct {
		Name string `json:"name"`
	}{}
	err := json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = application.GetTinduck().AddDuck(data.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func getDuckMatch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	rawData, _ := c.GetRawData()
	data := struct {
		X int `json:"x"`
		Y int `json:"y"`
	}{}
	err = json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	match, err := application.GetTinduck().GetDuckMatch(id, data.X, data.Y)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, match)
}
