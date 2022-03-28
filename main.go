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

// func ipPage(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "ipValidator.html")
// }

func playLottery(w http.ResponseWriter, r *http.Request) {
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
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
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
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
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
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
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
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

func sendSms(w http.ResponseWriter, r *http.Request) {
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
	number := (r.URL.Query().Get("number"))
	message := (r.URL.Query().Get("message"))
	result := util.SendSms(string(number), string(message))

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
	http.HandleFunc("/sendSms", sendSms)
	http.HandleFunc("/", homePage)

	fs := http.FileServer(http.Dir("../myApp"))
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", fs)
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
