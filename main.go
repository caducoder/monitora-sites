package main

import (
	"fmt"
	"net/http"
	"os"
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
	sites := []string{
		"https://www.alura.com.br",
		"https://www.youtube.com",
		"https://github.com/",
	}

	for i := 0; i < MONITORAMENTOS; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "está online")
	} else {
		fmt.Println("O site", site, "retornou o status: ", resp.StatusCode)
	}
}
