Queue\GetExtAttributes
======================

Retrieve a list of extended attributes for a specific queue.

Usage
-----

Use this command to retrieve a list of extended attributes for a specific queue.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/tld-queue/pages/preconfigure.aspx](https://resellertest.enom.com/tld-queue/pages/preconfigure.aspx)

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                        | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                     | 20 |
| PW | Required           | Account password | 20    |
| QIDList     | Required | Queue identification list in comma-separated format. Use Queue\_GetInfo command to retrieve queue identification number | - |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| ID        | Queue indetifier |
| TLD | Top level domain                                         |
| Attributes    | Attributes collection in XML format |
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

The following query retrieves a list of extended attributes for a TLD available in a queue, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw
&Command=Queue_GetExtAttributes&QIDList=53&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw
&Command=Queue_GetExtAttributes&QIDList=53&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw
&Command=Queue_GetExtAttributes&QIDList=53&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Queue_GetExtAttributes>
 <Queues>
 <Queue>
  <ID>53</ID>
  <TLD>NINJA</TLD>
  <Attributes>
  <Attribute>
   <ID>331</ID>
   <Name>ninja_username</Name>
   <Application>2</Application>
   <UserDefined>True</UserDefined>
   <Required>1</Required>
   <Description>Username</Description>
   <IsChild>0</IsChild>
   <Options />
  </Attribute>
.
.
.
  <Attribute>
   <ID>333</ID>
   <Name>ninja_emailaddress</Name>
   <Application>2</Application>
   <UserDefined>True</UserDefined>
   <Required>1</Required>
   <Description>Email Address</Description>
   <IsChild>0</IsChild>
   <Options />
  </Attribute>
  </Attributes>
 </Queue>
 </Queues>
</Queue_GetExtAttributes>
<Command>QUEUE_GETEXTATTRIBUTES</Command>
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
<ExecTime>0.156</ExecTime>
<Done>true</Done>
<TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
<RequestDateTime>4/30/2013 4:12:07 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
ID: 53
TLD:
ID: 343
Name: Ninja_username
Application: 2
UserDefined: True
Required: 1
Description: Specify Registration Date
IsChild: 0
Options:
Command: QUEUE_GETEXTATTRIBUTES
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.344
Done: true
TrackingKey: 2301fc97-2168-4d98-b110-1afad7b68243
RequestDateTime: 2/5/2015 11:21:55 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
ID=53
TLD=
ID=343
Name=Ninja_username
Application=2
UserDefined=True
Required=1
Description=Specify Registration Date
IsChild=0
Options=
Command=QUEUE_GETEXTATTRIBUTES
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
TrackingKey=a3afbbbb-0372-45f4-904c-aecbc8258a73
RequestDateTime=2/5/2015 11:25:54 AM
```
Related Commands
----------------

GetAgreementPage

PE\_GetPremiumPricing

Queue\_DomainPurchase

Queue\_GetDomains

Queue\_GetInfo

Queue\_GetOrderDetail

Queue\_GetOrders