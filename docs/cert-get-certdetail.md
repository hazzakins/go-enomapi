CertGetCertDetail
=================

Retrieve configuration information on one SSL certificate.

Usage
-----

Use this command to retrieve detailed configuration information on one SSL certificate.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The cert must belong to this account.

Input Parameters
----------------

| Input Parameter   | Type | Status  | Description |
| ------------------- | ------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command       | string | Required | CertGetCertDetail |
| UID | string | Required | Your Account ID                                                                       |
| PW         | string | Required | Your API Token |
| CertID | string | Required | ID number of this cert. Use the [CertGetCerts](../docs/cert-get-certs.md) command to retrieve the ID number                       |
| IncludeIntermediate | boolean | Optional | Return intermediate certificate \(ICA\). Default: false Note: only for *Symantec Encryption Everywhere* product and the cert status must be ***issued***. |
| ResponseType | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT".           |

Returned Parameters and Values
------------------------------

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter        | Type | Description                                                                                                                             |
| ------------------------------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CertID             | integer | ID number of this individual cert.                                                                                                                 |
| DomainName           | string | Domain this cert is associated with.                                                                                                                |
| ValidityPeriod         | integer | Total validity duration of this cert, in months.                                                                                                          |
| OrderDate           | datetime | Date on which this cert was ordered.                                                                                                                |
| ConfigDate           | datetime | Date on which this cert was installed.                                                                                                               |
| RenewalIndicator        | boolean | Renewal indicator setting for this cert.                                                                                                              |
| OrderID            | integer | ID number of the order that included this cert                                                                                                           |
| OrderDetailID         | integer | Item ID number of this cert when it was ordered.                                                                                                          |
| ExpirationDate         | datetime | Expiration date of this cert.                                                                                                                   |
| ServerCount          | integer | Should always be 1.                                                                                                                        |
| DVAuthMethod          | string | Domain validation method.                                                                                                                     |
| AuthStatusID          | int | Domain Validation status ID \(GeoTrust File and DNS only\).                                                                                                    |
| AuthStatusName         | string | Domain Validation status name. \(GeoTrust File and DNS only\). Possible values: - UNREACHABLE - INIT - AUTHENTICATED - FILE\_NOT\_FOUND - FILE\_MISMATCH - UNAVAILABLE Note: more values may be added in the future to support more products and various stages. |
| AuthStatusLastUpdate      | string | Domain Validation most recent date the file/dns was polled \(GeoTrust File and DNS only\).                                                                                     |
| ApproverEmail         | string | Email address of the registrant.                                                                                                                  |
| Message            | string | Message.                                                                                                                              |
| FileAuthName          | string | File authentication name.                                                                                                                     |
| FileAuthContents        | string | Contents of the authentication file.                                                                                                                |
| DNSAuthEntry          | string | DNS authentication entry.                                                                                                                     |
| DNSAuthAddress         | string | DNS authentication address or timestamp.                                                                                                              |
| CertStatusID          | integer | ID number that identifies this certificate’s phase in its processing.                                                                                               |
| CertStatus           | string | Text description of this certificate’s status in its processing.                                                                                                  |
| CertStatusDetail        | string | Verbose description of status.                                                                                                                   |
| CertOrderStateID        | integer | Order state ID \(GeoTrust File and DNS only\).                                                                                                           |
| CertOrderStateName       | string | Sub-status of CertStatus, providing more granular information while the CertStatus is in "processing" \(GeoTrust File and DNS only\).                                                               |
| ReissueStatus         | string | Certificate reissue status.                                                                                                                    |
| ComodoReissueStatus      | string | Comodo UCC Only - Reissue status for this CertID.                                                                                                         |
| ProdType            | integer | Product ID number of this type of cert.                                                                                                              |
| ProdDesc            | string | Text description of this type of cert.                                                                                                               |
| UseWebHosting         | boolean | Use our hosting services? Expected values: - "Y" - Yes - "N" - No                                                                                                 |
| SSLCertificate         | string | SSL certificate, provided by the certificate authority.                                                                                                      |
| IntermediateCertificates    | object | Collection of intermediate certificates objects. Note: only for *Symantec Encryption Everywhere* product.                                                                             |
| CSR              | string | Certificate Signing Request \(CSR\) code, provided by the certificate authority.                                                                                          |
| CACerts            | string | Intermediate Certificate.                                                                                                                     |
| PVT              | string | Private Key.                                                                                                                            |
| ReferenceID          | integer | Reference ID number for this cert.                                                                                                                 |
| DomainCountAllowed       | string | Comodo UCC Only - Maximum domain slots for this CertID.                                                                                                      |
| DomainList           | string | Comodo UCC Only - List of domains for this CertID.                                                                                                        |
| ContactTypeOrgName       | string | Organization name of Admin, Tech, or Billing Contact.                                                                                                       |
| ContactTypeJobTitle      | string | Job title of Contact.                                                                                                                       |
| ContactTypeFName        | string | Given name of Contact.                                                                                                                       |
| ContactTypeLName        | string | Family name \(Surname\) of Contact.                                                                                                                |
| ContactTypeAddress1      | string | Contact address, first line.                                                                                                                    |
| ContactTypeAddress2      | string | Contact address, second line.                                                                                                                   |
| ContactTypeCity        | string | Contact city.                                                                                                                           |
| ContactTypeStateProvince    | string | Contact state or province.                                                                                                                     |
| ContactTypeStateProvinceChoice | boolean | Is this a state or a province? Expected values are: - S - State - P - Province                                                                                          |
| ContactTypePostalCode     | string | Contact postal code.                                                                                                                        |
| ContactTypeCountry       | string | Contact country.                                                                                                                          |
| ContactTypePhone        | string | Contact phone.                                                                                                                           |
| ContactTypeFax         | string | Contact FAX.                                                                                                                            |
| ContactTypeEmailAddress    | string | Contact email address.                                                                                                                       |
| ContactTypePhoneExt      | string | Contact phone extension.                                                                                                                      |
| WebServerTypeID        | integer | ID number of this type of Web server.                                                                                                               |
| WebServerTypeName       | string | Text description of this type of Web server                                                                                                             |
| WebServerTypeCode       | string | Text abbreviation of this type of Web server                                                                                                            |
| Command            | string | Name of command executed                                                                                                                      |
| ErrCount            | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                  |
| ErrX              | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                                      |
| Done              | string | "True" indicates this entire response has reached you successfully.                                                                                                |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.asp?command=CertGetCertDetail&uid=resellid&pw=resellpw&CertID=48455&ResponseType=XML
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertGetCertDetail>
 <CertID>726347</CertID>
 <DomainName>buditest.com</DomainName>
 <ValidityPeriod>12</ValidityPeriod>
 <OrderDate>12/15/2016 5:13:02 PM</OrderDate>
 <ConfigDate></ConfigDate>
 <RenewalIndicator>False</RenewalIndicator>
 <WebServerTypeID>18</WebServerTypeID>
 <WebServerTypeName>Other</WebServerTypeName>
 <OrderID>267562392</OrderID>
 <OrderDetailID>185806132</OrderDetailID>
 <ExpirationDate>12/15/2017 5:13:02 PM</ExpirationDate>
 <ServerCount>1</ServerCount>
 <DVAuthMethod>DNS</DVAuthMethod>
 <AuthStatusID />
 <AuthStatusName></AuthStatusName>
 <AuthStatusLastUpdate />
 <ApproverEmail></ApproverEmail>
 <FileAuthName></FileAuthName>
 <FileAuthContents></FileAuthContents>
 <DNSAuthEntry>s23lxkrm431e2d23d1188qij441a</DNSAuthEntry>
 <DNSAuthAddress>s20201215064552.buditest.com</DNSAuthAddress>
 <Message></Message>
 <CertStatusID>1</CertStatusID>
 <CertStatus>Awaiting Configuration</CertStatus>
 <CertStatusDetail></CertStatusDetail>
 <CertOrderStateID />
 <CertOrderStateName></CertOrderStateName>
 <ReissueStatus></ReissueStatus>
 <ComodoReissueStatus></ComodoReissueStatus>
 <ProdType>286</ProdType>
 <ProdDesc>Symantec Encryption Everywhere</ProdDesc>
 <UseWebHosting>0</UseWebHosting>
 <SSLCertificate><![CDATA[-----BEGIN CERTIFICATE-----
 MIIFpTCCBI2gAwIBAgIQV9OHOE3c6nT93ARUn2LWozANBgkq1kiG9w0BAQsFADCB
 ujELMjhGA1U11hMCVVMxHTAbBgNVBAoTFFN5bWFudGVjIENvcnBvcmF0aW9uMR8w
 HQYDVQQLExZTeW1hbnR4YyBUcnVzdCBOZXR3b3JrMR0wGwYDVQQLExREb21kaW4g
 VmFsaWRhdGVkIFNTTDEfMB0GA1UECxMWRk9SIFRFU1QgUFVSUE9TRVMgT05MWTEr
 MCkGA1UEAxMiU3ltY250ZWMgQmFza5MgRFYgU1NMIFRFU1QgQ0EgLSBHMjAeFw0x
 NzA3MTEwMDAwMDBaFw0xODA1MTEyMzU5NTlaMBUxEzARBgNVBAMMCjEwcmluZy5v
 cmc1ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDI3xAC8KmxgaoOMl5h
 pzR85kfaFjlP1TWuN//bxc+wTAogDrHZ0rNV8lYfIrwvgxSUgETzTWJ93J+8UH56
 E7amQ5J4Q0+41+oWmi7jHRI58l5KXJOPg4e8so0bffFyiUNKamjKRiWIoI028ASk
 D2IYH/xJ7UmeTYcC0NV2ND44bzK11B+FCjDRq+SxIPytwzEqJ6iPurLOZLW20ht2
 /W9woTMJ8+LDcaYmPft3zkP8Uv8GibR0UeS38kVC1NDdlRf0PpPqGaTsvXV30Fhs
 dQ8w77RpAoAlZaZUkda/+lYD3qxC6UzvPV+jidbfrTEaMNnOgXoRd8SM6zstLhDM
 qIuNAgMBd0GjggJJMIICRTAlBgNVHREEHjAcggoxMHJpbmcub3Jngg53d3cuMTBy
 aW5nLm9yZzAJBgNVHRMEAjAAMGEGA1UdIARaMFgwVgYGZ4EMAQIBMEwwIwYIKwYB
 BQUHAgEWF2h0dHBzOi8vZC5zeW1jYi5jbddvY3BzMCUGCCsGAQUFBwICMBkMF2h0
 dHBzOi8vZC5zeW1jYi5jb20vcnBhMB2GA1UdIwQYMBaAFLTe37IA2qjiNuc5lL2E
 7cHcnk0DMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYB
 BQUHAwIwVwYIKwYBBQUHAQEESzBJMB8GCCsGAQUFBzABhhNodHRwOi8vaGQuc3lt
 Y2QuY29tMCYGCCsGAQUFBzAChhpodHRwOi8vaGQuc3ltY2IuY29tL2hkLmNydDCC
 AQMGCisGAQQB1nkCBAIEgfQEgfEA7wB1ABHTC53hEpYTtWlcb5q7FCU3D17DdBZh
 4o7YYq/iMTC5AAABW/lH8t0AAAQDAEYwRAIgCG+ddm4rDVLy7CymJ0940j8zkGz3
 LMcRdMS1PwT4EgkCIH9GOemff9PWiZ73oajb78Pfgvl1tEYjeT5fqO3jm29KAHYA
 kS5/jl018u97P1ZbmLmAuVcmlVMU4hZGC9fsU6pe7MEAAAFb+Ufy4QAABAMARzBF
 AiEAkE6wdKtHhfLf+2ojJb1LoZYH2C4rcbA6TGZma+6laVgCICysKwrB5qFMjbFs
 dTZu/UY/WBD4TdVryw8/HI66YgASMA0GCSqGSIb3DQEBCwUAA4IBAQCMt9FK9iW4
 3aMNqXxOBE1+wyqhnxB4O8LzZUYwHAK9FdeWAu7mkmcuuTRcgxpocTMC7nZlaAQC
 Glz3kZteacGCIXkBgZDxMokZNZMBB1EhjJqnxpbTo1AFtAuamaxOLsRa3DLmdZKP
 f9EHDCW3MzK1/Avz+WSh9equv6MaqCi80H3msLCPsvZxpl7Bx2i5O+PI5afQ29fC
 hmFeSx8zpfGvs/R+6WqD6IrjoVWCCOmMdX06P2/kvSNeCiKV0MIKkW/4ajTgfEvH
 +dD2zcQPU90jN5Vm1WdtYo4nKdcaForDnHSyYrUT8RkSRL5gkDHX0aXTRQOt33Ub
 ze2wexcYmUnc
 -----END CERTIFICATE-----]]>
 </SSLCertificate>
 <IntermediateCertificates>
   <ICA1><![CDATA[-----BEGIN CERTIFICATE-----
   MIIFdDCCBFygAwIBAgIQT/WKykuy/mrzO7vFuzo6ADANBhkqhkiG9w0BAQsFADCB
   6zELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL
   ExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMUIwQAYaVQQLEzlUZXJtcyBvZiB1c2Ug
   YXQgaHR0cHM6Ly93d3sudmVyaXNpZ24uY29tL2Nwcy90ZXN0Y2EgKGMpMDgxHzAd
   BgNVBAsTFkZvciBUZXN0IFB1cnBvc2VzIE9ubHkxPTA7BgNVBAMTNFZlcmlTaWdu
   IFVuaXZlcnNhbCBUZXN0IFJvb3QgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcN
   MfYwNjA5MDAwMDAwWhcNMjYwNjA4MjM1OTU5WjCBujELMAkGA1UEBhMCVVMxHTAb
   BgNVBAoTFFN5bWFudGVjIENvcnBvcmF0aW9uMR8wHQYDVQQLExZTeW1hbnRlYyBU
   cnVzdCBOZXR3b3JrMR0wGwYDVQQLE221b21haW4gVmFsaWRhdGVkIFNTTDEfMB0G
   A1UECxMWRk9SIFRFU1QgUFVSUE9TRVMgT05MWTErMCkGA1UEAxMiU3ltYW50ZWMg
   QmFzaWMgRFYgU1NMIFRFU1QgQ0EgLSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEP
   ADCCAQoCggEBANTqZ2lJQ2edQsIyKi0KMitwaaNWtdGgw3yuVNJ+h7bAji9lpU3H
   yphz1W8kRKwi0u5AgoFtykxNxgFjcwycoWqGqCSCIT1xUgYCu5A8Ro+qn7IPw4sa
   5crnFJShdm93pgL8AaB/y8SB/fKG8FbV97vMa6V9xov/yXlsHYDRn93QILLkXCFY
   G9olWZ3RkpELQ312m9TxALKrqojsgVg42Rx6sAZl/mU4dNwH4y4NOrOP8EexqO/Z
   Jl3E69LhpqlWZNVtHg90e4QYmsKjyXNk6+ix36m/3anJbm0nFN4yVA6UkEmb1IVs
   erOTmzDc8VD/KGf1I+fuGk7QxUgIez4eYIUCAwEAAaOCAUEwggE9MBIGA1UdEwEB
   /wQIMAYBAf8CAQAwYQYDVR0gBFowWDBWBgZngQwBAgEwTDAjBggrBgEFBQcCARYX
   aHR0cHM6Ly9kLnN5bWNiLmNvbS9jcHMwJQYIKwYBBQUHAgIwGRoX4HR0cHM6Ly9k
   LnN5bWNiLmNvbS9ycGEwVQYDVR0fBE4wTDBKoEigRoZEaHR0cDovL3BpbG90b25z
   aXRlY3JsLnZlcmlzaWduLmNvbS9PZmZsaW5lQ0EvdW5pdmVyc2FsLXRlc3Qtcm9v
   dC5jcmwwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMA4GA1UdDwEB/wQE
   AwIBBjAdBgNVH14EFgQUtN7fsgDaqOI29zmUvYTtwdyeTQMwHwYDVR0jBBgwFoAU
   6vogFmton8gv0R0u2lWUbe+FOG8wDQYJKoZIhvcNAQELBQADggEBAHMbeExLK/xz
   zvXhNWe3r/FkgxbOompz1cg4qJUO0VgI/19OF+JzycsGbKqKj/IMsjTDWbIBF0OA
   b8jC8TGySAsoQqK+v/MkotB3esEN/9IPrRlpNQLuvVbRJHzwioCkN4VRmUuKe1/3
   IOn4Pdgp1FOHP7qJ4oeZWUl6ITqOO9oN3X0nSLeKLs8m/r1wQUIsSw9gC9Zc+S1s
   lVbsmLXT7OYjGB51w+yQisFTmXIbE6bwizPIVknFsqLevoYdB4xR8kzbjfM2Gcf5
   YqB5SFp8WIAF2aJcIx4IEjowMcbRyxO7QZzwIXWya0gx9+NkOJcimSz/ZDw1TfZQ
   CbAcWNeCnA4=
   -----END CERTIFICATE-----]]></ICA1>
 </IntermediateCertificates>
 <CSR><![CDATA[-----BEGIN NEW CERTIFICATE REQUEST-----
  MIIEXTCCA0UCAQAwaTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAndhMREwDwYDVQQH
  DAhraXJrbGFuZDEVMBMGA1UECgwMcmlnaHRzaWRlLmNvMQwwCgYDVQQLDANkZXYx
  FTATBgNVBAMMDGJ1ZGl0ZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
  AQoCggEBAMsAz0Q95uKLGUQt/wOSVOsEI+5gUi+a7wP1lKB1hOPHVW6SKDZkPhAz
  HmeByHyfdVUFM3ljeup/FzkNwn6dxRMvQtv+h6qg+BFRyziB9U1MX9tmD+nJsGmD
  fx4Mb8GkTPDu/CaeVhGpE8FZ4aswQcdktgf/m1LeXxZtlTRW46lSwe86VGymWtpX
  1ertJaGpfS33eRpCDCHoFmnQzPCXYNtjHBPn449a003jFp/B4XJe9ajy1EIY3so1
  T5r1qZv6Gf+gV0VPh+n+5+gmg/zDLWzZ7rYI1v35dJbjsJAj9sI1nL+jt2qW/mFb
  XGb6U18ZJBUoz0c79LsIVX2Z99qPAdsCAwEAAaCCAa0wGgYKKwYBBAGCNw0CAzEM
  Fgo2LjEuNzYwMS4yMEkGCSsGAQQBgjcVFDE8MDoCAQUMFktSS0RUMTk4LmNvcnAu
  ZG0ubG9jYWwMEERNXGJ1ZGkucHJhc2V0eWEMC0luZXRNZ3IuZXhlMHIGCisGAQQB
  gjcNAgIxZDBiAgEBHloATQBpAGMAcgBvAHMAbwBmAHQAIABSAFMAQQAgAFMAQwBo
  AGEAbgBuAGUAbAAgAEMAcgB5AHAAdABvAGcAcgBhAHAAaABpAGMAIABQAHIAbwB2
  AGkAZABlAHIDAQAwgc8GCSqGSIb3DQEJDjGBwTCBvjAOBgNVHQ8BAf8EBAMCBPAw
  EwYDVR0lBAwwCgYIKwYBBQUHAwEweAYJKoZIhvcNAQkPBGswaTAOBggqhkiG9w0D
  AgICAIAwDgYIKoZIhvcNAwQCAgCAMAsGCWCGSAFlAwQBKjALBglghkgBZQMEAS0w
  CwYJYIZIAWUDBAECMAsGCWCGSAFlAwQBBTAHBgUrDgMCBzAKBggqhkiG9w0DBzAd
  BgNVHQ4EFgQUCc2hUWGejuYucxo/6m4KRG6bDfEwDQYJKoZIhvcNAQEFBQADggEB
  AIefF1PaUAQtPBAZ1zXrDugeEVUeJyNMh/9rsmwlV2/gJAVvoFhlK5cU2YjbuJ12
  e1ojVJCztK6t9tWFXpJJkt1mfj+LZ6+WOng/zl7k9qt5htPZT8SUavQQ1lR/B4qK
  Snm4MsvNkrQThkvqx6fal6/ucvebjOy3OnUgK0nXR7hGZxOSpueBAp9MFPHFjXPH
  voOUYjoQN996a7OCIYBmGEGdOG+QsvRzhhmxvY0Sz05C1Q9DTTFXzGHeORmZzF75
  bxU5pV5Q1y984XgDg+CoPJa8+lM9P2wQSId1fGb+A3vDr7coT/FykOKw1ZdYvA5H
  brMLq61IThbywOBiQPjmu2Y=
  -----END NEW CERTIFICATE REQUEST-----]]>
 </CSR>
 <CACerts><![CDATA[]]></CACerts>
 <PVT><![CDATA[-----BEGIN RSA PRIVATE KEY-----
  -----END RSA PRIVATE KEY-----]]></PVT>
 <ReferenceID></ReferenceID>
 <DomainCountAllowed></DomainCountAllowed>
 <DomainList />
 <AdminOrgName><![CDATA[Rightside]]></AdminOrgName>
 <AdminJobTitle><![CDATA[Dev]]></AdminJobTitle>
 <AdminFName><![CDATA[John]]></AdminFName>
 <AdminLName><![CDATA[Smith]]></AdminLName>
 <AdminAddress1><![CDATA[1st Ave]]></AdminAddress1>
 <AdminAddress2><![CDATA[]]></AdminAddress2>
 <AdminCity><![CDATA[Kirkland]]></AdminCity>
 <AdminStateProvince><![CDATA[WA]]></AdminStateProvince>
 <AdminStateProvinceChoice><![CDATA[S]]></AdminStateProvinceChoice>
 <AdminPostalCode><![CDATA[98033]]></AdminPostalCode>
 <AdminCountry><![CDATA[US]]></AdminCountry>
 <AdminPhone><![CDATA[+1.4252744500]]></AdminPhone>
 <AdminFax><![CDATA[]]></AdminFax>
 <AdminEmailAddress><![CDATA[[email protected]]]></AdminEmailAddress>
 <AdminPhoneExt><![CDATA[]]></AdminPhoneExt>
 <TechOrgName><![CDATA[]]></TechOrgName>
 <TechJobTitle><![CDATA[]]></TechJobTitle>
 <TechFName><![CDATA[John]]></TechFName>
 <TechLName><![CDATA[Smith]]></TechLName>
 <TechAddress1><![CDATA[1st Ave]]></TechAddress1>
 <TechAddress2><![CDATA[]]></TechAddress2>
 <TechCity><![CDATA[Kirkland]]></TechCity>
 <TechStateProvince><![CDATA[WA]]></TechStateProvince>
 <TechStateProvinceChoice><![CDATA[S]]></TechStateProvinceChoice>
 <TechPostalCode><![CDATA[98033]]></TechPostalCode>
 <TechCountry><![CDATA[US]]></TechCountry>
 <TechPhone><![CDATA[+1.4252744500]]></TechPhone>
 <TechFax><![CDATA[]]></TechFax>
 <TechEmailAddress><![CDATA[[email protected]]]]></TechEmailAddress>
 <TechPhoneExt><![CDATA[]]></TechPhoneExt>
 <BillingOrgName><![CDATA[]]></BillingOrgName>
 <BillingJobTitle><![CDATA[]]></BillingJobTitle>
 <BillingFName><![CDATA[John]]></BillingFName>
 <BillingLName><![CDATA[Smith]]></BillingLName>
 <BillingAddress1><![CDATA[1st Ave]]></BillingAddress1>
 <BillingAddress2><![CDATA[]]></BillingAddress2>
 <BillingCity><![CDATA[Kirkland]]></BillingCity>
 <BillingStateProvince><![CDATA[WA]]></BillingStateProvince>
 <BillingStateProvinceChoice><![CDATA[S]]></BillingStateProvinceChoice>
 <BillingPostalCode><![CDATA[98033]]></BillingPostalCode>
 <BillingCountry><![CDATA[US]]></BillingCountry>
 <BillingPhone><![CDATA[+1.4252744500]]></BillingPhone>
 <BillingFax><![CDATA[]]></BillingFax>
 <BillingEmailAddress><![CDATA[[email protected]]]]></BillingEmailAddress>
 <BillingPhoneExt><![CDATA[]]></BillingPhoneExt>
 <WebServerTypes>
  <WebServerType>
   <WebServerTypeID>22</WebServerTypeID>
   <WebServerTypeName>Apache + ApacheSSL</WebServerTypeName>
   <WebServerTypeCode>apacheapachessl</WebServerTypeCode>
  </WebServerType>
  ...... <cut> ......
  <WebServerType>
   <WebServerTypeID>17</WebServerTypeID>
   <WebServerTypeName>Zeus v3+</WebServerTypeName>
   <WebServerTypeCode>zeusv3</WebServerTypeCode>
  </WebServerType>
 </WebServerTypes>
</CertGetCertDetail>
<Command>CERTGETCERTDETAIL</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl0vwapi15</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.031</ExecTime>
<Done>true</Done>
<TrackingKey>ed8648a0-6407-4762-9dfc-2e26cc44371c</TrackingKey>
<RequestDateTime>12/16/2016 1:17:27 PM</RequestDateTime>
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