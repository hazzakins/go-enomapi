GetSubAccountDetails
====================

For one subaccount, get the wholesale prices you charge this subaccount for products and services, and whether enabled settings for one subaccount.

Usage
-----

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/SubAccount-List.asp](https://resellertest.enom.com/resellers/SubAccount-List.asp)

On the sub-account page, clicking a link in the LoginID column calls the GetSubAccountDetails command without the Action=Manage parameter.

[https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=154-dz-5567](https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=154-dz-5567)

On the Sub-account configuration page, clicking the save changes button calls the GetSubAccountDetails command with the Action=Manage parameter.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

[https://resellertest.enom.com/pdq/RE\_Default.asp?maintab=overview](https://resellertest.enom.com/pdq/RE_Default.asp?maintab=overview)

Clicking the my site tab calls the GetServiceContact command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The account must be a reseller account.

The subaccount must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                                                                                                                 | Max Size |
| --------------- | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                                                                                                              | 20 |
| PW | Required            | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                                                                                  | 4 |
| Account | Required            | Subaccount ID number, in NNN-aa-NNNN format. You can retrieve this number using the GetSubAccounts command | 20    |
| Action     | Optional; default value is List | Action to take. Permitted values are List and Info. Action=List returns a long list of current price settings. Action=Info returns a short list of information about the subaccount and parent account                                                                  | 20 |
| UseQtyEngine | Optional            | Year bracket to retrieve prices for. For example, UseQtyEngine=2 retrieves the annual price for domain name registrations of 2 to 4 years; UseQtyEngine=5 retrieves the annual price for 5 to 9 years. Permitted values are 1, 2, 5, and 10. You may use only one UseQtyEngine parameter per query. Use this parameter with Action=List | ?    |

Returned Parameters and Values
------------------------------

| Output Parameter               | Description |
| -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command                   | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                               |
| ErrX                     | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                              |
| Action                    | Type of action you specified in the query string |
| Display | Which content to display here                                                                |
| PartyID                   | Party identification number expressed as 32 hexadecimal characters, hyphenated |
| LoginID | Account login ID of the subaccount                                                              |
| DomainCount                 | Number of domains in this subaccount |
| Account | Subaccount ID number, in NNN-aa-NNNN format                                                        |
| TLD                     | Top-level domain; for example, com for .com domains |
| RegisterPrice | Price for registering a domain of this type for one year                                                   |
| RegisterEnabled               | True/False—True indicates that this subaccount allows customers to register domain names |
| RenewPrice | Price for renewing a domain of this type for one year                                                    |
| RenewEnabled                 | True/False — True indicates that this subaccount allows customers to renew domain names |
| TransferPrice | paying for one year of registration                                                             |
| TransferEnabled               | True/False — True indicates that this subaccount allows customers to transfer domain names |
| POP3 | Price for one 10-mailbox POP pak, per year, in $US                                                     |
| DotNameBundle                | Price for a one-year registration of a .name domain, with the .name Registry’s email forwarding |
| DNSHosting | Price for one year of DNS hosting                                                              |
| WPPS                     | Price for one year of ID Protect Certificate-GeoTrust-QuickSSL-Premium Price for one year registration of a GeoTrust QuickSSL Premium SSL certificate |
| Certificate-GeoTrust-TrueBizID | Price for one year registration of a GeoTrust True BusinessID SSL certificate                                        |
| Certificate-RapidSSL-RapidSSL       | Price for one year registration of a RapidSSL certificate |
| Certificate-RapidSSL-RapidSSL-Wildcard | Price for one year registration of a RapidSSL Wildcard certificate                                              |
| Certificate-GeoTrust-TrueBizID-EV     | Price for one year registration of a GeoTrust True BusinessID EV certificate |
| Certificate-GeoTrust-QuickSSL | Price for one year registration of a GeoTrust QuickSSL certificate                                              |
| Certificate-GeoTrust-TrueBizID-Wildcard  | Price for one year registration of a GeoTrust True BusinessID Wildcard SSL certificate |
| Certificate-VeriSign-Secure-Site | Price for one year registration of a VeriSign Secure Site SSL certificate                                          |
| Certificate-VeriSign-Secure-Site-Pro   | Price for one year registration of a VeriSign Secure Site PRO SSL certificate |
| Certificate-VeriSign-Secure-Site-EV | Price for one year registration of a VeriSign Secure Site Extended Validation SSL certificate                                |
| Certificate-VeriSign-Secure-Site-Pro-EV | Price for one year registration of a VeriSign Secure Site Pro with Extended Validation SSL certificate |
| Certificate-Comodo-Essential | Price for one year registration of a Comodo Essential SSL certificate                                            |
| Certificate-Comodo-Instant         | Price for one year registration of a Comodo Instant SSL certificate |
| Certificate-Comodo-Premium-Wildcard | Price for one year registration of a Comodo PremiumWildcard SSL certificate                                         |
| Certificate-Comodo-Essential-Wildcard   | Price for one year registration of a Comodo Essential Wildcard SSL certificate |
| Certificate-Comodo-EV | Price for one year registration of a Comodo EV SSL certificate                                                |
| Certificate-Comodo-EV-SGC         | Price for one year registration of a Comodo EV SGC SSL certificate |
| Seal-VeriSign-Trust-Seal | Price for one year subscription to VeriSign Trust Seal                                                    |
| LWSC-Basic                 | Price for one month of Web Site Creator Basic |
| LWSC-Full | Price for one month of Web Site Creator Full                                                         |
| LWSC-Ecommerce               | Price for one month of Web Site Creator eCommerce |
| RichContentFreeTrial | Price for RichContent free trial, per month                                                         |
| RichContent                 | Price for RichContent, per month. Additional prices for further down in the response, in the ValueAdd nodes and ProductType and following return parameters |
| DNSHostingRenew | Price to renew one year of DNS hosting                                                            |
| POPStorage                  | Price for one unit of extra POP storage, for one year |
| POPStorageRenew | Price to renew one year of extra POP storage, for one year                                                  |
| POPRenew                   | Price to renew POP service for one year, for one 10-mailbox POP pak |
| DotNameMail | Price for one year of email forwarding by the .name Registry, for one domain name                                      |
| DotNameRegMailRenew             | Price to renew one year of registration and email forwarding by the .name Registry |
| EmailForwardingRenew | Price to renew our email forwarding service, for one year and one domain name                                        |
| URLForwardingRenew              | Price to renew our URL forwarding service, for one year and one domain name |
| WPPSRenew | Price to renew ID Protect for one year, for one domain                                                    |
| WBL                     | Price for Business Listing for one year, for one domain |
| WBLRenew | Price for a one-year renewal of Business Listing                                                      |
| ProductType                 | Product ID number |
| Product | Product name                                                                         |
| ProductEnabled                | Enabled status of this product |
| Price | Price of this product. Both Email Forwarding and URL Forwarding are yearly subscriptions.                                 |
| TimeBracket                 | Billing period for this product-billing cycle-price set |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests for a subaccount of resellid, subaccount 659-fs-2869. The query retrieves the prices for all products and requests the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetSubaccountDetails&uid=resellid&pw=resellpw
&Account=659-fs-2869&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetSubaccountDetails&uid=resellid&pw=resellpw
&Account=659-fs-2869&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetSubaccountDetails&uid=resellid&pw=resellpw
&Account=659-fs-2869&responsetype=text
```
The response lists the prices given in the query for each of the domain name products, and then lists identification information and the updated prices for subaccount 154-dz-5567:

```html
<?xml version="1.0" ?>

 <YearBracket />

 <Action />

 <Display>all</Display>

  <PartyID>{18FC24BC-07D5-4591-8A16-A4D32DF934B4}</PartyID>

  <LoginID>resellidsub</LoginID>

  <DomainCount>0</DomainCount>

  <Account>659-fs-2869</Account>

  <Reseller>1</Reseller>

  <Site>enom</Site>

   <tld>com</tld>

   <tldid>0</tldid>

   <registerprice>29.95</registerprice>

   <resellerpricereg>29.95</resellerpricereg>

   <registerenabled>True</registerenabled>

   <renewprice>29.95</renewprice>

   <resellerpricerenew>29.95</resellerpricerenew>

   <renewenabled>True</renewenabled>

   <transferprice>29.95</transferprice>

   <resellerpricetran>29.95</resellerpricetran>

   <transferenabled>True</transferenabled>

  </product>

   <tld>net</tld>

   <tldid>1</tldid>

.

.

.

  <count>78</count>

  </pricestructure>

  <pop3>9.95</pop3>

  <dotnamebundle>19.95</dotnamebundle>

  <DNSHosting>4.00</DNSHosting>

  <wpps>6.00</wpps>

.

.

.

   <producttype>190</producttype>

   <productenabled>True</productenabled>

   <price>12.50</price>

   <timebracket>1</timebracket>

  </valueadd>

   <price>10.41</price>

   <timebracket>12</timebracket>

.

.

.

  </ValueAdds>

 </Get>

 </SubAccountsManage>

 <Command>GETSUBACCOUNTDETAILS</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLER1-STG</Server>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.500</ExecTime>

 <Done>true</Done>

 <![CDATA[ ]]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Command: GETSUBACCOUNTDETAILS

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.063

Done: true

RequestDateTime: 2/4/2015 12:38:43 PM
```
RequestDateTime=2/4/2015 12:39:17 PM
```