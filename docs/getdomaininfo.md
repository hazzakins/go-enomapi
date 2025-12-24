GetDomainInfo
=============

Get information about a single domain name.

Usage
-----

Use this command to display current information about a single domain name. This command also retrieves the DomainNameID value, which is required by some other commands in our API.

Availability
------------

All resellers have access to this command.

Constraints
-----------

None

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetDomainInfo&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Type | Description                                |
| --------------- | ------ | -------------------------------------------------------------------------- |
| command     | string | **GetDomainInfo**                             |
| uid       | string | Your Account ID                              |
| pw       | string | Your API Token                               |
| sld       | string | Second-level domain name \(for example, enom in enom.com\)        |
| tld       | string | Top-level domain name \(extension\)                    |
| ResponseType  | string | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter           | Type | Description                                                                                                                                                                                                                                                                                                                                                            |
| ------------------------------------ | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| SLD                 | string | Second-level domain name \(for example, enom in enom.com\)                                                                                                                                                                                                                                                                                                                                   |
| TLD                 | string | Top-level domain name \(extension\)                                                                                                                                                                                                                                                                                                                                               |
| DomainNameID             | integer | ID number of this individual domain                                                                                                                                                                                                                                                                                                                                                |
| Multy-LangSLD            | boolean | Is this a multi-language SLD?                                                                                                                                                                                                                                                                                                                                                  |
| Expiration              | string | Expiration date of this domain registration                                                                                                                                                                                                                                                                                                                                            |
| Registrar              | string | Registrar of this domain                                                                                                                                                                                                                                                                                                                                                     |
| RegistrationStatus          | string | Registration status of this domain                                                                                                                                                                                                                                                                                                                                                |
| Purchase-Status           | string | Purchase status of this domain                                                                                                                                                                                                                                                                                                                                                  |
| Party-ID              | string | Party ID of the account in which this domain is registered                                                                                                                                                                                                                                                                                                                                    |
| Belongs-To             | string | Login ID of the account in which this domain is registered                                                                                                                                                                                                                                                                                                                                    |
| EntryName              | string | Name of the product or service described in this node                                                                                                                                                                                                                                                                                                                                       |
| Changable              | booelan | Can this service be changed?                                                                                                                                                                                                                                                                                                                                                   |
| Service               | object | Setting for the service.                                                                                                                                                                                                                                                                                                                                                     |
| DNSServer              | integer | Domain Name Servers - 1006 use our name servers - 1012 use user-specified domain servers, including none DNSSettings \(Host Records\) - 1021 host records WSB \(Web site services\) - 1060 no Web site building or hosting services - 1063 Web Site Builder - 1066 Web site hosting accounts EmailSet \(Email Services\) - 1048 no email - 1051 email forwarding \(to a POP or WebMail address\) - 1054 user \(mail server name required\) - 1105 user simplified \(mail server’s IP address required\) - 1114 POP3/WebMail plus email forwarding WPPS \(ID Protect Whois Privacy Protection\) - 1120 WhoIs information is masked - 1123 WhoIs information is viewable WBL \(Business Listing\) - 1130 Business Listing settings |
| ServiceChangable           | boolean | Can this service be toggled on or off, or switched?                                                                                                                                                                                                                                                                                                                                        |
| ConfigurationChangable        | boolean | Is this a configurable product or service?                                                                                                                                                                                                                                                                                                                                            |
| WBLID                | string | Business Listing identification number                                                                                                                                                                                                                                                                                                                                              |
| StatusID               | string | Status number of this Business Listing                                                                                                                                                                                                                                                                                                                                              |
| StatusDescr             | string | Status ID of this Business Listing                                                                                                                                                                                                                                                                                                                                                |
| ExpDate               | string | Expiration date of this Business Listing                                                                                                                                                                                                                                                                                                                                             |
| Enabled               | boolean | Visibility status of this Business Listing in the WhoisBusinessListings.com directory. - 1 indicates Business Listing is visible; - 0 indicates Business Listing is hidden.                                                                                                                                                                                                                                                                          |
| Renew                | boolean | Automatic renewal setting of this Business Listing                                                                                                                                                                                                                                                                                                                                        |
| CompanyName             | string | Company name for this Business Listing                                                                                                                                                                                                                                                                                                                                              |
| CompanyDescription          | string | Description of company                                                                                                                                                                                                                                                                                                                                                      |
| DomainName              | string | Domain name to which this Business Listing is attached                                                                                                                                                                                                                                                                                                                                      |
| Street                | string | Street address of this Business Listing                                                                                                                                                                                                                                                                                                                                              |
| City                 | string | City of this Business Listing                                                                                                                                                                                                                                                                                                                                                   |
| PostalCode              | string | Postal code of this Business Listing                                                                                                                                                                                                                                                                                                                                               |
| Country               | string | Country of this Business Listing                                                                                                                                                                                                                                                                                                                                                 |
| CategoryIDX             | string | Category identification number\(s\) of this Business Listing                                                                                                                                                                                                                                                                                                                                   |
| FieldName              | string | Name of descriptive field                                                                                                                                                                                                                                                                                                                                                     |
| Value                | string | Descriptive content associated with this field name                                                                                                                                                                                                                                                                                                                                        |
| Name                 | string | Name of this host record                                                                                                                                                                                                                                                                                                                                                     |
| Type                 | string | Type of product, service, or host record                                                                                                                                                                                                                                                                                                                                             |
| Address               | string | Host record address                                                                                                                                                                                                                                                                                                                                                        |
| MXPref                | string | Priority value of this host record                                                                                                                                                                                                                                                                                                                                                |
| IRTPSettings             | object | IRTP settings for this domain                                                                                                                                                                                                                                                                                                                                                   |
| `<IRTPSettings>` ICANNCompliant   | boolean | Is the TLD of this domain under ICANN Compliant?                                                                                                                                                                                                                                                                                                                                         |
| `<IRTPSettings>` OptOut       | boolean | IRTP opt-out setting at the account level. OptOut=True ***auto-approves*** any changes made in the Registrant contact for any ICANN Compliant domains.                                                                                                                                                                                                                                                                                    |
| `<IRTPSettings>` TransferLock    | boolean | True if the domain is currently locked due to Registrant contact changes under IRTP enforcement.                                                                                                                                                                                                                                                                                                                 |
| `<IRTPSettings>` TransferLockExpDate | string | Transfer lock expiration date \(Pacific Time\) When the Registrant contact is changed \(*first name*, *last name*, *organization* or *email address*\) AND the domain is under ICANN Compliant, the domain is set to be locked for 60 days. If another Registrant contact change is triggered afterward, the lock will be reset back to 60 days. Attributes: - *DaysRemaining:* the days remaining until the 60 days lock is removed. - *UTC:* expiration date time value and format in UTC standard. - *Epoch:* expiration date time value and format in Linux standard.                                                                          |
| `<WhoisPublicity>` VASItemID     | integer | ID number of the VAS Item.                                                                                                                                                                                                                                                                                                                                                    |
| `<WhoisPublicity>` Enabled      | boolean | Whois Publicity flag. Expected values: - True - False Note: in order WPS to be enabled \(active state\), registrant must consent the GDPR requirement for WPS product. Please check the consent status using [VAS\_GetDetail](../docs/vas-get-detail.md) command.                                                                                                                                                                                                                              |
| Command               | string | Name of command executed                                                                                                                                                                                                                                                                                                                                                     |
| ErrCount               | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                                                                                                                                                                                                                                 |
| ErrX                 | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                                                                                                                                                                                                                                     |
| Done                 | boolean | True indicates this entire response has reached you successfully.                                                                                                                                                                                                                                                                                                                                |

