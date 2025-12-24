CreateSubAccount
================

Create a subaccount.

Usage
-----

Use this command to create a new retail subaccount.

Resellers can also create subaccounts to help organize their business.

CreateSubAccount differs from CreateAccount in that CreateAccount offers credit card processing;

CreateSubAccount does not.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://www.enom.com/resellers/NewSubAccount.asp](https://www.enom.com/resellers/NewSubAccount.asp)?

The submit button calls the CreateSubAccount command.

Constraints
-----------

The query must meet the following requirements:

The account must have reseller status in eNom’s database.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter         | Status | Description                                                                                                                                           | Max Size |
| ------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID               | Required | Account login ID                                                                                                                                         | 20 |
| PW | Required | Account password | 20    |
| ResponseType          | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                                                             | 4 |
| NewUID | Required | Subaccount login ID. Permitted values are 6 to 20 characters in length; permitted characters include letters, numbers, hyphen, and underscore | 20    |
| NewPW              | Required | Subaccount password. Permitted characters are letters, numbers, hyphen, and underscore                                                                                                     | 20 |
| ConfirmPW | Required | Confirm subaccount password. Permitted characters are letters, numbers, hyphen, and underscore | 20    |
| AuthQuestionType        | Optional | Nature of the question used for identity verification. Permitted values are: smaiden mother’s maiden name sbirth city of birth ssocial last 4 digits of SSN shigh high school fteach favorite teacher fvspot favorite vacation spot fpet favorite pet fmovie favorite movie fbook favorite book | 10 |
| AuthQuestionAnswer | Optional | The answer to the question type \(mother's maiden name, city of birth, last 4 digits of social security, or high school\) | 50    |
| RegistrantAddress1       | Required | Registrant address                                                                                                                                        | 60 |
| RegistrantAddress2 | Optional | Registrant additional address info | 60    |
| RegistrantCity         | Required | Registrant city                                                                                                                                         | 60 |
| RegistrantCountry | Required | Registrant country | 60    |
| RegistrantEmailAddress     | Required | Email address for WhoIs                                                                                                                                     | 128 |
| RegistrantEmailAddress\_Contact | Optional | Email address for us to contact you about your domain name account | 128   |
| RegistrantFax          | Optional | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).                                                     | 17 |
| RegistrantFirstName | Required | Registrant first name | 60    |
| RegistrantLastName       | Required | Registrant last name                                                                                                                                       | 60 |
| RegistrantJobTitle | Optional | Registrant job title | 60    |
| RegistrantOrganizationName   | Optional | Registrant organization                                                                                                                                     | 60 |
| RegistrantPhone | Required | Registrant phone. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\) | 17    |
| RegistrantPostalCode      | Optional | Registrant postal code                                                                                                                                      | 16 |
| RegistrantStateProvince | Optional | Registrant state or province | 60    |
| RegistrantStateProvinceChoice  | Optional | Registrant state or province choice: S state P province                                                                                                                     | 1 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ------------------ | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done        | True indicates this entire response has reached you successfully. |
| StatusCustomerInfo | Returns Successful if account was created                            |
| PartyID      | Party ID for the new subaccount |
| Account | Account ID for the new subaccount                                |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query creates a new subaccount, resellsub4, and requests the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=CreateSubAccount&uid=resellid&pw=resellpw
&newuid=resellsub4&newpw=resellsub4&confirmpw=resellsub4
&RegistrantOrganizationName=Reseller%20Documents%20Inc.
&RegistrantFirstName=john&RegistrantLastName=doe
&RegistrantAddress1=111%20Main%20St.&RegistrantCity=
Hometown&RegistrantStateProvince=WA
&RegistrantStateProvinceChoice=S&RegistrantPostalCode=
98003&RegistrantCountry=United+States
&RegistrantEmailAddress=john%2Edoe%40resellerdocs
%2Ecom&RegistrantPhone=+1.5555555555
&RegistrantFax=+1.5555555556&AuthQuestionType=smaiden
&AuthQuestionAnswer=jones&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=CreateSubAccount&uid=resellid&pw=resellpw
&newuid=resellsub4&newpw=resellsub4&confirmpw=resellsub4
&RegistrantOrganizationName=Reseller%20Documents%20Inc.
&RegistrantFirstName=john&RegistrantLastName=doe
&RegistrantAddress1=111%20Main%20St.&RegistrantCity=
Hometown&RegistrantStateProvince=WA
&RegistrantStateProvinceChoice=S&RegistrantPostalCode=
98003&RegistrantCountry=United+States
&RegistrantEmailAddress=john%2Edoe%40resellerdocs
%2Ecom&RegistrantPhone=+1.5555555555
&RegistrantFax=+1.5555555556&AuthQuestionType=smaiden
&AuthQuestionAnswer=jones&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=CreateSubAccount&uid=resellid&pw=resellpw
&newuid=resellsub4&newpw=resellsub4&confirmpw=resellsub4
&RegistrantOrganizationName=Reseller%20Documents%20Inc.
&RegistrantFirstName=john&RegistrantLastName=doe
&RegistrantAddress1=111%20Main%20St.&RegistrantCity=
Hometown&RegistrantStateProvince=WA
&RegistrantStateProvinceChoice=S&RegistrantPostalCode=
98003&RegistrantCountry=United+States
&RegistrantEmailAddress=john%2Edoe%40resellerdocs
%2Ecom&RegistrantPhone=+1.5555555555
&RegistrantFax=+1.5555555556&AuthQuestionType=smaiden
&AuthQuestionAnswer=jones&ResponseType=text
```
The response confirms the successful creation of the subaccount:

```
<?xml version="1.0" ?>

 <StatusCustomerInfo>Successful</StatusCustomerInfo>

 <PartyID>{B301E3A9-1BFF-4600-9B2D-D214C29325A9}</PartyID>

 <Account>661-zw-1374</Account>

 </NewAccount>

 <Command>CREATESUBACCOUNT</Command>

 <ErrCount>0</ErrCount>

 <Server>ResellerTest</Server>

 <Site>enom</Site>

 <Done>true</Done>

 <![CDATA [ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

DuplicateLoginID: True

OrganizationName: Reseller Documents Inc.

Fname: john

Lname: doe

JobTitle:

Address1: 111 Main St.

Address2:

City: Hometown

StateProvince: WA

StateProvinceChoice: S

PostalCode: 98003

Country: US

Phone: +1.5555555555

PhoneExt:

Fax: +1.5555555556

EmailAddress: [email protected]

EmailAddress_Contact:

EmailInfo:

Reseller:

NewUID: resellsub4

NewPW: resellsub4

ConfirmPW: resellsub4

AuthQuestionType: smaiden

AuthQuestionAnswer: jones

DefPeriod:

Command: CREATESUBACCOUNT

APIType: API

Language: eng

ErrCount: 1

Err1: Failed to create new customer. LoginID already exists.

ResponseCount: 1

ResponseNumber1: 302155

ResponseString1: Validation error; duplicate; loginid

MinPeriod:

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site:

IsLockable:

IsRealTimeTLD:

TimeDifference: +0.00

ExecTime: 0.109

Done: true

RequestDateTime: 2/3/2015 3:04:35 PM
```
[email protected]

ExecTime=0.094

RequestDateTime=2/3/2015 3:05:08 PM
```