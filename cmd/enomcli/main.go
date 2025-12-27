package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/domains"
	domainmanagement "github.com/hazzakins/go-enomapi/domains/domain-management"
	"github.com/hazzakins/go-enomapi/response"
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
	fmt.Print("  3: DNSSEC Operations\n")
	fmt.Print("  4: Get TLD Pricing\n")
	fmt.Print("Enter mode number: ")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	switch mode {
	case "1":
		buyDomain(client, reader)
	case "2":
		getDomainInfo(client, reader)
	case "3":
		dnssecOperations(client, reader)
	case "4":
		getTLDPricing(client, reader)
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
			".%s requires name servers at registration (min %d, max %d).\n",
			domain.Extension,
			tld.Registration.DNSMinServers,
			tld.Registration.DNSMaxServers,
		)
	}

	extAttrs, err := domainsClient.GetExtAttributes(domain.Extension)
	if err != nil {
		fmt.Printf("Error getting extended attributes: %v\n", err)
		return
	}
	var extendedAttributes map[string]string
	if len(extAttrs.Attributes) > 0 {
		var ok bool
		extendedAttributes, ok = readExtendedAttributes(reader, domain.Extension, extAttrs.Attributes)
		if !ok {
			return
		}
	}
	if value, ok := extendedAttributes["registered_for"]; ok {
		extendedAttributes["RegisteredFor"] = value
	}

	if p.IsPremium || p.IsPlatinum || p.IsEAP {
		fmt.Println("Premium, platinum, and EAP domains require additional purchase parameters not supported by the CLI.")
		return
	}

	var useDNS string
	var nameServers []string
	if tld.Registration.DNSRequired {
		if useEnomDNS(reader) {
			useDNS = "default"
		} else {
			nameServers = readNameServers(reader, tld.Registration.DNSMinServers, tld.Registration.DNSMaxServers)
			if len(nameServers) == 0 {
				return
			}
		}
	}

	if len(extendedAttributes) > 0 {
		fmt.Println("Extended attributes to submit:")
		for key, value := range extendedAttributes {
			fmt.Printf("  %s=%s\n", key, value)
		}
	}

	var purchase *response.DomainPurchase
	if useDNS != "" || len(nameServers) > 0 || len(extendedAttributes) > 0 {
		purchase, err = domainsClient.PurchaseWithOptions(domains.PurchaseRequest{
			Domain:             domain,
			UseDNS:             useDNS,
			NameServers:        nameServers,
			ExtendedAttributes: extendedAttributes,
		})
	} else {
		purchase, err = domainsClient.Purchase(domain)
	}
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

func dnssecOperations(client *enomapi.Client, reader *bufio.Reader) {
	managementClient := &domainmanagement.Client{Client: client}

	fmt.Print("Select DNSSEC Operation;\n")
	fmt.Print("  1: Get DNSSEC\n")
	fmt.Print("  2: Add DNSSEC\n")
	fmt.Print("  3: Delete DNSSEC\n")
	fmt.Print("Enter operation number: ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	switch operation {
	case "1":
		getDnsSec(managementClient, reader)
	case "2":
		addDnsSec(managementClient, reader)
	case "3":
		deleteDnsSec(managementClient, reader)
	default:
		fmt.Println("Invalid DNSSEC operation selected.")
	}
}

