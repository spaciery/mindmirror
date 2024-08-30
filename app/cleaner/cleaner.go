package cleaner

import (
	"errors"
	"fmt"
	"os"
)

type Cleaner struct {
	Directory string
}

var NothingToClean = errors.New("nothing to clean")

func (c *Cleaner) Clean() error {

	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		return fmt.Errorf("nothing to clean %w", NothingToClean)
	}

	err := os.RemoveAll(c.Directory)
	if err != nil {
		return err
	}
	return nil
}
