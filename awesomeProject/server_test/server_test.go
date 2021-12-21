package server_test_test

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/mazen160/go-random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"strconv"
	"testing"
	"time"
)

var db *gorm.DB

func pathById(id string) string {
	return "/books/" + id
}

type ExampleTestSuite struct {
	suite.Suite
	bookTeste *models.Books
	e         *httpexpect.Expect
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
func (suite *ExampleTestSuite) SetupTest() {
	suite.e = httpexpect.New(suite.T(), "http://localhost:5000/api/v1")
	data, _ := random.String(20)
	randint,_:= random.GetInt(100000)
	x := models.Books{
		Name:        "LivroTeste",
		Description: data,
		MediumPrice: float32(randint),
		Author:      "Eu",
		ImageURL:    "",
	}
	json.Unmarshal([]byte(suite.e.POST("/books").WithJSON(x).Expect().Body().Raw()), &suite.bookTeste)
	//fmt.Println(suite.bookTeste)
}

func ZerarNano(x models.Books) (y models.Books) {
	x.CreatedAt = time.Date(
		x.CreatedAt.Year(),
		x.CreatedAt.Month(),
		x.CreatedAt.Day(),
		x.CreatedAt.Hour(),
		x.CreatedAt.Minute(),
		x.CreatedAt.Second(),
		0,
		time.UTC,
	)
	x.UpdatedAt = time.Date(
		x.UpdatedAt.Year(),
		x.UpdatedAt.Month(),
		x.UpdatedAt.Day(),
		x.UpdatedAt.Hour(),
		x.UpdatedAt.Minute(),
		x.UpdatedAt.Second(),
		0,
		time.UTC,
	)
	return x
}
func DbOpen() (db *gorm.DB, err error){
	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	fmt.Println(err)
	return db, err

}
func (suite *ExampleTestSuite) TestGet() {

	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	fmt.Println(err)



	suite.T().Run("Testando Get", func(t *testing.T) {
		arm := suite.e.GET("/books").Expect().Body().Raw()
		assert.Equal(t, 200, suite.e.GET("/books").Expect().Raw().StatusCode)
		var response []models.Books
		json.Unmarshal([]byte(arm), &response)
		//fmt.Println(response[0].DeletedAt)
		var book models.Books

		_ = database.Where("name = ? AND id = ?", "LivroTeste", "1").Find(&book)
		//fmt.Println(book)

		assert.Equal(t, book, response[0])

	})
	suite.T().Run("Testando chaves", func(t *testing.T) {
		//values:=[8]string{"name", "description", "medium_price", "author", "image_url", "created_at", "updated_at", "deleted_at"}
		obj := suite.e.GET("/books/1").Expect().Status(http.StatusOK).JSON().Object()
		//assert.Equal(t, values, obj.Keys().ContainsOnly())
		obj.Keys().ContainsOnly("name", "description", "medium_price", "author", "img_url", "created", "updated", "deleted", "id")
		obj.Value("description").String().Equal("Um livro legal")
		if obj.Value("medium_price").Number().Raw() != 0 {
			t.Error("Deu erro forte 2 :)")
		}
	})

}
func (suite *ExampleTestSuite) TestPost() {
	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	fmt.Println(err)

	suite.T().Run("Testando Post", func(t *testing.T) {
		arm := suite.e.POST("/books").WithJSON(suite.bookTeste).Expect()
		//assert.Equal(t, 200, suite.e.POST("/books").WithJSON(suite.bookTeste).Expect().Raw().StatusCode)
		var response models.Books
		json.Unmarshal([]byte(arm.Body().Raw()), &response)
		var book models.Books
		s := strconv.FormatUint(response.ID, 10)
		_ = database.Where("id = ?", s).Find(&book)
		response = ZerarNano(response)
		book = ZerarNano(book)
		fmt.Println(book)
		fmt.Println(response.DeletedAt)
		assert.Equal(t, book, response)
		t.Logf(suite.e.POST("/books").WithJSON(suite.bookTeste).Expect().Body().Raw())

	})
}
func (suite *ExampleTestSuite) TestDelete() {
	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	fmt.Println(err)

	suite.T().Run("Testando delete", func(t *testing.T) {
		var s string = strconv.FormatUint(uint64(suite.bookTeste.ID), 10)
		arm := suite.e.GET(pathById(s)).Expect().Body().Raw()
		if suite.e.DELETE(pathById(s)).Expect().Raw().StatusCode != 204 {
			t.Error("Deu erro forte :)")

		}
		var response models.Books
		json.Unmarshal([]byte(arm), &response)
		fmt.Println(response.DeletedAt)
		var book models.Books
		_ = database.Where("id = ?", s).Find(&book)
		if book.DeletedAt == nil{
			fmt.Println("Erro ao deletar livro com id:", s)
		}

		//assert.Contains(t, )

	})
	suite.T().Run("Testando delete com id não criado", func(t *testing.T) {
		//var s string = strconv.FormatUint(uint64(suite.bookTeste.ID), 10)
		if suite.e.DELETE(pathById("20")).Expect().Raw().StatusCode != 400 {
			t.Error("Deu erro forte :)")

		}

	})
	suite.T().Run("Testando delete com id de string", func(t *testing.T) {
		//var s string = strconv.FormatUint(uint64(suite.bookTeste.ID), 10)
		if suite.e.DELETE(pathById("oi")).Expect().Raw().StatusCode != 204 {
			t.Error("Deu erro forte :)")

		}
	})

}

