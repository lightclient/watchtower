package root

// type MongoConfig struct {
// 	Ip     string `json:"ip"`
// 	DbName string `json:"dbName"`
// }

type ServerConfig struct {
	Port string `json:"port"`
}

type KubernetesConfig struct {
	Namespace string `json:"namespace"`
}

// type AuthConfig struct {
// 	Secret string `json:"secret"`
// }

type Config struct {
	Server     *ServerConfig     `json:"server"`
	Kubernetes *KubernetesConfig `json:"kubernetes"`
}
