package authentication

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"products/communication"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//Need to call the userhandle server for auth and should return back the user claims

func InitAuthVariables() {

	envRequired := []string{"USERHANDLE_ADDRESS"}

	_, err := os.Stat(".env")
	if err == nil {
		secret, err := godotenv.Read()
		if err != nil {
			log.Panic("Error reading .env file")
		}

		for _, key := range envRequired {
			if secret[key] != "" {
				os.Setenv(key, secret[key])
			}
		}
	}

	for _, key := range envRequired {
		if os.Getenv(key) == "" {
			log.Panic("Environment variable " + key + " not set")
		}
	}
}

func CheckUserAuthMiddleware(c *gin.Context) {

	userhandle_address := os.Getenv("USERHANDLE_ADDRESS")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+userhandle_address+"/authcheck", nil)

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in cookie, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return
	}

	req.Header.Add("Cookie", "token="+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

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

	userhandle_address := os.Getenv("USERHANDLE_ADDRESS")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+userhandle_address+"/authcheck", nil)

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"message": "Auth token not found in cookie, will report to admins.",
		})
		log.Println("[WARN] Request without any auth attempt tried gaining access!!!")
		c.Abort()
		return nil
	}

	req.Header.Add("Cookie", "token="+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

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
