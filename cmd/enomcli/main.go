package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/domains"
	domainmanagement "github.com/hazzakins/go-enomapi/domains/domain-management"
	"github.com/joho/godotenv"
)

// EnvKey represents an environment variable key used by the CLI.
type EnvKey string

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select Env;\n")
	fmt.Print("  1: PROD\n")
	fmt.Print("  2: DEV\n")
	env, _ := reader.ReadString('\n')

	switch env {
	case "2\n":
		env = ".env.dev"
	case "1\n":
		env = ".env"
	}

	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("Connecting client with Interface %v\n", os.Getenv("INTERFACE"))
	client, err := enomapi.NewClient(os.Getenv("INTERFACE"), os.Getenv("RESELLERID"), os.Getenv("APIKEY"))
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	fmt.Print("Select Mode;\n")
	fmt.Print("  1: Buy Domain\n")
	fmt.Print("  2: Get Domain Info\n")
	fmt.Print("Enter mode number: ")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	switch mode {
	case "1":
		buyDomain(client, reader)
	case "2":
		getDomainInfo(client, reader)
	default:
		fmt.Println("Invalid mode selected.")
	}

}

func buyDomain(client *enomapi.Client, reader *bufio.Reader) {
	domainsClient := &domains.Client{Client: client}

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
		fmt.Printf(
			".%s requires name servers at registration (min %d, max %d); the CLI does not support this yet.\n",
			domain.Extension,
			tld.Registration.DNSMinServers,
			tld.Registration.DNSMaxServers,
		)
		return
	}

	extAttrs, err := domainsClient.GetExtAttributes(domain.Extension)
	if err != nil {
		fmt.Printf("Error getting extended attributes: %v\n", err)
		return
	}
	if len(extAttrs.Attributes) > 0 {
		fmt.Printf("Extended attributes are required for .%s. The CLI does not support collecting them yet.\n", domain.Extension)
		for _, attr := range extAttrs.Attributes {
			reqLabel := "optional"
			if attr.Required > 0 {
				reqLabel = "required"
			}
			fmt.Printf("  - %s (%s)\n", attr.Name, reqLabel)
		}
		return
	}

	if p.IsPremium || p.IsPlatinum || p.IsEAP {
		fmt.Println("Premium, platinum, and EAP domains require additional purchase parameters not supported by the CLI.")
		return
	}

	purchase, err := domainsClient.Purchase(domain)
	if err != nil {
		fmt.Printf("Error purchasing domain: %v\n", err)
		return
	}

	fmt.Printf("Order ID: %s\n", purchase.OrderID)
	fmt.Printf("Order Status: %s\n", purchase.OrderStatus)
	fmt.Printf("Order Description: %s\n", purchase.OrderDescription)
	fmt.Printf("Charged Price: $%.2f\n", purchase.Price)
	if purchase.OrderCompleted {
		if !purchase.RegistrationDate.IsZero() {
			fmt.Printf("Registration Date: %s\n", purchase.RegistrationDate)
		}
		if !purchase.ExpirationDate.IsZero() {
			fmt.Printf("Expiration Date: %s\n", purchase.ExpirationDate)
		}
	} else {
		fmt.Println("Order was queued; check order status later.")
	}
}

func getDomainInfo(client *enomapi.Client, reader *bufio.Reader) {
	manaementClient := &domainmanagement.Client{Client: client}

	fmt.Print("Enter domain name: ")
	dname, _ := reader.ReadString('\n')
	dname = strings.TrimSpace(dname)

	domain := enomapi.NewDomain(dname)
	info, err := manaementClient.GetDomainInfo(domain)
	if err != nil {
		fmt.Printf("Error getting domain info: %v\n", err)
		return
	}

	fmt.Printf("Domain Info for %s:\n", dname)
	fmt.Printf("  Domain Name ID:      %d\n", info.DomainName.DomainNameID)
	fmt.Printf("  SLD:                 %s\n", info.DomainName.SLD)
	fmt.Printf("  TLD:                 %s\n", info.DomainName.TLD)
	fmt.Printf("  Expiration:          %s\n", info.Status.Expiration)
	fmt.Printf("  Registrar:           %s\n", info.Status.Registrar)
	fmt.Printf("  Registration Status: %s\n", info.Status.RegistrationStatus)
}
