TP\CancelOrder
==============

Cancel a transfer order that has been submitted by eNom, but not yet processed by the Registry.

Usage
-----

Use this command to cancel the parts of a transfer order that has been submitted by us, but not yet processed by the Registry \(there is an interval of approximately 5 days between the two events\).

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/TransferStatus.asp?transferorderid=445425](https://resellertest.enom.com/domains/TransferStatus.asp?transferorderid=445425)

On the Transfer Order Detail page, the Cancel This Order link calls the TP\_CancelOrder command.

\[[https://resellertest.enom.com/domains/TransferOrderReview.asp](https://resellertest.enom.com/domains/TransferOrderReview.asp)?

TransferOrderID=445428&OrderType=Auto\+Verification&RegistrantFirstName=&

RegistrantLastName=\]\([https://resellertest.enom.com/domains/TransferOrderReview.asp](https://resellertest.enom.com/domains/TransferOrderReview.asp)?

RegistrantLastName=\)

On the Pending orders page, the cancel button calls the TP\_CancelOrder command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The transfer order must be valid and must belong to this account.
- A transfer order can be cancel only for domains that have a TransferOrderDetail StatusID of 0, 9, 10,11, 12, 13, 28, or 29. Use TP\_GetOrderDetail to retrieve the statuses of each domain in the order.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                        | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                     | 20 |
| PW | Required           | Account password | 20    |
| TransferOrderID | Required | Transfer order ID number. To retrieve this value, use the TP\_GetOrderStatuses command. | 10 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Success     | Returns True if transfer request was successfully cancelled |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- To be eligible for cancellation, each domain must be in one of the following statuses. Retrieve an order’s status by calling TP\_GetOrder. Eligible TransferOrderDetail StatusIDs:
 - ```
0 Transfer request created - awaiting fax
```
 - ```
9 Awaiting auto verification of transfer request
```
 - ```
10 Unable to retrieve current domain contacts from UWhois
```
 - ```
11 Auto verification of transfer request initiated
```
 - ```
12 Awaiting for auto transfer string validation
```
 - ```
13 Domain awaiting transfer initiation
```
 - ```
28 Fax received - awaiting registrant verification
```
 - ```
29 Awaiting manual fax verification.
```
Example
-------

The following query requests that transfer order ID 445413 be canceled, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_CancelOrder&uid=resellid&pw=resellpw
&transferOrderID=445413&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_CancelOrder&uid=resellid&pw=resellpw
&transferOrderID=445413&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_CancelOrder&uid=resellid&pw=resellpw
&transferOrderID=445413&responsetype=text
```
In the response, the return value True for the Success parameter confirms the successful cancellation of the order:

```
<?xml version="1.0" ?>
<interface-response>
<success>True</success>
<Command>TP_CANCELORDER</Command>
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
success: True
Command: TP_CANCELORDER
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site:
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.094
Done: true
RequestDateTime: 2/6/2015 1:17:59 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
success=True
Command=TP_CANCELORDER
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
ExecTime=0.063
Done=true
RequestDateTime=2/6/2015 1:18:46 PM
```
Related Commands
----------------

PushDomain

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrder

TP\_GetOrderDetail

TP\_GetOrderReview

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList