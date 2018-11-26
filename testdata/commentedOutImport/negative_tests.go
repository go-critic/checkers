package checker_test

import (
	"fmt"
)

// Check that it doesn't try to check comments after the imports spec.

var (
	//"fmt"
	_ = fmt.Sprint
)

// "fmt"
//"fmt"

/*"fmt"*/
/* "fmt" */
