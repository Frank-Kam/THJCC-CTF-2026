package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.New("slot").Parse(`
<!DOCTYPE html>
<html lang="zh-TW">
<head>
<meta charset="UTF-8">
<title>Las Vegas</title>
<style>
body {
    background: #111;
    color: #fff;
    font-family: monospace;
    text-align: center;
    padding-top: 50px;
}
#slot {
    font-size: 4em;
    letter-spacing: 20px;
}
button {
    font-size: 1.5em;
    margin-top: 20px;
}
#message {
    margin-top: 20px;
    font-size: 1.5em;
}
</style>
</head>
<body>
<h1>Slot Machine 拉霸機 🎰</h1>
<div id="slot">0 0 0</div>
<button id="spin">Pull!</button>
<p id="message">{{.Message}}</p>

<script>
const slot = document.getElementById("slot");
const btn = document.getElementById("spin");
const message = document.getElementById("message");

let interval;

function getRandomDigit() {
    return Math.floor(Math.random() * 10);
}

btn.onclick = function() {
    btn.disabled = true;
    let digits = [0,0,0];
    let count = 0;

    interval = setInterval(() => {
        for (let i = 0; i < 3; i++) {
            digits[i] = getRandomDigit();
        }
        slot.textContent = digits.join(' ');
        count++;

        if(count > 20){
            clearInterval(interval);
            const n = digits.join('');
            fetch("/?n=" + n, {method: "POST"})
                .then(resp => resp.text())
                .then(txt => {
                    message.innerHTML = txt;
                    btn.disabled = false;
                });
        }
    }, 100);
};
</script>
</body>
</html>
`))

func generateFlag() string {
	b := make([]byte, 8) // 8 bytes = 16 hex
	_, err := rand.Read(b)
	if err != nil {
		return "THJCC{LUcKy_sEVen_deadbeefdeadbeef}"
	}
	return fmt.Sprintf("THJCC{LUcKy_sEVen_%s}", hex.EncodeToString(b))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num := r.URL.Query().Get("n")

		if num == "777" {
			flag := generateFlag()
			fmt.Fprintf(w, "What a Lucky man! %s", flag)
		} else {
			fmt.Fprint(w, "Nope, Try Again!")
		}
		return
	}

	tpl.Execute(w, struct{ Message string }{Message: ""})
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("run at :14514")
	log.Fatal(http.ListenAndServe(":14514", nil))
}
