package cell

type Cell struct {
	DBPath string
	Tx     string // Need import Tx data
}

func NewCell() *Cell {
	return &Cell{"", ""}
}
