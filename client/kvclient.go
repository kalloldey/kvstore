package main

import (
    	"fmt"
    	"net/http"
	"io/ioutil"
	"strings"
	"os"
	"bufio"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Key Value Server\nType 'help' to know more")
	reader := bufio.NewReader(os.Stdin)
    	for {
		fmt.Print("|| KeyValueClient>> ")
	        line, err := reader.ReadString('\n')

        	if err != nil {
	            fmt.Println("[Error] in taking input")
        	    break
	        }
		if line=="\n"{
			fmt.Println("[Exit]")
			break
		}

//        	fmt.Print(line)
		var arr []string	
		arr=strings.Fields(line)
		if arr[0]=="exit"{	
			fmt.Println("[Exit]")
			break
		}
		if arr[0]=="get" {
			if len(arr)!=2{
				fmt.Println("[GET] {Error}Format mismatch of get request. See help")
			} else {
				
			 	resp, err := http.Get("http://localhost:8080/get/"+arr[1])
        			if err != nil {
//			                log.Fatal(err)
					fmt.Printf("[GET] Value not found for the key : %s",arr[1])
		        	}
		        	defer resp.Body.Close()
			        body, err := ioutil.ReadAll(resp.Body)
			        fmt.Printf("[GET] Value of key %s: %s", arr[1], body)
				fmt.Println()
			}
    		} else if arr[0]=="set" {
			if(len(arr)<2){
				  fmt.Println("[SET] {Error}Format mismatch of set request. See help")
			}else{
				resp, err := http.Get("http://localhost:8080/set/"+arr[1]+" "+strings.SplitAfterN(line, " ", 3)[2])
 		       	        if err != nil {
//              	        	log.Fatal(err)                             
                        	}
	        	        defer resp.Body.Close()
//		      	        body, err := ioutil.ReadAll(resp.Body)
//		              	fmt.Printf("[SET]: Returned by server = %s", body)
				fmt.Println("[SET] Value set for the key: ",arr[1])
			}
		}else if arr[0]=="help" {
			fmt.Println("\nKey Value Client Help")
			fmt.Println("$get <key>: Will print the value of the key if it exists in the server")
			fmt.Println("$set <key> <value>: Will set the value to the key in the server")
			fmt.Println("$exit: Will exit the program")
			fmt.Println("$help: Will print this screen")
			fmt.Println("Empty line: Will exit the program\n")
		}else {
			fmt.Println("[Error] Invalid Menu Entry. Type help to see more")
		}
	}
}
