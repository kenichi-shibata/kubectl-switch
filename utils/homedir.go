package utils

import (
	"log"
	"os/user"
)
// Homedir return the homedir string
func Homedir() string {
usr, err := user.Current()
if err != nil {
		log.Fatal(err)
}

return usr.HomeDir 
}