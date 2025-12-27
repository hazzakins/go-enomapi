TP\UpdateOrderDetail
====================

Update a preconfigured transfer order before submitting the order.

Usage
-----

Use this command to configure registrar lock, auto-renew, domain password, or contact information for a domain name that is in the process of being transferred to eNom.

This command is typically used after an order has been created using the TP\_CreateOrder command with the parameter PreConfig=1. Once the transfer order has been configured using this command, the next step is typically to use the TP\_SubmitOrder command, which causes the transfer order to be processed.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/TransferNew.asp](https://resellertest.enom.com/domains/TransferNew.asp)

Enter a domain name and click next.

On the Pending orders page, click the domain name.

In the Editing transfer order for domain name box, the check boxes set the parameter values for the TP\_UpdateOrderDetail command, and the next button calls the command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The transfer order detail ID must be valid and must be part of an order created under this account.
- The transfer order must be in a StatusID of 5 \(Order not submitted\).

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter        | Status | Description                                                                                                                     | Max Size |
| ----------------------------- | -------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID              | Required | Account login ID                                                                                                                   | 20 |
| PW | Required                                                 | Account password | 20    |
| TransferOrderDetailID     | Required | Transfer order detail ID. Use the TP\_GetOrder command to retrieve this value.                                                                                   | 10 |
| Lock | Optional; default is True                                        | Set registrar lock on the name. Permitted values are True or False. | 5    |
| Renew             | Optional; default is True | Set auto-renew. Permitted values are True or False.                                                                                                | 5 |
| DomainPassword | Optional                                                 | Set a domain password on the name | 60    |
| DeleteOrderDetail       | Optional | Remove all names from this order,. To remove all names, supply DeleteOrderDetail=Yes                                                                                | 3 |
| RegistrantAddress1 | Optional; Required Fax orders                                      | Registrant address | 60    |
| RegistrantAddress2      | Optional | Registrant additional address info                                                                                                          | 60 |
| RegistrantCity | Optional; Required Fax orders                                      | Registrant city | 60    |
| RegistrantCountry       | Optional; Required Fax orders | Registrant country                                                                                                                  | 60 |
| RegistrantEmailAddress | Optional; Required Fax orders and for .org names for which you are supplying new Registrant information | Registrant email address | 128   |
| RegistrantFax         | Optional | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). .                             | 20 |
| RegistrantFirstName | Optional; Required Fax orders                                      | Registrant first name | 60    |
| RegistrantLastName      | Optional; Required Fax orders | Registrant last name                                                                                                                 | 60 |
| RegistrantJobTitle | Optional                                                 | Registrant job title | 60    |
| RegistrantOrganizationName  | Optional; Required Fax orders | Registrant organization                                                                                                               | 60 |
| RegistrantPhone | Optional; Required Fax orders                                      | Registrant phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URLencoded as a plus sign \(%2B\). . | 20    |
| RegistrantPostalCode     | Optional | Registrant postal code                                                                                                                | 16 |
| RegistrantStateProvince | Optional                                                 | Registrant state or province | 60    |
| RegistrantStateProvinceChoice | Optional | Registrant state or province choice: S state P province                                                                                               | 1 |
| UseContacts | Optional; default is 0                                          | Set =1 to transfer existing Whois contacts when the transfer is complete. If you supply UseContacts=0 for a domain that has extended attributes, you must supply contact information and extended attributes before submitting the transfer order. | 1    |
| ExtendedAttributes      | Required | TLDs that use extended attributes, when UseContacts=0 Extended attributes, required for some country code TLDs. You can retrieve a list of required extended attributes for any TLD using the GetExtAttributes command.              | n/a |
| RegistrantUseContact | Optional                                                 | If set =none then Billing contact will be used, otherwise Registrant contact data must be submitted | 3    |
| Tech\*            | Optional | Technical contact data. See Note.                                                                                                         | - |
| AuxBilling\* | Optional                                                 | Auxiliary Billing contact data. See Note. | -    |
| TechUseContact        | Optional | If set =none then no contact will be used, otherwise Tech contact data must be submitted.                                                                              | 3 |
| AuxBillingUseContact | Optional                                                 | If set =none then no contact will be used, otherwise Auxbilling contact data must be submitted. | 3    |
| ResponseType         | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                                                                                | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| --------------------- | ------------------------------------------------------------------------------------------------ |
| TransferOrderDetailID | Transfer order detail number                                   |
| SLD          | Second-level domain name |
| TLD | Top-level domain name                                      |
| lock         | Registrar lock status |
| renew | Auto-renew status                                        |
| domainpassword    | Password to set for the domain name |
| statusid | Status ID of the order                                      |
| statusdesc      | Description of the status |
| price | Transfer price                                          |
| usecontacts      | Use the current contacts flag |
| Registrant\* | Registrant contact data                                     |
| AuxBilling\*     | Auxilliary Billing contact data |
| Tech\* | Technical contact data                                      |
| Admin\*        | Administrative contact data |
| Billing\* | Billing contact data                                       |
| Command        | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed.
- The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- To update Technical or Auxilliary Billing contacts submit the contact data by replacing Registrant *param names with Tech* or AuxBilling\* param names.

