package main

import "fmt"

/*
Markdown is the implemented type of Formatter.
*/
type Markdown struct {
}

/*
Hugo is the implemented type of Formatter.
*/
type Hugo struct {
}

func (*Markdown) format(fileName, thumbnail string) string {
	return fmt.Sprintf("[![Title](%s)](%s)", thumbnail, fileName)
}

func (*Hugo) format(fileName, thumbnail string) string {
	return fmt.Sprintf("{{< figure src=\"%s\" link=\"%s\" title=\"\" >}}", thumbnail, fileName)
}
