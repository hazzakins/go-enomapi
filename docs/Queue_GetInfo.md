Queue\GetInfo
=============

Retrieve a list of all queues, which can be filtered as appropriate.

Usage
-----

Use this command to retrieve a list of all queues, which can be filtered as appropriate.

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
https://resellertest.enom.com/interface.asp?command=Queue_GetInfo&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter   | Type | Status  | Description |
| ------------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command       | string | Required | SetRenew |
| uid | string | Required | Your Account ID                              |
| pw         | string | Required | Your API Token |
| Category | string | Optional | Filter by category                            |
| FilterTLD      | string | Optional | Filter by TLD |
| FilterTLDStatusDesc | string | Optional | Filter by TLD status description                     |
| FilterQStatusName  | string | Optional | Filter by queue status name |
| FilterExtAttributes | string | Optional | Filter by extended attributes requirement                 |
| FilterStartDate   | string | Optional | Filter by queue start date |
| FilterEndDate | string | Optional | Filter by queue end date                         |
| DisplayComingSoon  | string | Optional | Filter by queue availability |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

If you submit this command without filter parameters, the server will return every TLD and queue for which you have a watchlist item. If you want to only receive the watchlist items for active queues, you must include the DisplayComingSoon=False parameter.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| TLD       | string | Top level domain                                         |
| NativeIDN    | string | IDN TLD in native characters                                   |
| CategoryName   | string | TLD category                                           |
| CategoryDesc   | string | TLD category description                                     |
| QID       | string | Queue indetifier                                         |
| QName      | string | Queue name                                            |
| QStatusID    | string | Queue status identifier                                     |
| QStatusName   | string | Queue status name                                        |
| ExtAttributes  | string | Does this Queue require extended attributes?                           |
| StartDate    | string | Start date                                            |
| EndDate     | string | End date                                             |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&ResponseType=xml
&Command=Queue_GetInfo&FilterTLD=Ninja
```
```
https://resellertest.enom.com/interface.asp?
command=setrenew&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&renewflag=1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=setrenew&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&renewflag=1&responsetype=text
```
```csharp
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
 <Queue_GetInfo>
 <Queue>
  <TLD>NINJA</TLD>
  <NativeIDN></NativeIDN>
  <CategoryName>Novelty</CategoryName>
  <CategoryDesc></CategoryDesc>
  <TLDDescription><![CDATA[For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.]]></TLDDescription>
  <TLDStatusID>2</TLDStatusID>
  <TLDStatusDesc>Evaluation Passed</TLDStatusDesc>
  <QID>53</QID>
  <QName>Sunrise</QName>
  <QStatusID>1</QStatusID>
  <QStatusName>Open</QStatusName>
  <ExtAttributes>false</ExtAttributes>
  <StartDate>2013-04-16T15:59:00</StartDate>
  <EndDate>2013-10-28T14:32:21</EndDate>
  <AnnouncementDate>2013-02-05T08:00:00</AnnouncementDate>
 </Queue>
 <TotalRecords>1</TotalRecords>
 </Queue_GetInfo>
 <Command>QUEUE_GETINFO</Command>
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
 <ExecTime>0.031</ExecTime>
 <Done>true</Done>
 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
 <RequestDateTime>4/30/2013 3:43:35 PM</RequestDateTime>
</interface-response>
```
```csharp
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>TLD: </STRONG>NINJA<br />
<STRONG>NativeIDN:</STRONG><br />
<STRONG>CategoryName: </STRONG>Novelty<br />
<STRONG>CategoryDesc:</STRONG><br />
<STRONG>TLDDescription: </STRONG>For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.<br />
<STRONG>TLDStatusID: </STRONG>4<br />
<STRONG>TLDStatusDesc: </STRONG>General Availability<br />
<STRONG>QID:</STRONG><br />
<STRONG>QName:</STRONG><br />
<STRONG>QStatusID: </STRONG>2<br />
<STRONG>QStatusName: </STRONG>Close<br />
<STRONG>ExtAttributes: </STRONG>false<br />
<STRONG>StartDate:</STRONG><br />
<STRONG>EndDate:</STRONG><br />
<STRONG>AnnouncementDate:</STRONG><br />
<STRONG>Price:</STRONG><br />
<STRONG>InEAP: </STRONG>false<br />
<STRONG>TLD: </STRONG>NINJA<br />
<STRONG>NativeIDN:</STRONG><br />
<STRONG>CategoryName: </STRONG>Novelty<br />
<STRONG>CategoryDesc:</STRONG><br />
<STRONG>TLDDescription: </STRONG>For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.<br />
<STRONG>TLDStatusID: </STRONG>4<br />
<STRONG>TLDStatusDesc: </STRONG>General Availability<br />
<STRONG>QID:</STRONG><br />
<STRONG>QName: </STRONG>General Availability<br />
<STRONG>QStatusID: </STRONG>0<br />
<STRONG>QStatusName: </STRONG>Open<br />
<STRONG>ExtAttributes: </STRONG>false<br />
<STRONG>StartDate: </STRONG>2014-05-28T10:00:00<br />
<STRONG>EndDate:</STRONG><br />
<STRONG>AnnouncementDate:</STRONG><br />
<STRONG>Price:</STRONG><br />
<STRONG>InEAP: </STRONG>false<br />
<STRONG>TotalRecords: </STRONG>2<br />
<STRONG>Command: </STRONG>QUEUE_GETINFO<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable:</STRONG><br />
<STRONG>IsRealTimeTLD:</STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.234<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>a73b5c86-03aa-4028-84e9-eb2fc479067f<br />
<STRONG>RequestDateTime: </STRONG>2/5/2015 11:30:01 AM<br />
```
```csharp
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
TLD=NINJA
NativeIDN=
CategoryName=Novelty
CategoryDesc=
TLDDescription=For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.
TLDStatusID=4
TLDStatusDesc=General Availability
QID=
QName=
QStatusID=2
QStatusName=Close
ExtAttributes=false
StartDate=
EndDate=
AnnouncementDate=
Price=
InEAP=false
TLD=NINJA
NativeIDN=
CategoryName=Novelty
CategoryDesc=
TLDDescription=For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.
TLDStatusID=4
TLDStatusDesc=General Availability
QID=
QName=General Availability
QStatusID=0
QStatusName=Open
ExtAttributes=false
StartDate=2014-05-28T10:00:00
EndDate=
AnnouncementDate=
Price=
InEAP=false
TotalRecords=2
Command=QUEUE_GETINFO
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
ExecTime=0.656
Done=true
TrackingKey=4f8f7be0-487b-4606-9307-1824ba61042b
RequestDateTime=2/5/2015 11:30:44 AM
```
Related Commands
----------------

GetAgreementPage

PE\_GetPremiumPricing

Queue\_DomainPurchase

Queue\_GetDomains

Queue\_GetExtAttributes

Queue\_GetOrderDetail

Queue\_GetOrders