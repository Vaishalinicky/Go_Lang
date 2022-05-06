package modulec

import "fmt"

// Hello returns a greeting for the named person.
func Module_c(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hello, %v. Welcome in modulec", name)
    return message
}