GetExtendInfo
=============

Get extension \(renewal\) information.

Usage
-----

Use this command to display information about the expiration date of a domain name.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

On the domain control panel page, the add years button calls the GetExtendInfo command.

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

| Output Parameter | Description                                                                                   |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| RegistrarHold | The enabled status of a domain. If RegistrarHold is True, the domain is disabled.                                               |
| Expiration      | Current expiration date. |
| MaxExtension | Maximum \# of years that can be added.                                                                     |
| CCAuthorized     | Credit card transaction successfully authorized. |
| Price | Price for one-year extension for this type of TLD.                                                              |
| AvailableBalance   | Available balance in this account. |
| DefPeriod | Account-level setting for default renewal period for domain name registrations                                                 |
| AllowDNS       | Account-level setting—True allows domain name servers other than eNom’s |
| ShowPopups | Account-level setting—True shows popup menus                                                                  |
| AutoRenew      | Account-level setting—True automatically renews the domain registration 30 days before it expires |
| RegLock | Account-level setting—True requires the account holder’s permissiont to transfer the domain to another registrar                                |
| AutoPakRenew     | Account-level setting—True automatically renews POP paks 30 days before they expire |
| UseDNS | Account-level setting—True uses eNom’s name servers                                                              |
| ResellerStatus    | Account-level setting—Reseller status of this account |
| RenewalSetting | Account-level setting 0 Send no renewal reminder when domain expiration approaches 1 Send renewal reminder by email 2 Contact customer, renew registration, and charge account |
| RenewalBCC      | Account-level setting— 1 Send blind carbon copy of renewal email to reseller 0 Do not send blind carbon copy to reseller |
| RenewalURLForward | Account-level setting—True automatically renews URL forwarding 30 days before the subscription expires                                     |
| RenewalEmailForward | Account-level setting—True automatically renews email forwarding 30 days before the subscription expires |
| MailNumLimit | Account-level setting—True automatically renews email forwarding 30 days before the subscription expires                                    |
| IDProtect      | Account-level setting—True automatically renews ID Protect 30 days before the subscription expires |
| DefIDProtectRenew | Account-level setting—True automatically renews ID Protect 30 days before the subscription expires                                       |
| HostNameX      | Account-level setting—Name of default host record |
| AddressX | Account-level setting—Default host record address                                                               |
| RecordTypeX     | Account-level setting—Default host record type |
| DefaultHostRecordOwn | Indicates whether this domain is associated with a Web Hosting account                                                     |
| DefaultHostRecordOwn | Account-level setting—By default, does this account assign user-defined host records? |
| UseOurDNS | Account-level setting—By default, does this account assign our DNS to domain names?                                              |
| DNSX         | Account-level setting—Default name servers |
| AcceptTerms | True has signed a credit card processing agreement with eNom                                                          |
| URL         | URL for reseller site, used in email notices |
| Balance | Current balance in this account.                                                                        |
| Command       | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                        |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                                       |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests renewal information for resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getextendinfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=getextendinfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getextendinfo&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows:

