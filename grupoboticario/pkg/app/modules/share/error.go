package share

import (
	"fmt"
	"strings"
)

type DomainError struct {
	Domain      string
	Module      string
	Err         string
	Description string
} //@name ErrorResponse

func (err DomainError) Error() string {
	return fmt.Sprintf(
		"%s|%s|%s",
		strings.ToUpper(err.Domain),
		strings.ToUpper(err.Module),
		strings.ToUpper(err.Err),
	)
}
