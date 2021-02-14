package parsing

type Parser interface {
	Name() string
	ParseBuckwheats() ([]Buckwheat, error)
}
