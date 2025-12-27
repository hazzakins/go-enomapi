SetPassword
===========

Set a password for a domain name.

Usage
-----

Use this command to set a password for access to a domain name.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain name password must use ASCII characters, and must be 6 to 60 characters in length.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetPassword&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&DomainPassword=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | --------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | SetPassword |
| uid | string | Required | Your Account ID                                                 |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                          |
| TLD       | string | Required | Top-level domain name \(extension\) |
| DomainPassword | string | Required | Password to access and manage the domain name. Must use ASCII characters and be 6 to 60 characters in length. |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=setpassword&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&domainpassword=resellerdocspw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=setpassword&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&domainpassword=resellerdocspw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=setpassword&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&domainpassword=resellerdocspw&ResponseType=text
```
```
<interface-response>
 <Command>SETPASSWORD</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>0.344</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/11/2011 11:35:47 PM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Command: </ STRONG>SETPASSWORD<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.078<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>804faac0-3607-4062-bd72-6e4828ae1a5e<br>
<STRONG>RequestDateTime: </ STRONG>2/5/2015 3:13:44 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command=SETPASSWORD
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
TrackingKey=31bf1a25-e9b5-4b9c-aeeb-7d176b27f01b
RequestDateTime=2/5/2015 3:14:04 PM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

[GetDomainExp](../docs/domains/renewal/GetDomainExp.md)

[GetDomainInfo](../docs/domains/domain-management/general/getdomaininfo.md)

GetDomains

GetExtendInfo

GetPasswordBit

GetRegistrationStatus

[SetRegLock](../docs/domains/domain-management/general/SetRegLock.md)

[SetRenew](../docs/domains/renewal/SetRenew.md)

StatusDomain

[ValidatePassword](../docs/domains/domain-management/general/ValidatePassword.md)