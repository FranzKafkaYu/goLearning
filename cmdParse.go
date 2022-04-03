/*
 * @Author: FranzKafka
 * @Date: 2022-04-03 10:09:47
 * @LastEditTime: 2022-04-03 10:29:57
 * @LastEditors: Franzkafka
 * @Description:this excersie will help you learn how to parse cmd line and do something
 * @FilePath: \goStudy\cmdParse.go
 */
package main

//包名定义之后导包，使用import关键字导入所依赖的包
import (
	"flag"
	"fmt"
	"os"
)

//使用方法：cmdParse show -localShowArgsint=2

func main() {
	//flag module init,flag.ExitOnError is an const variable
	runCmd:=flag.NewFlagSet("run",flag.ExitOnError)
	showCmd:=flag.NewFlagSet("show",flag.ExitOnError)
	var localRunArgsint int
	var localRunArgsbool bool

	var localShowArgsint int
	var localShowArgsBool bool
	runCmd.IntVar(&localRunArgsint,"localRunArgsint",0,"check runCmd int value")
	runCmd.BoolVar(&localRunArgsbool,"localRunArgsbool",false,"check runCmd bool value")
	showCmd.IntVar(&localShowArgsint,"localShowArgsint",0,"show showCmd int value")
	showCmd.BoolVar(&localShowArgsBool,"localShowArgsBool",false,"show showCmd bool value")
	//all flags settled donw,we need use flag.Parse() to enable parse
	flag.Parse()
	//通过os模块识别命令行
	switch os.Args[1]{
	case "run":
		fmt.Println("here we entered run cmd")
		err:=runCmd.Parse(os.Args[2:])
		if err!=nil{
			fmt.Println(err)
			return
		}
		//if we got localRunArgsint value,do something
		if(localRunArgsint!=0){
			fmt.Println("we got a runCmd int value as follows:")
			fmt.Println(localRunArgsint)
		}
		//if we got localRunArgsbool value,do something
		if(localRunArgsbool){
			fmt.Println("we got a runCmd bool value as true")
		}

	case "show":
		fmt.Println("here we entered show cmd")
		err:=showCmd.Parse(os.Args[2:])
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Println(localShowArgsint)
		fmt.Println(localShowArgsBool)
		if(localShowArgsint!=0){
			fmt.Println("we got a showCmd int value as follows:")
			fmt.Println(localShowArgsint)
		}
		if(localShowArgsBool){
			fmt.Println("we got a showCmd bool value as true")
		}
	default:
		fmt.Println("here we entered default")
		return
	}
}