package main

//func main() {
//	counts := make(map[string]map[string]int)
//
//	if len(os.Args[1:])<1 {
//		temp := make(map[string]int)
//		scanLine(os.Stdin,temp)
//		counts[""] = temp
//	}else {
//		for _,fileName := range os.Args[1:]{
//			f,err := os.Open(fileName)
//
//			if err != nil {
//				fmt.Fprint(os.Stderr,"打开文件错误，错误信息为：%v\n",err)
//				continue
//			}
//
//			temp := make(map[string]int)
//			scanLine(f,temp)
//
//			counts[fileName] = temp
//		}
//	}
//
//	for fileName,tempMap := range counts{
//		for line,count := range tempMap{
//			if count > 1 {
//				if fileName == ""{
//					fmt.Printf("重复行的内容为：%v\t重复的次数为：%v\n",line,count)
//				}else {
//					name := strings.Split(string(fileName),"\\")
//					fmt.Printf("重复行的内容为：%v\t重复的次数为：%v\t所属文件是：%s\n",line,count,name[len(name)-1])
//				}
//			}
//		}
//	}
//}
//
//func scanLine(f *os.File,lines map[string]int)  {
//	input := bufio.NewScanner(f)
//
//	for input.Scan(){
//		lines[input.Text()]++
//	}
//}
