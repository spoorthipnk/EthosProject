package main
import(

"ethos/efmt"
"ethos/ethos"
"ethos/syscall"
"log"

)
func isDir(fd syscall.Fd,dirName string)(bool) {
	efmt.Println("isDir", dirName)
	fd1,status1 := ethos.OpenDirectory(fd,dirName)
	useless(fd1)
	if status1 == 0 {
		efmt.Println("success")
		return true
	}
	return false	
}
func useless(fd syscall.Fd){
	}

func Walk(fd syscall.Fd,p string){
	efmt.Println("walking")
        filename := " "
        listFileName := " "
        for {
		count := 0
                name,status := ethos.GetNextName(fd,filename)
                if status!=syscall.StatusOk {
                        break
                }
                filename = string(name)
		
		if filename == "." {
			efmt.Println("Filename. ",filename)
			count = 1
		}
		if filename == ".." {
			count = 1
		}
		if count == 0 {
		efmt.Println("Filename1",filename)
		if isDir(fd,filename) {
			listFileName = listFileName + "  " + filename
			efmt.Println("filename",filename)
			efmt.Println("spoorthi",listFileName)
			p = p + "/"+filename
			efmt.Println("path",p)
			fd,status := ethos.OpenDirectoryPath(p)
			efmt.Println(status)
			Walk(fd,p)
		}
		}       
        }
}

func main(){
	me:=syscall.GetUser()
	p:="/user/"+me+"/myDir"
	efmt.Println("Path: ",p)
	fd, status := ethos.OpenDirectoryPath(p)
	if status != syscall.StatusOk {
	log.Fatalf ("Error opening %v: %v\n", p, status)
	}
	Walk(fd,p)
}
