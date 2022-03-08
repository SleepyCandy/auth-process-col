package authen

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type header struct {
	Sender    string `header:"sender"`
	Refer     string `header:"refer"`
	Forward   string `header:"forward"`
	Branch    string `header:"branch"`
	SendDate  string `header:"sendDate"`
	Signature string `header:"signature"`
}

type body struct {
	Channal     string `json:"channal"`
	Url         string `json:"url"`
	BodyRequest string `json:"bodyRequest"`
}

type response struct {
	Authentication string `json:"authentication"`
	KongApiKey     string `json:"kongApiKey"`
}

type AuthHandler struct {
}

func (h *AuthHandler) ServiceAuth(c *gin.Context) {
	header := header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(200, err)
	}
	body := body{}
	if err := c.ShouldBindHeader(&body); err != nil {
		c.JSON(200, err)
	}

	privateKey := "mockPrivateKey"
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(200, err)
	}

	hash1 := sha1.New()
	hash1.Write([]byte(header.Sender + header.Refer + header.Forward + header.SendDate + privateKey))
	sha1 := base64.URLEncoding.EncodeToString(hash1.Sum(nil))
	hash1.Write(jsonBody)
	sha2 := base64.URLEncoding.EncodeToString(hash1.Sum(nil))
	hash1.Write([]byte(sha1 + sha2))
	sha3 := base64.URLEncoding.EncodeToString(hash1.Sum(nil))
	fmt.Println(sha1)
	fmt.Println(sha2)
	fmt.Println(sha3)

	if header.Signature == sha3 {
		c.JSON(200, response{Authentication: "Success", KongApiKey: "mockKongApiKey"})
	}
}
