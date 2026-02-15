package matrix

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func ReadYaml(filepath string, es *EquationSystem) error {
	if _, err := os.Stat(filepath); filepath != "" && os.IsNotExist(err) {
		return fmt.Errorf("can't find config file at path: %s", filepath)
	}

	if err := cleanenv.ReadConfig(filepath, es); err != nil {
		return err
	}

	return nil
}
