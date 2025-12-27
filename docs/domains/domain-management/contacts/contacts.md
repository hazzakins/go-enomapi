Contacts
========

Update contact information for a domain name.

Usage
-----

Use this command to set or change contact information for a domain.

If you want to update one contact type and leave the others as is, use this command with the "ContactType" parameter.

When you create contacts for .EU and .BE domain names, we recommend that you always provide Registrant contact information that is separate from Billing contact information. Don’t use the "same as Billing" default.

> ### Tip
>
>
>
> If the Billing and Registrant contact information are the same, we recommend changing the use or spelling of abbreviations in the street address to help our system recognize that it needs to create multiple contacts. Separating the Registrant and Billing information makes it easier to update Registrant contact information in the future.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The domain must exist in the account that is modifying it.
- Some country code TLDs require extended attributes \(parameters that are specific to the country code\). To find out whether a country code TLD requires extended parameters and what they are, run "GetExtAttributes" on the country code.
- The BILLING contact type cannot be updated with this command. To update billing information, use "UpdateAccountInfo".
- If you use the "ContactType" parameter, you can reset only one contact type.
- If you use the "ContactType" parameter, the contact type of the other parameters must match it \(for example, if you set "ContactType=AUXBILLING", you must use parameters "AuxBillingOrganizationName", "AuxBillingFirstName", and so on.
- If you don’t use the "ContactType" parameter and don’t supply new information for all contacts, those that you don’t supply will be empty in the database and will use our defaults in the interface.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=Contacts&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ContactType={Optional}&ContactTypeFirstName={Required}&ContactTypeLastName={Required}&ContactTypeOrganizationName={Required}&ContactTypeJobTitle={Optional}&ContactTypeAddress1={Required}&ContactTypeAddress2={Optional}&ContactTypeCity={Required}&ContactTypeStateProvinceChoice={Optional}&ContactTypeStateProvince={Optional}&ContactTypePostalCode={Optional}&ContactTypeCountry={Required}&ContactTypeEmailAddress={Required}&ContactTypePhone={Required}&ContactTypePhoneExt={Optional}&ContactTypeFax={Optional}&ExtendedAttributes={Required}&responsetype={Optional}
```
| Input Parameter        | Type | Status                   | Description |
| ------------------------------ | ------- | ------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command            | string | Required                  | Contacts |
| uid | string | Required | Your Account ID                                                                                                                           |
| pw               | string | Required                  | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                     |
| TLD              | string | Required                  | Top-level domain name \(extension\) |
| ContactType | string | Optional | Type of contact to add/update. If no contact parameters are supplied, all contacts are set to the "Account Billing" contact. Permitted values are: - REGISTRANT - AUXBILLING - TECH - ADMIN                                    |
| ContactTypeFirstName      | string | Required                  | First Name of ContactType that you add or update. In place of "ContactType" in the parameter name, substitute "Registrant", "AuxBilling", "Tech", or "Admin". For example, to add or update the "Registrant" contact first name, use "RegistrantFirstName=John". |
| ContactTypeLastName | string | Required | Last Name of the "ContactType" to update.                                                                                                             |
| ContactTypeOrganizationName  | string | Required                  | Organization of the "ContactType" to update. |
| ContactTypeJobTitle | string | Optional | Job Title of the "ContactType" to update.                                                                                                             |
| ContactTypeAddress1      | string | Required                  | Address of the "ContactType" to update. |
| ContactTypeAddress2 | string | Optional | Additional address information of the "ContactType" to update.                                                                                                   |
| ContactTypeCity        | string | Required                  | City of the "ContactType" to update. |
| ContactTypeStateProvinceChoice | string | Optional | State or province choice of ContactType contact. Permitted values are: - S -- State - P -- Province                                                                              |
| ContactTypeStateProvince    | string | Optional                  | State or province of the "ContactType" contact. |
| ContactTypePostalCode | string | Optional for most TLDs. Required for .org | Postal code of the "ContactType" contact.                                                                                                             |
| ContactTypeCountry       | string | Required                  | Country of the "ContactType" contact. |
| ContactTypeEmailAddress | string | Required | Email address of the "ContactType" contact.                                                                                                            |
| ContactTypePhone        | string | Required                  | Phone Number of the "ContactType" contact. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). |
| ContactTypePhoneExt | string | Optional | Phone Extension of the "ContactType" contact.                                                                                                           |
| ContactTypeFax         | string | Optional                  | Fax Number of the "ContactType" contact. |
| ExtendedAttributes | string | Required | Country codes data required by the Registry for some ccTLD domains. Use "GetExtAttributes" to determine whether this TLD requires extended attributes.                                                      |
| IRTPOptOut           | boolean | Optional                  | IRTPOptOut=True will omit the 60-day transfer lock on any changes made in the Registrant contact for any ICANN Compliant domain. |
| IRTPOptOutReason | string | Optional | Reason for IRTP OptOut                                                                                                                       |
| IRTPEmailLanguageCode     | string | Optional                  | IRTP email language option. Permitted values are: - EN – English - IT – Italian - FR – French - PT – Portuguese - ES – Spanish - DE – German |
| ResponseType | string | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML.                                                                                            |

Returned Parameters and Values
------------------------------

> ### Billing Contacts
>
>
>
> The BILLING contact type cannot be updated with this command. To update billing information, use "UpdateAccountInfo".

> ### Multiple Contact Type updates
>
>
>
> You can submit multiple contact data in a single request by omitting the "ContactType" parameter and sending additional contact data by replacing "Registrant" in the below optional parameters with the ContactTypes
>
>
>
> - AUXBILLING
> - TECH
> - ADMIN.
- If you use the "ContactType" parameter, you can reset only one contact type.
- If you use the "ContactType" parameter, the contact type of the other parameters must match it \(for example, if you set "ContactType=AUXBILLING", you must use parameters "AuxBillingOrganizationName", "AuxBillingFirstName", and so on\).
- If you don’t use the "ContactType" parameter and don’t supply new information for all contacts, those that you don’t supply will be empty in the database and will use our defaults in interfaces.
- "AUXBILLING" masks the true billing contact in WhoIs output.
- Some country code TLDs require extended attributes \(parameters that are specific to the country code\). To find out whether a country code TLD requires extended parameters and what they are, run "GetExtAttributes" on the country code TLD.
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                            |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------- |
| command     | string | Contacts                                             |
| errcount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| errX       | string | Error messages explaining the failure. These can be presented as is back to the client.     |
| ConsentStatus  | string | The Consent Level of the registrant contact, possible values outlined [here.](../docs/gdpr.md) |
| done       | string | True value indicates this entire response has reached you successfully.             |

Example Output
--------------

The following query provides new technical contact information for resellerdocs.com and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=contacts&UID=YourAccountID&PW=YourApiToken&sld=resellerdocs&tld=com&ContactType=Tech&TechFirstName=John&TechLastName=Smith&TechPhone=+1.5555555555&TechAddress1=100%20Main20St.&TechCity=Townsville&TechCountry=USA&[email protected]&ResponseType={Optional}
```
```
<interface-response>
<Command>CONTACTS</Command>
 <ConsentStatus>PENDING</ConsentStatus>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl21wresellt01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.156</ExecTime>
<Done>true</Done>
<debug/>
<TrackingKey>fc1114b9-1773-4714-ae3a-db3a882f7cab</TrackingKey>
<RequestDateTime>12/7/2011 4:56:54 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>PendingVerification: </STRONG>False<br />
<STRONG>Command: </STRONG>CONTACTS<br />
<STRONG>ConsentStatus: </STRONG>PENDING<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>3.188<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>cd974a5b-215b-431a-a777-7cb6ada03c7b<br />
<STRONG>RequestDateTime: </STRONG>2/3/2015 2:50:44 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
PendingVerification=False
Command=CONTACTS
ConsentStatus=PENDING
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.828
Done=true
TrackingKey=2506cbff-4217-42e7-a995-2386432255cb
RequestDateTime=2/3/2015 2:51:26 PM
```
Related Commands
----------------

AddContact

[GetContacts](../docs/domains/domain-management/contacts/getcontacts.md)

[GetCusPreferences](../docs/get-customer-preferences.md)

GetExtAttributes

GetServiceContact

GetWhoisContact

Preconfigure

[UpdateCusPreferences](../docs/update-customer-preferences.md)