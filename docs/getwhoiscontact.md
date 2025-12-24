GetWhoisContact
===============

Get Whois contact information for a domain name.

Usage
-----

Use this command to get contact information for any domain name registered through eNom or its resellers.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetWhoisContact&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | TLD\_GetTLD |
| UID | string | Required | Your Account ID                              |
| PW       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

If additional contact types are returned other than Registrant they will be "ContactType" Technical,Administrative or Billing.

| Output Parameter        | Type | Description                                                                               |
| ------------------------------ | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command            | string | Name of command executed                                                                         |
| ContactType          | string | Type of contact data being returned                                                                   |
| RegistrantOrganization     | string | Organization information                                                                         |
| RegistrantFName        | string | First name                                                                                |
| RegistrantLName        | string | Last name                                                                                |
| RegistrantAddress1       | string | Address line 1                                                                              |
| RegistrantAddress2       | string | Address line 2                                                                              |
| RegistrantCity         | string | City                                                                                   |
| RegistrantStateProvince    | string | State or province                                                                            |
| RegistrantPostalCode      | string | Postal Code                                                                               |
| RegistrantCountry       | string | Country                                                                                 |
| RegistrantPhone        | string | Phone number                                                                               |
| RegistrantFax         | string | Fax number                                                                                |
| RegistrantEmailAddress     | string | Email address                                                                              |
| Registrar           | string | Registrar for the name                                                                          |
| RegistrarWhoisServer      | string | Registrar whois server URL                                                                        |
| Updated-by          | string | Registrar that last updated the Whois info                                                                |
| Updated-date         | string | Date of last update                                                                           |
| Created-date         | string | Creation date for the name                                                                        |
| Registration-expiration-date | string | Current expiration date                                                                         |
| nameserverX          | string | Name servers. Indexed X when ResponseType=Text or HTML.                                                        |
| WPPSEnabled          | boolean | Whois Privacy Protection Service flag.                                                                 |
| WhoisPublicityEnabled     | boolean | Whois Publicity flag. Expected values: - True - False Note: in order WPS to be enabled \(active state\), registrant must consent the GDPR requirement for WPS product. |
| ErrCount            | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                     |
| ErrX              | string | Error messages explaining the failure. These can be presented as is back to the client.                                        |
| Done              | boolian | "True" indicates this entire response has reached you successfully.                                                   |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getwhoiscontact&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getwhoiscontact&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=getwhoiscontact&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
```html
<interface-response>
<GetWhoisContacts>
<domainname domainnameid="152533676" sld="resellerdocs" tld="com">resellerdocs.com</domainname>
<DomainExpired>false</DomainExpired>
<AboutusLink/>
<row RecordFound="1" IsExpired="0" RRProcessor="E" Registrar="eNom, Inc." LockStatus="ok" AbuseContactEmail="[email protected]" AbuseContactPhone="+1.4252744500" DomainRoid="10888527_DOMAIN_COM-VRSN" ExpDate="10 Jun 2019 15:56:56" CreateDate="25 Jun 2002 16:24:35" LastUpdatedDate="16 May 2018 12:08:48" NSStatus="NA" DnsSec="unSigned" DnsSecStatus="unSigned" ResellerOrganizationName="" ResellerURL="" AboutUsLink="1" ResellerEmailAddress="" RCOMPitch="" AbuseURL="http://wdprs.internic.net/" IanaID="48" RegistrarWhoisServer="whois.enom.com" RegistrarURL="www.enom.com" Idp="1" Wps="0"/>
<contacts>
<contact ContactType="Registrant">
<ContactId/>
<Organization>Whois Privacy Protection Service, Inc.</Organization>
<FName>Whois</FName>
<LName>Agent</LName>
<Address1>PO Box 639</Address1>
<Address2>C/O resellerdocs.com</Address2>
<City>Kirkland</City>
<StateProvince>WA</StateProvince>
<StateProvinceChoice/>
<PostalCode>98083</PostalCode>
<Country>US</Country>
<Phone>+1.4252740657</Phone>
<PhoneExt/>
<Fax>+1.4256960234</Fax>
<EmailAddress>[email protected]</EmailAddress>
<ContactRoid/>
</contact>
<contact ContactType="Administrative">
<ContactId/>
<Organization>Whois Privacy Protection Service, Inc.</Organization>
<FName>Whois</FName>
<LName>Agent</LName>
<Address1>PO Box 639</Address1>
<Address2>C/O resellerdocs.com</Address2>
<City>Kirkland</City>
<StateProvince>WA</StateProvince>
<StateProvinceChoice/>
<PostalCode>98083</PostalCode>
<Country>US</Country>
<Phone>+1.4252740657</Phone>
<PhoneExt/>
<Fax>+1.4256960234</Fax>
<EmailAddress>[email protected]</EmailAddress>
<ContactRoid/>
</contact>
<contact ContactType="Technical">
<ContactId/>
<Organization>Whois Privacy Protection Service, Inc.</Organization>
<FName>Whois</FName>
<LName>Agent</LName>
<Address1>PO Box 639</Address1>
<Address2>C/O resellerdocs.com</Address2>
<City>Kirkland</City>
<StateProvince>WA</StateProvince>
<StateProvinceChoice/>
<PostalCode>98083</PostalCode>
<Country>US</Country>
<Phone>+1.4252740657</Phone>
<PhoneExt/>
<Fax>+1.4256960234</Fax>
<EmailAddress>[email protected]</EmailAddress>
<ContactRoid/>
</contact>
<reseller/>
</contacts>
<rrp-info>
<registration-expiration-date>2019-06-10T15:56:56.00Z</registration-expiration-date>
<created-date>2002-06-25T16:24:35.00Z</created-date>
<updated-date>2018-05-16T12:08:48.00Z</updated-date>
<status>
<status>ok</status>
</status>
<domain>resellerdocs.com</domain>
<nameserver>
<nameserver>-</nameserver>
<nameserver>-</nameserver>
</nameserver>
</rrp-info>
<businesslisting>
<listing/>
</businesslisting>
<legal>
<![CDATA[
The data in this whois database is provided to you for information purposes only, that is, to assist you in obtaining information about or related to a domain name registration record. We make this information available "as is," and do not guarantee its accuracy. By submitting a whois query, you agree that you will use this data only for lawful purposes and that, under no circumstances will you use this data to: (1) enable high volume, automated, electronic processes that stress or load this whois database system providing you this information; or (2) allow, enable, or otherwise support the transmission of mass unsolicited, commercial advertising or solicitations via direct mail, electronic mail, or by telephone. The compilation, repackaging, dissemination or other use of this data is expressly prohibited without prior written consent from us.<br /><br />We reserve the right to modify these terms at any time. By submitting this query, you agree to abide by these terms.<br /><br /> Version 6.3 4/3/2002
]]>
</legal>
<WPPSEnabled>false</WPPSEnabled>
<Version>2</Version>
</GetWhoisContacts>
<Success>True</Success>
<Command>GETWHOISCONTACT</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>1.203</ExecTime>
<Done>true</Done>
<TrackingKey>12ad8dba-1a19-4937-a026-02edd3191fde</TrackingKey>
<RequestDateTime>9/6/2018 1:52:56 AM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Name1: </ STRONG>5earlyaccess<br>
<STRONG>ContactType: </ STRONG>Registrant<br>
<STRONG>RegistrantOrganization: </ STRONG>WHOIS PRIVACY PROTECTION SERVICE, INC.<br>
<STRONG>RegistrantFName: </ STRONG>WHOIS<br>
<STRONG>RegistrantLName: </ STRONG>AGENT<br>
<STRONG>RegistrantAddress1: </ STRONG>PO BOX 639<br>
<STRONG>RegistrantAddress2: </ STRONG>C/O RESELLERDOCS.COM<br>
<STRONG>RegistrantCity: </ STRONG>KIRKLAND<br>
<STRONG>RegistrantStateProvince: </ STRONG>WA<br>
<STRONG>RegistrantPostalCode: </ STRONG>98083<br>
<STRONG>RegistrantCountry: </ STRONG>US<br>
<STRONG>RegistrantPhone: </ STRONG>+1.4252740657<br>
<STRONG>RegistrantPhoneExt:</ STRONG><br>
<STRONG>RegistrantFax: </ STRONG>+1.4256960234<br>
<STRONG>RegistrantEmailAddress: </ STRONG>[email protected]<br>
<STRONG>ContactType: </ STRONG>Administrative<br>
<STRONG>AdministrativeOrganization: </ STRONG>WHOIS PRIVACY PROTECTION SERVICE, INC.<br>
<STRONG>AdministrativeFName: </ STRONG>WHOIS<br>
<STRONG>AdministrativeLName: </ STRONG>AGENT<br>
<STRONG>AdministrativeAddress1: </ STRONG>PO BOX 639<br>
<STRONG>AdministrativeAddress2: </ STRONG>C/O RESELLERDOCS.COM<br>
<STRONG>AdministrativeCity: </ STRONG>KIRKLAND<br>
<STRONG>AdministrativeStateProvince: </ STRONG>WA<br>
<STRONG>AdministrativePostalCode: </ STRONG>98083
<STRONG>AdministrativeCountry: </ STRONG>US<br>
<STRONG>AdministrativePhone: </ STRONG>+1.4252740657<br>
<STRONG>AdministrativePhoneExt:</ STRONG><br>
<STRONG>AdministrativeFax: </ STRONG>+1.4256960234<br>
<STRONG>AdministrativeEmailAddress: </ STRONG>[email protected]<br>
<STRONG>ContactType: </ STRONG>Technical<br>
<STRONG>TechnicalOrganization: </ STRONG>WHOIS PRIVACY PROTECTION SERVICE, INC.<br>
<STRONG>TechnicalFName: </ STRONG>WHOIS<br>
<STRONG>TechnicalLName: </ STRONG>AGENT<br>
<STRONG>TechnicalAddress1: </ STRONG>PO BOX 639<br>
<STRONG>TechnicalAddress2: </ STRONG>C/O RESELLERDOCS.COM<br>
<STRONG>TechnicalCity: </ STRONG>KIRKLAND<br>
<STRONG>TechnicalStateProvince: </ STRONG>WA<br>
<STRONG>TechnicalPostalCode: </ STRONG>98083<br>
<STRONG>TechnicalCountry: </ STRONG>US<br>
<STRONG>TechnicalPhone: </ STRONG>+1.4252740657<br>
<STRONG>TechnicalPhoneExt:</ STRONG><br>
<STRONG>TechnicalFax: </ STRONG>+1.4256960234<br>
<STRONG>TechnicalEmailAddress: </ STRONG>[email protected]<br>
<STRONG>Reseller:</ STRONG><br>
<STRONG>Item: </ STRONG>Value<br>
<STRONG>nameserver1: </ STRONG>DNS1.NAME-SERVICES.COM<br>
<STRONG>nameserver2: </ STRONG>DNS5.NAME-SERVICES.COM<br>
<STRONG>nameserver3: </ STRONG>DNS4.NAME-SERVICES.COM<br>
<STRONG>nameserver4: </ STRONG>DNS2.NAME-SERVICES.COM<br>
<STRONG>nameserver5: </ STRONG>DNS3.NAME-SERVICES.COM<br>
<STRONG>updated-date: </ STRONG>2015-01-27 01:17:57.000<br>
<STRONG>created-date: </ STRONG>2002-06-25 20:24:35.000<br>
<STRONG>hostCount: </ STRONG>14<br>
<STRONG>registration-expiration-date: </ STRONG>2015-08-21 17:45:35.000<br>
<STRONG>domain: </ STRONG>resellerdocs.com<br>
<STRONG>server: </ STRONG>SJL21WLGDEV05<br>
<STRONG>status: </ STRONG>ok<br>
<STRONG>DomainExpired: </ STRONG>False<br>
<STRONG>RRPCode: </ STRONG>200<br>
<STRONG>RRPText: </ STRONG>Command completed successfully<br>
<STRONG>Source: </ STRONG>ENOM<br>
<STRONG>Command: </ STRONG>GETWHOISCONTACT<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+08.00<br>
<STRONG>ExecTime: </ STRONG>0.516<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/4/2015<br> 1:01:01 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL1VWRESELL_T1
;Encoding Type is utf-8
domainname=resellerdocs.com
DomainExpired=false
AboutusLink=
row=
ContactId=
Organization=Whois Privacy Protection Service, Inc.
FName=Whois
LName=Agent
Address1=PO Box 639
Address2=C/O resellerdocs.com
City=Kirkland
StateProvince=WA
StateProvinceChoice=
PostalCode=98083
Country=US
Phone=+1.4252740657
PhoneExt=
Fax=+1.4256960234
[email protected]
ContactRoid=
ContactId=
Organization=Whois Privacy Protection Service, Inc.
FName=Whois
LName=Agent
Address1=PO Box 639
Address2=C/O resellerdocs.com
City=Kirkland
StateProvince=WA
StateProvinceChoice=
PostalCode=98083
Country=US
Phone=+1.4252740657
PhoneExt=
Fax=+1.4256960234
[email protected]
ContactRoid=
ContactId=
Organization=Whois Privacy Protection Service, Inc.
FName=Whois
LName=Agent
Address1=PO Box 639
Address2=C/O resellerdocs.com
City=Kirkland
StateProvince=WA
StateProvinceChoice=
PostalCode=98083
Country=US
Phone=+1.4252740657
PhoneExt=
Fax=+1.4256960234
[email protected]
ContactRoid=
reseller=
registration-expiration-date=2019-06-10T15:56:56.00Z
created-date=2002-06-25T16:24:35.00Z
updated-date=2018-05-16T12:08:48.00Z
status=ok
domain=resellerdocs.com
nameserver=-
nameserver=-
listing=
legal=The data in this whois database is provided to you for information purposes only, that is, to assist you in obtaining information about or related to a domain name registration record. We make this information available "as is," and do not guarantee its accuracy. By submitting a whois query, you agree that you will use this data only for lawful purposes and that, under no circumstances will you use this data to: (1) enable high volume, automated, electronic processes that stress or load this whois database system providing you this information; or (2) allow, enable, or otherwise support the transmission of mass unsolicited, commercial advertising or solicitations via direct mail, electronic mail, or by telephone. The compilation, repackaging, dissemination or other use of this data is expressly prohibited without prior written consent from us.<br /><br />We reserve the right to modify these terms at any time. By submitting this query, you agree to abide by these terms.<br /><br /> Version 6.3 4/3/2002
WPPSEnabled=false
Version=2
Success=True
Command=GETWHOISCONTACT
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl1vwresell_t1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=1.172
Done=true
TrackingKey=ca9f5e17-aaa3-450e-8247-c919ab116f43
RequestDateTime=9/6/2018 1:53:46 AM
```
Related Commands
----------------

AddContact

Contacts

GetAddressBook

GetContacts

GetExtAttributes

GetResellerInfo

GetServiceContact

Preconfigure