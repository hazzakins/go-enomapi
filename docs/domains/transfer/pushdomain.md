PushDomain
==========

Push a domain name into another account.

Usage
-----

Availability
------------

All resellers have access to this command.

Constraints
-----------

If the domain is in pending verification due to RAA or IRTP and there is a possibility of contact change, the domain push will be failed.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=PushDomain&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml&AccountID=destination_account&pushcontact=1
```
| Input Parameter    | Type | Description                                                                                                        |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command        | string | PushDomain                                                                                                        |
| uid          | string | Your Account ID                                                                                                      |
| pw          | string | Your API Token                                                                                                      |
| AccountID       | string | Login ID of the account to push the name into. Use GetSubAccounts to list subaccounts.                                                                 |
| PushContact      | boolean | Push the contact information for this domain from the old account to the new account. Permitted values are: - 0 Do not push contact information with the domain - 1 Push contact information with the domain \(default\) |
| IRTPOptOut      | boolean | IRTPOptOut=True ***auto-approves*** any changes made in the Registrant contact for any ICANN Compliant domains.                                                     |
| IRTPOptOutReason   | string | Reason for IRTP OptOut                                                                                                  |
| IRTPEmailLanguageCode | string | IRTP email language option. Permitted values are: - EN – English - IT – Italian - FR – French - PT – Portuguese - ES – Spanish - DE – German                                       |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                |
| ---------------- | ------ | ------------------------------------------ |
| DomainName    | string | Domain name, SLD, TLD and DomainNameID   |
| PushDomain    | int | Returns =1 if successful or =0 if it fails |

Response Params
---------------

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Push>
<domainname sld="4ecessio0aryrickety" tld="com" id="366121791">4ecessio0aryrickety.com</domainname>
<PushDomain>1</PushDomain>
</Push>
<Command>PUSHDOMAIN</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>KRKDT198</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+08.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>cc383e6f-e4d0-485d-9a3b-2714f8254b96</TrackingKey>
<RequestDateTime>9/20/2016 11:31:13 AM</RequestDateTime>
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