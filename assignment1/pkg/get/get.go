package get

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Words struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percentages  map[string]float64 `json:"percentages"`
	Special      []string           `json:"special"`
	ExtraSpecial []interface{}      `json:"extraSpecial"`
}

type GetAPIIface interface {
	MakeGetRequest() Words
}

type GetApi struct {
	Client     http.Client
	RequestUrl string
}

func New(requestUrl string) GetAPIIface {
	return &GetApi{
		Client:     *http.DefaultClient,
		RequestUrl: requestUrl,
	}
}

func (a *GetApi) MakeGetRequest() Words {
	if _, err := url.ParseRequestURI(a.RequestUrl); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	response, err := a.Client.Get(a.RequestUrl)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var words Words
	if err = json.Unmarshal(body, &words); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return words
}
