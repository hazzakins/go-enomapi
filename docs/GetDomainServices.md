GetDomainServices
=================

Retrieve the settings for domain services and value-added services for a domain.

Usage
-----

Use this command to retrieve information about email forwarding, URL forwarding, and other domain and value added services for a domain

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
https://resellertest.enom.com/interface.asp?command=GetDomainServices&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetDomainServices |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| sld | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| tld       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter   | Type | Description                                           |
| -------------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command       | string | Name of command executed                                     |
| WebSite       | string | Web site setting                                         |
| Phone        | string | Name-My-Phone setting                                     |
| IPResolver      | string | Setting for resolving dynamic IP addresses                            |
| Map         | string | Name-My-Map setting                                      |
| DomainNameID     | string | ID number for this domain, from our internal records                       |
| EMailForwarding   | string | Email forwarding setting for this domain                             |
| EmailForwardExpDate | string | Expiration date for email forwarding                               |
| EMailAutoRenew    | string | Auto-renew setting for email forwarding                             |
| URLForwarding    | string | URL forwarding setting for this domain                              |
| URLForwardExpDate  | | Expiration date for email forwarding                               |
| URLAutoRenew     | string | Auto-renew setting for email forwarding                             |
| EMailForwardingPrice | string | Price for email forwarding                                    |
| URLForwardingPrice  | string | Price for URL forwarding                                     |
| ErrCount       | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX         | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done         | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=GetDomainServices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainServices&uid=resellid&pw=resellpw
&sld=resellerdocs&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainServices&uid=resellid&pw=resellpw
&sld=resellerdocs&ResponseType=text
```
```
<interface-response>
 <domainservices>
 <website>1</website>
 <phone>false</phone>
 <ipresolver>1</ipresolver>
 <map>true</map>
 <domainnameid>152533676</domainnameid>
 <mobilizer>1117</mobilizer>
 <mobilizersourcedomain/>
 <valueadd>
  <EmailForwarding>0</EmailForwarding>
  <EmailForwardExpDate/>
  <EmailAutoRenew/>
  <URLForwarding>0</URLForwarding>
  <URLForwardExpDate/>
  <URLAutoRenew/>
  <emailForwardingPrice>0.00</emailForwardingPrice>
  <urlForwardingPrice>0.00</urlForwardingPrice>
 </valueadd>
 </domainservices>
 <Command>GETDOMAINSERVICES</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+8.00</TimeDifference>
 <ExecTime>0.188</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>fa977e24-b1d0-4837-b145-9c26b5f96990</TrackingKey>
 <RequestDateTime>12/8/2011 4:28:18 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
 <STRONG>Reseller: </ STRONG>1<br>
<STRONG>website: </ STRONG>1<br>
<STRONG>phone: </ STRONG></ STRONG>false<br>
<STRONG>ipresolver: </ STRONG>1<br>
<STRONG>map: </ STRONG>true<br>
<STRONG>domainnameid: </ STRONG>152533676<br>
<STRONG>mobilizer: </ STRONG>1117<br>
<STRONG>mobilizersourcedomain:</ STRONG><br>
<STRONG>EmailForwarding: </ STRONG>0<br>
<STRONG>EmailForwardExpDate:</ STRONG><br>
<STRONG>EmailAutoRenew:</ STRONG><br>
<STRONG>URLForwarding: </ STRONG>0<br>
<STRONG>URLForwardExpDate:</ STRONG><br>
<STRONG>URLAutoRenew:</ STRONG><br>
<STRONG>emailForwardingPrice: </ STRONG>0.00<br>
<STRONG>urlForwardingPrice: </ STRONG>0.00<br>
<STRONG>Command: </ STRONG>GETDOMAINSERVICES<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+8.00<br>
<STRONG>ExecTime: </ STRONG>0.078<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>1977a456-808c-4ba9-b426-f371e702fc99<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 6:01:13 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
website=1
phone=false
ipresolver=1
map=true
domainnameid=152533676
mobilizer=1117
mobilizersourcedomain=
EmailForwarding=0
EmailForwardExpDate=
EmailAutoRenew=
URLForwarding=0
URLForwardExpDate=
URLAutoRenew=
emailForwardingPrice=0.00
urlForwardingPrice=0.00
Command=GETDOMAINSERVICES
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
TrackingKey=ec9d74e6-3564-4951-9ef3-a6ccb9e9b7d6
RequestDateTime=2/3/2015 6:01:40 PM
```
Related Commands
----------------

DisableServices

EnableServices

GetIPResolver

ServiceSelect

SetDomainSubServices

SetIPResolver