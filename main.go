package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const MONITORAMENTOS = 3
const DELAY = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) // sair com sucesso
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	fmt.Println("*--------------- MONITORAMENTO ---------------*")
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido:", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{
	// 	"https://www.alura.com.br",
	// 	"https://www.youtube.com",
	// 	"https://github.com/",
	// }

	sites := leSitesDoArquivo()

	for i := 0; i < MONITORAMENTOS; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "está online")
	} else {
		fmt.Println("O site", site, "retornou o status: ", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	//arquivo, err := os.Open("sitess.txt") devolve o ponteiro pro arquivo
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de sites: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}
