RenewPOPBundle
==============

Renew the annual subscription on a POP mail 10-pak in real time.

Usage
-----

Use this command to renew, in real time, the annual subscription on a POP mail 10-pak.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/POPConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/POPConfig.asp?DomainNameID=152533676)

The add years button behaves similarly to the RenewPOPBundle command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- This command is available for reseller accounts only. Retail accounts must use the queue-based
- shopping cart process to renew POP bundles.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter    | Status | Description                                                                                                                                                                                                                             | Max Size |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID          | Required | Account login ID                                                                                                                                                                                                                          | 20 |
| PW | Required                                                                               | Account password | 20    |
| SLD          | Required | Second-level domain name \(for example, enom in enom.com\)                                                                                                                                                                                                    | 63 |
| TLD | Required                                                                               | Top-level domain name \(extension\) | 15    |
| PakID         | Required | ID number of this POP pak. You can retrieve the ID numbers of all POP paks in the account with the GetPOPExpirations command.                                                                                                                                                                   | 10 |
| Quantity | Required                                                                               | Number of years to renew this POP pak subscription | 2    |
| UseCreditCard     | Optional overall, but Required for resellers who use our credit card processing AND want to charge this transaction to the credit card included in this query string | Permitted values are yes and no. The credit card supplied in this query string is charged only if UseCreditCard=yes. If this param is omitted or if UseCreditCard=no, the account balance rather than the credit card is debited for this transaction. This is true even if the query string includes all the Registrant contact and credit card information. When you pass credit card information with this command, you must use the secure HTTPS protocol. | 3 |
| EndUserIP | Required for our credit card processing                                                               | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. | 15    |
| RegistrantFirstName  | Required for our CC processing | Registrant first name                                                                                                                                                                                                                        | 60 |
| RegistrantLastName | Required for our CC processing                                                                    | Registrant last name | 60    |
| RegistrantAddress1   | Required for our CC processing | First line of Registrant address                                                                                                                                                                                                                  | 60 |
| RegistrantAddress2 | Optional for our CC processing                                                                    | Second line of Registrant address | 60    |
| RegistrantCity     | Required for our CC processing | Registrant city                                                                                                                                                                                                                           | 60 |
| RegistrantCountry | Required for our CC processing                                                                    | Registrant country. Two-letter country code is a permitted format. | 60    |
| RegistrantPostalCode  | Required for our CC processing | Registrant postal code                                                                                                                                                                                                                       | 16 |
| RegistrantPhone | Required for our CC processing                                                                    | Registrant phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). | 20    |
| RegistrantEmailAddress | Required for our CC processing | Registrant email address                                                                                                                                                                                                                      | 128 |
| CardType | Required for our CC processing                                                                    | Credit card type. Permitted values are Visa, Mastercard, AmEx, Discover | 20    |
| CCName         | Required for our CC processing | Cardholder’s name                                                                                                                                                                                                                          | 60 |
| CreditCardNumber | Requiredfor our CC processing                                                                    | Credit card number | 128   |
| CreditCardExpMonth   | Required for our CC processing | Expiration month of the credit card, in format MM                                                                                                                                                                                                          | 2 |
| CreditCardExpYear | Required for our CC processing                                                                    | Expiration year of the credit card, in format YYYY | 4    |
| CVV2          | Required for our CC processing | Credit card verification code                                                                                                                                                                                                                    | 4 |
| CCAddress | Required for our CC processing                                                                    | Credit card billing address | 60    |
| ChargeAmount      | Required for our CC processing | Amount to charge this credit card. Required format is DD.cc                                                                                                                                                                                                    | 6 |
| ResponseType | Optional Format of response.                                                                    | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| OrderID     | Identification number for tracking this order in our system |
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

The following query renews POP pak 5105 for 1 year, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=renewpopbundle&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&pakid=5105
&quantity=1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=renewpopbundle&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&pakid=5105
&quantity=1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=renewpopbundle&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&pakid=5105
&quantity=1&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <orderid>156227365</orderid>
 <Command>RENEWPOPBUNDLE</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <ExecTime>0.34375</ExecTime>
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
Command: RENEWPOPBUNDLE
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site:
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.313
Done: true
RequestDateTime: 2/5/2015 2:19:02 PM
```
Related Commands
----------------

Extend

GetPOP3

GetPOPExpirations

GetRenew

Purchase

PurchasePOPBundle

RenewServices