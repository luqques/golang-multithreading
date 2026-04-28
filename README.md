# Desafio de multithreading do curso Go Expert da Full Cycle

## Como rodar o programa

Como rodar o programa:
1. Clonar o repositório na sua máquina
2. Acessar a pasta raiz do projeto
3. Rodar o comando `go run main.go`
4. Digite um CEP quando for solicitado e aperte a tecla Enter para iniciar a requisição

## Descrição do desafio

Neste desafio, você deve aplicar conhecimentos de Multithreading e APIs para construir um sistema que busca o endereço mais rápido entre duas APIs distintas. O foco é lidar com requisições simultâneas, tratamento de concorrência e timeouts.

APIs Alvo Você deve consumir as seguintes APIs de consulta de CEP:

BrasilAPI: https://brasilapi.com.br/api/cep/v1/{cep}
ViaCEP: http://viacep.com.br/ws/{cep}/json/
Requisitos Técnicos

1. Requisições Simultâneas: O seu sistema deve fazer a requisição para as duas APIs ao mesmo tempo (paralelamente).

2. Race Condition (Corrida): O sistema deve aceitar apenas a resposta da API que responder mais rápido e descartar a resposta da outra (mais lenta).

3. Output (Saída): O resultado da requisição deve ser exibido na linha de comando (terminal), contendo:
Os dados do endereço recebido.
Qual API entregou a resposta (BrasilAPI ou ViaCEP).

4. Timeout: O tempo limite de resposta é de 1 segundo.

Caso nenhuma das APIs responda dentro desse tempo, o sistema deve exibir um erro de timeout.

Tecnologias

Linguagem: Go (Golang)

Conceitos: Goroutines, Channels, Select, Package net/http.
 

Entregável

Link do Repositório: O link para o seu repositório no GitHub contendo o código fonte.

README: Instruções de como rodar o projeto.
 

Regras de Entrega

Repositório Único: É obrigatória a entrega de apenas um projeto por repositório.

Branch Principal: Todo o código deve estar na branch main.

Documentação: O README deve explicar claramente como executar e testar sua aplicação.

