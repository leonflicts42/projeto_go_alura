package database

import (
	"log"
	"os"
	"time"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")

	stringDeConexao := "host=" + endereco +
		" user=" + usuario +
		" password=" + senha +
		" dbname=" + nomeBanco +
		" port=" + portaBanco +
		" sslmode=disable"

	// 🔁 tenta conectar várias vezes
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(stringDeConexao))

		if err == nil {
			log.Println("✅ Conectado ao banco de dados!")
			break
		}

		log.Println("⏳ Banco ainda não está pronto, tentando novamente...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Panic("❌ Não foi possível conectar ao banco de dados")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}