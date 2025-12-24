Preconfigure
============

Configure domain in the cart that requires extended attributes

Usage
-----

Use this command when registering domains that require extended attributes \(information required by some registries for some TLDs\). To retrieve extended attributes, use the **GetExtAttributes** command.

> ### This command operates only if you use our cart, and only if the status is active.
>
>
>
> If you use the **Purchase** command, supply the extended attributes directly to the command along with other parameters.

Optionally, you can also use this command to configure many universal attributes for domains in your cart that are in status **active**. For example *Auto Renew*, *Registrar Lock*, *IDN Codes*, *Name Servers* \(*DNS*\), *Host Records*, *Access Passwords*, and *Contacts*.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- Domain must be in the cart and in status **active**

Input Parameters
----------------

| Parameter        | Type | Status             | Description |
| ------------------------ | ------- | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command         | string | Required            | Preconfigure. |
| UID | string | Required | Your Account ID.                                                                                                                                                    |
| PW            | string | Required            | Your API Token. |
| Load | string | Required | Load the existing configuration. Permitted values: - 1: use the previously configured values - 2: submit new configuration information                                                                                         |
| *\[ExtendedAttributes\]* | object | Required for some country codes | Data required by the Registry. Use **GetExtAttributes** to return parameter names tagged as *`<Name>`* and permitted values tagged as *`<Value>`* |
| AutoRenew | integer | Required | Auto Renew flag. Automatically renew the domain before it expires. Permitted values: - 0: auto renew off - 1: auto renew on \(recommended\)                                                                                      |
| RegLock         | integer | Required            | Registrar Lock flag. Prevent unauthorized transfers. Permitted values: - 0: registrar lock off - 1: registrar lock on \(recommended\) |
| IDN*x* | string | Optional\* | *\[Required for International Domain Name\]* PUNY encoded names that use characters other than the English alphabet, numbers, and hyphen. For a list of IDN codes, use **GetIDNCodes** command.                                                            |
| OptContactReg      | string | Required            | Permitted values: - UseBilling: use the account billing contact as the Registrant contact for this domain - UseRegistrant: use the account default Registrant contact as the Registrant contact for this domain - UseExisting: use the Registrant contact information supplied in this query string |
| OptTechnical | string | Required | Permitted values: - UseBilling: use the account billing contact as the Technical contact for this domain - UseRegistrant: use the account default Registrant contact as the Technical contact for this domain - UseExisting: use the Technical contact information supplied in this query string             |
| OptAdministrative    | string | Required            | Permitted values: - UseBilling: use the account billing contact as the Administrative contact for this domain - UseRegistrant: use the account default Registrant contact as the Administrative contact for this domain - UseExisting: use the Administrative contact information supplied in this query string |
| OptContactAux | string | Required | Permitted values: - UseBilling: use the account billing contact as the Auxiliary Billing contact for this domain - UseRegistrant: use the account default Registrant contact as the Auxiliary Billing contact for this domain - UseExisting: use the Auxiliary Billing contact information supplied in this query string |
| ResponseType       | string | Optional            | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Addition parameters for custom name server \(DNS\)
--------------------------------------------------

```
&PreConfigDNS=other
&NS1=NS1.mydomain.com
&NS2=NS2.mydomain.com
&NS3=NS3.mydomain.com
```
| Parameter | Type  | Status | Description                                                                        |
| ------------ | ------ | ---------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| PreConfigDNS | string | Optional | Which name servers this domain uses. Permitted values: - Default: use our name servers - Other: use other name servers specified in the NSx parameters. |
| NSx     | string | Required\* | Required if PreConfigDNS is set to use other. Names of the name servers a domain is using. Maximum of 12 name servers. |

Addition parameters for custom host records
-------------------------------------------

> ### Host record components
>
>
>
> Parameters *HostName*, *RecordType* and *Address* are set of comma-delimited list. Please make sure all sets are paired correctly.

> ### Host record address requirements
>
>
>
> - If RecordType=A, Address must be an IP address.
> - If RecordType=AAAA, Address must be an IPv6 address.
> - If RecordType=CNAME, Address must be a fully qualified domain name or a host name defined in this domain.
> - If RecordType=URL, Address must be the exact URL of the page you redirect to, or an IP address, or a fully qualified domain name.
> - If RecordType=FRAME, Address is the actual URL, or the IP address, or the fully qualified domain name of the page you want to display when someone types mydomain.com.
> - If RecordType=MX, Address must be a fully qualified domain name or a host record name defined in this domain.
> - If RecordType=MXE, Address must be an IP address.
> - If RecordType=TXT, Address is a text record.

