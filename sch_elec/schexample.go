package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
    "io"
    "astuart.co/goq"
    "log"
    
)

type Product struct {
    OrderNumberForMatching string `goquery:"#start > div.pdp-product-info > div.pdp-product-wrapper--desktop > div.pdp-product-info__id > span"`
    Family string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(1) > td.pes-text-left > div"`
    AssetType string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    NumberOfSlots string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(4) > td.pes-text-left > div"`    
    Description string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(5) > td.pes-text-left > div"`
    Description1 string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(6) > td.pes-text-left > div"`
    Interface string `goquery:"#characteristics > div > table:nth-child(2) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    Standards string `goquery:"#characteristics > div > table:nth-child(3) > tbody > tr:nth-child(6) > td.pes-text-left > div"` 
}

func main() {
    File()
    Goq()
}

func getHTML(url string) string {

	resp, _ := http.Get(url) //http kütüphanesinin get metodu kullanılırak girilen url okunuyor 
	bytes, _ := ioutil.ReadAll(resp.Body)// response değerinin body metodu kullanılarak ioutil kütüphanesi ile byte olarak okunuyor 

	resp.Body.Close()

    return string(bytes)

}

func File(){
 
    
    content:= getHTML("https://www.se.com/tr/tr/product/BMXP341000/i%C5%9Flemci-mod%C3%BCl%C3%BC-m340---maks-512-dijital-%2B-128-analog-g-%C3%A7---modbus/?range=1468-modicon-m340&node=12145172188-standart-ortam")
    fmt.Print("************\n")
    fmt.Print("Metin: ")
    fmt.Print("************\n")

  
    // Dosyaya verilecek isim için adını ve uzantısını belirle.
    ext := ".txt"
    fileName := "./dummy-" +  ext
    // Dosyayı oluştur.
    file, err := os.Create(fileName)
    // Bir hata oluşursa err dolu gelecek, onu checkError() metoduna gönder.
    checkError(err)
    // Dosyayı serbest bırak.
    defer file.Close()
 
    // Oluşturulan dosyaya kullanıcıdan alınan metni yaz.
    ln, err := io.WriteString(file, content)
    checkError(err)
 
    fmt.Printf("Dosyada %v karakter var.", ln)// işin süs kısmı
   
}

func checkError(err error){
    if err != nil {
        panic(err)
    }

}

func Goq() {
	res, err := http.Get("https://www.se.com/tr/tr/product/BMXP341000/i%C5%9Flemci-mod%C3%BCl%C3%BC-m340---maks-512-dijital-%2B-128-analog-g-%C3%A7---modbus/?range=1468-modicon-m340&node=12145172188-standart-ortam")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var ex Product
	
	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}
 
    log.Println(ex.OrderNumberForMatching, "\n" ,ex.Family,"\n" ,ex.AssetType,"\n" ,ex.NumberOfSlots,"\n", ex.Description, "\n", ex.Description1, "\n", ex.Interface , "\n", ex.Standards)
    
}