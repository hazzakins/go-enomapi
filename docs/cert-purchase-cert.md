CertPurchaseCert
================

Submit certificate request

Usage
-----

Send a cert configuration to Certificate Authority for final approval and issuance of the cert.

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

| Parameter   | Type | Status  | Description |
| ------------- | ------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command    | string | Required | CertPurchaseCert |
| UID | string | Required | Your Account ID.                                                                                                                                                                                                            |
| PW      | string | Required | Your API Token. |
| CertID | string | Required | ID number for this individual certificate. Retrieve this number using the [CertGetCerts](../docs/cert-get-certs.md) command.                                                                                                                                                     |
| ApproverEmail | string | Required | *Not required for Symantec Encryption Everywhere \(EE\)* Email address of the registrant of record for the domain to be associated with this cert. Use the [CertParseCSR](../docs/cert-parse-csr.md) command to retrieve the domain name, and then use the [CertGetApproverEmail](../docs/cert-get-approver-email.md) command to retrieve the registrant’s email address \(the “approver”\) from the authoritative Whois database. |
| DVAuthMethod | string | Optional | Authentication method for Domain Validation. Permitted values: - File - DNS - Email \(default\) Supported products: - RapidSSL - QuickSSL - QuickSSL Premium                                                                                                                                      |
| ResponseType | string | Optional | Format of response. Permitted values: - Text \(default\) - HTML - XML |

Additional parameters for Comodo Certificates and Unified Communications Certificate \(UCC\)
--------------------------------------------------------------------------------------------

| Parameter | Type  | Status | Description                                   |
| --------------------------------- | ------- | -------- | -------------------------------------------------------------------------------- |
| CSROrganization | string | Required | Name of the organization that will use this cert.                |
| CSROrganizationUnit        | string | Required | Unit within the organization that will use this cert. |
| CSRAddress1 | string | Required | Address, line 1, of the organization that will use this cert.          |
| CSRAddress2            | string | Required | Address, line 2. |
| CSRLocality | string | Required | Location \(often a city\) of the organization.                 |
| CSRStateProvince         | string | Required | State or province in which the organization is located. |
| CSRPostalCode | string | Required | Postal code of the organization.                        |
| CSRCountry            | string | Required | Country code in which the organization is located. |
| DUNSNumber | string | Optional | D-U-N-S Number. Speeds validation of Secure, Secure Plus, and all EV certs. |
| DomainListNumber         | integer | Required | - \[UCC only\]\* Number of domains to be added to the list. |
| UCCDomainList*X* \(*X=1 to 100*\) | string | Required | - \[UCC only\]\* Domain name.                          |
| UCCEmailList*X* \(*X=1 to 100*\) | string | Required | - \[UCC only\]\* Approver email. |

Symantec Encryption Everywhere upgrade path
-------------------------------------------

> ### This feature is currently under development

This cross-platform upgrade is only available for upgrading Symantec Encryption Everywhere certificate to other certificates available through the same Certificate Authority.

> ### \ You can upgrade certificates at any time. If the upgrade occurs outside the renewal window, the upgrade is not discounted.\ Wildcard certificates can only be upgraded to other wildcard products.

| Parameter | Type  | Status | Description                                                                         |
| ------------- | ------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Upgrade | boolean | Optional | Upgrade flag. Permitted values: - False \(default\) - True                                                 |
| OriginalRefID | integer | Optional | *Required if Upgrade is set to True* The Reference Order ID for the currently active Symantec Encryption Everywhere certificate that you intend to upgrade. |

Symantec Encryption Everywhere error codes
------------------------------------------

> ### Error format:
>
>
>
> - **\[key\]** \(field\): **\[reason\]**
>
>
>
> Sample:
>
>
>
> -

DNS keys:
- DNS\_INVALID\_ENTRY
- DNS\_MISMATCH
- DNS\_ENTRY\_MISSING
- DNS\_INVALID\_DOMAIN
- DNS\_INTERNAL\_ERROR

FILE keys:
- FILE\_SERVER\_NOT\_REACHABLE
- FILE\_NOT\_FOUND
- FILE\_INVALID\_FORMAT
- FILE\_INCORRECT\_CONTENT
- FILE\_OUTDATED\_CONTENT

