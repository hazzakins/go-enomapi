IRTP
====

Inter-Registrar Transfer Policy

> ### Effective date: December 1st, 2016
>
>
>
> ICANN \(the Internet Corporation for Assigned Names and Numbers\), the organization that effectively governs Internet infrastructure, announced an official date for a **new policy regarding domain name transfers for all gTLDs \(non-country code domains\)**. The change will go into effect for all users on December 1, 2016. Read more at... [https://www.icann.org/news/announcement-2016-06-01-en](https://www.icann.org/news/announcement-2016-06-01-en)

> ### IRTP applies when...
>
>
>
> ...you or your customer updates any of the following **Registrant Contact Information** on any gTLD domain names \(non-country code domains\) in your Enom account.
>
>
>
> - Registrant First Name
> - Registrant Last Name
> - Registrant Organization Name
> - Registrant Email

Enom will fully support IRTP through all of our retail properties and via the API for our channel partners. Below is a high-level diagram of the customer experience and corresponding developer experience to better help you visualize the upcoming changes and what you may need to do to support them in your own integrations.
1. Before Registrant Contact Changes
-------------------------------------

The **Customer Experience** is how we will support these changes in all of our retail properties. The **Developer Experience** is how you can provide the same or similar experience in your own custom integrations, **although no changes are required to support IRTP**.

**This is the state before any registrant changes are made to an existing domain name in your account.** Two very common commands to get domain information and contact information for a domain will now include additional information about IRTP.

| API Commands                   | |
| ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| [GetDomainInfo](../docs/getdomaininfo.md#IRTP) | Returns the current domain information along with the current IRTP Settings for this domain. |
| [GetContacts](../docs/getcontacts.md#IRTP) | Returns the current contact information along with the current IRTP OptOut status for this domain's registrant contact information. |
2. Making Registrant Contact Changes
-------------------------------------

In the UI, we will provide prompts to help the customer understand that any changes they make which result in updating a domain name's **Registrant Contact Information** will invoke IRTP.

Changes will be seen immediately, but could potentially put the domain name in transfer lock status for 60-days if the account owner has not applied a global opt-out setting in their account settings.

Note that the actual change event can override the global opt-out setting regardless. \(i.e. Enable the lock if globally it's disabled, or disable the transfer lock if globally it's enabled.\) This override option remains at the domain level per IRTP specifications.

**At the API, there are 3 commands that can potentially invoke IRTP.**

| API Commands                       | |
| -------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [PushDomain](../docs/pushdomain.md#IRTP)        | When pushing a domain from one Enom account to another Enom account, you have the choice to push the domain's existing contact information too and this will not invoke IRTP. However, if you elect to push a domain and that causes the domain's registrant contact information to change in the process, then this will invoice IRTP for each domain in the push. |
| [UpdateAccountInfo](../docs/updateaccountinfo.md#IRTP) | When you update your account's global information **AND** that account information is used as any of your domain's registrant contact information, then you will invoke IRTP for each domain where this is true.                                                                           |
| [Contacts](../docs/contacts.md#IRTP)          | When you update the registrant's contact information on a domain name, you may invoke IRTP if you are changing the first name, last name, org name, or email address. |
3. After Registrant Contact Changes \(Change Notification\)
------------------------------------------------------------

Once changes are applied, a change notice email is sent to the prior email and the new email per IRTP specifications. If no one elected to opt-out of the 60-day transfer lock, then the domain will be locked. If they did opt-out, then the domain will not be locked and can be transferred to another registrar/owner if that is desired.

FAQ
---

> ### How does this affect WHOIS privacy updates \(enable, disable, "rolling contact" details\)?
>
>
>
> According to ICANN, these changes are no different than customer-invoked changes. They will trigger IRTP if the first name, last name, org name, or email changes, which is typically the case for enabling, disabling, or rolling the contact details on a domain name with WHOIS privacy. However, for IDP provided by and resold through Enom, as the **designated agent**, we will be able to apply these changes on your behalf without invoking IRTP. If you are managing your own WHOIS privacy service, which is not IDP provided by Enom, you will invoke IRTP when making the changes above.