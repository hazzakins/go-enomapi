VAS\Delete
==========

Delete an item

Usage
-----

Delete an item from an existing VAS \(Value Added Services\) account or subscription.

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
- [Google Apps](#vas-delete-googleapps)

Google Apps
-----------

> ### Seats and Users definition
>
>
>
> ***Seat***
>
> An empty placeholder or slot that has not been used or assigned to a user. You may purchase multiple seats without any users assigned to.
>
> API parameter \(output\): *QtyPurchased*
>
>
>
> ***User***
>
> An account \(user\) that is assigned to a seat. This user becomes an active license or active seat and has full access to Google's service.
>
> API parameter \(output\): *QtyUsed*

> ### Seats and Users action description
>
>
>
> ***Delete\_Seat***
>
> Remove a seat from an existing Google Apps' account.
>
>
>
> - Seat number \(QtyPurchased\) will be decreased by 1 \(one\)
>
>
>
> ***Delete\_User***
>
> Delete and un-assign a user from an existing Google Apps' account
>
>
>
> - User number \(QtyUsed\) will be decreased by 1 \(one\)
> - Seat number \(QtyPurchased\) will be decreased by 1 \(one\)
>
>
>
> ***Delete\_User\_Only***
>
> Delete a user from an existing Google Apps' account
>
>
>
> - User number \(QtyUsed\) will be decreased by 1 \(one\)

#### Input Parameters

| Parameter  | Type | Status  | Description |
| ------------ | ------ | -------- | ------------------------------------------------------------------------------------------------------ |
| Command   | string | Required | VAS\_Delete. |
| UID | string | Required | Your Account ID.                                           |
| PW      | string | Required | Your API Token. |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\).                      |
| TLD     | string | Required | Top-level domain name \(extension\). |
| ProductType | string | Required | Type of VAS product. Permitted values: - GoogleApps                          |
| ActionType  | string | Required | Action to take with the product. Permitted values: - Delete\_User - Delete\_User\_Only - Delete\_Seat |
| Username | string | Required | Username to delete **For Example:** **user**@*example.com*                      |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML. |

#### Returned Parameters and Values

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter | Type  | Description |
| ------------ | ------- | ------------------------------------------------------------------------------------------------ |
| ProductType | string | Returns the product type as specified in your query. |
| ActionType | string | Returns the action type as specified in your query. |
| EmailAddress | string | Email address that's been deleted. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | True indicates this entire response has reached you successfully. |

#### Examples

```
https://resellertest.enom.com/interface.asp?command=VAS_Delete&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType={Required}&ActionType=Delete_User&UserName={Required}&responsetype=xml
```
```
<interface-response>
<VAS_Delete>
  <ProductType>GoogleApps</ProductType>
  <ActionType>Delete</ActionType>
  <EmailAddress>[email protected]</EmailAddress>
</VAS_Delete>
<Command>VAS_DELETE</Command>
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
<ExecTime>4.516</ExecTime>
<Done>true</Done>
<TrackingKey>ec178de5-447c-47ab-8b86-5d2e47e1b47c</TrackingKey>
<RequestDateTime>7/14/2016 3:59:58 PM</RequestDateTime>
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