Purchase
========

Purchase domain name

Usage
-----

Use this command to purchase a **domain name** or **premium domain** in real time, or put in a real time order into the pre-registration queue for an EAP domain.

The **[Purchase](../docs/domains/registration/purchase.md)** command is typically used for a single-name purchase and returns an immediate success/failure response. Commonly is used by resellers who maintain their own databases.

> ### GDPR
>
>
>
> The European Union’s General Data Protection Regulation \(GDPR\) lays out a new set of rules for how the personal data of people living within the EU \(“EU-local individuals”\) should be handled. The policy comes into full effect on **May 25, 2018**, and we recommend that you start preparing now by speaking with a lawyer and familiarizing yourself with the information we’ve provided here.
>
>
>
> Please see [GDPR Guide](../docs/gdpr.md) for detailed information and the output sample below.

> ### Premium Domains
>
>
>
> Because of the high prices for some Premium Domains, we offer you a wider variety of payment methods than for most purchases:
>
>
>
> - You can pay immediately using your account balance or credit card, just as you do for non-premium domains \(but note that we place a maximum on credit card transactions\).
> - For Premium Domains, you can also pay by wire transfer, or you can pay later using your account balance. Use the *UseWireTransfer* parameter in the Purchase command to defer payment. When you’re ready to pay, either wire your payment, or use the **[NM\_ProcessOrder](../docs/nm-processorder.md)** command to pay with your account balance. Just be aware that we will not release the Premium Domain to your account until we receive payment.

> ### .eu and .be
>
>
>
> When you register a .eu or .be domain name, we recommend that you always provide Registrant contact information that is separate from Billing contact information; don’t use the “same as Billing” default.
>
>
>
> Tip: If the Billing and Registrant contact information are the same, we recommend changing the use or spelling of abbreviations in the street address to help our system recognize that it needs to create multiple contacts. Separating the Registrant and Billing information makes it easier to update Registrant contact information in the future.

Availability
------------

