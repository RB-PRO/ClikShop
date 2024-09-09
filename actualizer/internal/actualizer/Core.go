package actualizer

type Shop interface {
	Scraper() (string, error)
}
