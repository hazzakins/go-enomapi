InsertNewOrder
==============

Get a list of orders including Closed for the last 6 months.

Usage
-----

Use this command if you use our shopping cart, to finalize purchase of the cart contents.

This command checks out the items in a shopping cart that are in A \(active\) status, and puts the order in a queue.

Purchase, a similar command, completes the purchase of selected products in real time, without using our shopping cart or waiting for the order processing queue.

Because of the high prices for some Premium Domains, we offer you a wider variety of payment methods than for most purchases. In addition to immediate payment using your account balance, as you do for other domains and services, we offer two additional payment options for high-value Premium Domains: wire transfer, or delayed payment using your account balance. Use the UseWireTransfer parameter in the InsertNewOrder command to signal that you want to pay later, then either wire your payment or use the NM\_ProcessOrder command to pay using your account balance.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- If the shopping cart is empty, ItemCount must equal 1 or the query will return an error message.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=
InsertNewOrder&uid=(Required)&pw=(Required)&EndUserIP=(Required)&responsetype=xml
```
| Input Parameter | Type | Status                                 | Description |
| --------------- | ------ | ---------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | Required                                | TP\_GetOrderStatuses |
| uid | string | Required | Your Account ID                                                                                                                                                                                            |
| pw       | string | Required                                | Your API Token |
| EndUserIP | string | Required | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN.                                                                                                                                   |
| ItemCount    | string | Optional                                | Number of items ready to purchase. All with a status of A \(active\). To retrieve the statuses of items in the shopping cart, use the GetCartContent command |
| UseWireTransfer | string | Optional overall, but Required defer payment of a Premium Domain order | If UseWireTransfer=Yes, process this Premium Domain order as a deferred payment order. Options for later payment are wire transfer or account balance. If you pay by wire transfer, contact your sales representative for wire transfer instructions. If you pay using your account balance \(but after submitting the order using the InsertNewOrder command\), use the NM\_ProcessOrder command. |
| ResponseType  | string | Optional                                | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| -------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| PremiumDomainOrderID | int   | Identification number for the order that includes all Premium Domains and the services directly associated with those Premium Domains. |
| OrderID | int   | Identification number for the order that includes all items not associated with Premium Domains |
| ProdTypeX | string | Product type identification number for domain names |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=insertneworder&uid=resellid&pw=resellpw
&UseWireTransfer=Yes&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=insertneworder&uid=resellid&pw=resellpw
&UseWireTransfer=Yes&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=insertneworder&uid=resellid&pw=resellpw
&UseWireTransfer=Yes&responsetype=text
```
```
<interface-response>
 <OrderID>157781016</OrderID>
 <ProdType10>10</ProdType10>
 <Command>INSERTNEWORDER</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>4.922</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/9/2011 3:24:56 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>OrderID: </ STRONG>161808106<br>
<STRONG>ProdType10: </ STRONG>10<br>
<STRONG>Command: </ STRONG>INSERTNEWORDER<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+07.00<br>
<STRONG>ExecTime: </ STRONG>4.250<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/4/2015 2:11:23 PM<br>
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command=INSERTNEWORDER
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.250
Done=true
RequestDateTime=2/4/2015 2:11:43 PM
```
Related Commands
----------------

AddBulkDomains

AddToCart

Check

DeleteFromCart

Extend

Extend\_RGP

ExtendDomainDNS

GetDomainExp

GetExtendInfo

GetRenew

Purchase

PurchasePreview

SetRenew

UpdateCart

UpdateExpiredDomains