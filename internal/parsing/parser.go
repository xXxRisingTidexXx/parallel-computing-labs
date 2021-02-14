package parsing

type Parser interface {
	ParseBuckwheats() ([]Buckwheat, error)
}