Example
-------

The following query requests that for transfer order detail ID 301770, the registrar lock and automatic renew options both be set to On. It also instructs that contact information should remain unchanged, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_UpdateOrderDetail&uid=resellid&pw=resellpw
&TransferOrderDetailID=301770&Lock=On&Renew=On
&RegistrantUseContact=none&TechUseContact=none
&AuxBillingUseContact=none&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TP_UpdateOrderDetail&uid=resellid&pw=resellpw
&TransferOrderDetailID=301770&Lock=On&Renew=On
&RegistrantUseContact=none&TechUseContact=none
&AuxBillingUseContact=none&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_UpdateOrderDetail&uid=resellid&pw=resellpw
&TransferOrderDetailID=301770&Lock=On&Renew=On
&RegistrantUseContact=none&TechUseContact=none
&AuxBillingUseContact=none&responsetype=text
```
The response indicates the successful setting of the registrar lock and automatic renew options:

```html
<?xml version="1.0" ?>
<interface-response>
<transferorderdetail>
 <transferorderdetailid>301770</transferorderdetailid>
 <sld>resellerdocs2</sld>
 <tld>net</tld>
 <lock>True</lock>
 <renew>True</renew>
 <domainpassword />
 <statusid>10</statusid>
 <statusdesc>Unable to retrieve current domain contacts from UW</statusdesc>
 <price>8.95</price>
 <usecontacts>0</usecontacts>
 <ordertype />
 <contacts>
 <Registrant>
  <RegistrantPartyID>
  {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
  </RegistrantPartyID>
  <RegistrantAddress1>111 Main Street</RegistrantAddress1>
  <RegistrantAddress2 />
  <RegistrantCity>Hometown</RegistrantCity>
  <RegistrantCountry>US</RegistrantCountry>
  <RegistrantEmailAddress>[email protected]</RegistrantEmailAddress>
  <RegistrantFax>555-555-5556</RegistrantFax>
  <RegistrantFirstName>John</RegistrantFirstName>
  <RegistrantJobTitle>President</RegistrantJobTitle>
  <RegistrantLastName>Doe</RegistrantLastName>
  <RegistrantOrganizationName>Reseller Documents Inc.</RegistrantOrganizationName>
  <RegistrantPhone>555-555-5555</RegistrantPhone>
  <RegistrantPostalCode>99999</RegistrantPostalCode>
  <RegistrantStateProvince>WA</RegistrantStateProvince>
  <RegistrantStateProvinceChoice>S</RegistrantStateProvinceChoice>
 </Registrant>
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
  <BillingFax>555-555-5556</BillingFax>
  <BillingFirstName>John</BillingFirstName>
  <BillingJobTitle>President</BillingJobTitle>
  <BillingLastName>Doe</BillingLastName>
  <BillingOrganizationName>Reseller Documents Inc.</BillingOrganizationName>
  <BillingPhone>555-555-5555</BillingPhone>
  <BillingPostalCode>99999</BillingPostalCode>
  <BillingStateProvince>WA</BillingStateProvince>
  <BillingStateProvinceChoice>S</BillingStateProvinceChoice>
  <BillingFullCountry>United States</BillingFullCountry>
  <UseBelowAuxInfo>True</UseBelowAuxInfo>
  <auxID />
  <regID>{BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}</regID>
 </Billing>
 </contacts>
</transferorderdetail>
<success>True</success>
<Command>TP_UPDATEORDERDETAIL</Command>
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
renew: True
domainpassword:
statusid: 10
statusdesc: Unable to retrieve current domain contacts from UW
price: 8.95
usecontacts: 0
ordertype:
RegistrantPartyID: {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
RegistrantAddress1: 111 Main Street
RegistrantAddress2:
RegistrantCity: Hometown
RegistrantCountry: US
RegistrantEmailAddress: [email protected]
RegistrantFax: 555-555-5556
RegistrantFirstName: John
RegistrantJobTitle: President
RegistrantLastName: Doe
RegistrantOrganizationName: Reseller Documents Inc.
RegistrantPhone: 555-555-5555
RegistrantPostalCode: 99999
RegistrantStateProvince: WA
RegistrantStateProvinceChoice: S

AuxBilling: None
Tech: None
Admin: None
BillingPartyID: {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
BillingAddress1: 111 Main Street
BillingAddress2:
BillingCity: Hometown
BillingCountry: US
BillingEmailAddress: [email protected]
BillingFax: 555-555-5556
BillingFirstName: John
BillingJobTitle: President
BillingLastName: Doe
BillingOrganizationName: Reseller Documents Inc.
BillingPhone: 555-555-5555
BillingPostalCode: 99999
BillingStateProvince: WA
BillingStateProvinceChoice: S
BillingFullCountry: United States
UseBelowAuxInfo: True
auxID:
regID: {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
success: True
Command: TP_UPDATEORDERDETAIL
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
ExecTime: 0.047
Done: true
RequestDateTime: 2/9/2015 1:16:10 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
transferorderdetailid=301770
sld=resellerdocs2
tld=net
lock=True
renew=True
domainpassword=
statusid=10
statusdesc=Unable to retrieve current domain contacts from UW
price=8.95
usecontacts=0
ordertype=
RegistrantPartyID={BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
RegistrantAddress1= 111 Main Street
RegistrantAddress2=
RegistrantCity=Hometown
RegistrantCountry=US
[email protected]
RegistrantFax=555-555-5556
RegistrantFirstName=John
RegistrantJobTitle=President
RegistrantLastName=Doe
RegistrantOrganizationName=Reseller Documents Inc.
RegistrantPhone=555-555-5555
RegistrantPostalCode=99999
RegistrantStateProvince=WA
RegistrantStateProvinceChoice=S

AuxBilling=None
Tech=None
Admin=None
BillingPartyID={BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
BillingAddress1=111 Main Street
BillingAddress2=
BillingCity=Hometown
BillingCountry=US
[email protected]
BillingFax=555-555-5556
BillingFirstName=John
BillingJobTitle=President
BillingLastName=Doe
BillingOrganizationName=Reseller Documents Inc.
BillingPhone=555-555-5555
BillingPostalCode=99999
BillingStateProvince=WA
BillingStateProvinceChoice=S
BillingFullCountry=United States
UseBelowAuxInfo=True
auxID=
regID={BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
success=True
Command=TP_UPDATEORDERDETAIL
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
ExecTime=0.047
Done=true
RequestDateTime=2/9/2015 1:16:10 PM
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

TP\_ResubmitLocked

TP\_SubmitOrder

UpdatePushList