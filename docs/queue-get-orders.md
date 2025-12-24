Queue\GetOrders
===============

Retrieve a list of queue domain order history.

Usage
-----

Use this command to retrieve a list of queue domain order history.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/reports.aspx](https://resellertest.enom.com/myaccount/reports.aspx)

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
| Input Parameter  | Status | Description                      | Max Size |
| ----------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID        | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| PortalUserPartyID | Optional | Portal user identification              | |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| OrderID     | Order ID |
| StatusID | Status ID                                            |
| StatusName    | Status Name |
| StatusDesc | Status Description                                        |
| Success     | Total number of successful item in this order |
| Failed | Total number of failed item in this order                            |
| Pending     | Total number of pending item in this order |
| OrderDate | Order date                                            |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieve a list of queue domain order history, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&Command=Queue_GetOrders
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&Command=Queue_GetOrders
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&Command=Queue_GetOrders
&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Queue_GetOrders>
 <Orders>
 <Order>
  <OrderID>157896057</OrderID>
  <StatusID>0</StatusID>
  <StatusName>New</StatusName>
  <StatusDesc></StatusDesc>
  <Success>0</Success>
  <Failed>0</Failed>
  <Pending>1</Pending>
  <OrderDate>4/11/2013 5:11 PM</OrderDate>
 </Order>
 </Orders>
 <ItemCount>1</ItemCount>
 <ItemTotal>1</ItemTotal>
</Queue_GetOrders>
<Command>QUEUE_GETORDERS</Command>
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
<ExecTime>3.047</ExecTime>
<Done>true</Done>
<TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
<RequestDateTime>6/5/2013 3:27:10 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
OrderID: 161808481
StatusID: 0
StatusName: New
StatusDesc:
Success: 0
Failed: 0
Pending: 2
OrderDate: 2/5/2015 11:16 AM
OrderID: 161808479
StatusID: 0
StatusName: New
StatusDesc:
Success: 0
Failed: 0
Pending: 2
OrderDate: 2/5/2015 11:15 AM
OrderID: 161786811
StatusID: 0
StatusName: New
StatusDesc:
Success: 0
Failed: 0
Pending: 2
OrderDate: 11/27/2014 9:38 PM
.
.
.
ItemCount: 67
ItemTotal: 67
Command: QUEUE_GETORDERS
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
ExecTime: 9.828
Done: true
TrackingKey: 3698d5c5-8ad4-4786-acb4-a6072e1fa3cd
RequestDateTime: 2/5/2015 11:45:00 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
OrderID=161808481
StatusID=0
StatusName=New
StatusDesc=
Success=0
Failed=0
Pending=2
OrderDate=2/5/2015 11:16 AM
OrderID=161808479
StatusID=0
StatusName=New
StatusDesc=
Success=0
Failed=0
Pending=2
OrderDate=2/5/2015 11:15 AM
OrderID=161786811
StatusID=0
StatusName=New
StatusDesc=
Success=0
Failed=0
Pending=2
OrderDate=11/27/2014 9:38 PM
.
.
.
ItemCount=67
ItemTotal=67
Command=QUEUE_GETORDERS
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
ExecTime=1.281
Done=true
TrackingKey=cff86e45-9eb2-45c2-ba28-0394934d1308
RequestDateTime=2/5/2015 11:46:16 AM
```
Related Commands
----------------

GetAgreementPage

PE\_GetPremiumPricing

Queue\_DomainPurchase

Queue\_GetDomains

Queue\_GetExtAttributes

Queue\_GetInfo

Queue\_GetOrderDetail