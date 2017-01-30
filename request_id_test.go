package requestid_test

import (
	requestid "github.com/michele/echo-requestid"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRequestIdMiddleware(t *testing.T) {
	e := echo.New()
	next := func(c echo.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req, _ := http.NewRequest(echo.GET, "http://localhost:8080/ping", nil)

	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	requestid.RequestIdMiddleware(next)(c)

	rid := c.Get("RequestId").(string)

	_, err := uuid.Parse(rid)

	assert.Equal(t, nil, err)
	assert.Equal(t, rid, c.Request().Header.Get("X-Request-Id"))
	assert.Equal(t, rid, res.Header().Get("X-Request-Id"))

	testUuid := "3e638683-b36b-4dd5-b0df-bf9cbc291b7c"
	req.Header.Set("X-Request-Id", testUuid)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	requestid.RequestIdMiddleware(next)(c)
	rid = c.Get("RequestId").(string)
	assert.Equal(t, testUuid, rid)
	assert.Equal(t, rid, c.Request().Header.Get("X-Request-Id"))
	assert.Equal(t, rid, res.Header().Get("X-Request-Id"))
}
