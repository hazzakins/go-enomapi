GetContacts
===========

Get all contact data for a domain name.

Usage
-----

Use this command to display contact information for a domain within your account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

None

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetContacts&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Type | Description                        |
| --------------- | -------- | ---------------------------------------------------------- |
| command     | string | GetContacts                        |
| uid       | string | Your Account ID                      |
| pw       | string | Your API Token                       |
| TLD       | Required | Top-level domain name \(extension\)            |
| SLD       | Required | Second-level domain name \(e.g. "enom" in "enom.com"\) |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter         | Type | Description                                                                                                                                                                                                                                                                                  |
| --------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| DomainName            | string | Domain name, SLD, TLD and DomainNameID                                                                                                                                                                                                                                                                     |
| `<Registrant>`          | object | Collection of Registrant contact information, including RegistrantConsentStatus to reflect the current Consent Level of the registrant contact, as outlined [here.](../docs/gdpr.md)                                                                                                                                                                                            |
| `<AuxBilling>`          | object | Collection of Auxiliary Billing information                                                                                                                                                                                                                                                                  |
| `<Tech>`             | object | Collection of Technical contact information                                                                                                                                                                                                                                                                  |
| `<Admin>`             | object | Collection of Administrative contact information                                                                                                                                                                                                                                                                |
| `<Billing>`            | object | Collection of Billing contact information                                                                                                                                                                                                                                                                   |
| ICANNCompliant          | boolean | Is the TLD of this domain under ICANN Compliant?                                                                                                                                                                                                                                                                |
| PendingVerification        | boolean | Is there a pending contact update for this domain? See RAA and IRTP Response Params.                                                                                                                                                                                                                                             |
| RAA\_Status            | boolean | Are there any RAA related activities? Possible values: - True - False                                                                                                                                                                                                                                                     |
| RAA\_StatusDesc          | string | RAA status description. Possible values: - Null - Pending Suspension - Suspended                                                                                                                                                                                                                                               |
| IRTPOptOut            | boolean | IRTP opt-out setting at the account level. IRTPOptOut=True ***auto-approves*** any changes made in the Registrant contact for any ICANN Compliant domains                                                                                                                                                                                                          |
| TransferLock           | string | True if the domain is currently locked due to Registrant contact changes under IRTP enforcement.                                                                                                                                                                                                                                       |
| TransferLockExpDate        | string | Transfer lock expiration date \(Pacific Time\) When the Registrant contact is changed \(*first name*, *last name*, *organization* or *email address*\) AND the domain is under ICANN Compliant, the domain is set to be locked for 60 days. If another Registrant contact change is triggered afterward, the lock will be reset back to 60 days. Attributes: - *DaysRemaining:* the days remaining until the 60 days lock is removed. - *UTC:* expiration date time value and format in UTC standard. - *Epoch:* expiration date time value and format in Linux standard. |
| Nexus               | string | Nexus category                                                                                                                                                                                                                                                                                 |
| Purpose              | string | Nexus purpose                                                                                                                                                                                                                                                                                 |
| CurrentAttributes         | string | Nexus attributes                                                                                                                                                                                                                                                                                |
| WPPSAllowed            | boolean | Is Whois Privacy Protection Service allowed for this domain?                                                                                                                                                                                                                                                          |
| WPPSExists            | boolean | Is Whois Privacy Protection Service exist for this domain?                                                                                                                                                                                                                                                           |
| WPPSEnabled            | boolean | Is Whois Privacy Protection Service enabled for this domain?                                                                                                                                                                                                                                                          |
| WPPSExpDate            | datetime | Whois Privacy Protection Service expiration date                                                                                                                                                                                                                                                                |
| WPPSAutoRenew           | boolean | Whois Privacy Protection Service auto-renew flag                                                                                                                                                                                                                                                               |
| `<WPPSContactData>`        | object | Collection of Whois Privacy Protection Service contact information                                                                                                                                                                                                                                                       |
| EscrowLiftDate          | datetime | Escrow lift date for this domain                                                                                                                                                                                                                                                                        |
| EscrowHold            | boolean | Escrow holding flag                                                                                                                                                                                                                                                                              |
| `<Attributes>`          | object | Collection of extended attributes                                                                                                                                                                                                                                                                       |
| IsContactShared          | boolean | Is this domain contact shared?                                                                                                                                                                                                                                                                         |
| ContactRestrictedTLD       | boolean | Does this TLD have restricted contact regulation?                                                                                                                                                                                                                                                               |
| `<WhoisPublicity>` VASItemID   | integer | ID number of the VAS Item.                                                                                                                                                                                                                                                                          |
| `<WhoisPublicity>` ProdStatusID  | integer | Product Status ID. Expected values: - 1:Awaiting Configuration - 2:Service Active - 3:Billing Failed - 4: Pending Renewal - 5: Cancellation Pending - 6: Service Canceled - 7: Pending Expiration - 8: Service Expired - 9: Service Deleted                                                                                                                                                                  |
| `<WhoisPublicity>` ProdStatusDesc | string | Product Status Description. See above.                                                                                                                                                                                                                                                                    |
| `<WhoisPublicity>` ProdEnabled  | boolean | User control for WHOIS \(WPS\) visibility. The registrant must consent to the GDPR Whois Publicity before this flag is honored by the system. Expected values: - True - False                                                                                                                                                                                                |
| `<WhoisPublicity>` ProdConsented | boolean | Flag or indicator to show if the registrant has consented to the GDPR Whois Publicity. Expected values: - True - False                                                                                                                                                                                                                            |
| `<WhoisPublicity>` AutoRenew   | boolean | Auto renewal flag. Expected values: - True - False                                                                                                                                                                                                                                                              |
| `<WhoisPublicity>` ExpDate    | datetime | Value added service **expiration** date time. Format: MM/DD/YYYY HH:mm:SS AM/PM \(Pacific Time\) UTC and Epoch times are also available in the output.                                                                                                                                                                                                            |

