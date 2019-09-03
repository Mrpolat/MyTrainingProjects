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

type Data struct {
    OrderNumberForMatching string `goquery:"#start > div.pdp-product-info > div.pdp-product-wrapper--desktop > div.pdp-product-info__id > span"`
    Family string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(1) > td.pes-text-left > div"`
    AssetType string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    NumberOfSlots string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(4) > td.pes-text-left > div"`    
    Description string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(5) > td.pes-text-left > div"`
    Description1 string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(6) > td.pes-text-left > div"`
    Interface string `goquery:"#characteristics > div > table:nth-child(2) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    Standards string `goquery:"#characteristics > div > table:nth-child(3) > tbody > tr:nth-child(6) > td.pes-text-left > div"` 
}

type Data1 struct {
    OrderNumberForMatching string `goquery:"#start > div.pdp-product-info > div.pdp-product-wrapper--desktop > div.pdp-product-info__id > span"`
    Family string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(1) > td.pes-text-left > div"`
    AssetType string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    NumberOfSlots string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(4) > td.pes-text-left > div"`    
    Description string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(5) > td.pes-text-left > div"`
    Description1 string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(6) > td.pes-text-left > div"`
    Interface string `goquery:"#characteristics > div > table:nth-child(2) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    Standards string `goquery:"#characteristics > div > table:nth-child(3) > tbody > tr:nth-child(6) > td.pes-text-left > div"` 
}

type Data2 struct {
    OrderNumberForMatching string `goquery:"#start > div.pdp-product-info > div.pdp-product-wrapper--desktop > div.pdp-product-info__id > span"`
    Family string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(1) > td.pes-text-left > div"`
    AssetType string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    NumberOfSlots string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(4) > td.pes-text-left > div"`    
    Description string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(5) > td.pes-text-left > div"`
    Description1 string `goquery:"#characteristics > div > table:nth-child(1) > tbody > tr:nth-child(6) > td.pes-text-left > div"`
    Interface string `goquery:"#characteristics > div > table:nth-child(2) > tbody > tr:nth-child(2) > td.pes-text-left > div"`
    Standards string `goquery:"#characteristics > div > table:nth-child(3) > tbody > tr:nth-child(6) > td.pes-text-left > div"` 
}


type Data3 struct {
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
    data()
    dataone()
    datatwo()
    datathree()

    // struct fieldlarını okudun
    // struct fieldları içinde donerek bir değer oku
    // field baslangıcı slot olsun diyelim
    // slot........onemli bilgi > onemli bilgiyi slot değişkenine atıp kullanabilirsin

    // for  x in data() {
    //     if string.contains(x,"slot")
    //     if string.contains(x."ürün serisi" " ürün ya da ")
    //         slot = x[20][30]
    // }

}

func getHTML(url string) string {

	resp, _ := http.Get(url) //http kütüphanesinin get metodu kullanılırak girilen url okunuyor 
	bytes, _ := ioutil.ReadAll(resp.Body)// response değerinin body metodu kullanılarak ioutil kütüphanesi ile byte olarak okunuyor 

	resp.Body.Close()

    return string(bytes)

}

func File(){
 
    
    content:= getHTML("https://www.se.com/tr/tr/product-range/1468-modicon-m340/?filter=business-1-endüstriyel-otomasyon-ve-kontrol")
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
 
    log.Println("OrderNumberForMatching:", (ex.OrderNumberForMatching))
    log.Println("Family:", (ex.Family))
    log.Println("AssetType:", (ex.AssetType))
    log.Println("NumberOfSlots:", (ex.NumberOfSlots))
    log.Println("Description:", (ex.Description))
    log.Println("Description1:", (ex.Description1))
    log.Println("Interface:", (ex.Interface))
    log.Println("Standards:", (ex.Standards))
    println("-----------------------")
   
    
}

func data() {
	res, err := http.Get("https://www.se.com/tr/tr/product/BMXP342000/i%C5%9Flemci-mod%C3%BCl%C3%BC-m340---maks-1024-dijital-%2B-256-analog-g-%C3%A7---modbus/?range=1468-modicon-m340&node=12145172188-standart-ortam")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var ex Data
	
	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}
    log.Println("OrderNumberForMatching:", (ex.OrderNumberForMatching))
    log.Println("Family:", (ex.Family))
    log.Println("AssetType:", (ex.AssetType))
    log.Println("NumberOfSlots:", (ex.NumberOfSlots))
    log.Println("Description:", (ex.Description))
    log.Println("Description1:", (ex.Description1))
    log.Println("Interface:", (ex.Interface))
    log.Println("Standards:", (ex.Standards))
    println("-----------------------")
   
    
}

func dataone() {
	res, err := http.Get("https://www.se.com/tr/tr/product/BMXP3420102/i%C5%9Flemci-mod%C3%BCl%C3%BC-m340---maks-1024-dijital-%2B-256-analog-g-%C3%A7/?range=1468-modicon-m340&node=12145172188-standart-ortam")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var ex Data1
	
	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}
 
    log.Println("OrderNumberForMatching:", (ex.OrderNumberForMatching))
    log.Println("Family:", (ex.Family))
    log.Println("AssetType:", (ex.AssetType))
    log.Println("NumberOfSlots:", (ex.NumberOfSlots))
    log.Println("Description:", (ex.Description))
    log.Println("Description1:", (ex.Description1))
    log.Println("Interface:", (ex.Interface))
    log.Println("Standards:", (ex.Standards))
    println("-----------------------")
   
    
}

func datatwo() {
	res, err := http.Get("https://www.se.com/tr/tr/product/BMXP342020/i%C5%9Flemci-mod%C3%BCl%C3%BC-m340---maks-1024-dijital-%2B-256-analog-g-%C3%A7---modbus---ethernet/?range=1468-modicon-m340&node=12145172188-standart-ortam")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var ex Data2
	
	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}
 
    log.Println("OrderNumberForMatching:", (ex.OrderNumberForMatching))
    log.Println("Family:", (ex.Family))
    log.Println("AssetType:", (ex.AssetType))
    log.Println("NumberOfSlots:", (ex.NumberOfSlots))
    log.Println("Description:", (ex.Description))
    log.Println("Description1:", (ex.Description1))
    log.Println("Interface:", (ex.Interface))
    log.Println("Standards:", (ex.Standards))
    println("-----------------------")
   
    
}

func datathree() {
	res, err := http.Get("https://www.se.com/tr/tr/product/BMXP3420302/i%C5%9Flemci-mod%C3%BCl%C3%BC-m340---maks-1024-dijital-%2B-256-analog-g-%C3%A7---canopen/?range=1468-modicon-m340&node=12145172188-standart-ortam")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var ex Data3
	
	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}
 
    log.Println("OrderNumberForMatching:", (ex.OrderNumberForMatching))
    log.Println("Family:", (ex.Family))
    log.Println("AssetType:", (ex.AssetType))
    log.Println("NumberOfSlots:", (ex.NumberOfSlots))
    log.Println("Description:", (ex.Description))
    log.Println("Description1:", (ex.Description1))
    log.Println("Interface:", (ex.Interface))
    log.Println("Standards:", (ex.Standards))
    println("-----------------------")
   
    
}