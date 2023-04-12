package authentication

import (
<<<<<<< HEAD
	// "fmt"
	"encoding/json"
	"io/ioutil"
=======
	"encoding/json"
	"io/ioutil"
	"log"
>>>>>>> b1491b0 (order mservice starting)
	"net/http"
	"orders/communication"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	log "github.com/urishabh12/colored_log"
)

func CheckUserAuthMiddleware(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5001/authcheck", nil)
	current_token_header := c.Request.Header["Token"]
	if len(current_token_header) <
		1 {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"status":  false,
			"message": "Please provide a proper token in header!",
		})
		c.Abort()
		return
	}
	req.Header = http.Header{
		"Token": {current_token_header[0]},
	}
=======
)

//Need to call the userhandle server for auth and should return back the user claims

func CheckUserAuthMiddleware(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5001/authcheck", nil)
	current_token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in cookie, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}
	req.AddCookie(&http.Cookie{Name: "token", Value: current_token})
>>>>>>> b1491b0 (order mservice starting)
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

func GetClaims(c *gin.Context) *communication.AuthResponse {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5001/authcheck", nil)
<<<<<<< HEAD
	current_token_header := c.Request.Header["Token"]
	req.Header = http.Header{
		"Token": {current_token_header[0]},
	}
=======
	current_token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in cookie, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return nil
	}
	req.AddCookie(&http.Cookie{Name: "token", Value: current_token})
>>>>>>> b1491b0 (order mservice starting)
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
		return &resp_struct
	} else {
		c.JSON(http.StatusNetworkAuthenticationRequired, resp_struct)
		c.Abort()
		return nil
	}
}
