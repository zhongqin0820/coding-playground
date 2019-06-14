package method

// abstrct product
type Server interface {
	GetServerType() string
}

// concrete product
type RestfulServer struct{}

func (r *RestfulServer) GetServerType() string {
	return "Restful"
}

type SoaServer struct{}

func (s *SoaServer) GetServerType() string {
	return "SOA"
}

// abstract factory
type Factory interface {
	CreateServer() Server
}

// concrete factory
type RestfulFactory struct{}

func (r *RestfulFactory) CreateServer() Server {
	return &RestfulServer{}
}

type SoaFactory struct{}

func (s *SoaFactory) CreateServer() Server {
	return &SoaServer{}
}
