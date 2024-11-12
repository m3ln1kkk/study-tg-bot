package main

import (
	"bytes"
	"encoding/json"

	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type cryptoCurrency string
type fiatCurrency string

const (
	Bitcoin  cryptoCurrency = "bitcoin"
	Ethereum cryptoCurrency = "ethereum"
	Tron     cryptoCurrency = "tron"
)

const (
	USD fiatCurrency = "usd"
	EUR fiatCurrency = "eur"
)

type listCrypto struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
		EUR float64 `json:"eur"`
	} `json:"bitcoin"`
	Ethereum struct {
		USD float64 `json:"usd"`
		EUR float64 `json:"eur"`
	} `json:"ethereum"`
	Tron struct {
		USD float64 `json:"usd"`
		EUR float64 `json:"eur"`
	} `json:"tron"`
}

var DataCur listCrypto //variable to store the currency data

func checkErr(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

var client *Client

func newClient(baseURL string, headers map[string]string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		Headers:    headers,
	}
}

func (c *Client) DoRequest(method, route string, body []byte) (*http.Response, error) {
	url := c.BaseURL + route
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	checkErr("ошибка при создании запроса ", err)

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return c.HTTPClient.Do(req)
}

func init() {
	headers := map[string]string{
		"x-cg-pro-api-key": string(os.Getenv("xCgProApiKey")),
		"accept":           "application/json",
	}
	client = newClient("https://api.coingecko.com/api/v3", headers)
}

func display(w http.ResponseWriter, r *http.Request) {
	getConversionCurrency(Bitcoin, USD)
	getConversionCurrency(Ethereum, EUR)
	getConversionCurrency(Tron, EUR)
	getConversionCurrency("tronn", EUR)

	jsonData, err := json.Marshal(DataCur)
	checkErr("ошибка при конвертации json", err)
	w.Write(jsonData)
}

func getConversionCurrency(inCoin cryptoCurrency, outCoin fiatCurrency) {
	route := "/simple/price?" + "ids=" + string(inCoin) + "&vs_currencies=" + string(outCoin)

	res, err := client.DoRequest("GET", route, nil)
	checkErr("ошибка при выполненииtest запроса", err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	checkErr("ошибка при чтении ответа", err)

	err = json.Unmarshal(body, &DataCur)
	checkErr("ошибка при парсинге JSON", err)
}

func main() {

	err := godotenv.Load() //reading env
	checkErr("ошибка при прочтении переменных окружения", err)

	http.HandleFunc("/", display)
	log.Fatal(http.ListenAndServe(":8092", nil))

}