func getTLDPricing(client *enomapi.Client, reader *bufio.Reader) {
	domainsClient := &domains.Client{Client: client}

	fmt.Print("Enter TLD (e.g. com). Leave blank for all: ")
	tldInput, _ := reader.ReadString('\n')
	tldInput = strings.TrimSpace(tldInput)
	tldInput = strings.TrimPrefix(tldInput, ".")

	var years *int
	var useQtyEngine *bool
	fmt.Print("Years for multi-year pricing (1, 2, 5, 10; blank for default): ")
	yearsInput, _ := reader.ReadString('\n')
	yearsInput = strings.TrimSpace(yearsInput)
	if yearsInput != "" {
		value, err := strconv.Atoi(yearsInput)
		if err != nil {
			fmt.Printf("Invalid years value: %s\n", yearsInput)
			return
		}
		years = &value
		useQtyEngineValue := true
		useQtyEngine = &useQtyEngineValue
	}

	resp, err := domainsClient.PEGetDomainPricing(domains.PEGetDomainPricingRequest{
		UseQtyEngine: useQtyEngine,
		Years:        years,
	})
	if err != nil {
		fmt.Printf("Error getting TLD pricing: %v\n", err)
		return
	}

	products := resp.Products
	if tldInput != "" {
		filtered := make([]response.PEDomainPricingProduct, 0, 1)
		for _, product := range products {
			if strings.EqualFold(product.TLD, tldInput) {
				filtered = append(filtered, product)
			}
		}
		products = filtered
		if len(products) == 0 {
			fmt.Printf("No pricing found for .%s\n", tldInput)
			return
		}
	}

	fmt.Println("TLD pricing:")
	for _, product := range products {
		fmt.Printf("  .%s (ID %d)\n", product.TLD, product.TLDID)
		fmt.Printf("    Register: $%.2f (reseller $%.2f) %s\n", product.RegisterPrice, product.ResellerPriceReg, formatEnabled(product.RegisterEnabled))
		fmt.Printf("    Renew:    $%.2f (reseller $%.2f) %s\n", product.RenewPrice, product.ResellerPriceRenew, formatEnabled(product.RenewEnabled))
		fmt.Printf("    Transfer: $%.2f (reseller $%.2f) %s\n", product.TransferPrice, product.ResellerPriceTran, formatEnabled(product.TransferEnabled))
		if product.RGPEnabled || product.RGPPrice != 0 || product.ResellerPriceRGP != 0 {
			fmt.Printf("    RGP:      $%.2f (reseller $%.2f) %s\n", product.RGPPrice, product.ResellerPriceRGP, formatEnabled(product.RGPEnabled))
		}
	}
}

func formatEnabled(enabled bool) string {
	if enabled {
		return "[enabled]"
	}
	return "[disabled]"
}

func getDnsSec(client *domainmanagement.Client, reader *bufio.Reader) {
	domain, ok := readDomain(reader)
	if !ok {
		return
	}

	resp, err := client.GetDnsSec(domain)
	if err != nil {
		fmt.Printf("Error getting DNSSEC info: %v\n", err)
		return
	}

	fmt.Printf("DNSSEC for %s:\n", formatDomain(domain))
	fmt.Printf("  Response Code: %d\n", resp.ResponseCode)
	fmt.Printf("  Response Message: %s\n", resp.ResponseMessage)
	if len(resp.Keys) == 0 {
		fmt.Println("  No DNSSEC keys found.")
		return
	}

	for i, key := range resp.Keys {
		fmt.Printf("  Key %d:\n", i+1)
		fmt.Printf("    Algorithm:  %d\n", key.Algorithm)
		fmt.Printf("    Digest:     %s\n", key.Digest)
		fmt.Printf("    DigestType: %d\n", key.DigestType)
		fmt.Printf("    KeyTag:     %d\n", key.KeyTag)
	}
}

