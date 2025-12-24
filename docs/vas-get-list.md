VAS\GetList
===========

Retrieve a list of account or subscription

Usage
-----

Retrieve a list of the VAS \(Value Added Services\) account or subscription for a reseller.

Availability
------------

All resellers have access to this command.

Contraints
----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The VAS account or subscription must belong to this account.

Products
--------
- [Whois Publicity Service](#vas-get-list-whoispublicityservice)

Whois Publicity Service
-----------------------

> ### Enabled flag
>
>
>
> This is the flag that can be enabled or disabled by the domain owner to display public or private data. in order WPS to be enabled \(active state\), registrant must consent the GDPR requirement for WPS product. Please check the consent status using VAS\_GetDetail command: [Private Data Visibility Logic](../docs/vas-get-detail.md#vas-get-detail-whoispublicityservice).

#### Input Parameters

| Parameter    | Type | Status  | Description |
| --------------- | ------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command     | string | Required | VAS\_GetList. |
| UID | string | Required | Your Account ID.                                                                  |
| PW       | string | Required | Your API Token. |
| ProductType | string | Required | Type of VAS product. Permitted values: - WhoisPublicity                                               |
| RecordsToReturn | integer | Optional | *Not implemented yet.* The number of records being returned in the output. Default value: 25. |
| SortBy | string | Optional | *Not implemented yet.* Sorting parameter. Permitted values: - DomainName - ExpDate Default value: DomainName.                   |
| SortOrder    | string | Optional | *Not implemented yet.* The order in which the records are returned. Permitted values: - asc - ascending - desc - descending Default value: asc. |
| StartPosition | string | Optional | *Not implemented yet.* The starting point where the records are returned. Default value: 1.                            |
| ResponseType  | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML. |

#### Returned Parameters and Values

> Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output | Type   | Description |
| ------------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string  | Name of command executed. |
| ErrorCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string  | Error messages explaining the failure. These can be presented as is back to the client. |
| VASItemID | integer | ID number of the VAS Item. |
| DomainName | string  | Domain name. |
| DomainNameID | integer | Domain name ID. |
| Enabled | boolean | Whois Publicity flag. Expected values: - True - False Note: in order WPS to be enabled, registrant must consent the GDPR requirement for WPS product. |
| ExpDate | datetime | Value added service expiration date time. Format: MM/DD/YYYY HH:mm:SS AM/PM \(Pacific Time\) UTC and Epoch times are also available in the output. |

#### Example

```
https://resellertest.enom.com/interface.asp?command=VAS_GetList&uid=YourAccountID&pw=YourApiToken&ProductType={Required}&RecordsToReturn={Optional}&SortBy={Optional}&SortOrder={Optional}&StartPosition={Optional}&responsetype=xml&ProductType=WhoisPublicityService
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <VAS_GetList>
    <ProductType>WhoisPublicity</ProductType>
    <Account>
      <VASItemID>1017152</VASItemID>
      <DomainName>postman-reseller-20180816t172517394z.com</DomainName>
      <DomainNameID>366364261</DomainNameID>
      <ProdStatusID>2</ProdStatusID>
      <ProdStatusDesc>Service Active</ProdStatusDesc>
      <ProdEnabled>false</ProdEnabled>
      <ExpDate UTC="2019-08-17T00:25:00.000Z" Epoch="1565976300">8/16/2019 5:25:00 PM</ExpDate>
    </Account>
    <Account>
      <VASItemID>1017153</VASItemID>
      <DomainName>postman-purchaseplus-20180816t172803785z.com</DomainName>
      <DomainNameID>366364263</DomainNameID>
      <ProdStatusID>2</ProdStatusID>
      <ProdStatusDesc>Service Active</ProdStatusDesc>
      <ProdEnabled>false</ProdEnabled>
      <ExpDate UTC="2019-08-17T00:28:00.000Z" Epoch="1565976480">8/16/2019 5:28:00 PM</ExpDate>
    </Account>
    ... <cut> ...
    <Count>291</Count>
    <TotalRecords>291</TotalRecords>
  </VAS_GetList>
  <Command>VAS_GETLIST</Command>
  <APIType>API.NET</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>krkdt198</Server>
  <Site>eNom</Site>
  <IsLockable/>
  <IsRealTimeTLD/>
  <TimeDifference>+0.00</TimeDifference>
  <ExecTime>1.181</ExecTime>
  <Done>true</Done>
  <TrackingKey>dec065e7-1ffc-4a26-b094-405e21a47444</TrackingKey>
  <RequestDateTime>10/25/2018 3:45:34 PM</RequestDateTime>
```
Related Commands
----------------
- [VAS\_Verification](../docs/vas-verification.md)
- [VAS\_GetList](../docs/vas-get-list.md)
- [VAS\_GetDetail](../docs/vas-get-detail.md)
- [VAS\_Update](../docs/vas-update.md)
- [VAS\_Add](../docs/vas-add.md)
- [VAS\_Delete](../docs/vas-delete.md)
- [VAS\_Cancel](../docs/vas-cancel.md)
- [VAS\_GetPricing](../docs/vas-get-pricing.md)
- [Purchase](../docs/purchase.md)
- [AddToCart](../docs/add-to-cart.md)
- [PurchaseServices](../docs/purchaseservices.md)
- [GetHosts](../docs/gethosts.md)
- [SetHosts](../docs/sethosts.md)

**[Back to Top](#top)