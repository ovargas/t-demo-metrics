package dummy

type Error string

var (
	ErrUpsIAteABug Error = "ups I ate a bug"
)

var _ error = (*Error)(nil)

func (e Error) Error() string {
	return string(e)
}
