TP\CreateOrder
==============

Transfer domains into an account. Accepts Fax and AutoVerification order types.

Usage
-----

> ### Asynchronous Transfers
>
>
>
> If the "statusdesc" output parameter is "Pending consent", or the command returns a "statusid" of 100, additional action is required by the registrant to complete the transfer. For more information about consent-related statuses, please refer to our [GDPR](../docs/gdpr.md) landing page.

Use this command to create an order to transfer domains from another registrar to eNom or one of its resellers.

You can also use this command to create but not submit a transfer order, using the PreConfig parameter.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

When you transfer a .eu or .be domain name from another registrar, we recommend that you always provide Registrant contact information that is separate from Billing contact information; don’t use the “same as Billing” default. Tip: If the Billing and Registrant contact information are the same, we recommend changing the use or spelling of abbreviations in the street address to help our system recognize that it needs to create multiple contacts. Separating the Registrant and Billing information makes it easier to update Registrant contact information in the future.

> ### CCTLD Transfers
>
>
>
> Many Country Code TLD's require the creation of a new registrant contact during the creation of an incoming transfer order. Notable examples are **.EU** and **.CA.**
>
>
>
> When passing a new Registrant's contact information, all Registrant fields must be supplied.

Premium Domain
--------------

A unique, memorable or noteworthy name might be considered as a premium domain. This type of domain will be more expensive depending on the name and TLD. If the TP\_CreateOrder command returns the error "CustomerSuppliedPrice is required", then this domain is considered as a premium domain. The exact price for this domain must be declared in the parameter CustomerSuppliedPrice.

