RenewServices
=============

Renew one value-added service in real time.

Usage
-----

Use this command to renew one value-added service in real time, as opposed to going through our shopping cart.

This real-time command currently supports ID Protect and DNS hosting.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

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
- The service must currently be subscribed for this domain. If not already subscribed, do so using the PurchaseServices or AddToCart commands.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter  | Status | Description                                                                                                                                                                                                                                                    | Max Size |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID        | Required | Account login ID                                                                                                                                                                                                                                                 | 20 |
| PW | Required                                                                                                           | Account password | 20    |
| Service      | Required | Service to be renewed. Permitted values: WPPS ID Protect DNSHosting DNS Hosting                                                                                                                                                                                                                 | 15 |
| NumYears | Optional; default is 1                                                                                                    | Number of years to renew this service | 2    |
| SLD        | Required for WPPS | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with                                                                                                                                                                                                    | 63 |
| TLD | Required for WPPS                                                                                                      | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| UseCreditCard   | Optional overall, but Required resellers who have a credit card processing agreement with us AND want to use our CC processing to charge this transaction to the credit card included in this query string. Default is no. | Available for WPPS DNSHosting Permitted values are yes and no. The credit card supplied in this query string is charged only if UseCreditCard=yes. If this param is omitted, or if the value supplied is anything other than yes, the account balance rather than the credit card is debited for this transaction. This is true even if the query string includes all the required credit card information. When you pass credit card information with this command, you must use the secure HTTPS protocol. | 3 |
| EndUserIP | Required for our credit card processing                                                                                           | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. | 15    |
| ChargeAmount    | Required our CC processing | Amount to charge for each year; ChargeAmount is multiplied by NumYears. Required format is DD.cc                                                                                                                                                                                                        | 6 |
| CardType | Required for our CC processing                                                                                                | Credit card type. Permitted values are Visa, Mastercard, AmEx, Discover | 20    |
| CreditCardNumber  | Required our CC processing | Credit card number                                                                                                                                                                                                                                                | 128 |
| CreditCardExpMonth | Required for our CC processing                                                                                                | Expiration month of the credit card, in format MM | 2    |
| CreditCardExpYear | Required our CC processing | Expiration year of the credit card, in format YYYY                                                                                                                                                                                                                                | 4 |
| CVV2 | Required for our CC processing, if printed on the card                                                                                    | Credit card verification value | 4    |
| CCName       | Required for our CC processing | Cardholder’s name                                                                                                                                                                                                                                                 | 60 |
| CCAddress | Required for our CC processing                                                                                                | Credit card billing street address | 60    |
| CCCity       | Optional for our CC processing | Credit card billing city                                                                                                                                                                                                                                             | 60 |
| CCStateProvince | Optional for our CC processing                                                                                                | Credit card billing state or province | 40    |
| CCCountry     | Required for our CC processing | Credit card billing country. Twocharacter country code is a permitted format                                                                                                                                                                                                                   | 60 |
| CCZip | Required for our CC processing                                                                                                | Credit card billing postal code | 60    |
| CCPhone      | Optional for our CC processing | Credit card billing phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).                                                                                                                                                           | 20 |
| ResponseType | Optional                                                                                                           | Format of response. Permitted values are Text \(default\), HTML, or XML | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| OrderID     | Identification number for tracking this order in our system |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query renews ID Protect for two years and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=renewservices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=WPPS
&numyears=2&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=renewservices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=WPPS
&numyears=2&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=renewservices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=WPPS
&numyears=2&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <OrderID>157405437</OrderID>
 <Command>RENEWSERVICES</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLER1-STG</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>0.797</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA[ ]]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
OrderID: 161808576
Command: RENEWSERVICES
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.828
Done: true
RequestDateTime: 2/5/2015 2:21:55 PM
```