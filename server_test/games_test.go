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

func pathByIdGames(id string) string {
	return "/games/" + id
}

type ExampleTestSuite2 struct {
	suite.Suite
	gameTeste *models.Games
	e         *httpexpect.Expect
}

func TestExampleTestSuite2(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite2))
}

func ZerarNanoGames(x models.Games) (y models.Games) {
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
func (suite *ExampleTestSuite2) SetupTest() {
	suite.e = httpexpect.New(suite.T(), "http://localhost:5000/api/v1")
	data, _ := random.String(20)
	randint, _ := random.GetInt(100000)
	x := models.Games{
		Title:     data,
		Genre:     "Testing",
		Publisher: "Joao",
		Price:     float32(randint),
	}
	fmt.Println("1234")
	json.Unmarshal([]byte(suite.e.POST("/games").WithJSON(x).Expect().Body().Raw()), &suite.gameTeste)
	fmt.Println(suite.gameTeste)
}

func (suite *ExampleTestSuite2) TestGetGames() {

	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	fmt.Println(err)

	suite.T().Run("Testando Get (Games)", func(t *testing.T) {
		t.Logf(suite.gameTeste.Genre)
		arm := suite.e.GET("/games").Expect().Body().Raw()
		assert.Equal(t, 200, suite.e.GET("/games").Expect().Raw().StatusCode)
		var response []models.Games
		json.Unmarshal([]byte(arm), &response)
		fmt.Println(response[0].Title)
		var game models.Games

		_ = database.Table("games").Where("genre = ? AND id = ?", "Adventure", "1").Find(&game)
		//fmt.Println(book)

		assert.Equal(t, game, response[0])

	})
	suite.T().Run("Testando chaves (Games)", func(t *testing.T) {
		obj := suite.e.GET("/games/1").Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("title", "genre", "price", "publisher", "created", "updated", "deleted", "id")
		obj.Value("genre").String().Equal("Testing")
		if obj.Value("author").String().Raw() != "Joao" {
			t.Error("Deu erro forte 2 :)")
		}
	})

}
func (suite *ExampleTestSuite2) TestPostGames() {
	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	fmt.Println(err)

	suite.T().Run("Testando Post (Games)", func(t *testing.T) {
		arm := suite.e.POST("/games").WithJSON(suite.gameTeste).Expect()
		//assert.Equal(t, 200, suite.e.POST("/books").WithJSON(suite.bookTeste).Expect().Raw().StatusCode)
		var response models.Games
		json.Unmarshal([]byte(arm.Body().Raw()), &response)
		var game models.Games
		s := strconv.FormatUint(response.ID, 10)
		_ = database.Table("games").Where("id = ?", s).Find(&game)
		response = ZerarNanoGames(response)
		game = ZerarNanoGames(game)
		fmt.Println(game)
		fmt.Println(response.DeletedAt)
		assert.Equal(t, game, response)
		t.Logf(suite.e.POST("/games").WithJSON(suite.gameTeste).Expect().Body().Raw())

	})
}
func (suite *ExampleTestSuite2) TestDeleteGames() {
	database, _ := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=books sslmode=disable password=postgres"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	database.DB()
	//fmt.Println(err)

	suite.T().Run("Testando delete (Game)", func(t *testing.T) {
		var s string = strconv.FormatUint(uint64(suite.gameTeste.ID), 10)
		arm := suite.e.GET(pathByIdGames(s)).Expect().Body().Raw()
		if suite.e.DELETE(pathByIdGames(s)).Expect().Raw().StatusCode != 204 {
			t.Error("Deu erro forte :)")

		}
		var response models.Games
		json.Unmarshal([]byte(arm), &response)
		fmt.Println(response.DeletedAt)
		var game models.Games
		_ = database.Table("games").Where("id = ?", s).Find(&game)
		if game.DeletedAt == nil {
			fmt.Println("Erro ao deletar jogo com id:", s)
		}

		//assert.Contains(t, )

	})
	suite.T().Run("Testando delete com id não criado (Games)", func(t *testing.T) {
		//var s string = strconv.FormatUint(uint64(suite.bookTeste.ID), 10)
		if suite.e.DELETE(pathByIdGames("20")).Expect().Raw().StatusCode != 400 {
			t.Error("Deu erro forte :)")

		}

	})
	suite.T().Run("Testando delete com id de string", func(t *testing.T) {
		//var s string = strconv.FormatUint(uint64(suite.bookTeste.ID), 10)
		if suite.e.DELETE(pathByIdGames("oi")).Expect().Raw().StatusCode != 204 {
			t.Error("Deu erro forte :)")

		}

	})
	suite.T().Run("Testando delete com id de string", func(t *testing.T) {

	})

}
func (suite *ExampleTestSuite2) TestExemplo() {
	return
}
