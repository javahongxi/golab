package parser

import (
	_ "embed"
)

//go:embed profile_test_data.html.txt
var profileTestData string

//go:embed citylist_test_data.html.txt
var citylistTestData string

func GetProfileTestData() string {
	return profileTestData
}

func GetCitylistTestData() string {
	return citylistTestData
}
