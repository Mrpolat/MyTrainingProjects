package main

import (
	"strings"
	"net/http"
	"fmt"
    "io/ioutil"
    "regexp"
	"os"
    "io"
)

func main() {
    //File(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")),"test")
    
    //fmt.Println(between(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment"), `pdp-product-info__id`, `a`))

    r,_ := regexp.Compile(`pdp-product-info__id"><span>(.*)</span></div><div class="green`)
	strs := r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("OrderNumberForMatching:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`<div class="lockUpperCase">(.*)</div><td></tr><tr><th class="pes-text-left char-table__title">product or`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("Family:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`product or component type</th><td class="pes-text-left"><div class="lockUpperCase">(.*)</div><td></tr><tr><th class="pes-text-left char-table__title">number of racks`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("AssetType:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`number of slots</th><td class="pes-text-left"><div class="lockUpperCase">(.*)</div><td></tr><tr><th class="pes-text-left char-table__title">discrete I/O processor capacity`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("NumberOfSlots:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`discrete I/O processor capacity</th><td class="pes-text-left"><div class="lockUpperCase">(.*)</div><td></tr><tr><th class="pes-text-left char-table__title">analogue I/O processor capacity`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("Description:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`analogue I/O processor capacity</th><td class="pes-text-left"><div class="lockUpperCase">(.*)</div><td></tr><tr><th class="pes-text-left char-table__title">number of application specific channel`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("Description1:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`integrated connection type</th><td class="pes-text-left"><div class="lockUpperCase">(.*)</div><td></tr><tr><th class="pes-text-left char-table__title">communication module processor capacity`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("Interface:",strs[1])
    println("-----------------------")
    r,_ = regexp.Compile(`standards</th><td class="pes-text-left"><div class="lockUpperCase">(.*)</div><td></tr></tbody></table><table class="pes-table"><caption class="pes-text-left">Offer Sustainability`)
	strs = r.FindStringSubmatch(HtmlMinify(getHTML("https://www.schneider-electric.com/en/product/BMXP341000/processor-module-m340---max-512-discrete-%2B-128-analog-i-o---modbus/?range=1468-modicon-m340&node=166345095-standard-environment")))
    fmt.Println("Standards:",strs[1])
}

func getHTML(url string) string {

	resp, _ := http.Get(url) //http kütüphanesinin get metodu kullanılırak girilen url okunuyor 
	bytes, _ := ioutil.ReadAll(resp.Body)// response değerinin body metodu kullanılarak ioutil kütüphanesi ile byte olarak okunuyor 

	resp.Body.Close()

    return string(bytes)

}

func File(content string, filename string){
    fmt.Print("************\n")
  
    fmt.Print("************\n")

    file, err := os.Create(filename)
    // Bir hata oluşursa err dolu gelecek, onu checkError() metoduna gönder.
    if err != nil {
        panic(err)
    }
    // Dosyayı serbest bırak.
    defer file.Close()
 
    // Oluşturulan dosyaya kullanıcıdan alınan metni yaz.
    _, err = io.WriteString(file, content)
    if err != nil {
        panic(err)
    }
}

func HtmlMinify(html string) string {
    htmlArray := strings.Fields(html)
    htmlJoined := strings.Join(htmlArray, " ")
    
    return strings.Replace(htmlJoined, "> <", "><", -1)
}
