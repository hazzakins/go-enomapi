Queue\GetOrderDetail
====================

Retrieve a detailed queue domain order.

Usage
-----

Use this command to retrieve a detailed queue domain order.

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
| Input Parameter  | Status | Description                                                              | Max Size |
| ----------------- | ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID        | Required | Account login ID                                                           | 20 |
| PW | Required           | Account password | 20    |
| OrderID      | Required | The identification number of the order for which you want details. You can retrieve this number using the Queue\_GetOrders command. | 10 |
| PortalUserPartyID | Optional           | Portal user identification |     |
| ResponseType   | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                         | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| OrderID | Order ID                                             |
| OrderDate    | Order Date |
| PaidAmount | Amount charged for this transaction                               |
| DomainName    | Domain Name |
| OrderType | Order type. In a queue process, this value will be 'Register'                  |
| ExtTime     | Registration time applied to this item |
| OrderStatus | Order status                                           |
| RegistrationFee | Registration Fee |
| ApplicationFee | Application Fee                                         |
| TotalFee     | Total fee for this item |
| RecordStart | Record start position                                      |
| PagingSize    | Record size per page |
| NextRecord | Record next position                                       |
| TotalResults   | Total records for this account |
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

The following query retrieve a list of queue domain order history, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&OrderID=157896057
&Command=Queue_GetOrderDetail&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&OrderID=157896057
&Command=Queue_GetOrderDetail&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
UID=ResellID&PW=resellpw&OrderID=157896057
&Command=Queue_GetOrderDetail&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Queue_GetOrderDetail>
 <Orders>
 <Order>
  <OrderID>157896057</OrderID>
  <OrderDate>4/11/2013 5:11 PM</OrderDate>
  <PaidAmount>54.00</PaidAmount>
  <DomainName>superdomain.ninja</DomainName>
  <OrderType>Register</OrderType>
  <ExtTime>1</ExtTime>
  <OrderStatus>Pending</OrderStatus>
  <RegistrationFee>32.00</RegistrationFee>
  <ApplicationFee>22.00</ApplicationFee>
  <TotalFee>54.00</TotalFee>
 </Order>
 </Orders>
 <RecordStart>1</RecordStart>
 <PagingSize>25</PagingSize>
 <NextRecord>1</NextRecord>
 <TotalResults>1</TotalResults>
</Queue_GetOrderDetail>
<Command>QUEUE_GETORDERDETAIL</Command>
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
<RequestDateTime>6/5/2013 3:40:56 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Orders:
OrderID: 157896057
OrderDate: 4/11/2013 5:11 PM
PaidAmount: 54.00
DomainName: superdomain.ninja
OrderType: Register
ExtTime: 1
OrderStatus: Pending
RegistrationFee: 32.00
ApplicationFee: 22.00
TotalFee: 54.00
RecordStart: 1
PagingSize: 25
NextRecord: 1
TotalResults: 0
Command: QUEUE_GETORDERDETAIL
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
ExecTime: 0.797
Done: true
TrackingKey: 585202de-aafd-4dac-9a70-337e937a2e88
RequestDateTime: 2/5/2015 11:36:53 AM
```
Related Commands
----------------

GetAgreementPage

PE\_GetPremiumPricing

Queue\_DomainPurchase

Queue\_GetDomains

Queue\_GetExtAttributes

Queue\_GetInfo

Queue\_GetOrders