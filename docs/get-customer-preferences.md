GetCusPreferences
=================

Usage
-----

Retrieve global preferences for an account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The login ID and API Token must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetCusPreferences&uid=YourAccountID&pw=YourApiToken&responsetype=xml
```
| Input Parameter | Type | Description                                          |
| --------------- | ------ | --------------------------------------------------------------------------------------------- |
| Command     | string | GetCusPreferences                                       |
| UID       | string | Your Account ID                                        |
| PW       | string | Your API Token                                        |
| ResellerKey   | string | Reseller EC key for this Instant Reseller or Registry Rocket site, if different from this UID |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter    | Type | Description                                                                  |
| ---------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| DefPeriod       | integer | Default period for registrations, renewals, and transfers, in years.                                     |
| AllowDNS        | boolean | True allows name servers other than eNom’s.                                                 |
| ShowPopups       | boolean | True shows pop-up menus.                                                          |
| AutoRenew       | boolean | True automatically renews the domain 30 days before it expires.                                       |
| RegLock        | boolean | True requires the account holder’s permission to transfer the domain to another registrar.                          |
| AutoPakRenew      | boolean | True automatically renews POP paks 30 days before they expire.                                        |
| UseDNS         | boolean | True uses eNom’s name servers.                                                        |
| ResellerStatus     | boolean | Is this a reseller account.                                                         |
| RenewalSetting     | integer | Possible values: - 0 indicates no email - 1 indicates send notice - 2 indicates contact and charge customer                  |
| RenewalBCC       | integer | Possible values: - 0 indicates no copy to reseller - 1 indicates send copy of email to reseller                        |
| RenewalURLForward   | boolean | True automatically renews URL forwarding 30 days before the URL forwarding subscription expires.                       |
| RenewalEmailForward  | boolean | True automatically renews email forwarding 30 days before the email forwarding subscription expires.                     |
| MailNumLimit      | integer | Maximum email records allowed for EmailForward.                                               |
| IDProtect       | boolean | Auto add IDProtect to domain registration added to the cart.                                         |
| DefIDProtectRenew   | integer | Default renewals for IDProtect in years.                                                   |
| NameJetSales      | boolean | Is this account part of the NameJet sales?                                                  |
| IRTPOptOut       | boolean | IRTPOptOut=True auto-approve any changes made in the Registrant contact for any ICANN Compliant domains.                  |
| IRTPEmailLanguageCode | string | IRTP email language option. Permitted values are: - EN – English - IT – Italian - FR – French - PT – Portuguese - ES – Spanish - DE – German |
| `<DefaultHostRecords>` | object | Default host record setting for new domain registration. Items: - hostname - address - recordtype Example:\\                 |
| DefaultHostRecordOwn  | boolean | Is this account using its own set of default host records?                                          |
| UseOurDNS       | boolean | Does this account use eNom’s DNS servers by default?                                             |
| `<NameServers>`    | object | Default nameserver setting for new domain registration if UseOurDSN=false Example:`<DNS1>dns1.name-services.com</DNS1>`            |
| AcceptTerms      | boolean | True indicates this account has signed a credit-card processing agreement with us.                             |
| URL          | string | URL for reseller site listed in email notices.                                                |
| ParentAccount     | string | This account’s parent.                                                            |
| ParentLogin      | string | This account’s parent login ID.                                                       |
| NoService       | boolean | Service status.                                                               |
| BulkRegLimit      | integer | Maximum number of names this account can register in a single query.                                     |
| Account        | string | This account ID.                                                               |

Response Params
---------------

```
<?xml version="1.0" encoding="UTF-8"?>
<interface-response>
  <CustomerPrefs>
   <DefPeriod>2</DefPeriod>
   <AllowDNS>False</AllowDNS>
   <ShowPopups>False</ShowPopups>
   <AutoRenew>False</AutoRenew>
   <RegLock>True</RegLock>
   <AutoPakRenew>True</AutoPakRenew>
   <UseDNS>True</UseDNS>
   <ResellerStatus />
   <RenewalSetting>2</RenewalSetting>
   <RenewalBCC>0</RenewalBCC>
   <RenewalURLForward>False</RenewalURLForward>
   <RenewalEmailForward>False</RenewalEmailForward>
   <MailNumLimit>99</MailNumLimit>
   <IDProtect>True</IDProtect>
   <DefIDProtectRenew>False</DefIDProtectRenew>
   <DefWBLRenew>False</DefWBLRenew>
   <NameJetSales>False</NameJetSales>
   <IRTPOptOut>True</IRTPOptOut>
   <IRTPEmailLanguageCode>EN</IRTPEmailLanguageCode>
   <defaulthostrecords>
     <hostrecord hostname="*" address="69.25.142.5" recordtype="A" />
     <hostrecord hostname="WWW" address="" recordtype="A" />
     <hostrecord hostname="@" address="" recordtype="A" />
     <hostrecord hostname="" address="" recordtype="A" />
     <hostrecord hostname="" address="" recordtype="A" />
   </defaulthostrecords>
   <defaulthostrecordown>False</defaulthostrecordown>
   <UseOurDNS>True</UseOurDNS>
   <NameServers>
     <DNS1>dns1.name-services.com</DNS1>
     <DNS2>dns2.name-services.com</DNS2>
     <DNS3>dns3.name-services.com</DNS3>
     <DNS4>dns4.name-services.com</DNS4>
     <DNS5>dns5.name-services.com</DNS5>
   </NameServers>
  </CustomerPrefs>
  <CustomerInformation>
   <AcceptTerms>True</AcceptTerms>
   <URL />
   <ParentAccount>000-00-0000</ParentAccount>
   <ParentLogin>Main-eNom</ParentLogin>
   <NoService>False</NoService>
   <BulkRegLimit>3000</BulkRegLimit>
   <Account>662-jo-2898</Account>
  </CustomerInformation>
  <Command>GETCUSPREFERENCES</Command>
  <APIType>API</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod />
  <MaxPeriod>10</MaxPeriod>
  <Server>RESELLERTEST</Server>
  <Site>eNom</Site>
  <IsLockable />
  <IsRealTimeTLD />
  <TimeDifference>+0.00</TimeDifference>
  <ExecTime>0.000</ExecTime>
  <Done>true</Done>
  <TrackingKey>4d0eda04-74f6-4268-ab2a-9052769ce408</TrackingKey>
  <RequestDateTime>10/19/2016 11:23:35 AM</RequestDateTime>
  <debug />
</interface-response>
```
Related Commands
----------------

GetConfirmationSettings

[UpdateCusPreferences](../docs/update-customer-preferences.md)

UpdateRenewalSettings