All resellers can sell domain names that are available at the Registry, but only our direct ETP \(*Enom Technology Partner*\) resellers can sell Premium Domains.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- To use our credit card processing, this must be a reseller account directly under eNom, and must have signed a credit card agreement with us.
- The domain name to be purchased must be valid \(see requirements under the **[Check](../docs/domains/availability/check.md)** command\).
- eNom must be licensed to sell the names you attempt to register.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=Purchase&uid=YourAccountID&pw=YourApiToken&EndUserIP={Required}&SLD={Required}&TLD={Required}&responsetype=xml
```
| Input Parameter        | Type | Status                                                                           | Description |
| ----------------------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command            | string | Required                                                                          | Purchase |
| UID | string | Required | Your Account ID                                                                                                                                                                                                                                                                                      |
| PW              | string | Required                                                                          | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                                                                                                                                                                                |
| TLD              | string | Required                                                                          | Top-level domain name \(extension\) |
| IDNCode | string | Optional. Required for PUNY-encoded names that use characters other than the English alphabet, Arabic numbers, and hyphen. | International Domain Name code for each language used by a domain in the cart that has status Active. To retrieve the Active domains in the cart, use command [GetCartContent](../docs/getcartcontent.md) To retrieve the IDN Codes for a particular TLD, use command [GetIDNCodes](../docs/domains/tld/get-idn-codes.md) For a list of IDN TLDs, go to [http://www.enom.com/domains/idn-domains](http://www.enom.com/domains/idn-domains)                                                                          |
| CustomerSuppliedPrice     | integer | Optional. Required for purchasing a **Premium Domain**                                                  | The exact price for this domain to acknowledge this domain comes at a price higher than most registrations. Use the [PE\_GetPremiumPricing](../docs/pe-getpremiumpricing.md) command to retrieve the price. |
| PremiumDomain | boolean | Optional. Required for purchasing a **Premium Domain via NM\_\* commands** | To acknowledge this domain comes at a price higher than most registrations                                                                                                                                                                                                                                                        |
| UseWireTransfer        | boolean | Optional overall, but Required if you want to defer payment for a Premium Domain                                      | If ***UseWireTransfer=Yes*** , process this Premium Domain order as a deferred payment order. Use this parameter if you want to use either of our pay-later options: wire transfer or account balance. If you pay by wire transfer, contact your sales representative for wire transfer instructions. If you pay using your account balance \(but after submitting the order using the [Purchase](../docs/domains/registration/purchase.md) command\), use the [NM\_ProcessOrder](../docs/nm-processorder.md) command. |
| RegPeriod | integer | Optional. Required for purchasing an **EAP domain** | Which EAP Registration period you are attempting to purchase. Supported values: - pre-registration                                                                                                                                                                                                                                           |
| EapDay            | integer | Optional. Required for purchasing an **EAP domain**                                                    | Specify which EAP day you are intending to purchase. Permitted values: - 1 - 7 |
| Signed\_Mark\_Data | string | Optional. Required for purchasing **domain with a trademark** | The contents of the SMD \(*Signed Mark Data*\) file that the TMCH \(*Trade Mark Clearing House*\) has provided you.                                                                                                                                                                                                                                   |
| UseDNS            | string | Either UseDNS or NSX is Required                                                              | Specify *UseDNS=default* to use our name servers \(most of the services that we provide require our name servers\). See Note. |
| NS\(*x*\) | string | Either UseDNS or NSX is Required | Use this parameter to assign user-specified name servers. Permitted values are the use names of registered name servers, for example, *NS1=ns1.name-services.com*. See Note.                                                                                                                                                                                                    |
| UnLockRegistrar        | boolean | Optional; default value is 0                                                                | Set to 1 to unlock or 0 to lock the domain name. |
| RenewName | boolean | Optional; default value is 0 | Set to 1 to renew the name automatically before it expires.                                                                                                                                                                                                                                                               |
| DomainPassword        | string | Optional; default value is no password                                                           | Set a password on the domain name. For Premium Domains, you cannot set a password in the [Purchase](../docs/domains/registration/purchase.md) command, but you can set a password once the domain name is in your account. |
| EmailNotify | boolean | Optional; default is 0 | Set to 1 to receive email notification of customer orders, set to 0 or don't use it otherwise.                                                                                                                                                                                                                                              |
| NumYears           | integer | Optional for non-premium domains; default is the number you set with **UpdateCusPreferences**. For Premium Domains, we automatically set this value to 1 | Number of years to register the name. Permitted values: - 1 - 10 Some TLDs vary from this range. For Premium Domains, we ignore this parameter if you supply it, and automatically set NumYears to 1 |
| QueueOrder | boolean | Optional | If *QueueOrder=1*, register the domain at the Registry and return an order ID, then put this order in our order queue instead of processing in real time. This option reduces the risk of the order timing out when our system is under stress. This option is suitable for real-time TLDs like .com and .net but less relevant for manually processed TLDs like .de and .co.uk. The slight delay associated with this option makes it unsuitable if your processes immediately execute secondary calls like setting host records, they must be done after the order completes. |
| AllowQueuing         | boolean | Optional; default is 1                                                                   | If the Registry would fail this order due to stresses it is experiencing, submit *AllowQueuing=1* to queue and process this order when the Registry recovers. Submit *AllowQueuing=0* to fail the order when the Registry is under stress. |
| IgnoreNSFail | boolean | Optional; default is No | Continue processing even if name servers cannot be applied. If *IgnoreNSFail=Yes*, this purchase request will succeed even if the Registry does not recognize the name servers listed in this query. If you use *IgnoreNSFail=Yes* and failures are returned, you should confirm the status of name servers with the Registry.                                                                                                                             |
| Service            | string | Optional                                                                          | Type of service to add along with domain registration. Permitted values are: - WhoisPublicity |
| UseCreditCard | boolean | Optional | If *UseCreditCard=Yes*, use our credit card processing services. This service is available only to resellers who have entered into a credit card processing agreement with us. When you pass credit card information with this command, you must use the secure HTTPS protocol.                                                                                                                                                    |
| ChargeAmount         | decimal | Required if using our credit card processing to purchase a non- premium domain name.                                   | Amount to charge per year for the registration \(this value will be multiplied by *NumYears*to calculate the total charge to the credit card\). Required format: ***DD.cc*** For Premium Domains, we ignore this parameter if you supply it, and charge the posted price. We cover your merchant services fees for Premium Domains so that the convenience charges do not reduce your commission. |
| EndUserIP | string | Required if using our credit card processing | End user’s IP address. This is used in fraud checking, as part of our order processing service. Required format: ***NNN.NNN.NNN.NNN***.                                                                                                                                                                                                                       |
| CardType           | string | Required if using our credit card processing                                                        | Type of credit card. Permitted values: - Visa - Mastercard - AmEx - Discover |
| CCName | string | Required if using our credit card processing | Cardholder's name                                                                                                                                                                                                                                                                                     |
| CreditCardNumber       | string | Required if using our credit card processing                                                        | Credit card number |
| CreditCardExpMonth | integer | Required if using our credit card processing | Credit card expiration month. Required format: ***MM***                                                                                                                                                                                                                                                                 |
| CreditCardExpYear       | integer | Required if using our credit card processing                                                        | Credit card expiration year. Required format: ***YYYY*** |
| CVV2 | integer | Required if using our credit card processing | Credit card verification code                                                                                                                                                                                                                                                                               |
| CCAddress           | string | Required if using our credit card processing                                                        | Credit card billing address |
| CCCity | string | Required if using our credit card processing | Credit card billing city                                                                                                                                                                                                                                                                                 |
| CCStateProvince        | string | Required if using our credit card processing                                                        | Credit card billing state or province |
| CCZip | string | Required if using our credit card processing | Credit card billing postal code                                                                                                                                                                                                                                                                              |
| CCPhone            | string | Required if using our credit card processing                                                        | Credit card billing phone number. Required format: ***\+CountryCode.PhoneNumber*** - CountryCode *and*PhoneNumber \* use only numeric characters The \+ \(plus\) symbol is URLEncoded as a plus sign \(%2B\). |
| CCCountry | string | Required if using our credit card processing | Credit card billing country. The two-letter country codes.                                                                                                                                                                                                                                                               |
| ExtendedAttributes      | string | Required for some country code TLDs                                                            | Data required by the Registry for some country codes. Use **GetExtAttributes** to determine whether this TLD requires extended attributes. |
| RegistrantFirstName | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\) | Registrant first name                                                                                                                                                                                                                                                                                   |
| RegistrantLastName      | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\)                   | Registrant last name |
| RegistrantOrganizationName | string | Optional | Registrant organization                                                                                                                                                                                                                                                                                  |
| RegistrantJobTitle      | string | Required if param RegistrantOrganizationName is used                                                    | Registrant job title |
| RegistrantAddress1 | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\) | Registrant Address                                                                                                                                                                                                                                                                                    |
| RegistrantAddress2      | string | Optional                                                                          | Registrant additional address information |
| RegistrantCity | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\) | Registrant city                                                                                                                                                                                                                                                                                      |
| RegistrantStateProvinceChoice | string | Optional                                                                          | Registrant state or province. Permitted values: - S for state - P for province |
| RegistrantStateProvince | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\) | Registrant state or province name                                                                                                                                                                                                                                                                             |
| RegistrantPostalCode     | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\)                   | Registrant postal code |
| RegistrantCountry | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\) | Registrant country. The two-letter country codes.                                                                                                                                                                                                                                                                   |
| RegistrantEmailAddress    | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\)                   | Registrant email address |
| RegistrantPhone | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\) | Registrant phone. Required format: ***\+CountryCode.PhoneNumber*** - CountryCode *and*PhoneNumber \* use only numeric characters The \+ \(plus\) symbol is URLEncoded as a plus sign \(%2B\).                                                                                                                                                                                             |
| RegistrantFax         | string | Required if extended attributes are required OR if you supply one of the other core Registrant attributes \(see Note\)                   | Registrant fax number. Required format: ***\+CountryCode.PhoneNumber*** - CountryCode *and*PhoneNumber \* use only numeric characters The \+ \(plus\) symbol is URLEncoded as a plus sign \(%2B\). |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML.                                                                                                                                                                                                                                                        |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter  | Type | Description                                                |
| ------------------ | -------- | --------------------------------------------------------------------------------------------------------- |
| OrderDelayed    | boolean | Returned if this order was queued due to stresses at the Registry                     |
| ContactTypeContact | string | Returned if new contacts were created with this query string                       |
| OrderID      | double | Order number if successful. We recommend that you store this number for future use            |
| TotalCharged    | decimal | Total points or $US charged for this order                                |
| RegistrantPartyID | guid | Party ID number for registrant, for our database                             |
| RRPCode      | integer | Success code. - 200: successful processing in real time. - 1300: successful processing to queue order. |
| RRPText      | string | Text description of the processing.                                   |
| OrderStatus    | string | Order status.                                              |
| OrderDescription  | string | Verbose description of order status.                                   |
| RegistryExpDate  | datetime | Expiration date of this domain when response code is 200 and processing is in real time.         |
| RegistryCreateDate | datetime | Create date of this domain when response code is 200 and processing is in real time.           |
| Command      | string | Name of command executed                                         |
| ErrCount      | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.     |
| ErrX        | string | Error messages explaining the failure. These can be presented as is back to the client.         |
| Done        | boolean | True indicates this entire response has reached you successfully.                    |

Example Output
--------------

```
https://resellertest.enom.com/interface.asp?command=contacts&UID=YourAccountID&PW=YourApiToken&sld=resellerdocs&tld=com&ResponseType={Optional}
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <OrderID>162373942</OrderID>
  <DomainInfo>
    <RegistryCreateDate>2018-05-17 01:26:10.000</RegistryCreateDate>
    <RegistryExpDate>2019-05-17 01:26:10.000</RegistryExpDate>
  </DomainInfo>
  <TotalCharged>9</TotalCharged>
  <RegistrantPartyID></RegistrantPartyID>
  <OrderStatus>Success</OrderStatus>
  <OrderDescription>Order has completed</OrderDescription>
  <RRPCode>200</RRPCode>
  <RRPText>Command completed successfully - 162373942</RRPText>
  <Contacts>
    <Registrant>
      <OrganizationName>Data Protected</OrganizationName>
      <JobTitle>Data Protected</JobTitle>
      <FirstName>Data</FirstName>
      <LastName>Protected</LastName>
      <Address1>123 Data Protected</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.0000000000</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.0000000000</Fax>
      <EmailAddress>[email protected]</EmailAddress>
      <ConsentStatus>PENDING</ConsentStatus>
    </Registrant>
    <Admin>
      <OrganizationName>Enom</OrganizationName>
      <JobTitle></JobTitle>
      <FirstName>James</FirstName>
      <LastName>Smith</LastName>
      <Address1>2710 Clark Avenue</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.123456789</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.123456789</Fax>
      <EmailAddress>[email protected]</EmailAddress>
    </Admin>
    <Tech>
      <OrganizationName>Enom</OrganizationName>
      <JobTitle></JobTitle>
      <FirstName>James</FirstName>
      <LastName>Smith</LastName>
      <Address1>2710 Clark Avenue</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.123456789</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.123456789</Fax>
      <EmailAddress>[email protected]</EmailAddress>
    </Tech>
    <Billing>
      <OrganizationName>Enom</OrganizationName>
      <JobTitle></JobTitle>
      <FirstName>James</FirstName>
      <LastName>Smith</LastName>
      <Address1>2710 Clark Avenue</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.123456789</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.123456789</Fax>
      <EmailAddress>[email protected]</EmailAddress>
    </Billing>
  </Contacts>
  <Command>PURCHASE</Command>
  <APIType>API</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>RESELLERTEST</Server>
  <Site>eNom</Site>
  <IsLockable>True</IsLockable>
  <IsRealTimeTLD>True</IsRealTimeTLD>
  <TimeDifference>+08.00</TimeDifference>
  <ExecTime>0.000</ExecTime>
  <Done>true</Done>
  <TrackingKey></TrackingKey>
  <RequestDateTime>5/16/2018 6:26:13 PM</RequestDateTime>
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <OrderID>162373943</OrderID>
  <DomainInfo>
    <RegistryCreateDate></RegistryCreateDate>
    <RegistryExpDate></RegistryExpDate>
  </DomainInfo>
  <TotalCharged>9</TotalCharged>
  <RegistrantPartyID></RegistrantPartyID>
  <OrderStatus>Order Pending</OrderStatus>
  <OrderDescription>Order is in pending status until the Registrant Contact provides consent</OrderDescription>
  <RRPCode>200</RRPCode>
  <RRPText>Command completed successfully - 162373943</RRPText>
  <Contacts>
    <Registrant>
      <OrganizationName>Data Protected</OrganizationName>
      <JobTitle>Data Protected</JobTitle>
      <FirstName>Data</FirstName>
      <LastName>Protected</LastName>
      <Address1>123 Data Protected</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.0000000000</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.0000000000</Fax>
      <EmailAddress>[email protected]</EmailAddress>
      <ConsentStatus>PENDING</ConsentStatus>
    </Registrant>
    <Admin>
      <OrganizationName>Enom</OrganizationName>
      <JobTitle></JobTitle>
      <FirstName>James</FirstName>
      <LastName>Smith</LastName>
      <Address1>2710 Clark Avenue</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.123456789</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.123456789</Fax>
      <EmailAddress>[email protected]</EmailAddress>
    </Admin>
    <Tech>
      <OrganizationName>Enom</OrganizationName>
      <JobTitle></JobTitle>
      <FirstName>James</FirstName>
      <LastName>Smith</LastName>
      <Address1>2710 Clark Avenue</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.123456789</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.123456789</Fax>
      <EmailAddress>[email protected]</EmailAddress>
    </Tech>
    <Billing>
      <OrganizationName>Enom</OrganizationName>
      <JobTitle></JobTitle>
      <FirstName>James</FirstName>
      <LastName>Smith</LastName>
      <Address1>2710 Clark Avenue</Address1>
      <Address2></Address2>
      <City>Kirkland</City>
      <StateProvince>WA</StateProvince>
      <PostalCode>98033</PostalCode>
      <Country>US</Country>
      <Phone>+1.123456789</Phone>
      <PhoneExt></PhoneExt>
      <Fax>+1.123456789</Fax>
      <EmailAddress>[email protected]</EmailAddress>
    </Billing>
  </Contacts>
  <Command>PURCHASE</Command>
  <APIType>API</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>RESELLERTEST</Server>
  <Site>eNom</Site>
  <IsLockable>True</IsLockable>
  <IsRealTimeTLD>True</IsRealTimeTLD>
  <TimeDifference>+08.00</TimeDifference>
  <ExecTime>0.000</ExecTime>
  <Done>true</Done>
  <TrackingKey></TrackingKey>
  <RequestDateTime>5/16/2018 6:34:07 PM</RequestDateTime>
