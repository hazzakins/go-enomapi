Queue\GetDomains
================

Retrieve a list of domains currently in the specified queue for the account.

Usage
-----

Use this command to retrieve a list of domains currently in the specified queue in the account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=Queue_GetDomains&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status           | Description |
| ---------------- | ------- | -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | Required          | Queue\_GetDomains |
| uid | string | Required | Your Account ID                                                                          |
| pw        | string | Required          | Your API Token |
| DisplayMetrics | boolean | Optional; dafault is true | Display queue domain's metric. Permitted values are true and false                                                |
| DisplayDomains  | boolean | Optional; dafault is false | Display a list of domains. Permitted values are true and false |
| RecordStart | int   | Optional; default is 1 | Return results beginning with this position in the sorted list. For example, StartPosition=26&PagingPageSize=25 returns accounts 26 through 50 in the sorted list |
| PagingSize    | int | Optional default is 25   | Number of accounts to return in this response. Maximum permitted value is 500. |
| SortBy | string | Optional | Sorting parameter. Permitted values are: DomainName TLD Price                                                   |
| SortOrder    | string | Optional; default is Asc  | Sort order. Permitted values are Asc and Desc |
| FilterTLD | string | Optional | Filter by TLD                                                                           |
| FilterQueue   | string | Optional          | Filter by Queue |
| FilterStatus | string | Optional | Filter by Status ID                                                                        |
| FilterStatusDesc | string | Optional          | Filter by Status Description |
| DomainNameFilter | string | Optional | Filter by domain name, contains specific word                                                           |
| ResponseType   | string | Optional          | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------------- | ------- | -------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| WatchlistTotal | int   | Total items in Watchlist \(not a distinct list\). For example, domain.ninja and .ninja are counted as two \(2\) entries |
| WatchlistPreorder | int   | Total items in Watchlist available for Pre-Order |
| WathclistDomainWatched | int   | Total domains \(SLD.TLD\) watched |
| WatchlistTLDWatched | int   | Total TLD watched |
| NewTLDOrders | int   | Number of orders entered to the Queue |
| NewTLDValidated | int   | Number of orders have been Validated |
| NewTLDPending | int   | Number of orders in Pending status |
| NewTLDRejected | int   | Number of orders in Rejected status |
| QueuesDisabled | int   | Is Queue disabled for this account? |
| PreregDisabled |     | Is PreRegistration disabled for this account? |
| OrderID |     | Order ID |
| DomainName |     | Domain name |
| TLDID |     | TLD ID |
| TLDName |     | TLD name |
| TLDDesc |     | TLD description |
| DateStart |     | Queue opening date |
| DateEnd |     | Queue closing date |
| Price |     | Price |
| QID |     | Queue ID |
| QName |     | Queue name |
| StatusID |     | Status ID |
| StatusName |     | Status name |
| ItemCount |     | Total items return in this search \(filter\) |
| ItemTotal |     | Total items in this account |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&Command=Queue_GetDomains&DisplayDomains=true&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&Command=Queue_GetDomains&DisplayDomains=true&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&Command=Queue_GetDomains&DisplayDomains=true&ResponseType=text
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
 <Queue_GetDomains>
 <Metrics>
  <WatchlistTotal>42</WatchlistTotal>
  <WatchlistPreorder>23</WatchlistPreorder>
  <WathclistDomainWatched>24</WathclistDomainWatched>
  <WatchlistTLDWatched>21</WatchlistTLDWatched>
  <NewTLDOrders>786</NewTLDOrders>
  <NewTLDValidated>450</NewTLDValidated>
  <NewTLDPending>70</NewTLDPending>
  <NewTLDRejected>68</NewTLDRejected>
  <QueuesDisabled>false</QueuesDisabled>
  <PreregDisabled>false</PreregDisabled>
 </Metrics>
 <Domains>
  <Domain>
  <OrderID>157910627</OrderID>
  <DomainName>domain.ninja</DomainName>
  <TLDID />
  <TLDName>NINJA</TLDName>
  <TLDDesc></TLDDesc>
  <DateStart>6/12/2013 2:32 PM</DateStart>
  <DateEnd>6/28/2013 2:32 PM</DateEnd>
  <Price>55.00</Price>
  <QID>57</QID>
  <QName>Sunrise B</QName>
  <StatusID>139</StatusID>
  <StatusName>Awaiting Submission</StatusName>
  </Domain>
  <Domain>
  <OrderID>157910451</OrderID>
  <DomainName>teacher.ninja</DomainName>
  <TLDID />
  <TLDName>NINJA</TLDName>
  <TLDDesc></TLDDesc>
  <DateStart>6/12/2013 2:32 PM</DateStart>
  <DateEnd>6/28/2013 2:32 PM</DateEnd>
  <Price>55.00</Price>
  <QID>57</QID>
  <QName>Sunrise B</QName>
  <StatusID>139</StatusID>
  <StatusName>Awaiting Submission</StatusName>
  </Domain>
 </Domains>
 <ItemCount>120</ItemCount>
 <ItemTotal>120</ItemTotal>
 </Queue_GetDomains>
 <Command>QUEUE_GETDOMAINS</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>2.375</ExecTime>
 <Done>true</Done>
 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
 <RequestDateTime>6/3/2013 3:55:35 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>WatchlistTotal: </ STRONG>1289<br>
