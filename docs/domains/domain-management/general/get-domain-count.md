GetDomainCount
==============

Get the count for registered, hosted, processing, watch list, cart item and expired domains.

Usage
-----

Use this command to return the number of domains in each category.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

GetDomainCount is not implemented on eNom.com, but clicking the log-in button returns information that includes counts of the domains in each category.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ------------------- | ------------------------------------------------------------------------------------------------ |
| RegisteredCount | Number of registered domain names                                |
| HostCount      | Number of hosted domain names |
| ExpiringCount | Number of expiring domain names                                 |
| ExpiredDomainsCount | Number of expired domain names |
| RGP | Number of domain names in Redemption Grace Period                        |
| ExtendedRGP     | Number of domain names in Extended Redemption Grace Period |
| KeywordCount | Number of keywords \(obsolete feature\)                             |
| ProcessCount    | Number of domain names that are still processing |
| WatchlistCount | Number of domain names in the watchlist                             |
| CartItemCount    | Number of items in the shopping cart |
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

The following query requests a tally of domains in account resellid, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetDomainCount&uid=resellid&pw=resellpw
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainCount&uid=resellid&pw=resellpw
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainCount&uid=resellid&pw=resellpw
&ResponseType=text
```
The response indicates that the account has 3084 registered and 190 hosted domains:

```html
<interface-response>
<RegisteredCount>3084</RegisteredCount>
<HostCount>190</HostCount>
<ExpiringCount>565</ExpiringCount>
<ExpiredDomainsCount>5379</ExpiredDomainsCount>
<RGP>0</RGP>
<ExtendedRGP>0</ExtendedRGP>
<KeywordCount>0</KeywordCount>
<ProcessCount>0</ProcessCount>
<WatchlistCount>160</WatchlistCount>
<CartItemCount>3</CartItemCount>
<trafficmsg>
 <VistaCustomer>True</VistaCustomer>
 <RedirectorData>False</RedirectorData>
 <Month/>
 <PageViews/>
 <Visitors/>
 <FreeTrial>False</FreeTrial>
 <PDQVista/>
</trafficmsg>
<Command>GETDOMAINCOUNT</Command>
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
<ExecTime>0.578</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 4:05:04 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RegisteredCount: 2692
HostCount: 204
ExpiringCount: 502
ExpiredDomainsCount: 7608
RGP: 0
ExtendedRGP: 0
KeywordCount: 0
ProcessCount: 0
WatchlistCount: 238
CartItemCount: 34
VistaCustomer: True
RedirectorData: False
Month:
PageViews:
Visitors:
FreeTrial: False
PDQVista:
Command: GETDOMAINCOUNT
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.406
Done: true
TrackingKey: e4ba9ab9-0919-4166-bcd4-c54c1558ed2a
RequestDateTime: 2/3/2015 5:34:22 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
RegisteredCount=2692
HostCount=204
ExpiringCount=502
ExpiredDomainsCount=7608
RGP=0
ExtendedRGP=0
KeywordCount=0
ProcessCount=0
WatchlistCount=238
CartItemCount=34
VistaCustomer=True
RedirectorData=False
Month=
PageViews=
Visitors=
FreeTrial=False
PDQVista=
Command=GETDOMAINCOUNT
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.125
Done=true
TrackingKey=4d6b3557-7edc-47e7-9292-497262d644c5
RequestDateTime=2/3/2015 5:35:34 PM
```
Related Commands
----------------

GetAllDomains

GetDomainExp

GetDomainInfo

GetDomains

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