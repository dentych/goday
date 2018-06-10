package main

import (
	"github.com/dentych/goday/apprication/api"
)

func main() {
	api.StartApi()
	//buf := make([]byte, 20)
	//for {
	//	read, err := res.Read(buf)
	//	if read > 0 {
	//		fmt.Println("Amount read: ", read)
	//		//fmt.Println("Text:", string(buf))
	//	} else {
	//		fmt.Println("Such empty")
	//	}
	//
	//	if err != nil {
	//		fmt.Printf("%T\n", err)
	//		fmt.Println(err.Error())
	//		if err == io.EOF {
	//			fmt.Println("End of line met!")
	//		}
	//		break
	//	}
	//}
}
