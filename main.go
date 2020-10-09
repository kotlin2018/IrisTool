//package irisTool
package main

import (
	_ "irisTool/commands/new"
	commands "irisTool/commands/new"
)
func main() {

	//pflag.Parse()
	//currentpath, _ := os.Getwd()
	//if !utils.IsExist(currentpath) {
	//	log.Printf("Application '%s' already exists", currentpath)
	//	os.Exit(0)
	//}
	//appName := ""
	//
	//for v := range pflag.Args() {
	//	if v == 0 {
	//		if pflag.Args()[v] == "new" {
	//			if len(pflag.Args()) > 1 {
	//				if pflag.Args()[v+1] != "" {
	//					appName = pflag.Args()[v+1]
	//					commands.CreatedApp(currentpath, appName)
	//				}
	//			} else {
	//				appName = "irisApp"
	//				commands.CreatedApp(currentpath, appName)
	//			}
	//
	//		}
	//	}
	//}
	commands.Mysql.Create(&commands.UserInfo{1,"lisi"})

}
