StatusDomain
============

Get information about the status of a single domain name: registrar, expiration date, and whether it is in your account.

Usage
-----

Use this command when you have received an RRPCode value of 724 when attempting to register a domain. The response for the Purchase command gives additional information explaining the 724 code.

For other aspects of domain status, use the GetDomainStatus command for the fastest response.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The query returns an error if the domain does not belong to this account.

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
| OrderType    | Optional | Options are Purchase\(default\), Transfer or Extend.    | 10 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| DomainName    | Name being statused Registrar Returns Known, Unknown, or None. |
| InAccount | Returns 0 if not in an eNom account, 1 if in your eNom account or 2 if in another eNom account. |
| OrderID     | ID number of the most recent order that included this domain. |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XMLin your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

The following query requests information about domain resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=statusdomain&uid=resellid&pw=resellpw&
sld=resellerdocs&tld=com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=statusdomain&uid=resellid&pw=resellpw&
sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=statusdomain&uid=resellid&pw=resellpw&
sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows:

```
<interface-response>
<DomainStatus>
 <DomainName>resellerdocs.com</DomainName>
 <Registrar>Known</Registrar>
 <InAccount>1</InAccount>
 <ExpDate>6/10/2019 3:56:56 PM</ExpDate>
 <OrderID/>
</DomainStatus>
<RRPCode>200</RRPCode>
<RRPText>Command completed successfully</RRPText>
<Command>STATUSDOMAIN</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+08.00</TimeDifference>
<ExecTime>0.453</ExecTime>
<Done>true</Done>
<RequestDateTime>12/12/2011 12:59:00 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainName: resellerdocs.com
Registrar: Known
InAccount: 1
ExpDate: 6/10/2019 3:56:56 PM
OrderID:
RRPCode: 200
RRPText: Command completed successfully
Command: STATUSDOMAIN
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: enom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.844
Done: true
RequestDateTime: 2/6/2015 10:26:46 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainName=resellerdocs.com
Registrar=Known
InAccount=1
ExpDate=6/10/2019 3:56:56 PM
OrderID=
RRPCode=200
RRPText=Command completed successfully
Command=STATUSDOMAIN
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=enom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.172
Done=true
RequestDateTime=2/6/2015 10:27:12 AM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetDomainStatus

GetExtendInfo

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

ValidatePassword