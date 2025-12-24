CreateAccount
=============

Create a new sub-account.

Usage
-----

Use this command to create a new retail sub-account.

"CreateAccount" differs from "CreateSubAccount" in that "CreateAccount" offers credit card processing.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- Parameter values submitted in the query must pass validation tests.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=CreateAccount&uid=Your Account ID&pw=Your API Token&NewUID={Required}&NewPW={Required}&ConfirmPW={Required}&RegistrantOrganizationName={Required}&RegistrantJobTitle={Optional}&RegistrantFirstName={Required}&RegistrantLastName={Required}&RegistrantAddress1={Required}&RegistrantAddress2={Optional}&RegistrantCity={Required}&RegistrantPostalCode={Required}&RegistrantStateProvince={Optional}&RegistrantStateProvinceChoice={Optional}&RegistrantCountry={Optional}&RegistrantPhone={Required}&RegistrantFax={Optional}&RegistrantEmailAddress={Required}&RegistrantEmailAddress_Contact={Optional}&CardType={Required}&CCName={Required}&CreditCardNumber={Required}&CreditCardExpMonth={Required}&CreditCardExpYear={Required}&CVV2={Required}&CCAddress={Required}&CCZip={Required}&CCCountry={Required}&AuthQuestionType={Required}&AuthQuestionAnswer={Required}&EmailInfo={Optional}&ResponseType={Optional}
```
| Input Parameter         | Type | Status  | Description |
| ------------------------------- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command             | string | Required | CreateAccount |
| uid | string | Required | Your Account ID                                                                                                                                                                |
| pw               | string | Required | Your API Token |
| NewUID | string | Required | Sub-account login ID. Permitted values are 6 to 20 characters in length; permitted characters include letters, numbers, hyphen, and underscore.                                                                                              |
| NewPW              | string | Required | Sub-account password. Passwords must be from 10 to 20 characters in length with the following parameters: - Alpha-Numeric characters - - @ - \_ . / |
| ConfirmPW | string | Required | Sub-account password. Passwords must be from 10 to 20 characters in length with the following parameters: - Alpha-Numeric characters - - @ - \_ . /                                                                                           |
| RegistrantOrganizationName   | string | Required | Registrant organization |
| RegistrantJobTitle | string | Optional | Registrant job title                                                                                                                                                             |
| RegistrantFirstName       | string | Required | Registrant first name |
| RegistrantLastName | string | Required | Registrant last name                                                                                                                                                             |
| RegistrantAddress1       | string | Required | Registrant address |
| RegistrantAddress2 | string | Optional | Registrant additional address info                                                                                                                                                      |
| RegistrantCity         | string | Required | Registrant city |
| RegistrantPostalCode | string | Required | Registrant postal code                                                                                                                                                            |
| RegistrantStateProvince     | string | Optional | Registrant state or province |
| RegistrantStateProvinceChoice | string | Optional | Registrant state or province choice. Permitted values are: - S - state - P - province                                                                                                                           |
| RegistrantCountry        | string | Optional | Registrant country |
| RegistrantPhone | string | Required | Registrant phone number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).                                                                          |
| RegistrantFax          | string | Optional | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). |
| RegistrantEmailAddress | string | Required | Email address for Whois                                                                                                                                                            |
| RegistrantEmailAddress\_Contact | string | Optional | Email address for us to contact you about your domain name account |
| CardType | string | Required | Type of credit card. Permitted values are - Visa - Mastercard - AmEx - Discover When you pass credit card information with this command, you must use the secure HTTPS protocol.                                                                              |
| CCName             | string | Required | Cardholder's name |
| CreditCardNumber | int  | Required | Customer's credit card number                                                                                                                                                         |
| CreditCardExpMonth       | int | Required | Credit card expiration month in MM format |
| CreditCardExpYear | int  | Required | Credit card expiration year in YYYY format                                                                                                                                                  |
| CVV2              | int | Required | Credit card verification code |
| CCAddress | string | Required | Credit card billing address                                                                                                                                                          |
| CCZip              | string | Required | Credit card billing postal code |
| CCCountry | string | Required | Credit card billing country. The two-letter country code is a permitted format.                                                                                                                              |
| AuthQuestionType        | string | Required | Nature of the question used for identity verification. Permitted values are: - smaiden - mother’s maiden name - sbirth - city of birth - ssocial - last 4 digits of SSN - shigh - high school - fteach - favorite teacher - fvspot - favorite vacation spot - fpet - favorite pet - fmovie - favorite movie - fbook - favorite book |
| AuthQuestionAnswer | string | Required | The answer to the question type \(mother's maiden name, city of birth, last 4 digits of social security, or high school\)                                                                                                           |
| EmailInfo            | string | Optional | The value of a send e-mail notification checkbox. If "EmailInfo=checked" is set, an email will be sent to the registrant email parameter posted with this request, containing this account information |
| ResponseType | string | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML.                                                                                                                                 |
| VatEntityType          | String | Optional | Vat Entity Type. Examples: - business - charity |
| VatId | string | Optional | Enter a Valid VAT number                                                                                                                                                           |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.
- If there are errors validating the data passed in, the same parameters and values are returned to rebuild the form.

| Output Parameter   | Type | Description                                           |
| --------------------- | ------ | ------------------------------------------------------------------------------------------------ |
| OrganizationName   | string | Organization name of the Billing contact                             |
| FName         | string | First name \(given name\) of the Billing contact                         |
| LName         | string | Last name \(family name\) of the Billing contact                         |
| JobTitle       | string | Job title for the Billing contact                                |
| Address1       | string | Address, line 1, for the Billing contact                             |
| Address2       | string | Address, line 2, for the Billing contact                             |
| City         | string | City for the Billing contact                                   |
| StateProvince     | string | State or province for the Billing contact                            |
| StateProvinceChoice  | string | Tells whether StateProvince value is a state or a province                    |
| PostalCode      | string | Postal code for the Billing contact                               |
| Country        | string | County for the Billing contact                                  |
| Phone         | string | Phone number for the Billing contact                               |
| PhoneExt       | string | Phone extension for the Billing contact                             |
| Fax          | string | Fax number for the Billing contact                                |
| EmailAddress     | string | Email address for public WhoIs record                              |
| EmailAddress\_Contact | string | Email address that we use when contacting account owner about this account            |
| EmailInfo       | string | Tells whether confirmation emails should be sent to these email addresses            |
| CCType        | string | Type of credit card                                       |
| CCName        | string | Cardholder’s name                                        |
| CCNumber       | string | Credit card number                                        |
| CCMonth        | int | Expiration month for this credit card                              |
| CCYear        | int | Expiration year for this credit card                               |
| CCAddress       | string | Billing address street address for this credit card                       |
| CCCity        | string | Billing address city for this credit card                            |
| CCStateProvince    | string | Billing address state or province for this credit card                      |
| CCZip         | string | Billing address postal code for this credit card                         |
| CCCountry       | string | Billing address country for this credit card                           |
| CCPhone        | string | Phone number on record for this credit card                           |
| Reseller       | int | Indicates the sub-account status. Expected values are: - 1 - sub-reseller - 0 - sub-retail |
| NewUID        | string | Login ID for this subaccount                                   |
| NewPW         | string | Password for this subaccount                                   |
| ConfirmPW       | string | Repeat password, for confirmation purposes                            |
| AuthQuestionType   | string | Type of authorization question for this subaccount                        |
| AuthQuestionAnswer  | string | Answer to the authorization question                               |
| DefPeriod       | int | Default registration period for domain names, in years                      |
| StatusCustomerInfo  | string | Returns "Successful" if the subaccount was created                        |
| PartyID        | int | "PartyID" of this subaccount, for our records                          |
| Account        | int | Account ID code of this sub-account, for our records. We recommend that you store this value. |
| sLoginPass      | string | Code we use to log the user into the new sub-account immediately.               |
| VatEntityType     | string | Valid VAT entity                                         |
| VatId         | string | Valid VAT number                                         |
| command        | string | CreateAccount                                          |
| errcount       | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| errX         | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| done         | string | True value indicates this entire response has reached you successfully.             |

Example Output
--------------

The following query creates sub-account "ichiro" and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=CreateAccount&uid=Your Account ID&pw=Your API Token&NewUID=ichiro&NewPW=ichiropw&ConfirmPW=ichiropw&RegistrantOrganizationName=Reseller20Documents20Inc.&RegistrantFirstName=John&RegistrantLastName=Doe&RegistrantJobTitle=First%20Baseman&RegistrantAddress1=111%20Main%20St.&RegistrantCity=Hometown&RegistrantPostalCode=99999&RegistrantCountry=United+States&RegistrantEmailAddress=john%2Edoe%40resellerdocs%2Ecom&RegistrantPhone=+1.5555555555&RegistrantFax=+1.5555555556&AuthQuestionType=smaiden&AuthQuestionAnswer=jones&CardType=visa&CCName=test1&ResponseType={Optional}
```
```
<interface-response>
<NewAccount>
<OrganizationName>Reseller Documents Inc.</OrganizationName>
<Fname>John</Fname>
<Lname>Doe</Lname>
<JobTitle>First Baseman</JobTitle>
<Address1>111 Main St.</Address1>
<Address2/>
<City>Hometown</City>
<StateProvince/>
<StateProvinceChoice/>
<PostalCode>99999</PostalCode>
<Country>US</Country>
<Phone>+1.5555555555</Phone>
<PhoneExt/>
<Fax>+1.5555555556</Fax>
<EmailAddress>[email protected]</EmailAddress>
<EmailAddress_Contact/>
<EmailInfo/>
<GetCustomerPaymentInfo>
<CCType>visa</CCType>
<CCName>test1</CCName>
<CCNumber/>
<CCMonth/>
<CCYear/>
<CCAddress/>
<CCCity/>
<CCStateProvince/>
<CCZip/>
<CCCountry/>
<CCPhone/>
</GetCustomerPaymentInfo>
<GetAccountInfo>
<Reseller>0</Reseller>
<NewUID>ichiro</NewUID>
<NewPW>ichiropw</NewPW>
<ConfirmPW>ichiropw</ConfirmPW>
<AuthQuestionType>smaiden</AuthQuestionType>
<AuthQuestionAnswer>jones</AuthQuestionAnswer>
</GetAccountInfo>
<GetCustomerPreferences>
<DefPeriod/>
</GetCustomerPreferences>
<StatusCustomerInfo>Successful</StatusCustomerInfo>
<PartyID>{779CA342-D320-E111-A3F0-005056BC7747}</PartyID>
<Account>599-nj-0803</Account>
<sLoginPass>7A1218131209144A4948401218131209140B0C404A49544C54494B4A4A5B4F414E42414A4F5B3A36</sLoginPass>
</NewAccount>
<Command>CREATEACCOUNT</Command>
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
<ExecTime>1.313</ExecTime>
<Done>true</Done>
<RequestDateTime>12/7/2011 4:59:14 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>OrganizationName=</STRONG>Reseller Documents Inc.<br />
<STRONG>Fname= </STRONG>John<br />
<STRONG>Lname= </STRONG>Doe<br />
<STRONG>JobTitle=</STRONG>First Baseman<br />
<STRONG>Address1= </STRONG>111 Main St.<br />
<STRONG>Address2= </STRONG><br />
<STRONG>City=</STRONG>Hometown<br />
<STRONG>StateProvince= </STRONG><br />
<STRONG>StateProvinceChoice= </STRONG><br />
<STRONG>PostalCode= </STRONG>99999<br />
<STRONG>Country= </STRONG>US<br />
<STRONG>Phone= </STRONG>+1.5555555555<br />
<STRONG>PhoneExt= </STRONG><br />
<STRONG>Fax= </STRONG>+1.5555555556<br />
<STRONG>EmailAddress= </STRONG>[email protected]<br />
<STRONG>EmailAddress_Contact= </STRONG><br />
<STRONG>EmailInfo= </STRONG><br />
<STRONG>CCType= </STRONG>visa<br />
<STRONG>CCName= </STRONG>test1<br />
<STRONG>CCNumber= </STRONG><br />
<STRONG>CCMonth= </STRONG><br />
<STRONG>CCYear= </STRONG><br />
<STRONG>CCAddress= </STRONG><br />
<STRONG>CCCity= </STRONG><br />
<STRONG>CCStateProvince= </STRONG><br />
<STRONG>CCZip= </STRONG><br />
<STRONG>CCCountry= </STRONG><br />
<STRONG>CCPhone= </STRONG><br />
<STRONG>Reseller= </STRONG>0<br />
<STRONG>NewUID= </STRONG>ichiro<br />
<STRONG>NewPW= </STRONG>ichiropw<br />
<STRONG>ConfirmPW= </STRONG>ichiropw<br />
<STRONG>AuthQuestionType= </STRONG>smaiden<br />
<STRONG>AuthQuestionAnswer= </STRONG>jones<br />
<STRONG>DefPeriod= </STRONG><br />
<STRONG>StatusCustomerInfo= </STRONG>Successful<br />
<STRONG>PartyID= </STRONG>{779CA342-D320-E111-A3F0-005056BC7747}<br />
<STRONG>Account= </STRONG>599-nj-0803<br />
<STRONG>sLoginPass= </STRONG>7A1218131209144A4948401218131209140B0C404A49544C54494B4A4A5B4F414E42414A4F5B3A36<br />
<STRONG>Command= </STRONG>CREATEACCOUNT<br />
<STRONG>APIType= </STRONG>API<br />
<STRONG>Language= </STRONG>eng<br />
<STRONG>ErrCount= </STRONG>0<br />
<STRONG>ResponseCount= </STRONG>0<br />
<STRONG>MinPeriod= </STRONG>1<br />
<STRONG>MaxPeriod= </STRONG>10<br />
<STRONG>Server= </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site= </STRONG>eNom<br />
<STRONG>IsLockable= </STRONG>True<br /><br />
<STRONG>IsRealTimeTLD= </STRONG>True<br />
<STRONG>TimeDifference= </STRONG>+08.00<br />
<STRONG>ExecTime= </STRONG>1.313<br />
<STRONG>Done= </STRONG>true<br />
<STRONG>RequestDateTime= </STRONG>2/3/2015 3:00:18 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
OrganizationName=Reseller Documents Inc.
Fname=John
Lname=Doe
JobTitle=First Baseman
Address1=111 Main St.
Address2=
City=Hometown
StateProvince=
StateProvinceChoice=
PostalCode=99999
Country=US
Phone=+1.5555555555
PhoneExt=
Fax=+1.5555555556
[email protected]
EmailAddress_Contact=
EmailInfo=
CCType=visa
CCName=test1
CCNumber=
CCMonth=
CCYear=
CCAddress=
CCCity=
CCStateProvince=
CCZip=
CCCountry=
CCPhone=
Reseller=0
NewUID=ichiro
NewPW=ichiropw
ConfirmPW=ichiropw
AuthQuestionType=smaiden
AuthQuestionAnswer=jones
DefPeriod=
StatusCustomerInfo=Successful
PartyID={779CA342-D320-E111-A3F0-005056BC7747}
Account=599-nj-0803
sLoginPass=7A1218131209144A4948401218131209140B0C404A49544C54494B4A4A5B4F414E42414A4F5B3A36
Command=CREATEACCOUNT
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
ResponseNumber1=304157
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=1.313
Done=true
RequestDateTime=2/3/2015 3:00:18 PM
```
Related Commands
----------------

[CheckLogin](../docs/checklogin.md)

CreateSubAccount

DeleteSubaccount

GetAccountInfo

GetAccountPassword

GetAllAccountInfo

GetOrderDetail

GetOrderList

GetReport

GetSubAccountDetails

GetSubAccounts

GetTransHistory

SubAccountDomains

UpdateAccountInfo

UpdateCusPreferences