GetAllDomains
=============

Get all domain names in an account.

Usage
-----

Use this command to list all the domain names in an account, with their domain name IDs and expiration dates.

We recommend that you use this command only for accounts with fewer than 200 names, because longer lists can time out. For large accounts, use GetDomains.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- We recommend that you use this command only for accounts with fewer than 200 names, becauselonger lists can time out. For large accounts, use GetDomains.
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                               | Max Size |
| --------------- | ----------------------------- | -------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                             | 20 |
| PW | Required           | Account password | 20    |
| UseDNS     | Optional | Returns a list of domains that use custom \(user-specified\) name servers. Permitted value is custom. | 10 |
| GetDefaultOnly | Optional           | Returns default name servers. Value is 1. | 1    |
| Letter     | Optional | Returns domains with the specified first character. Permitted values are 0 through 9 and A through Z. | 1 |
| RegistrarLock | Optional           | Returns domains with the specified Registrar Lock status. Permitted values are Locked and Not Locked. | 10    |
| AutoRenew    | Optional | Returns domains with the specified Auto Renew status. Permitted values are Yes and No.         | 3 |
| NameServer | Optional           | Returns domains that use the specified name server. Use format dns1.nameserver.com. | 135   |
| UseEnomNS    | Optional | Returns domains that use the specified name servers. Permitted values are Yes or No.          | 3 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| DomainName    | Name of the domain described in this node |
| DomainNameID | Database ID of this domain                                    |
| Expiration-Date | Expiration date of this domain |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests all domains in account resellid, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetAllDomains&uid=resellid
&pw=resellpw&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetAllDomains&uid=resellid
&pw=resellpw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetAllDomains&uid=resellid
&pw=resellpw&ResponseType=text
```
The response lists all domain names in the account, with the ID number and expiration date for each domain:

```
<interface-response>
<GetAllDomains>
 <DomainDetail>
 <DomainName>00000000001111.com</DomainName>
 <DomainNameID>340691704</DomainNameID>
 <expiration-date>7/7/2012 8:37:00 AM</expiration-date>
 <lockstatus>Locked</lockstatus>
 <AutoRenew>No</AutoRenew>
 </DomainDetail>
 <DomainDetail>
 <DomainName>00000startajay.info</DomainName>
 <DomainNameID>152708845</DomainNameID>
 <expiration-date>2/4/2011 6:25:00 AM</expiration-date>
 <lockstatus>Not Locked</lockstatus>
 <AutoRenew>No</AutoRenew>
 </DomainDetail>
.
.
.
 <domaincount>3077</domaincount>
 <UserRequestStatus>DomainBox</UserRequestStatus>
</GetAllDomains>
<Command>GETALLDOMAINS</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod/>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>74.456</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 3:38:26 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainName1: 00000101010101fourth.com
DomainNameID1: 345245402
expiration-date1: 8/7/2015 7:03:00 AM
lockstatus1: Not Locked
AutoRenew1: No
DomainName2: 00000bbb.com
DomainNameID2: 345730105
expiration-date2: 8/13/2015 3:01:33 AM
lockstatus2: Locked
AutoRenew2: No
DomainName3: 000ajaycouk.us
DomainNameID3: 152708789
expiration-date3: 1/30/2012 11:59:00 PM
lockstatus3: Locked
AutoRenew3: No
.
.
.
domaincount: 2691
UserRequestStatus: DomainBox
Command: GETALLDOMAINS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 10.172
Done: true
RequestDateTime: 2/3/2015 4:29:33 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
DomainName1=00000101010101fourth.com
DomainNameID1=345245402
expiration-date1=8/7/2015 7:03:00 AM
lockstatus1=Not Locked
AutoRenew1=No
DomainName2=00000bbb.com
DomainNameID2=345730105
expiration-date2=8/13/2015 3:01:33 AM
lockstatus2=Locked
AutoRenew2=No
DomainName3=000ajaycouk.us
DomainNameID3=152708789
expiration-date3=1/30/2012 11:59:00 PM
lockstatus3=Locked
AutoRenew3=No
.
.
.
domaincount=2691
UserRequestStatus=DomainBox
Command=GETALLDOMAINS
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=4.953
Done=true
RequestDateTime=2/3/2015 4:30:23 PM
```
Related Commands
----------------

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetExtendInfo

GetHomeDomainList

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

StatusDomain

ValidatePassword