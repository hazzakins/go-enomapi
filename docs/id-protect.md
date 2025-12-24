ID Protect
==========

Whois Privacy Protection Service \(WPPS\)

What is ID Protect?
-------------------

Registry regulations require that valid contact information is provided for every domain name, and stored in the publicly accessible Whois database. Without identity protection, your personal or business contact information is available to anyone, at any time.

What domains are eligible for ID Protect?
-----------------------------------------

ID Protect can be applied to most domain names, but registry rules prohibit the masking of .us, .ca, .uk, .de, .eu, and some other country-code TLDs. To use our ID Protect product, the domain name must be registered with us or use our DNS hosting, and must use our name servers.

> ### You can use the GetTLDDetails command to check if a TLD supports ID Protect.
>
>
>
> [GETTLDDETAILS](http://www.enom.com/api/API%20topics/api_GetTLDDetails.htm)
>
>
>
> Parse your results for:\\ or

How do I enable ID Protect on a domain?
---------------------------------------

There are three ways to enable IDP for a domain name.

|                                                  | Command\(s\) |
| ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1. Add ID Protect at the time of purchase with the following parameter: WPPS           | [PURCHASE](http://www.enom.com/api/API%20topics/api_Purchase.htm) |
| 2. Add ID Protect when transferring a domain into your account with the following parameter: WPPS | [TP\_CREATEORDER](http://www.enom.com/api/API%20topics/api_TP_CreateOrder.htm)                                         |
| 3. Add ID Protect to an existing domain in your account                      | [PURCHASESERVICES](http://www.enom.com/api/API%20topics/api_PurchaseServices.htm) [ENABLESERVICES](http://www.enom.com/api/API%20topics/api_EnableServices.htm) |

> ### Adding ID Protect to an existing domain is a two-step process:
>
>
>
> **Step 1:** [PURCHASESERVICES](http://www.enom.com/api/API%20topics/api_PurchaseServices.htm) \(Purchase ID Protect for a domain in your account\)
>
> **Step 2:** [ENABLESERVICES](http://www.enom.com/api/API%20topics/api_EnableServices.htm) \(Enable ID Protect for the domain used in step 1.\)