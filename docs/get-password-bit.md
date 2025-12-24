GetPasswordBit
==============

Check to see if a password is set for a domain name.

Usage
-----

Use this command to determine whether a password is set for a domain name, and what the password is.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

On the domain control panel, in the utilities box, the Domain Access Password link calls the GetPasswordBit command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.    | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| PasswordSet | Returns 1 if a password is set, 0 otherwise.                          |
| DomainPassword  | If PasswordSet=1, then a password is returned in this parameter. |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the domain name password for resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getpasswordbit&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=getpasswordbit&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getpasswordbit&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
In the response, the password-set value of 1 indicates that a password is set. The DomainPassword value indicates that the domain password is userpw.

```
<?xml version="1.0" ?>
<interface-response>
<DomainPassword>userpw</DomainPassword>
<password-set>1</password-set>
<Command>GETPASSWORDBIT</Command>
<ErrCount>0</ErrCount>
<Server>ResellerTest</Server>
<Site>enom</Site>
<Done>true</Done>
<debug>
 <![CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
DomainPassword: resellerdocspw
PasswordSet: 1
Command: GETPASSWORDBIT
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +8.00
ExecTime: 0.063
Done: true
TrackingKey: 15c817bc-45ac-4942-a67d-fd20077b3821
RequestDateTime: 2/4/2015 11:13:56 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainPassword=resellerdocspw
PasswordSet=1
Command=GETPASSWORDBIT
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.031
Done=true
TrackingKey=8c888f47-7d8d-43b9-a9f2-bdb0f69d7e75
RequestDateTime=2/4/2015 11:14:18 AM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetExtendInfo

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

StatusDomain

ValidatePassword