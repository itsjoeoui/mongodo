/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
