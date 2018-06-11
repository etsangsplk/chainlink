package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/smartcontractkit/chainlink/services"
	"github.com/smartcontractkit/chainlink/store/presenters"
)

// ConfigController returns information for the active account
type ConfigController struct {
	App *services.ChainlinkApplication
}

// Show returns the address, plus it's ETH & LINK balance
// Example:
//  "<application>/config"
func (jsc *ConfigController) Show(c *gin.Context) {
	pc := presenters.Config{}
	if json, err := jsonapi.Marshal(pc); err != nil {
		c.AbortWithError(500, fmt.Errorf("failed to marshal config using jsonapi: %+v", err))
	} else {
		c.Data(200, MediaType, json)
	}
}
