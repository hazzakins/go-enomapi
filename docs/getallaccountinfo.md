GetAllAccountInfo
=================

Combines GetAccountInfo, GetCustomerPaymentInfo, and GetCustomerPreferences.

Usage
-----

Use this command to access all information for an account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=GetAllAccountInfo&uid=Your Account ID&pw=Your API Token&ResponseType={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetAllAccountInfo |
| uid | string | Required | Your Account ID                                                            |
| pw       | string | Required | Your API Token |
| ResponseType | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter   | Type | Description                                                                                                                                                                  |
| --------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| OrganizationName   | string | Name of organization                                                                                                                                                             |
| JobTitle       | string | Job title                                                                                                                                                                   |
| FName         | string | First Name                                                                                                                                                                  |
| LName         | string | Last Name                                                                                                                                                                   |
| Address1       | string | Address line 1                                                                                                                                                                |
| Address2       | string | Address line 2                                                                                                                                                                |
| City         | string | City                                                                                                                                                                     |
| StateProvince     | string | State or province                                                                                                                                                               |
| StateProvinceChoice  | string | State or province choice. Expected values are: - s - State - p - Province                                                                                                                                 |
| PostalCode      | integer | Postal code                                                                                                                                                                  |
| Country        | string | Country code. Formatted in two character ISO context.                                                                                                                                            |
| Phone         | string | Telephone number                                                                                                                                                               |
| Fax          | string | Fax number                                                                                                                                                                  |
| EmailAddress     | string | Email address for Whois                                                                                                                                                            |
| EmailAddress\_Contact | string | Email address for Enom to contact you about your domain name account.                                                                                                                                    |
| UserID        | integer | Current account ID                                                                                                                                                              |
| Password       | string | Current account password                                                                                                                                                           |
| AuthQuestionType   | string | Nature of the question used for identity verification. Permitted values are: - smaiden - mother’s maiden name - sbirth - city of birth - ssocial - last 4 digits of SSN - shigh - high school - fteach - favorite teacher - fvspot - favorite vacation spot - fpet - favorite pet - fmovie - favorite movie - fbook - favorite book |
| AuthQuestionAnswer  | string | Current question answer                                                                                                                                                            |
| Account        | integer | Account number                                                                                                                                                                |
| Reseller       | string | Is this a reseller account or not. Expected values are: - True - Reseller - False - Not a reseller                                                                                                                     |
| CardType       | string | Type of credit card                                                                                                                                                              |
| CCName        | string | Cardholder's name                                                                                                                                                               |
| CreditCardNumber   | integer | The last four digits of the customer's credit card number                                                                                                                                           |
| CreditCardExpMonth  | integer | Credit card expiration month                                                                                                                                                         |
| CreditCardExpYear   | integer | Credit card expiration year                                                                                                                                                          |
| CCAddress       | string | Credit card billing address                                                                                                                                                          |
| CCZip         | integer | Credit card billing postal code                                                                                                                                                        |
| CCCountry       | string | Credit card billing country                                                                                                                                                          |
| DefPeriod       | integer | Default number of years to register a name. Expected values are from "1" to "10".                                                                                                                              |
| IRTPOptOut      | boolean | IRTPOptOut=True auto-approve any changes made in the Registrant contact for any ICANN Compliant domains.                                                                                                                  |
| VatID         | string | Valid VAT number                                                                                                                                                               |
| VatEntityType     | string | VAT Entity Type. Example: - Business - Charity                                                                                                                                                |
| PIN          | integer | Personal identification number for our technical support                                                                                                                                           |
| Enabled        | string | Is the technical support PIN enabled.                                                                                                                                                    |
| Command        | string | GetAllAccountIInfo                                                                                                                                                              |
| ErrCount       | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                                       |
| ErrX         | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                                           |
| Done         | string | "True" indicates this entire response has reached you successfully.                                                                                                                                     |

Example Output
--------------

The following query requests all account information for account "resellid" and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=GetAllAccountInfo&uid=resellid&pw=resellpw&ResponseType={Optional}
```
```
<interface-response>
<OrganizationName>Extraordinary Sales</OrganizationName>
<JobTitle>Manager</JobTitle>
<Fname>Johnny</Fname>
<Lname>Doety</Lname>
<Address1>15801 NE 24th Street</Address1>
<Address2/>
<City>Bellevue</City>
<StateProvince>WA</StateProvince>
<StateProvinceChoice>S</StateProvinceChoice>
<PostalCode>98008</PostalCode>
<Country>US</Country>
<Phone>+1.4252744500</Phone>
<PhoneExt/>
<Fax/>
<EmailAddress>[email protected]</EmailAddress>
<EmailAddress_Contact/>
<URL/>
<CreationDate>6/25/2002</CreationDate>
<GetAccountInfo>
<UserID>resellid</UserID>
<Password>resellpw</Password>
<AuthQuestionType>smaiden</AuthQuestionType>
<AuthQuestionAnswer>Jones</AuthQuestionAnswer>
<Account>217-no-0647</Account>
<Reseller>True</Reseller>
<AcceptTerms>True</AcceptTerms>
</GetAccountInfo>
<GetCustomerPaymentInfo>
<CCName>ww</CCName>
<CCNumber>************1111</CCNumber>
<CCType>VISA</CCType>
<CCMonth>04</CCMonth>
<CCYear>2017</CCYear>
<CCAddress>asd</CCAddress>
<CCZip>234</CCZip>
<CCCity>asd</CCCity>
<CCStateProvince>asd</CCStateProvince>
<CCCountry>AI</CCCountry>
<CCPhoneDial/>
<CCPhone>+United States.34534</CCPhone>
</GetCustomerPaymentInfo>
<GetCustomerPreferences>
<DefPeriod/>
</GetCustomerPreferences>
<GetCustomerSupportPin>
<pin>2619237</pin>
<enabled>True</enabled>
</GetCustomerSupportPin>
<Command>GETALLACCOUNTINFO</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod/>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.172</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 3:34:38 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>OrganizationName: </STRONG>Extraordinary Sales<br />
<STRONG>JobTitle: </STRONG>Manager<br />
<STRONG>Fname: </STRONG>Rosh<br />
<STRONG>Lname: </STRONG>Bach<br />
<STRONG>Address1: </STRONG>15801 NE 24th Street<br />
<STRONG>Address2: </STRONG><br />
<STRONG>City: </STRONG>ekm<br />
<STRONG>StateProvince: </STRONG>WA<br />
<STRONG>StateProvinceChoice: </STRONG>S<br />
<STRONG>PostalCode: </STRONG>98008<br />
<STRONG>Country: </STRONG>US<br />
<STRONG>Phone: </STRONG>+91.9544048048<br />
<STRONG>PhoneExt: </STRONG><br />
<STRONG>Fax: </STRONG><br />
<STRONG>EmailAddress: </STRONG>[email protected]<br />
<STRONG>EmailAddress_Contact: </STRONG><br />
<STRONG>URL: </STRONG><br />
<STRONG>CreationDate: </STRONG>6/25/2002<br />
<STRONG>UserID: </STRONG>resellid<br />
<STRONG>Password: </STRONG>resellpw<br />
<STRONG>AuthQuestionType: </STRONG>fpet<br />
<STRONG>AuthQuestionAnswer: </STRONG>kitty<br />
<STRONG>Account: </STRONG>217-no-0647<br />
<STRONG>Reseller: </STRONG>True<br />
<STRONG>AcceptTerms: </STRONG>True<br />
<STRONG>CCName: </STRONG>JohnDoe<br />
<STRONG>CCNumber: </STRONG>************5215<br />
<STRONG>CCType: </STRONG>MASTERCARD<br />
<STRONG>CCMonth: </STRONG>02<br />
<STRONG>CCYear: </STRONG>2016<br />
<STRONG>CCAddress: </STRONG>100 Main St.<br />
<STRONG>CCZip: </STRONG>99999<br />
<STRONG>CCCity: </STRONG>ekm<br />
<STRONG>CCStateProvince: </STRONG>WA<br />
<STRONG>CCCountry: </STRONG>us<br />
<STRONG>CCPhoneDial: </STRONG><br />
<STRONG>CCPhone: </STRONG>1.5555559999<br />
<STRONG>DefPeriod: </STRONG><br />
<STRONG>pin: </STRONG>2619237<br />
<STRONG>enabled: </STRONG>True<br />
<STRONG>PendingVerification: </STRONG>False<br />
<STRONG>Command: </STRONG>GETALLACCOUNTINFO<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG><br />
<STRONG>IsRealTimeTLD: </STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.234<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/3/2015 4:26:41 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
OrganizationName=Extraordinary Sales
JobTitle=Manager
Fname=Rosh
Lname=Bach
Address1=15801 NE 24th Street
Address2=
City=ekm
StateProvince=WA
StateProvinceChoice=S
PostalCode=98008
Country=US
Phone=+91.9544048048
PhoneExt=
Fax=
[email protected]
EmailAddress_Contact=
URL=
CreationDate=6/25/2002
UserID=resellid
Password=resellpw
AuthQuestionType=fpet
AuthQuestionAnswer=kitty
Account=217-no-0647
Reseller=True
AcceptTerms=True
CCName=JohnDoe
CCNumber=************5215
CCType=MASTERCARD
CCMonth=02
CCYear=2016
CCAddress=100 Main St.
CCZip=99999
CCCity=ekm
CCStateProvince=WA
CCCountry=us
CCPhoneDial=
CCPhone=1.5555559999
DefPeriod=
pin=2619237
enabled=True
PendingVerification=False
Command=GETALLACCOUNTINFO
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.172
Done=true
RequestDateTime=2/3/2015 4:27:19 PM
```
Related Commands
----------------

[CheckLogin](../docs/checklogin.md)

[CreateAccount](../docs/createaccount.md)

CreateSubAccount

GetAccountInfo

GetAccountPassword

GetCustomerPaymentInfo

GetOrderDetail

GetOrderList

GetReport

GetResellerInfo

GetSubAccountDetails

GetSubAccounts

GetTransHistory

RPT\_GetReport

UpdateAccountInfo

UpdateCusPreferences