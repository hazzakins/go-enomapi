Check
=====

Check the availability of a domain name.

Usage
-----

Use this command to check whether a domain name is already registered.

> ### Version 2 Overview
>
>
>
> There are few new parameters introduced in this version to allow customers obtaining more comprehensive domain information delivered in one single action \(this command\), instead of running multiple commands to accomplish the same results. Additionally, the XML output structures are improved to handle multiple domain objects.
>
> The new components are:
>
>
>
> - Returns **General Availability prices** for domain registration, renewal, restore and transfer on demand.
> - Returns **TLD properties** for the requested domain on demand.
> - Returns **EAP** detailed information and prices on demand.

> ### Customer Supplied Price
>
>
>
> The domain registry expects the customers to provide the *CustomerSuppliedPrice* to be sent to [AddToCart](../docs/addtocart.md) or **Purchase** command to confirm agreement with the purchase price. This command returns *ExpectedCustomerSuppliedPrice*, calculated based on the required price components.
>
>
>
> Here is the formula to calculate the price:
>
>
>
> ##### Premium in GA \(General Availability\)
>
>
>
> **CustomerSuppliedPrice = premium domain price for 1-year registration**
>
> Multiple year domain registrations are not allowed.
>
>
>
> #### EAP \(Early Access Program\)
>
>
>
> **CustomerSuppliedPrice =*n*-year domain registration \+ EAP participation fee**
>
> By default, it is 1-year domain registration.
>
>
>
> #### EAP and Premium
>
>
>
> **CustomerSuppliedPrice = premium domain price for 1-year registration \+ EAP participation fee**
>
> Multiple year domain registrations are not allowed.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- Second-level domain \(SLD\) must be composed of the letters a through z, the numbers 0 through 9, and the hyphen \(-\) character.
- The SLD must not begin or end with the hyphen character.
- The SLD must not contain spaces.
- The SLD must not contain special characters other than the hyphen character.
- The third and fourth characters of the SLD must not both be hyphens unless it is an encoded international character domain name.
- The SLD must contain 2 to 63 characters, inclusive.
- SLDs are not case sensitive.
- The SLD-TLD combination must be unique.

Input Parameters
----------------

| Parameter  | Type | Status   | Description |
| ------------ | ------- | ---------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| Command   | string | Required  | Check. |
| UID | string | Required | Your Account ID.                                                         |
| PW      | string | Required  | Your API Token. |
| SLD | string | Required\* | Second-level domain name \(e.g. "enom" in "enom.com"\) Note: required unless *DomainList* is used               |
| TLD     | string | Required\* | Top-level domain name \(extension\). Note: required unless *DomainList* or *TLDList* is used |
| Version | integer | Optional | Command's version. Default value is 1. Please see [Version 2](#check-version-2) note below.                   |
| DomainList  | string | Optional  | *\[Version 1 only\]* A comma- or newline-delimited list of domains to check, up to 30 names. If used, SLD and TLD are ignored. |
| TLDList | string | Optional | *\[Version 1 only\]* A comma- or newline-delimited list of TLDs to check, up to 30 names. If used, TLD is ignored.       |
| ResponseType | string | Optional  | Format of response. Permitted values: - Text \(default\) - HTML - XML |

Addition parameters for Version 2
---------------------------------

