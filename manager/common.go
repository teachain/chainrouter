package manager

import "strings"

func toPureName(name string)string{
	name=strings.ReplaceAll(name,"\\", "")
	name=strings.ReplaceAll(name,"/", "")
	name=strings.ReplaceAll(name,".", "")
	return name
}
