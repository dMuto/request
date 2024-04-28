package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("O status da requisição foi: ", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp.Body.Close()

	fmt.Println("O corpo da resposta foi: ", string(body))

	callDistopiaExecOne()
}

func callDistopiaExecOne(){
	resp, err := http.Get("https://distopia.savi2w.workers.dev/")
	
	if err != nil{
		fmt.Println(err.Error())
		return
	}
    
	
	fmt.Println("Header escondido: ", resp.Request.Response.Header.Get("Distopia"))
}