```
&UseHostRecords=1
&HostName=@,www,home
&RecordType=A,A,CNAME
&Address=209.185.108.165,209.185.108.165,myhome.domain.com
```
| Parameter | Type  | Status | Description                                                                                                                                                                                |
| -------------- | ------- | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| UseHostRecords | integer | Optional | Use host records provided in this query string. Permitted values: - 0: use the account's default host records - 1: use the host records provided in this string                                                                                                     |
| HostName    | string | Required\* | Host record name. |
| RecordType | string | Required\* | Host record type. Permitted values are: - A: IP address - AAAA: IPv6 address - CNAME: *Canonical Name* or alias record type, to associate a host name with another host - URL: URL redirect - FRAME: frame redirect - MX: Mail. Can be a host name under this domain name or the name of a mail server - MXE: Mail Easy \(email forwarding\) - TXT: Text \(SPF\) record |
| Address    | string | Required\* | Host record address. Please [host record address requirements](#host-record-address-requirements) note above. |

Addition parameters for custom contacts
---------------------------------------

> ### ContactType parameters
>
>
>
> - Registrant
> - AuxBilling
> - Tech
> - Admin

```
&RegistrantOrg=Enom
&RegistrantJobTitle=Developer
&RegistrantFName=First
&RegistrantLName=Last
&RegistrantAddress=5808 Lake Washington Blvd NE
&RegistrantAddress2=Suite 201
&RegistrantCity=Kirkland
&RegistrantCountry=US
&RegistrantStateProvinceChoice=S
&RegistrantStateProvince=WA
&RegistrantZip=98033
&[email protected]
&RegistrantPhone=+1.4252744500
&AdminOrg=Tucows
&AdminJobTitle=Developer
&AdminFName=First
&AdminLName=Last
&AdminAddress=96 Mowat Avenue
&AdminAddress2=
&AdminCity=Toronto
&AdminCountry=CA
&AdminStateProvinceChoice=P
&AdminStateProvince=ON
&AdminZip=M6K 3M1
&[email protected]
&AdminPhone=+1.4165350123
```
| Parameter | Type  | Status | Description                                                                                          |
| ---------------------------------- | ------- | ---------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| - ContactType\*Org | string | Optional | Organization name.                                                                                      |
| - ContactType\*JobTitle      | string | Optional  | Job title. |
| - ContactType\*FName | string | Required\* | First name.                                                                                          |
| - ContactType\*LName        | string | Required\* | Last name. |
| - ContactType\*Address | string | Required\* | Address.                                                                                           |
| - ContactType\*Address2      | string | Optional  | Address 2. |
| - ContactType\*City | string | Required\* | City.                                                                                             |
| - ContactType\*State        | string | Required\* | State. |
| - ContactType\*Province | string | Required\* | Province.                                                                                           |
| - ContactType\*StateProvinceChoice | string | Required\* | State or province choice. Permitted values: - S if contact is located in the United States - P if contact is located in a state, province or other regional type other than the United States |
| - ContactType\*Zip | string | Required\* | Postal code.                                                                                         |
| - ContactType\*Country       | string | Required\* | 2-letter country code. |
| - ContactType\*Phone | string | Required\* | Phone number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).        |
| - ContactType\*PhoneExt      | integer | Optional  | Phone extension. |
| - ContactType\*Fax | string | Optional | Fax number. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).         |
| - ContactType\*Email        | string | Required\* | Email address. |

Addition parameters for custom access password
----------------------------------------------

```
&AccessPassword1=mySecretPasswordABC
&AccessPassword2=mySecretPasswordABC
```
| Parameter | Type  | Status | Description                          |
| --------------- | ------ | ---------- | ------------------------------------------------------------- |
| AccessPassword1 | string | Optional | Domain name password.                    |
| AccessPassword2 | string | Required\* | Domain name password. Must be identical to AccessPassword1. |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed. |
| PreConfigSuccess | boolean | 1 indicates no extended attributes required. 2 indicates query failed at Registry. |
| Count | int   | Number of TLDs affected. |
| IsLockable | string | 0 indicates the domain is not lockable; 1 indicates the domain is lockable. |
| IsRealTimeTLD | string | Indicates whether this is a TLD that registers in real time. 0 indicates no; 1 indicates yes. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=preconfigure&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=us&Load=2&us_nexus=c11
&us_purpose=p3&PreConfigDNS=default
&ExpressCheckout=1&responsetype=xml
```
```
<?xml version="1.0" ?>
<interface-response>
 <PreConfigBillingInfo>
 <BillingOrganizationName>Reseller Documents Inc.</BillingOrganizationName>
 <BillingJobTitle>President</BillingJobTitle>
 <BillingFirstName>John</BillingFirstName>
 <BillingLastName>Doe</BillingLastName>
 <BillingAddress1>111 Main St.</BillingAddress1>
 <BillingAddress2 />
 <BillingCity>Hometown</BillingCity>
 <BillingStateProvince>WA</BillingStateProvince>
 <BillingStateProvinceChoice>S</BillingStateProvinceChoice>
 <BillingPostalCode>99999</BillingPostalCode>
 <BillingCountry>United States</BillingCountry>
 <BillingPhone>5555555555</BillingPhone>
 <BillingFax>5555555556</BillingFax>
 <BillingEmailAddress>[email protected]</BillingEmailAddress>
 </PreConfigBillingInfo>
 <PreConfigSuccess>2</PreConfigSuccess>
.
.
.
 <Command>PRECONFIGURE</Command>
 <ErrCount>0</ErrCount>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>False</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
</interface-response>
```
Related Commands
----------------

AddContact

[Contacts](../docs/contacts.md)

[GetContacts](../docs/getcontacts.md)

GetExtAttributes

GetWhoisContact