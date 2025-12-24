GetDomainStatus
===============

Check the registration status of TLDs that do not register in real time.

Usage
-----

Use this command to check the status of domains that do not register in real time \(including com.au, net.au, org.au, .sg, com.sg and others\). Because of the delay inherent in the non-real-time registrations, wait at least five minutes after yourtransaction to run this command, and run it at intervals of five minutes or longer. This command runs more quicklythan a similar command, StatusDomain.

Use this command to retrieve the most recent order ID for a domain name.

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
https://resellertest.enom.com/interface.asp?command=GetDomainStatus&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&OrderID=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------------ |
| command     | string | Required | TLD\_GetTLD |
| uid | string | Required | Your Account ID                                |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)          |
| TLD       | string | Required | Top-level domain name \(extension\) |
| OrderID | string | Required | Order ID of the most recent transaction for this domain            |
| OrderType    | string | Required | Type of order. Permitted values are Purchase \(default\), Transfer, or Extend |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML   |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| DomainName    | string | Type of contact data being returned                               |
| InAccount    | string | Organization information                                     |
| StatusDesc    | string | First name                                            |
| ExpDate     | string | Last name                                            |
| OrderID     | string | Address line 1                                          |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GETDOMAINSTATUS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=net&orderID=156375934
&ordertype=purchase&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GETDOMAINSTATUS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=net&orderID=156375934
&ordertype=purchase&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GETDOMAINSTATUS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=net&orderID=156375934
&ordertype=purchase&ResponseType=text
```
```
<interface-response>
 <DomainStatus>
 <DomainName>resellerdocs.net</DomainName>
 <InAccount>1</InAccount>
 <ExpDate>10/25/2007 5:18:12 AM</ExpDate>
 <OrderID>156375934</OrderID>
 </DomainStatus>
 <Command>GETDOMAINSTATUS</Command>
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
 <ExecTime>0.563</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 5:05:44 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>DomainName: </ STRONG>resellerdocs.net<br>
<STRONG>InAccount: </ STRONG>1<br>
<STRONG>ExpDate: </ STRONG>10/25/2007 5:18:12 AM<br>
<STRONG>OrderID: </ STRONG>156375934<br>
<STRONG>Command: </ STRONG>GETDOMAINSTATUS<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+8.00<br>
<STRONG>ExecTime: </ STRONG>0.234<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>63f25373-9670-405f-981d-4f484c160222<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 6:07:24 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainName=resellerdocs.net
InAccount=1
ExpDate=10/25/2007 5:18:12 AM
OrderID=156375934
Command=GETDOMAINSTATUS
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
TrackingKey=f84ab999-4381-4aa0-ac9e-74e7c9983b61
RequestDateTime=2/3/2015 6:07:48 PM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetExpiredDomains

GetExtendInfo

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