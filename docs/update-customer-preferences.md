UpdateCusPreferences
====================

Usage
-----

Update the customer preference settings for an account.

> ### For most settings, if no value is supplied for an individual setting, this command reverts that setting to its default value, with two exceptions:
>
>
>
> - Host Records: the UseParentDefault parameter allows you to choose between persisting current host records, and setting new host records.
> - Name Servers: always persisted by default, unless new name servers are specified in the query string.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The login ID and API Token must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=UpdateCusPreferences&uid=YourAccountID&pw=YourApiToken&responsetype=xml&{param1}={value1}
```
| Input Parameter    | Type | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| --------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command        | string | GetCusPreferences                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| UID          | string | Your Account ID                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| PW          | string | Your API Token                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| DefPeriod       | integer | Default period for auto-renew and registrar lock settings, in years. A domain’s registration period cannot extend more than 10 years beyond today. - Default is 1\*.                                                                                                                                                                                                                                                                                                                                                                                                           |
| RegLock        | boolean | RegLock=on prevents a domain from being transferred to a new registrar. RegLock=off allows unrestricted transfer of a domain from one registrar to another. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                      |
| AutoRenew       | boolean | AutoRenew=on renews domain names automatically. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| IDProtect       | boolean | IDProtect=on automatically purchases ID Protect for all eligible domains. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| DefIDProtectRenew   | boolean | DefIDProtectRenew=on automatically renews ID Protect before it expires. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| DefWBLRenew      | boolean | DefWBLRenew=on automatically renews Business Listing before it expires. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| AutoPakRenew     | boolean | AutoPakRenew=on renews POP email paks automatically. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| EMailForwardRenew   | boolean | EMailForwardRenew=on renews email forwarding automatically. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| URLForwardingRenew  | boolean | URLForwardingRenew=on renews URL forwarding automatically. - Default is off\*.                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| AllowDNS       | boolean | \(Obsolete\) Allow domain name servers.                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| UseDNS        | boolean | Which name servers a domain is using. - UseDNS=1, use our name servers. - UseDNS=0, use the name servers specified in this query string. Note: If UseDNS or DNS*x* are present in the UpdateCusPreferences other settings \(DefPeriod through AllowDNS\) will be updated.                                                                                                                                                                                                                                                                                                                                                         |
| DNS*x*        | string | Use name of domain name server. For example, DNS1=ns1.name-services.com. Note: If UseDNS or DNS*x* are present in the UpdateCusPreferences other settings \(DefPeriod through AllowDNS\) will be updated.                                                                                                                                                                                                                                                                                                                                                                                        |
| UseParentDefault   | boolean | Which host records should be used for this account’s defaults. Permitted values are: - 0 Use the host records in this query string - 1 Persist the existing host records                                                                                                                                                                                                                                                                                                                                                                                                           |
| HostName       | string | Name of host record, for example, HostName1=[www](http://www/). Permitted format is a comma-delimited list.                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| RecordType      | string | Record type of host record. Permitted format is comma-delimited list of record types for each of the host records specified with the HostName parameter. List record types in the same order as the host name they correspond to. Permitted values are: - A IP address - AAAA IPv6 address - CNAME Alias record type, to associate a host name with another host - URL URL redirect - FRAME Frame redirect - MX Mail. Can be a host name under this domain name or the name of a mail server - MXE Mail Easy \(email forwarding\) TXT Text \(SPF\) record                                                                                                                                                                                                                |
| Address        | string | Address to redirect to. Permitted format is a comma-delimited list of the record types for each host record. List addresses in the same order as the host name they correspond to. Permitted values are: - RecordType=A, Address must be an IP address - RecordType=AAAA, Address must be an IPv6 address - RecordType=CNAME, Address must be a fully qualified domain name \(see Note\) or a host name defined in this domain - RecordType=URL, Address must be the exact URL of the page you redirect to, or an IP address, or a fully qualified domain name \(see Note\) - RecordType=FRAME, Address is the actual URL, or the IP address, or the fully qualified domain name \(see Note\) of the page you want to display when someone types Your\_Domain.com - RecordType=MX, Address must be a fully qualified domain name \(see Note\) or a host name defined in this domain - RecordType=MXE, Address must be an IP address - RecordType=TXT, Address is a text \(SPF\) record. |
| ShowPopUps      | boolean | \(Obsolete\) Show menus. - ShowPopUps=0 hides menus - ShowPopUps=1 shows menus                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| IRTPOptOut      | boolean | Global setting for IRTP OptOut. Only applicable to ICANN Compliant domains. - IRTPOptOut=1 ***auto-approves*** any changes made in the Registrant contact. - IRTPOptOut=0 enforces 60 days transfer lock for any changes made in the Registrant contact. - Default is empty\*. Our system will adjust the value automatically.                                                                                                                                                                                                                                                                                                                             |
| IRTPEmailLanguageCode | string | IRTP email language option. Permitted values are: - EN – English - IT – Italian - FR – French - PT – Portuguese - ES – Spanish - DE – German                                                                                                                                                                                                                                                                                                                                                                                                                         |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description       |
| ---------------- | ------ | ------------------------ |
| Command     | string | Name of command executed |

Response Params
---------------

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Command>UPDATECUSPREFERENCES</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod></MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>eNom</Site>
<IsLockable></IsLockable>
<IsRealTimeTLD></IsRealTimeTLD>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>c3be87ab-2f27-45e1-83c3-05027a25514d</TrackingKey>
<RequestDateTime>10/19/2016 12:02:26 PM</RequestDateTime>
```
Related Commands
----------------

GetConfirmationSettings

[GetCusPreferences](../docs/get-customer-preferences.md)

UpdateRenewalSettings