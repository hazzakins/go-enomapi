TP\GetOrderStatuses
===================

Get a list of orders including Closed for the last 6 months.

Usage
-----

Use this command to list all transfer orders for an account ID.

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
https://resellertest.enom.com/interface.asp?command=TP_GetOrderStatuses&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter   | Type | Status  | Description |
| ------------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command       | string | Required | TP\_GetOrderStatuses |
| uid | string | Required | Your Account ID                              |
| pw         | string | Required | Your API Token |
| IncludeClosedOrders | int  | Optional | Set =1 to return closed orders in the list.               |
| ResponseType    | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

TransferOrderDetail StatusID is the status of this domain in the transfer process:

| Output Parameter | Type  | Description |
| ---------------- | ------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| TransferOrderID | int   | Transfer order number |
| OrderDate | string | Date the order was entered |
| OrderTypeID | string | Type ID of order. |
| OrderTypeDesc | string | Type of order |
| StatusID | string | Status ID of transfer order. Options: 0 Transfer request has been made 1 Fax has been received 2 Order canceled 3 Order complete 4 Processing 5 Order not submitted |
| StatusDesc | string | Status description of transfer order. See text descriptions of statuses in StatusID parameter, above. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderStatuses&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderStatuses&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderStatuses&uid=resellid
&pw=resellpw&responsetype=text
```
```html
<?xml version="1.0" ?>
<interface-response>
 <transferorder>
 <transferorderid>445413</transferorderid>
 <orderdate>2002-07-29T17:19:29.437</orderdate>
 <statusid>4</statusid>
 <tos>
  <statusdesc>Processing</statusdesc>
  <ordertypeid>1</ordertypeid>
  <tot>
  <ordertypedesc>Auto Verification</ordertypedesc>
  </tot>
 </tos>
 </transferorder>
 <transferorder>
.
.
.
 </transferorder>
 <FoundAtLeastOnePendingOrder>False</FoundAtLeastOnePendingOrder>
 <Command>TP_GETORDERSTATUSES</Command>
 <ErrCount>0</ErrCount>
 <Server>Dev Workstation</Server>
 <Site>enom</Site>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>orderid1: </STRONG>175581269<br />
<STRONG>TransferOrderCount: </STRONG>1000<br />
<STRONG>transferorderid1: </STRONG>175696133<br />
<STRONG>orderdate1: </STRONG>2/6/2015 1:23:29 PM<br />
<STRONG>statusid1: </STRONG>4<br />
<STRONG>statusdesc1: </STRONG>Processing<br />
<STRONG>ordertypeid1: </STRONG>1<br />
<STRONG>ordertypedesc1: </STRONG>Auto Verification<br />
<STRONG>transferinitdate1:</STRONG><br />
<STRONG>transferorderid2: </STRONG>175696130<br />
<STRONG>orderdate2: </STRONG>2/6/2015 1:22:47 PM<br />
<STRONG>statusid2: </STRONG>4<br />
<STRONG>statusdesc2: </STRONG>Processing<br />
<STRONG>ordertypeid2: </STRONG>1<br />
<STRONG>ordertypedesc2: </STRONG>Auto Verification<br />
<STRONG>transferinitdate2:</STRONG><br />
<STRONG>transferorderid3: </STRONG>175695689<br />
<STRONG>orderdate3: </STRONG>1/16/2015 2:02:53 AM<br />
<STRONG>statusid3: </STRONG>3<br />
<STRONG>statusdesc3: </STRONG>Order complete<br />
<STRONG>ordertypeid3: </STRONG>1<br />
<STRONG>ordertypedesc3: </STRONG>Auto Verification<br />
<STRONG>transferinitdate3:</STRONG><br />
.
.
.
<STRONG>count: </STRONG>1000<br />
<STRONG>FoundAtLeastOnePendingOrder: </STRONG>True<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>Command: </STRONG>TP_GETORDERSTATUSES<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable:</STRONG><br />
<STRONG>IsRealTimeTLD:</STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.781<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>3c14c5e1-88dd-49da-bf28-403353ae929b<br />
<STRONG>RequestDateTime: </STRONG>2/9/2015 1:00:28 PM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
TransferOrderCount=1000
transferorderid1=175696133
orderdate1=2/6/2015 1:23:29 PM
statusid1=4
statusdesc1=Processing
ordertypeid1=1
ordertypedesc1=Auto Verification
transferinitdate1=
transferorderid2=175696130
orderdate2=2/6/2015 1:22:47 PM
statusid2=4
statusdesc2=Processing
ordertypeid2=1
ordertypedesc2=Auto Verification
transferinitdate2=
transferorderid3=175695689
orderdate3=1/16/2015 2:02:53 AM
statusid3=3
statusdesc3=Order complete
ordertypeid3=1
ordertypedesc3=Auto Verification
transferinitdate3=
.
.
.
count=1000
FoundAtLeastOnePendingOrder=True
ErrCount=0
ResponseCount=0
Command=TP_GETORDERSTATUSES
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
ExecTime=0.391
Done=true
TrackingKey=efcd8bb9-f770-4478-b27f-7bbf0a4140cf
RequestDateTime=2/9/2015 1:01:20 PM
```
Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

TP\_CreateOrder

TP\_GetOrder

TP\_GetOrderDetail

TP\_GetOrderReview

TP\_GetOrdersByDomain

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList