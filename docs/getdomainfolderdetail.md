GetDomainFolderDetail
=====================

Retrieve an extensive set of information on a domain folder.

Usage
-----

Use this command to retrieve an extensive set of information on a domain folder.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder](https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder)

Clicking a Settings link calls the GetDomainFolderDetail command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The folder must belong to this account

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
| FolderName | Required | Name of the folder you want to delete | 125   |

Returned Parameters and Values
------------------------------

| Output Parameter  | Description |
| ------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command       | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                                                                   |
| ErrX        | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                                                                                                                                                  |
| Result       | Result of this query. Values are: 0 Failure 1 Success 2 Folder does not exist in this account |
| FolderID | Identification number of this folder, used by our database                                                                                                                                                                      |
| FolderName     | Name of this folder |
| FolderDescription | Description entered by user when this folder was created                                                                                                                                                                       |
| FolderType     | Folder type. Values are: 0 regular folder \(for organizing domains\) 1 Magic Folder \(for managing domain settings\) |
| FolderStatus | Folder status. Values are: 1 Standard folder 2 Magic Folder that is in sync 3 Magic Folder—user has revised settings, but not processed yet 4 Magic Folder that is out of sync                                                                                                            |
| FolderDomainCountan | Number of domains in this folder |
| SyncDNS | ShouShould DNS settings be the same for all domains in this folder? Values are: 0 Leave DNS records for each domain as is 1 Update DNS records for all domains to this folder’s settings                                                                                                       |
| SyncDNSStatusan   | SyncSynchronization status of DNS settings. Values are: 1 In sync and eligible for updates to DNS settings 2 Newly created and eligible for updates to DNS settings 3 Pending—revisions to DNS settings are not yet processed |
| NSStatusan | NameName server status. Values are: Yes Domains use our name servers NA Domains use custom name servers                                                                                                                                               |
| SyncRenewan     | Should auto-renew settings be the same for all domains in this folder. If auto-renew is on, domains automatically renew 30 days before they expire. Values are: 0 Leave auto-renew settings for each domain as is 1 Update auto-renew settings for all domains to this folder’s setting |
| SyncRenewStatusan | SyncSynchronization status of auto-renew settings. Values are: 1 In sync and eligible for updates to auto-renew settings 2 Newly created and eligible for updates to auto-renew settings 3 Pending—revisions to auto-renew settings are not yet processed                                                                    |
| RenewStatusan    | AutoAuto-renew setting for this folder, if one is set. Values are: 0 Do not auto-renew domains in this folder 1 Automatically renew domains in this folder |
| SyncRegLockan | ShouShould registrar lock settings be the same for all domains in this folder? If registrar lock is on, domains can be transferred only after the registrant >grants permigrants permission. Values are: 0 Leave registrar lock settings for each domain as is 1 Update registrar lock settings for all domains to this folder’s setting                               |
| SyncRegLockStatus  | Synchronization status of registrar lock settings. Values are: 1 In sync and eligible for updates to registrar lock settings 2 Newly created and eligible for updates to registrar lock settings 3 Pending—revisions to registrar lock settings are not yet processed |
| RegLockStatus | Registrar lock setting for this folder, if one is set. Values are: 0 Do not lock domains in this folder 1 Automatically lock domains in this folder                                                                                                                         |
| SyncDomainPwd    | Should domain password settings be the same for all domains in this folder. If domain password is on, domains folder? If domain password is on, domains can be managed individually by their registrants, without logging in to the domain name account. Values are: 0 Leave domain password settings for each domain as is 1 Update domain password settings for all domains to the folder setting |
| SyncDomainPwdStatus | Synchronization status of domain password settings. Values are: 1 In sync and eligible for updates to domain password settings 2 Newly created and eligible for updates to domain password settings 3 Pending—revisions to domain password settings are not yet processed                                                              |
| DomainPwdValue   | Domain password setting for this folder, if one is set. Values are: 0 Do not require a password for domains in this folder 1 Require a password for all domains in this folder |
| SyncHostse | Should host records be the same for all domains in this folder? Host records are also known as third-level domain names, like [www](http://www/). Values each are: 0 Leave host records for each domain as is 1 Update host records for all domains to the folder settings                                                             |
| SyncHostsStatus   | Synchronization status of host records. Values are: 1 In sync and eligible for updates to host records 2 Newly created and eligible for updates to host records 3 Pending—revisions to host records are not yet processed |
| SyncContact | Should any contacts be reset to match the Should any contacts be reset to match the account’s Billing contact? Values are: 0 Leave contact information for each domain as is 1 Update any contacts that have a SyncContactContactType value 1                                                                             |
| SyncContactStatus  | Synchronization status of contact information. Values are: 1 In sync and eligible for updates to contact settings 2 Newly created and eligible for updates to contact settings 3 Pending—revisions to contact settings are not yet processed |
| SyncContactContact | Should the contact information for this contact type \(Registrant, Tech, Admin, or AuxBilling\) be revised to match another contact type. Values are: 0 Leave contact information for each domain as is 1 Update the contact information of this ContactType for all domains in this folder                                                     |
| PartyID       | Party ID of this contact |
| OrganizationName | Organization name of this contact                                                                                                                                                                                   |
| JobTitle      | Job title of this contact |
| FName | First \(given\) name of this contact                                                                                                                                                                                 |
| LName        | Last \(family\) name of this contact |
| Address1 | Address, first line, of this contact                                                                                                                                                                                 |
| Address2      | Address, second line, of this contact |
| City | City of this contact                                                                                                                                                                                         |
| StateProvince    | State or province of this contact |
| StateProvinceChoice | Is StateProvince value a state or a province? Values are: S State S State P Province                                                                                                                                                         |
| Country       | Country of this contact |
| Phone | Phone number of this contact                                                                                                                                                                                     |
| Fax         | Fax number of this contact |
| EmailAddress | Email address of this contact                                                                                                                                                                                     |
| PhoneExt      | Phone extension of this contact |
| SameAs | Same As setting of this contact                                                                                                                                                                                    |
| FolderNameEnc    | Name of folder |
| TrafficOnly | True indicates that this is a Traffic Only account                                                                                                                                                                          |
| NameOnly      | True indicates that this is a Name Only account |
| ParkingEnabled | True indicates that parking is enabled                                                                                                                                                                                |
| Reseller      | Customer Type. True indicates that this account is in our reseller category \(Reseller, Name Only, or Traffic Only\). |

Notes
-----

The default response format is plain text. To receive the response in HTML or XMThe default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter e parameterErr Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves details on a folder, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetDomainFolderDetail&uid=resellid
&pw=resellpw&FolderName=Favorites&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainFolderDetail&uid=resellid
&pw=resellpw&FolderName=Favorites&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainFolderDetail&uid=resellid
&pw=resellpw&FolderName=Favorites&responsetype=text
```
In the response, the presence of data and an ErrCount value 0 indicate that the query was successful:

<Result>1</Result>

<FolderID>9140</FolderID>

<FolderName>Favorites</FolderName>

<FolderDescription>FolderForMyFavoriteNames</FolderDescription>

<FolderType>1</FolderType>

<FolderStatus>2</FolderStatus>

<FolderDomainCount>0</FolderDomainCount>

<SyncDNS>1</SyncDNS>

<SyncDNSStatus>2</SyncDNSStatus>

<NSStatus>NA</NSStatus>

<NameServer>ns769.hostgator.com</NameServer>

<NameServer>ns770.hostgator.com</NameServer>

</DNS>

<SyncRenew>0</SyncRenew>

<SyncRenewStatus>2</SyncRenewStatus>

<RenewStatus/>

</Renew>

<SyncRegLock>0</SyncRegLock>

<SyncRegLockStatus>2</SyncRegLockStatus>

<RegLockStatus/>

</RegLock>

<SyncDomainPwd>0</SyncDomainPwd>

<SyncDomainPwdStatus>2</SyncDomainPwdStatus>

<DomainPwdValue></DomainPwdValue>

</DomainPwd>

<SyncHosts>0</SyncHosts>

<SyncHostsStatus>2</SyncHostsStatus>

</Hosts>

<SyncContact>0</SyncContact>

<SyncContactStatus>2</SyncContactStatus>

<PartyID>BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629</PartyID>

<OrganizationName>Extraordinary Sales</OrganizationName>

<JobTitle>Manager</JobTitle>

<FName>Johnny</FName>

<LName>Doety</LName>

<Address1>15801 NE 24th Street</Address1>

<Address2></Address2>

<City>Bellevue</City>

<StateProvince>WA</StateProvince>

<StateProvinceChoice>S</StateProvinceChoice>

<PostalCode>98008</PostalCode>

<Country>US</Country>

<FullCountry>United States</FullCountry>

<Phone>+1.4252744500</Phone>

<Fax></Fax>

<EmailAddress>[email protected]</EmailAddress>

<PhoneExt></PhoneExt>

<SameAs/>

</BillingContact>

<SyncContactRegistrant>0</SyncContactRegistrant>

<PartyID/>

<OrganizationName></OrganizationName>

<JobTitle></JobTitle>

<FName></FName>

<LName></LName>

<Address1></Address1>

<City></City>

<StateProvince></StateProvince>

<StateProvinceChoice></StateProvinceChoice>

<PostalCode></PostalCode>

<Country></Country>

<FullCountry></FullCountry>

<Phone></Phone>

<EmailAddress></EmailAddress>

</RegistrantContact>

<SyncContactAdmin>0</SyncContactAdmin>

</AdminContact>

<SyncContactTech>0</SyncContactTech>

</TechContact>

<SyncContactAuxBilling>0</SyncContactAuxBilling>

</AuxBillingContact>

</Contacts>

</Folder>

<FolderNameEnc>Favorites</FolderNameEnc>

<TrafficOnly>False</TrafficOnly>

<NameOnly>False</NameOnly>

<ParkingEnabled>False</ParkingEnabled>

<Command>GETDOMAINFOLDERDETAIL</Command>

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

<ExecTime>0.234</ExecTime>

<Done>true</Done>

<RequestDateTime>12/8/2011 4:14:10 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Result: 1

FolderNameEnc: Favorites

TrafficOnly: False

NameOnly: False

ParkingEnabled: False

Command: GETDOMAINFOLDERDETAIL

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.234

Done: true

RequestDateTime: 2/11/2015 2:03:56 PM
```
;Machine is SJL0VWRESELL_T

Server=SJL0VWRESELL_T

ExecTime=0.078

RequestDateTime=2/11/2015 2:05:10 PM
```