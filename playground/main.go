package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Token struct {
	GUID string
	jwt.StandardClaims
}

var SECRET_KEY string = "secret"

func main() {
	c := gin.Default()

	c.GET("/page1/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	c.GET("/s/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})

	c.Run(":8080")
}

func GenerateTokens(uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &Token{
		GUID: base64.RawURLEncoding.EncodeToString([]byte(uid)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Fatal(err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Fatal(err)
		return
	}

	return token, refreshToken, err
}

/*
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)

	data := User{
		Name:    "Ainur",
		Country: "KZ",
	}

	if err := enc.Encode(data); err != nil {
		log.Fatal(err)
	}

	file.Write(buf.Bytes())
	bslice, _ := os.ReadFile(file.Name())
	fmt.Println(string(bslice))
*/
