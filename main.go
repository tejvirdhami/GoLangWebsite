package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/util"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playLottery(w http.ResponseWriter, r *http.Request) {
	result := util.PlayLottery()

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func validateIp(w http.ResponseWriter, r *http.Request) {
	ip := (r.URL.Query().Get("c"))
	result := util.ValidateIp(string(ip))

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := (r.URL.Query().Get("c"))
	result := util.GetWeather(string(city))

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func convertCurrency(w http.ResponseWriter, r *http.Request) {
	from := (r.URL.Query().Get("from"))
	to := (r.URL.Query().Get("to"))
	have := (r.URL.Query().Get("have"))
	result := util.ConvertCurrency(string(from), string(to), string(have))

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func encode(w http.ResponseWriter, r *http.Request) {
	value := (r.URL.Query().Get("c"))
	result := util.Encode(string(value))

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func decode(w http.ResponseWriter, r *http.Request) {
	value := (r.URL.Query().Get("c"))
	result := util.Decode(string(value))

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/play", playLottery)
	http.HandleFunc("/verify", validateIp)
	http.HandleFunc("/getweather", getWeather)
	http.HandleFunc("/convertCurrency", convertCurrency)
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/", homePage)

	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))

	// fs := http.FileServer(http.Dir("../myApp"))
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
