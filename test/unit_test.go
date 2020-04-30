package test

import (
	"github.com/gavv/httpexpect/v2"
	"goci-example/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnit(t *testing.T) {
	handler := api.GetApiService()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/").WithQuery("num1","1").WithQuery("num2","5").
		Expect().Body().Contains("ready")

	e.GET("/status").WithQuery("num1","1").WithQuery("num2","5").
		Expect().Body().Contains("ready")

	e.GET("/add").WithQuery("num1","1").WithQuery("num2","5").
		Expect().Body().Equal("6")

	e.GET("/add").WithQuery("num1","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/add").WithQuery("num2","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/subtract").WithQuery("num1","1").WithQuery("num2","5").
		Expect().Body().Equal("-4")

	e.GET("/subtract").WithQuery("num1","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/subtract").WithQuery("num2","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/multiply").WithQuery("num1","1").WithQuery("num2","5").
		Expect().Body().Equal("5")

	e.GET("/multiply").WithQuery("num1","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/multiply").WithQuery("num2","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/divide").WithQuery("num1","1").WithQuery("num2","5").
		Expect().Body().Equal("0")

	e.GET("/divide").WithQuery("num1","10").WithQuery("num2","5").
		Expect().Body().Equal("2")

	e.GET("/divide").WithQuery("num1","1").
		Expect().Status(http.StatusBadRequest);

	e.GET("/divide").WithQuery("num2","1").
		Expect().Status(http.StatusBadRequest);


}
