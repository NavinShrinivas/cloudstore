package authentication

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"products/communication"

	"github.com/gin-gonic/gin"
)

//Need to call the userhandle server for auth and should return back the user claims

func CheckUserAuthMiddleware(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5001/auth", nil)
	current_token_header := c.Request.Header["Token"]
	req.Header = http.Header{
		"Token": {current_token_header[0]},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	resp_struct := communication.AuthResponse{}
	json.Unmarshal(body, &resp_struct)
	if resp_struct.Status {
		c.Next()
		return
	} else {
		c.JSON(http.StatusNetworkAuthenticationRequired, resp_struct)
		c.Abort()
		return
	}

}
