TLD\Overview
============

Retrieve a detailed overview of a TLD \(Top Level Domain\).

Usage
-----

Use this command to retrieve a detailed overview of a TLD \(Top Level Domain\).

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
https://resellertest.enom.com/interface.asp?command=TLD_Overview&uid=(Required)&pw=(Required)&TLD=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ----------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | TLD\_Overview |
| uid | string | Required | Your Account ID                                                        |
| pw       | string | Required | Your API Token |
| tld | string | Required | Top Level Domain. Permitted values are any type of domain name we support. Examples: - com - net - org - info - co.uk - eu |
| ResponseType  | string | Required | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ------------------------- | ------- | ------------------------------------------------------------------------------------------------ |
| ID | string | TLD identification \(if applicable\) |
| Name | string | TLD name |
| NativeIDN | string | TLD name in native language |
| Description | string | TLD description |
| TimeFrameDescription | string | TLD time-frame description \(pre-order phase\) |
| Requirements | string | TLD requirements |
| DrawNumbers | string | TLD draw-numbers \(pre-order phase\) |
| Categories | string | TLD categories |
| QueueInfo | string | Collection of queue\(s\) information |
| StatusID | string | Queue status ID |
| StatusDesc | string | Queue status description |
| StatusAdditional | string | Queue status sub-description |
| QID | string | Queue ID |
| Phase | string | Queue name or phase |
| DateStart | string | Queue start date |
| DateEnd | string | Queue end date |
| IsActive | string | Is queue open for registration? |
| ApplicationFee | string | Queue application fee |
| ApplicationFeeRefundable | string | Is queue application fee refundable? |
| RegistrationFee | string | Queue registration fee |
| RegistrationFeeRefundable | string | Is queue registration fee refundable? |
| PopularDomains | string | Popular domain\(s\) suggestion based on this TLD |
| SimilarNewTLDs | string | Similar new TLD\(s\) suggestion based on this TLD |
| SimilarTLDs | string | Similar existing TLD\(s\) suggestion based on this TLD |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query retrieves detailed overview of TLD, and requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?command=TLD_Overview&UID=ResellID&PW=resellpw&ResponseType=XML&TLD=ninja
```
```
https://resellertest.enom.com/interface.asp?command=TLD_Overview&UID=ResellID&PW=resellpw&ResponseType=HTML&TLD=ninja
```
```
https://resellertest.enom.com/interface.asp?command=TLD_Overview&UID=ResellID&PW=resellpw&ResponseType=text&TLD=ninja
```
```csharp
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<TLD_Overview>
<ID />
<Name>ninja</Name>
<NativeIDN><![CDATA[]]></NativeIDN>
<Description><![CDATA[For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.]]></Description>
<TimeFrameDescription><![CDATA[This new TLD is not yet available, but it is expected to open for registration Spring of 2013. Please sign up to watch this new gTLD and we will notify you when registration opens. ]]></TimeFrameDescription>
<Requirements />
<DrawNumbers>856</DrawNumbers>
<Categories>Featured,gTLD,Novelty</Categories>
<QueueInfo>
<Queue>
<StatusID>2</StatusID>
<StatusDesc>Evaluation Passed</StatusDesc>
<StatusAdditional>All applications for TLD have passed ICANN’s evaluation.</StatusAdditional>
<QID />
<Phase>Evaluation Passed</Phase>
<DateStart />
<DateEnd />
<IsActive>false</IsActive>
<ApplicationFee></ApplicationFee>
<ApplicationFeeRefundable />
<RegistrationFee></RegistrationFee>
<RegistrationFeeRefundable />
</Queue>
</QueueInfo>
<PopularDomains><![CDATA[car, computer, design, management, training, yourname, yourcompany, info]]></PopularDomains>
<SimilarNewTLDs><![CDATA[.guru, .help, .expert]]></SimilarNewTLDs>
<SimilarTLDs><![CDATA[none]]></SimilarTLDs>
</TLD_Overview>
<Command>TLD_OVERVIEW</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl0vwresell_t1</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.016</ExecTime>
<Done>true</Done>
<TrackingKey>5e84a3c3-f928-4076-9239-c8c2255ca4ee</TrackingKey>
<RequestDateTime>11/22/2013 6:08:46 PM</RequestDateTime>
<debug/>
</interface-response>
```
```csharp
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>ID: </STRONG><br />
<STRONG>Name: </STRONG>ninja<br />
<STRONG>NativeIDN: </STRONG><br />
<STRONG>Description: </STRONG>For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.<br />
<STRONG>TimeFrameDescription: </STRONG>This new TLD is not yet available, but it is expected to open for registration Spring of 2013. Please sign up to watch this new gTLD and we will notify you when registration opens.<br />
<STRONG>Requirements: </STRONG><br />
<STRONG>DrawNumbers: </STRONG>856<br />
<STRONG>Categories: </STRONG>Novelty<br />
<STRONG>CustomMarketing: </STRONG>true<br />
<STRONG>GeneralAvailability: </STRONG>true<br />
<STRONG>StatusID: </STRONG>4<br />
<STRONG>StatusDesc: </STRONG>General Availability<br />
<STRONG>StatusAdditional: </STRONG><br />
<STRONG>QID: </STRONG><br />
<STRONG>Phase: </STRONG>General Availability<br />
<STRONG>DateStart: </STRONG>2014-05-28T10:00:00<br />
<STRONG>DateEnd: </STRONG><br />
<STRONG>IsActive: </STRONG>true<br />
<STRONG>ApplicationFee: </STRONG><br />
<STRONG>ApplicationFeeRefundable: </STRONG><br />
<STRONG>RegistrationFee: </STRONG><br />
<STRONG>RegistrationFeeRefundable: </STRONG><br />
<STRONG>InEAP: </STRONG>false<br />
<STRONG>PopularDomains: </STRONG>car, computer, design, management, training, yourname, yourcompany, info<br />
<STRONG>SimilarNewTLDs: </STRONG><br />
<STRONG>SimilarTLDs: </STRONG><br />
<STRONG>Command: </STRONG>TLD_OVERVIEW<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>0.016<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>57ef0f6a-a458-44db-8b59-840a7cf5ac85<br />
<STRONG>RequestDateTime: </STRONG>2/6/2015 12:23:09 PM<br />
</BODY></HTML>
```
```csharp
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
ID=
Name=ninja
NativeIDN=
Description=For the self-professed masters of the world, .NINJA provides a fun, unique, and flexible TLD option. Use it as blog.NINJA, homeimprovement.NINJA, homecooking.NINJA, and more. .NINJA can also be used for more traditional applications of the word, to promote ninja books, comics, movies, and cultural references, such as comics.NINJA. Whatever the use, .NINJA is a memorable domain namespace guaranteed to help any business, group, or individual using it make a lasting impression.
TimeFrameDescription=This new TLD is not yet available, but it is expected to open for registration Spring of 2013. Please sign up to watch this new gTLD and we will notify you when registration opens.
Requirements=
DrawNumbers=856
Categories=Novelty
CustomMarketing=true
GeneralAvailability=true
StatusID=4
StatusDesc=General Availability
StatusAdditional=
QID=
Phase=General Availability
DateStart=2014-05-28T10:00:00
DateEnd=
IsActive=true
ApplicationFee=
ApplicationFeeRefundable=
RegistrationFee=
RegistrationFeeRefundable=
InEAP=false
PopularDomains=car, computer, design, management, training, yourname, yourcompany, info
SimilarNewTLDs=
SimilarTLDs=
Command=TLD_OVERVIEW
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
ExecTime=0.016
Done=true
TrackingKey=95eb46df-e2b0-4adf-b2f2-af598dc3b5c9
RequestDateTime=2/6/2015 12:23:34 PM
```
Related Commands
----------------

TLD\_AddWatchlist

TLD\_DeleteWatchlist

TLD\_GetTLD

TLD\_GetWatchlist

TLD\_GetWatchlistTlds