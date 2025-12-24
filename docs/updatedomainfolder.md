UpdateDomainFolder
==================

Change the name of a folder, configure or change the settings of a Magic Folder, and/or change a folder from Standard to Magic or from Magic to Standard.

Usage
-----

Use this command to change a folder from Standard to Magic or from Magic to Standard, configure or change the settings of a Magic Folder, and/or change the folder name or description of any folder.

This command lets a user choose which settings are Magic, and what those settings should be \(on, off, and so on\). Settings that can be Magic include the following:

Auto-renew

Registrar lock

Name servers

Host records

Domain password

Contact information

When a domain is put into a Magic Folder, the domain’s settings synchronize to those that are specified as Magic for the folder.

For example, if a folder is only Magic with respect to auto-renew, and auto-renew is on, all domains put into that folder have their auto-renew setting turned on. Since that is the only Magic setting for the folder, none of the domain’s other settings change.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/folder-control-panel/default.asp?FolderName=ResellFolder](https://resellertest.enom.com/domains/folder-control-panel/default.asp?FolderName=ResellFolder)

Click any of the links in the left-hand column. On the resulting pages, the save button calls the UpdateDomainFolder command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The folder must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand& uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter        | Status | Description                                                                                                                                                                                                                                                                                                                                             | Max Size |
| ------------------------------ | ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID              | Required | Account login ID                                                                                                                                                                                                                                                                                                                                           | 20 |
| PW | Required        | Account password | 20    |
| ResponseType          | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                                                                                                                                                                                                                                                               | 4 |
| FolderName | Optional        | Should this folder be Standard or Magic. Permitted values are: 0 Standard 1 Magic | 1    |
| FolderType           | Optional | New name for this existing folder. Permitted characters are letters, numbers, hyphens, underscores, and spaces; name must begin with a letter or number.                                                                                                                                                                                                                                                                      | 125 |
| NewFolderName | Optional        | Description for this folder. | 250   |
| Description          | Optional | Do you want DNS settings to be a Magic setting of this folder. Permitted values are: 0 No 1 Yes                                                                                                                                                                                                                                                                                                   | 1 |
| DNSSync | Optional        | Do you want name servers synchronized to our name servers. Permitted values are: 0 No, I will specify custom name servers 1 Yes, use eNom’s name servers | 78    |
| UseNameServer         | Optional | Custom name servers for this folder. Up to 12 may be specified. Required format is DNS1=ns1.sld.tld or dns1.sld.tld                                                                                                                                                                                                                                                                                      | 1 |
| DNSX X=1 to maximum 12 | Optional        | Do you want auto-renew to be a Magic setting of this folder. Permitted values are: 0 No 1 Yes | 1    |
| AutoRenewSync         | Optional | Auto-renew setting for this folder. Permitted values are: 0 Auto-renew off 1 Auto-renew on                                                                                                                                                                                                                                                                                                    | 1 |
| AutoRenew | Optional        | Do you want Registrar Lock to be a Magic setting of this folder. Permitted values are: 0 No 1 Yes | 1    |
| RegLockSync          | Optional | Registrar Lock setting for this folder. Registrar Lock protects you from unapproved transfers of your domains. Permitted values are: 0 Registrar Lock off 1 Registrar Lock on                                                                                                                                                                                                                                                           | 1 |
| RegLock | Optional        | Do you want domain passwords to be a Magic setting of this folder. Permitted values are: 0 No 1 Yes | 1    |
| DomainPWSync          | Optional; default is 0 | Password for all domains in this folder.                                                                                                                                                                                                                                                                                                                              | 20 |
| DomainPassword | Optional        | Do you want domain passwords to be a Magic setting of this folder. Permitted values are: 0 No 1 Yes | 1    |
| HostSync            | Optional | Host record name. Permitted values are www, @, and \*                                                                                                                                                                                                                                                                                                                        | 60 |
| HostNameX X=1 to maximum 3 | Optional        | Record type of this host record. Permitted values are: A Authoritative address. If RecordTypeX=A, AddressX must be an IP address. AAAA IPv6 address. If RecordTypeX=AAAA, AddressX must be an IPv6 address. CName Canonical name \(alias\). If RecordTypeX=CName, AddressX must be a fully qualified domain name or a host name defined in this domain. URL URL redirect. If RecordTypeX=URL, AddressX must be the exact URL of the page you want to redirect to, or an IP address, or a fully qualified domain name. Frame Frame redirect. If RecordTypeX=Frame, AddressX must be the exact URL of the page you want to redirect to, or an IP address, or a fully qualified domain name. | 5    |
| RecordTypeX X=1 to maximum 3  | Optional | Address to send this host record to.                                                                                                                                                                                                                                                                                                                                | 260 |
| AddressX X=1 to maximum 3 | Optional        | Do you want contacts to be a Magic setting of this folder. Permitted values are: 0 No 1 Yes | 1    |
| ContactSync          | Optional | Do you want this contact type to be a Magic setting of this folder. Permitted ContactTypes are Registrant, AuxBilling, Tech, or Admin. Permitted values are: 0 Do not synchronize this contact 1 Synchronize this contact For example, to make Registrant a Magic contact, use RegistrantContactSync=1                                                                                                                                                                                               | 1 |
| ContactTypeContactSync | Optional        | Use the Billing contact information for this contact. Permitted values are: 0 Do not use the Billing contact information for this contact 1 Use the Billing contact information for this contact | 60    |
| ContactTypeUseBilling     | Optional | Name of this contact’s organization                                                                                                                                                                                                                                                                                                                                 | 60 |
| ContactTypeOrganizationName | Optional        | Job title of this contact | 60    |
| ContactTypeJobTitle      | Optional | First \(given\) name of this contact                                                                                                                                                                                                                                                                                                                                 | 60 |
| ContactTypeFirstName | Optional        | Last \(family\) name of this contact | 60    |
| ContactTypeLastName      | Optional | Address, line 1, of this contact                                                                                                                                                                                                                                                                                                                                   | 60 |
| ContactTypeAddress1 | Optional        | Address, line 2, of this contact Optional | 60    |
| ContactTypeAddress2      | Optional | City of this contact                                                                                                                                                                                                                                                                                                                                         | 60 |
| ContactTypeCity | Optional        | State or province of this contact | 60    |
| ContactTypeProvince      | Optional | Is the value for ContactTypeProvince a state or a province? Permitted values are: S State P Province                                                                                                                                                                                                                                                                                                 | 1 |
| ContactTypeStateProvinceChoice | Optional        | Postal code of this contact | 16    |
| ContactTypePostalCode     | Optional | Country of this contact. Two-letter country code is a permitted value.                                                                                                                                                                                                                                                                                                              | 60 |
| ContactTypeCountry | Optional        | Phone number of this contact. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). | 17    |
| ContactTypePhone        | Optional | Phone extension of this contact.                                                                                                                                                                                                                                                                                                                                  | 8 |
| ContactTypePhoneExt | Optional        | Fax number of this contact. Required format is \+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\). | 17    |
| ContactTypeFax         | Optional | Email address of this contact. Permitted format is [\[email protected\]](../cdn-cgi/l/email-protection.html#450028242c2904212137203636053629216b312921)                                                                                                                                                                                                                                                                       | 128 |
| ContactTypeEmailAddress | Optional        | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter   | Description |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command        | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| Result        | Result of this operation. Return values are: 0 Failure 1 Success 2 New folder name already exists 3 Folder name does not exist in this account 4 Cannot convert this folder to Magic because at least one domain is already in another Magic folder 5 Cannot convert this folder to Magic because at least one domain has a registration status other than registered 6 Cannot convert this folder to Standard because at least one domain in this folder is pending processing 7 Unable to update folder 8 Invalid name server status value 9 Invalid auto-renew status value 10 Invalid registrar lock status value 11 Invalid domain password value 12 Cannot update DNS because at least one domain is pending processing 13 Cannot update auto-renew because at least one domain is pending processing 14 Cannot update registrar lock because at least one domain is pending processing 15 Cannot update domain password because at least one domain is pending processing 16 Cannot update contact information because at least one domain is pending processing 17 Cannot update host records because at least one domain is pending processing 18 Invalid host value |
| UpdateText | Encloses your input parameters in XML tags                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| FolderID       | Identification number of this folder |
| FolderName | Your name for this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| FolderDescription   | Your description of this folder |
| FolderType | Standard or Magic folder. 0 Standard 1 Magic                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| FolderStatus     | Status of this folder as a whole. Return values are: 1 Not a Magic folder 2 Folder is currently in sync \(all Magic settings have been updated to the values specified for the folder\) 3 Folder is locked pending a global edit 4 Something in the folder is out of sync |
| FolderDomainCount | Number of domains in this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| SyncDNS        | Should this folder synchronize name servers. Return values are: 0 No, DNS is not Magic for this folder 1 Yes, DNS is Magic for this folder |
| SyncDNSStatus | Status of DNS synchronization for this folder. Return values are: 1 This is not a Magic folder 2 This is a Magic folder but DNS is not set to Magic, or DNS is set to Magic and is in sync 3 DNS is set to Magic and this folder is pending synchronization 4 DNS is set to Magic and is out of sync                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| NSStatus       | Name servers for this folder. Return values are: NA This folder uses custom name servers Yes This folder uses our name servers |
| NameServer | Name server you specified                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| SyncRenew       | Should this folder synchronize auto-renew settings. Return values are: 0 No 1 Yes |
| SyncRenewStatus | Status of auto-renew synchronization for this folder. Return values are: 1 This is not a Magic folder 2 This is a Magic folder but auto-renew is not set to magic, or auto-renew is set to Magic and is in sync 3 Auto-renew is set to Magic and this folder is pending synchronization 4 Auto-renew is set to Magic and is out of sync                                                                                                                                                                                                                                                                                                                                                                                                                  |
| RenewStatus      | Auto-renew setting of this folder. Return values are: No Auto-renew is off for all domains in this folder Yes Auto-renew is on for all domains in this folder |
| SyncRegLock | Should this folder synchronize registrar lock settings. Return values are: 0 No, registrar lock is not Magic for this folder 1 Yes, registrar lock is Magic for this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| SyncRegLockStatus   | Status of registrar lock synchronization for this folder. Return values are: 1 This is not a Magic folder 2 This is a Magic folder but registrar lock is not set to Magic, or registrar lock is set to Magic and is in sync 3 Registrar lock is set to Magic and this folder is pending synchronization 4 Registrar lock is set to Magic and is out of sync |
| RegLockStatus | Registrar lock setting of this folder. Return values are: Locked Registrar lock is on for this folder Not Locked Registrar lock is off for this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| SyncDomainPwd     | Should this folder synchronize domain passwords. Return values are: 0 No, domain password is not Magic for this folder 1 Yes, domain password is Magic for this folder |
| SyncDomainPwdStatus | Status of password synchronization for this folder. Return values are: 1 This is not a Magic folder 2 This is a Magic folder but password is not set to Magic, or password is set to Magic and is in sync 3 Password is set to Magic and this folder is pending synchronization 4 Password is set to Magic and is out of sync                                                                                                                                                                                                                                                                                                                                                                                                                          |
| DomainPwdValue    | Domain password for this folder |
| SyncHosts | Should host records be Magic for this folder. Return values are: 0 No, do not synchronize host records in this folder 1 Yes, synchronize host records in this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| SyncHostsStatus    | Status of host record synchronization for this folder. Return values are: 1 This is not a Magic folder 2 This is a Magic folder but host records are not set to Magic, or host records are set to Magic and are in sync 3 Host records are set to Magic and this folder is pending synchronization 4 Host records are set to Magic and are out of sync |
| SyncContact | Are contacts Magic in this folder. Return values are: 0 No, contacts are not synchronized in this folder 1 Yes, contacts are synchronized in this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| SyncContactStatus   | Status of synchronization of this contact type. Return values are: 1 This is not a Magic folder 2 This is a Magic folder but this contact type is not set to Magic, or this contact type is set to Magic and is in sync 3 This contact type is set to Magic and this folder is pending synchronization 4 This contact type is set to Magic and is out of sync |
| PartyID | PartyID of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| OrganizationName   | Organization name of this contact |
| JobTitle | Job title of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| FName         | First \(given\) name of this contact |
| LName | Last \(family\) name of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| Address1       | Address, first line, of this contact |
| Address2 | Address, second line, of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| City         | City of this contact |
| StateProvince | State or province of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| StateProvinceChoice  | Is the StateProvince value a state or a province. Return values are: S State P Province |
| PostalCode | Postal code of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| Country        | Country of this contact |
| FullCountry | Full country name of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| Phone         | Phone number of this contact |
| Fax | Fax number of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| EmailAddress     | Email address of this contact |
| PhoneExt | Phone extension of this contact                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| SameAs        | Model for this contact |
| SyncContactRegistrant | Is Registrant contact information Magic in this folder. Return values are: 0 No, Registrant is not synchronized in this folder 1 Yes, Registrant is synchronized in this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| SyncContactAdmin   | Is Admin contact information Magic in this folder. Return values are: 0 No, Admin contact is not synchronized in this folder 1 Yes, Admin contact is synchronized in this folder |
| SyncContactTech | Is Tech contact information Magic in this folder. Return values are: 0 No, Tech contact is not synchronized in this folder 1 Yes, Tech contact is synchronized in this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| SyncContactAuxBilling | Is Auxiliary Billing contact information Magic in this folder. Return values are: 0 No, Auxiliary Billing contact is not synchronized in this folder 1 Yes, Auxiliary Billing contact is synchronized in this folder |
| FolderNameEnc | Name of this folder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| TrafficOnly      | Is this a Traffic Only account. Return values are: True This is a Traffic Only account False This is not a Traffic Only account |
| NameOnly | Is this a Name Only account. Return values are: True This is a Name Only account False This is not a Name Only account                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ParkingEnabled    | Is parking enabled for this account. Return values are: True Parking is enabled for this account False Parking is not enabled for this account |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query makes a folder Magic, adds a folder description, makes auto-renew Magic and turns autorenew on for the folder, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=UpdateDomainFolder& uid=resellid&pw=resellpw
&FolderName=ResellFolder2&FolderType=1&Description=hello
&AutoRenewSync=1&AutoRenew=1&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateDomainFolder& uid=resellid&pw=resellpw
&FolderName=ResellFolder2&FolderType=1&Description=hello
&AutoRenewSync=1&AutoRenew=1&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateDomainFolder& uid=resellid&pw=resellpw
&FolderName=ResellFolder2&FolderType=1&Description=hello
&AutoRenewSync=1&AutoRenew=1&ResponseType=text
```
In the response, a Result code 1 and an ErrCount value 0 confirm that the query was successful:

```
<?xml version="1.0" ?>

 <FolderName>ResellFolder2</FolderName>

 <Result>1</Result>

 <UpdateText><Folder><FolderDescription>test</FolderDescription>

 <FolderType>1</FolderType><Renew><SyncRenew>1</SyncRenew>

 <RenewStatus>Yes</RenewStatus></Renew></Folder></UpdateText>

 </UpdateResult>

 <FolderID>9080</FolderID>

  <![CDATA[ hello]]>

 </FolderDescription>

 <FolderType>1</FolderType>

 <FolderStatus>2</FolderStatus>

 <FolderDomainCount>0</FolderDomainCount>

  <SyncDNS>0</SyncDNS>

 </DNS>

  <SyncRenew>1</SyncRenew>

  <SyncRenewStatus>2</SyncRenewStatus>

  <RenewStatus>Yes</RenewStatus>

 </Renew>

  <SyncRegLock>0</SyncRegLock>

 </RegLock>

  <SyncDomainPwd>0</SyncDomainPwd>

 </DomainPwd>

  <SyncHosts>0</SyncHosts>

 </Hosts>

  <SyncContacts>0</SyncContacts>

 </Contacts>

 </Folder>

 <FolderNameEnc>ResellFolder2</FolderNameEnc>

 <TrafficOnly>False</TrafficOnly>

 <NameOnly>False</NameOnly>

 <ParkingEnabled>False</ParkingEnabled>

 <Command>UPDATEDOMAINFOLDER</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLER1-STG</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>8.203</ExecTime>

 <Done>true</Done>

 <![CDATA[ ]]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

FolderName: ResellFolder2

Result: 1

UpdateText: <Folder><FolderDescription>hello</FolderDescription><FolderType>1</FolderType><Renew><SyncRenew>1</SyncRenew><RenewStatus>Yes</RenewStatus></Renew></Folder>

FolderID: 12039

FolderDescription: hello

FolderStatus: 2

FolderType: 1

FolderDomainCount: 0

SyncDNS: 0

SyncDNSStatus: 2

NSStatus:

SyncRenew: 1

SyncRenewStatus: 2

RenewStatus: Yes

SyncRegLock: 0

SyncRegLockStatus: 2

RegLockStatus:

SyncDomainPwd: 0

SyncDomainPwdStatus: 2

SyncHosts: 0

SyncHostsStatus: 2

SyncContact: 0

SyncContactStatus: 2

BillingOrganizationName: Extraordinary Sales

BillingJobTitle: Manager

BillingFName: Rosh

BillingLName: Bach

BillingAddress1: 15801 NE 24th Street

BillingAddress2:

BillingCity: ekm

BillingStateProvince: WA

BillingStateProvinceChoice: S

BillingPostalCode: 98008

BillingCountry: US

BillingPhone: +91.9544048048

BillingFax:

BillingEMailAddress: [email protected]

BillingPhoneExt:

BillingSameAs:

SyncContactRegistrant: 0

RegistrantPartyID:

RegistrantOrganizationName:

RegistrantJobTitle:

RegistrantFName:

RegistrantLName:

RegistrantAddress1:

RegistrantAddress2:

RegistrantCity:

RegistrantStateProvince:

RegistrantStateProvinceChoice:

RegistrantPostalCode:

RegistrantCountry:

RegistrantPhone:

RegistrantFax:

RegistrantEMailAddress:

RegistrantPhoneExt:

RegistrantSameAs:

SyncContactAdmin: 0

AdminPartyID:

AdminOrganizationName:

AdminJobTitle:

AdminFName:

AdminLName:

AdminAddress1:

AdminAddress2:

AdminCity:

AdminStateProvince:

AdminStateProvinceChoice:

AdminPostalCode:

AdminCountry:

AdminPhone:

AdminFax:

AdminEMailAddress:

AdminPhoneExt:

AdminSameAs:

SyncContactTech: 0

TechPartyID:

TechOrganizationName:

TechJobTitle:

TechFName:

TechLName:

TechAddress1:

TechAddress2:

TechCity:

TechStateProvince:

TechStateProvinceChoice:

TechPostalCode:

TechCountry:

TechPhone:

TechFax:

TechEMailAddress:

TechPhoneExt:

TechSameAs:

SyncContactAuxBilling: 0

AuxBillingPartyID:

AuxBillingOrganizationName:

AuxBillingJobTitle:

AuxBillingFName:

AuxBillingLName:

AuxBillingAddress1:

AuxBillingAddress2:

AuxBillingCity:

AuxBillingStateProvince:

AuxBillingStateProvinceChoice:

AuxBillingPostalCode:

AuxBillingCountry:

AuxBillingPhone:

AuxBillingFax:

AuxBillingEMailAddress:

AuxBillingPhoneExt:

AuxBillingSameAs:

FolderNameEnc: ResellFolder2

TrafficOnly: False

NameOnly: False

ParkingEnabled: False

Command: UPDATEDOMAINFOLDER

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.547

Done: true

RequestDateTime: 2/9/2015 2:03:54 PM
```
;Machine is SJL0VWRESELL_T1

[email protected]

Server=SJL0VWRESELL_T1

ExecTime=0.078

RequestDateTime=2/9/2015 2:04:37 PM
```