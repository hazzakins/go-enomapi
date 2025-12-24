VAS\GetPricing
==============

Get product pricing and pro-rated value

Usage
-----

Get VAS \(Value Added Services\) product pricing and pro-rated value

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
- [Google Apps](#vas-get-pricing-googleapps)

Google Apps
-----------

#### Input Parameters

| Parameter  | Type | Status              | Description |
| ------------ | ------ | -------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| Command   | string | Required             | VAS\_Add. |
| UID | string | Required | Your Account ID.                                                            |
| PW      | string | Required             | Your API Token. |
| ProductType | string | Required | Type of VAS product. Permitted values: - GoogleApps - GoogleApps-Freetrial - GoogleApps-Unlimited - GoogleApps-Unlimited-Freetrial |
| ActionType  | string | Required             | Action to take with the product. Permitted values: - Create - Add\_User |
| SLD | string | Required if ActionType=Add\_User | Second-level domain name \(e.g. "enom" in "enom.com"\).                                      |
| TLD     | | Required if ActionType=Add\_User | Top-level domain name \(extension\). |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.                               |

#### Returned Parameters and Values

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter   | Type | Description                                                                 |
| -------------- | ------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Command    | string | Name of command executed.                                                          |
| ErrorCount   | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                       |
| ErrX      | string | Error messages explaining the failure. These can be presented as is back to the client.                          |
| Done      | string | "True" value indicates this entire response has reached you successfully.                                  |
| ProductType  | string | Name of the product added. Expected values: - GoogleApps - GoogleApps-Freetrial - GoogleApps-Unlimited - GoogleApps-Unlimited-Freetrial |
| ActionType   | string | Name of the action taken. Expected values: - Create - Add\_User                                       |
| VASItemID   | integer | Value Added Services item ID.                                                        |
| OrderAmount  | decimal | The order amount for this action type                                                    |
| ProdTypeID   | integer | Enom's product type ID                                                            |
| ProdDesc    | string | Google's product description                                                         |
| Period     | integer | Periodic value                                                                |
| BillingPeriod | string | Type of billing period. Permitted values: - Month \(monthly\) - Year \(yearly\)                               |
| ExpirationDate | string | Expected expiration date value                                                        |
| TotalUsers   | integer | Number of total user after the action is taken                                                |
| TaxAmount   | decimal | Tax amount for this order                                                          |

#### Examples

```
https://resellertest.enom.com/interface.asp?command=VAS_GetPricingAdd&uid=YourAccountID&pw=YourApiToken&ProductType=GoogleApps&ActionType=Create&SLD={Optional}&TLD={Optional}&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_GetPricing>
 <ProductType>GoogleApps-Unlimited</ProductType>
 <ActionType>Create</ActionType>
 <VasItemID />
 <OrderAmount>9.00</OrderAmount>
 <ProdTypeID>2412</ProdTypeID>
 <ProdDesc>Google Apps for Work Unlimited - Monthly</ProdDesc>
 <Period>1</Period>
 <BillingPeriod>Month</BillingPeriod>
 <ExpDate>9/25/2016</ExpDate>
 <TotalUsers>1</TotalUsers>
 <TaxAmount>0</TaxAmount>
</VAS_GetPricing>
<Command>VAS_GETPRICING</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl2vwapi01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>2.132</ExecTime>
<Done>true</Done>
<TrackingKey>514817bc-cc71-4153-b0ff-d3518e1abfa7</TrackingKey>
<RequestDateTime>8/25/2016 10:06:27 PM</RequestDateTime>
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
- [Purchase](../docs/purchase.md)
- [AddToCart](../docs/add-to-cart.md)
- [PurchaseServices](../docs/purchaseservices.md)
- [GetHosts](../docs/gethosts.md)
- [SetHosts](../docs/sethosts.md)

**[Back to Top](#top)