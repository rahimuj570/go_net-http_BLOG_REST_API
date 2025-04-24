package enums

type Role int

const (
	ADMIN Role = iota
	READER
	AUTHOR
)

func (r Role) String() string {
	switch r {
	case ADMIN:
		return "ADMIN"
	case READER:
		return "READER"
	case AUTHOR:
		return "AUTHOR"
	default:
		return "UNKNOWN ROLE"
	}
}