Response Params
---------------

```html
<?xml version="1.0" encoding="UTF-8"?>
<interface-response>
  <GetDomainInfo>
   <domainname sld="01817" tld="com" domainnameid="152809550">01817.com</domainname>
   <multy-langSLD />
   <status>
     <expiration>4/1/2017 7:34:00 PM</expiration>
     <deletebydate>4/1/2017 7:34:00 PM</deletebydate>
     <deletetype />
     <restorable>True</restorable>
     <renewbeforeexpiration />
     <registrar>eNom, Inc.</registrar>
     <registrationstatus>Registered</registrationstatus>
     <purchase-status>Paid</purchase-status>
     <belongs-to party-id="{39AE68C0-D019-4690-9999-FD632BC1AFAA}">enomtest_reseller</belongs-to>
     <escrowhold />
     <escrowliftdate />
     <auctionhold>False</auctionhold>
     <auctionliftdate />
   </status>
   <ParkingEnabled>False</ParkingEnabled>
   <services>
     <entry name="dnsserver">
      <enomDNS value="YES" isDotName="NO" />
      <service changable="1">1006</service>
      <configuration changable="0" type="dns">
        <dns>dns1.name-services.com</dns>
        <dns>dns2.name-services.com</dns>
        <dns>dns3.name-services.com</dns>
        <dns>dns4.name-services.com</dns>
        <dns>dns5.name-services.com</dns>
      </configuration>
     </entry>
     <entry name="dnssettings">
      <service changable="0">1021</service>
      <configuration changable="1" type="host">
        <host>
         <name><![CDATA[*]]></name>
         <type><![CDATA[A]]></type>
         <address><![CDATA[10.7.37.52]]></address>
         <mxpref><![CDATA[10]]></mxpref>
         <iseditable><![CDATA[1]]></iseditable>
        </host>
        <host>
         <name><![CDATA[@]]></name>
         <type><![CDATA[A]]></type>
         <address><![CDATA[10.7.37.52]]></address>
         <mxpref><![CDATA[10]]></mxpref>
         <iseditable><![CDATA[1]]></iseditable>
        </host>
      </configuration>
     </entry>
     <entry name="wsb">
      <service changable="1">1063</service>
      <configuration changable="1" type="wsb">
        <wsb>LWSC442x5665</wsb>
        <siteid>1826</siteid>
        <prodtype>86</prodtype>
        <nextbilldate>Nov 11 2011 12:00AM</nextbilldate>
      </configuration>
     </entry>
     <entry name="emailset">
      <service changable="1">1048</service>
     </entry>
     <entry name="wpps">
      <service changable="1">1120</service>
      <configuration changable="1" type="id protect">
        <wpps>
         <cloakedemail><![CDATA[[email protected]]]></cloakedemail>
         <forward-to><![CDATA[[email protected]]]></forward-to>
         <expiredate>Apr 01, 2013</expiredate>
         <autorenew>No</autorenew>
        </wpps>
      </configuration>
     </entry>
     <entry name="wbl">
      <wbl>
        <statusid><![CDATA[0]]></statusid>
        <statusdescr><![CDATA[Available]]></statusdescr>
      </wbl>
     </entry>
     <entry name="mobilizer">
      <service changable="0">1117</service>
      <mobilizer />
     </entry>
     <entry name="parking">
      <service changable="1">1030</service>
     </entry>
     <entry name="messaging">
      <service changable="1">1087</service>
     </entry>
     <entry name="map">
      <service changable="1">1108</service>
     </entry>
     <entry name="raasettings">
      <service changable="0" />
      <raasetting>
        <verificationstatus><![CDATA[Suspended]]></verificationstatus>
        <domainsuspended><![CDATA[1]]></domainsuspended>
        <suspensiondate><![CDATA[2016-05-23T08:50:21.673]]></suspensiondate>
        <isqueuedchange><![CDATA[0]]></isqueuedchange>
        <statusexpdate><![CDATA[2016-05-21T13:46:08.617]]></statusexpdate>
      </raasetting>
     </entry>
     <entry name="irtpsettings">
      <irtpsetting>
        <optout>True</optout>
        <transferlock>True</transferlock>
        <transferlockexpdate daysremaining="54" utc="2016-12-13T18:10:01.000Z" epoch="1481623801">12/13/2016 10:10 AM</transferlockexpdate>
      </irtpsetting>
     </entry>
     <entry name="portalsettings">
      <service changable="1" />
     </entry>
     <entry name="whoispublicity">
       <service changable="1"></service>
       <whoispublicity>
         <vasitemid>1017648</vasitemid>
         <enabled>False</enabled>
       </whoispublicity>
     </entry>
   </services>
  </GetDomainInfo>
  <Command>GETDOMAININFO</Command>
  <APIType>API</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>RESELLERTEST</Server>
  <Site>eNom</Site>
  <IsLockable>True</IsLockable>
  <IsRealTimeTLD>True</IsRealTimeTLD>
  <TimeDifference>+08.00</TimeDifference>
  <ExecTime>0.000</ExecTime>
  <Done>true</Done>
  <TrackingKey>b4c29783-15a9-4f02-b9ea-a7fcaf6a51c3</TrackingKey>
  <RequestDateTime>10/19/2016 1:50:50 PM</RequestDateTime>
</interface-response>
```
Related Commands
----------------

[GetContacts](../docs/getcontacts.md)

[Contacts](../docs/contacts.md)

[GetCusPreferences](../docs/get-customer-preferences.md)

[UpdateCusPreferences](../docs/update-customer-preferences.md)