Use the PE\_GetPremiumPricing command to retrieve this price.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- To use our credit card processing, this must be an ETP reseller account.
- All domain names in the order must be in top-level domains supported by this registrar.
- To transfer EPP names, the query must include the authorization key from the Registrar.
- When using the Fax order type, the registrant contact information must match the current Whois registrant.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=TP_CreateOrder&uid=(Required)&pw=(Required)&OrderType=(Required)&DomainCount=(Required)&SLD1=(Required)&TLD1=(Required)&AuthInfo1=(Required)&responsetype=(Optional)
```
| Input Parameter        | Type | Status                                                                     | Description |
| ----------------------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command            | string | Required                                                                    | TP\_CreateOrder |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                           |
| pw              | string | Required                                                                    | Your API Token |
| PreConfig | string | Optional; default is 0 | Set PreConfig=1 to create, but not send, the order at this time. Using PreConfig=1 allows you to use TP\_UpdateOrderDetail to modify DomainPassword, Lock, Renew, contacts, and extended attributes before submitting the order. When you are ready to submit the order, call TP\_SubmitOrder.                                                                                  |
| OrderType           | string | Required                                                                    | Permitted values are Fax or Autoverification |
| CustomerSuppliedPrice | int   | Optional | For premium domain\(s\), the exact price for this domain must be declared in this parameter. Use the PE\_GetPremiumPricing command to retrieve the price.                                                                                                                                                     |
| IncludeIDP          | int | Optional                                                                    | Includes ID Protect on domain transfers for TLDs that support ID Protect. Permitted values are 1 = include ID Protect, and 0 = Do not include ID Protect |
| IDPPrice | int   | Optional; required if IncludeIDP = 1 | Price of ID Protect \(the price can be retrieved from the appropriate price command\)                                                                                                                                                                                        |
| DomainCount          | int | Required                                                                    | The number of domain names to be submitted on the order. This number must match the actual number of names submitted. |
| SLDX X=1 to DomainCount | string | Required | Second-level domain name \(for example, enom in enom.com\)                                                                                                                                                                                                    |
| TLDX             | string | Required                                                                    | Top-level domain name \(extension\) |
| AuthInfoX | string | Required | EPP TLDs Current \(“losing”\) registrar’s authorization key                                                                                                                                                                                                     |
| DomainPassword        | string | Optional                                                                    | Set a domain access password on the domain name |
| Lock | int   | Optional | Set Lock=1 to turn on RegistrarLock option                                                                                                                                                                                                             |
| Renew             | int | Optional                                                                    | Set Renew=1 to turn on Auto-Renew option |
| AuthString | string | Optional | Authorization string for automated transfer order entries \(approved accounts only\)                                                                                                                                                                                        |
| ExtendedAttributes      | string | Required for TLDs that use extended attributes                                                 | Extended attributes, required for some country code TLDs. You can retrieve a list of required extended attributes for any TLD using the GetExtAttributes command. |
| RegistrantAddress1 | string | Optional; Required for Fax orders or TLDs that require extended attributes. | Registrant address                                                                                                                                                                                                                         |
| RegistrantAddress2      | string | Optional; Required for Fax orders or TLDs that require extended attributes.                                  | Registrant additional address info |
| RegistrantCity | string | Optional; Required for Fax orders or TLDs that require extended attributes. | Registrant city                                                                                                                                                                                                                           |
| RegistrantCountry       | string | Optional; Required for Fax orders or TLDs that require extended attributes.                                  | Registrant country |
| RegistrantEmailAddress | string | Optional; Required for Fax orders or TLDs that require extended attributes. | Registrant email address                                                                                                                                                                                                                      |
| RegistrantFax         | string | Optional                                                                    | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URLencoded as a plus sign \(%2B\). |
| RegistrantFirstName | string | Optional; Required for Fax orders or TLDs that require extended attributes. | Registrant first name                                                                                                                                                                                                                        |
| RegistrantLastName      | string | Optional; Required for Fax orders or TLDs that require extended attributes.                                  | Registrant last name |
| RegistrantJobTitle | string | Optional | Registrant job title                                                                                                                                                                                                                        |
| RegistrantOrganizationName  | string | Optional; Required for Fax orders or TLDs that require extended attributes.                                  | Registrant organization |
| RegistrantPhone | string | Optional; Required for Fax orders or TLDs that require extended attributes. | Registrant phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URLencoded as a plus sign \(%2B\).                                                                                                                                          |
| RegistrantPostalCode     | string | Optional; Required for Fax orders or TLDs that require extended attributes.                                  | Registrant postal code |
| RegistrantStateProvince | string | Optional; Required for Fax orders or TLDs that require extended attributes. | Registrant state or province                                                                                                                                                                                                                    |
| RegistrantStateProvinceChoice | string | Optional                                                                    | Registrant state or province choice: S state P province |
| UseCreditCard | boolean | Optional for resellers who use our credit card processing AND want to charge this transaction to the credit card included in this query string | Permitted values are yes and no. The credit card supplied in this query string is charged only if UseCreditCard=yes. If this param is omitted or if UseCreditCard=no, the account balance rather than the credit card is debited for this transaction. This is true even if the query string includes all the Registrant contact and credit card information. When you pass credit card information with this command, you must use the secure HTTPS protocol. |
| EndUserIP           | string | Required if using our CC processing                                                      | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| CardType | string | Required if using our CC processing | Credit card type. Permitted values are Visa, Mastercard, AmEx, Discover                                                                                                                                                                                              |
| CreditCardNumber       | int | Required if using our CC processing                                                      | Credit card number |
| CreditCardExpMonth | int   | Required if using our CC processing | Expiration month of the credit card, in format MM                                                                                                                                                                                                          |
| CreditCardExpYear       | int | Required if using our CC processing                                                      | Expiration year of the credit card, in format YYYY |
| CVV2 | int   | Required if using our CC processing | Credit card verification code                                                                                                                                                                                                                    |
| CCName            | string | Required if using our CC processing                                                      | Cardholder’s name |
| CCAddress | string | Required if using our CC processing | Credit card billing street address                                                                                                                                                                                                                 |
| CCCity            | string | Required if using our CC processing                                                      | Credit card billing city |
| CCStateProvince | string | Required if using our CC processing | Credit card billing state or province                                                                                                                                                                                                                |
| CCCountry           | string | Required if using our CC processing                                                      | Credit card billing country. Two-character country code is a permitted format |
| CCZip | string | Required if using our CC processing | Credit card billing postal code                                                                                                                                                                                                                   |
| CCPhone            | string | Required if using our CC processing                                                      | Credit card billing phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL- encoded as a plus sign \(%2B\). |
| ChargeAmount | int   | Required if using our CC processing | Amount to charge this credit card. Required format is DD.cc                                                                                                                                                                                                    |
| ResponseType         | string | Optional                                                                    | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

Example of SLD and TLD params:

SLD1=first SLD

TLD1=first TLD

SLD2=second SLD

TLD2=secondTLD

Additional params for contact information can be passed by replacing Registrant in the param names above with Tech, Admin or AuxBilling.

If UseContacts is set =1 and contact information is passed in the URL, current Whois contacts are transferred and preconfigured contacts \(passed in the URL\) are ignored.

Automatic transfer of Whois information is available only for the largest registrars.

An OrderType of Fax requires a signed fax to process the order, an OrderType of Autoverification uses an electronic verification process to authorize and initiate the transfer.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| TransferOrderID | int   | Transfer order number. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_CreateOrder&uid=resellid&pw=resellpw
&orderType=AutoVerification&sld1=resellerdocs2
&tld1=net&AuthInfo1=ros8enQi&sld2=resellerdocs3&tld2=info
&AuthInfo2=pkv7ihRb& domaincount=2&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TP_CreateOrder&uid=resellid&pw=resellpw
&orderType=AutoVerification&sld1=resellerdocs2
&tld1=net&AuthInfo1=ros8enQi&sld2=resellerdocs3&tld2=info
&AuthInfo2=pkv7ihRb& domaincount=2&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_CreateOrder&uid=resellid&pw=resellpw
&orderType=AutoVerification&sld1=resellerdocs2
&tld1=net&AuthInfo1=ros8enQi&sld2=resellerdocs3&tld2=info
&AuthInfo2=pkv7ihRb& domaincount=2&responsetype=text
```
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
 <success>True</success>
 <Command>TP_CREATEORDER</Command>
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
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>success: </STRONG>True<br />
<STRONG>transferorderid: </STRONG>175696130<br />
<STRONG>orderdate: </STRONG>2/6/2015 1:22:47 PM<br />
<STRONG>ordertypeid: </STRONG>1<br />
<STRONG>ordertypedesc: </STRONG>Auto Verification<br />
<STRONG>statusid: </STRONG>4<br />
<STRONG>statusdesc: </STRONG>Processing<br />
<STRONG>authamount: </STRONG>17.95<br />
<STRONG>version: </STRONG>1<br />
<STRONG>transferorderdetailid1: </STRONG>77519246<br />
<STRONG>sld1: </STRONG>resellerdocs2<br />
<STRONG>tld1: </STRONG>net<br />
<STRONG>statusid1: </STRONG>9<br />
<STRONG>statusdesc1: </STRONG>Awaiting auto verification of transfer request<br />
<STRONG>price1: </STRONG>9<br />
<STRONG>usecontacts1: </STRONG>0<br />
<STRONG>premiumdomain1:</STRONG><br />
<STRONG>customersuppliedprice1:</STRONG><br />
<STRONG>transferorderdetailid2: </STRONG>77519247<br />
<STRONG>sld2: </STRONG>resellerdocs3<br />
<STRONG>tld2: </STRONG>info<br />
<STRONG>statusid2: </STRONG>13<br />
<STRONG>statusdesc2: </STRONG>Domain awaiting transfer initiation<br />
<STRONG>price2: </STRONG>8.95<br />
<STRONG>usecontacts2: </STRONG>0<br />
<STRONG>premiumdomain2:</STRONG><br />
<STRONG>customersuppliedprice2:</STRONG><br />
<STRONG>transferorderdetailcount: </STRONG>2<br />
<STRONG>success: </STRONG>True<br />
<STRONG>Command: </STRONG>TP_CREATEORDER<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+07.00<br />
<STRONG>ExecTime: </STRONG>2.902<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/6/2015 1:22:48 PM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
success=True
transferorderid=175696133
orderdate=2/6/2015 1:23:30 PM
ordertypeid=1
ordertypedesc=Auto Verification
statusid=4
statusdesc=Processing
authamount=17.95
version=1
transferorderdetailid1=77519248
sld1=resellerdocs2
tld1=net
statusid1=9
statusdesc1=Awaiting auto verification of transfer request
price1=9
usecontacts1=0
premiumdomain1=
customersuppliedprice1=
transferorderdetailid2=77519249
sld2=resellerdocs3
tld2=info
statusid2=13
statusdesc2=Domain awaiting transfer initiation
price2=8.95
usecontacts2=0
premiumdomain2=
customersuppliedprice2=
transferorderdetailcount=2
success=True
Command=TP_CREATEORDER
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+07.00
ExecTime=0.500
Done=true
RequestDateTime=2/6/2015 1:23:29 PM
```
Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

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