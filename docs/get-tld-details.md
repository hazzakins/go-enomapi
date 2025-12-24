GetTLDDetails
=============

Retrieve TLD characteristics in detail for a specified TLD.

Usage
-----

Use this command to retrieve TLD characteristics in detail for a specified TLD.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Multiple pages to drive specific characteristic based on TLD.

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
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| TLD       | Required | Top-level domain name \(extension\)         | 16 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter               | Description |
| -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| TLDID                    | Top Level Domain identification number |
| TLD | Top Level Domain name                                                            |
| AbleToLock                  | Support domain lock to prevent transfer? |
| AllowWPPS | Allow domain's contacts to use WHOIS Privacy Protection Service \(WPPS\)?                                  |
| AutoRenewOnly                | Allow user to renew domain manually at any time? |
| ExtAttributes | Require user to add extended attributes on domain registration?                                       |
| AllowWBL                   | Allow domain to use WHOIS Business Listing \(WBL\)? |
| TransRequiresNewContact | Require a new registrant contact for transfer                                                |
| EncodingType                 | Support Internationalized Domain Name \(IDN\)? If not empty, the TLD supports IDN |
| LangCode | IDN natural character encoding type                                                     |
| NativeDisplay                | Native characters for the IDN TLD |
| ClearPhoneNumber | Clear phone number as part of contact creation                                               |
| ClearFax                   | Clear fax number as part of contact creation |
| ValidateDNSHosting | Perform check to the Registry if this TLD can be hosted                                           |
| HasPremiumNames               | Support premium domain name category? |
| Registration - Realtime | Support realtime registration?                                                       |
| Registration - Unit             | Registration quantity unit |
| Registration - Minimum | Registration minimum unit                                                          |
| Registration - Maximum           | Registration maximum unit |
| Registration - HasPremiumNames | Support premium domain name registration?                                                  |
| Registration - ExtAttributes        | Require extended attributes for registration? |
| Registration - DomainLevel | Support domain level extended attributes?                                                  |
| Registration - DNSRequired         | Require DNS for registration? |
| Registration - DNSMinimum | Minimum number of DNS                                                            |
| Registration - DNSMaximum          | Maximum number of DNS |
| Registration - GeneralAvailabilityStartDate | Date and time the TLD enters General Availability                                              |
| Renewal - AutoRenewOnly           | Allow user to renew domain manually at any time? |
| Renewal - AutoRenewed | Does Registry renew the domain automatically on day of expiration?                                     |
| Renewal - RenewBeforeExpMonths       | How many months does the registry allow user to renew the domain prior to the expiration date? |
| Renewal - DeleteType | Deletion type. Legacy parameter                                                      |
| Renewal - DeleteDay             | Number of days from the expiration date, that we delete a domain name. If negative, we delete the domain name prior to the expiration date |
| Renewal - GracePeriod | How many days after registration/renewal can the domain be deleted and a full refund issued?                        |
| Renewal - Reactivate            | Can the domain be renewed after the expiration date? Flag used by order processing |
| Renewal - Restorable | Can the domain be renewed after the expiration date? Flag used by user interface                              |
| Renewal - RGP                | Can domain be restored after deletion. Flag used by auto renewals |
| Renewal - RGPDays | After domain is deleted, how long does the Registry allow for redemption period                               |
| Renewal - ExtendedRGP            | Support Extended RGP? |
| Renewal - TransferPeriod | The number of years the expiration date will be extended when a domain transfers.                             |
| Transfer - Transferable           | Support transfer? |
| Transfer - AuthInfo | Support Authorization Info or Extensible Provisioning Protocol \(EPP\) Key                                 |
| Transfer - Realtime             | Support realtime transfer request? |
| Transfer - AutoVerification | Support auto verification transfer?                                                     |
| Transfer - AutoFax             | Support auto fax transfer? |
| Transfer - TransferByFOA | Support transfer approval using an FOA email to the losing registrant                                    |
| Transfer - RequiresNewContact        | Require a new registrant contact for transfer |
| PremiumNames - HasPremiumNames | Support premium domain name registration?                                                  |
| PremiumNames - MaxPremiumRegYears      | Maximum number of years a premium domain can be registered |
| PremiumNames - TrademarkStartDate | The date when the registry will support TM claims check. If null, the Registry does not support TM claims check              |
| NameServers - RequiresDNS          | Require DNS for registration? |
| NameServers - MinNameServers | Minimum number of DNS                                                            |
| NameServers - MaxNameServers        | Maximum number of DNS |
| Command | Name of command executed                                                          |
| ErrCount                   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                          |
| Done                     | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves TLD characteristics in detail for a specified TLD, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
Command=GetTLDDetails&UID=resellid&PW=resellpw
&TLD=Ninja&Responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
Command=GetTLDDetails&UID=resellid&PW=resellpw
&TLD=Ninja&Responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
Command=GetTLDDetails&UID=resellid&PW=resellpw
&TLD=Ninja&Responsetype=text
```
The response is as follows:

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<tlds>
 <tld>
 <TLDID>298</TLDID>
 <TLD>ninja</TLD>
 <AbleToLock>True</AbleToLock>
 <AllowWPPS>True</AllowWPPS>
 <AutoRenewOnly>False</AutoRenewOnly>
 <ExtAttributes>False</ExtAttributes>
 <AllowWBL>True</AllowWBL>
 <TransRequiresNewContact>False</TransRequiresNewContact>
 <EncodingType>PUNY</EncodingType>
 <LangCode></LangCode>
 <NativeDisplay></NativeDisplay>
 <ClearPhoneNumber>False</ClearPhoneNumber>
 <ClearFax>False</ClearFax>
 <ValidateDNSHosting>False</ValidateDNSHosting>
 <HasPremiumNames>True</HasPremiumNames>
 <Registration>
  <Realtime>True</Realtime>
  <Unit>Year</Unit>
  <Minimum>1</Minimum>
  <Maximum>10</Maximum>
  <HasPremiumNames>True</HasPremiumNames>
  <ExtAttributes>False</ExtAttributes>
  <ExtAttributesDomainLevel>False</ExtAttributesDomainLevel>
  <DNSRequired>True</DNSRequired>
  <DNSMinimum>2</DNSMinimum>
  <DNSMaximum>13</DNSMaximum>
  <GeneralAvailabilityStartDate>2014-05-28T10:00:00</GeneralAvailabilityStartDate>
 </Registration>
 <Renewal>
  <AutoRenewOnly>False</AutoRenewOnly>
  <AutoRenewed>True</AutoRenewed>
  <RenewBeforeExpMonths>0</RenewBeforeExpMonths>
  <DeleteType>0</DeleteType>
  <DeleteDay>42</DeleteDay>
  <GracePeriod>5</GracePeriod>
  <Reactivate>True</Reactivate>
  <Restorable>True</Restorable>
  <RGP>True</RGP>
  <RGPDays>32</RGPDays>
  <ExtendedRGP>True</ExtendedRGP>
  <TransferPeriod>1</TransferPeriod>
 </Renewal>
 <Transfer>
  <Transferable>True</Transferable>
  <AuthInfo>True</AuthInfo>
  <Realtime>False</Realtime>
  <AutoVerification>True</AutoVerification>
  <AutoFax>True</AutoFax>
  <TransferByFOA>False</TransferByFOA>
  <RequiresNewContact>False</RequiresNewContact>
 </Transfer>
 <PremiumNames>
  <HasPremiumNames>True</HasPremiumNames>
  <MaxPremiumRegYears>True</MaxPremiumRegYears>
  <TrademarkStartDate>2014-04-30T10:00:00</TrademarkStartDate>
 </PremiumNames>
 <NameServers>
  <RequiresDNS>True</RequiresDNS>
  <MinNameServers>False</MinNameServers>
  <MaxNameServers>False</MaxNameServers>
 </NameServers>
 </tld>
</tlds>
<Command>GETTLDDETAILS</Command>
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
<ExecTime>0.016</ExecTime>
<Done>true</Done>
<TrackingKey>1bfc0cc0-69c4-446c-93a0-79bafe666ee3</TrackingKey>
<RequestDateTime>1/1/2014 2:42:35 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
TLDID: 298
TLD: ninja
AbleToLock: True
AllowWPPS: True
AutoRenewOnly: False
ExtAttributes: False
AllowWBL: True
TransRequiresNewContact: False
EncodingType: PUNY
LangCode:
NativeDisplay:
ClearPhoneNumber: False
ClearFax: False
ValidateDNSHosting: False
HasPremiumNames: True
SupportsDnsSec: True
Realtime: True
Unit: Year
Minimum: 1
Maximum: 10
HasPremiumNames: True
ExtAttributes: False
ExtAttributesDomainLevel: False
DNSRequired: True
DNSMinimum: 2
DNSMaximum: 13
GeneralAvailabilityStartDate: 2014-05-28T10:00:00
AutoRenewOnly: False
AutoRenewed: True
RenewBeforeExpMonths: 0
DeleteType: 0
DeleteDay: 42
GracePeriod: 5
Reactivate: True
Restorable: True
RGP: True
RGPDays: 32
ExtendedRGP: True
TransferPeriod: 1
Transferable: True
AuthInfo: True
Realtime: False
AutoVerification: True
AutoFax: True
TransferByFOA: False
RequiresNewContact: False
HasPremiumNames: True
MaxPremiumRegYears: True
TrademarkStartDate: 2014-04-30T10:00:00
TrademarkEndDate: 2014-07-29T10:00:00
RequiresDNS: True
MinNameServers: False
MaxNameServers: False
Command: GETTLDDETAILS
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +8.00
ExecTime: 0.047
Done: true
TrackingKey: 2ab9cede-9bab-4389-bcc7-11a87f8564da
RequestDateTime: 2/4/2015 12:50:27 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
TLDID=298
TLD=ninja
AbleToLock=True
AllowWPPS=True
AutoRenewOnly=False
ExtAttributes=False
AllowWBL=True
TransRequiresNewContact=False
EncodingType=PUNY
LangCode=
NativeDisplay=
ClearPhoneNumber=False
ClearFax=False
ValidateDNSHosting=False
HasPremiumNames=True
SupportsDnsSec=True
Realtime=True
Unit=Year
Minimum=1
Maximum=10
HasPremiumNames=True
ExtAttributes=False
ExtAttributesDomainLevel=False
DNSRequired=True
DNSMinimum=2
DNSMaximum=13
GeneralAvailabilityStartDate=2014-05-28T10:00:00
AutoRenewOnly=False
AutoRenewed=True
RenewBeforeExpMonths=0
DeleteType=0
DeleteDay=42
GracePeriod=5
Reactivate=True
Restorable=True
RGP=True
RGPDays=32
ExtendedRGP=True
TransferPeriod=1
Transferable=True
AuthInfo=True
Realtime=False
AutoVerification=True
AutoFax=True
TransferByFOA=False
RequiresNewContact=False
HasPremiumNames=True
MaxPremiumRegYears=True
TrademarkStartDate=2014-04-30T10:00:00
TrademarkEndDate=2014-07-29T10:00:00
RequiresDNS=True
MinNameServers=False
MaxNameServers=False
Command=GETTLDDETAILS
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.016
Done=true
TrackingKey=feb728c2-c188-4bd7-809d-aaae149f836b
RequestDateTime=2/4/2015 12:51:44 PM
```
Related Commands
----------------

GetTLDList