<STRONG>WatchlistPreorder: </ STRONG>24<br>
<STRONG>WatchlistDomainWatched: </ STRONG>1273<br>
<STRONG>WatchlistTLDWatched: </ STRONG>641<br>
<STRONG>NewTLDOrders: </ STRONG>11<br>
<STRONG>NewTLDValidated: </ STRONG>11<br>
<STRONG>NewTLDPending: </ STRONG>0<br>
<STRONG>NewTLDRejected: </ STRONG>0<br>
<STRONG>RegisteredDomains: </ STRONG>14<br>
<STRONG>ExpiringDomains: </ STRONG>11<br>
<STRONG>ExpiredDomains: </ STRONG>7<br>
<STRONG>RedemptionDomains: </ STRONG>0<br>
<STRONG>QueuesDisabled: </ STRONG>false<br>
<STRONG>PreregDisabled: </ STRONG>false<br>
<STRONG>OrderID: </ STRONG>161808481<br>
<STRONG>DomainName: </ STRONG>mydomain2.ninja<br>
<STRONG>TLDID:</ STRONG><br>
<STRONG>TLDName: </ STRONG>ninja<br>
<STRONG>TLDDesc:</ STRONG><br>
<STRONG>DateStart: </ STRONG>2/17/2014 3:00 PM<br>
<STRONG>DateEnd: </ STRONG>4/21/2014 10:00 AM<br>
<STRONG>OrderDate: </ STRONG>2/5/2015 11:16 AM<br>
<STRONG>Price: </ STRONG>140.00<br>
<STRONG>QID: </ STRONG>227<br>
<STRONG>QName: </ STRONG>Sunrise<br>
<STRONG>StatusID: </ STRONG>679<br>
<STRONG>StatusName: </ STRONG>Awaiting Submission<br>
<STRONG>StatusDesc: </ STRONG>Your Pre-order for Sunrise and Landrush phases will be submitted shortly. For Pre-registration orders go to General Availability, your Pre-order will be submitted the moment the registry is open.<br>
.
.
.
<STRONG>ItemTotal: </ STRONG>72<br>
<STRONG>Command: </ STRONG>QUEUE_GETDOMAINS<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>1.484<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>f978412a-242e-4c77-95cc-1093de9238ff<br>
<STRONG>RequestDateTime: </ STRONG>2/5/2015 11:18:25 AM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
WatchlistTotal=1289
WatchlistPreorder=24
WatchlistDomainWatched=1273
WatchlistTLDWatched=641
NewTLDOrders=11
NewTLDValidated=11
NewTLDPending=0
NewTLDRejected=0
RegisteredDomains=14
ExpiringDomains=11
ExpiredDomains=7
RedemptionDomains=0
QueuesDisabled=false
PreregDisabled=false
OrderID=161808481
DomainName=mydomain2.ninja
TLDID=
TLDName=ninja
TLDDesc=
DateStart=2/17/2014 3:00 PM
DateEnd=4/21/2014 10:00 AM
OrderDate=2/5/2015 11:16 AM
Price=140.00
QID=227
QName=Sunrise
StatusID=679
StatusName=Awaiting Submission
StatusDesc=Your Pre-order for Sunrise and Landrush phases will be submitted shortly. For Pre-registration orders go to General Availability, your Pre-order will be submitted the moment the registry is open.
OrderID=161808479
DomainName=mydomain1.ninja
TLDID=
TLDName=ninja
TLDDesc=
DateStart=2/17/2014 3:00 PM
DateEnd=4/21/2014 10:00 AM
OrderDate=2/5/2015 11:15 AM
Price=140.00
QID=227
QName=Sunrise
StatusID=679
StatusName=Awaiting Submission
StatusDesc=Your Pre-order for Sunrise and Landrush phases will be submitted shortly. For Pre-registration orders go to General Availability, your Pre-order will be submitted the moment the registry is open.
.
.
.
ItemTotal=72
Command=QUEUE_GETDOMAINS
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.109
Done=true
TrackingKey=acfdb046-251d-40ae-9950-f4da03202617
RequestDateTime=2/5/2015 11:19:58 AM
```
Related Commands
----------------

GetAgreementPage

PE\_GetPremiumPricing

Queue\_DomainPurchase

Queue\_GetExtAttributes

Queue\_GetInfo

Queue\_GetOrderDetail

Queue\_GetOrders