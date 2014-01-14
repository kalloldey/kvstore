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

func deletehandler(w http.ResponseWriter, r *http.Request) {
        spt:=strings.Trim(r.URL.Path[5:], " ")
	delete(kvmap,spt)
	fmt.Println("[DELETE] key:",spt)
}

func sethandler(w http.ResponseWriter, r *http.Request) {
	spt:=string(r.URL.Path[5:])
	ind:=strings.Index(spt," ")
	key:=strings.Trim(spt[0:ind], " ")
	set(key,strings.Trim(spt[ind+1:]," "))
	fmt.Println("[SET] key: ",key)
//	fmt.Fprintf(w,"#%s#%s",key,kvmap[key])
}

func copyhandler(w http.ResponseWriter, r *http.Request) {
        spt:=string(r.URL.Path[6:])
        ind:=strings.Index(spt," ")
        key1:=strings.Trim(spt[0:ind], " ")
	key2:=strings.Trim(spt[ind+1:]," ")
        set(key2,get(key1))
        fmt.Println("[COPY] keySrc: ",key1," keyDest: ",key2)
//      fmt.Fprintf(w,"#%s#%s",key,kvmap[key])
}

func main() {
	fmt.Println("\nKey value server is running . . . ")
	kvmap=make(map[string]string)
	http.HandleFunc("/get/", gethandler)
	http.HandleFunc("/set/", sethandler)
	http.HandleFunc("/copy/", copyhandler)	
	http.HandleFunc("/del/", deletehandler)
	http.ListenAndServe(":8080", nil)
}
