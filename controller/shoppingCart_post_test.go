package controller

import (
	"alta/project/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func testPostShoppingCartController(t *testing.T, shoppingCartController echo.HandlerFunc) {
	// try to request
	reqBody, err := json.Marshal(model.Shopping_cart{
		User_id:    1,
		Product_id: 1,
		Name:       "Redmi Note 10",
		Category:   "Handphone",
		Type:       "Xiaomi",
		Price:      10000000,
		Qty:        1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	req, err := http.NewRequest("POST", "carts/:user_id/:product_id", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	shoppingCartController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var cart model.Shopping_cart
	if err := json.Unmarshal(body, &cart); err != nil {
		t.Error(err)
	}
	if cart.Name != "Redmi Note 10" {
		t.Errorf("expected Redmi Note 10, got : %#v", cart.Name)
	}
}

func TestDBPostShoppingCartController(t *testing.T) {
	// create DB
	test_connectionString := "root:02021996Doni*@tcp(localhost:3306)/alta_shop_project_testing?charset=utf8&parseTime=True&loc=Local" // doni local computer
	var err error
	DB, err := gorm.Open(mysql.Open(test_connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.Shopping_cart{})
	DB.Delete(&model.Shopping_cart{}, "1=1")
	m := model.NewGormShoppingCartModel(DB)
	// create controller
	shoppingCartController := CreatePostShoppingCartController(m)
	if err != nil {
		t.Error(err)
	}
	testPostShoppingCartController(t, shoppingCartController)
	DB.Delete(&model.Shopping_cart{}, "1=1")
}
