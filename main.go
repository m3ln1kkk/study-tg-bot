package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

// Обработать и сделать вывод, типизировать ошибки

// что делает?
type dataCoinPrice struct {
	Bitcoin  string
	Ethereum string
	BNB      string
	Tron     string
}

// нейминг функций??
func general(w http.ResponseWriter, r *http.Request) {
	// нейминг??
	dataCoinPrice := dataCoinPrice{
		Bitcoin:  getPrice("bitcoin", "usd"),
		Ethereum: getPrice("ethereum", "usd"),
	}
	// обработать
	// нейминг как из пакета ??
	template, _ := template.ParseFiles("general.html")

	// обработать
	template.Execute(w, dataCoinPrice)
}

// Что за нейминг переменных? | А какие валюты доступны?? | enum -> посмотреть что это такое как это реализовать в go
func getPrice(coin string, cur string) string {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + coin + "&vs_currencies=" + cur

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")                      // вынести в отдельный клиент
	req.Header.Add("x-cg-pro-api-key", "CG-2qGpqEPGVxfrCu5PCSTKj1Rk") // вынести в отдельный клиент

	// обработать
	res, _ := http.DefaultClient.Do(req) // перепиши на go доку
	defer res.Body.Close()

	// обработать
	textRes, _ := io.ReadAll(res.Body)
	return string(textRes)
}

// Структура, которая будет сериализована в JSON
//type Response struct {
//	Message string `json:"message"`
//	Status  int    `json:"status"`
//}
//
//func jsonHandler(w http.ResponseWriter, r *http.Request) {
//	// Проверяем, что это метод GET
//	if r.Method != http.MethodGet {
//		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
//		return
//	}
//
//	// Создаем объект ответа
//	response := Response{
//		Message: "Hello, this is a JSON response!",
//		Status:  http.StatusOK,
//	}
//
//	// Устанавливаем заголовок Content-Type
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//
//	// Кодируем объект в JSON и отправляем клиенту
//	err := json.NewEncoder(w).Encode(response)
//	if err != nil {
//		http.Error(w, "Ошибка при создании JSON", http.StatusInternalServerError)
//	}
//}

func main() {

	http.HandleFunc("/", general)
	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	log.Printf("%s %s", r.Method, r.URL)
	//	w.Write([]byte("Hello from /hello route!")) // Добавляем ответ для маршрута /hello
	//})
	//http.HandleFunc("/json", jsonHandler)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
