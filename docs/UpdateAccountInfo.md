UpdateAccountInfo
=================

Update a variety of account information, for this account or a retail subaccount.

Usage
-----

Use this command to update Billing contact information for this account.

Use this command to change the password or authorization question for a retail subaccount.

Use this command to update credit card information for this account or a retail subaccount.

When you pass credit card information with this command, you must use the secure HTTPS protocol.

To add funds to a reseller account, use the RefillAccount command.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The account identified in the UID parameter must be a reseller account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=UpdateAccountInfo&uid=(Required)&pw=(Required)&AuthQuestionType=(Required)&AuthQuestionAnswer=(required)&RegistrantFirstName=(Required)&RegistrantLastName=(Required)&RegistrantAddress1=(Required)&RegistrantCity=(Required)&RegistrantStateProvinceChoice=(Required)&RegistrantStateProvince=(Required)&RegistrantEmailAddress=(Required)&RegistrantPhone=(Required)&responsetype=(Optional)
```
| Input Parameter         | Type | Status                                             | Description |
| ------------------------------- | ------ | ----------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command             | string | Required                                            | Queue\_GetDomains |
| uid | string | Required | Your Account ID                                                                                                                                         |
| pw               | string | Required                                            | Your API Token |
| NewUID | string | Required for changing the information of a retail subaccount | Retail subaccount login ID. Permitted format is alphanumeric characters. Must be unique in our system                                                                                             |
| NewPW              | string | Required for changing the password of a retail subaccount                    | Desired password for the retail subaccount for the future. Permitted values are 6 to 20 alphanumeric characters and symbols. Not permitted: space, \#, %, & |
| ConfirmNewPW | string | Required for changing the password of a retail subaccount | Confirm desired password for the retail subaccount for the future                                                                                                                |
| AuthQuestionType        | string | Required                                            | Nature of the question used for identity verification. Permitted values are: smaiden mother’s maiden name sbirth city of birth ssocial last 4 digits of SSN shigh high school fteach favorite teacher fvspot favorite vacation spot fpet favorite pet fmovie favorite movie fbook favorite book |
| AuthQuestionAnswer | string | Required | The answer to the question type \(mother's maiden name, city of birth, last 4 digits of social security or high school\)                                                                                     |
| Account             | string | Optional                                            | Account ID of the account to be updated |
| RegistrantOrganizationName | string | Required for us to apply the Web site URL | Registrant organization                                                                                                                                     |
| WebSiteURL           | string | Optional overall, but Required if we bill your customers for domain renewals and other services | URL of the domain registration Web site, for billing purposes |
| RegistrantFirstName | string | Required | Registrant first name                                                                                                                                      |
| RegistrantLastName       | string | Required                                            | Registrant last name |
| RegistrantJobTitle | string | Required if RegistrantOrganization Name is supplied | Registrant job title                                                                                                                                       |
| RegistrantAddress1       | string | Required                                            | Registrant Address |
| RegistrantAddress2 | string | Optional | Registrant additional address info                                                                                                                                |
| RegistrantCity         | string | Required                                            | Registrant city |
| RegistrantStateProvinceChoice | string | Required | Registrant state or province choice: S state P province                                                                                                                     |
| RegistrantStateProvince     | string | Required                                            | Registrant state or province |
| RegistrantPostalCode | string | Required | Registrant postal code                                                                                                                                      |
| RegistrantCountry        | string | Required                                            | Registrant country |
| RegistrantEmailAddress | string | Required | Email address for WhoIs                                                                                                                                     |
| RegistrantEmailAddress\_Contact | string | Required                                            | Email address for us to contact to you about your domain name account |
| RegistrantPhone | string | Required | Registrant phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL- encoded as a plus sign \(%2B\).                                                       |
| RegistrantFax          | string | Required if RegistrantOrganization Name is supplied                       | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\) |
| VatId | string | Optional | Valid VAT ID number                                                                                                                                       |
| VatEntityType          | string | Optional                                            | Vat Entity Type. Example: - Business - Charity |
| CardType | string | Required if updating credit card information of a retail account | Type of credit card When you pass credit card information with this command, you protocol.                                                                                                   |
| CCName             | string | Required if updating credit card information of a retail account                | Cardholder's name |
| CreditCardNumber | string | Required if updating credit card information of a retail account | Customer's credit card number                                                                                                                                  |
| CreditCardExpMonth       | string | Required if updating credit card information of a retail account                | Credit card expiration month |
| CreditCardExpYear | string | Required if updating credit card information of a retail account | Credit card expiration year                                                                                                                                   |
| CVV2              | string | Required if updating credit card information of a retail account                | Credit card verification code |
| CCAddress | string | Required if updating credit card information of a retail account | Credit card billing address                                                                                                                                   |
| CCCity             | string | Optional if updating credit card information of a retail account                | Credit card billing city |
| CCStateProvince | string | Optional if updating credit card information of a retail account | Credit card billing state or province                                                                                                                              |
| CCZip              | string | Required if updating credit card information of a retail account                | Credit card billing postal code |
| CCCountry | string | Required if updating credit card information of a retail account | Credit card billing country                                                                                                                                   |
| ResponseType          | string | Optional                                            | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| Fname | string | First Name |
| Lname | string | Last Name |
| JobTitle | string | Job Title |
| Address1 | string | Street Address Line 1 |
| Address2 | string | Street Address Line 2 |
| City | string | City |
| StateProvince | string | State or Province Name |
| StateProvinceChoice | string | Where StateProvince is a State \(S\) or a Province \(P\) |
| PostalCode | string | Postal or Zip Code |
| Country | string | Country |
| Phone | string | Phone Number |
| PhoneExt | string | Optional Phone Extension |
| Fax | string | Fax Number |
| VatId | string | Valid VAT ID number |
| VatEntityType | string | Vat Entity Type. Example: - Business - Charity |
| EmailAddress | string | Email address for WhoIs |
| EmailAddress\_Contact | string | Email address for us to contact to you about your domain name account |
| EmailInfo | string | |
| CCType | string | Type of credit card |
| CCName | string | Cardholder's name |
| CCNumber | string | Customer's credit card number |
| CCMonth | string | Credit card expiration month |
| CCYear | string | Credit card expiration year |
| CCAddress | string | Credit card billing address |
| CCCity | string | Credit card billing city |
| CCStateProvince | string | Credit card billing state or province |
| CCZip | string | Credit card billing postal code |
| CCCountry | string | Credit card billing country |
| CCPhone | string | Credit card holder's phone number |
| AuthQuestionType | string | Type of Security question |
| AuthQuestionAnswer | string | Answer to security Question |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=UpdateAccountInfo&uid=resellid&pw=resellpw
&NewUID=resellid2&NewPW=resellpw2&ConfirmNewPW=resellpw2
&AuthQuestionType=smaiden&AuthQuestionAnswer=Jones
&RegistrantAddress1=111+Main+St.&RegistrantCity=Hometown
&RegistrantCountry=United+States
&[email protected]
&RegistrantFax=+1.5555555556&RegistrantFirstName=John
&RegistrantLastName=Doe&RegistrantJobTitle=President
&RegistrantOrganizationName=Reseller+Documents+Inc.
&RegistrantPhone=+1.5555555555&RegistrantPostalCode=99999
&RegistrantStateProvince=WA&RegistrantStateProvinceChoice=S
&RegistrantNexus=United+States
&RegistrantPurpose=&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateAccountInfo&uid=resellid&pw=resellpw
&NewUID=resellid2&NewPW=resellpw2&ConfirmNewPW=resellpw2
&AuthQuestionType=smaiden&AuthQuestionAnswer=Jones
&RegistrantAddress1=111+Main+St.&RegistrantCity=Hometown
&RegistrantCountry=United+States
&[email protected]
&RegistrantFax=+1.5555555556&RegistrantFirstName=John
&RegistrantLastName=Doe&RegistrantJobTitle=President
&RegistrantOrganizationName=Reseller+Documents+Inc.
&RegistrantPhone=+1.5555555555&RegistrantPostalCode=99999
&RegistrantStateProvince=WA&RegistrantStateProvinceChoice=S
&RegistrantNexus=United+States
&RegistrantPurpose=&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateAccountInfo&uid=resellid&pw=resellpw
&NewUID=resellid2&NewPW=resellpw2&ConfirmNewPW=resellpw2
&AuthQuestionType=smaiden&AuthQuestionAnswer=Jones
&RegistrantAddress1=111+Main+St.&RegistrantCity=Hometown
&RegistrantCountry=United+States
&[email protected]
&RegistrantFax=+1.5555555556&RegistrantFirstName=John
&RegistrantLastName=Doe&RegistrantJobTitle=President
&RegistrantOrganizationName=Reseller+Documents+Inc.
&RegistrantPhone=+1.5555555555&RegistrantPostalCode=99999
&RegistrantStateProvince=WA&RegistrantStateProvinceChoice=S
&RegistrantNexus=United+States
&RegistrantPurpose=&ResponseType=text
```
```
<interface-response>
 <OrganizationName>Reseller Documents Inc.</OrganizationName>
 <Fname>John</Fname>
 <Lname>Doe</Lname>
 <JobTitle>President</JobTitle>
 <Address1>111 Main St.</Address1>
 <Address2/>
 <City>Hometown</City>
 <StateProvince>WA</StateProvince>
 <StateProvinceChoice>S</StateProvinceChoice>
 <PostalCode>99999</PostalCode>
 <Country>US</Country>
 <Phone>+1.5555555555</Phone>
 <PhoneExt/>
 <Fax>+1.5555555556</Fax>
 <EmailAddress>[email protected]</EmailAddress>
 <EmailAddress_Contact/>
 <EmailInfo/>
 <GetCustomerPaymentInfo>
 <CCType/>
 <CCName/>
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
 <Reseller>True</Reseller>
 <UserID>resellid2</UserID>
 <Password>resellpw2</Password>
 <AuthQuestionType>smaiden</AuthQuestionType>
 <AuthQuestionAnswer>Jones</AuthQuestionAnswer>
 </GetAccountInfo>
 <StatusEditContact>Fail</StatusEditContact>
 <Command>UPDATEACCOUNTINFO</Command>
 <Language>eng</Language>
 <ErrCount>1</ErrCount>
 <errors>
 <Err1>Only retail subaccounts can be modified by the parent account.</Err1>
 </errors>
 <ResponseCount>1</ResponseCount>
 <responses>
 <response>
  <ResponseNumber>313157</ResponseNumber>
  <ResponseString>Validation error; unauthorized; sub account(s)</ResponseString>
 </response>
 </responses>
 <MinPeriod/>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site/>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.641</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 4:13:13 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>OrganizationName: </ STRONG>Reseller Documents Inc.<br>
<STRONG>Fname: </ STRONG>John<br>
<STRONG>Lname: </ STRONG>Doe<br>
<STRONG>JobTitle: </ STRONG>President<br>
<STRONG>Address1: </ STRONG>111 Main St.<br>
<STRONG>Address2:</ STRONG><br>
<STRONG>City: </ STRONG>Hometown<br>
<STRONG>StateProvince: </ STRONG>WA<br>
<STRONG>StateProvinceChoice: </ STRONG>S<br>
<STRONG>PostalCode: </ STRONG>99999<br>
<STRONG>Country: </ STRONG>US<br>
<STRONG>Phone: </ STRONG>+1.5555555555<br>
<STRONG>PhoneExt:</ STRONG><br>
<STRONG>Fax: </ STRONG>+1.5555555556<br>
<STRONG>EmailAddress: </ STRONG>[email protected]<br>
<STRONG>EmailAddress_Contact:</ STRONG><br>
<STRONG>EmailInfo:</ STRONG><br>
<STRONG>CCType:</ STRONG><br>
<STRONG>CCName:</ STRONG><br>
<STRONG>CCNumber:</ STRONG><br>
<STRONG>CCMonth:</ STRONG><br>
<STRONG>CCYear:</ STRONG><br>
<STRONG>CCAddress:</ STRONG><br>
<STRONG>CCCity:</ STRONG><br>
<STRONG>CCStateProvince:</ STRONG><br>
<STRONG>CCZip:</ STRONG><br>
<STRONG>CCCountry:</ STRONG><br>
<STRONG>CCPhone:</ STRONG><br>
<STRONG>Reseller: </ STRONG>False<br>
<STRONG>UserID: </ STRONG>resellid2<br>
<STRONG>Password: </ STRONG>resellpw2<br>
<STRONG>AuthQuestionType: </ STRONG>smaiden<br>
<STRONG>AuthQuestionAnswer: </ STRONG>Jones<br>
<STRONG>StatusEditContact: </ STRONG>Fail<br>
<STRONG>Command: </ STRONG>UPDATEACCOUNTINFO<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>1<br>
<STRONG>Err1: </ STRONG>Only the parent account can update this subaccount.<br>
<STRONG>ResponseCount: </ STRONG>1<br>
<STRONG>ResponseNumber1: </ STRONG>313155<br>
<STRONG>ResponseString1: </ STRONG>Validation error; unauthorized; loginid<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site:</ STRONG><br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.734<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/9/2015 1:48:14 PM<br>
</BODY></HTML>
```
```
OrganizationName=Reseller Documents Inc.
Fname=John
Lname=Doe
JobTitle=President
Address1=111 Main St.
Address2=
City=Hometown
StateProvince=WA
StateProvinceChoice=S
PostalCode=99999
Country=US
Phone=+1.5555555555
PhoneExt=
Fax=+1.5555555556
[email protected]
EmailAddress_Contact=
EmailInfo=
CCType=
CCName=
CCNumber=
CCMonth=
CCYear=
CCAddress=
CCCity=
CCStateProvince=
CCZip=
CCCountry=
CCPhone=
Reseller=False
UserID=resellid2
Password=resellpw2
AuthQuestionType=smaiden
AuthQuestionAnswer=Jones
StatusEditContact=Fail
Command=UPDATEACCOUNTINFO
APIType=API
Language=eng
ErrCount=1
Err1=Only the parent account can update this subaccount.
ResponseCount=1
ResponseNumber1=313155
ResponseString1=Validation error; unauthorized; loginid
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.063
Done=true
RequestDateTime=2/9/2015 1:49:01 PM
```
Related Commands
----------------

CheckLogin

CreateAccount

CreateSubAccount

GetAccountInfo

GetAccountPassword

GetAllAccountInfo

GetCustomerPaymentInfo

GetOrderDetail

GetOrderList

GetReport

GetSubAccountDetails

GetSubAccounts

GetTransHistory

UpdateCusPreferences