```
Notes
-----
- If using our credit card option, all registrant and credit card information is required.
- For most TLDs \(those that don’t require extended attributes\), Registrant contact information is optional. If not supplied, it is inherited from the Billing contact. However, if you supply any of the core Registrant values, you must supply them all. The core Registrant values are:
 - RegistrantFirstName
 - RegistrantLastName
 - RegistrantAddress1
 - RegistrantCity
 - PostalCode
 - RegistrantCountry \(two-letter country codes\)
 - RegistrantEmailAddress
 - RegistrantPhone
- Additional parameters for contact information can be passed by replacing ***Registrant*** in the parameter names above with ***Tech***, ***Admin*** or ***AuxBilling***.
- To set name servers to eNom’s, set the *UseDNS=default* param and don't pass NS\(x\) name servers. To set name servers to your name servers, set *NS\(x\)=YourNameServerX* and don't pass *UseDNS=default*. You can set up to 13 of your own name servers.
- Some TLDs, such as .de, requires you to specify name servers. The Registry will automatically delete the domain name if no name servers are specified.
- In the return, an *RRPCode=200* indicates a successful registration. Otherwise, check the return parameter *ErrCount*. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- In the return, if the *RRPCode=1300* and *IsRealTimeTLD=false*, this is a non real-time TLD. For these names, use the GetOrderDetail command roughly every 24 hours to check the progress of the registration.