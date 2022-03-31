package util

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strconv"
)

type Round struct {
	WinningNumbers    string `json:"winning_numbers"`
	IpValidatorResult string `json:"ip_validator_result"`
	WeatherReport     string `json:"weather_report"`
	ConvertedCurrency string `json:"converted_currency"`
	MessageStatus     string `json:"message_status"`
	EncodedValue      string `json:"encoded_value"`
	DecodedValue      string `json:"decoded_value"`
}

func PlayLottery() Round {
	var nums []int
	var min int = 0
	var max int = 50
	var i int = 0
	winningNumbers := ""
	var randomNumber int = 0

	for i = 0; i < 8; i++ {
		randomNumber = (rand.Intn(max-min) + min)
		nums = append(nums, randomNumber)
	}

	winningNumbers = strconv.Itoa(nums[0])

	for i = 1; i < 8; i++ {
		winningNumbers = winningNumbers + ", " + strconv.Itoa(nums[i])
	}

	var result Round
	result.WinningNumbers = winningNumbers
	return result

}

func ValidateIp(ip string) Round {
	ipValidatorResult := ""

	//--------------JUST FOUND OUT THIS LIBRARY IN GO THAT AUTOMATICALLY CHECKS IP WITHOUT USING REGEX!-------------

	if net.ParseIP(ip) != nil {
		ipValidatorResult = "It is a Valid IP!"
	} else {
		ipValidatorResult = "It is an Invalid IP!"
	}

	// obj, err := regexp.Match("(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])", []byte(ip))

	var result Round
	result.IpValidatorResult = ipValidatorResult
	return result
}

func GetWeather(city string) Round {

	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + city

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "weatherapi-com.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "50816c55bemsha74515a42a5acf6p160902jsn8da02f65d690")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Round
	result.WeatherReport = (string(body))
	return result

}

func ConvertCurrency(from string, to string, have string) Round {
	url := "https://currency-converter-by-api-ninjas.p.rapidapi.com/v1/convertcurrency?have=" + from + "&want=" + to + "&amount=" + have

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "currency-converter-by-api-ninjas.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "50816c55bemsha74515a42a5acf6p160902jsn8da02f65d690")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Round
	result.ConvertedCurrency = (string(body))
	return result
}

func Encode(s string) Round {
	se := base64.StdEncoding.EncodeToString([]byte(s))

	var result Round
	result.EncodedValue = se
	return result
}

func Decode(s string) Round {
	sd, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		fmt.Println(e)
	}
	var result Round
	result.DecodedValue = string(sd)
	return result
}
