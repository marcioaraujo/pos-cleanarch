package configs

import (
	"github.com/spf13/viper"
)

// Conf armazena as configurações do sistema
type conf struct { // Deixe como está, não altere o nome
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RABBITMQ_USER     string `mapstructure:"RABBITMQ_USER"`
	RABBITMQ_PASS     string `mapstructure:"RABBITMQ_PASS"`
	RABBITMQ_IP       string `mapstructure:"RABBITMQ_IP"`
	RABBITMQ_PORT     string `mapstructure:"RABBITMQ_PORT"`
}

// LoadConfig carrega a configuração do arquivo .env usando o viper
func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config") // Define o nome do arquivo de configuração
	viper.SetConfigType("env")        // Define o tipo de arquivo como .env
	viper.AddConfigPath(path)         // Define o caminho para procurar o arquivo
	viper.SetConfigFile(".env")       // Define o arquivo específico
	viper.AutomaticEnv()              // Carrega as variáveis de ambiente automaticamente

	// Tenta ler a configuração
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Tenta deserializar os dados do arquivo para a estrutura conf
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
