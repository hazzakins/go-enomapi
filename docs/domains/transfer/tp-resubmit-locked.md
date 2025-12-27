TP\ResubmitLocked
=================

Resubmit a request for a previously locked domain name.

Usage
-----

Use this command after you have previously submitted a transfer request, received an error message saying the domain was locked, and have had the registrant of the domain unlock it. This command may also be used to resubmit a transfer that was rejected because the domain was registered less than 60 days ago.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

On the Transfer order detail page, the Resubmit locked domain link \(when it is present\) calls the TP\_ResubmitLocked command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The transfer order detail ID must be valid.
- The transfer order detail ID must belong to a domain that previously failed to transfer because it was locked.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter    | Status | Description                                    | Max Size |
| --------------------- | ----------------------------- | --------------------------------------------------------------------------------- | -------- |
| UID          | Required | Account login ID                                 | 20 |
| PW | Required           | Account password | 20    |
| TransferOrderDetailID | Required | Transfer Order Detail ID. To retrieve this value, use the TP\_GetOrder command. | 10 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| TransferOrderID | New Transfer Order ID |
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

Previously, account resellid had requested that a domain be transferred into this account, and received an e-mail stating that the transfer could not complete because the domain was locked. The owner of account resellid contacted the owner of the domain and had the lock removed. The following query resubmits the transfer request, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_ResubmitLocked&uid=resellid&pw=resellpw
&TransferOrderDetailID=607291&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_ResubmitLocked&uid=resellid&pw=resellpw
&TransferOrderDetailID=607291&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_ResubmitLocked&uid=resellid&pw=resellpw
&TransferOrderDetailID=607291&responsetype=text
```
The Success=True parameter value in the response, and the new TransferOrderID, confirm a successful request:

```html
<?xml version="1.0" ?>
<interface-response>
<transferorder>
 <transferorderid>952245</transferorderid>
</transferorder>
<success>True</success>
<Command>TP_RESUBMITLOCKED</Command>
<ErrCount>0</ErrCount>
<Server>ResellerTest</Server>
<Site>enom</Site>
<Done>true</Done>
<debug>
 <![ CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
success: True
Command: TP_RESUBMITLOCKED
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site:
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.078
Done: true
RequestDateTime: 2/9/2015 1:09:47 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
success=True
Command=TP_RESUBMITLOCKED
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.078
Done=true
RequestDateTime=2/9/2015 1:10:33 PM
```
Related Commands
----------------

PushDomain

TP\_CancelOrder

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrder

TP\_GetOrderDetail

TP\_GetOrderReview

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList