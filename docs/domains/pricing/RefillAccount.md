RefillAccount
=============

Add funds to your account using a credit card.

Usage
-----

To add funds to your account, set parameter Debit=True and include the credit card parameters.

To remove your credit card information from our database, set parameter Debit=True and omit the credit card parameters.

To retrieve your account balance without adding funds, and to retrieve the current credit card information for the account, set parameter Debit=False.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The account must be a reseller account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=RefillAccount&uid=(Required)&pw=(Required)&Debit=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status                | Description |
| --------------- | ------- | ------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required               | RefillAccount |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                                                    |
| pw       | string | Required               | Your API Token |
| Debit | boolean | Required | Debit the credit card; options are True or False. Debit=True uses the credit card in this query string for this transaction, and replaces the credit card information for the account record with the values in this query string. Debit=False retrieves your account balance without changing it, and without charging your credit card; it also retrieves the current credit card information for the account. When you pass credit card information with this command, you must use the secure HTTPS protocol. |
| EndUserIP    | string | Required to add funds to this account | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| CCAmount | int   | Required to add funds to this account | Amount to charge to credit card, in DD.cc format. We charge a 3% convenience fee which is deducted immediately from the CCAmount. Min $100                                                                                                                                                                                     |
| CCType     | string | Required to add funds to this account | Credit card type. Options are AmEx, Discover, MasterCard, Visa. |
| CCName | string | Required to add funds to this account | Credit card holder’s name as imprinted on the credit card                                                                                                                                                                                                                               |
| CCNumber    | int | Required to add funds to this account | 16-digit credit card number |
| CCMonth | int   | Required to add funds to this account | Month in which credit card expires, in MM format                                                                                                                                                                                                                                   |
| CCYear     | int | Required to add funds to this account | Year in which credit card expires, in YYYY format |
| CVV2 | int   | Required to add funds to this account | Credit card security verification code, 3- or 4- digit number from the back of the credit card                                                                                                                                                                                                           |
| CCAddress    | string | Required to add funds to this account | Optional Street address as shown on credit card bill |
| CCCity | string | Required to add funds to this account | Optional City as shown on credit card bill                                                                                                                                                                                                                                      |
| CCStateProvince | string | Required to add funds to this account | State or province as shown on credit card bill |
| CCZip | string | Required to add funds to this account | Zip code or postal code as shown on credit card bill                                                                                                                                                                                                                                 |
| CCCountry    | string | Required to add funds to this account | Country of credit card billing address, expressed as 2-character country code |
| CCPhone | string | Required to add funds to this account | Phone number as shown on credit card bill. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).                                                                                                                                                     |
| SendMail    | string | Optional; default is 1        | Send email confirmation to the billing contact for this account. 0 or No turns off the email; any other value or omitting this parameter sends the email. |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML                                                                                                                                                                                                                       |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| Reseller     | int | 1 indicates that this is a reseller account                           |
| CCTRANSRESULT  | string | Result of the credit card transaction                              |
| ResellerRefill  | string | Success status of the account refill                               |
| CreditCardStatus | string | Success status of the credit card transaction                          |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=REFILLACCOUNT&uid=resellid&pw=resellpw
&CCAmount=100&CCType=MasterCard&CCName=JohnDoe
&CCNumber=5200520052005200&CCMonth=02
&CCYear=2013&cvv2=200&ccaddress=100+Main+St.
&CCStateProvince=WA&cczip=99999&debit=true
&CCCountry=us&CCPhone=+1.5555559999&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
command=REFILLACCOUNT&uid=resellid&pw=resellpw
&CCAmount=100&CCType=MasterCard&CCName=JohnDoe
&CCNumber=5200520052005200&CCMonth=02
&CCYear=2013&cvv2=200&ccaddress=100+Main+St.
&CCStateProvince=WA&cczip=99999&debit=true
&CCCountry=us&CCPhone=+1.5555559999&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=REFILLACCOUNT&uid=resellid&pw=resellpw
&CCAmount=100&CCType=MasterCard&CCName=JohnDoe
&CCNumber=5200520052005200&CCMonth=02
&CCYear=2013&cvv2=200&ccaddress=100+Main+St.
&CCStateProvince=WA&cczip=99999&debit=true
&CCCountry=us&CCPhone=+1.5555559999&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <CurrentDate>Wednesday, August 20, 2007</CurrentDate>
 <RefillAccount>
 <FName>John</FName>
 <LName>Doe</LName>
 <Reseller>1</Reseller>
 <CCTRANSRESULT>APPROVED</CCTRANSRESULT>
 <ResellerRefill>Transactions processed successfully.</ResellerRefill>
 <CreditCardStatus>Successful</CreditCardStatus>
 <CCYear>2004</CCYear>
 <CCMonth>2</CCMonth>
 <CCNumber>************5200</CCNumber>
 <CCName>JohnDoe</CCName>
 <PaymentType>MasterCard</PaymentType>
 <CCAddress>100 Main St.</CCAddress>
 <CCCity>Hometown</CCCity>
 <CCStateProvince>WA</CCStateProvince>
 <CCCountry>us</CCCountry>
 <CCPhoneDial />
 <CCPhone>5555559999</CCPhone>
 <CCZip>99999</CCZip>
 </RefillAccount>
 <Command>REFILLACCOUNT</Command>
 <ErrCount>0</ErrCount>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>0</IsLockable>
 <IsRealTimeTLD>0</IsRealTimeTLD>
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
<STRONG>CurrentDate: </ STRONG>Thursday, February 05, 2015<br>
<STRONG>ParentAccount: </ STRONG>000-00-0000<br>
<STRONG>BulkMemberShip: </ STRONG>0<br>
<STRONG>BulkMembershipType:</ STRONG><br>
<STRONG>FName: </ STRONG>John<br>
<STRONG>LName: </ STRONG>Doe<br>
<STRONG>Reseller: </ STRONG>1<br>
<STRONG>CustomerSite: </ STRONG>e<br>
<STRONG>IsInvoiceCustomer: </ STRONG>no<br>
<STRONG>CCTRANSRESULT: </ STRONG>APPROVED<br>
<STRONG>ResellerRefill: </ STRONG>Transactions processed successfully.<br>
<STRONG>CreditCardStatus: </ STRONG>Successful<br>
<STRONG>CCYear: </ STRONG>2015<br>
<STRONG>CCMonth: </ STRONG>2<br>
<STRONG>CCNumber: </ STRONG>************5200<br>
<STRONG>CCName: </ STRONG>JohnDoe<br>
<STRONG>PaymentType: </ STRONG>MasterCard<br>
<STRONG>CCAddress: </ STRONG>100 Main St.<br>
<STRONG>CCCity: </ STRONG>Hometown<br>
<STRONG>CCStateProvince: </ STRONG>WA<br>
<STRONG>CCCountry: </ STRONG>us<br>
<STRONG>CCPhoneDial:</ STRONG><br>
<STRONG>CCPhone: </ STRONG>5555559999<br>
<STRONG>CCZip: </ STRONG>99999<br>
<STRONG>Command: </ STRONG>REFILLACCOUNT<br>
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
<STRONG>ExecTime: </ STRONG>5.426<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/5/2015 1:59:54 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
CurrentDate=Thursday, February 05, 2015
ParentAccount=000-00-0000
BulkMemberShip=0
BulkMembershipType=
FName=John
LName=Doe
Reseller=1
CustomerSite=e
IsInvoiceCustomer=no
CCTRANSRESULT=APPROVED
ResellerRefill=Transactions processed successfully.
CreditCardStatus=Successful
CCYear=2015
CCMonth=2
CCNumber=************5200
CCName=JohnDoe
PaymentType=MasterCard
CCAddress=100 Main St.
CCCity=Hometown
CCStateProvince=WA
CCCountry=us
CCPhoneDial=
CCPhone=5555559999
CCZip=99999
Command=REFILLACCOUNT
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
ExecTime=5.426
Done=true
RequestDateTime=2/5/2015 1:59:54 PM
```
Related Commands
----------------

GetBalance

GetTransHistory

UpdateNotificationAmount