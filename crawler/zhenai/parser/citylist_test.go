package parser

import (
	"testing"

	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html.txt")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")

	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}
}
