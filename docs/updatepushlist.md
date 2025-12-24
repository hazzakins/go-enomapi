UpdatePushList
==============

Push a domain name into another account.

Usage
-----

Push a domain name list into another account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

If the domain is in pending verification due to RAA or IRTP and there is a possibility of contact change, the domain push will be failed.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=UpdatePushList&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml&AccountID=destination_account&pushcontact=1
```
| Input Parameter    | Type | Description                                                                                                         |
| --------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command        | string | UpdatePushList                                                                                                       |
| uid          | string | Your Account ID                                                                                                       |
| pw          | string | Your API Token                                                                                                       |
| PushToLoginID     | string | Login ID of the account to push the name into. Use GetSubAccounts to list subaccounts.                                                                  |
| DomainList      | string | CRLF-delimited list of domain names \(URLEncode value for CRLF is %0D%0A\).                                                                        |
| PushContact      | boolena | Push the contact information for this domain from the old account to the new account. Permitted values are: ``` 0 Do not push contact information with the domain 1 Push contact information with the domain (default) ``` |
| IRTPOptOut      | boolean | IRTPOptOut=True ***auto-approves*** any changes made in the Registrant contact for any ICANN Compliant domains.                                                      |
| IRTPOptOutReason   | string | Reason for IRTP OptOut                                                                                                   |
| IRTPEmailLanguageCode | string | IRTP email language option. Permitted values are: - EN – English - IT – Italian - FR – French - PT – Portuguese - ES – Spanish - DE – German                                        |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter      | Type | Description                               |
| -------------------------- | ------ | ------------------------------------------------------------------------ |
| DomainName         | string | Domain name, SLD, TLD and DomainNameID                  |
| DomainListWrongFormat   | string | Returns =Valid if in right format or =InValid if it fails        |
| DomainNameNotInYourAccount | string | Returns=Yes if in your account or throws an error if not in your account |
| DomainNameID        | int | DomainNameID                               |
| SuccessfulPush       | string | Returns =Yes if successfully pushed or =No if fails           |
| Status           | string | Returns = "Ok to push" if no lock is on the domain or throws an error  |

Response Params
---------------

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<UPDATEPUSHLIST>
<PUSH-DOMAIN>
<DomainName>api2010526-135234.org</DomainName>
<DomainListWrongFormat>Valid</DomainListWrongFormat>
<DomainNameNotInYourAccount>Yes</DomainNameNotInYourAccount>
<DomainNameID>345862486</DomainNameID>
<SuccessfulPush>Yes</SuccessfulPush>
<Status>Ok to push</Status>
</PUSH-DOMAIN>
<PUSH-DOMAIN>
<DomainName>api2009227-121541.net</DomainName>
<DomainListWrongFormat>Valid</DomainListWrongFormat>
<DomainNameNotInYourAccount>Yes</DomainNameNotInYourAccount>
<DomainNameID>340777946</DomainNameID>
<SuccessfulPush>Yes</SuccessfulPush>
<Status>Ok to push</Status>
</PUSH-DOMAIN>
</UPDATEPUSHLIST>
<Command>UPDATEPUSHLIST</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod></MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL10VWRESELL_T</Server>
<Site>eNom</Site>
<IsLockable></IsLockable>
<IsRealTimeTLD></IsRealTimeTLD>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>159ef3b3-cc74-407d-90ea-c48ef1b1faa6</TrackingKey>
<RequestDateTime>12/16/2016 11:05:33 AM</RequestDateTime>
```
RAA and IRTP Response Params
----------------------------

**PendingVerification** status is based on RAA\_Status and IRTP\_Status. Below is the logic:

| RAA\_Status | IRTP\_Status | PendingVerification |
| ----------- | ------------ | ------------------- |
| false    | false | false        |
| false    | true | true        |
| true    | false | true        |
| true    | true | true        |

```
<PendingVerification>True</PendingVerification>
<RAA_Status>True</RAA_Status>
<RAA_StatusDesc><![CDATA[Suspended]]></RAA_StatusDesc>
<IRTP_Status>False</IRTP_Status>
<IRTP_StatusDesc />
```