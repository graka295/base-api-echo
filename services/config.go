package services

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// GetConfiguration get function from conf/config.json
func GetConfiguration(code string) string {
	e := echo.New()
	viper.SetConfigType("json")
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	return viper.GetString(code)
}

// Message get message for message.json
func Message(code string, data map[string]interface{}) string {
	country := "id"
	viper.SetConfigType("json")
	viper.AddConfigPath("./resources/lang/" + country)
	viper.SetConfigName("global.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("err Message" + err.Error() + " code = " + code + " os.Getenv(ENV)" + os.Getenv("ENV"))
		return ""
	}
	text := viper.GetString(code)

	var tpl bytes.Buffer
	t, _ := template.New("").Parse(text)
	err = t.Execute(&tpl, data)
	if err != nil {
		log.Println("err Message" + err.Error() + " code = " + code + " os.Getenv(ENV)" + os.Getenv("ENV"))
		return ""
	}
	return tpl.String()
}
