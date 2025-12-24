GetResellerInfo
===============

Retrieve the contact information for the account holder.

Usage
-----

Use this command to retrieve contact information for a reseller account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/EditContact.asp](https://resellertest.enom.com/myaccount/EditContact.asp)

On the my info page, the Organization \(or Individual\) Information box displays the results of the GetResellerInfo command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |
| SLD | Required | Second-level domain name \(for example, enom in enom.com\) | 63    |
| TLD       | Required | Top-level domain name \(extension\)                   | 15 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ------------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done        | True indicates this entire response has reached you successfully. |
| OrganizationName | Reseller organization name                                    |
| FName        | Reseller first name |
| LName | Reseller last name                                        |
| Address1      | Address first line |
| Address2 | Address second line                                       |
| City        | City |
| StateProvince | Name of state or province                                    |
| StateProvinceChoice | S indicates StateProvince is a state; P, a province |
| PostalCode | Postal code                                           |
| Country       | Two-letter country code |
| Phone | Phone number                                           |
| Fax         | Fax number |
| EMailAddress | Email address                                          |
| PhoneExt      | Phone extension |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves reseller contact information for the specified account and domain, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetResellerInfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetResellerInfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetResellerInfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows

<OrganizationName>Extraordinary Sales</OrganizationName>

<FName>Johnny</FName>

<LName>Doety</LName>

<Address1>15801 NE 24th Street</Address1>

<Address2/>

<City>Bellevue</City>

<StateProvince>WA</StateProvince>

<StateProvinceChoice>S</StateProvinceChoice>

<PostalCode>98008</PostalCode>

<Country>US</Country>

<Phone>+1.4252744500</Phone>

<Fax/>

<EmailAddress>[email protected]</EmailAddress>

<PhoneExt/>

<Site>E</Site>

<LoginID>resellid</LoginID>

</ResellerInfo>

<Command>GETRESELLERINFO</Command>

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

<ExecTime>0.672</ExecTime>

<Done>true</Done>

<RequestDateTime>12/9/2011 2:55:03 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

OrganizationName: Extraordinary Sales

FName: Rosh

LName: Bach

Address1: 15801 NE 24th Street

City: ekm

StateProvince: WA

StateProvinceChoice: S

PostalCode: 98008

Country: US

Phone: +91.9544048048

EmailAddress: [email protected]

Site: E

LoginID: resellid

Command: GETRESELLERINFO

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.141

Done: true

RequestDateTime: 2/4/2015 12:27:38 PM
```
;Machine is SJL0VWRESELL_T

[email protected]

Server=SJL0VWRESELL_T

ExecTime=0.047

RequestDateTime=2/4/2015 12:28:00 PM
```