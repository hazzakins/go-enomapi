CertReissueCert
===============

Reissue an existing Symantec, Verisign or GeoTrust certificate.

Usage
-----

Use this command to reissue an existing Symantec, Verisign or GeoTrust certificate.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/secure/configure-ssl-certificate.aspx?certid=295](https://resellertest.enom.com/secure/configure-ssl-certificate.aspx?certid=295)

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The cert must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                            | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                | 4 |
| CSR | Required | Certificate Signing Request \(CSR\) generated as an input parameter for the CertConfigureCert command. | 4500   |
| CertID     | Required | ID number of this cert. To retrieve this number, use the CertGetCerts command.            | 8 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| Success | Success status of this parsing operation                             |
| Organization   | Organization name embedded in the CSR |
| DomainName | Domain name embedded in the CSR                                 |
| Email      | Email address embedded in the CSR |
| Locality | Locality \(usually a city\) embedded in the CSR                         |
| OrganizationUnit | Organizational unit embedded in the CSR |
| State | State embedded in the CSR                                    |
| Country     | Country embedded in the CSR |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query parses a CSR and sends the response in XML format:

```
https://resellertest.enom.com/interface.asp?
command=CertReissueCert&uid=resellid&pw=resellpw
&ResponseType=XML&CertID=111&CSR=%2D%2D%2D%2D%2D
BEGIN+NEW+CERTIFICATE+REQUEST%2D%2D%2D%2D%2D%0D%0AMI
IDVzCCAsACAQAwfDELMAkGA1UEBhMCVVMxCzAJBgNVBA
gTAkRFMRMwEQYDVQQH%0D%0AEwpXaWxtaW5ndG9uMRIw
EAYDVQQKEwlBY21lIEluYy4xGDAWBgNVBAsTD1NTTCBD%
0D%0AZXJ0aWZpY2F0ZTEdMBsGA1UEAxMUd3d3LnJlc2Vs
bGVyZG9jcy5jb20wgZ8wDQYJ%0D%0AKoZIhvcNAQEBBQA
DgY0AMIGJAoGBAL0aEkBD2RkKOm82yJGy%2FIhLRH7pYIG
ZCazh%0D%0ACxp731BjemXjSaVJLu0%2BoE6sI
MRhV04X%2FJjVFgGGbswVvHT5qWZdpODe2EEBG9
71%0D%0ABJlUmBGEX%2Flgkd%2BjjHbP3MnOqbkuRRYrd
ao2CPjB3dcv46IcjFvCl5P%2BSEVx7Y9c
%0D%0Awjk3n%2FxLAgMBAAGgggGZMBoGCisGAQQBgjc
NAgMxDBYKNS4yLjM3OTAuMjB7Bgor%0D%0ABgEEAYI3Ag
EOMW0wazAOBgNVHQ8BAf8EBAMCBPAwRAYJKoZIhvcNA
QkPBDcwNTAO%0D%0ABggqhkiG9w0DAgICAIAwDgYIKoZ
IhvcNAwQCAgCAMAcGBSsOAwIHMAoGCCqGSIb3%0D%0AD
QMHMBMGA1UdJQQMMAoGCCsGAQUFBwMBMIH9BgorBgEEA
YI3DQICMYHuMIHrAgEB%0D%0AHloATQBpAGMAcgBvAHMA
bwBmAHQAIABSAFMAQQAgAFMAQwBoAGEAbgBuAGUAbAAg
%0D%0AAEMAcgB5AHAAdABvAGcAcgBhAHAAaABpAGMAIAB
QAHIAbwB2AGkAZABlAHIDgYkA%0D%0AAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAA>A%0D%0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA%0D%0AAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
DANBgkqhkiG%0D%0A9w0B9w0BAQUFAAOBgQCfN86YfjwT
GCMj84DFkvQBPX0zQ8815bgqaZSNHQRYl67Gj4A%2B%0D
%0Awmg3O18lX0NwBt%2BT%2F57OaJS4HB6x6JvAo96N
%2B1vX%2F%2BiG2thcu1PqTb%2Fu%2BmUbapaa%0D
%0Ackas6Ubbe2MVHKRz7o0ZNfIKUrRBdCRtZBEV
xz1ZrPMqRHVo1oq5N17SEA%3D%3D%0D%0A%2D%2D
%2D%2D%2DEND+NEW+CERTIFICATE+REQUEST%2D%2D
%2D%2D%2D
```
```
https://resellertest.enom.com/interface.asp?
command=CertReissueCert&uid=resellid&pw=resellpw
&ResponseType=html&CertID=111&CSR=%2D%2D%2D%2D%2D
BEGIN+NEW+CERTIFICATE+REQUEST%2D%2D%2D%2D%2D%0D%0AMI
IDVzCCAsACAQAwfDELMAkGA1UEBhMCVVMxCzAJBgNVBA
gTAkRFMRMwEQYDVQQH%0D%0AEwpXaWxtaW5ndG9uMRIw
EAYDVQQKEwlBY21lIEluYy4xGDAWBgNVBAsTD1NTTCBD%
0D%0AZXJ0aWZpY2F0ZTEdMBsGA1UEAxMUd3d3LnJlc2Vs
bGVyZG9jcy5jb20wgZ8wDQYJ%0D%0AKoZIhvcNAQEBBQA
DgY0AMIGJAoGBAL0aEkBD2RkKOm82yJGy%2FIhLRH7pYIG
ZCazh%0D%0ACxp731BjemXjSaVJLu0%2BoE6sI
MRhV04X%2FJjVFgGGbswVvHT5qWZdpODe2EEBG9
71%0D%0ABJlUmBGEX%2Flgkd%2BjjHbP3MnOqbkuRRYrd
ao2CPjB3dcv46IcjFvCl5P%2BSEVx7Y9c
%0D%0Awjk3n%2FxLAgMBAAGgggGZMBoGCisGAQQBgjc
NAgMxDBYKNS4yLjM3OTAuMjB7Bgor%0D%0ABgEEAYI3Ag
EOMW0wazAOBgNVHQ8BAf8EBAMCBPAwRAYJKoZIhvcNA
QkPBDcwNTAO%0D%0ABggqhkiG9w0DAgICAIAwDgYIKoZ
IhvcNAwQCAgCAMAcGBSsOAwIHMAoGCCqGSIb3%0D%0AD
QMHMBMGA1UdJQQMMAoGCCsGAQUFBwMBMIH9BgorBgEEA
YI3DQICMYHuMIHrAgEB%0D%0AHloATQBpAGMAcgBvAHMA
bwBmAHQAIABSAFMAQQAgAFMAQwBoAGEAbgBuAGUAbAAg
%0D%0AAEMAcgB5AHAAdABvAGcAcgBhAHAAaABpAGMAIAB
QAHIAbwB2AGkAZABlAHIDgYkA%0D%0AAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAA>A%0D%0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA%0D%0AAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
DANBgkqhkiG%0D%0A9w0B9w0BAQUFAAOBgQCfN86YfjwT
GCMj84DFkvQBPX0zQ8815bgqaZSNHQRYl67Gj4A%2B%0D
%0Awmg3O18lX0NwBt%2BT%2F57OaJS4HB6x6JvAo96N
%2B1vX%2F%2BiG2thcu1PqTb%2Fu%2BmUbapaa%0D
%0Ackas6Ubbe2MVHKRz7o0ZNfIKUrRBdCRtZBEV
xz1ZrPMqRHVo1oq5N17SEA%3D%3D%0D%0A%2D%2D
%2D%2D%2DEND+NEW+CERTIFICATE+REQUEST%2D%2D
%2D%2D%2D
```
```
https://resellertest.enom.com/interface.asp?
command=CertReissueCert&uid=resellid&pw=resellpw
&ResponseType=text&CertID=111&CSR=%2D%2D%2D%2D%2D
BEGIN+NEW+CERTIFICATE+REQUEST%2D%2D%2D%2D%2D%0D%0AMI
IDVzCCAsACAQAwfDELMAkGA1UEBhMCVVMxCzAJBgNVBA
gTAkRFMRMwEQYDVQQH%0D%0AEwpXaWxtaW5ndG9uMRIw
EAYDVQQKEwlBY21lIEluYy4xGDAWBgNVBAsTD1NTTCBD%
0D%0AZXJ0aWZpY2F0ZTEdMBsGA1UEAxMUd3d3LnJlc2Vs
bGVyZG9jcy5jb20wgZ8wDQYJ%0D%0AKoZIhvcNAQEBBQA
DgY0AMIGJAoGBAL0aEkBD2RkKOm82yJGy%2FIhLRH7pYIG
ZCazh%0D%0ACxp731BjemXjSaVJLu0%2BoE6sI
MRhV04X%2FJjVFgGGbswVvHT5qWZdpODe2EEBG9
71%0D%0ABJlUmBGEX%2Flgkd%2BjjHbP3MnOqbkuRRYrd
ao2CPjB3dcv46IcjFvCl5P%2BSEVx7Y9c
%0D%0Awjk3n%2FxLAgMBAAGgggGZMBoGCisGAQQBgjc
NAgMxDBYKNS4yLjM3OTAuMjB7Bgor%0D%0ABgEEAYI3Ag
EOMW0wazAOBgNVHQ8BAf8EBAMCBPAwRAYJKoZIhvcNA
QkPBDcwNTAO%0D%0ABggqhkiG9w0DAgICAIAwDgYIKoZ
IhvcNAwQCAgCAMAcGBSsOAwIHMAoGCCqGSIb3%0D%0AD
QMHMBMGA1UdJQQMMAoGCCsGAQUFBwMBMIH9BgorBgEEA
YI3DQICMYHuMIHrAgEB%0D%0AHloATQBpAGMAcgBvAHMA
bwBmAHQAIABSAFMAQQAgAFMAQwBoAGEAbgBuAGUAbAAg
%0D%0AAEMAcgB5AHAAdABvAGcAcgBhAHAAaABpAGMAIAB
QAHIAbwB2AGkAZABlAHIDgYkA%0D%0AAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAA>A%0D%0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA%0D%0AAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
DANBgkqhkiG%0D%0A9w0B9w0BAQUFAAOBgQCfN86YfjwT
GCMj84DFkvQBPX0zQ8815bgqaZSNHQRYl67Gj4A%2B%0D
%0Awmg3O18lX0NwBt%2BT%2F57OaJS4HB6x6JvAo96N
%2B1vX%2F%2BiG2thcu1PqTb%2Fu%2BmUbapaa%0D
%0Ackas6Ubbe2MVHKRz7o0ZNfIKUrRBdCRtZBEV
xz1ZrPMqRHVo1oq5N17SEA%3D%3D%0D%0A%2D%2D
%2D%2D%2DEND+NEW+CERTIFICATE+REQUEST%2D%2D
%2D%2D%2D
```
In the response, parsed information and an ErrCount value 0 confirm that the query was successful:

```
<?xml version="1.0" ?>

 <Success>True</Success>

 <Message />

 <Command>CERTREISSUECERT</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.578125</ExecTime>

 <Done>true</Done>

 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>

 <RequestDateTime>1/11/2012 11:18:42 AM</RequestDateTime>

</interface-response>
```