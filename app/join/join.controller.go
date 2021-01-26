package join

import (
	"jsc-join-api/handlers"
	"net/http"
)

type JoinController struct {
	joinService JoinService
	response    handlers.Response
}

func ProvideJoinController(JoinService JoinService) JoinController {
	return JoinController{
		joinService: JoinService,
	}
}

func (c *JoinController) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.joinService.Find()
	if err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.response.ResponseJSON(w, http.StatusOK, data)
}
