AddToCart
=========

Add an item to the shopping cart.

Usage
-----

To purchase items, once they are in the shopping cart, use the "InsertNewOrder" command.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- If the item is a domain name, use the "Check" command to confirm that the name is available.
- The domain names must be valid \(see the domain-name constraints under "Check"\).
- The domain names must use a top-level domain supported by this registrar.
- The number of SLDs must match the number of TLDs.
- There is a limit of 200 items \(active plus inactive\) in the cart.

Input Parameters
----------------

Click the links below to go directly to specific products:

**[Input parameters for Domains](#add-to-cart-domains)**

**[Input parameters for ID Protect](#add-to-cart-protect)**

**[Input parameters for SSL certificates](#add-to-cart-certificates)**

**[Input parameters for POP Email](#add-to-cart-popemail)**

**[Input parameters for Email and URL Forwarding](#add-to-cart-euforward)**

**[Input parameters for Whois Publicity Service](#add-to-cart-whois-publicity-service)**

Input parameters for Domains
----------------------------

```
https://resellertest.enom.com/interface.asp?command=AddToCart&uid=YourAccountID&pw=YourApiToken&EndUserIP={Required}&ProductType={Required}&SLD={Required}&TLD={Required}&Quantity={Optional}&AutoRenew={Optional}&RegLock={Opitonal}&UsePerm={Required for Extended RGP}&ClearItems={Optional}&responsetype={xml, text, or HTML}
```
| Input    | Type | Status                                | Description |
| ------------ | ------ | -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command   | string | Required                               | AddToCart |
| uid | string | Required | Your Account ID                                                                                                                        |
| pw      | string | Required                               | Your API Token |
| EndUserIP | string | Required | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN.                                                               |
| ProductType | string | Required                               | Type of product to add. Permitted values are: - Register - Renew \(Use Renew for all TLDs including .name bundles\) - RGP \(Redemption Grace Period\) - Extended RGP - Host \(DNS hosting\) - Host Renew - PremiumDomain |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                  |
| TLD     | string | Required                               | Top-level domain name \(extension\) |
| Quantity | int  | Optional for: - Register - Renew - Host - Host Renew - PremiumDomain | Permitted values are 1 through 10 for most domains; some TLDs vary. Registration cannot extend beyond 10 years from today for any TLD. Permitted value for PremiumDomain is 1. Default is duration you can set in "UpdateCusPreferences" Number of years. |
| AutoRenew  | int | Optional for: - Register                       | Do you want the name to attempt to renew automatically? \(renewal will succeed if available account balance is sufficient\)? Permitted values are 0 \(No\) and 1 \(Yes\) |
| RegLock | int  | Optional for: - Register | Do you want protection against unauthorized transfer of this domain name? Permitted values are 0 \(No\) and 1 \(Yes\)                                                                     |
| UsePerm   | string | Required for: - Extended RGP                     | Required value is "UsePerm=ok" |
| ClearItems | string | Optional | Make all items currently in the cart inactive \(but keep them in the cart\) and add a new item. Permitted value is "yes". To reactivate items in the cart that are inactive, use the "UpdateCart" command.                         |
| ResponseType | string | Optional                               | Format of response. Permitted values are "Text" \(default\), "HTML", or "XML". |

**[Back to Top](#top)**

Input parameters for ID Protect
-------------------------------

Use the input parameters in the table below to add ID Protect to the cart.
Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=AddToCart&uid=YourAccountID&pw=YourApiToken&EndUserIP={Required}&ProductType={Required}&SLD={Required}&TLD={Required}&Quantity={Optional}&ClearItems={Optional}&responsetype={xml, text, or HTML}
```
| Input | Type  | Status | Description                                                                                                 |
| ------------ | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command | string | Required | AddToCart                                                                                                  |
| UID     | string | Required | Your Account ID |
| PW | string | Required | Your API Token                                                                                                |
| EndUserIP  | string | Required | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| ProductType | string | Required | Type of product to add. Permitted values are: - IDProtect \(Whois Privacy Protection\) - IDProtectRenewal                                                  |
| SLD     | string | Required | Second-level domain name \(for example, "enom" in "enom.com"\) |
| TLD | string | Required | Top-level domain name \(extension\)                                                                                     |
| Quantity   | int | Optional | Number of years. Default value is: 1 |
| ClearItems | string | Optional | Make all items currently in the cart inactive \(but keep them in the cart\), and add a new item. Permitted value is yes. To reactivate items in the cart that are inactive, use the "UpdateCart" command. |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Input parameters for SSL certificates
-------------------------------------

Use the input parameters in the table below to add an SSL certificate to the cart.
Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=AddToCart&uid=YourAccountID&pw=YourApiToken&EndUserIP={Required}&ProductType={Required}&Quantity={Optional}&ClearItems={Optional}&responsetype={xml, text, or HTML}
```
| Input | Type  | Status | Description                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------ | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command | string | Required | AddToCart                                                                                                                                                                                                                                                                                                                                                                                       |
| UID     | string | Required | Your Account ID |
| PW | string | Required | Your API Token                                                                                                                                                                                                                                                                                                                                                                                    |
| EndUserIP  | string | Required | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| ProductType | string | Required | Type of product to add. Permitted values are: - Certificate-Comodo-Essential - Certificate-Comodo-Instant - Certificate-Comodo-Essential-Wildcard - Certificate-Comodo-Premium-Wildcard - Certificate-Comodo-EV - Certificate-Comodo-EV-SGC - Certificate-GeoTrust-QuickSSL - Certificate-GeoTrust-QuickSSL-Premium - Certificate-GeoTrust-TrueBizID - Certificate-GeoTrust-TrueBizID-Wildcard - Certificate-GeoTrust-TrueBizID-EV - Certificate-RapidSSL-RapidSSL - Certificate-RapidSSL-RapidSSL-Wildcard - Certificate-VeriSign-Secure-Site - Certificate-VeriSign-Secure-Site-Pro - Certificate-VeriSign-Secure-Site-EV - Certificate-VeriSign-Secure-Site-Pro-EV - Certificate-Comodo-UCC-DV - Certificate-Comodo-UCC-OV |
| Quantity   | int | Required | Number of years to purchase this SSL certificate. Permitted values are 1 to the to the following maximum values: - 3 - Comodo Essential - 3 - Comodo Instant - 3 - Comodo Essential Wildcard - 3 - Comodo Premium Wildcard - 2 - Comodo EV - 2 - Comodo EV SGC - 3 - GeoTrust QuickSSL - 3 - GeoTrust QuickSSL Premium - 3 - GeoTrust True BusinessID - 3 - GeoTrust True BusinessID Wildcard - 2 - GeoTrust True BusinessID with EV - 3 - RapidSSL - 3 - RapidSSL Wildcard - 3 - VeriSign Secure Site - 3 - VeriSign Secure Site Pro - 2 - VeriSign Secure Site with EV - 2 - VeriSign Secure Site Pro with EV - 3 - Comodo UCC DV - 3 - Comodo UCC OV |
| ClearItems | string | Optional | Makes all items, currently in the cart, inactive \(but keep them in the cart\) and add a new item. Permitted value is "yes". To reactivate items in the cart that are inactive, use the "UpdateCart" command.                                                                                                                                                                                                                                                                                   |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Input parameters for POP Email
------------------------------

Use the input parameters in the table below to add POP Email to the cart.
Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=AddToCart&uid=YourAccountID&pw=YourApiToken&EndUserIP={Required}&ProductType={Required}&SLD={Required}&TLD={Required}&Quantity={Optional}&ProductID={Required for...}&ActionType={Optional for...}&TimeUnit={Optional for...}&TimeQuantity={Optional for...}&Capacity={Optional for...}&ClearItems={Optional}&responsetype={xml, text, or HTML}
```
| Input | Type  | Status | Description                                                                                                                                                            |
| ------------- | ------ | ----------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command | string | Required | AddToCart                                                                                                                                                             |
| UID      | string | Required                                   | Your Account ID |
| PW | string | Required | You API Token                                                                                                                                                           |
| EndUserIP   | string | Required                                   | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| ProductType | string | Required | Type of product to add. Permitted values are: - POP \(10-mailbox pak with 1GB storage per mailbox\) - Renew-POP - EmailStorage \(512MB extra storage per mailbox\) - Mailbox \(This is the new email type which allows for both creating new mailboxes and renewals.\)                            |
| SLD      | string | Required for: - POP                              | Second-level domain name \(for example, "enom" in "enom.com"\). |
| TLD | string | Required for: - POP | Top-level domain name \(extension\).                                                                                                                                              |
| Quantity   | int | Optional for: - POP - Renew-POP - EmailStorage - Mailbox           | For "ProductType=POP", the number of 10-mailbox paks to purchase for this domain. For "ProductType=Renew-POP", the number of years to add to this POP pak subscription. For "ProductType=EmailStorage", the number of 512MB units of storage to add to this pak. For "ProductType=Mailbox", the number of mailboxes to add. |
| QuantityUnit | string | Optional | Permitted values are: - **packs**: pack of 10 email addresses \(default\) - **boxes**: single email address                                                                                                            |
| ProductID   | int | Required for: - Renew-POP - EmailStorage. - Mailbox with "ActionType=Renew" | POP pak ID number. To retrieve this value, use the "GetPOP3" command. |
| ActionType | string | Optional for: - Mailbox | Possible values: - Create \(default\) - Renew \(ProductID is required for renewal\)                                                                                                                        |
| BillingPeriod | string | Optional for: - Mailbox                            | Possible values: - Year \(default\) - Month |
| Period | int  | Optional for: - Mailbox | Possible values: 1-12 number of months or years \(default is 1\)                                                                                                                                 |
| Capacity   | int | Optional for: - Mailbox                            | Possible values: - 1 - 1 Gigabyte of storage \(default\) - 3 - 3 Gigabytes of storage - 5 - 5 Gigabytes of storage - 10 - 10 Gigabytes of storage |
| ClearItems | string | Optional | Make all items currently in the cart inactive \(but keep them in the cart\), and add a new item. Permitted value is yes. To reactivate items in the cart that are inactive, use the "UpdateCart" command.                                                           |
| ResponseType | string | .Optional                                  | Format of response. Permitted values are Text \(default\), HTML, or XML |

Input parameters for Email and URL Forwarding
---------------------------------------------

Use the input parameters in the table below to add Email Forwarding or URL Forwarding to the cart.
Build the query string using this syntax:

| Input | Type  | Status | Description                                                                                                 |
| ------------ | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command | string | Required | AddToCart                                                                                                  |
| UID     | string | Required | Your Account ID |
| PW | string | Required | Your API Token                                                                                                |
| EndUserIP  | string | Required | End user’s IP address. This is used in fraud checking, as part of our order processing service. Use format NNN.NNN.NNN.NNN. |
| ProductType | string | Required | Type of product to add. Permitted values are: - EmailForwarding - EmailForwardingRenew - URLForwarding - URLForwardingRenew                                         |
| SLD     | string | Required | Second-level domain name \(for example, "enom" in "enom.com"\) |
| TLD | string | Required | Top-level domain name \(extension\)                                                                                     |
| Quantity   | int | Optional | Number of years. Default value is: 1 |
| ClearItems | string | Optional | Make all items currently in the cart inactive \(but keep them in the cart\), and add a new item. Permitted value is yes. To reactivate items in the cart that are inactive, use the "UpdateCart" command. |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Input parameters for Whois Publicity Service
--------------------------------------------

Use the input parameters in the table below to add Whois Publicity Service to the cart.
Build the query string using this syntax:

| Input | Types  | Status | Description                                                                                                 |
| ------------ | ------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command | string | Required | AddToCart.                                                                                                 |
| UID     | string | Required | Your Account ID. |
| PW | string | Required | Your API Token.                                                                                               |
| SLD     | string | Required | Second-level domain name \(for example, "enom" in "enom.com"\). |
| TLD | string | Required | Top-level domain name \(extension\).                                                                                    |
| ProductType | string | Required | Type of product to add. *Permitted values:* - WhoisPublicity |
| ClearItems | boolean | Optional | Make all items currently in the cart inactive \(but keep them in the cart\), and add a new item. Permitted value is yes. To reactivate items in the cart that are inactive, use the "UpdateCart" command. |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0, the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output | Type  | Description |
| ---------- | ------- | ------------------------------------------------------------------------------------------------ |
| ItemName | string | Domain name or service |
| ItemId | int   | Item number |
| Price | float  | Price for this product or service |
| ICANNFees | float  | Fees charged by ICANN for this product or service |
| CartItemID | int   | ID number, assigned by us, of this item in your shopping cart |
| ItemAdded | boolean | True indicates this entire response has reached you successfully. |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | "True" indicates this entire response has reached you successfully. |

Related Commands
----------------

[AddBulkDomains](../docs/domains/registration/addbulkdomains.md)

[DeleteFromCart](../docs/DeleteFromCart.md)

[GetCartContent](../docs/getcartcontent.md)

[GetHomeDomainList](../docs/domains/domain-management/general/get-home-domain-list.md)

[InsertNewOrder](../docs/domains/renewal/InsertNewOrder.md)

[PurchasePreview](../docs/purchasepreview.md)

[UpdateCart](../docs/UpdateCart.md)