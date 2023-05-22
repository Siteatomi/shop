package address

import "errors"

var (
	NoCitiesFound    error = errors.New("no cities found in database")
	NoAddressesFound error = errors.New("no addresses for selected user found")
	AddressNotFound  error = errors.New("requested address doesnt exists")
	YouAreNotAllowed error = errors.New("operation is not allowed for you")
)
