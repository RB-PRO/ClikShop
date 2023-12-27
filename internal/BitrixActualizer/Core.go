package actualizer

// Интерфейс магазина
type Shop interface {
	screper() (string, error)
}
