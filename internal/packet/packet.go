package packet

type Packet interface {
	IdType() int
	Value() int
	Version() int
	VersionSum() int
}
