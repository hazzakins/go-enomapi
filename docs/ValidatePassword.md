ValidatePassword
================

Validate a password for a domain name.

Usage
-----

Use this command to determine whether the password for a domain name is valid.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The domain name must exist.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetDomainServices&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&DomainPassword=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | ValidatePassword |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| sld | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| tld       | string | Required | Top-level domain name \(extension\) |
| DomainPassword | string | Required | Password to access and manage the domain name.              |
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

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=validatepassword&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&domainpassword=resellerdocs2pw
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=validatepassword&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&domainpassword=resellerdocs2pw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=validatepassword&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&domainpassword=resellerdocs2pw&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <Command>VALIDATEPASSWORD</Command>
 <ErrCount>0</ErrCount>
 <Server>Dev Workstation</Server>
 <Site>enom</Site>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
 <STRONG>Command: </ STRONG>VALIDATEPASSWORD<br>
 <STRONG>APIType: </ STRONG>API<br>
 <STRONG>Language: </ STRONG>eng<br>
 <STRONG>ErrCount: </ STRONG>0<br>
 <STRONG>ResponseCount: </ STRONG>0<br>
 <STRONG>MinPeriod: </ STRONG>1<br>
 <STRONG>MaxPeriod: </ STRONG>10<br>
 <STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
 <STRONG>Site:</ STRONG><br>
 <STRONG>IsLockable: </ STRONG>True<br>
 <STRONG>IsRealTimeTLD: </ STRONG>True<br>
 <STRONG>TimeDifference: </ STRONG>+08.00<br>
 <STRONG>ExecTime: </ STRONG>0.078<br>
 <STRONG>Done: </ STRONG>true<br>
 <STRONG>RequestDateTime: </ STRONG>2/9/2015 2:33:40 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Command=VALIDATEPASSWORD
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.043
Done=true
RequestDateTime=2/9/2015 2:34:21 PM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

[GetDomainExp](../docs/GetDomainExp.md)

[GetDomainInfo](../docs/getdomaininfo.md)

GetDomains

GetExtendInfo

GetPasswordBit

GetRegistrationStatus

[GetRegLock](../docs/getreglock.md)

[GetRenew](../docs/getrenew.md)

GetSubAccountPassword

SetPassword

[SetRegLock](../docs/SetRegLock.md)

[SetRenew](../docs/SetRenew.md)

StatusDomain