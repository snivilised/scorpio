package domain

import "fmt"

// this file an entry point into the domain of
// the application. It need not remain, neither even
// the domain package. It is just here for illustration
// purposes.
//

func EnterWidget(ps *WidgetParameterSet) error {
	if ps.Directory != "" {
		fmt.Printf("---> Enter(widget): directory: '%v'\n", ps.Directory)
	}

	return nil
}
