package main

import (
        "os"
        "encoding/hex"
        "math/big"
        "fmt"
        "net"
        "strings"
)

func IsIPv4(address string) bool {
        return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
        return strings.Count(address, ":") >= 2
}


// Inet6_Aton converts an IP Address (IPv4 or IPv6) net.IP object to a hexadecimal
// representaiton. This function is the equivalent of
// inet6_aton({{ ip address }}) in MySQL.
func Inet6_Aton(ip net.IP) string {
        ipv4 := false
        if ip.To4() != nil {
                ipv4 = true
        }

        ipInt := big.NewInt(0)
        if ipv4 {
                ipInt.SetBytes(ip.To4())
                ipHex := hex.EncodeToString(ipInt.Bytes())
                return ipHex
        }

        ipInt.SetBytes(ip.To16())
        ipHex := hex.EncodeToString(ipInt.Bytes())
        return ipHex
}

func Reverse[T any](original []T) (reversed []T) {
        reversed = make([]T, len(original))
        copy(reversed, original)

        for i := len(reversed)/2 - 1; i >= 0; i-- {
                tmp := len(reversed) - 1 - i
                reversed[i], reversed[tmp] = reversed[tmp], reversed[i]
        }

        return
}

func Usage() {
        fmt.Println(os.Args[0] + " <ip>")
        os.Exit(1)
}

func main() {
        if(len(os.Args) != 2) {
                Usage()
        }

        remote_ip := os.Args[1]

        if(IsIPv4(remote_ip)) {
                ipparts := strings.Split(remote_ip, ".")
                ipparts = ipparts[:len(ipparts) - 1]
                query := strings.Join(Reverse(ipparts), ".") + ".origin.asn.cymru.com"

                //txtrecords, _ := net.LookupTXT("102.160.192.origin.asn.cymru.com")
                txtrecords, _ := net.LookupTXT(query)

                //fmt.Println(txtrecords)
                //fmt.Println("The length of the slice is:", len(txtrecords))
                if(len(txtrecords) > 0) {
                        // do stuff with the first result
                        split:= strings.Split(txtrecords[0], "|")
                        asn := split[0]
                        fmt.Println(asn)
                } else {
                        // no records found
                        fmt.Println("ERR")
                }
        } else if(IsIPv6(remote_ip)) {
                // do IPv6 stuff
                hexstring := Inet6_Aton(net.ParseIP(remote_ip))
                ipparts := strings.Split(hexstring[:12], "")
                //fmt.Println(ipparts)

                query := strings.Join(Reverse(ipparts), ".") + ".origin6.asn.cymru.com"

                //txtrecords, _ := net.LookupTXT("102.160.192.origin.asn.cymru.com")
                txtrecords, _ := net.LookupTXT(query)

                //fmt.Println(txtrecords)
                //fmt.Println("The length of the slice is:", len(txtrecords))
                if(len(txtrecords) > 0) {
                        // do stuff with the first result
                        split:= strings.Split(txtrecords[0], "|")
                        asn := split[0]
                        fmt.Println(asn)
                } else {
                        // no records found
                        fmt.Println("ERR")
                }
        } else {
                fmt.Println("ERR")
        }
}
