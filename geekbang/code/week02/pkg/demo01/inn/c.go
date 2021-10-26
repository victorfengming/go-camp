package inn

import (
	"github.com/pkg/errors"
	"os"
)

func C() error {
	_, err := os.Open("abc")
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}
