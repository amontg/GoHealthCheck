package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configuration struct {
	User     string
	Pass     string
	Name     string
	Surname  string
	OnCampus int
	Vaxxed   bool
	Sick     int
	Exposed  int
	TestPos  int
}

var config *configuration

// Load env variables
func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// because I need to convert some shit and they also return errors
	OnCampus, _ := strconv.Atoi(os.Getenv("ONCAMPUS"))
	Vaxxed, _ := strconv.ParseBool(os.Getenv("VAXXED"))
	Sick, _ := strconv.Atoi(os.Getenv("SICK"))
	Exposed, _ := strconv.Atoi(os.Getenv("EXPOSED"))
	TestPos, _ := strconv.Atoi(os.Getenv("TESTPOS"))

	config = &configuration{
		User:     os.Getenv("CLU_USER"),
		Pass:     os.Getenv("CLU_PASS"),
		Name:     os.Getenv("FNAME"),
		Surname:  os.Getenv("LNAME"),
		OnCampus: OnCampus,
		Vaxxed:   Vaxxed,
		Sick:     Sick,
		Exposed:  Exposed,
		TestPos:  TestPos,
	}
}

// returns (User string, Pass string, Name string, Surname string)
func GetPersonals() (string, string, string, string) {
	return config.User, config.Pass, config.Name, config.Surname
}

// returns (OnCampus int, Vax bool, Sick int, Exposed int, Test int)
func GetFields() (int, bool, int, int, int) {
	return config.OnCampus, config.Vaxxed, config.Sick, config.Exposed, config.TestPos
}
