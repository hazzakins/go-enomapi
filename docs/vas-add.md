VAS\Add
=======

Add an item

Usage
-----

Add an item \(real time purchase\) to an existing VAS \(Value Added Services\) account or subscription.

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
- [Google Apps](#vas-add-googleapps)

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
> ***Add\_Seat\_Only***
>
> Purchase a seat to an existing Google Apps' account.
>
>
>
> - Seat number \(*QtyPurchased*\) will be increased by 1 \(one\)
> - Billing event will executed
>
>
>
> ***Add\_User\_Only***
>
> Add or create a new user and assign it to an empty seat
>
>
>
> - User number \(*QtyUser*\) will be increased by 1 \(one\)
>
>
>
> ***Add\_User***
>
> Process both actions above.

#### Input Parameters

| Parameter   | Type | Status  | Description |
| -------------- | ------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command    | string | Required | VAS\_Add. |
| UID | string | Required | Your Account ID.                                                                                                                                                                                                             |
| PW       | string | Required | Your API Token. |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\).                                                                                                                                                                                        |
| TLD      | string | Required | Top-level domain name \(extension\). |
| ProductType | string | Required | Type of VAS product. Permitted values: - GoogleApps                                                                                                                                                                                            |
| ActionType   | string | Required | Action to take with the product. Permitted values: - Add\_User - Add\_Seat\_Only - Add\_User\_Only |
| FirstName | string | Required | - \*First Name\*\* of the added user.                                                                                                                                                                                                   |
| LastName    | string | Required | - \*Last Name\*\* of the added user. |
| UserName | string | Required | - \*User Name\*\* for the added user. - Usernames can contain letters \(a-z\), numbers \(0-9\), dashes \(-\), underscores \(\_\), apostrophes \('\), and periods \(.\). - Usernames can't contain an ampersand \(&\), equal sign \(=\), brackets \(<,>\), plus sign \(\+\), comma \(,\), or more than one period \(.\) in a row. - Usernames can begin or end with non-alphanumeric characters, with a maximum of 64 characters. |
| Password    | string | Required | - \*Password\*\* for the added user. - Passwords can contain any combination of ASCII characters and must contain a minimum of 8 characters. |
| AutoRetry | boolean | Optional | If set to true, API will automatically retry the failed request. Default value is false.                                                                                                                                                                         |
| AutoRetryCount | integer | Optional | Number of auto-retry attempts. Paused time value between trials is exponential number. Default value is 5. Maximum value is 10. |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.                                                                                                                                                                                |

#### Returned Parameters and Values

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter   | Type | Description                                           |
| -------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command    | string | Name of command executed.                                    |
| ErrorCount   | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX      | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done      | string | "True" value indicates this entire response has reached you successfully.            |
| ProductType  | string | Name of the product added. Expected values: - GoogleApps                    |
| ActionType   | string | Name of the action taken. Expected values: - Add\_User - Add\_Seat\_Only - Add\_User\_Only   |
| EmailAddress  | string | Email address associated with the value added service.                     |
| OrderID    | integer | Order number associated with the transaction.                          |
| OrderAmount  | float | The amount charged for adding the user.                             |
| AutoRetryCount | integer | Number of auto-retry attempts.                                 |

#### Examples

```
https://resellertest.enom.com/interface.asp?command=VAS_Add&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType=GoogleApps&ActionType=Add_User&FirstName={Required}&LastName={Required}&UserName={Required}&Password={Required}&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_Add>
  <ProductType>GoogleApps</ProductType>
  <ActionType>Add_User</ActionType>
  <EmailAddress>[email protected]</EmailAddress>
  <OrderID>162046659</OrderID>
  <OrderAmount>5.65</OrderAmount>
  <AutoRetryCount>1</AutoRetryCount>
</VAS_Add>
<Command>VAS_ADD</Command>
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
<ExecTime>4.641</ExecTime>
<Done>true</Done>
<TrackingKey>4dbcb0f3-bb9e-411c-851a-3a5ddc525a3b</TrackingKey>
<RequestDateTime>7/14/2016 3:51:57 PM</RequestDateTime>
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