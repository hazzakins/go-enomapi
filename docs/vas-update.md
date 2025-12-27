VAS\Update
==========

Update an item

Usage
-----

Update an item from an existing VAS \(Value Added Services\) account or subscription.

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
- [Whois Publicity Service](#vas-update-whoispublicityservice)

Whois Publicity Service
-----------------------

> ### ENABLE definition
>
>
>
> Whois Publicity is considered in **enable** or **active** state if these 2 \(two\) conditions are met:
>
>
>
> 1. Registrant has consented the GDPR for Whois Publicity Service \(WPS\).
> 2. The domain WPS flag is set to Enable.
>
>
>
> Let's consider this scenario:
>
> Domain owner sends update command to enable WPS and gets response back as successful. The success flag means that our system receives the requested action and adds the action to the internal queue to be processed.
>
>
>
> You might need to run VAS\_GetDetail to make sure the back-end processes have been completed and the expected result is delivered as requested. There is a chance that the GDPR WPS has not been consented by the registrant. In this case, our system would not allow the WPS for this domain to be enabled \(active\).

#### Input Parameters

| Parameter  | Type | Status  | Description |
| ------------ | ------- | -------- | --------------------------------------------------------------------------------------------------------------- |
| Command   | string | Required | VAS\_Update. |
| UID | string | Required | Your Account ID.                                                |
| PW      | string | Required | Your API Token. |
| ProductType | string | Required | Type of VAS product. Permitted values: - WhoisPublicity                            |
| VASItemID  | integer | Required | ID number of the VAS Item. Use the [VAS\_GetList](../docs/vas-get-list.md) command to retrieve the ID number |
| ActionType | boolean | Required | Name of the action taken. Expected value: - Enable - Disable                          |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML. |

#### Returned Parameters and Values

> Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output | Type  | Description |
| ----------- | ------- | -------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed. |
| ErrorCount | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| ProductType | string | Name of the product. Expected values: - WhoisPublicity |
| ActionType | boolean | Name of the action taken. Expected value: - Enable - Disable |
| VASItemID | integer | ID number of the VAS Item. |
| Success | boolean | Update request status transmission. **Important: please see information box above about ENABLE definition.** |

#### Example

```
https://resellertest.enom.com/interface.asp?command=VAS_Update&uid=YourAccountID&pw=YourApiToken&VASItemID={Required}&ProductType=WhoisPublicityService&responsetype=xml&ActionType=Enable
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <VAS_Update>
    <ProductType>WhoisPublicityService</ProductType>
    <ActionType>Enable</ActionType>
    <VASItemID>1017154</VASItemID>
    <Success>true</Success>
  </VAS_Update>
  <Command>VAS_UPDATE</Command>
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
  <ExecTime>0.515</ExecTime>
  <Done>true</Done>
  <TrackingKey>e7806d66-5e79-45ab-bcd9-8e3e3c122558</TrackingKey>
  <RequestDateTime>10/17/2018 11:39:08 AM</RequestDateTime>
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
- [Purchase](../docs/domains/registration/purchase.md)
- [AddToCart](../docs/add-to-cart.md)
- [PurchaseServices](../docs/purchaseservices.md)
- [GetHosts](../docs/domains/domain-management/host-records/gethosts.md)
- [SetHosts](../docs/sethosts.md)

**[Back to Top](#top)