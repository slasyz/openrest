package formatters

type Formatter interface {
	Format(filename string) error
}
