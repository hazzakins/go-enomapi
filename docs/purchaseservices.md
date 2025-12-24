PurchaseServices
================

Purchase Value Added Services product in real time

Usage
-----

Use this command to purchase one value added service in real time, as opposed to going through a shopping cart and our order queue.

Most commonly, real time commands are used by resellers who maintain their own databases.

Availability
------------

All resellers have access to this command.

Constraints
-----------
- The login ID and API token must be valid.
- Some products are available to only some TLDs. For example, ID Protect is not available for .us.
- To use our credit card processing, this must be an ETP reseller account that has signed a credit card agreement with us.

Input Parameters
----------------

**[General Parameters](#general-params)** : *Parameters that all variations of this command will accept.*

**[Credit Card Processing](#cc-proc)** : *Parameters to pass when our Credit Card Services will be used.*
* * *

For product specific input parameters, click the appropriate link below.
* * *

\*\*\[ID protect \*\* : *Specific to ID protect*

**[SSL Certificates](#ssl-certs)** : *Specific to SSL certificate orders*

**[POP Email](#pop-email)** : *Specific to POP Email orders*

**[DNS Hosting](#dns-hosting)** : *Specific to DNS Hosting purchases.*

**[Whois Publicity Service](#whois-publicity-service)**: *Specific to Whois Publicity service.*
* * *

Returned Parameters and Example Usage
-------------------------------------

**[Returned Parameters](#return)** - *Parameters and values returned by this command*

**[Example Output](#example-output)** - *Example usage of this Command*.
* * *

General Parameters
------------------

These parameters are used for any service purchased through PurchaseServices.

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------- |
| command     | string | Required | PurchaseServices |
| uid | string | Required | Your Account ID                               |
| pw       | string | Required | Your API Token |
| ResponseType | string | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.* |

**[Back to top](#top)**

Credit Card Processing
----------------------

These parameters are to be used when utilizing our credit card services.

> **When you pass credit card information with this command, you must use the secure HTTPS protocol.**
>
>
>
> *Our credit card services are only available for[ID Protect](#wpps-wbl), [Business Listing](#wpps-wbl) and [DNS hosting](#dns-hosting).*

| Input Parameter  | Type | Status            | Description |
| ------------------ | ---------- | ----------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| UseCreditCard   | string | Optional           | **The credit card supplied in this query string is charged only if UseCreditCard=yes.** If this param is omitted, or if the value supplied is anything other than yes, the account balance rather than the credit card is debited for this transaction. This is true even if the query string includes all the required credit card information. *Permitted values are yes and no.* |
| EndUserIP | string   | Required if UseCreditCard=yes | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format *NNN.NNN.NNN.NNN.*                                                                                                                           |
| CardType      | string | Required if UseCreditCard=yes | Credit card type. *Permitted values are Visa, Mastercard, AmEx, Discover* |
| CreditCardNumber | int    | Required if UseCreditCard=yes | Credit card number                                                                                                                                                                                   |
| CreditCardExpMonth | string | Required if UseCreditCard=yes | Expiration month of the credit card. *Format as MM* |
| CreditCardExpYear | int    | Required if UseCreditCard=yes | Expiration year of the credit card. *Format as YYYY*                                                                                                                                                                  |
| CVV2        | int | Required if UseCreditCard=yes | Credit card verification value. **Required if printed on card** |
| CCName | string   | Required if UseCreditCard=yes | Cardholder’s name                                                                                                                                                                                    |
| CCAddress     | string | Required if UseCreditCard=yes | Credit card billing street address |
| CCCity | string   | Required if UseCreditCard=yes | Credit card billing city                                                                                                                                                                                |
| CCStateProvince  | string | Required if UseCreditCard=yes | Credit card billing state or province |
| CCCountry | string   | Required if UseCreditCard=yes | Credit card billing country. *Permitted values are two-character ISO 3166-1 alpha-2 Country Codes*                                                                                                                                         |
| CCZip       | int, mixed | Required if UseCreditCard=yes | Credit card billing postal code. |
| CCPhone | mixed   | Required if UseCreditCard=yes | Credit card billing phone. *Format as:* *\+CountryCode.PhoneNumber, where CountryCode and PhoneNumber use only numeric characters and the \+ is URL-encoded as a plus sign \(%2B\).* ***For example:*** *%2B001.4252744500*                                                                            |

ID protect
----------

> You may use our Credit Card processing for these commands.
>
>
>
> See **[Credit Card Processing](#cc-proc)**.

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Service     | string | Required | For ID Protect, use *WPPS*. |
| SLD | string | Required | Second-level domain name - \**For example*\*\*: *'Enom' in Enom.com*                                                 |
| TLD       | string | Required | Top-level domain name \(extension\). |
| NumYears | int  | Optional | Number of years to subscribe. *Default value is 1.*                                                         |
| RenewName    | int | Optional | Use RenewName=1 to set this service to renew 30 days prior to its expiration. *Default value is 0.* |
| EmailNotify | int  | Optional | If included in your query, sends you an email confirmation when a customer makes a purchase using this command. - Permitted values are 0 and 1\_. **Default is 0.** |

SSL Certificates
----------------

| Input Parameter | Type | Status             | Description |
| --------------- | ------ | ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Service     | string | Required            | Type of SSL certificate. *Permitted values are:* - Certificate-Comodo-Essential - Certificate-Comodo-Instant - Certificate-Comodo-Essential-Wildcard - Certificate-Comodo-Premium-Wildcard - Certificate-Comodo-EV - Certificate-Comodo-EV-SGC - Certificate-GeoTrust-QuickSSL - Certificate-GeoTrust-QuickSSL-Premium - Certificate-GeoTrust-TrueBizID - Certificate-GeoTrust-TrueBizID-Wildcard - Certificate-GeoTrust-TrueBizID-EV - Certificate-RapidSSL - Certificate-RapidSSL-RapidSSL-Wildcard - Certificate-VeriSign-Secure-Site - Certificate-VeriSign-Secure-Site-Pro - Certificate-VeriSign-Secure-Site-EV - Certificate-VeriSign-Secure-Site-Pro-EV - Certificate-Comodo-UCC-DV - Certificate-Comodo-UCC-OV |
| NumYears | int  | Optional, default value is 1. | Number of years to purchase this SSL certificate. *Permitted values are**1** to the following maximum values:* \*The following certificate types allow a maximum of**2**: \* - Comodo Essential - Comodo Instant - Comodo Essential Wildcard - Comodo Premium Wildcard - GeoTrust QuickSSL - GeoTrust QuickSSL Premium - GeoTrust True BusinessID - GeoTrust True BusinessID Wildcard - RapidSSL Wildcard - VeriSign Secure Site - VeriSign Secure Site Pro - Comodo UCC DV - Comodo UCC OV - Comodo EV - Comodo EV SGC - GeoTrust True BusinessID with EV - VeriSign Secure Site with EV - VeriSign Secure Site Pro with EV                                                                         |

POP Email
---------

| Input Parameters | Type | Status  | Description |
| ---------------- | ------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| Service     | string | Required | Type of service to purchase. For POP Email, Service=MAILBOX or POP3 |
| EndUserIP | string | Required | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format *NNN.NNN.NNN.NNN.* |
| SLD       | string | Required | Second-level domain name - \**For example*\*\*: *'Enom' in Enom.com* |
| TLD | string | Required | Top-level domain name \(extension\).                                               |
| Quantity     | int | Optional | Number of email address\(es\) based on the quantity unit |
| QuantityUnit | string | Optional | Quantity unit. Permitted values are: - \*packs\*\*: pack of 10 email addresses \(default\) - \*boxes\*\*: single email address   |
| ActionType    | string | Optional | Possible values: - Create \(default\) - Renew \(ProductID is required for renewal\) |
| ProductID | int  | Optional | POP pak ID number. To retrieve this value, use the [GetPOP3](http://www.enom.com/api/API%20topics/api_GetPOP3.htm) command.    |
| TimeUnit     | string | Optional | Possible values: - Year \(default\) - Month |
| TimeQuantity | int  | Optional | Possible values: *1-12 number of months or years \(default is 1\)*                                 |
| Capacity     | int | Optional | Possible values: - 10 - 10 Gigabytes of storage - 25 - 25 Gigabytes of storage 50 - 50 Gigabytes of storage |
| Period | int  | Optional | Number of period/time \(year\)                                                   |

DNS Hosting
-----------

> You may use our Credit Card processing for this command.
>
>
>
> See **[Credit Card Processing](#cc-proc)**.

| Input Parameter | Type | Status         | Description |
| --------------- | ------ | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Service     | string | Required        | DNSHosting |
| NumYears | int  | Required | Number of years to subscribe to this service. Permitted values are *1 to 10.*                                                               |
| SLD       | string | Required        | Second-level domain name - \**For example*\*\*: *'Enom' in Enom.com* |
| TLD | string | Required | Top-level domain name \(extension\).                                                                                    |
| DomainPassword | string | Optional        | Set a password on the domain name |
| RenewName | int  | Optional | Set to 1 to auto-renew the service.                                                                                    |
| EmailNotify   | int | Optional; default is 0 | Permitted values are: - **0**: \_No email is sent when a customer makes a purchase using this command.\* - **1**: \_Sends you an email confirmation when a customer makes a purchase using this command.\* |

Whois Publicity Service
-----------------------

| Input Parameter | Type  | Status | Description                                             |
| --------------- | ------- | -------------------------------------- | --------------------------------------------------------------------------------------------------- |
| Service | string | Required | Type of service to add. *Permitted values are:* - WhoisPublicity                  |
| SLD       | string | Required if DomainNameID is not passed | Second-level domain name |
| TLD | string | Required if DomainNameID is not passed | Top-level domain name                                       |
| ActionType   | string | Optional, default value is Create.  | Type of product to add. *Permitted values are:* - Create \(default\) - Renew \(not available yet\) |
| Period | integer | Optional, default value is 1. | Number of period/time \(year\)                                   |

Returned Parameters and Values
------------------------------

> Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter  | Type | Description                                                                                                |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CCTransResult   | string | Credit card preauthorization result; included when using our credit card processing. See **[Credit Card Processing](#cc-proc)** parameters for services that allow use of our credit processing service. |
| SubscriptionID   | int | Our identification number for the Domain Name associated with a **[GoMobi](#gomobi)** purchase.                                                      |
| BundleID      | int | Our identification number for a **[POP](#pop-email)** email purchase.                                                                   |
| BundleCount    | int | Total Number of POP packs purchased with this query.                                                                           |
| PackageID     | int | - \*[Web hosting](#webhosting)\*\* package ID number assigned by us.                                                                   |
| ServerIP      | string | IP of the Web hosting server for a Web hosting purchase.                                                                         |
| Successful     | string | Success status of a query to purchase Web hosting service.                                                                        |
| CertID       | int | Cert ID for the cert purchased in this order; included for any **[SSL certificate](#ssl-certs)** purchase.                                                |
| OrderID      | int | Order ID number.                                                                                             |
| ProductType    | string | Product type                                                                                                |
| ActionType     | string | Action type                                                                                                |
| DomainName     | string | Domain name                                                                                                |
| OrderID      | int | Order ID                                                                                                  |
| Quantity      | int | Number of users                                                                                              |
| BillingPeriod   | string | Billing period. Possible values are monthly and yearly                                                                          |
| VasItemID     | int | ID number of the Google Apps account                                                                                    |
| OrderAmount    | decimal | Order amount                                                                                                |
| ExpDate      | datetime | Expiration date                                                                                              |
| TotalQuantity   | int | Total quantity                                                                                               |
| UpgradeToVASItemID | int | ID number of the Google Apps account to be updated or upgraded to.                                                                    |
| Command      | string | Name of command executed.                                                                                         |
| ErrCount      | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                      |
| ErrX        | string | Error messages explaining the failure. These can be presented as is back to the client.                                                         |
| Done        | string | True indicates this entire response has reached you successfully.                                                                     |

Example Output
--------------

The following query purchases new GoogleApps service.

```
https://resellertest.enom.com/interface.asp?Command=PurchaseServices&uid=resellid&pw=tester&service=googleapps&sld=example&tld=com&responsetype={text, XML, or HTML}
```
```
<interface-response>
<Success>True</Success>
<ProductType>Google Apps for Work - Monthly</ProductType>
<ActionType>Create</ActionType>
<DomainName>example.com</DomainName>
<OrderID>162046629</OrderID>
<Quantity>1</Quantity>
<BillingPeriod>Monthly</BillingPeriod>
<VasItemID>1008918</VasItemID>
<OrderAmount>4.50</OrderAmount>
<Command>PURCHASESERVICES</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL2VWAPI01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+08.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>dce159ef-61a8-4d2c-8c99-e00cafe13745</TrackingKey>
<RequestDateTime>7/14/2016 2:36:39 PM</RequestDateTime>
```
Related Commands
----------------
- [AddToCart](../docs/add-to-cart.md)