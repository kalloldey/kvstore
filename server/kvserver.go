package main

import (
    	"fmt"
    	"net/http"
	"strings"
)

var kvmap map[string]string

func get( key string) string{
	val,exist:=kvmap[key]
        if exist==true{
        	return val
        } else {
        	return "#NOT_EXIST#"
        }
}

func set( key, value string){
	kvmap[key]=strings.Trim(value,"\n")
}

func gethandler(w http.ResponseWriter, r *http.Request) {
	spt:=strings.Trim(r.URL.Path[5:], " ")
	fmt.Fprintf(w, "%s",get(spt))
	fmt.Println("[GET] key: ",spt)
}

func sethandler(w http.ResponseWriter, r *http.Request) {
	spt:=string(r.URL.Path[5:])
	ind:=strings.Index(spt," ")
	key:=strings.Trim(spt[0:ind], " ")
	set(key,strings.Trim(spt[ind+1:]," "))
	fmt.Println("[SET] key: ",key)
//	fmt.Fprintf(w,"#%s#%s",key,kvmap[key])
}

func main() {
	fmt.Println("\nKey value server is running . . . ")
	kvmap=make(map[string]string)
	http.HandleFunc("/get/", gethandler)
	http.HandleFunc("/set/", sethandler)
	http.ListenAndServe(":8080", nil)
}