DATA keys:
- INTERNAL\_DATA\_CHECK\_FAILED
- WEAK\_KEY
- KEY\_CHECK\_FAILED
- CAA\_CHECK\_FAILED

An internal data check failed error occurs when an order does not pass Symantec validation and security checks. These security checks are in place to prevent the issuance of certificates to embargoed countries, entities denied by the U.S. government, and sites that are vulnerable to phishing activities.

The field value shows the area of your request that needs to be updated:
- domain.cn: the requested common name.
- domain.sans: any of the requested SANs.
- contacts: any of the order contacts.
- csr: CSR.
- org: organization or admin.

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output | Type  | Description |
| ------------------ | ------- | --------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed. |
| Success | boolean | Has this cert been configured successfully? |
| CertID | integer | Identification number of this individual cert. |
| GeoTrustOrderID | integer | Order ID for Symantec, GeoTrust and Verisign product. |
| DVAuthMethod | string | Authentication method for Domain Validation. |
| ApproverEmail | string | Approver email. |
| FileAuthName | string | Authentication file name. |
| FileAuthContents | string | Authentication file contents. |
| DNSAuthEntry | string | DNS authentication entry. |
| DNSAuthAddress | string | DNS authentication address or timestamp. |
| ValidityPeriodDays | string | - \[EE only\]\* Validity period of this certificate \(days\). |
| SealInfo | string | - \[EE only\]\* Certificate Seal information \(URL\). |
| CertType | string | - \[EE only\]\* Certificate type \(ICA1 or ICA2\). |
| CertSerialNum | string | - \[EE only\]\* Will always contain an empty value. Node is present only to maintain backwards compatibility. |
| SSLCertificate | string | - \[EE only\]\* SSL Certificate |
| ErrorCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | string | True value indicates this entire response has reached you successfully. |
| TotalRecords | integer | Total number of record lines returned in the query. |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.asp?command=certpurchasecert&uid=resellid&pw=resellpw&responsetype=xml&certid=48455&[email protected]
```
```
<interface-response>
<CertPurchaseCert>
  <Success>True</Success>
  <GeoTrustOrderID>14228176</GeoTrustOrderID>
  <ApproverEmail>[email protected]</ApproverEmail>
  <DVAuthMethod></DVAuthMethod>
  <FileAuthName></FileAuthName>
  <FileAuthContents></FileAuthContents>
  <DNSAuthEntry></DNSAuthEntry>
  <DNSAuthAddress></DNSAuthAddress>
  <CertID>48455</CertID>
