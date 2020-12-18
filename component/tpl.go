package component

import "fmt"

func Tpl(s string) string {
	return fmt.Sprintf(`{"type": "tpl", "tpl":"%s"}`, s)
}
