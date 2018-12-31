package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "A Net Lookup CLI Tools"
	app.Usage = "net lookup tools"
	app.Version = "0.0.1"

	// set a destination variable for a flag
	cliFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "whoami",
		},
	}

	// declare a group command
	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Lookup DNS records for the give domain name.",
			Flags: cliFlags,
			Action: func(c *cli.Context) error {
				dns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println("[!] ", err)
					return err
				}
				for i := 0; i < len(dns); i++ {
					fmt.Println("[*] dns : ", dns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up host IPV4 or IPV6 address.",
			Flags: cliFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println("[*] ip address :", ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Lookups the canonical name for the give host.",
			Flags: cliFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println("[!]", err)
					return err
				}
				fmt.Println("[*] cname : ", cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Lookups the DNS MX records for the given domain name sorted by preference.",
			Flags: cliFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println("[!]", err)
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println("[*] mx records : ", mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
		{
			Name:  "DNS",
			Usage: "Lookups the DNS TXT records for the given domain name.",
			Flags: cliFlags,
			Action: func(c *cli.Context) error {
				dns_text, err := net.LookupTXT(c.String("host"))
				if err != nil {
					fmt.Println("[!] ", err)
					return err
				}
				for i := 0; i < len(dns_text); i++ {
					fmt.Println("[*] dns list : ", dns_text[i])
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
