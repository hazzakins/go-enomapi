Extend
======

Extend \(renew\) the registration period for a domain name.

Usage
-----

Use this command to renew domain registrations in real time. We recommend that you renew at least a week before the domain expiration date.

> ### When Using eNom's credit card processing services
>
>
>
> When you pass credit card information with this command, you must use the secure HTTPS protocol.

Most commonly, this command is used by resellers who maintain their own databases. Resellers who use a shopping cart and our order processing queue generally use the "AddToCart" command with input parameter "ProductType=Renew".

After a domain expires, use "UpdateExpiredDomains" to reactivate it and renew its registration.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- To use our credit card processing, this must be an ETP reseller account.
- The domain name must be valid and active.
- The new expiration date cannot be more than 10 years in the future.

> ### When used to Extend Premium domain names
>
>
>
> The "Extend" command can now be used to extend Premium domain names. The "CustomerSuppliedPrice" parameter must be passed to extend a premium domain name.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=Extend&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&NumYears={Required}&responsetype={Optional}
```
| Input Parameter        | Type | Status                                       | Description | Max Size |
| ----------------------------- | ------ | ---------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| command            | string | Required                                      | Extend | 6    |
| uid              | string | Required                                      | Your Account ID | 20    |
| pw              | string | Required                                      | Your API Token | 20    |
| SLD              | string | Required                                      | Second-level domain name \(e.g. "enom" in "enom.com"\) | 63    |
| TLD              | string | Required                                      | Top-level domain name \(extension\) | 15    |
| NumYears           | int | Optional                                      | Number of years to extend | 10    |
| OverrideOrder         | string | Optional                                      | Use "OverrideOrder=1" to submit more than one renewal request for this domain name within 24 hours of the first transaction. | 1    |
| UseCreditCard         | string | Optional                                      | If "UseCreditCard=yes", use eNom’s credit-card processing services. When you pass credit card information with this command, you must secure HTTPS protocol. | 3    |
| EndUserIP           | string | Required If "UseCreditCard=yes " is used.                     | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. | 15    |
| ChargeAmount         | float | Required If "UseCreditCard=yes " is used.                     | Amount to charge per year for the renewal. Use DD.cc format | 6    |
| CustomerSuppliedPrice     | float | Required If "UseCreditCard=yes " is used and the domain name is a Premium Domain. | The price of the domain's renewal, as shown in the ouput of the [Check](../docs/domains/availability/check.md) command. |     |
| CardType           | string | Required If "UseCreditCard=yes " is used.                     | Type of credit card. Permitted values are: - Visa - Mastercard - AmEx - Discover | 20    |
| CCName            | string | Required If "UseCreditCard=yes " is used.                     | Cardholder's name | 60    |
| CreditCardNumber       | int | Required If "UseCreditCard=yes " is used.                     | Customer's credit card number | 128   |
| CreditCardExpMonth      | int | Required If "UseCreditCard=yes " is used.                     | Credit card expiration month in MM format | 2    |
| CreditCardExpYear       | int | Required If "UseCreditCard=yes " is used.                     | Credit card expiration year in YYYY format | 4    |
| CVV2             | int | Required If "UseCreditCard=yes " is used.                     | Credit card verification code | 4    |
| CCAddress           | string | Required If "UseCreditCard=yes " is used.                     | Credit card billing address | 60    |
| CCZip             | int | Required If "UseCreditCard=yes " is used.                     | Credit card billing postal code | 15    |
| CCCountry           | string | Required If "UseCreditCard=yes " is used.                     | Credit card billing country. The two-letter country code is a permitted format. | 60    |
| RegistrantFirstName      | string | Required If "UseCreditCard=yes " is used.                     | Registrant first name | 60    |
| RegistrantLastName      | string | Required If "UseCreditCard=yes " is used.                     | Registrant last name | 60    |
| RegistrantAddress1      | string | Required If "UseCreditCard=yes " is used.                     | Registrant address | 60    |
| RegistrantAddress2      | string | Optional                                      | Registrant additional address info | 60    |
| RegistrantCity        | string | Required If "UseCreditCard=yes " is used.                     | Registrant city | 60    |
| RegistrantStateProvinceChoice | string | Required If "UseCreditCard=yes " is used.                     | Registrant state or province choice. Permitted values are: - S - state - P - province | 1    |
| RegistrantStateProvince    | string | Required If "UseCreditCard=yes " is used.                     | Registrant state or province | 60    |
| RegistrantPostalCode     | string | Required If "UseCreditCard=yes " is used.                     | Registrant postal code | 16    |
| RegistrantCountry       | string | Required If "UseCreditCard=yes " is used.                     | Registrant country | 60    |
| RegistrantEmailAddress    | string | Required If "UseCreditCard=yes " is used.                     | Registrant email address | 128   |
| RegistrantOrganizationName  | string | Optional                                      | Registrant organization | 60    |
| RegistrantJobTitle      | string | Optional                                      | Registrant job title | 60    |
| RegistrantPhone        | string | Required If "UseCreditCard=yes " is used.                     | Registrant phone number. Required format is "\+CountryCode.PhoneNumber", where "CountryCode" and "PhoneNumber" use only numeric characters and the "\+ symbol" is URL-encoded as a plus sign \(%2B\). | 20    |
| RegistrantFax         | string | Optional                                      | Registrant fax number. Required format is "\+CountryCode.PhoneNumber", where "CountryCode" and "PhoneNumber" use only numeric characters and the "\+ symbol" is URL-encoded as a plus sign \(%2B\). | 20    |
| ResponseType         | string | .Optional                                     | Format of response. Permitted values are: Text \(default\), HTML, or XML | 4    |

Returned Parameters and Values
------------------------------

> ### If using the "UseCreditCard=yes" option, then ALL of the credit card information is required.

The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

We recommend that you store the "OrderID" value, at least the most recent one for each domain, from the return. Several other commands use this value as a required input parameter.

| Output Parameter | Type | Description                                             |
| ---------------- | ------- | --------------------------------------------------------------------------------------------------- |
| command     | string | Extend                                               |
| OrderID     | int | An order number is returned if successful. We recommend that you store this value for future use. |
| Extension    | string | Returns "Successful", otherwise this parameter isn't returned.                   |
| RRPCode     | int | Success code. Only a "200" indicates success.                           |
| RRPText     | string | Text which accompanies and describes the "RRPCode".                        |
| UseCreditCard  | boolean | Use our credit card processing services                               |
| TotalCharged   | float | Total points or $US charged for this order                             |
| RegistryExpDate | string | Domain expiration date at the Registry level.                           |
| errcount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.  |
| errX       | string | Error messages explaining the failure. These can be presented as is back to the client.      |
| done       | string | True value indicates this entire response has reached you successfully.              |

Example Output
--------------

The following query requests that the registration period for "resellerdocs.com" be extended for 1 year, and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=modifyns&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ns1=ns1.name-services.com&ns2=ns2.name-ervices.com&responsetype={Optional}
```
```
<interface-response>
<Extension>successful</Extension>
<DomainName>resellerdocs.com</DomainName>
<OrderID>157781677</OrderID>
<RRPCode>200</RRPCode>
<RRPText>Command completed successfully</RRPText>
<DomainInfo>
<RegistryExpDate>2015-03-31 09:17:29.000</RegistryExpDate>
</DomainInfo>
<Command>EXTEND</Command>
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
<ExecTime>1.452</ExecTime>
<Done>true</Done>
<RequestDateTime>12/12/2011 3:32:35 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Extension: </STRONG>successful<br>
<STRONG>DomainName: </STRONG>resellerdocs.com<br>
<STRONG>OrderID: </STRONG>157781677<br>
<STRONG>RRPCode: </STRONG>200<br>
<STRONG>RRPText: </STRONG>Command completed successfully<br>
<STRONG>RegistryExpDate: </STRONG>2015-03-31 09:17:29.000<br>
<STRONG>Command: </STRONG>EXTEND<br>
<STRONG>APIType: </STRONG>API<br>
<STRONG>Language: </STRONG>eng<br>
<STRONG>ErrCount: </STRONG>0<br>
<STRONG>ResponseCount: </STRONG>0<br>
<STRONG>MinPeriod: </STRONG>1<br>
<STRONG>MaxPeriod: </STRONG>10<br>
<STRONG>Server: </STRONG>SJL0VWRESELL_T1<br>
<STRONG>Site: </STRONG><br>
<STRONG>IsLockable: </STRONG>True<br>
<STRONG>IsRealTimeTLD: </STRONG>True<br>
<STRONG>TimeDifference: </STRONG>+08.00<br>
<STRONG>ExecTime: </STRONG>0.359<br>
<STRONG>Done: </STRONG>true<br>
<STRONG>RequestDateTime: </STRONG>2/11/2015 1:19:45 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Extension=successful
DomainName=resellerdocs.com
OrderID=157781677
RRPCode=200
RRPText=Command completed successfully
RegistryExpDate=2015-03-31 09:17:29.000
Command=EXTEND
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.391
Done=true
RequestDateTime=2/11/2015 1:21:53 PM
```
Related Commands
----------------

Extend\_RGP

ExtendDomainDNS

GetDomainExp

GetExtendInfo

GetPOPExpirations

GetRenew

InsertNewOrder

RenewPOPBundle

SetRenew

UpdateExpiredDomains