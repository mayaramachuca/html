package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// func passando uma lista(...) de urls de string() - retornando um canal somente leitura(<- ch) com dados do tipo string
// que são os titulos das urls para isso vai ser feito uma chamada http do tipo GET em cima das URLS
// Titulo obtem titulo de uma pagina html
func Titulo(urls ...string) <-chan string {

	c := make(chan string)
	//criar um laço para cada url recebida no parametro na lista urls ...string
	for _, url := range urls {
		//criando uma func anonima (sem nome)
		//goroutine(go antes de func())
		go func(url string) {
			//ignorando o erro
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) //invocar() função por ela ser anonima, passando a url
	}
	return c
}
