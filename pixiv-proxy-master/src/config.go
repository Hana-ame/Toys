package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type config struct {
	Host      string            `json:"Host"`
	IPList    []string          `json:"IPList"`
	OverWrite map[string]string `json:"OverWrite"`
}

func initPars() {
	pairs = make(map[string]string, 200)
	pairs["https://cdnjs.cloudflare.com"] = "https://cdnjs.moonchan.xyz"
	pairs["https://i.pximg.net"] = "http://i.pximg.net"
	pairs[`https://account.pixiv.net`] = "http://account.pixiv.net"
	pairs["https://d.pixiv.org"] = "http://d.pixiv.org"
	pairs[`https:\/\/i.pximg.net`] = "http://i.pximg.net"
	pairs[`https:\/\/www.pixiv.net`] = "http://www.pixiv.net"
}
func initHost() {
	host = make(map[string]([]string), 200)
	host["exhentai.org"] = []string{
		"178.175.129.252",
		"178.175.129.254",
		"178.175.128.252",
		"178.175.128.254",
		"178.175.132.20",
		"178.175.132.22",
	}
	host["nyaa.si"] = []string{
		"185.178.208.182",
	}
	host["sukebei.nyaa.si"] = []string{
		"198.251.89.38",
	}
	host["www.pixiv.net"] = []string{
		"210.140.131.199",
		"210.140.131.219",
		"210.140.131.201",
		"210.140.131.223",
	}
	host["accounts.pixiv.net"] = []string{
		"210.140.131.199",
		"210.140.131.219",
		"210.140.131.201",
		"210.140.131.223",
	}
	host["i.pximg.net"] = []string{
		"210.140.92.141",
		"210.140.92.142",
		"210.140.92.140",
		"210.140.92.144",
		"210.140.92.145",
		"210.140.92.138",
		"210.140.92.146",
		"210.140.92.139",
		"210.140.92.143",
		"210.140.92.147",
	}
	host["d.pixiv.org"] = []string{
		"210.140.131.159",
		"210.140.131.157",
		"210.140.131.158",
	}
	host["www.xvideos.com"] = []string{
		"185.88.181.7",
		"185.88.181.8",
		"185.88.181.9",
		"185.88.181.10",
		"185.88.181.11",
		"185.88.181.2",
		"185.88.181.3",
		"185.88.181.4",
		"185.88.181.5",
		"185.88.181.6",
	}
	host["www.pornhub.com"] = []string{
		"66.254.114.41",
	}
	host["cn.pornhub.com"] = []string{
		"66.254.114.41",
	}
}

func loadConfig() {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// fmt.Println(string(byteValue))

	var cf []config
	json.Unmarshal((byteValue), &cf)

	// fmt.Println(cf)

	host = make(map[string]([]string), len(cf))
	pairsByHost = make(map[string](map[string]string), len(cf))

	for _, icf := range cf {
		host[icf.Host] = icf.IPList
		pairsByHost[icf.Host] = icf.OverWrite
	}

	// fmt.Println(host)
	// fmt.Println(pairsByHost)

}

func shouldProxy(s string) bool {
	return host[s] != nil
}

func getIPbyHost(s string) string {
	return host[s][rand.Intn(len(host[s]))]
}
