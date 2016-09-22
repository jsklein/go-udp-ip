package main
import (
    "fmt" 
    "net"  
    "reflect"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {   
	response := addr.IP.String()
    _,err := conn.WriteToUDP([]byte(response), addr)
    // fmt.Printf("Success")
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 1024,
        IP: net.ParseIP("0.0.0.0"),
    }
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    for {
        _,remoteaddr,err := ser.ReadFromUDP(p)
        // fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
	h := "\xff\xff\xff\xff\x67\x65\x74\x73\x74\x61\x74\x75\x73\x0a"
    //m := []byte(h)
	//pp := []byte(p)[0:14]
	// use this if we want to add a key
	if reflect.DeepEqual(h, h) {
		go sendResponse(ser, remoteaddr)	
		}
	// fmt.Println(pp,m)
	if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        continue
      }
}
