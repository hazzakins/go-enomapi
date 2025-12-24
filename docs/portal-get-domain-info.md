Portal\GetDomainInfo
====================

Get information about a domain registered through portal.

Usage
-----

Use this command to get information about a domain registered through portal.

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
| ----------------------- | ------------------------------------------------------------------------------------------------ |
| DomainName | Domain name                                           |
| EmailAddress      | Email address of the portal account |
| ExpirationDate | Expiration date                                         |
| RegisterDate      | Registration date |
| ForeignLoginId | Portal account's login ID                                    |
| PortalDomainId     | Domain name ID |
| RegisterPrice | Registration price                                        |
| RenewPrice       | Renewal price |
| RegistrationPeriod | Registration period \(year\)                                   |
| ResellerProvisioned   | Is provisioned by the reseller? |
| ResellerProvisionedDate | Provisioned date                                         |
| IsPremium        | Is premium domain? |
| RegistrationStatus | Domain registration status                                    |
| `<domains>`       | Collections of domain activities, such as registrations, renewals and transfers. |
| orders | Collections of order histories.                                 |
| pushes         | Collections of push activities. |
| entry | Collections of domain services including nameserrver, host records and other value added items. |
| RAASettings       | RAA Information about contact update. |
| Payment | Portal account payment information.                               |
| Contacts        | Collections of domain contacts. |
| Command | Name of command executed                                     |
| ErrCount        | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done          | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query gets information about a domain registered through portal, and sends the response in XML, HTML, or Text format.

```
http://resellertest.enom.com/interface.asp?
command=portal_getdomaininfo&uid=resellid&pw=resellpw
&sld=domaintest&tld=camera&ResponseType=XML
```
```
http://resellertest.enom.com/interface.asp?
command=portal_getdomaininfo&uid=resellid&pw=resellpw
&sld=domaintest&tld=camera&ResponseType=html
```
```
http://resellertest.enom.com/interface.asp?
command=portal_getdomaininfo&uid=resellid&pw=resellpw
&sld=domaintest&tld=camera&ResponseType=text
```
The response is as follows:

