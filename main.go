package main

import (
	"encoding/json"

	"io"
	"log"
	"net/http"
	"os"
	"time"

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

var Client *http.Client

func init() {
	Client = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func display(w http.ResponseWriter, r *http.Request) {
	getConversionCurrency(Bitcoin, USD)
	getConversionCurrency(Ethereum, EUR)
	getConversionCurrency(Tron, EUR)

	jsonData, err := json.Marshal(DataCur)
	checkErr("ошибка при конвертации json", err)
	w.Write(jsonData)
}

func getConversionCurrency(inCoin cryptoCurrency, outCoin fiatCurrency) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + string(inCoin) + "&vs_currencies=" + string(outCoin)

	req, err := http.NewRequest("GET", url, nil)
	checkErr("ошибка при создании запроса", err)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-pro-api-key", string(os.Getenv("xCgProApiKey")))

	res, err := Client.Do(req)
	checkErr("ошибка при выполнении запроса", err)
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
