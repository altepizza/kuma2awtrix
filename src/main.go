package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AwtrixMessage struct {
	Text     string `json:"text"`
	Icon     string `json:"icon"`
	Lifetime int    `json:"lifetime"`
	Duration int    `json:"duration"`
}

func isAppUp(msg string) bool {
	if strings.Contains(msg, "Down") {
		return false
	} else {
		return true
	}
}

func extractKumaMessageFromBody(body []byte) (app string, up bool) {
	var result map[string]interface{}
	var appName string

	json.Unmarshal([]byte(body), &result)
	msg := result["msg"].(string)
	monitoringInfo := result["monitor"].(map[string]interface{})
	appName = monitoringInfo["name"].(string)
	return appName, isAppUp(msg)
}

// setupViper initializes the configuration for Viper.
func setupViper() {
	viper.SetEnvPrefix("KUMA2AWTRIX")
	viper.AutomaticEnv()
}

func main() {
	setupViper()
	opts := mqtt.NewClientOptions().AddBroker(viper.GetString("brokerHost") + ":" + viper.GetString("brokerPort"))
	opts.SetUsername(viper.GetString("username"))
	opts.SetPassword(viper.GetString("password"))
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			c.String(500, "Error reading body")
			return
		}
		c.String(200, "POST done")

		log.Print(extractKumaMessageFromBody(body))

		appName, up := extractKumaMessageFromBody(body)
		var icon string
		if up {
			icon = "up"
		} else {
			icon = "down"
		}

		x := AwtrixMessage{
			Text:     appName + " is " + map[bool]string{true: "up", false: "down"}[up],
			Icon:     icon,
			Lifetime: 86400,
			Duration: 10,
		}

		jsonData, err := json.Marshal(x)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			return
		}

		if token := client.Publish(viper.GetString("AWTRIX_PREFIX")+"/custom/kuma", 0, false, jsonData); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Healthy",
		})
	})

	r.Run(":8181")
}