```html
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Domain>
 <DomainName>domaintest.camera</DomainName>
 <EmailAddress>[email protected]</EmailAddress>
 <ExpirationDate>10/14/2015</ExpirationDate>
 <RegisterDate>10/14/2013</RegisterDate>
 <ForeignLoginId>[email protected]</ForeignLoginId>
 <PortalDomainId>155238298</PortalDomainId>
 <RegisterPrice>20.00</RegisterPrice>
 <RenewPrice>10.99</RenewPrice>
 <RegistrationPeriod>1</RegistrationPeriod>
 <ResellerProvisioned>0</ResellerProvisioned>
 <ResellerProvisionedDate />
 <IsPremium>0</IsPremium>
 <RegistrationStatus>Registered</RegistrationStatus>
</Domain>
<Transactions>
 <domains>
 <domain AutoRenew="Yes" ExpDate="2015-10-14T18:46:00" RegistrationStatus="Registered" Cred="eNom, Inc." DomainName="domaintest.camera" DomainNameID="152718298" SiteID="E" LoginID="resellid" />
 </domains>
 <orders />
 <pushes />
</Transactions>
<services>
 <entry name="dnsserver">
 <enomDNS value="NA" isDotName="NO" />
 <service changable="1">1012</service>
 <configuration changable="1" type="dns">
  <dns />
 </configuration>
 </entry>
 <entry name="wpps">
 <service changable="1">1123</service>
 </entry>
 <entry name="wbl">
 <wbl>
  <statusid>0</statusid>
  <statusdescr>Available</statusdescr>
 </wbl>
 </entry>
 <entry name="mobilizer">
 <service changable="0">1117</service>
 <mobilizer />
 </entry>
</services>
<RAASettings>
 <FirstName />
 <LastName />
 <EmailAddress />
 <VerificationStatus>Pending Suspension</VerificationStatus>
 <ValidationType>New Domain</ValidationType>
 <IsSuspended>False</IsSuspended>
 <IsPendingSuspension>True</IsPendingSuspension>
 <SuspensionDate/>
 <ToBeSuspendedDate>12/18/2013 7:35:57 AM</ToBeSuspendedDate>
 <ResultText/>
 <EmailsSent/>
</RAASettings>
<Payment>
 <CCType>Mastercard</CCType>
 <CCName>Jim Smith</CCName>
 <CCNumber>************5215</CCNumber>
 <CCExpDate>01/2017</CCExpDate>
 <CCAddress1>1st Ave</CCAddress1>
 <CCCity>Kirkland</CCCity>
 <CCZip>98033</CCZip>
 <CCStateProvince>WA</CCStateProvince>
 <CCCountry>US</CCCountry>
 <CCPhone>+1.4252744500</CCPhone>
</Payment>
<Contacts>
 <Registrant>
 <RegistrantPartyID>{3c850aea-ff6b-e311-acb1-0050568631ab}</RegistrantPartyID>
 <RegistrantOrganizationName>-</RegistrantOrganizationName>
 <RegistrantJobTitle>-</RegistrantJobTitle>
 <RegistrantFirstName>Jim</RegistrantFirstName>
 <RegistrantLastName>Smith</RegistrantLastName>
 <RegistrantAddress1>1st Ave</RegistrantAddress1>
 <RegistrantAddress2 />
 <RegistrantCity>Kirkland</RegistrantCity>
 <RegistrantCountry>US</RegistrantCountry>
 <RegistrantStateProvince>WA</RegistrantStateProvince>
 <RegistrantStateProvinceChoice>S</RegistrantStateProvinceChoice>
 <RegistrantPostalCode>98033</RegistrantPostalCode>
 <RegistrantEmailAddress>[email protected]</RegistrantEmailAddress>
 <RegistrantPhone>+1.4252744500</RegistrantPhone>
 <RegistrantPhoneExt />
 <RegistrantFax />
 </Registrant>
 <AuxBilling>
 <AuxBillingPartyID>{3c850aea-ff6b-e311-acb1-0050568631ab}</AuxBillingPartyID>
 <AuxBillingOrganizationName>-</AuxBillingOrganizationName>
 <AuxBillingJobTitle>-</AuxBillingJobTitle>
 <AuxBillingFirstName>Jim</AuxBillingFirstName>
 <AuxBillingLastName>Smith</AuxBillingLastName>
 <AuxBillingAddress1>1st Ave</AuxBillingAddress1>
 <AuxBillingAddress2 />
 <AuxBillingCity>Kirkland</AuxBillingCity>
 <AuxBillingCountry>US</AuxBillingCountry>
 <AuxBillingStateProvince>WA</AuxBillingStateProvince>
 <AuxBillingStateProvinceChoice>S</AuxBillingStateProvinceChoice>
 <AuxBillingPostalCode>98033</AuxBillingPostalCode>
 <AuxBillingEmailAddress>[email protected]</AuxBillingEmailAddress>
 <AuxBillingPhone>+1.4252744500</AuxBillingPhone>
 <AuxBillingPhoneExt />
 <AuxBillingFax />
 <UseRegInfoAbove>True</UseRegInfoAbove>
 <auxID>{3c850aea-ff6b-e311-acb1-0050568631ab}</auxID>
 <regID>{3c850aea-ff6b-e311-acb1-0050568631ab}</regID>
 </AuxBilling>
 <Technical>
 <TechnicalPartyID>{3c850aea-ff6b-e311-acb1-0050568631ab}</TechnicalPartyID>
 <TechnicalOrganizationName>-</TechnicalOrganizationName>
 <TechnicalJobTitle>-</TechnicalJobTitle>
 <TechnicalFirstName>Jim</TechnicalFirstName>
 <TechnicalLastName>Smith</TechnicalLastName>
 <TechnicalAddress1>1st Ave</TechnicalAddress1>
 <TechnicalAddress2 />
 <TechnicalCity>Kirkland</TechnicalCity>
 <TechnicalCountry>US</TechnicalCountry>
 <TechnicalStateProvince>WA</TechnicalStateProvince>
 <TechnicalStateProvinceChoice>S</TechnicalStateProvinceChoice>
 <TechnicalPostalCode>98033</TechnicalPostalCode>
 <TechnicalEmailAddress>[email protected]</TechnicalEmailAddress>
 <TechnicalPhone>+1.4252744500</TechnicalPhone>
 <TechnicalPhoneExt />
 <TechnicalFax />
 </Technical>
 <Admin>
 <AdminPartyID>{3c850aea-ff6b-e311-acb1-0050568631ab}</AdminPartyID>
 <AdminOrganizationName>-</AdminOrganizationName>
 <AdminJobTitle>-</AdminJobTitle>
 <AdminFirstName>Jim</AdminFirstName>
 <AdminLastName>Smith</AdminLastName>
 <AdminAddress1>1st Ave</AdminAddress1>
 <AdminAddress2 />
 <AdminCity>Kirkland</AdminCity>
 <AdminCountry>US</AdminCountry>
 <AdminStateProvince>WA</AdminStateProvince>
 <AdminStateProvinceChoice>S</AdminStateProvinceChoice>
 <AdminPostalCode>98033</AdminPostalCode>
 <AdminEmailAddress>[email protected]</AdminEmailAddress>
 <AdminPhone>+1.4252744500</AdminPhone>
 <AdminPhoneExt />
 <AdminFax />
 </Admin>
</Contacts>
<Success>True</Success>
<Command>PORTAL_GETDOMAININFO</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.188</ExecTime>
<Done>true</Done>
<TrackingKey>7d438b0f-3fd1-4d5d-9da8-4af45468bd76</TrackingKey>
<RequestDateTime>5/22/2014 11:17:34 AM</RequestDateTime>
</interface-response>
```
Related Commands
----------------

Portal\_GetAwardedDomains

Portal\_GetToken

Portal\_UpdateAwardedDomains