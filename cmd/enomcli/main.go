package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/domains"
	"github.com/joho/godotenv"
)

type EnvKey string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := enomapi.NewClient("https://resellertest.enom.com/", os.Getenv("RESELLERID"), os.Getenv("APIKEY"))
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	domainsClient := &domains.Client{Client: client}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter domain name: ")
	dname, _ := reader.ReadString('\n')
	dname = strings.TrimSpace(dname)
	fmt.Printf("Checking availability for domain: %s\n", dname)

	domain := enomapi.NewDomain(dname)
	p, err := domainsClient.Check(domain)
	if err != nil {
		fmt.Printf("Error checking domain: %v\n", err)
		return
	}
	if !p.IsAvailable {
		fmt.Printf("Domain %s is not available\n", dname)
		return
	}

	fmt.Printf("Domain %s is available\n", dname)
	fmt.Print("Prices:\n")
	fmt.Printf("  Registration:   $%.2f\n", p.Prices.Registration)
	fmt.Printf("  Renewal:        $%.2f\n", p.Prices.Renewal)
	fmt.Printf("  Transfer:       $%.2f\n", p.Prices.Transfer)
	fmt.Printf("  Expected Price: $%.2f\n", p.Prices.ExpectedPrice)

	fmt.Printf("Would you like to Purchase %s? [y/n]", p.FQDN)
	confirm, _ := reader.ReadString('\n')
	if confirm != "y\n" && confirm != "Y\n" {
		fmt.Println("Purchase cancelled.")
		return
	}

	fmt.Printf("Checking registration requirements for .%s\n", domain.Extension)
	tld, err := domainsClient.GetTLDDetails(domain.Extension)
	if err != nil {
		fmt.Printf("Error getting TLD details: %v\n", err)
		return
	}

	if tld.Registration.DNSRequired {
		// TODO: Implement DNS server input handling
		fmt.Println("Domain requires DNS servers to be set at registration, but this feature is not yet implemented.")

		// fmt.Println("Domain requires DNS servers to be set at registration.")
		// fmt.Print("Enter comma-separated list of DNS servers: ")
		// dnsInput, _ := reader.ReadString('\n')
		// dnsInput = strings.TrimSpace(dnsInput)
		// dnsServers := strings.Split(dnsInput, ",")
		// for i := range dnsServers {
		// 	dnsServers[i] = strings.TrimSpace(dnsServers[i])
		// }
		// domain.NS = dnsServers
	}

	fmt.Printf("Result: %v\n", *tld)
}
