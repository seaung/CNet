#### About CNet go command line tools
>>>
    
    this is a network lookup command line tools.

#### Build
>>>
    
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o cnet cnet.go

    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o cnet cnet.go

#### Usage
>>>
    
    1. $ cnet --help

    NAME:
    A Net Lookup CLI Tools - net lookup tools

    USAGE:
    cnet [global options] command [command options] [arguments...]

    VERSION:
    0.0.1

    COMMANDS:
        ns       Lookup DNS records for the give domain name.
        ip       Looks up host IPV4 or IPV6 address.
        cname    Lookups the canonical name for the give host.
        mx       Lookups the DNS MX records for the given domain name sorted by preference.
        DNS      Lookups the DNS TXT records for the given domain name.
        help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
    --help, -h     show help
    --version, -v  print the version

    2. $ cnet DNS --host xxx.com

    [*] dns list :  v=spf1 include:spf1.xxx.com include:spf2.xxx.com include:spf3.xxx.com a mx ptr -all
    [*] dns list :  xxx-site-verification=GHb98-6msqyx_qqjGl5eRatD3QTHyVB6-xQ3gJB5UwM

    3. $ cnet ns --host xxx.com

    [*] dns.xxx.com.
    [*] ns7.xxx.com.
    [*] ns2.xxx.com.
    [*] ns3.xxx.com.
    [*] ns4.xxx.com.
    
    4. $ cnet cname --host xxx.com

    [*] cname :  xxx.com.

    5. $ cnet mx --host xxx.com

    [*] mx records :  mx.maillb.xxx.com. 10
    [*] mx records :  mx.n.xxx.com. 15
    [*] mx records :  mx1.xxx.com. 20
    [*] mx records :  jpmx.xxx.com. 20
    [*] mx records :  mx50.xxx.com. 20

    6. $ cnet ip --host xxx.com

    [*] ip address  0.0.0.0
    [*] ip address  0.0.0.0

    7. $ cnet --version

    A Net Lookup CLI Tools version 0.0.1