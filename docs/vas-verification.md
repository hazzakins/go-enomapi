VAS\Verification
================

Verification action for an item

Usage
-----

Verification action for an item for an existing VAS \(Value Added Services\) account or subscription.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The VAS account or subscription must belong to this account.

Products
--------
- [Google Apps](#vas-verification-googleapps)

Google Apps
-----------

#### Domain Validation

***ActionType=Domain\_Validation*** runs both steps below and returns *IsValid* parameter, true or false.
- Internal - API needs to know if the selected domain is registered in the owner's account. In some cases, it might be located under sub-retail account. There is a function to check multiple scenarios such as if domain is registered in our system and registrar partners, allowed to be hosted, sub-account access validation or other business rules.
- External - API checks with Google to the domain is not being used in other accounts outside of our system.

#### Site Verification

Google verifies that you own your domain to ensure that no one else is using your domain for Google Apps without your permission \(to send email, for example\). These are the steps for completing site verification:
1. Get token value from Google -- ***ActionType=GetToken***
2. Update the domain's host record with the token
3. Send command to Google to verify the host record -- ***ActionType=SetToken***

> Extra parameter *ForceTrue=True* could be used in **ActionType=SetToken** to always return valid \(true\) result for testing purposes.

#### Input Parameters

| Parameter  | Type | Status  | Description |
| ------------ | ------ | -------- | ----------------------------------------------------------------------------------------------------- |
| Command   | string | Required | VAS\_Verification. |
| UID | string | Required | Your Account ID.                                           |
| PW      | string | Required | Your API Token. |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                      |
| TLD     | string | Required | Top-level domain name \(extension\) |
| ProductType | string | Required | Type of VAS product. Permitted values: - GoogleApps                         |
| ActionType  | string | Required | Action to take with the product. Permitted values: - Domain\_Validation - Get\_Token - Submit\_Token |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.              |

#### Returned Parameters and Values

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter  | Type | Description                                           |
| ----------- | ------- | ------------------------------------------------------------------------------------------------ |
| ProductType | string | Returns the product type as specified in your query.                      |
| ActionType | string | Returns the action type as specified in your query.                       |
| DomainName | string | Returns the DomainName that was specified in the query.                     |
| IsValid   | string | Returns true if verification is successful, false if verification fails.            |
| Message   | string | Returns a descriptive error message if IsValid is false.                    |
| ErrCount  | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX    | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done    | boolean | True indicates this entire response has reached you successfully.                |

#### Examples

```
https://resellertest.enom.com/interface.asp?command=VAS_Add&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType=GoogleApps&ActionType=Domain_Validation&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_Verification>
  <ProductType>GoogleApps</ProductType>
  <ActionType>Domain_Validation</ActionType>
  <DomainName>buditest1511a.net</DomainName>
  <IsValid>true</IsValid>
  <Message />
</VAS_Verification>
<Command>VAS_VERIFICATION</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl2vwapi01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.688</ExecTime>
<Done>true</Done>
<TrackingKey>c066226b-6e83-4182-a180-f6d53302e07b</TrackingKey>
<RequestDateTime>8/25/2016 10:40:35 AM</RequestDateTime>
```
```
https://resellertest.enom.com/interface.asp?command=VAS_Add&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType=GoogleApps&ActionType=Get_Token&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_Verification>
  <ProductType>GoogleApps</ProductType>
  <ActionType>Get_Token</ActionType>
  <Type>INET_DOMAIN</Type>
  <Method>DNS_TXT</Method>
  <Success>true</Success>
  <Message />
  <Token>google-site-verification=TrqtLVkvTuzPocHd0u08qpfZhF6e75v4MU4UMw253pg</Token>
</VAS_Verification>
<Command>VAS_VERIFICATION</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl2vwapi01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.703</ExecTime>
<Done>true</Done>
<TrackingKey>c049bb33-75c6-4ab0-8b21-31bfd76f163a</TrackingKey>
<RequestDateTime>8/25/2016 12:42:56 PM</RequestDateTime>
```
```
https://resellertest.enom.com/interface.asp?command=VAS_Add&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType=GoogleApps&ActionType=Submit_Token&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_Verification>
  <ProductType>GoogleApps</ProductType>
  <ActionType>Submit_Token</ActionType>
  <Type>INET_DOMAIN</Type>
  <Method>DNS_TXT</Method>
  <Success>true</Success>
  <Message />
  <Token></Token>
</VAS_Verification>
<Command>VAS_VERIFICATION</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl2vwapi01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.141</ExecTime>
<Done>true</Done>
<TrackingKey>2c69cd66-001c-4f47-a324-cc8bcfcbecfb</TrackingKey>
<RequestDateTime>8/25/2016 12:45:04 PM</RequestDateTime>
```
#### Related Commands
- [VAS\_Verification](../docs/vas-verification.md)
- [VAS\_GetList](../docs/vas-get-list.md)
- [VAS\_GetDetail](../docs/vas-get-detail.md)
- [VAS\_Update](../docs/vas-update.md)
- [VAS\_Add](../docs/vas-add.md)
- [VAS\_Delete](../docs/vas-delete.md)
- [VAS\_Cancel](../docs/vas-cancel.md)
- [VAS\_GetPricing](../docs/vas-get-pricing.md)
- [Purchase](../docs/domains/registration/purchase.md)
- [AddToCart](../docs/add-to-cart.md)
- [PurchaseServices](../docs/purchaseservices.md)
- [GetHosts](../docs/domains/domain-management/host-records/gethosts.md)
- [SetHosts](../docs/sethosts.md)

**[Back to Top](#top)