## mmdb_block_search

 `mmdb_block_search`, is designed to search for IP addresses in a MaxMind DB (MMDB) file and, if found, pass the IP address to the `nali` command for further analysis. 

 ## Building

 ```
 go build -o mmdb_block_search main.go
 ```

 ## Usage

 ```
 ./mmdb_block_search /path/to/blocked.mmdb "101.40.68.3/32"
 ```

 ## Example output

 ```
 IP 101.40.68.3 is located in country code: CN
Network: 101.40.68.3
Mask: 255.255.255.255
 ```

 ```
 IP 1.1.1.1 is not found in the MMDB file
 ```

 ```
 2405:201:8000:6805::123 [印度 Reliance Jio Infocomm Limited (信实工业)]
Network: 2405:201:8000:6805::123
Mask: ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff
 ```