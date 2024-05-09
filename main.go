package main

import (
        "fmt"
        "log"
        "net"
        "os"
        "os/exec"

        "github.com/oschwald/maxminddb-golang"
)

func main() {
        if len(os.Args) != 3 {
                fmt.Println("Usage: ./mmdb_block_search /path/to/mmdb ip")
                os.Exit(1)
        }

        mmdbPath := os.Args[1]
        ipStr := os.Args[2]

        // Parse the IP address or CIDR range
        ip, ipNet, err := net.ParseCIDR(ipStr)
        if err != nil {
                // Try parsing as single IP address
                ip = net.ParseIP(ipStr)
                if ip == nil {
                        fmt.Println("Invalid IP address or CIDR range")
                        os.Exit(1)
                }
        }

        // Open the MMDB file
        db, err := maxminddb.Open(mmdbPath)
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        // Define the struct for storing MMDB data
        var record struct {
                Country struct {
                        IsoCode string `maxminddb:"iso_code"`
                } `maxminddb:"country"`
        }

        // Lookup the IP address in the MMDB file
        err = db.Lookup(ip, &record)
        if err != nil {
                log.Fatal(err)
        }

        // If IP exists, pass it to nali
        if record.Country.IsoCode != "" {
                cmd := exec.Command("nali", ip.String())
                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr
                err := cmd.Run()
                if err != nil {
                        log.Fatal(err)
                }
        } else {
                fmt.Printf("IP %s is not found in the MMDB file\n", ip)
        }

        // If the input was a CIDR range, also print out the network address and mask
        if ipNet != nil {
                fmt.Printf("Network: %s\n", ipNet.IP)
                fmt.Printf("Mask: %s\n", net.IP(ipNet.Mask))
        }
}