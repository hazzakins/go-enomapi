TP\GetOrderDetail
=================

Get information for a single domain on a transfer order.

Usage
-----

Use this command to retrieve a long list of information on a single domain that is in the process of transferring.

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
- The transfer order detail ID must belong to a transfer order created under this account.
- Input parameters

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter    | Status | Description                           | Max Size |
| --------------------- | ----------------------------- | --------------------------------------------------------------- | -------- |
| UID          | Required | Account login ID                        | 20 |
| PW | Required           | Account password | 20    |
| TransferOrderDetailID | Required | Transfer order detail number returned by calling TP\_GetOrder. | 10 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter   | Description |
| --------------------- | ------------------------------------------------------------------------------------------------ |
| TransferOrderDetailID | Transfer order detail number |
| SLD | Second level name                                        |
| TLD          | Top level name |
| Lock | Lock status of the name                                     |
| Renew         | Renew status of the name |
| DomainPassword | Password to be set for the name                                 |
| StatusID       | Status ID of this order. See Notes. |
| StatusDesc | Status description of this order. See Notes.                          |
| Price         | Charge amount for this order |
| UseContacts | Use original contacts or not                                   |
| Command        | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- TransferOrderDetail StatusID and StatusDesc:
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

The following query requests information on transfer order detail \(one item in a transfer order\) ID 301770, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderDetail&uid=resellid&pw=resellpw
&TransferOrderDetailID=301770&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderDetail&uid=resellid&pw=resellpw
&TransferOrderDetailID=301770&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderDetail&uid=resellid&pw=resellpw
&TransferOrderDetailID=301770&responsetype=text
```
The response indicates that tranfer order detail ID 301770 is for the transfer of domain resellerdocs2.net, and provides details of the order:

```html
<?xml version="1.0" ?>
<interface-response>
<transferorderdetail>
 <transferorderdetailid>301770</transferorderdetailid>
 <sld>resellerdocs2</sld>
 <tld>net</tld>
 <lock>True</lock>
 <renew>False</renew>
 <domainpassword />
 <statusid>9</statusid>
 <statusdesc>Awaiting auto verification of transfer request</statusdesc>
 <price>8.95</price>
 <usecontacts>0</usecontacts>
 <ordertype />
 <contacts>
 <Registrant>None</Registrant>
 <AuxBilling>None</AuxBilling>
 <Tech>None</Tech>
 <Admin>None</Admin>
 <Billing>
  <BillingPartyID>{BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}</BillingPartyID>
  <BillingAddress1>111 Main Street</BillingAddress1>
  <BillingAddress2 />
  <BillingCity>Hometown</BillingCity>
  <BillingCountry>US</BillingCountry>
  <BillingEmailAddress>[email protected]</BillingEmailAddress>
  <BillingFax>+1.5555555556</BillingFax>
  <BillingFirstName>John</BillingFirstName>
  <BillingJobTitle>President</BillingJobTitle>
  <BillingLastName>Doe</BillingLastName>
  <BillingOrganizationName>Reseller Documents Inc.</BillingOrganizationName>
  <BillingPhone>+1.5555555555</BillingPhone>
  <BillingPostalCode>99999</BillingPostalCode>
  <BillingStateProvince>WA</BillingStateProvince>
  <BillingStateProvinceChoice>S</BillingStateProvinceChoice>
  <BillingFullCountry>United States</BillingFullCountry>
  <UseRegInfoAbove>True</UseRegInfoAbove>
  <auxID />
  <regID />
 </Billing>
 </contacts>
</transferorderdetail>
<Command>TP_GETORDERDETAIL</Command>
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
transferorderdetailid: 301770
sld: resellerdocs2
tld: net
lock: True
renew: False
domainpassword:
statusid: 9
statusdesc: Awaiting auto verification of transfer request
price: 8.95
usecontacts: 0
ordertype:
Registrant: None
AuxBilling: None
Tech: None
Admin: None
BillingPartyID: {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}

BillingAddress1: 111 Main Street
BillingAddress2:
BillingCity: Hometown
BillingCountry: US
BillingEmailAddress: [email protected]
BillingFax: +1.5555555556
BillingFirstName: John
BillingJobTitle: President
BillingLastName: Doe
BillingOrganizationName: Reseller Documents Inc.

BillingPhone: +1.5555555555
BillingPostalCode: 99999
BillingStateProvince: WA
BillingStateProvinceChoice: S
BillingFullCountry: United States
UseRegInfoAbove: True
auxID:
regID:
Command: TP_GETORDERDETAIL
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
TrackingKey: a5d536c8-e18a-465b-b0e3-052a0283c026
RequestDateTime: 2/9/2015 12:43:55 PM
```
BillingAddress1=111 Main Street
BillingAddress2=
BillingCity=Hometown
BillingCountry=US
[email protected]
BillingFax=+1.5555555556
BillingFirstName=John
BillingJobTitle=President
BillingLastName=Doe
BillingOrganizationName=Reseller Documents Inc.

Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrder

TP\_GetOrderReview

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList