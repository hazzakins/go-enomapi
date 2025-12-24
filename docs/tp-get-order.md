TP\GetOrder
===========

Get a list of domains in a single transfer order.

Usage
-----

Use this command to retrieve status information on each item in a transfer order.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/TransferStatus.asp](https://resellertest.enom.com/domains/TransferStatus.asp)

On the transfer a name page, clicking a link in the Transfer Order ID column calls the TP\_GetOrder command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The transfer order ID must be valid.
- The transfer order must have originated from this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                              | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                           | 20 |
| PW | Required           | Account password | 20    |
| TransferOrderID | Required | Transfer order ID number. You can retrieve this number by calling the TP\_GetOrderStatuses command. | 10 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter   | Description |
| --------------------- | ------------------------------------------------------------------------------------------------ |
| TransferOrderID    | Transfer order number |
| OrderDate | Date the order was entered                                    |
| OrderTypeID      | Type ID of order |
| OrderTypeDesc | Type of order                                          |
| StatusID       | Status ID of order. See Notes. |
| StatusDesc | Status description of order. See Notes.                            |
| AuthAmount      | Charge amount for this order |
| TransferOrderDetailID | Transfer order detail number                                   |
| SLD          | Second-level domain name \(for example, enom in enom.com\) |
| TLD | Top-level domain name \(extension\)                               |
| Price         | Cost of the name transfer |
| UseContacts | Use original contacts or not                                   |
| Command        | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTMLor ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.
- TransferOrder StatusID:
*0 Transfer request has been made&\#xA;* 1 Fax has been received
*2 Order canceled&\#xA;* 3 Order complete
*4 Processing&\#xA;* 5 Order not submitted
- TransferOrderDetail StatusID:
*0 Transfer request created - awaiting fax&\#xA;* 1 WhoIs information matches
*2 Canceled due to WhoIs error&\#xA;* 3 Pending due to domain status
*4 Canceled due to domain status&\#xA;* 5 Transferred and paid successfully
*6 Transfer incomplete - charge problem&\#xA;* 7 Frozen due to charge problem
*8 NSI rejected transfer&\#xA;* 9 Awaiting auto verification of transfer request \(no longer used due to GDPR\)
*10 Unable to retrieve current domain contacts from UWhois \(no longer used due to GDPR\)&\#xA;* 11 Auto verification of transfer request initiated \(no longer used due to GDPR\)
*12 Awaiting for auto transfer string validation&\#xA;* 13 Domain awaiting transfer initiation
*14 Domain transfer initiated and awaiting approval&\#xA;* 15 Canceled - cannot obtain domain contacts from UWhois
*16 Canceled - domain contacts did not respond to verification e-mail&\#xA;* 17 Canceled - domain contacts did not approve transfer of domain
*18 Canceled - domain validation string is invalid&\#xA;* 19 Canceled - Whois information provided does not match current registrant
*20 Canceled - Domain is currently not registered and cannot be transferred&\#xA;* 21 Canceled - Domain is already registered in account and cannot be transferred
*22 Canceled - Domain is locked at current registrar, or is not yet 60 days old&\#xA;* 23 Canceled - Transfer already initiated for this domain
*24 Canceled - Unable to transfer due to unknown error&\#xA;* 25 Canceled - The current registrar has rejected transfer \(please contact them for details\)
*26 Canceled - Transfer authorization fax not received&\#xA;* 27 Canceled by customer
*28 Fax received - awaiting registrant verification&\#xA;* 29 Awaiting manual fax verification
*30 Canceled - Domain name is invalid or is Invalid for Transfers&\#xA;* 31 Canceled - Domain is currently undergoing transfer by another Registrar
*32 Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key&\#xA;* 33 Canceled - Cannot transfer domain from name-only account
*34 Unable to complete transfer. Transfers must include a change in registrar.&\#xA;* 35 Transfer request not yet submitted
*36 Canceled - Account is not authorized to perform domain transfers&\#xA;* 37 Canceled - Domain was not retagged or not retagged in time by losing registrar
*45 Order cancelled&\#xA;* 100 Pending consent - The transfer order is waiting for GDPR consent to be given.
\* 101 Canceled - Registrant denied - GDPR consent was refused, and so the order was cancelled.

Example
-------

The following query requests transfer information for transfer order ID 445413, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrder&uid=resellid&pw=resellpw
&TransferOrderID=445413&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrder&uid=resellid&pw=resellpw
&TransferOrderID=445413&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrder&uid=resellid&pw=resellpw
&TransferOrderID=445413&responsetype=text
```
The response provides information on transfer order ID 445413:

```html
<?xml version="1.0" ?>
<interface-response>
<transferorder>
 <transferorderid>445413</transferorderid>
 <orderdate>7/29/2002 5:19:29 PM</orderdate>
 <ordertypeid>1</ordertypeid>
 <ordertypedesc>Auto Verification</ordertypedesc>
 <statusid>4</statusid>
 <statusdesc>Processing</statusdesc>
 <authamount>17.90</authamount>
 <version>1</version>
 <transferorderdetail>
 <transferorderdetailid>301770</transferorderdetailid>
 <sld>resellerdocs2</sld>
 <tld>net</tld>
 <statusid>9</statusid>
 <statusdesc>Awaiting auto verification of transfer request</statusdesc>
 <price>8.95</price>
 <usecontacts>0</usecontacts>
 </transferorderdetail>
 <transferorderdetail>
 <transferorderdetailid>301771</transferorderdetailid>
 <sld>resellerdocs3</sld>
 <tld>info</tld>
 <statusid>9</statusid>
 <statusdesc>Awaiting auto verification of transfer request</statusdesc>
 <price>8.95</price>
 <usecontacts>0</usecontacts>
 </transferorderdetail>
 <transferorderdetailcount>2</transferorderdetailcount>
</transferorder>
<Command>TP_GETORDER</Command>
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
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
transferorderid: 445413
orderdate: 7/29/2002 5:19:29 PM
ordertypeid: 1
ordertypedesc: Auto Verification
statusid: 4
statusdesc: Processing
authamount: 17.90
version: 1
transferorderdetailid: 301770
sld: resellerdocs2
tld: net
statusid: 9
statusdesc: Awaiting auto verification of transfer request
price: 8.95
usecontacts: 0
transferorderdetailid: 301771
sld: resellerdocs3
tld: info
statusid: 9
statusdesc: Awaiting auto verification of transfer request
price: 8.95
usecontacts: 0
transferorderdetailcount: 2
Command: TP_GETORDER
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
ExecTime: 0.031
Done: true
TrackingKey: 9d3fa45d-648c-461a-ac95-bcf2ead5a8ad
RequestDateTime: 2/9/2015 12:34:28 PM
```
Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrderDetail

TP\_GetOrderReview

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList