AddContact
==========

Add a contact to the address book for an account.

Usage
-----

Use this command to add a contact to the address book for an account.

To retrieve the contacts in an address book, use the GetAddressBook command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid
&pw=yourpassword&paramname=paramvalue
&nextparamname=nextparamvalue
```
| Input Parameter        | Status | Description                                                                                        | Max Size |
| ----------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID              | Required | Account login ID                                                                                      | 20 |
| PW | Required | Account password | 20    |
| ResponseType         | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                          | 4 |
| RegistrantOrganizationName | Optional | Registrant organization | 60    |
| RegistrantJobTitle      | Optional | Registrant job title                                                                                    | 60 |
| RegistrantFirstName | Required | Registrant first name | 60    |
| RegistrantLastName      | Required | Registrant last name                                                                                    | 60 |
| RegistrantAddress1 | Required | Registrant address | 60    |
| RegistrantAddress2      | Optional | Registrant additional address info                                                                             | 60 |
| RegistrantCity | Required | Registrant city | 60    |
| RegistrantPostalCode     | Required | Registrant postal code                                                                                   | 16 |
| RegistrantStateProvince | Optional | Registrant state or province | 60    |
| RegistrantStateProvinceChoice | Optional | Registrant state or province choice: S state P province                                                                  | 1 |
| RegistrantCountry | Optional | Registrant country | 60    |
| RegistrantPhone        | Required | Registrant phone number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). | 20 |
| RegistrantFax | Optional | Registrant fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). | 20    |
| RegistrantEmailAddress    | Required | Email address for Whois                                                                                  | 128 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| PartyID | Party ID number for this contact, assigned by us                         |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query adds a party to the address book for account resellid and sends the response in XML format:

```
https://resellertest.enom.com/interface.asp?command=addcontact
&uid=resellid&pw=resellpw&RegistrantFirstName=john
&RegistrantLastName=doe&RegistrantAddress1=111%20Main%20Street
&RegistrantAddress2=Suite%20100&RegistrantCity=Hometown
&RegistrantStateProvince=WA&RegistrantStateProvinceChoice=state
&RegistrantPostalCode=99999&RegistrantCountry=us
&RegistrantPhone=+800.5554444&RegistrantFax=+800.5554445
&[email protected]
&responsetype=xm
```
```
https://resellertest.enom.com/interface.asp?command=addcontact
&uid=resellid&pw=resellpw&RegistrantFirstName=john
&RegistrantLastName=doe&RegistrantAddress1=111%20Main%20Street
&RegistrantAddress2=Suite%20100&RegistrantCity=Hometown
&RegistrantStateProvince=WA&RegistrantStateProvinceChoice=state
&RegistrantPostalCode=99999&RegistrantCountry=us
&RegistrantPhone=+800.5554444&RegistrantFax=+800.5554445
&[email protected]
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?command=addcontact
&uid=resellid&pw=resellpw&RegistrantFirstName=john
&RegistrantLastName=doe&RegistrantAddress1=111%20Main%20Street
&RegistrantAddress2=Suite%20100&RegistrantCity=Hometown
&RegistrantStateProvince=WA&RegistrantStateProvinceChoice=state
&RegistrantPostalCode=99999&RegistrantCountry=us
&RegistrantPhone=+800.5554444&RegistrantFax=+800.5554445
&[email protected]
&responsetype=text
```
In the response, a Party ID and an ErrCount value of 0 confirm that the query was successful:

<RegistrantPartyID>{2B3CFB2C-7125-E111-A3F0-005056BC7747}

</RegistrantPartyID>

</Contact>

<Command>ADDCONTACT</Command>

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

<ExecTime>0.625</ExecTime>

<Done>true</Done>

<RequestDateTime>12/13/2011 1:59:43 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

RegistrantPartyID: {AA211EAB-DEAB-E411-ACB1-0050568631AB}

Command: ADDCONTACT

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.484

Done: true

RequestDateTime: 2/3/2015 11:55:57 AM
```
RegistrantPartyID={3C3FE1E2-DEAB-E411-ACB1-0050568631AB}

ExecTime=0.109

RequestDateTime=2/3/2015 11:57:30 AM
```