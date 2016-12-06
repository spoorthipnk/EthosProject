package main

import(
"ethos/efmt"
"ethos/ethos"
"ethos/syscall"
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

func Walk(p string){
	efmt.Println("walking")
        filename := " "
        listFileName := " "
	fd,status := ethos.OpenDirectoryPath(p)
	efmt.Println("Status ",status)
        for {	
                name,status := ethos.GetNextName(fd,filename)
                if status!=syscall.StatusOk {
                        break
                }
                filename = string(name)
		if filename == "." {
			continue
		}
		if filename == ".." {	
			continue
		}
		new_path := p + "/" +filename
		efmt.Println("FINALPATH: ",new_path)
		if isDir(fd,filename) {
			listFileName = listFileName + "  " + filename
			efmt.Println("filename",filename)
			efmt.Println("spoorthi",listFileName)
			Walk(new_path)
		}      
        }
}

func main(){
	me:=syscall.GetUser()
	p:="/user/"+me+"/myDir"
	efmt.Println("Path: ",p)
	Walk(p)
}
