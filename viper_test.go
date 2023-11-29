package GoViper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {

	config := viper.New()
	assert.NotNil(t, config)
}

func TestJson(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	//func read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Feri Susmiyanto", config.Get("application.name"))
	assert.Equal(t, 20, config.GetInt("application.umur"))
	assert.Equal(t, false, config.GetBool("application.menikah"))
	assert.Equal(t, "suyono", config.Get("family.ayah"))
}

func TestYaml(t *testing.T) {
	config := viper.New()
	//config.SetConfigName("config")
	//config.SetConfigType("json")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	//func read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Feri Susmiyanto", config.Get("application.name"))
	assert.Equal(t, 20, config.GetInt("application.umur"))
	assert.Equal(t, false, config.GetBool("application.menikah"))
	assert.Equal(t, "suyono", config.Get("family.ayah"))
}

func TestEnv(t *testing.T) {
	config := viper.New()
	//config.SetConfigName("config")
	//config.SetConfigType("json")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	//func read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.Get("APP_NAME"))
	assert.Equal(t, 5432, config.GetInt("PORT"))
	assert.Equal(t, true, config.GetBool("CONNECTION"))
	assert.Equal(t, "feri susmiyanto", config.Get("APP_AUTHOR"))
}

func TestAutomaticENV(t *testing.T) {
	config := viper.New()
	//config.SetConfigName("config")
	//config.SetConfigType("json")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	//func read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, 5432, config.GetInt("PORT"))
	assert.Equal(t, true, config.GetBool("CONNECTION"))
	assert.Equal(t, "feri susmiyanto", config.GetString("APP_AUTHOR"))

	assert.Equal(t, "Hello", config.Get("FROM_ENV"))
}
