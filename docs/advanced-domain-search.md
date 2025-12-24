AdvancedDomainSearch
====================

Search the domains in your account.

Usage
-----

Use this command to search domains that are in your account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=search](https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=search)

The search button calls the AdvancedDomainSearch command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter  | Status | Description                                                                                                                                                                                  | Max Size |
| ------------------ | ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID        | Required | Account login ID                                                                                                                                                                               | 20 |
| PW | Required           | Account password | 20    |
| TLDList      | Optional; default is \* | Comma-delimited list of TLDs to include in search results. Also permitted is \*, which returns all TLDs that we offer.                                                                                                                           | 120 |
| SearchCriteria | Optional           | What part of domains to match? For example, SearchCriteria=Start&SLD=A returns all domains that start with A. Permitted values are: Start End In Exact | 5    |
| SLD        | Optional | String to match in domain name. Use with SearchCriteria to constrain results. For example, SearchCriteria=Start&SLD=A returns all domains that start with A. Permitted values are letters, numbers, and hyphens.                                                                             | 8000 |
| ExcludeNumbers | Optional           | Return domains with numbers in the domain name? Permitted values are: 0 Exclude domain names containing numbers 1 Include domain names containing numbers | 1    |
| ExcludeDashes   | Optional | Return domains with hyphens in the domain name? Permitted values are: 0 Exclude domain names containing hyphens 1 Include domain names containing hyphens                                                                                                           | 1 |
| ParkingStatus | Optional           | Return parked domains? Permitted values are: 0 Return domains that are not parked 1 Return domains that are parked | 1    |
| RegistrationStatus | Optional | What registration status to return? Permitted values are: Registered Hosted Expired                                                                                                                                              | 10 |
| AutoRenew | Optional           | What auto-renew status to return? Permitted values are: 0 Auto-renew off 1 Auto-renew on | 1    |
| Locked       | Optional | What registrar lock status to return? Permitted values are: 0 Registrar lock off 1 Registrar lock on                                                                                                                                     | 1 |
| CreationDate | Optional           | Date the domain was originally registered. Permitted format is MM/DD/YYYY | 10    |
| DaysTillExpires  | Optional | Return domains that are within this many days before expiration date. Permitted values are integers 0 to 65535                                                                                                                                | 7 |
| DaysExpired | Optional           | Return domains that are within this many days after expiration date. Permitted values are integers 0 to 65535 | 7    |
| NSStatus      | Optional | Return only domains that use our name servers? Permitted values are: Yes No                                                                                                                                                  | 5 |
| NameServer | Optional           | Return domains that use this name server. | 60    |
| HasIDProtect    | Optional | Return domains with ID Protect? Permitted values are: 0 Domains with no subscription to ID Protect 1 Domains with ID Protect                                                                                                                         | 1 |
| DaysUntilExpires | Optional           | Return domains that have ID Protect that’s due to expire within the number of days specified here. Permitted values are integers 0 to 65535 | 7    |
| HasPOPMail     | Optional | Return domains with subscriptions to POPmail? Permitted values are: 0 No POP mail subscription 1 Subscribed to POP mail                                                                                                                            | 1 |
| HasWebHosting | Optional           | Return domains associated with Web hosting? Permitted values are: 0 No Web hosting subscription 1 Subscribed to Web hosting | 1    |
| IncludeSubaccounts | Optional | Include domains that fit the search criteria that are in retail subaccounts? Permitted values are: 0 Exclude domain names in subaccounts 1 Include domain names in subaccounts                                                                                                | 1 |
| SubaccountLogin | Optional           | Login name of subaccounts to include 20 OrderBy Optional Permitted values are: SLD TLD NSStatus ExpDate Renew | 8    |
| StartPosition   | Optional | What number in the list to begin with in this response? For example, StartPosition=26&RecordsToReturn=25 returns domains 26 through 50 in the list.                                                                                                             | 8 |
| RecordsToReturn | Optional           | Number of domains to return in each response. For example, StartPosition=26&RecordsToReturn=25 returns domains 26 through 50 in the list. Maximum permitted value is 100. | 3    |
| FolderOption    | Optional | Retrieve domains associated with what types of folders? Permitted values are: 1 Domains that are not in any folders 2 Domains that are in at least one folder 3 Domains that are in any Magic folder 4 Domains that are in any standard folder 5 Domains that are only in a Magic folder 6 Domains that are only in a standard folder 7 Domains that are in folder FolderName | 1 |
| FolderName | Required           | when FolderOption=7 Return domains in this folder. Permitted value is a folder name. | 125   |
| FolderSyncStatus  | Optional | Synchronization status. Permitted values are: 1Domains are in the process of synchronization 2 Domains are in sync 3 At least one domain is out of sync                                                                                                           | 1 |
| MultiLang | Optional           | Permitted values are On or Off. | 3    |
| ResponseType    | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                                                                                                                                             | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| -------------------------- | ------------------------------------------------------------------------------------------------ |
| SP-TLDList | Domains included in this return match these TLDs                         |
| SP-SLD          | Domains included in this return include this character string in the SLD |
| SP-SearchCriteria | Domains included in this return include the SP-SLD value in this part of the SLD        |
| SP-ParkingStatus     | Returned domains match this parking status |
| SP-XML | XML formatting options for this return                              |
| SP-RegistrationStatus   | Returned domains match this registration status |
| SP-AutoRenew | Returned domains match this auto-renew setting                         |
| SP-Locked         | Returned domains match this registrar lock setting |
| SP-DaysTillExpires | Returned domains are within this many days of expiration                     |
| SP-DaysExpired      | Returned domains have expired in the last X days |
| SP-NSStatus | Returned domains match this name server status                          |
| SP-HostRecord       | Type Not used |
| SP-HostName | Not used                                             |
| SP-HostAddress      | Not used |
| SP-NameServer | Returned domains use this name server                              |
| SP-Has          | Returned domains match this ID Protect status |
| SP-HasPOPMail | Returned domains match this POP mail status                           |
| SP-EmailForwarding    | Not used |
| SP-HasWebHosting | Returned domains match this Web hosting status                          |
| SP-ExcludeNumbers     | If SP-ExcludeNumbers=1, returned domains do not contain numbers in the SLD |
| SP-ExcludeDashes | If SP-ExcludeDashes=1, returned domains do not contain hyphens in the SLD            |
| SP-IncludeSubAccounts   | If SP-IncludeSubAccounts=1, returned domains include domains in retail subaccounts |
| SP-SubAccountLogin | Returned domains are in this subaccount                             |
| SP-XMLResponse      | XMLResponse setting |
| SP-RecordsToReturn | This response includes this number of domains in a sorted list                  |
| SP-StartPosition     | This response starts at this number in the sorted list |
| SP-OrderBy | Sort criterion for this list                                   |
| SP-CustomerGroupName   | Not used |
| SP-FolderOption | Returned domains match this folder option                            |
| SP-FolderSyncStatus    | Returned domains match this folder sync status |
| SP-FolderName | Returned domains match this folder name                             |
| SP-EmailResultsOnly    | Send response by email only |
| SP-CreationDate | Returned domains were originally registered on this date                     |
| SP-DaysUntilExpires    | Returned domains are within X days before ID Protect |
| SP-ContactXML |                                                 |
| SP-WBLStatusID      | Returned domains have this status ID |
| SP-WBLEnabled | Returned domains have this enabled status for Business Listing                  |
| SP-WBLAutoRenew      | Returned domains have this Auto Renew status for Business Listing |
| TotalResults | Total number of domains that match these search criteria                     |
| StartPosition       | The first domain in this response is in this position in the overall sorted list of domains |
| NextPosition | The next response will start at this position in the overall sorted list             |
| MultiRRP          | Multiple RRPs in this response? |
| TLDOverride | TLD override setting                                       |
| DomainNameID        | ID number for this domain |
| SLD | SLD of this domain                                        |
| TLD            | TLD of this domain |
| AutoRenew | Auto-renew setting of this domain                                |
| ExpDate          | Expiration date of this domain |
| DomainRegistrationStatus | Registration status of this domain                                |
| DeleteType         | Delete status of this domain |
| NSStatus | Is this domain using our name servers?                              |
| FolderStatus        | Is this domain in any folders? |
| RRProcessor | Registrar credential of this domain name                             |
| RRCompanyName       | Registrar of record for this domain name |
| HasIDProtect | Does this domain have ID Protect, and is it enabled?                       |
| DomainFolderStatus     | Sync status of this domain |
| AbleToReactivate | If expired, can this name be reactivated?                            |
| IsPremiumName       | Is this a premium .tv name? |
| PremiumPrice | Price for this premium .tv name                                 |
| PremiumAboveThresholdPrice | Is the price for this premium .tv name above the premium threshold price? |
| PremiumCategory | Premium category for a .tv name                                 |
| ReactivatePrice      | Price to reactivate this expired .tv name |
| WBLStatusID | Status identifier for Business Listing for this domain name                   |
| WBLStatus         | Description of status for Business Listing |
| WBLExpDate | Expiration date for Business Listing                               |
| WBLAutoRenew        | Auto-renew setting for Business Listing |
| Command | Name of command executed                                     |
| ErrCount          | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done            | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves domains that match the specified search criteria, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=AdvancedDomainSearch&uid=resellid&pw=resellpw
&SearchCriteria=Start&SLD=r&ExcludeNumbers=1
&ExcludeDashes=1&OrderBy=ExpDate&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=AdvancedDomainSearch&uid=resellid&pw=resellpw
&SearchCriteria=Start&SLD=r&ExcludeNumbers=1
&ExcludeDashes=1&OrderBy=ExpDate&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=AdvancedDomainSearch&uid=resellid&pw=resellpw
&SearchCriteria=Start&SLD=r&ExcludeNumbers=1
&ExcludeDashes=1&OrderBy=ExpDate&responsetype=text
```
In the response, a domain list and an ErrCount value 0 indicate that the query was successful:

```
<interface-response>
<DomainSearch>
 <SearchParams>
 <SP-TLDList>*</SP-TLDList>
 <SP-SLD>r</SP-SLD>
 <SP-SearchCriteria>Start</SP-SearchCriteria>
 <SP-ParkingStatus/>
 <SP-XML>
  <DomainNames>
  <DomainName sld="r" location="Start"/>
  </DomainNames>
 </SP-XML>
 <SP-RegistrationStatus>Registered</SP-RegistrationStatus>
 <SP-AutoRenew/>
 <SP-Locked/>
 <SP-DaysTillExpires/>
 <SP-DaysExpired/>
 <SP-NsStatus/>
 <SP-HostRecordType/>
 <SP-HostName/>
 <SP-HostAddress/>
 <SP-NameServer/>
 <SP-HasIDProtect/>
 <SP-HasPOPMail/>
 <SP-EmailForwarding/>
 <SP-HasWebHosting/>
 <SP-ExcludeNumbers>True</SP-ExcludeNumbers>
 <SP-ExcludeDashes>True</SP-ExcludeDashes>
 <SP-IncludeSubAccounts/>
 <SP-SubAccountLogin/>
 <SP-XMLResponse>1</SP-XMLResponse>
 <SP-RecordsToReturn>100</SP-RecordsToReturn>
 <SP-StartPosition>1</SP-StartPosition>
 <SP-OrderBy>ExpDate</SP-OrderBy>
 <SP-CustomerGroupName/>
 <SP-FolderOption/>
 <SP-FolderSyncStatus/>
 <SP-FolderName/>
 <SP-EmailResultsOnly/>
 <SP-CreationDate/>
 <SP-DaysUntilIDProtectExpires/>
 <SP-ContactXML/>
 <SP-WBLStatusID/>
 <SP-WBLEnabled/>
 <SP-WBLAutoRenew/>
 </SearchParams>
 <TotalResults>33</TotalResults>
 <StartPosition>1</StartPosition>
 <NextPosition>1</NextPosition>
 <MultiRRP>False</MultiRRP>
 <TLDOverride>0</TLDOverride>
 <Domains>
 <Domain>
  <DomainNameID>318265464</DomainNameID>
  <SLD>rtrtretqq</SLD>
  <TLD>com</TLD>
  <AutoRenew>0</AutoRenew>
  <ExpDate>12/30/2005</ExpDate>
  <DomainRegistrationStatus>Registered
