/*
Copyright © 2022 Joey Yu <joey@itsjoeoui.com>
*/
package main

import (
	"mongodo/cmd"
	"mongodo/db"
)

func main() {
	db.Setup()
	cmd.Execute()
}
