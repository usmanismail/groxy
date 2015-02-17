package meta

type ContainerMetaStruct struct {
	Id      string
	Image   string
	Created int64
	Status  string
	Ports   []PortStruct
}

type PortStruct struct {
	IP          string
	PrivatePort int32
	PublicPort  int32
	Type        string
}