</DomainRegistrationStatus>
  <DeleteType/>
  <LoginID>resellid</LoginID>
  <AccountID>217-no-0647</AccountID>
  <NSStatus>Yes</NSStatus>
  <FolderStatus>0</FolderStatus>
  <RRProcessor>E</RRProcessor>
  <RRCompanyName>eNom, Inc.</RRCompanyName>
  <HasIDProtect>disabled</HasIDProtect>
  <IDProtectExpires/>
  <DomainFolderStatus>0</DomainFolderStatus>
  <WBLStatusID>0</WBLStatusID>
  <WBLStatus>Available</WBLStatus>
  <WBLExpDate/>
  <WBLAutoRenew/>
  <NameServers>dns1.name-services.com,
dns2.name-services.com,dns3.name-services.com,
dns4.name-services.com,dns5.name-services.com
</NameServers>
  <Vas/>
 </Domain>
 <Domain>
  <DomainNameID>318273905</DomainNameID>
  <SLD>roscoisgek</SLD>
  <TLD>com</TLD>
  <AutoRenew>1</AutoRenew>
  <ExpDate>5/13/2006</ExpDate>
  <DomainRegistrationStatus>Registered
</DomainRegistrationStatus>
  <DeleteType/>
  <LoginID>resellid</LoginID>
  <AccountID>217-no-0647</AccountID>
  <NSStatus>Yes</NSStatus>
  <FolderStatus>0</FolderStatus>
  <RRProcessor>E</RRProcessor>
  <RRCompanyName>eNom, Inc.</RRCompanyName>
  <HasIDProtect>enabled</HasIDProtect>
  <IDProtectExpires>7/29/2007</IDProtectExpires>
  <DomainFolderStatus>0</DomainFolderStatus>
  <WBLStatusID>0</WBLStatusID>
  <WBLStatus>Available</WBLStatus>
  <WBLExpDate/>
  <WBLAutoRenew/>
  <NameServers>dns1.name-services.com,
