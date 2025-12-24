TP\SubmitOrder
==============

Submit a preconfigured transfer order.

Usage
-----

Use this command to submit a transfer order that has been created and configured using the TP\_CreateOrder command, but not submitted. Domain names in this status appear on the Pending orders page. When you pass credit card information with this command, you must use the secure HTTPS protocol.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- To use our credit card processing, this must be an ETP reseller account.
- The TransferOrderID must identify an order that has been successfully created but not submitted.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=TP_SubmitOrder&uid=(Required)&pw=(Required)&TransferOrderID=(Required)&responsetype=(Optional)
```
| Input Parameter  | Type | Status                                                                     | Description |
| ------------------ | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command      | string | Required                                                                    | TP\_SubmitOrder |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                            |
| pw         | string | Required                                                                    | Your API Token |
| TransferOrderID | int  | Required | Transfer order ID number. To retrieve this value, use the TP\_GetOrderStatuses command                                                                                                                                                                                        |
| EndUserIP     | string | Required our CC processing                                                           | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| UseCreditCard | string | Optional for resellers who use our credit card processing AND want to charge this transaction to the credit card included in this query string | Permitted values are yes. or no. The credit card supplied in this query string is charged only if UseCreditCard=yes. If this param is omitted or if UseCreditCard=no , the account balance rather than the credit card is debited for this transaction. This is true even if the query string includes all the Registrant contact and credit card information. When you pass credit card information with this command, you must use the secure HTTPS protocol. |
| CardType      | string | Required our CC processing                                                           | processing Credit card type. Permitted values are Visa, Mastercard, AmEx, Discover |
| CCName | string | Required our CC processing | Cardholder’s name                                                                                                                                                                                                                           |
| CreditCardNumber  | string | Required our CC processing                                                           | Credit card number |
| CreditCardExpMonth | string | Required our CC processing | Expiration month of the credit card, in format MM                                                                                                                                                                                                           |
| CreditCardExpYear | string | Required our CC processing                                                           | Expiration year of the credit card, in format YYYY |
| CVV2 | string | Required our CC processing | Credit card verification code                                                                                                                                                                                                                     |
| CCName       | string | Required our CC processing                                                           | Cardholder’s name |
| CCAddress | string | Required our CC processing | Credit card billing street address                                                                                                                                                                                                                  |
| CCCity       | string | Required our CC processing                                                           | Credit card billing city |
| CCStateProvince | string | Required our CC processing | Credit card billing state or province                                                                                                                                                                                                                 |
| CCCountry     | string | Required our CC processing                                                           | Credit card billing country. Two-character country code is a permitted format |
| CCZip | string | Required our CC processing | Credit card billing postal code                                                                                                                                                                                                                    |
| CCPhone      | string | Optional for our CC processing                                                         | Credit card billing phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). |
| ChargeAmount | int  | Required our CC processing | Amount to charge this credit card. Required format is DD.cc                                                                                                                                                                                                     |
| ResponseType    | string | Optional                                                                    | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_SubmitOrder&uid=resellid&pw=resellpw
&TransferOrderID=445440&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TP_SubmitOrder&uid=resellid&pw=resellpw
&TransferOrderID=445440&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_SubmitOrder&uid=resellid&pw=resellpw
&TransferOrderID=445440&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <success>True</success>
 <Command>TP_SUBMITORDER</Command>
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
<STRONG>Command: </ STRONG>TP_SUBMITORDER<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site:</ STRONG><br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.109<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/9/2015 1:12:27 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command=TP_SUBMITORDER
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.063
Done=true
RequestDateTime=2/9/2015 1:13:05 PM
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

TP\_UpdateOrderDetail

UpdatePushList