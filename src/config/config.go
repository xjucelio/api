package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MySQL
	StringConexaoBanco = ""
	// Porta onde a API vai estar rodando
	Porta = 0
	// SecretKey e a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Load() {
	var erro error

	// Carrega as variáveis do arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", erro)
	}

	// Atribui a porta da API
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal("Erro ao converter a porta da API: ", erro)
	}

	// Constrói a string de conexão com o banco de dados
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	// Note: O host e a porta do MySQL no Docker podem ser diferentes.
	// O serviço 'db' do Docker Compose pode ser resolvido pelo nome do container.
	// Vamos supor que o host seja o próprio serviço 'db'.
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	// Formato da string de conexão:
	// <user>:<password>@tcp(<host>:<port>)/<dbname>?charset=utf8&parseTime=True&loc=Local
	StringConexaoBanco = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
