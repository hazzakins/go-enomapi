CertGetCerts
============

Retrieve a list of SSL certificates

Usage
-----

Retrieve a list of the SSL certificates currently registered in this account.

Availability
------------

All resellers have access to this command.

Contraints
----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The certificate must belong to this account.

Input parameters
----------------

| Parameter      | Type | Status          | Description |
| ------------------- | -------- | ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command       | string | Required         | CertGetApproverEmail. |
| UID | string  | Required | Your Account ID.                                                                                                                                                                                                                                            |
| PW         | string | Required         | Your API Token. |
| ProdType | integer | Optional | Product identification number. - \*GeoTrust\*\*: - 20: QuickSSL Premium - 21: True BusinessID - 24: TrueBizID with EV - 26: QuickSSL - 27: TrueBizID Wildcard - 23: RapidSSL - 285: RapidSSL Wildcard - \*Symantec\*\*: - 180: Secure Site - 181: Secure Site Pro - 182: Secure Site EV - 183: Secure Site Pro EV - 286: Encryption Everywhere - \*Comodo\*\*: - 211: Essential - 212: Instant - 213: Premium Wildcard - 214: Essential Wildcard - 221: EV - 222: EV SGC - 270: UCC DV - 274: UCC OV  |
| StartPosition    | integer | Optional *default=1*   | Return results beginning with this position in the sorted list. Example, *StartPosition=26&PagingPageSize=25* returns accounts 26 through 50 in the sorted list. |
| PagingPageSize | integer | Optional *default=25* | Number of accounts to return in this response. Maximum permitted value is 250.                                                                                                                                                                                                             |
| SortBy       | string | Optional *default=CertID* | Sorting parameter. Permitted values: - CertID - Domain - CertStatus - Expiration |
| SortByDirection | string  | Optional *default=Asc* | Sort order. Permitted values: - Asc - Desc                                                                                                                                                                                                                               |
| CertStatusID    | integer | Optional         | Filter output by a specific cert status ID. Permitted values: - 0: Return all certs - 1: Awaiting Configuration - 2: Processing - 4: Certificate Issued - 6: Rejected by Customer - 7: Refunded - No Cert Issued - 8: Refunded - Cert Issued - 9: Approval Email Sent \(Comodo only\) - 10: Approved by Domain Owner \(Comodo only\) - 12: Pending Installation \(associate with our hosting\) - 13: Cert Installed \(associate with our hosting\) Note: CertStatusID=3, 5 and 11 are no longer valid |
| DomainNameContains | string  | Optional | Specific word in domain names.                                                                                                                                                                                                                                     |
| OrderDateStart   | datetime | Optional         | Order date start filter. Use format MM/DD/YYYY. |
| OrderDateEnd | datetime | Optional | Order date end filter. Use format MM/DD/YYYY.                                                                                                                                                                                                                             |
| ExpirationDateStart | datetime | Optional         | Expiration date start filter. Use format MM/DD/YYYY. |
| ExpirationDateEnd | datetime | Optional | Expiration date end filter. Use format MM/DD/YYYY.                                                                                                                                                                                                                           |
| ResponseType    | string | Optional         | Format of response. Permitted values: - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output | Type   | Description |
| ------------------ | -------- | ------------------------------------------------------------------------------------------------ |
| Command | string  | Name of command executed. |
| IsExpired | boolean | Is this cert expired? |
| CertID | integer | Identification number of this cert. |
| OrderDate | datetime | Date on which this cert was purchased. |
| ConfigDate | datetime | Date on which this cert was configured . |
| DomainName | string  | Domain that this cert is associated with. |
| ValidityPeriod | integer | Validity period in month. |
| RenewalIndicator | boolean | Renewal indicator setting. |
| ProdType | integer | Product ID number. |
| ProdDesc | string  | Text description of product. |
| ExpirationDate | datetime | Expiration date of this cert. |
| ServerCount | integer | Should always return 1. |
| CertStatus | string  | Current status of this cert. |
| CertStatusID | integer | Current status of this cert. |
| OrderID | integer | Identification number of the order that included this cert. |
| DomainCountAllowed | integer | Number of domain allowed for UCC cert. |
| ReissueStatus | boolean | Reissue status. |
| ErrorCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string  | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | string  | True value indicates this entire response has reached you successfully. |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.asp?command=CertGetCerts&uid=resellid&pw=resellpw&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertGetCerts>
 <Certs>
  <Cert>
   <IsExpired>False</IsExpired>
   <CertID>726340</CertID>
   <OrderDate>12/15/2016</OrderDate>
   <ConfigDate></ConfigDate>
   <DomainName><![CDATA[buditest.com]]></DomainName>
   <ValidityPeriod>12</ValidityPeriod>
   <RenewalIndicator>False</RenewalIndicator>
   <ProdType>23</ProdType>
   <ProdDesc>SSL Certificate - RapidSSL</ProdDesc>
   <ExpirationDate>12/15/2017</ExpirationDate>
   <ServerCount>1</ServerCount>
   <CertStatus>Cert Issued</CertStatus>
   <CertStatusID>7</CertStatusID>
   <OrderID>267562071</OrderID>
   <DomainCountAllowed></DomainCountAllowed>
   <ReissueStatus></ReissueStatus>
  </Cert>
  <Cert>
   <IsExpired>True</IsExpired>
   <CertID>435118</CertID>
   <OrderDate>12/9/2013</OrderDate>
   <ConfigDate></ConfigDate>
   <DomainName><![CDATA[buditest.com]]></DomainName>
   <ValidityPeriod>12</ValidityPeriod>
   <RenewalIndicator>False</RenewalIndicator>
   <ProdType>23</ProdType>
   <ProdDesc>SSL Certificate - RapidSSL</ProdDesc>
   <ExpirationDate>12/9/2014</ExpirationDate>
   <ServerCount>1</ServerCount>
   <CertStatus>Awaiting Configuration</CertStatus>
   <CertStatusID>1</CertStatusID>
   <OrderID>237772242</OrderID>
   <DomainCountAllowed></DomainCountAllowed>
   <ReissueStatus></ReissueStatus>
  </Cert>
  <ProductType>all</ProductType>
  <StatusID>all</StatusID>
  <StartPosition>1</StartPosition>
  <PagingPageSize>2</PagingPageSize>
  <Count>2</Count>
  <Sortby></Sortby>
  <SortByDirection>desc</SortByDirection>
  <TotalRecords>105</TotalRecords>
 </Certs>
</CertGetCerts>
<TotalCertCount>105</TotalCertCount>
<Command>CERTGETCERTS</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl0vwapi05</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.031</ExecTime>
<Done>true</Done>
<TrackingKey>cce2f424-8a40-465b-851a-845e19700391</TrackingKey>
<RequestDateTime>12/16/2016 1:11:00 PM</RequestDateTime>
```
Related Commands
----------------
- [CertChangeApproverEmail](../docs/cert-change-approveremail.md)
- [CertConfigureCert](../docs/cert-configure-cert.md)
- [CertGetApproverEmail](../docs/cert-get-approver-email.md)
- [CertGetCertDetail](../docs/cert-get-certdetail.md)
- [CertGetCerts](../docs/cert-get-certs.md)
- [CertModifyOrder](../docs/cert-modify-order.md)
- [CertParseCSR](../docs/cert-parse-csr.md)
- [CertPurchaseCert](../docs/cert-purchase-cert.md)
- [CertResendApproverEmail](../docs/cert-resend-approveremail.md)