```
<interface-response>
<RegistrarHold>False</RegistrarHold>
<Expiration>6/10/2019</Expiration>
<MaxExtension>2</MaxExtension>
<MinAllowed>1</MinAllowed>
<CCAuthorized>True</CCAuthorized>
<Price>8.95</Price>
<Balance>2,491.60</Balance>
<AvailableBalance>2,106.75</AvailableBalance>
<CustomerPrefs>
 <DefPeriod>1</DefPeriod>
 <AllowDNS>False</AllowDNS>
 <ShowPopups>False</ShowPopups>
 <AutoRenew>False</AutoRenew>
 <RegLock>False</RegLock>
 <AutoPakRenew>False</AutoPakRenew>
 <UseDNS>True</UseDNS>
 <ResellerStatus/>
 <RenewalSetting>1</RenewalSetting>
 <RenewalBCC>1</RenewalBCC>
 <RenewalURLForward>False</RenewalURLForward>
 <RenewalEmailForward>False</RenewalEmailForward>
 <MailNumLimit>100</MailNumLimit>
 <IDProtect>False</IDProtect>
 <DefIDProtectRenew>False</DefIDProtectRenew>
 <DefWBLRenew>False</DefWBLRenew>
 <NameJetSales>False</NameJetSales>
 <defaulthostrecords>
 <hostrecord hostname="@" address="69.25.142.5" recordtype="A"/>
 <hostrecord hostname="*" address="69.25.142.5" recordtype="A"/>
 <hostrecord hostname="www" address="69.25.142.5" recordtype="A"/>
 <hostrecord hostname="" address="" recordtype="A"/>
 <hostrecord hostname="" address="" recordtype="A"/>
 </defaulthostrecords>
 <defaulthostrecordown>False</defaulthostrecordown>
 <UseOurDNS>False</UseOurDNS>
 <NameServers>
 <DNS1>ns1.nicmus.com</DNS1>
 <DNS2>ns2.nicmus.com</DNS2>
 <DNS3>ns3.nicmus.com</DNS3>
 </NameServers>
</CustomerPrefs>
<CustomerInformation>
 <AcceptTerms>True</AcceptTerms>
 <URL/>
 <ParentAccount>000-00-0000</ParentAccount>
 <ParentLogin>Main-eNom</ParentLogin>
 <NoService>False</NoService>
 <BulkRegLimit>100</BulkRegLimit>
 <Account>217-no-0647</Account>
</CustomerInformation>
<RRPCode>200</RRPCode>
<RRPText/>
<Command>GETEXTENDINFO</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+08.00</TimeDifference>
<ExecTime>0.393</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 4:54:28 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```bash
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RegistrarHold: False
Expiration: 6/10/2019
MaxExtension: 5
MinAllowed: 1
CCAuthorized: True
Price: 8.95
Balance: -1,832.23
AvailableBalance: 6,193.37
DefPeriod: 1
AllowDNS: False
ShowPopups: False
AutoRenew: True
RegLock: False
AutoPakRenew: False
UseDNS: True
ResellerStatus:
RenewalSetting: 1
RenewalBCC: 0
RenewalURLForward: False
RenewalEmailForward: False
MailNumLimit: 100
IDProtect: False
DefIDProtectRenew: False
DefWBLRenew: False
NameJetSales: False
hostname1: @
address1: 85.92.87.178
recordtype1: A
hostname2: *
address2: 85.92.87.179
recordtype2: A
hostname3: www
address3: 85.92.87.180
recordtype3: A
hostrecordcount: 3
defaulthostrecordown: True
UseOurDNS: True
DNS1: dns1.name-services.com
DNS2: dns2.name-services.com
DNS3: dns3.name-services.com
DNS4: dns4.name-services.com
DNS5: dns5.name-services.com
DNSCount: 5
AcceptTerms: True
URL:
ParentAccount: 000-00-0000
ParentLogin: Main-eNom
NoService: False
BulkRegLimit: 100
Account: 217-no-0647
RRPCode: 200
RRPText:
Command: GETEXTENDINFO
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.828
Done: true
RequestDateTime: 2/3/2015 6:20:41 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RegistrarHold=False
Expiration=6/10/2019
MaxExtension=5
MinAllowed=1
CCAuthorized=True
Price=8.95
Balance=-1,832.23
AvailableBalance=6,193.37
DefPeriod=1
AllowDNS=False
ShowPopups=False
AutoRenew=True
RegLock=False
AutoPakRenew=False
UseDNS=True
ResellerStatus=
RenewalSetting=1
RenewalBCC=0
RenewalURLForward=False
RenewalEmailForward=False
MailNumLimit=100
IDProtect=False
DefIDProtectRenew=False
DefWBLRenew=False
NameJetSales=False
hostname1=@
address1=85.92.87.178
recordtype1=A
hostname2=*
address2=85.92.87.179
recordtype2=A
hostname3=www
address3=85.92.87.180
recordtype3=A
hostrecordcount=3
defaulthostrecordown=True
UseOurDNS=True
DNS1=dns1.name-services.com
DNS2=dns2.name-services.com
DNS3=dns3.name-services.com
DNS4=dns4.name-services.com
DNS5=dns5.name-services.com
DNSCount=5
AcceptTerms=True
URL=
ParentAccount=000-00-0000
ParentLogin=Main-eNom
NoService=False
BulkRegLimit=100
Account=217-no-0647
RRPCode=200
RRPText=
Command=GETEXTENDINFO
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
ExecTime=0.109
Done=true
RequestDateTime=2/3/2015 6:21:13 PM
```
Related Commands
----------------

Extend

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

StatusDomain

ValidatePassword