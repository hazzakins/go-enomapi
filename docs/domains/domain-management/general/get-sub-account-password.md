GetSubAccountPassword
=====================

Get the password for a domain name.

Usage
-----

Use this command to retrieve the password for a domain name that belongs to your account. The password is emailed to the Registrant contact address.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/help/](https://resellertest.enom.com/help/)

On the help page, the Password recovery link calls the GetSubAccountPassword command.

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

The following query requests the password for domain resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetSubAccountPassword&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetSubAccountPassword&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetSubAccountPassword&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The 0 value for ErrCount indicates that the request has been processed successfully. The password is e-mailed to the registrant almost instantly.

```
<?xml version="1.0" ?>
<interface-response>
<Command>GETSUBACCOUNTPASSWORD</Command>
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
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command: GETSUBACCOUNTPASSWORD
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
ExecTime: 0.531
Done: true
RequestDateTime: 2/4/2015 12:41:07 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command=GETSUBACCOUNTPASSWORD
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.281
Done=true
RequestDateTime=2/4/2015 12:41:44 PM
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

GetRegistrationStatus

GetRegLock

GetRenew

StatusDomain

ValidatePassword