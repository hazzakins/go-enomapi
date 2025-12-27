GetDomainNameID
===============

Retrieves the ID number for a domain.

Usage
-----

In some API commands, you can use the domain name ID instead of the SLD and TLD.

Our database associates a unique ID with each domain name, and all activity in the database is tracked by ID rather than by domain name.

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
https://resellertest.enom.com/interface.asp?command=
GetDomainNameID&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&responsetype=xml
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------- |
| command     | string | Required | GetDomainNameID |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)        |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| DomainRRP    | string | Processor identifier of the registrar which manages the domain                  |
| SLD       | string | Second-level domain name \(for example, enom in enom.com\)                   |
| TLD       | string | Top-level domain name \(extension\)                               |
| DomainNameID   | int | Identification number of this domain name                            |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getdomainnameid&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getdomainnameid&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getdomainnameid&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
```
<interface-response>
 <SLD>resellerdocs</SLD>
 <TLD>com</TLD>
 <DomainNameID>152533676</DomainNameID>
 <Command>GETDOMAINNAMEID</Command>
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
 <ExecTime>0.108</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/8/2011 4:24:19 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>SLD: </ STRONG>resellerdocs<br>
<STRONG>TLD: </ STRONG>com<br>
<STRONG>DomainNameID: </ STRONG>152533676<br>
<STRONG>Command: </ STRONG>GETDOMAINNAMEID<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+08.00<br>
<STRONG>ExecTime: </ STRONG>0.125<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 5:50:30 PM<br>
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
SLD=resellerdocs
TLD=com
DomainNameID=152533676
Command=GETDOMAINNAMEID
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.047
Done=true
RequestDateTime=2/3/2015 5:50:53 PM
```
Related Commands
----------------

GetDomainSLDTLD