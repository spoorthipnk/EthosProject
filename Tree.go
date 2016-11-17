package main
import(

"ethos/efmt"
"ethos/ethos"
"ethos/syscall"
"log"

)

func main(){
me:=syscall.GetUser()
p:="/root/EthosProject/"
fd, status := ethos.OpenDirectoryPath(p)
if status != syscall.StatusOk {
log.Fatalf ("Error opening %v: %v\n", p, status)
}

filename:= " "
listFileName := ""
data:= FileType{filename}
data.Write(fd)
data.WriteVar(p+"Direc")
for {
name,status := ethos.GetNextName(fd,filename)
filename = string(name)
if status!=syscall.StatusOk {
break
}
listFileName = listFileName + " "+filename
}
efmt.Println("spoorthi",listFileName,status,p)
efmt.Println(me)
}
