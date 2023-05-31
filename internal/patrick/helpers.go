package patrick

import (
	"fmt"
)

type Config struct {
	InputDir    string
	OutputDir   string
	ExcludeFile string
}

func dumpConfig(cfg Config) {
	doubleLine()

	fmt.Println(fmt.Sprintf("Input directory : %s", cfg.InputDir))
	fmt.Println(fmt.Sprintf("Output directory: %s", cfg.OutputDir))
	fmt.Println(fmt.Sprintf("Exclude files   : %s", cfg.ExcludeFile))

	doubleLine()
}

func doubleLine() {
	fmt.Println("================================================================================")
}

func singleLine() {
	fmt.Println("--------------------------------------------------------------------------------")
}
