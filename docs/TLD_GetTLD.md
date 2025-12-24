TLD\GetTLD
==========

Retrieve a list of TLDs currently offered for this account.

Usage
-----

Use this command to retrieve a list of TLDs currently offered for this account.

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
https://resellertest.enom.com/interface.asp?command=TLD_GetTLD&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | TLD\_GetTLD |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| Category | string | Required | Return a list of TLDs filtered by Primary Category            |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| Name | string | TLD name \(IDN is displayed as PUNY Code\). International Domain Name \(IDN\) allows more web users to navigate the Internet in their preferred native language. Most domain names are registered in ASCII characters. The non-Latin scripts such as Chinese, Russian, Korean or other languages cannot be rendered in ASCII. IDN is fully supported in PUNY Code and URL Encoded format. To get more information and conversion tool for PUNY Code, check the sites below: - Verisign's IDN Tool: [http://mct.verisign-grs.com](http://mct.verisign-grs.com/) - RFC 3492: [http://tools.ietf.org/html/rfc3492](http://tools.ietf.org/html/rfc3492) |
| CategoryNames | string | TLD Category. Each TLD may have more than one category. The first category in this list \(comma delimited format\) is the Primary |
| NativeIDN | string | TLD Name in native characters |
| Description | string | Description of NativeIDN |
| StatusID | string | TLD Status ID |
| StatusDesc | string | TLD Status Description |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TLD_GetTLD&UID=ResellID&PW=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_GetTLD&UID=ResellID&PW=resellpw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_GetTLD&UID=ResellID&PW=resellpw&ResponseType=text
```
```
<?xml version="1.0"?>
<interface-response>
 <TLD_GetTLD>
 <TLD>
  <Name><![CDATA[ ABOGADO ]]></Name>
  <CategoryNames>Services</CategoryNames>
  <NativeIDN><![CDATA[ ]]></NativeIDN>
  <Description><![CDATA[ ]]></Description>
  <TLDStatusID>1</TLDStatusID>
  <TLDStatusDesc><![CDATA[ Awaiting Evaluation ]]></TLDStatusDesc>
 </TLD>
.
.
.
 <TLD>
  <Name><![CDATA[ xn--io0a7i ]]></Name>
  <CategoryNames>IDN,Technology</CategoryNames>
  <NativeIDN><![CDATA[ 网络 ]]></NativeIDN>
  <Description>
  <![CDATA[ The string of "网络" means network. It is a collection
of hardware components and computers interconnected by communication
channels that allow sharing of resources and information. ]]>
  </Description>
  <TLDStatusID>1</TLDStatusID>
  <TLDStatusDesc><![CDATA[ Awaiting Evaluation ]]></TLDStatusDesc>
 </TLD>
 <TotalRecords>678</TotalRecords>
 </TLD_GetTLD>
 <Command>TLD_GETTLD</Command>
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
 <ExecTime>2.967</ExecTime>
 <Done>true</Done>
 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
 <RequestDateTime>7/17/2012 11:11:11 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Name1: </ STRONG>5earlyaccess<br>
<STRONG>CategoryNames1: </ STRONG>Shopping<br>
<STRONG>NativeIDN1:</ STRONG><br>
<STRONG>Description1:</ STRONG><br>
<STRONG>TLDStatusID1: </ STRONG>4<br>
<STRONG>TLDStatusDesc1: </ STRONG>General Availability<br>
<STRONG>Name2: </ STRONG>5earlyaccess<br>
<STRONG>CategoryNames2: </ STRONG>Shopping<br>
<STRONG>NativeIDN2:</ STRONG><br>
<STRONG>Description2:</ STRONG><br>
<STRONG>TLDStatusID2: </ STRONG>4<br>
<STRONG>TLDStatusDesc2: </ STRONG>General Availability<br>
<STRONG>.
<STRONG>.
<STRONG>.
<STRONG>TotalRecords: </ STRONG>1280<br>
<STRONG>count: </ STRONG>1281<br>
<STRONG>Command: </ STRONG>TLD_GETTLD<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.359<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>4bf9b99a-84cd-466e-8209-e44e7bc502b2<br>
<STRONG>RequestDateTime: </ STRONG>2/6/2015 12:07:01 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Name1=5earlyaccess
CategoryNames1=Shopping
NativeIDN1=
Description1=
TLDStatusID1=4
TLDStatusDesc1=General Availability
Name2=5earlyaccess
CategoryNames2=Shopping
NativeIDN2=
Description2=
TLDStatusID2=4
TLDStatusDesc2=General Availability
.
.
.
TotalRecords=1280
count=1281
Command=TLD_GETTLD
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
ExecTime=0.375
Done=true
TrackingKey=e3d3187b-d913-4f02-8dea-c81e7aa0443c
RequestDateTime=2/6/2015 12:08:12 PM
```
Related Commands
----------------

TLD\_AddWatchlist

TLD\_DeleteWatchlist

TLD\_GetWatchlist

TLD\_GetWatchlistTlds

TLD\_Overview