| Parameter | Type  | Status | Description                                                                             |
| ----------------- | ------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Version | integer | Optional | Permitted values are: - 1: original check command \(default\) - 2: version 2. Please see [Version 2 Overview](#check-version-2-overview) for detailed information. |
| IncludePrice   | boolean | Optional | Returns domain registration, renewal, restore and transfer prices. Permitted values: - 0 or False \(default\) - 1 or True |
| IncludeProperties | boolean | Optional | Returns TLD properties for the requested domain. Permitted values: - 0 or False \(default\) - 1 or True                               |
| IncludeEAP    | boolean | Optional | Returns EAP pricing and date. Permitted values: - 0 or False \(default\) - 1 or True |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

Version 1 \(default\)
---------------------

| Parameter | Type  | Description |
| -------------------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| DomainName | string | If return includes multiple domains, the name of domain *x*. Indexed *x* in HTML and Text output, but not in XML output. |
| RRPCodeX | integer | Success code. - 210: domain name is available at the Registry. - 211 domain is not available at the Registry. Please see [RRP Return Codes](../docs/rrp-return-codes.md) for more information. |
| RRPTextX | string | Text which accompanies and describes the RRPCode. |
| IsPremiumName | boolean | True indicates this is a premium name. |
| PremiumPrice | decimal | Registration price per year for this premium name in $US. |
| PremiumAboveThresholdPrice | boolean | True indicates the registration price for this premium name is higher than for non-premium names. |
| DomainCount | integer | The number of names checked. |
| ErrCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | True value indicates this entire response has reached you successfully. |

Version 2
---------

| Parameter | Type   | Description |
| ---------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name | string  | Domain name in request. |
| RRPCode | integer | Success code. - 210: domain name is available at the Registry. - 211 domain is not available at the Registry. Please see [RRP Return Codes](../docs/rrp-return-codes.md) for more information. |
| RRPText | string  | Text which accompanies and describes the RRPCode. |
| IsPremium | boolean | Is the requested domain a premium domain? |
| IsPlatinum | boolean | Is the requested domain a platinum domain? |
| IsEAP | boolean | Is the requested domain in EAP category? |
| `<Prices>` | object  | Price object for the requested domain. |
| `<Prices>` Currency | decimal | Price currency. Default value is US dollars. |
| `<Prices>` Registration | decimal | Registration price. |
| `<Prices>` Renewal | decimal | Renewal price. |
| `<Prices>` Restore | decimal | Restore \(redemption grace period or RGP\) price. |
| `<Prices>` Transfer | decimal | Transfer price. |
| `<Prices>` ExpectedCustomerSuppliedPrice | decimal | The exact price for this domain to acknowledge this domain comes at a price higher than most registrations. This price is required to be passed to [AddToCart](../docs/addtocart.md) or **Purchase** command if the domain is premium, platinum or EAP. |
| `<Properties>` | object  | TLD properties object for the requested domain. |
| `<Properties>` NativeSLD | string  | SLD in the native language if the domain is International Domain Name \(IDN\). |
| `<Properties>` MinRegYear | integer | Minimum registration year. |
| `<Properties>` MaxRegYear | integer | Maximum registration year. |
| `<Properties>` AbleToLock | boolean | Support domain lock to prevent transfer? |
| `<Properties>` ExtAttributes | boolean | Require user to add extended attributes on domain registration? |
| `<Properties>` Transferable | boolean | Support transfer? |
| `<Properties>` AllowWPPS | boolean | Allow domain's contacts to use WHOIS Privacy Protection Service \(WPPS\)? |
| `<Properties>` TrademarkStart | datetime | Start date of the TLD Trademark period \(if applicable\) in Pacific Time Zone, UTC and Epoch format. |
| `<Properties>` TrademarkEnd | datetime | End date of the TLD Trademark period \(if applicable\) in Pacific Time Zone, UTC and Epoch format. |
| `<EAP>` | object  | EAP object for the requested domain. |
| `<EAP>` CurrentDay | integer | Current EAP day. Possible values are 1 to 7. |
| `<EAP>` Day | integer | *\[Day specific information\]* EAP day. |
| `<EAP>` Fee | decimal | *\[Day specific information\]* EAP fee. |
| `<EAP>` Open | boolean | *\[Day specific information\]* Is it open for registration? |
| `<EAP>` DateStart | datetime | *\[Day specific information\]* Start date in Pacific Time Zone, UTC and Epoch format. |
| `<EAP>` DateEnd | datetime | *\[Day specific information\]* End date in Pacific Time Zone, UTC and Epoch format. |
| ErrCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| Err*X* | string  | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | True value indicates this entire response has reached you successfully. |

Example Input / Output
----------------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?command=check&sld=SuperDuperDomainName&tld=com&uid=resellid&pw=resellpw&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
	<DomainName>superduperdomainname.com</DomainName>
	<RRPCode>211</RRPCode>
 	<RRPText>Domain not available</RRPText>
 	<Command>CHECK</Command>
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
 	<ExecTime>1.734</ExecTime>
 	<Done>true</Done>
 	<TrackingKey>32f969e1-1475-4e5f-a42a-502c3f2d6599</TrackingKey>
 	<RequestDateTime>3/22/2017 3:58:23 PM</RequestDateTime>
 	<debug/>
</interface-response>
```
```
http://resellertest.enom.com/interface.asp?UID=resellid&PW=resellpw&SLD=all&TLD=email&Command=check&responsetype=xml&version=2&includeprice=1&includeproperties=1&includeeap=1
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
 <Domains>
  <Domain>
   <Name>all.email</Name>
   <RRPCode>210</RRPCode>
   <RRPText>Domain available</RRPText>
   <IsPremium>True</IsPremium>
   <IsPlatinum>False</IsPlatinum>
   <IsEAP>False</IsEAP>
   <Prices>
    <Currency />
    <Registration>3300.00</Registration>
    <Renewal>3300.00</Renewal>
    <Restore>250.00</Restore>
    <Transfer>3300.00</Transfer>
    <ExpectedCustomerSuppliedPrice>3300.00</ExpectedCustomerSuppliedPrice>
   </Prices>
   <Properties>
    <NativeSLD />
    <MinRegYear>1</MinRegYear>
    <MaxRegYear>10</MaxRegYear>
    <AbleToLock>True</AbleToLock>
    <ExtAttributes>False</ExtAttributes>
    <Transferable>True</Transferable>
    <AllowWPPS>True</AllowWPPS>
    <TrademarkStart UTC="2014-03-19T16:00:00Z" Epoch="1395219600">
     3/19/2014 9:00:00 AM
    </TrademarkStart>
    <TrademarkEnd UTC="2014-06-17T16:00:00Z" Epoch="1402995600">
     6/17/2014 9:00:00 AM
    </TrademarkEnd>
   </Properties>
   <EAP>
    <CurrentDay />
    <Pricing>
     <Day>1</Day>
     <Fee>10000.00</Fee>
     <Open>False</Open>
     <DateStart UTC="2014-03-19T16:00:00Z" Epoch="1395219600">
      3/19/2014 9:00:00 AM
     </DateStart>
     <DateEnd UTC="2014-03-20T15:59:59Z" Epoch="1395305999">
      3/20/2014 8:59:59 AM
     </DateEnd>
    </Pricing>
    ...
    <Pricing>
     <Day>7</Day>
     <Fee>100.00</Fee>
     <Open>False</Open>
     <DateStart UTC="2014-03-25T16:00:00Z" Epoch="1395738000">
      3/25/2014 9:00:00 AM
     </DateStart>
     <DateEnd UTC="2014-03-26T15:59:59Z" Epoch="1395824399">
      /26/2014 8:59:59 AM
     </DateEnd>
    </Pricing>
   </EAP>
  </Domain>
 </Domains>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t1</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>1.578</ExecTime>
<Done>true</Done>
<TrackingKey>f36c5444-ef5a-49e8-bfa4-0d43cbe2d879</TrackingKey>
<RequestDateTime>3/22/2017 4:16:34 PM</RequestDateTime>
<debug/>
</interface-response>
```
Premium
-------

```
https://resellertest.enom.com/interface.asp?command=check&sld=1plus1&tld=tv&uid=resellid&pw=resellpw&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
	<DomainName>1plus1.tv</DomainName>
	<IsPremiumName>true</IsPremiumName>
	<PremiumPrice>999.00</PremiumPrice>
	<RegistrationFee>999.00</RegistrationFee>
	<RenewalPrice>110.00</RenewalPrice>
	<RestorePrice>160.00</RestorePrice>
	<TransferPrice>39.95</TransferPrice>
	<RRPCode>210</RRPCode>
	<RRPText>Domain available</RRPText>
	<Command>CHECK</Command>
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
 	<ExecTime>1.134</ExecTime>
 	<Done>true</Done>
 	<TrackingKey>332369e1-1a75-4e5f-122a-502c32345599</TrackingKey>
 	<RequestDateTime>3/22/2017 2:21:20 PM</RequestDateTime>
 	<debug/>
</interface-response>
```
EAP
---

```
https://resellertest.enom.com/interface.asp?command=check&sld=1&tld=agency&uid=resellid&pw=resellpw&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
	<DomainName>1.agency</DomainName>
	<IsEAP>true</IsEAP>
	<IsPremiumName>false</IsPremiumName>
	<PremiumPrice>11000.00</PremiumPrice>
	<ExpectedCustomerSuppliedPrice>11000.00</ExpectedCustomerSuppliedPrice>
	<RegistrationFee>18.00</RegistrationFee>
	<EAPFee>11000.00</EAPFee>
	<RenewalPrice>18.00</RenewalPrice>
	<RestorePrice>250.00</RestorePrice>
	<TransferPrice>18.00</TransferPrice>
	<RRPCode>210</RRPCode>
	<RRPText>Domain available</RRPText>
	<Command>CHECK</Command>
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
 	<ExecTime>1.134</ExecTime>
 	<Done>true</Done>
 	<TrackingKey>aaa362e1-1335-123f-122a-502ab1245239</TrackingKey>
 	<RequestDateTime>3/22/2017 2:30:00 PM</RequestDateTime>
 	<debug/>
</interface-response>
```
EAP and Premium
---------------

```
https://resellertest.enom.com/interface.asp?command=check&sld=action&tld=agency&uid=resellid&pw=resellpw&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
 <DomainName>action.agency</DomainName>
 <IsEAP>true</IsEAP>
 <IsPremiumName>true</IsPremiumName>
 <PremiumPrice>11022.00</PremiumPrice>
 <ExpectedCustomerSuppliedPrice>11022.00</ExpectedCustomerSuppliedPrice>
 <RegistrationFee>22.00</RegistrationFee>
 <EAPFee>11000.00</EAPFee>
 <RenewalPrice>22.00</RenewalPrice>
 <RestorePrice>250.00</RestorePrice>
 <TransferPrice>22.00</TransferPrice>
 <RRPCode>210</RRPCode>
 <RRPText>Domain available</RRPText>
 <Command>CHECK</Command>
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
 <ExecTime>1.134</ExecTime>
 <Done>true</Done>
 <TrackingKey>123462aa-112s-7843-12aa-5234b12ab211</TrackingKey>
 <RequestDateTime>3/22/2017 3:55:12 PM</RequestDateTime>
 <debug/>
</interface-response>
```
Related Commands
----------------

[AddBulkDomains](../docs/AddBulkDomains.md)

[InsertNewOrder](../docs/InsertNewOrder.md)

[NameSpinner](../docs/namespinner.md)

[Purchase](../docs/purchase.md)