package dadJoke

import (
	"encoding/json"
	f "fmt"
	"io/ioutil"
	"net/http"
)

var URL string = "https://icanhazdadjoke.com/"

type Joke struct {
	Joke string `json:"joke"`
	ID   string `json:"id"`
}

func errCheck(e error) {
	if e != nil {
		f.Errorf("error: %g", e)
	}
}

func ApiCall() []byte {
	req, err := http.NewRequest("GET", URL, nil)
	errCheck(err)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	errCheck(err)

	body, err := ioutil.ReadAll(resp.Body)
	errCheck(err)

	defer resp.Body.Close()

	return body
}

func DecodeJson(payload []byte) (string, string) {
	var jk Joke
	err := json.Unmarshal(payload, &jk)
	errCheck(err)

	return jk.Joke, jk.ID
}

func displayJoke(text, id string) {
  f.Printf("Joke:\t%s\n\tID:\t%s", id, text)
}

func GetJoke() {
	rawResp := ApiCall()
	joke, id := DecodeJson(rawResp)
	displayJoke(joke, id)
}
