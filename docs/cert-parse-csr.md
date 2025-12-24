CertParseCSR
============

Retrieve information from CSR

Usage
-----

Parse a Certificate Signing Request \(CSR\) to determine the domain name and other information associated with this cert.

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

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| Command     | string | Required | CertGetCertDetail |
| UID | string | Required | Your Account ID                                                               |
| PW       | string | Required | Your API Token |
| CertID | string | Required | ID number of this cert. Use the [CertGetCerts](../docs/cert-get-certs.md) command to retrieve the ID number               |
| CSR       | string | Required | Certificate Signing Request \(CSR\) generated as an input parameter for the [CertConfigureCert](../docs/cert-configure-cert.md) command. |
| ResponseType | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT".   |

Returned Parameters and Values
------------------------------

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | CertParseCSR.                                          |
| Success     | boolean | Success status of this parsing operation.                            |
| Organization   | string | Organization name embedded in the CSR.                             |
| OrganizationUnit | string | Organizational unit embedded in the CSR                             |
| DomainName    | string | Domain name embedded in the CSR.                                |
| Email      | string | Email address embedded in the CSR.                               |
| Locality     | string | Locality \(usually a city\) embedded in the CSR.                        |
| State      | string | State embedded in the CSR.                                   |
| Country     | string | Country embedded in the CSR.                                  |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | string | "True" indicates this entire response has reached you successfully.               |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.asp?command=CertParseCSR&uid=resellid&pw=resellpw&CertID=32969&csr=-----BEGIN+NEW+CERTIFICATE+REQUEST-----%0aMIIEXTCCA0UCAQAwaTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAndhMREwDwYDVQQH%0aDAhraXJrbGFuZDEVMBMGA1UECgwMcmlnaHRzaWRlLmNvMQwwCgYDVQQLDANkZXYx%0aFTATBgNVBAMMDGJ1ZGl0ZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC%0aAQoCggEBAMsAz0Q95uKLGUQt%2fwOSVOsEI%2b5gUi%2ba7wP1lKBbhOPHVW6SKDZkPhAz%0aHmeByHyfdVUFM3ljeup%2fFzkNwn6dxRMvQtv%2bh6qg%2bBFRyziB9U1MX9tmD%2bnJsGmD%0afx4Mb8GkTPDu%2fCaeVhGpE8FZ4aswQcdktgf%2fm1LeXxZtlTRW46lSwe86VGymWtpX%0a1ertJaGpfS33eRpCDCHoFmnQzPCXYNtjHBPn449a003jFp%2fB4XJe9ajy1EIY3so1%0aT5r1qZv6Gf%2bgV0VPh%2bn%2b5%2bgmg%2fzDLWzZ7rYI1v35dJbjsJAj9sI1nL%2bjt2qW%2fmFb%0aXGb6U18ZJBUoz0c79LsIVX2Z99qPAcsCAwEAAaCCAa0wGgYKKwYBBAGCNw0CAzEM%0aFgo2LjEuNzYwMS4yMEkGCSsGAQQBgjcVFDE8MDoCAQUMFktSS0RUMTk4LmNvcnAu%0aZG0ubG9jYWwMEERNXGJ1ZGkucHJhc2V0eWEMC0luZXRNZ3IuZXhlMHIGCisGAQQB%0agjcNAgIxZDBiAgEBHloATQBpAGMAcgBvAHMAbwBmAHQAIABSAFMAQQAgAFMAQwBo%0aAGEAbgBuAGUAbAAgAEMAcgB5AHAAdABvAGcAcgBhAHAAaABpAGMAIABQAHIAbwB2%0aAGkAZABlAHIDAQAwgc8GCSqGSIb3DQEJDjGBwTCBvjAOBgNVHQ8BAf8EBAMCBPAw%0aEwYDVR0lBAwwCgYIKwYBBQUHAwEweAYJKoZIhvcNAQkPBGswaTAOBggqhkiG9w0D%0aAgICAIAwDgYIKoZIhvcNAwQCAgCAMAsGCWCGSAFlAwQBKjALBglghkgBZQMEAS0w%0aCwYJYIZIAWUDBAECMAsGCWCGSAFlAwQBBTAHBgUrDgMCBzAKBggqhkiG9w0DBzAd%0aBgNVHQ4EFgQUCc2hUWGejuYucxo%2f6m4KRG6bDfEwDQYJKoZIhvcNAQEFBQADggEB%0aAIefF1PaUAQtPBAZ1zXrDugeEVUeJyNMh%2f9rsmwlV2%2fgJAVtoFhlK5cU2YjbuJ12%0ae1ojVJCztK6t9tWFXpJJkt1mfj%2bLZ6%2bWOng%2fzl7k9qt5htPZT8SUavQQ1lR%2fB4qK%0aSnm4MsvNkrQThkvqx6fal6%2fucvebjOy3OnUgK0nXR7hGZxOSpueBAp9MFPHFjXPH%0avoOUYjoQN996a7OCIYBmGEGdOG%2bQsvRzhhmxvY0Sz05C1Q9DTTFXzGHeORmZzF75%0abxU5pV5Q1y984XgDg%2bCoPJa8%2blM9P2wQSId1fGb%2bA3vDr7coT%2fFykOKw1ZdYvA5H%0abrMLq61IThbywOBiQPjmu2Y%3d%0a-----END+NEW+CERTIFICATE+REQUEST-----&ResponseType=XML
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertParseCSR>
 <Success>True</Success>
 <Organization>rightside.co</Organization>
 <DomainName>buditest.com</DomainName>
 <Email></Email>
 <HasBadExtensions>False</HasBadExtensions>
 <Locality>kirkland</Locality>
 <OrganizationUnit>dev</OrganizationUnit>
 <State>wa</State>
 <Country>US</Country>
</CertParseCSR>
<Command>CERTPARSECSR</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t1</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.797</ExecTime>
<Done>true</Done>
<TrackingKey>1cf13d96-544a-4f00-87a4-bbe3dbbff0cf</TrackingKey>
<RequestDateTime>12/16/2016 1:30:53 PM</RequestDateTime>
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