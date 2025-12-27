GetDomainExp
============

Get domain expiration date.

Usage
-----

Use this command to get a domain's expiration date.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=getdomainexp&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetDomainExp |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| sld | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| tld       | string | Required | Top-level domain name |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| ExpirationDate  | string | Expiration date for the domain registration.                          |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetDomainExp&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainExp&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainExp&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
```
<interface-response>
 <ExpirationDate>6/10/2019 3:56:56 PM</ExpirationDate>
 <Command>GETDOMAINEXP</Command>
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
 <ExecTime>0.095</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/8/2011 4:06:42 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>ExpirationDate: </STRONG>6/10/2019 3:56:56 PM<br />
<STRONG>Command: </STRONG>GETDOMAINEXP<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+08.00<br />
<STRONG>ExecTime: </STRONG>0.125<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/3/2015 5:37:31 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
ExpirationDate=6/10/2019 3:56:56 PM
Command=GETDOMAINEXP
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.078
Done=true
RequestDateTime=2/3/2015 5:37:55 PM
```
Related Commands
----------------

Extend

Extend\_RGP

GetAllDomains

GetDomainCount

GetDomainInfo

GetDomains

GetExtendInfo

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

InsertNewOrder

SetPassword

SetRegLock

SetRenew

StatusDomain

UpdateExpiredDomains

ValidatePassword