func addDnsSec(client *domainmanagement.Client, reader *bufio.Reader) {
	domain, ok := readDomain(reader)
	if !ok {
		return
	}

	printDnsSecFieldOptions()
	if wantComputeDefaultYes(reader, "Generate DNSKEY and DS now (RSA-2048, flags 257)? [Y/n]: ") {
		digestType, ok := readInt(reader, "Digest Type: ")
		if !ok {
			return
		}
		privateKeyPEM, publicKey, err := generateRSADNSKEY()
		if err != nil {
			fmt.Printf("Error generating DNSKEY: %v\n", err)
			return
		}
		digest, keyTag, err := computeDnsSecFromDNSKEY(domain, 257, 3, 8, publicKey, digestType)
		if err != nil {
			fmt.Printf("Error computing DNSSEC values: %v\n", err)
			return
		}
		fmt.Println("Generated DNSKEY record:")
		fmt.Printf("%s IN DNSKEY 257 3 8 %s\n", formatDomain(domain), publicKey)
		fmt.Println("Generated private key (store securely):")
		fmt.Println(privateKeyPEM)
		fmt.Printf("Computed Digest: %s\n", digest)
		fmt.Printf("Computed Key Tag: %d\n", keyTag)

		sendDnsSec(client, domain, digestType, digest, keyTag, 8, reader)
		return
	}

	dnskeyRecord := readOptional(reader, "Paste DNSKEY record (optional): ")
	var algorithm int
	var digest string
	var digestType int
	var keyTag int

	if dnskeyRecord != "" {
		flags, protocol, dnskeyAlg, publicKey, err := parseDNSKEYRecord(dnskeyRecord)
		if err != nil {
			fmt.Printf("Invalid DNSKEY record: %v\n", err)
			return
		}
		algorithm = dnskeyAlg
		fmt.Printf("Using DNSKEY algorithm %d from record.\n", algorithm)
		digestType, ok = readInt(reader, "Digest Type: ")
		if !ok {
			return
		}
		digest, keyTag, err = computeDnsSecFromDNSKEY(domain, flags, protocol, algorithm, publicKey, digestType)
		if err != nil {
			fmt.Printf("Error computing DNSSEC values: %v\n", err)
			return
		}
		fmt.Printf("Computed Digest: %s\n", digest)
		fmt.Printf("Computed Key Tag: %d\n", keyTag)
	} else if wantComputeDefaultYes(reader, "Compute digest and key tag from DNSKEY fields? [Y/n]: ") {
		algorithm, ok = readInt(reader, "Algorithm: ")
		if !ok {
			return
		}
		digestType, ok = readInt(reader, "Digest Type: ")
		if !ok {
			return
		}
		flags, ok := readInt(reader, "DNSKEY Flags (e.g. 257): ")
		if !ok {
			return
		}
		protocol := 3
		fmt.Print("DNSKEY Protocol (default 3): ")
		protocolInput, _ := reader.ReadString('\n')
		protocolInput = strings.TrimSpace(protocolInput)
		if protocolInput != "" {
			value, err := strconv.Atoi(protocolInput)
			if err != nil {
				fmt.Println("Invalid protocol value.")
				return
			}
			protocol = value
		}
		publicKey := readOptional(reader, "DNSKEY Public Key (base64): ")
		if publicKey == "" {
			fmt.Println("Public key is required to compute the digest.")
			return
		}
		var err error
		digest, keyTag, err = computeDnsSecFromDNSKEY(domain, flags, protocol, algorithm, publicKey, digestType)
		if err != nil {
			fmt.Printf("Error computing DNSSEC values: %v\n", err)
			return
		}
		fmt.Printf("Computed Digest: %s\n", digest)
		fmt.Printf("Computed Key Tag: %d\n", keyTag)
	} else {
		algorithm, ok = readInt(reader, "Algorithm: ")
		if !ok {
			return
		}
		fmt.Print("Digest: ")
		digest, _ = reader.ReadString('\n')
		digest = strings.TrimSpace(digest)
		if digest == "" {
			fmt.Println("Digest is required.")
			return
		}
		digestType, ok = readInt(reader, "Digest Type: ")
		if !ok {
			return
		}
		keyTag, ok = readInt(reader, "Key Tag: ")
		if !ok {
			return
		}
	}

	var maxSigLife *int
	fmt.Print("MaxSigLife (optional, blank to skip): ")
	maxSigLifeInput, _ := reader.ReadString('\n')
	maxSigLifeInput = strings.TrimSpace(maxSigLifeInput)
	if maxSigLifeInput != "" {
		value, err := strconv.Atoi(maxSigLifeInput)
		if err != nil {
			fmt.Println("Invalid MaxSigLife value.")
			return
		}
		maxSigLife = &value
	}

	resp, err := client.AddDnsSec(domainmanagement.AddDnsSecRequest{
		Domain:     domain,
		MaxSigLife: maxSigLife,
		Algorithm:  algorithm,
		Digest:     digest,
		DigestType: digestType,
		KeyTag:     keyTag,
	})
	if err != nil {
		fmt.Printf("Error adding DNSSEC record: %v\n", err)
		return
	}

	fmt.Printf("Response Code: %d\n", resp.ResponseCode)
	fmt.Printf("Response Message: %s\n", resp.ResponseMessage)
}

