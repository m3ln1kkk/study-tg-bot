package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type dataCoinPrice struct {
	Bitcoin  string
	Ethereum string
}

func general(w http.ResponseWriter, r *http.Request) {
	dataCoinPrice := dataCoinPrice{
		Bitcoin:  getPrice("bitcoin", "usd"),
		Ethereum: getPrice("ethereum", "usd"),
	}
	template, _ := template.ParseFiles("general.html")
	template.Execute(w, dataCoinPrice)
}

func getPrice(coin string, cur string) string {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + coin + "&vs_currencies=" + cur

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-pro-api-key", "CG-2qGpqEPGVxfrCu5PCSTKj1Rk")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	textRes, _ := io.ReadAll(res.Body)
	return string(textRes)
}

func main() {

	http.HandleFunc("/", general)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
