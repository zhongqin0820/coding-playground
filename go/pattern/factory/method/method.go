package method

// abstrct product
type Server interface {
	GetServerType() string
}

// concrete product
type RestfulServer struct{}

type SoaServer struct{}

func (r *RestfulServer) GetServerType() string {
	return "Restful"
}

func (s *SoaServer) GetServerType() string {
	return "SOA"
}

// abstract factory
type Factory interface {
	CreateServer() Server
}

// concrete factory
type RestfulFactory struct{}

type SoaFactory struct{}

func (r *RestfulFactory) CreateServer() Server {
	return &RestfulServer{}
}

func (s *SoaFactory) CreateServer() Server {
	return &SoaServer{}
}
