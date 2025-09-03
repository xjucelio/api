package modelos

// Senha representa o formato da apresentacao de alteracao de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
