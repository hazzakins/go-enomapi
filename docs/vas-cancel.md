VAS\Cancel
==========

Cancel an item

Usage
-----

Cancel an item for an existing VAS \(Value Added Services\) account or subscription.

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
- [Google Apps](#vas-cancel-googleapps)

Google Apps
-----------

#### Input Parameters

| Parameter  | Type | Status  | Description |
| ------------ | ------ | -------- | --------------------------------------------------------------------------- |
| Command   | string | Required | VAS\_Cancel. |
| UID | string | Required | Your Account ID.                              |
| PW      | string | Required | Your API Token. |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\).        |
| TLD     | string | Required | Top-level domain name \(extension\). |
| ProductType | string | Required | Type of VAS product. Permitted values: - GoogleApps            |
| ActionType  | string | Required | Action to take with the product. Permitted values: - Cancel\_Subscription |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML. |

#### Returned Parameters and Values

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed.                                    |
| ErrorCount    | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | string | "True" value indicates this entire response has reached you successfully.            |
| ProductType   | string | Name of the product added. Expected values: - GoogleApps                    |
| ActionType    | string | Name of the action taken. Expected values: - Cancel\_Subscription                |
| Sucess      | boolean | Success flag.                                          |
| OrderAmount   | float | The amount charged for adding the user.                             |

#### Examples

```
https://resellertest.enom.com/interface.asp?command=VAS_Add&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType=GoogleApps&ActionType=Cancel_Subscription&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_Cancel>
  <ProductType>GoogleApps</ProductType>
  <ActionType>Cancel_Subscription</ActionType>
  <Success>true</Success>
</VAS_Cancel>
<Command>VAS_CANCEL</Command>
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
<ExecTime>2.734</ExecTime>
<Done>true</Done>
<TrackingKey>ee71a3d7-3f4e-4467-8e0b-9ce2578619b0</TrackingKey>
<RequestDateTime>8/24/2016 4:32:39 PM</RequestDateTime>
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