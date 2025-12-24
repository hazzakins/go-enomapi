TP\GetOrdersByDomain
====================

Retrieve transfer information for a domain name.

Usage
-----

Use this command to list a history and status of transfer orders for a domain name.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.    | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                                                                            |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| OrderCount | Number of transfer orders that have been submitted for this domain. TransferOrderIDX                                                                       |
| X=1 to OrderCount | Transfer order detail number. Indexed X when ResponseType=Text or HTML. |
| LoginIDX | The account this transfer order is in. Indexed X when ResponseType=Text or HTML.                                                                         |
| OrderDateX    | The date the order was submitted. Indexed X when ResponseType=Text or HTML. |
| OrderStatusX | Current status of the order. Indexed X when ResponseType=Text or HTML.                                                                              |
| StatusIDX     | TransferOrder StatusID number. Options are: 0 New 1 Authorization Succeeded 2 Authorization Failed 3 Processing 4 Order queued 5 Ready for Billing 6 Order complete 7 Order canceled. Indexed X when ResponseType=Text or HTML. |
| Command | Name of command executed                                                                                                      |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                                                     |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests transfer order information for domain name resellerdocs2.net, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrdersByDomain&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrdersByDomain&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrdersByDomain&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&responsetype=text
```
The response indicates that resellerdocs2.net is in transfer order ID 445413, initiated by account resellid:

```html
<?xml version="1.0" ?>
<interface-response>
<TransferOrder>
 <transferorderid>445413</transferorderid>
 <loginid>resellid</loginid>
 <orderdate>7/29/2002 5:19:29 PM</orderdate>
 <orderstatus>Processing</orderstatus>
 <statusid>4</statusid>
</TransferOrder>
<ordercount>1</ordercount>
<Command>TP_GETORDERSBYDOMAIN</Command>
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
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
transferorderid1: 175696133
loginid1: resellid
orderdate1: 2/6/2015 1:23:29 PM
orderstatus1: Processing
statusid1: 4
transferorderid2: 175696130
loginid2: resellid
orderdate2: 2/6/2015 1:22:47 PM
orderstatus2: Processing
statusid2: 4
transferorderid3: 175694255
loginid3: resellid
orderdate3: 11/30/2014 10:36:47 PM
orderstatus3: Order complete
statusid3: 3
.
.
.
ordercount: 88
ErrCount: 0
ResponseCount: 0
Command: TP_GETORDERSBYDOMAIN
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +8.00
ExecTime: 0.078
Done: true
TrackingKey: 73935cc5-65e3-4395-b25f-782b5b5029c4
RequestDateTime: 2/9/2015 12:57:27 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
transferorderid1=175696133
loginid1=resellid
orderdate1=2/6/2015 1:23:29 PM
orderstatus1=Processing
statusid1=4
transferorderid2=175696130
loginid2=resellid
orderdate2=2/6/2015 1:22:47 PM
orderstatus2=Processing
statusid2=4
transferorderid3=175694255
loginid3=resellid
orderdate3=11/30/2014 10:36:47 PM
orderstatus3=Order complete
statusid3=3
.
.
.
ordercount=88
ErrCount=0
ResponseCount=0
Command=TP_GETORDERSBYDOMAIN
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
TrackingKey=914b969c-ecd9-4773-b5a1-b6c3a27a293a
RequestDateTime=2/9/2015 12:58:16 PM
```
Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrder

TP\_GetOrderDetail

TP\_GetOrderReview

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList