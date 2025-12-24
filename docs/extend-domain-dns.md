ExtendDomainDNS
===============

In real time, renew DNS hosting services for a domain name that is registered elsewhere.

Usage
-----

Use this command to renew DNS services for a name that is registered elsewhere. This is a real-time command that is most commonly used by resellers who maintain their own databases. Resellers who use a shopping cart and our order queue generally use the AddToCart command with input parameter ProductType=HostRenew.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com](https://resellertest.enom.com/)

The ExtendDomainDNS command is not implemented on [www.enom.com.](http://www.enom.com./) However, on the checkout page, the purchase button calls InsertNewOrder, a similar command that initiates checkout. The difference is that InserNewOrder acquires contents of the shopping cart and puts them in a queue for checkout; the ExtendDomainDNS command bypasses the shopping cart and the queue.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

To use our credit card processing, this must be an ETP reseller account.

The domain name must not be registered with us, but must use our name servers.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter        | Status | Description                                                                                       | Max Size |
| ----------------------------- | -------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID              | Required | Account login ID                                                                                     | 20 |
| PW | Required                   | Account password | 20    |
| SLD              | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with                                       | 63 |
| TLD | Required                   | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| NumYears           | Optional | Number of years to add to the DNS hosting subscription. Default value is 1. Maximum total is 10.                                           | 2 |
| RegistrantFirstName | Required if using our credit card processing | Registrant first name | 60    |
| RegistrantLastName      | Required if using our credit card processing | Registrant last name                                                                                   | 60 |
| RegistrantAddress1 | Required if using our credit card processing | Registrant address | 60    |
| RegistrantAddress2      | Optional | Registrant additional address info                                                                            | 60 |
| RegistrantCity | Required if using our credit card processing | Registrant city | 60    |
| RegistrantCountry       | Optional | Registrant country                                                                                    | 60 |
| RegistrantEmailAddress | Required if using our credit card processing | Registrant email address | 128   |
| RegistrantOrganizationName  | Optional | Registrant organization                                                                                 | 60 |
| RegistrantJobTitle | Optional                   | Registrant job title | 60    |
| RegistrantPhone        | Optional | Registrant phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).   | 20 |
| RegistrantFax | Optional                   | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). | 20    |
| RegistrantPostalCode     | Required if using our credit card processing | Registrant postal code                                                                                  | 16 |
| RegistrantStateProvince | Optional                   | Registrant state or province | 60    |
| RegistrantStateProvinceChoice | Optional | Registrant state or province choice: S state P province                                                                 | 1 |
| ChargeAmount | Optional                   | Amount to charge to the credit card | 6    |
| UseCreditCard         | Optional | If UseCreditCard=yes, use our creditcard processing services. When you pass credit card information with this command, you must use the secure HTTPS protocol.             | 3 |
| EndUserIP | Required If UseCreditCard=yes        | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. | 15    |
| CardType           | Required If UseCreditCard=yes | Type of credit card. Permitted values are Visa, Mastercard, AmEx, and Discover                                                     | 20 |
| CCName | Required If UseCreditCard=yes        | Cardholder's name | 60    |
| CreditCardNumber       | Required If UseCreditCard=yes | Customer's credit card number                                                                              | 128 |
| CreditCardExpMonth | Required If UseCreditCard=yes        | Credit card expiration month, in MM format | 2    |
| CreditCardExpYear       | Required If UseCreditCard=yes | Credit card expiration year, in YYYY format                                                                       | 4 |
| CVV2 | Required UseCreditCard=yes          | Credit card verification code | 4    |
| CCAddress           | Required If UseCreditCard=yes | Credit card billing address                                                                               | 60 |
| CCZip | Required If UseCreditCard=yes        | Credit card billing postal code | 15    |
| CCCountry           | Required If UseCreditCard=yes | Credit card billing country                                                                               | 60 |
| ResponseType | Optional                   | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| CCTransResult  | Success status for the credit card transaction |
| OrderID | Order ID, assigned by us                                     |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query renews DNS hosting for domain resellerdnshost.com. It uses a credit card, submits the credit card information using the secure HTTPS protocol, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=EXTENDDOMAINDNS&uid=resellid&pw=resellpw
&sld=resellerdnshost&tld=com&numyears=1
&RegistrantFirstName=John&RegistrantLastName=Doe
&RegistrantAddress1=17462&RegistrantCity=Redmond
&RegistrantPostalCode=98052&RegistrantCountry=USA
&RegistrantPhone=+1.4255559999
&[email protected]
&UseCreditCard=yes&CardType=mastercard&CCName=John+Doe
&CredcVV2=555&ccAddress=17462&ChargeAmount=12
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=EXTENDDOMAINDNS&uid=resellid&pw=resellpw
&sld=resellerdnshost&tld=com&numyears=1
&RegistrantFirstName=John&RegistrantLastName=Doe
&RegistrantAddress1=17462&RegistrantCity=Redmond
&RegistrantPostalCode=98052&RegistrantCountry=USA
&RegistrantPhone=+1.4255559999
&[email protected]
&UseCreditCard=yes&CardType=mastercard&CCName=John+Doe
&CredcVV2=555&ccAddress=17462&ChargeAmount=12
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=EXTENDDOMAINDNS&uid=resellid&pw=resellpw
&sld=resellerdnshost&tld=com&numyears=1
&RegistrantFirstName=John&RegistrantLastName=Doe
&RegistrantAddress1=17462&RegistrantCity=Redmond
&RegistrantPostalCode=98052&RegistrantCountry=USA
&RegistrantPhone=+1.4255559999
&[email protected]
&UseCreditCard=yes&CardType=mastercard&CCName=John+Doe
&CredcVV2=555&ccAddress=17462&ChargeAmount=12
&responsetype=text
```
The response is as follows:

```
<interface-response>
 <OrderID>157781689</OrderID>
 <Command>EXTENDDOMAINDNS</Command>
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
 <ExecTime>1.109</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 4:05:56 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
OrderID: 157781689
Command: EXTENDDOMAINDNS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site:
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.141
Done: true
RequestDateTime: 2/11/2015 1:29:20 PM
```
Related Commands
----------------

AddToCart

DeleteHostedDomain

Extend

InsertNewOrder

Purchase

PurchaseHosting

PurchasePOPBundle

PurchaseServices