</CertPurchaseCert>
<Command>CERTPURCHASECERT</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod></MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL0VWAPI10</Server>
<Site>eNom</Site>
<IsLockable></IsLockable>
<IsRealTimeTLD></IsRealTimeTLD>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>3bdfd6fd-c335-4c62-883a-611b99862c52</TrackingKey>
<RequestDateTime>12/16/2016 10:39:20 AM</RequestDateTime>
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertPurchaseCert>
  <CertID>726342</CertID>
  <GeoTrustOrderID>1003094224</GeoTrustOrderID>
  <ValidityPeriodDays>365</ValidityPeriodDays>
  <SealInfo></SealInfo>
  <CertType>ICA1</CertType>
  <CertSerialNum>1877f66568a782bba9fb42dae44e33d7</CertSerialNum>
  <SSLCertificate>-----BEGIN CERTIFICATE-----
  MIIFhjCCBG6gAwIBAgIQanYV3Az0tNipCShR8yacPTANBgkqhkiG9w0BAQsFADCB
  lDELMAkGA1UEBhMCVVMxHTAbBgNVBAoTFFN5bWFudGVjIENvcnBvcmF0aW9uMR8w
  HQYDDU9efdZTeW1hbnRlYyBUcnVzdCBOZXR3b3JrMR0wGwYDVQQLExREb21haW4g
  VmFsaWRhdGVkIFNTTDEmMCQGA1UEAxMdU3ltYW50ZWMgQmFzaWMgRFYgU1NMIENB
  IC0gRzIwHJnx0TYxMjE2MDAwMDA12WhcNMTcxMjE2MjM1OTU5WjAXMRUwEwYDVQQ
  DAxidWRpdGVzdC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDL
  AM9EPebiixlELf8DklTrBCPuYFIvmu8D9ZSgW4Tjx1Vukig2ZD4QMx5ngch8n3VV
  BTN5Y3rqfxc5DcJ+ncUTL0Lb/oeqoPgRUcs4gfVNTF/bZg/pybBpg38eDG/BpEzw
  7vwmnlYRqRPBWeGrMEHHZLYH/5tS3l8WbZU0VuOpUsHvOlRsplraV9Xq7SWhqX0t
  93kaQgwh6BZp0Mzwl2DbYxwT5+OPWtNN4xafweFyXvWo8tRCGN7KNU+a9amb+hn/
  oFdFT4fp/ufoJoP8wy1s2e62CNb9+XSW47CQI/bCNZy/o7dqlv5hW1xm+lNfGSQV
  KM9HO/S7CFV9mffajwHLAgMBAAGjggJOMIICSjApBgNVHREEIjAgggxidWRpdGVz
  dC5jb22CEHd3dy5idWRpdGVzdC5jb20wCQYDV31sBAIwADBhBgNVHSAEWjBYMFYG
  BmeBDAECATBMMCMGCCsGAQUFBwIBFhdodHRwczovL2Quc3ltY2IuY29tL2NwczAl
  BggrBgEFBQcCAjAZDBdodHRwczovL2Quc3ltY2auY29tL3JwYTAfBgNVHSMEGDAW
  gBTKrF3hkC/x74zUnzUB4QE7oM7BdzAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYw
  FAYIKwYBBQUHAwEGCCsGAQUFBwMCMFcGCCsGAQUFBwEBBEswSTAfBggrBgEFBQcw
  AYYTaHR0cDovL2hkLnN5bWNkLmNvbTAmBggrBgEFBQcwAoYaaHR0cDovL2hkLnN5
  bWNiLmNvbS9oZC5jcnQwggEEBgorBgEEAdZ5AgQCBIH1BIHyAPAAdQDd6x0reg1P
  piCLga2BaHB+Lo6dAdVciI09EcTNtuy+zAAAAVkFHex2AAAEAwBGMEQCIHd2pZZd
  QQVRW07UlyMzr7Elvz17e6Vc8FNxKggVMgMlAiAMSjCrRDTE1glW41zpClBK2VyO
  HCojEoBN42wMRXKpDgB3AO5Lvbd1zmC64UJpH6vhnmajD35fsHLYgwDEe4l6qP3L
  AAABWQUd7KoAAAQDAEgwRgIhAPrBnIdACxLQJv9i+Nq6WIz2zvJQIhJqfiNfsqWc
  /6TOAiEAr9HqFR/zitgIw6FhSff4c5i+RnYz6MAfjhmKp6lOnDMwDQYJKoZIhvcN
  AQELBQADggEBADY3uXm+WsJ2Z6hW/zUCAMUqYU2rrcObGePZonnhGoQ67/Wzhy9G
  CNQVK//JvAP9WgQKJDoS0d7X2S+VKI1Y03kBlfHt5FXlxsOQbPb4aqUIj8wJ7Cnj
  /v8QM1DSz1a8UgO71cXrQGjzJ8LbLtWU3V9ZKizGc7Z2kk9zvYrnNJIJ3CWdb4ZR
  9Eh3+gmZ1KqxDUmJweKXdJchgS247S2GlYT0w/C9iNTCq7JCw/NJUHnwJIJOJcLy
  taWsu92n3npeDyFLHsM3cWFcCJiGuZuJpE7eL/C30AHyrwFFfYrpJol4WDTUsAR7
  ExnI1ee6g050Zv8uqzMsE5ma9rMEOPhEaQM=
  -----END CERTIFICATE-----</SSLCertificate>
</CertPurchaseCert>
<Command>CERTPURCHASECERT</Command>
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
<TrackingKey>d2ea5226-404a-4ab1-a505-4a4d828324cd</TrackingKey>
<RequestDateTime>12/15/2016 4:52:44 PM</RequestDateTime>
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