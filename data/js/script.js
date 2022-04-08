
  document.getElementById("weatherpanel").style.visibility = "hidden";

  function choose() {
    fetch("/play")
      .then(response => response.json())
      .then(data => {
        document.getElementById("winning_numbers").innerHTML = "Winning Numbers are: " + data.winning_numbers;
        document.getElementById("winning_numbers").style.color = "green";
      })
  }

  function validate() {
    var ip = document.getElementById("ipaddress").value;

    if (ip == "") {
      alert("Ip cannot be empty!");
      return;
    }

    fetch("/verify?c=" + ip)
      .then(response => response.json())
      .then(data => {
        document.getElementById("validate_ip").innerHTML = data.ip_validator_result;
        if (data.ip_validator_result == "It is a Valid IP!") {
          document.getElementById("validate_ip").style.color = "green";
        } else {
          document.getElementById("validate_ip").style.color = "red";
        }
      })
  }


  function getWeather() {
    var location = document.getElementById("city").value;

    if (location == "") {
      alert("Location cannot be empty!");
      return;
    }

    fetch("/getweather?c=" + location)
      .then(response => response.json())
      .then(data => {

        //TEST
        // document.getElementById("weatherReport").innerHTML = data.weather_report;

        var obj = JSON.parse(data.weather_report);

        document.getElementById("locationName").innerHTML = obj.location.name;
        document.getElementById("time").innerHTML = obj.location.localtime;
        document.getElementById("cuurentTemp").innerHTML = obj.current.temp_c + "°C";
        document.getElementById("currentCondition").innerHTML = obj.current.condition.text;
        document.getElementById("currentImage").src = obj.current.condition.icon;
        document.getElementById("windSpeed").innerHTML = "Wind Speed: " + obj.current.wind_kph + " km/h";
        document.getElementById("feelsLike").innerHTML = "Feels like: " + obj.current.feelslike_c + "°C";
        document.getElementById("visibility").innerHTML = "Visibility: " + obj.current.vis_km + " km";
        document.getElementById("weatherpanel").style.visibility = "visible";


      })
  }

  function covert() {
    var from = document.getElementById("from").value;
    var to = document.getElementById("to").value;
    var amt = document.getElementById("amt").value;

    if (amt == "") {
      alert("Amount cannot be empty!");
      return;
    }

    fetch("/convertCurrency?from=" + from + "&to=" + to + "&have=" + amt)
      .then(response => response.json())
      .then(data => {

        var obj = JSON.parse(data.converted_currency);

        document.getElementById("conversion_result").innerHTML = obj.old_amount + " " + obj.old_currency + " equals " + obj.new_amount + " " + obj.new_currency + ".";

      })
  }

  function encode() {
    var value = document.getElementById("text_value").value;

    if (value == "") {
      alert("Text cannot be empty!");
      return;
    }

    fetch("encode?c=" + value)
      .then(response => response.json())
      .then(data => {
        document.getElementById("encode_decode_result").innerHTML = "Encrypted Value : " + data.encoded_value;
      })
  }

  function decode() {
    var value = document.getElementById("text_value").value;

    if (value == "") {
      alert("Text cannot be empty!");
      return;
    }

    fetch("decode?c=" + value)
      .then(response => response.json())
      .then(data => {
        document.getElementById("encode_decode_result").innerHTML = "Decrypted Value : " + data.decoded_value;
      })
  }
