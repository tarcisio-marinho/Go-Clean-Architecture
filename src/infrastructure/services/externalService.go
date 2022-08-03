package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IGetResource interface {
	Get(clientId string) (interface{}, error)
}

type GetResource struct {
}

func (query GetResource) Get(customerId string) (interface{}, error) {

	url := fmt.Sprintf("https://resourceUrl.com/%s", customerId)

	output, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	response, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	myOutputObj := struct {
		Name string
		Id   string
		Date string
	}{}

	err = json.Unmarshal(response, &myOutputObj)
	if err != nil {
		return nil, err
	}

	return output, nil
}
