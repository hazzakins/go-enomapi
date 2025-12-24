CertConfigureCert
=================

Configure certificate information

Usage
-----

Add or update certificate information in preparation to submit to Certificate Authority.

Availability
------------

All resellers have access to this command.

Contraints
----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The certificate must belong to this account.

Input parameters
----------------

| Parameter                               | Type | Status  | Description |
| --------------------------------------------------------------------- | ------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command                                | string | Required | CertConfigureCert |
| UID | string | Required | Your Account ID.                                                                                               |
| PW                                  | string | Required | Your API Token. |
| CertID | string | Required | ID number for this individual certificate. Retrieve this number using the [CertGetCerts](../docs/cert-get-certs.md) command.                                       |
| WebServerType                             | integer | Required | Type of Web Server on which this cert will be installed. Please see [Web Server Table](#web-server-table) below to get the value. |
| CSR | string | Required | Certificate Signing Request. When creating a CSR for a ***wildcard*** certificate, be sure to submit the common name as \*.sld.tld.                                    |
| - ContactType\*FName Permitted ContactTypes: - Admin - Tech - Billing | string | Required | First Name of this contact for this cert. |
| - ContactType\*LName | string | Required | Last Name of this contact for this cert.                                                                                   |
| - ContactType\*OrgName                        | string | Required | Organization name of this contact for this cert. |
| - ContactType\*JobTitle | string | Required | Job title of this contact for this cert.                                                                                   |
| - ContactType\*Address1                        | string | Required | Address line 1 of this contact for this cert. |
| - ContactType\*Address2 | string | Required | Address line 2 of this contact for this cert.                                                                                |
| - ContactType\*City                          | string | Required | City of this contact for this cert. |
| - ContactType\*State | string | Required | State of this contact for this cert.                                                                                     |
| - ContactType\*Province                        | string | Required | Province of this contact for this cert. |
| - ContactType\*StateProvinceChoice | string | Required | Is this cert ContactType’s location a state or province. Permitted values: - S: state - P: province                                                     |
| - ContactType\*PostalCode                       | string | Required | Postal code of this contact for this cert. |
| - ContactType\*Country | string | Required | Country of this contact for this cert.                                                                                    |
| - ContactType\*Phone                         | string | Required | Phone number of this contact for this cert. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). |
| - ContactType\*PhoneExt | string | Required | Phone number extension of this contact for this cert.                                                                            |
| - ContactType\*Fax                          | string | Required | Fax number of this contact for this cert. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). |
| - ContactType\*EmailAddress | string | Required | Email address of this contact for this cert.                                                                                 |
| ResponseType                             | string | Optional | Format of response. Permitted values: - Text \(default\) - HTML - XML |

Additional parameters for Unified Communications Certificate \(UCC\)
--------------------------------------------------------------------

| Parameter | Type  | Status | Description                 |
| --------------------------------- | ------- | -------- | ------------------------------------------- |
| DomainListNumber | integer | Required | Number of domains to be added to the list. |
| UCCDomainList*X* \(*X=1 to 100*\) | string | Required | Approver emails. |

Additional parameters for Symantec Encryption Everywhere \(EE\)
---------------------------------------------------------------

> ### Private Key
>
>
>
> Private Key is created *on-the-fly* based on your domain and CSR country code input parameters. As security best practices, we **do not store** the Private Key in our system. Please keep safe the value in a secure location temporarily. It would be needed later when installing the certificate to the server.

| Parameter | Type  | Status | Description                                                     |
| -------------- | ------ | -------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| DVAuthMethod | string | Required | Authentication method for Domain Validation. Permitted values: - File - DNS                     |
| SLD      | string | Optional             | Second-level domain name. When you supply SLD and TLD, we will generate **Private Key** and **CSR** on-the-fly. |
| TLD | string | Optional | Top-level domain name.                                               |
| CSRCountryCode | string | Required if SLD and TLD are used | Two-letter country codes. |

Web Server Table
----------------