dns2.name-services.com,dns3.name-services.com,
dns4.name-services.com,dns5.name-services.com
</NameServers>
  <Vas/>
 </Domain>
.
.
.
 </Domains>
</DomainSearch>
<Command>ADVANCEDDOMAINSEARCH</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod/>
<MaxPeriod>10</MaxPeriod>
<Server>sjl21wresellt01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.141</ExecTime>
<Done>true</Done>
<debug/>
<TrackingKey>340cd124-a655-41bd-820c-d3e1b731a1b3</TrackingKey>
<RequestDateTime>12/7/2011 3:29:08 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
SP-TLDList: *
SP-SLD: r
SP-SearchCriteria: Start
SP-ParkingStatus:
DomainName:
SP-RegistrationStatus: Registered
SP-AutoRenew:
SP-Locked:
SP-DaysTillExpires:
SP-DaysExpired:
SP-NsStatus:
SP-HostRecordType:
SP-HostName:
SP-HostAddress:
SP-NameServer:
SP-HasIDProtect:
SP-HasPOPMail:
SP-EmailForwarding:
SP-HasWebHosting:
SP-ExcludeNumbers: True
SP-ExcludeDashes: True
SP-IncludeSubAccounts:
SP-SubAccountLogin:
SP-XMLResponse: 1
SP-RecordsToReturn: 100
SP-StartPosition: 1
SP-OrderBy: ExpDate
SP-CustomerGroupName:
SP-FolderOption:
SP-FolderSyncStatus:
SP-FolderName:
SP-EmailResultsOnly:
SP-CreationDate:
SP-DaysUntilIDProtectExpires:
SP-ContactXML:
SP-WBLStatusID:
SP-WBLEnabled:
SP-WBLAutoRenew:
TotalResults: 112
StartPosition: 1
NextPosition: 101
MultiRRP: False
TLDOverride: 0
DomainNameID1: 318265464
SLD1: rtrtretqq
TLD1: com
AutoRenew1: 0
ExpDate1: 12/30/2005
DomainRegistrationStatus1: Registered
DeleteType1:
LoginID1: resellid
AccountID1: 217-no-0647
IsProvisioned1: false
NSStatus1: NA
FolderStatus1: 0
RRProcessor1: E
RRCompanyName1: eNom, Inc.
HasIDProtect1: disabled
IDProtectExpires1:
DomainFolderStatus1: 0
WBLStatusID1: 0
WBLStatus1: Available
WBLExpDate1:
WBLAutoRenew1:
NameServers1:
Vas1:
DomainNameID2: 318273905
SLD2: roscoisgek
TLD2: com
AutoRenew2: 1
ExpDate2: 5/13/2006
.
.
.
count: 100
Command: ADVANCEDDOMAINSEARCH
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.531
Done: true
TrackingKey: d94fd92e-5cd3-4397-9c7d-dff73df0af56
RequestDateTime: 2/3/2015 12:37:24 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
SP-TLDList=*
SP-SLD=r
SP-SearchCriteria=Start
SP-ParkingStatus=
DomainName=
SP-RegistrationStatus=Registered
SP-AutoRenew=
SP-Locked=
SP-DaysTillExpires=
SP-DaysExpired=
SP-NsStatus=
SP-HostRecordType=
SP-HostName=
SP-HostAddress=
SP-NameServer=
SP-HasIDProtect=
SP-HasPOPMail=
SP-EmailForwarding=
SP-HasWebHosting=
SP-ExcludeNumbers=True
SP-ExcludeDashes=True
SP-IncludeSubAccounts=
SP-SubAccountLogin=
SP-XMLResponse=1
SP-RecordsToReturn=100
SP-StartPosition=1
SP-OrderBy=ExpDate
SP-CustomerGroupName=
SP-FolderOption=
SP-FolderSyncStatus=
SP-FolderName=
SP-EmailResultsOnly=
SP-CreationDate=
SP-DaysUntilIDProtectExpires=
SP-ContactXML=
SP-WBLStatusID=
SP-WBLEnabled=
SP-WBLAutoRenew=
TotalResults=112
StartPosition=1
NextPosition=101
MultiRRP=False
TLDOverride=0
DomainNameID1=318265464
SLD1=rtrtretqq
TLD1=com
AutoRenew1=0
ExpDate1=12/30/2005
DomainRegistrationStatus1=Registered
DeleteType1=
LoginID1=resellid
AccountID1=217-no-0647
IsProvisioned1=false
NSStatus1=NA
FolderStatus1=0
RRProcessor1=E
RRCompanyName1=eNom, Inc.
HasIDProtect1=disabled
IDProtectExpires1=
DomainFolderStatus1=0
WBLStatusID1=0
WBLStatus1=Available
WBLExpDate1=
WBLAutoRenew1=
NameServers1=
Vas1=
DomainNameID2=318273905
SLD2=roscoisgek
TLD2=com
AutoRenew2=1
ExpDate2=5/13/2006
.
.
.
count=100
Command=ADVANCEDDOMAINSEARCH
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.172
Done=true
TrackingKey=33d1ac28-6507-4993-a57f-6b5c5d143ea6
RequestDateTime=2/3/2015 12:40:58 PM
```
Related Commands
----------------

[AddDomainFolder](../docs/adddomainfolder.md)

[UpdateDomainFolder](../docs/updatedomainfolder.md)

[AssignToDomainFolder](../docs/AssignToDomainFolder.md)

[DeleteDomainFolder](../docs/deletedomainfolder.md)

[GetDomainFolderDetail](../docs/getdomainfolderdetail.md)

[GetDomainFolderList](../docs/getdomainfolderlist.md)

[RemoveUnsyncedDomains](../docs/removeunsynceddomains.md)