package apiserver

type APIServer struct {
	cfg *Config

}

func New(cfg *Config) *APIServer {
	return &APIServer{
		cfg: cfg,
	}
}

func (s *APIServer) Start() error {
	return nil
}