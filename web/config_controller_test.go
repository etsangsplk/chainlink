package web_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/internal/cltest"
	"github.com/smartcontractkit/chainlink/store/presenters"
	"github.com/smartcontractkit/chainlink/web"
	"github.com/stretchr/testify/assert"
)

func TestConfigController_Show(t *testing.T) {
	t.Parallel()

	app, cleanup := cltest.NewApplication()
	defer cleanup()

	resp := cltest.BasicAuthGet(app.Server.URL + "/config")
	cltest.AssertServerResponse(t, resp, 200)

	c := presenters.Config{}
	err = web.ParseJSONAPIResponse(cltest.ParseResponseBody(resp), &c)
	assert.NoError(t, err)
}