func sendDnsSec(client *domainmanagement.Client, domain enomapi.Domain, digestType int, digest string, keyTag int, algorithm int, reader *bufio.Reader) {
	var maxSigLife *int
	fmt.Print("MaxSigLife (optional, blank to skip): ")
	maxSigLifeInput, _ := reader.ReadString('\n')
	maxSigLifeInput = strings.TrimSpace(maxSigLifeInput)
	if maxSigLifeInput != "" {
		value, err := strconv.Atoi(maxSigLifeInput)
		if err != nil {
			fmt.Println("Invalid MaxSigLife value.")
			return
		}
		maxSigLife = &value
	}

	resp, err := client.AddDnsSec(domainmanagement.AddDnsSecRequest{
		Domain:     domain,
		MaxSigLife: maxSigLife,
		Algorithm:  algorithm,
		Digest:     digest,
		DigestType: digestType,
		KeyTag:     keyTag,
	})
	if err != nil {
		fmt.Printf("Error adding DNSSEC record: %v\n", err)
		return
	}

	fmt.Printf("Response Code: %d\n", resp.ResponseCode)
	fmt.Printf("Response Message: %s\n", resp.ResponseMessage)
}

func deleteDnsSec(client *domainmanagement.Client, reader *bufio.Reader) {
	domain, ok := readDomain(reader)
	if !ok {
		return
	}

	printDnsSecFieldOptions()
	algorithm, ok := readInt(reader, "Algorithm: ")
	if !ok {
		return
	}
	fmt.Print("Digest: ")
	digest, _ := reader.ReadString('\n')
	digest = strings.TrimSpace(digest)
	if digest == "" {
		fmt.Println("Digest is required.")
		return
	}
	digestType, ok := readInt(reader, "Digest Type: ")
	if !ok {
		return
	}
	keyTag, ok := readInt(reader, "Key Tag: ")
	if !ok {
		return
	}

	resp, err := client.DeleteDnsSec(domainmanagement.DeleteDnsSecRequest{
		Domain:     domain,
		Algorithm:  algorithm,
		Digest:     digest,
		DigestType: digestType,
		KeyTag:     keyTag,
	})
	if err != nil {
		fmt.Printf("Error deleting DNSSEC record: %v\n", err)
		return
	}

	fmt.Printf("Response Code: %d\n", resp.ResponseCode)
	fmt.Printf("Response Message: %s\n", resp.ResponseMessage)
}

func readDomain(reader *bufio.Reader) (enomapi.Domain, bool) {
	fmt.Print("Enter domain name: ")
	dname, _ := reader.ReadString('\n')
	dname = strings.TrimSpace(dname)
	if dname == "" {
		fmt.Println("Domain name is required.")
		return enomapi.Domain{}, false
	}
	return enomapi.NewDomain(dname), true
}

func readInt(reader *bufio.Reader, prompt string) (int, bool) {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	parsed, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Invalid integer value: %s\n", value)
		return 0, false
	}
	return parsed, true
}

func readExtendedAttributes(reader *bufio.Reader, tld string, attrs []response.ExtAttribute) (map[string]string, bool) {
	fmt.Printf("Extended attributes for .%s:\n", tld)
	values := make(map[string]string)
	for _, attr := range attrs {
		attrName := strings.TrimSpace(attr.Name)
		requiredLabel := "optional"
		if attr.Required == 1 {
			requiredLabel = "required"
		} else if attr.Required == 2 {
			requiredLabel = "conditional"
		}
		userDefinedLabel := "system"
		if attr.UserDefined {
			userDefinedLabel = "user"
		}
		fmt.Printf("  %s (%s, %s)\n", attrName, requiredLabel, userDefinedLabel)
		if attr.Description != "" {
			fmt.Printf("    %s\n", attr.Description)
		}
		if len(attr.Options) > 0 {
			fmt.Println("    Options:")
			for _, option := range attr.Options {
				value := strings.TrimSpace(option.Value)
				title := strings.TrimSpace(option.Title)
				line := fmt.Sprintf("      %s", value)
				if title != "" {
					line += " - " + title
				}
				fmt.Println(line)
				if option.Description != "" {
					fmt.Printf("        %s\n", option.Description)
				}
			}
		}

		for {
			prompt := "    Enter value"
			if len(attr.Options) > 0 {
				prompt = "    Enter option value"
			}
			prompt += " for " + attrName + " (blank to skip): "
			fmt.Print(prompt)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "" {
				if attr.Required == 1 {
					fmt.Printf("    %s is required.\n", attrName)
					continue
				}
				break
			}
			if len(attr.Options) > 0 {
				mapped, ok := mapExtAttributeOption(attr.Options, input)
				if !ok {
					fmt.Printf("    Invalid option for %s. Use one of the listed values.\n", attrName)
					continue
				}
				values[attrName] = mapped
				break
			}
			values[attrName] = input
			break
		}
	}
	return values, true
}

