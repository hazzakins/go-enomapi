GDPR
====

> ### For more information, please visit our GDPR articles:
>
>
>
> - [Overview](https://www.enom.com/support/the-gdpr)
> - [Reseller FAQ](https://www.enom.com/support/the-gdpr#reseller-faq)
> - [Support Center FAQ](https://help.enom.com/hc/en-us/articles/360003302691-GDPR-Reseller-FAQ)
> - \[WHOIS output\]\([https://www.enom.com/blog/wp-](https://www.enom.com/blog/wp-)
> content/uploads/2017/11/whois\_changes\_overview\_enom.pdf\)

The [European Union’s General Data Protection Regulation](https://www.eugdpr.org/) \(GDPR\) lays out a new set of rules for how the personal data of people living within the EU \(“EU-local individuals”\) should be handled. The policy comes into full effect on **May 25, 2018**, and we recommend that you start preparing now by speaking with a lawyer and familiarizing yourself with the information we’ve provided here.

Though it’s complex and far-reaching, at a high level, the GDPR can be understood in terms of three fundamental concepts:
1. Consent and Control
-----------------------

Clear, informed consent and individual control over the use of personal data are basic rights in the GDPR. Any business collecting or processing personal data must not only obtain consent to do so, but must also explain what they need the information for. What’s more, they’re only allowed to collect the minimum amount of information required to get the job done and can’t use the info for any purpose other than that to which the individual initially agreed. This puts the individual in charge of how their info is used from the very start.
2. Transparency
----------------

The GDPR imposes requirements around how companies should address security breaches that expose sensitive personal information. In the event of a breach, anyone whose information may have been exposed must be notified as soon as possible, and that notice should include an explanation of what happened, what’s being done to fix it, and what those affected should do to protect themselves. This type of information empowers each person to respond in the way they think is best in each circumstance in order to protect their own privacy.
3. The right to be forgotten
-----------------------------

Under these new rules, EU-local individuals have the right to revoke consent for a service provider to use their data. When this happens, the provider must essentially erase all record of the individual, giving them a fresh start. This requirement is not without consequences or limitations: some services can’t be provided without personal information, and sometimes personal information has to be kept for reasons of public interest or relating to legal claims.

Contact associated operations
-----------------------------

Any data that must be processed in order to register a domain, or provide any other type of service, will be covered under contract. We will be updating our **Registration Agreement** and **Reseller Agreement** to include mention of all these essential pieces of data:
- First name
- Last name
- Organization \(if provided\)
- Email address
- Country

> ### We will request consent from the data subject when:
>
>
>
> - We give the option of processing any piece of personal data that isn’t essential or necessary to provide the service. For example, for most domain registrations, we don’t require the registrant to provide their phone number, but by collecting this piece of data we are able to provide a backup verification method.
> - The data is required by a third party, with whom we do not yet have a GDPR-compliant contract. For example, a registry might require that the registrant’s postal address be on file in order to complete a domain registration. If we don’t have a GDPR-compliant contract with this particular registry, we would have to request consent from the data subject to process and share this extra piece of personal data before completing the registration.

API commands that may invoke GDPR consent
-----------------------------------------

More commands might be added if there are regulation or policy changes.

| Category      | API Commands |                                 |
| ------------------- | ------------------------------------------------- | ---------------------------------------------------------------- |
| Domain registration | [Purchase](../docs/domains/registration/purchase.md) | Purchase a domain name or premium domain in real time.     |
|           | [Preconfigure](../docs/domains/registration/Preconfigure.md) | Configure domain in the cart that requires extended attributes. |
|           | [InsertNewOrder](../docs/domains/renewal/InsertNewOrder.md) | Finalize purchase of the cart contents.             |
| Contacts update   | [Contacts](../docs/domains/domain-management/contacts/contacts.md) | Update contact information for a domain name.          |
|           | [GetContacts](../docs/domains/domain-management/contacts/getcontacts.md) | Get all contact data for a domain name.             |
| Domain transfer   | [TP\_CreateOrder](../docs/domains/transfer/TP_CreateOrder.md) | Transfer domains into an account.                |
| Domain push     | [PushDomain](../docs/domains/transfer/pushdomain.md) | Push a domain name into another account.            |
| Consent email    | [SendConsentEmail](../docs/sendconsentemail.md) | Send or resend consent email to the contact.          |

Consent level value is returned in some commands to inform user the status of the contact creation or changes.

| Consent level         | Description | Asynchronous |
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------ |
| NONE              | We have not yet requested or collected a consent choice for this product group from this data subject. | No      |
| PENDING            | Consent has been requested for this product group from this data subject but a selection has not yet been made. The product group is ‘synchronous’ so our system will complete the order using placeholder contact data for the consent-based fields unless and until consent to use real data is provided. Real data is used for contract-based fields. | No      |
| PENDING\_ASYNC         | Consent has been requested for this product group from this data subject but a selection has not yet been made. The product group is ‘asynchronous’ so our system cannot process the order unless and until consent to use real data is provided, at which time real data will be used for all contract-based fields. | Yes     |
| ACCEPTED\_CONTRACTUAL\_MINIMUM | The user \(data subject\) has indicated that they do not consent to any additional data use beyond that which is required by contract. Minimum information always held by Enom, required by contract includes first name, last name, organization, email address and country. Specific product groups may have different requirements for contract-based and consent-based data use, which will be indicated on the Data Use Consent Settings and Data Use Information pages. | No      |
| FORCED\_ALL\_CONTRACTUAL    | The product group only uses data based on a contract, there is no consent-based data use and so no consent is required. | No      |
| ACCEPTED\_FULL         | The data subject has provided full consent to use consent-based data elements; contract-based data elements are also used. | No      |
| DENIED             | This asynchronous product can only be ordered if the data subject consents to data use; they did not consent, so the order has been cancelled and the data is not used. Note: this option is only available for TLDs or products that are deemed to be asynchronously handled for GDPR, AND previously had a status of PENDING\_ASYNC \(once the pending period ends, the product moves to DENIED status\). | No      |