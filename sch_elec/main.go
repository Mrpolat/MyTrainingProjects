package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"io"
"bufio"
	
    
	
	
	
)

func main() {
	fmt.Println(getHTML("https://www.schneider-electric.com/ww/en/")) // byte tipi stringe dönüştürülüyor elde ettiğimiz değer girelen url nin html kodudur
    File()
}

func getHTML(url string) string {

	resp, _ := http.Get(url) //http kütüphanesinin get metodu kullanılırak girilen url okunuyor 
	bytes, _ := ioutil.ReadAll(resp.Body)// response değerinin body metodu kullanılarak ioutil kütüphanesi ile byte olarak okunuyor 

	resp.Body.Close()

    return string(bytes)

}

func File(){
 
    // Biz klavyeden girilen karakterleri content olarak alacağız.
    // Bu nedenle bu kısmı açıklama satırı olarak bıraktım.
    //content := 
 
    // Klavyeden girilen veriyi elde edebilmek için projeye eklediğimiz bufio kütüphanesini kullanıyoruz.
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("************\n")
    fmt.Print("Metin: ")
    content, _ := reader.ReadString('\n')
    fmt.Print("************\n")


    // Dosyaya verilecek isim için adını ve uzantısını belirle.
    ext := ".txt"
    fileName := "./dummy-" + ext
 
    // Dosyayı oluştur.
    file, err := os.Create(fileName)
    // Bir hata oluşursa err dolu gelecek, onu checkError() metoduna gönder.
    checkError(err)
    // Dosyayı serbest bırak.
    defer file.Close()
 
    // Oluşturulan dosyaya kullanıcıdan alınan metni yaz.
    ln, err := io.WriteString(file, content)
    checkError(err)
 
    fmt.Printf("Dosyada %v karakter var.", ln)
 
   
}



func checkError(err error){
    if err != nil {
        panic(err)
    }

}
