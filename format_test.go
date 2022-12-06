package stringvarformatter

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatter_Format(t *testing.T) {
	formatter := NewFormatter("$BUN $BUN $TIME", "$BUN", "$TIME")
	vars := formatter.GetVarMap()
	vars["$BUN"] = "bunbunbun"
	vars["$TIME"] = time.Now().Format(time.RFC3339)
	fmt.Println(formatter.Format(4, vars))
}
