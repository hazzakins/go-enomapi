Whois Publicity Service
=======================

The Whois Publicity service lets registrants choose to publish their real contact information in the public Whois record.

**Why would a registrant want Whois Publicity?**

Some domain owners, including domain investors and online business owners, want their contact information to be visible on their domain’s public Whois record. This can help demonstrate ownership of a domain, or indicate that it is available for sale.

> ### Important note about consent for Whois Publicity
>
>
>
> - The domain owner only needs to give consent for Whois Publicity once. If the service is later enabled on another domain, it will activate automatically.
> - Whois Publicity can be disabled and re-enabled on a per-domain basis or in bulk at any time via the Control Panel or API, and the domain owner can activate and deactivate the service for all their domains when needed simply by changing the consent choice on their Consent Settings page.

Step 1 - Purchase the service
------------------------------

The product or service must be purchased for the domain\(s\) that are located in the same account or sub-retail account.

| Action                                                               | API Command | Permitted API Param\(s\)                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Purchase a new domain along with Whois Publicity                                          | [Purchase](../docs/domains/registration/purchase.md) | - \*service=**WhoisPublicity** \*                                                                                      |
| Purchase Whois Publicity for a domain in your account or sub-retail account                            | [PurchaseServices](../docs/purchaseservices.md) | - \*service=**WhoisPublicity** \*                                                                                      |
| Purchase Whois Publicity for multiple domains in your account or sub-retail account \(Instant Purchase\)              | [AddBulkDomains](../docs/domains/registration/addbulkdomains.md) | - \*sldX=**first\_sld** \* - \*tldX=**first\_tld** \* - \*numyearsX=**first\_number\_of\_year** \* - \*producttype=**WhoisPublicity** \* - \*listcount=**total\_number\_of\_item** \* \*usecart=**false** \* |
| Purchase Whois Publicity for multiple domains in your account or sub-retail account \(Use Shopping Cart\)             | [AddBulkDomains](../docs/domains/registration/addbulkdomains.md) | - \*sldX=**first\_sld** \* - \*tldX=**first\_tld** \* - \*numyearsX=**first\_number\_of\_year** \* - \*producttype=**WhoisPublicity** \* - \*listcount=**total\_number\_of\_item** \* \*usecart=**true** \* |
| Add Whois Publicity to the shopping cart for domain\(s\) that already registered and located in your account or sub-retail account | [AddToCart](../docs/addtocart.md) | - \*producttype=**WhoisPublicity** \*                                                                                    |

Step 2 - Consent private data agreement
----------------------------------------

Once the order has been completed, the domain owner activates the service by providing consent for us to publish their personal data.
- A link to the Consent Settings page is sent via email as soon as the Whois Publicity service is ready to use for the domain\(s\).
- Domain owner can do this on the Data Use Consent Settings Page as well.

Step 3 - Control the private data visibility
---------------------------------------------

By default, the service enable flag is ON. However, if in any occasions the domain owner needs to turn it OFF \(or back to ON\), there are 3 options:

| Action                                                                                                              | API Command | Permitted API Param\(s\)                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Turn **OFF** private data for a domain                                                                                              | [VAS\_Update](../docs/vas-update.md) | - \*producttype=**WhoisPublicity** \* - \*vasitemid=**id\_for\_this\_item** \* - \*actiontype=**disable** \* Use the [VAS\_GetList](../docs/vas-get-list.md) command to retrieve the ID number. |
| Turn **ON** private data for a domain                                                                                              | [VAS\_Update](../docs/vas-update.md) | - \*producttype=**WhoisPublicity** \* - \*vasitemid=**id\_for\_this\_item** \* - \*actiontype=**enable** \* Use the [VAS\_GetList](../docs/vas-get-list.md) command to retrieve the ID number. |
| Turn **OFF** private data for a specific contact. This action will disable a particular contact that is attached or used by one or many domains. Domain owner needs to go to the Consent Setting page and revoke the consent. | - | N/A                                                                                                |

| Action                                                                                  | API Command | Permitted API Param\(s\)                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------- | ------------------------------------------------------------------------------ |
| Get item ID for a particular Whois Publicity attached to a domain name.                                                 | [VAS\_GetList](../docs/vas-get-list.md) | - \*producttype=**WhoisPublicity** \*                     |
| Get detailed information for a specific item. This is very useful to see if the product is in *good status* \(not expired\), *enable flag* status and *consent* status. | [VAS\_GetDetail](../docs/vas-get-detail.md) | - \*producttype=**WhoisPublicity** \* - \*vasitemid=**id\_for\_this\_item** \* |

Product life cycle and visibility logic
---------------------------------------

**Life cycle**

| ProdStatusID | ProdStatusDesc |
| ------------ | ---------------------- |
| 1      | Awaiting Configuration |
| 2 | Service Active     |
| 3      | Billing Failed |
| 4 | Pending Renewal    |
| 5      | Cancellation Pending |
| 6 | Service Canceled    |
| 7      | Pending Expiration |
| 8 | Service Expired    |
| 9      | Service Deleted |

**Private data visibility**

| ProdStatusID | ProdConsented | ProdEnabled | WHOIS display | Example scenarios |
| ------------ | ------------- | ----------- | ------------- | ------------------------------------------------------------------------------------------------------- |
| 2 | False     | False | Public    | Either the product is still in purchasing process or all flags are turned off by user. |
| 2 | False     | True | Public    | User just purchased the product. By default the switch is on, but has not consented. |
| 2 | True     | False | Public    | User switches off the product. |
| 2 | True     | True | **PRIVATE**  | Consented and enable flag is true. |
| Other than 2 | N/A      | N/A | Public    | Domain might have the product, it is either still in processing, deleted, expired or other conditions. |

**Private data and ID Protect**

If both ID Protect and Whois Publicity are enabled on the same domain, ID Protect will always take priority; queries for the domain in the public Whois will return the ID Protect fields instead of the registrant’s real information. ID Protect must be disabled for Whois Publicity to take effect.

| ID Protect | Whois Publicity | Display |
| ---------- | --------------- | --------------------------- |
| OFF | OFF       | Default \(GDPR-Protected\) |
| OFF | ON       | Real contact info |
| ON | OFF       | Privacy masking info |
| ON | ON       | Privacy masking info |

What contact info is displayed with Whois Publicity?
----------------------------------------------------

Only the registrant contact set will be published — Whois Publicity cannot be enabled for the domain’s admin, billing, or tech contact.

When Whois Publicity is enabled, all the standard data fields in the registrant contact set will be published. This includes:
- First and last name
- Organization \(if provided\)
- Street, City, State, and Zip or Postal Code
- Country
- Phone number
- Fax number \(if provided\)
- Email address

Which Whois Lookup results will be affected?
--------------------------------------------

Registrars and registries share their Whois data via “port 43 Whois servers”. Tucows/Enom will make the registrant data for all publicity-activated domains on platform accessible through this standard connection, meaning any Whois lookup service that pulls results from the Tucows/Enom Whois server, via port 43, will return the registrant details for any of our publicity-activated domains.