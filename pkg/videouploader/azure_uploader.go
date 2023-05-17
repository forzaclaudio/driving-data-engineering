import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//uploaderPort := 8080
//pingUploader("192.168.1.2", uploaderPort)

func pingUploader(host string, port int) {
	url := fmt.Sprintf("http://%s:%d/version", host, port)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func requestFileUpload(host string, port int, filePath string) {
	url := fmt.Sprintf("http://%s:%d/upload", host, port)
	requestBody, err := json.Marshal(map[string]string{"win_filepath": filePath})
	if err != nil {
		log.Println(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))
}