func mapExtAttributeOption(options []response.ExtAttributeOption, input string) (string, bool) {
	input = strings.TrimSpace(input)
	for _, option := range options {
		value := strings.TrimSpace(option.Value)
		title := strings.TrimSpace(option.Title)
		if strings.EqualFold(value, input) || strings.EqualFold(title, input) {
			return value, true
		}
	}
	return "", false
}

func useEnomDNS(reader *bufio.Reader) bool {
	fmt.Print("Use eNom DNS? [y/n]: ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	return strings.EqualFold(answer, "y")
}

func readNameServers(reader *bufio.Reader, minServers int, maxServers int) []string {
	fmt.Print("Enter name servers (comma-separated): ")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	if value == "" {
		fmt.Println("At least one name server is required.")
		return nil
	}

	parts := strings.Split(value, ",")
	nameServers := make([]string, 0, len(parts))
	for _, part := range parts {
		ns := strings.TrimSpace(part)
		if ns != "" {
			nameServers = append(nameServers, ns)
		}
	}
	if len(nameServers) == 0 {
		fmt.Println("At least one name server is required.")
		return nil
	}
	if minServers > 0 && len(nameServers) < minServers {
		fmt.Printf("At least %d name servers are required.\n", minServers)
		return nil
	}
	if maxServers > 0 && len(nameServers) > maxServers {
		fmt.Printf("No more than %d name servers are allowed.\n", maxServers)
		return nil
	}
	return nameServers
}

func formatDomain(domain enomapi.Domain) string {
	if domain.Extension == "" {
		return domain.Name
	}
	return domain.Name + "." + domain.Extension
}

func printDnsSecFieldOptions() {
	fmt.Println("DNSSEC field options:")
	fmt.Println("  Algorithm: 3 (DSA/SHA-1), 5 (RSA/SHA-1), 7 (RSASHA1-NSEC3-SHA1), 8 (RSA/SHA-256), 10 (RSA/SHA-512), 12 (GOST R 34.10-2001), 13 (ECDSA/SHA-256), 14 (ECDSA/SHA-384)")
	fmt.Println("  Digest: hex-encoded digest from the DS record")
	fmt.Println("  DigestType: 1 (SHA-1), 2 (SHA-256)")
	fmt.Println("  Tip: generate keys now or paste a DNSKEY record to auto-compute Digest and KeyTag.")
}

func wantComputeDefaultYes(reader *bufio.Reader, prompt string) bool {
	fmt.Print(prompt)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	if answer == "" {
		return true
	}
	return strings.EqualFold(answer, "y")
}

func readOptional(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func computeDnsSecFromDNSKEY(domain enomapi.Domain, flags int, protocol int, algorithm int, publicKey string, digestType int) (string, int, error) {
	if digestType != 1 && digestType != 2 {
		return "", 0, fmt.Errorf("unsupported digest type: %d", digestType)
	}
	if protocol < 0 || protocol > 255 {
		return "", 0, fmt.Errorf("invalid protocol: %d", protocol)
	}
	if flags < 0 || flags > 65535 {
		return "", 0, fmt.Errorf("invalid flags: %d", flags)
	}
	if algorithm < 0 || algorithm > 255 {
		return "", 0, fmt.Errorf("invalid algorithm: %d", algorithm)
	}

	nameWire, err := dnsNameToWire(formatDomain(domain))
	if err != nil {
		return "", 0, err
	}
	publicKey = strings.Map(func(r rune) rune {
		if r == ' ' || r == '\n' || r == '\r' || r == '\t' {
			return -1
		}
		return r
	}, publicKey)
	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", 0, fmt.Errorf("invalid public key base64: %w", err)
	}

	rdata := make([]byte, 0, 4+len(keyBytes))
	rdata = append(rdata, byte(flags>>8), byte(flags))
	rdata = append(rdata, byte(protocol))
	rdata = append(rdata, byte(algorithm))
	rdata = append(rdata, keyBytes...)

	keyTag := computeKeyTag(rdata)
	payload := append(append([]byte{}, nameWire...), rdata...)

	switch digestType {
	case 1:
		sum := sha1.Sum(payload)
		return fmt.Sprintf("%x", sum[:]), keyTag, nil
	case 2:
		sum := sha256.Sum256(payload)
		return fmt.Sprintf("%x", sum[:]), keyTag, nil
	default:
		return "", 0, fmt.Errorf("unsupported digest type: %d", digestType)
	}
}

func computeKeyTag(rdata []byte) int {
	var acc uint32
	for i, b := range rdata {
		if i%2 == 0 {
			acc += uint32(b) << 8
		} else {
			acc += uint32(b)
		}
	}
	acc += (acc >> 16) & 0xFFFF
	return int(acc & 0xFFFF)
}

func parseDNSKEYRecord(input string) (int, int, int, string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, 0, 0, "", fmt.Errorf("empty record")
	}
	if idx := strings.Index(input, ";"); idx >= 0 {
		input = strings.TrimSpace(input[:idx])
	}
	fields := strings.Fields(input)
	if len(fields) < 4 {
		return 0, 0, 0, "", fmt.Errorf("expected at least 4 fields")
	}
	dnskeyIndex := -1
	for i, field := range fields {
		if strings.EqualFold(field, "DNSKEY") {
			dnskeyIndex = i
			break
		}
	}
	if dnskeyIndex >= 0 {
		fields = fields[dnskeyIndex+1:]
	}
	if len(fields) < 4 {
		return 0, 0, 0, "", fmt.Errorf("missing DNSKEY fields")
	}
	flags, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, 0, 0, "", fmt.Errorf("invalid flags: %w", err)
	}
	protocol, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, 0, 0, "", fmt.Errorf("invalid protocol: %w", err)
	}
	algorithm, err := strconv.Atoi(fields[2])
	if err != nil {
		return 0, 0, 0, "", fmt.Errorf("invalid algorithm: %w", err)
	}
	publicKey := strings.Join(fields[3:], "")
	if publicKey == "" {
		return 0, 0, 0, "", fmt.Errorf("missing public key")
	}
	return flags, protocol, algorithm, publicKey, nil
}