Response Params
---------------

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
	<GetContacts>
	<domainname sld="enom" tld="com" domainnameid="12345" ><![CDATA[enom.com]]></domainname>
	<Registrant>...</Registrant>
	<AuxBilling>...</AuxBilling>
	<Tech>...</Tech>
	<Admin>...</Admin>
	<Billing>...</Billing>
	<PendingVerification>True</PendingVerification>
	<RAA_Status>True</RAA_Status>
	<RAA_StatusDesc><![CDATA[Suspended]]></RAA_StatusDesc>
 <IRTPOptOut>True</IRTPOptOut>
 <TransferLock>True</TransferLock>
 <TransferLockExpDate DaysRemaining="56" Epoch="1481623801" Utc="2016-12-13T18:10:01.677Z" >12/13/2016 10:10 AM</TransferLockExpDate>
	<Nexus category="" />
	<Purpose />
	<CurrentAttributes></CurrentAttributes>
	<WPPSAllowed><![CDATA[1]]></WPPSAllowed>
	<WPPSExists><![CDATA[1]]></WPPSExists>
	<WPPSEnabled><![CDATA[0]]></WPPSEnabled>
	<WPPSExpDate><![CDATA[6/28/2014]]></WPPSExpDate>
	<WPPSAutoRenew><![CDATA[No]]></WPPSAutoRenew>
	<WPPSContactData>...</WPPSContactData>
	<escrowliftdate />
	<escrowhold />
	<Attributes>
	</Attributes>
	<IsContactShared>False</IsContactShared>
</GetContacts>
```
Related Commands
----------------

[Contacts](../docs/contacts.md)

[GetDomainInfo](../docs/getdomaininfo.md)

[GetCusPreferences](../docs/get-customer-preferences.md)

[UpdateCusPreferences](../docs/update-customer-preferences.md)