| Certificate Type | Values                                                                                                                                                                                                                                                                                                                                                                                                |
| ----------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Symantec GeoTrust RapidSSL VeriSign | 1: Apache \+ MOD SSL 2: Apache \+ Raven 3: Apache \+ SSLeay 4: C2Net Stronghold 7: IBM HTTP 8: iPlanet Server 4.1 9: Lotus Domino Go 4.6.2.51 10: Lotus Domino Go 4.6.2.6\+ 11: Lotus Domino 4.6\+ 12: Microsoft IIS 4.0 13: Microsoft IIS 5.0 14: Netscape Enterprise/FastTrack 17: Zeus v3\+ 18: Other 20: Apache \+ OpenSSL 21: Apache 2 22: Apache \+ ApacheSSL 23: Cobalt Series 24: Cpanel 25: Ensim 26: Hsphere 27: Ipswitch 28: Plesk 29: Jakarta-Tomcat 30: WebLogic \(all versions\) 31: O’Reilly WebSite Professional 32: WebStar 33: Microsoft IIS 6.0                                                                                                            |
| Comodo               | 1000: Otherold 1001: AOL 1002: Apache/ModSSL 1003: Apache-SSL \(Ben-SSL, not Strong-hold\) 1004: C2Net Strongholdold 1005: Cobalt Raq 1006: Covalent Server Software 1007: IBM HTTP Server 1008: IBM Internet Connection Server 1009: iPlanet 1010: Java Web Server \(Javasoft / Sun\) 1011: Lotus Domino 1012: Lotus Domino Go\! 1013: Microsoft IIS 1.x to 4.x 1014: Microsoft IIS 5.x and later 1015: Netscape Enterprise Server 1016: Netscape FastTrack 1017: Novell Web Server 1018: Oracle 1019: Quid Pro Quo 1020: R3 SSL Server 1021: Raven SSL 1022: RedHat Linux 1023: SAP Web Application Server 1024: Tomcat 1025: Website Professional 1026: WebStar 4.x and later 1027: Web Ten \( from Tenon\) 1028: Zeus Web Server 1029: Ensim 1030: Plesk 1031: WHM/cPanel 1032: H-Sphere |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output | Type  | Description |
| -------------------------- | ------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed. |
| CertID | integer | Certificate ID. |
| ProdType | integer | Product Type. |
| *CSRData* Success | boolean | Has this cert been configured successfully? |
| *CSRData* Organization | string | Organization embedded in this CSR. |
| *CSRData* DomainName | string | Domain name embedded in this CSR. |
| *CSRData* Email | string | Email address embedded in this CSR. |
| *CSRData* HasBadExtensions | boolean | Any bad extensions in this CSR? |
| *CSRData* Locality | string | Locality embedded in this CSR. |
| *CSRData* OrganizationUnit | string | Organization unit embedded in this CSR. |
| *CSRData* State | string | State embedded in this CSR. |
| *CSRData* Country | string | Country embedded in this CSR. |
| Approver | object | List of valid approver emails from the Certificate Authority. |
| DVAuthMethod | string | - \[EE only\]\* Authentication method for Domain Validation. |
| ApproverEmail | string | - \[EE only\]\* Approver email. |
| FileAuthName | string | - \[EE only\]\* Authentication file name. |
| FileAuthContents | string | - \[EE only\]\* Authentication file contents. |
| DNSAuthEntry | string | - \[EE only\]\* DNS authentication entry. |
| DNSAuthType | string | - \[EE only\]\* DNS record type. |
| DNSAuthAddress | string | - \[EE only\]\* DNS authentication value. |
| DNSAuthHostUpdate | boolean | - \[EE only\]\* Is the host record for the domain successfully updated? If the domain is located in the account and using our DNS, our system updates it automatically. |
| CSR | string | - \[EE only\]\* Certificate Signing Request auto-generated by our system if SLD and TLD are the input parameters. |
| RSA | string | - \[EE only\]\* Private Key auto-generated by our system if SLD and TLD are the input parameters. Important: please see [Private Key Information](#symantec-encryption-everywhere). |
| ErrorCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | string | True value indicates this entire response has reached you successfully. |
| TotalRecords | integer | Total number of record lines returned in the query. |

Example Input / Output
----------------------

```
http://resellertest.enom.com/interface.asp?command=certconfigurecert&uid=resellid&pw=resellpw&responsetype=xml&certid=48455&AdminOrgName=Rightside&AdminJobTitle=Dev&AdminFName=John&AdminLName=Smith&AdminAddress1=1st%20Ave&AdminCity=Kirkland&AdminCountry=US&AdminPostalCode=98033&[email protected]&AdminState=WA&AdminPhone=+1.4252744500&TechFName=John&TechLName=Smith&TechAddress1=1st%20Ave&TechCity=Kirkland&TechCountry=US&TechPostalCode=98033&[email protected]&TechState=WA&TechPhone=+1.4252744500&BillingFName=John&BillingLName=Smith&BillingAddress1=1st%20Ave&BillingCity=Kirkland&BillingCountry=US&BillingPostalCode=98033&[email protected]&BillingState=WA&BillingPhone=+1.4252744500&csr=-----BEGIN+CERTIFICATE+REQUEST-----%0aMIICuDCCAaACADB0MQswCQYDVQQGEwJVUzEYMBYGA1UEAxMPcmVzZWxsZXJkb2Mu%0aY29tMREwDwYDVQQHEwhraXJrbGFuZDEVMBMGA1UEChMMcmlnaHRzaWRlLmNvMRMw%0aEQYDVQQIEwp3YXNoaW5ndG9uMQwwCgYDVQQLEwNkZXYwggEiMA0GCSqGSIb3DQEB%0aAQUAA4IBDwAwggEKAoIBAQDDWVb3y3pb8%2fxOewZHfYamfu%2bhoEmHNCaJ0s4O29Yv%0aa%2bwPhsqCiwed7uCjxYweMp8hCna9aCq9LlRDBRHaCLUyof%2f7PxxoEe5a83X5sHnZ%0a674hZOefhxMu7hBcgmnHRqCvWsVI4pqZ3y5sN70E8cadRJC3H%2fO5DZNllcXtn7bx%0a92sixlgW0oSDE0hs8VczH3g458QydCuBDi6aMvvkvVHlNi19Y1I6xqJuuP1Utvoz%0a%2bpusHd6IEH5gsWfvaMdnvDqrmWInolG4IT9uxhnRnTYo4xNJQ3fb3LjCkKrGlR4x%0a4o9wUPgmXir2xaRj%2fUjyqS4729Bqzp6jK6A6cTEECdwrAgMBAAGgADANBgkqhkiG%0a9w0BAQsFAAOCAQEAvL8hxzsvi%2ftYG%2bP7XRQqA2E914hdqHHoejhDxov1meoGS4a0%0a8D88Tz%2bRtyGqQfB9zCQlY5oH82s6Gsn%2bonKSF0%2fJs%2bucooN7dhJbSTuD9O665DUd%0azHA7r75DuZgdYB5u19UjYs1BtEhJ8lCzf3uHiTYImG6a%2fi3wZzYxC7vo7c4HzMW3%0adpYUF%2b2z1XRIdhKj4M1wqUB3mgMm%2fyBy62oioe6CUrclkRI8Ir4eZYtQJtPe9OGI%0aMdMC27phTJztRc8vGlbidUewsq7HO67%2ftBDbvPTs77ZcRH2b5fXVCrtfErwknTA4%0aul7NNXl4bWEcjkreVz%2bAoRxrtahGRUt57A2lfg%3d%3d%0a-----END+CERTIFICATE+REQUEST-----
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertConfigureCert>
  <CertID>725899</CertID>
  <ProdType>23</ProdType>
  <CSRData>
   <Success>True</Success>
   <Organization>rightside.co</Organization>
   <DomainName>resellerdocs.com</DomainName>
   <Email></Email>
   <HasBadExtensions>false</HasBadExtensions>
   <Locality>kirkland</Locality>
   <OrganizationUnit>dev</OrganizationUnit>
   <State>wa</State>
   <Country>US</Country>
  </CSRData>
  <ApproverEmail></ApproverEmail>
  <Success>True</Success>
  <Approver Type="Domain">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Domain">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Manual"></Approver>
</CertConfigureCert>
<Command>CERTCONFIGURECERT</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod></MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL0VWAPI05</Server>
<Site>eNom</Site>
<IsLockable></IsLockable>
<IsRealTimeTLD></IsRealTimeTLD>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>a25ce569-cc05-42aa-abf9-d1ae780d5588</TrackingKey>
<RequestDateTime>3/1/2017 1:02:10 PM</RequestDateTime>
```
```
<?xml version="1.0" encoding="utf-8"?>
 <interface-response>
 <CertConfigureCert>
  <CertID>726216</CertID>
  <ProdType>286</ProdType>
  <DVAuthMethod>DNS</DVAuthMethod>
  <ApproverEmail></ApproverEmail>
  <FileAuthName></FileAuthName>
  <FileAuthContents></FileAuthContents>
  <DNSAuthEntry>@</DNSAuthEntry>
  <DNSAuthType>TXT</DNSAuthType>
  <DNSAuthAddress>
   20170303210335hi8lp217ob29a2f4vh4joryuhiy3iv8qymsrvb2312922kcfe
   </DNSAuthAddress>
  <DNSAuthHostUpdate>True</DNSAuthHostUpdate>
  <CSR></CSR>
  <RSA></RSA>
 </CertConfigureCert>
 <Command>CERTCONFIGURECERT</Command>
 <APIType>API</APIType>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod></MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL0VWAPI11</Server>
 <Site>eNom</Site>
 <IsLockable></IsLockable>
 <IsRealTimeTLD></IsRealTimeTLD>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.000</ExecTime>
 <Done>true</Done>
 <TrackingKey>bd758da9-b211-4996-bd06-763820437277</TrackingKey>
 <RequestDateTime>3/1/2017 1:03:36 PM</RequestDateTime>
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertConfigureCert>
  <CertID>726216</CertID>
  <ProdType>286</ProdType>
  <DVAuthMethod>FILE</DVAuthMethod>
  <ApproverEmail></ApproverEmail>
  <FileAuthName>
   http://resellerdocs.com/.well-known/pki-validation/fileauth.txt
  </FileAuthName>
  <FileAuthContents>
   2017030321073946ncbo2k1r341ig30ud0h2jq3hyxt8hs22cvh1gox7x1yl2u5o
  </FileAuthContents>
  <DNSAuthEntry></DNSAuthEntry>
  <DNSAuthAddress></DNSAuthAddress>
  <DNSAuthHostUpdate></DNSAuthHostUpdate>
  <CSR></CSR>
  <RSA></RSA>
</CertConfigureCert>
<Command>CERTCONFIGURECERT</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod></MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>KRKDT198</Server>
<Site>eNom</Site>
<IsLockable></IsLockable>
<IsRealTimeTLD></IsRealTimeTLD>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>0227b6e0-3da6-4768-9950-acdede93556c</TrackingKey>
<RequestDateTime>3/1/2017 1:07:40 PM</RequestDateTime>
```
Related Commands
----------------

[CertChangeApproverEmail](../docs/cert-change-approveremail.md)

[CertConfigureCert](../docs/cert-configure-cert.md)

[CertGetApproverEmail](../docs/cert-get-approver-email.md)

[CertGetCertDetail](../docs/cert-get-certdetail.md)

[CertGetCerts](../docs/cert-get-certs.md)

[CertModifyOrder](../docs/cert-modify-order.md)

[CertParseCSR](../docs/cert-parse-csr.md)

[CertPurchaseCert](../docs/cert-purchase-cert.md)

[CertResendApproverEmail](../docs/cert-resend-approveremail.md)