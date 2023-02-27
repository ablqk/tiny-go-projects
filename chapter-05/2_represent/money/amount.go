package money

// Amount defines a decimal of money in a given currency.
type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("decimal is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{quantity: quantity, currency: currency}, nil
}
