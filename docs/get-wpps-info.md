GetWPPSInfo
===========

Retrieve the ID Protect \(Whois privacy protection\) status and contact information for a domain name.

Usage
-----

Use this command to retrieve ID Protect status for a domain name.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[http://enom.com/domains/DomainDetail.asp?DomainNameID=152533676](http://enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

In the id protect section, the configure button calls the GetWPPSInfo command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.    | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| DomainName | Domain name                                           |
| SLD       | Second-level domain name \(for example, enom in enom.com\) |
| TLD | Top-level domain name \(extension\)                               |
| DomainNameID   | Domain name ID number, from our database |
| WPPSAllowed | Whether ID Protect is available for this TLD                           |
| WPPSExists    | Whether this domain has an ID Protect subscription |
| WPPSEnabled | Whether ID Protect is enabled for this domain                          |
| WPPSExpDate   | Expiration date for this domain’s ID Protect subscription |
| WPPSAutoRenew | Whether auto-renew is enabled for this domain’s ID Protect subscription             |
| WPPSPrice    | Yearly fee for this domain’s ID Protect subscription |
| ContactType | The contact type for this set of contact information                       |
| Organization   | The organization name for this contact |
| FName | First name for this contact                                   |
| LName      | Last name for this contact |
| Address1 | Address, first line, for this contact                              |
| Address2     | Address, second line, for this contact |
| City | City for this contact                                      |
| StateProvince  | State or province for this contact |
| PostalCode | Postal code for this contact                                   |
| Country     | Country for this contact |
| Phone | Phone number for this contact                                  |
| Fax       | Fax number for this contact |
| EmailAddress | Email address for this contact                                  |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests ID Protect settings for resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetWPPSInfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetWPPSInfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetWPPSInfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
In the response, the ID Protect status information, and the ErrCount value 0, confirm that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
<GetWPPSInfo>
 <domainname sld="resellerdocs" tld="com" domainnameid="152533676">
resellerdocs.com</domainname>
 <WPPSAllowed>1</WPPSAllowed>
 <WPPSExists>1</WPPSExists>
 <WPPSEnabled>1</WPPSEnabled>
 <WPPSExpDate>Nov 26, 2014</WPPSExpDate>
 <WPPSAutoRenew>No</WPPSAutoRenew>
 <WPPSPrice>$12.00</WPPSPrice>
 <contacts>
 <contact ContactType="Administrative">
  <Organization>Reseller Documents Inc.</Organization>
  <FName>John</FName>
  <LName>Doe</LName>
  <Address1>111 Main St.</Address1>
  <Address2 />
  <City>Hometown</City>
  <StateProvince>WA</StateProvince>
  <PostalCode>99999</PostalCode>
  <Country>US</Country>
  <Phone>+1.5555555555</Phone>
  <Fax>+1.5555555556</Fax>
  <EmailAddress>[email protected]</EmailAddress>
 </contact>
.
.
.
 <contact ContactType="WPPS">
  <Organization>Whois Privacy Protection Service, Inc.</Organization>
  <FName>Whois</FName>
  <LName>Agent</LName>
  <Address1>PMB 368, 14150 NE 20th St - F1</Address1>
  <Address2>C/O resellerdocs.com</Address2>
  <City>Hometown</City>
  <StateProvince>WA</StateProvince>
  <PostalCode>99999</PostalCode>
  <Country>US</Country>
  <Phone>+1.4252740657</Phone>
  <Fax>+1.4256960234</Fax>
  <EmailAddress>[email protected]</EmailAddress>
 </contact>
 </contacts>
</GetWPPSInfo>
<Command>GETWPPSINFO</Command>
<Language>en</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<ExecTime>0.140625</ExecTime>
<Done>true</Done>
<debug>
 <![CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
WPPSAllowed: 1
WPPSExists: 1
WPPSEnabled: 1
WPPSExpDate: Dec 22, 2025
WPPSAutoRenew: No
WPPSPrice: $6.00
ContactType: Administrative
AdministrativeOrganization: Extraordinary Sales
AdministrativeFName: Rosh
AdministrativeLName: Bach
AdministrativeAddress1: 15801 NE 24th Street
AdministrativeAddress2:
AdministrativeCity: ekm
AdministrativeStateProvince: WA
AdministrativePostalCode: 98008
AdministrativeCountry: US
AdministrativePhone: +91.9544048048
AdministrativeFax:
AdministrativeEmailAddress: [email protected]
ContactType: Aux Billing
Aux BillingOrganization: Novaxis Solutions
Aux BillingFName: Carl
Aux BillingLName: Boivin
Aux BillingAddress1: 400 boulevard Langelier
Aux BillingAddress2:
Aux BillingCity: Quebec
Aux BillingStateProvince: AZ
Aux BillingPostalCode: 45632
Aux BillingCountry: US
Aux BillingPhone: +1.4186942033
Aux BillingFax:
Aux BillingEmailAddress: [email protected]
ContactType: Billing
BillingOrganization: Extraordinary Sales
BillingFName: Rosh
BillingLName: Bach
BillingAddress1: 15801 NE 24th Street
BillingAddress2:
BillingCity: ekm
BillingStateProvince: WA
BillingPostalCode: 98008
BillingCountry: US
BillingPhone: +91.9544048048
BillingFax:
BillingEmailAddress: [email protected]
ContactType: Registrant
RegistrantOrganization: Novaxis Solutions
RegistrantFName: Carl
RegistrantLName: Boivin
RegistrantAddress1: 400 boulevard Langelier
RegistrantAddress2:
RegistrantCity: Quebec
RegistrantStateProvince: AZ
RegistrantPostalCode: 45632
RegistrantCountry: US
RegistrantPhone: +1.4186942033
RegistrantFax:
RegistrantEmailAddress: [email protected]
ContactType: Technical
TechnicalOrganization:
TechnicalFName: John
TechnicalLName: Smith
TechnicalAddress1: 100 Main St.
TechnicalAddress2:
TechnicalCity: Townsville
TechnicalStateProvince: WA
TechnicalPostalCode: 98033
TechnicalCountry: US
TechnicalPhone: +1.5555555555
TechnicalFax:
TechnicalEmailAddress: [email protected]
ContactType: WPPS
WPPSOrganization: Whois Privacy Protection Service, Inc.
WPPSFName: Whois
WPPSLName: Agent
WPPSAddress1: PO Box 639
WPPSAddress2: C/O resellerdocs.com
WPPSCity: Kirkland
WPPSStateProvince: WA
WPPSPostalCode: 98083
WPPSCountry: US
WPPSPhone: +1.4252740657
WPPSFax: +1.4256960234
WPPSEmailAddress: [email protected]
Command: GETWPPSINFO
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.125
Done: true
RequestDateTime: 2/4/2015 1:04:25 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
WPPSAllowed=1
WPPSExists=1
WPPSEnabled=1
WPPSExpDate=Dec 22, 2025
WPPSAutoRenew=No
WPPSPrice=$6.00
ContactType=Administrative
AdministrativeOrganization=Extraordinary Sales
AdministrativeFName=Rosh
AdministrativeLName=Bach
AdministrativeAddress1=15801 NE 24th Street
AdministrativeAddress2=
AdministrativeCity=ekm
AdministrativeStateProvince=WA
AdministrativePostalCode=98008
AdministrativeCountry=US
AdministrativePhone=+91.9544048048
AdministrativeFax=
[email protected]
ContactType=Aux Billing
Aux BillingOrganization=Novaxis Solutions
Aux BillingFName=Carl
Aux BillingLName=Boivin
Aux BillingAddress1=400 boulevard Langelier
Aux BillingAddress2=
Aux BillingCity=Quebec
Aux BillingStateProvince=AZ
Aux BillingPostalCode=45632
Aux BillingCountry=US
Aux BillingPhone=+1.4186942033
Aux BillingFax=
Aux [email protected]
ContactType=Billing
BillingOrganization=Extraordinary Sales
BillingFName=Rosh
BillingLName=Bach
BillingAddress1=15801 NE 24th Street
BillingAddress2=
BillingCity=ekm
BillingStateProvince=WA
BillingPostalCode=98008
BillingCountry=US
BillingPhone=+91.9544048048
BillingFax=
[email protected]
ContactType=Registrant
RegistrantOrganization=Novaxis Solutions
RegistrantFName=Carl
RegistrantLName=Boivin
RegistrantAddress1=400 boulevard Langelier
RegistrantAddress2=
RegistrantCity=Quebec
RegistrantStateProvince=AZ
RegistrantPostalCode=45632
RegistrantCountry=US
RegistrantPhone=+1.4186942033
RegistrantFax=
[email protected]
ContactType=Technical
TechnicalOrganization=
TechnicalFName=John
TechnicalLName=Smith
TechnicalAddress1=100 Main St.
TechnicalAddress2=
TechnicalCity=Townsville
TechnicalStateProvince=WA
TechnicalPostalCode=98033
TechnicalCountry=US
TechnicalPhone=+1.5555555555
TechnicalFax=
[email protected]
ContactType=WPPS
WPPSOrganization=Whois Privacy Protection Service, Inc.
WPPSFName=Whois
WPPSLName=Agent
WPPSAddress1=PO Box 639
WPPSAddress2=C/O resellerdocs.com
WPPSCity=Kirkland
WPPSStateProvince=WA
WPPSPostalCode=98083
WPPSCountry=US
WPPSPhone=+1.4252740657
WPPSFax=+1.4256960234
[email protected]
Command=GETWPPSINFO
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.078
Done=true
RequestDateTime=2/4/2015 1:04:55 PM
```
Related Commands
----------------

DisableServices

EnableServices

PurchaseServices

ServiceSelect