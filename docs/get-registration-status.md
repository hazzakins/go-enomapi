GetRegistrationStatus
=====================

Get the registration and paid status of a domain name.

Usage
-----

Use this command to get the registration status and paid status of domains in your account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainManager.asp?tab=iown](https://resellertest.enom.com/domains/DomainManager.asp?tab=iown)

This command is not implemented on enom.com; however, the registered tab of the my domains page returns similar information.

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
| ------------------ | ------------------------------------------------------------------------------------------------ |
| RegistrationStatus | Registration Status of the Domain name. Options are Processing, Registered, Hosted, Null.   |
| PurchaseStatus   | Purchase status of the Domain name. Options are Processing, Paid, Null. |
| Command | Name of command executed                                     |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done        | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves information on whether the domain name resellerdocs.com is registered and paid for, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getregistrationstatus&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=getregistrationstatus&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=Extend_RGP&uid=resellid&pw=resellpw
&sld=VeryExpiredName&tld=com&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
<RegistrationStatus>Registered</RegistrationStatus>
<Command>GETREGISTRATIONSTATUS</Command>
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
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RegistrationStatus: Registered
Command: GETREGISTRATIONSTATUS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.063
Done: true
RequestDateTime: 2/4/2015 11:33:34 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
RegistrationStatus=Registered
Command=GETREGISTRATIONSTATUS
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
ExecTime=0.063
Done=true
RequestDateTime=2/4/2015 11:33:54 AM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetExtendInfo

GetPasswordBit

GetRegLock

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

StatusDomain

ValidatePassword