func generateRSADNSKEY() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}
	publicKey, err := encodeRSAPublicKeyDNSKEY(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateBytes})
	return string(privatePEM), publicKey, nil
}

func encodeRSAPublicKeyDNSKEY(publicKey *rsa.PublicKey) (string, error) {
	exponentBytes := bigEndianBytes(publicKey.E)
	var exponentLen []byte
	if len(exponentBytes) < 256 {
		exponentLen = []byte{byte(len(exponentBytes))}
	} else {
		exponentLen = []byte{0, byte(len(exponentBytes) >> 8), byte(len(exponentBytes))}
	}
	modulusBytes := publicKey.N.Bytes()
	key := append(append(exponentLen, exponentBytes...), modulusBytes...)
	return base64.StdEncoding.EncodeToString(key), nil
}

func bigEndianBytes(value int) []byte {
	if value == 0 {
		return []byte{0}
	}
	var bytes []byte
	for value > 0 {
		bytes = append([]byte{byte(value)}, bytes...)
		value >>= 8
	}
	return bytes
}

func dnsNameToWire(name string) ([]byte, error) {
	name = strings.TrimSuffix(strings.ToLower(strings.TrimSpace(name)), ".")
	if name == "" {
		return []byte{0}, nil
	}
	labels := strings.Split(name, ".")
	wire := make([]byte, 0, len(name)+2)
	for _, label := range labels {
		if label == "" {
			return nil, fmt.Errorf("invalid domain name: %s", name)
		}
		if len(label) > 63 {
			return nil, fmt.Errorf("label too long: %s", label)
		}
		wire = append(wire, byte(len(label)))
		wire = append(wire, label...)
	}
	wire = append(wire, 0)
	return wire, nil
}
