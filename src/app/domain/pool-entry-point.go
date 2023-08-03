package domain

import "fmt"

// this file an entry point into the domain of
// the application. It need not remain, neither even
// the domain package. It is just here for illustration
// purposes.
//

func EnterPool(_ *PoolParameterSet) error {
	fmt.Println("---> 🎯 orpheus(alpha) ...")

	return nil
}
