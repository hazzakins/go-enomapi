PE\GetRetailPricing
===================

Retrieve the retail prices that this account charges for all products, and their enabled status.

"Retail" prices apply to retail customers of this account. They also apply to any subaccounts in which the prices the subaccount pays have not been set for that individual subaccount using a command such as UpdateAccountPricing, GetSubAccountDetails, or PE\_SetPricing.

Usage
-----

Use this command to display the full list of retail prices for this account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/Settings.asp](https://resellertest.enom.com/myaccount/Settings.asp)

On the my enom > settings page, the Default Account/Sub-Account Pricing section displays the query results for PE\_GetRetailPricing.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                               | Max Size |
| --------------- | ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                            | 20 |
| PW | Required           | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                | 4 |
| TLDOnly | Optional; default is 0    | Use TLDOnly=1 if you want only domain name prices; otherwise this command returns prices for domain names and all other products and services. | 20    |
| Years      | Optional; default value is 1 | Beginning of year range for which to retrieve prices. Permitted values are 1, 2, 5, and 10. For example, Years=2 retrieves prices for registrations of 2 to 4 years. | 20 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                                                                                                                                                                                                                                                                                                                                                                         |
| ---------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                                                                                                                                                                                                                                                                                                                                                                   |
| ErrCount                          | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                                                                                                                                                                                                                                                                                  |
| Done                            | True indicates this entire response has reached you successfully. |
| TLD | Top-level domain name \(extension\)                                                                                                                                                                                                                                                                                                                                                                                             |
| TLDID                           | Our ID number for this TLD |
| ProductType for domains | Product type for domains. Permitted values are: 10 Domain registration 13 DNS hosting 14 DNS hosting renew 16 Domain renewal 17 Domain redemption grace period \(RGP\) 18 Domain Extended RGP \(available at our discretion, and decided by us on a name-by-name basis\) 19 transfer 41 Registration and email forwarding by the .name Registry 44 .name registration and email forwarding renewal                                                                                                                                                                                                           |
| ProductType for SSL certificates              | Product type for SSL certificates. Permitted values are: 211 SSL certificate - Comodo Essential 212 SSL certificate - Comodo Instant 214 SSL certificate - Comodo Essential Wildcard 213 SSL certificate - Comodo Premium Wildcard 221 SSL certificate - Comodo EV 222 SSL certificate - Comodo EV SGC 20 SSL certificate - GeoTrust QuickSSL Premium 21 SSL certificate - GeoTrust True BusinessID 24 SSL certificate - GeoTrust True BusinessID with EV 26 SSL certificate - GeoTrust QuickSSL 27 SSL certificate - GeoTrustTrueBizID Wildcard 23 SSL certificate - RapidSSL 180 SSL certificate - VeriSign Secure Site 181 SSL certificate - VeriSign Secure Site Pro 182 SSL certificate - VeriSign Secure Site EV 183 SSL certificate - VeriSign Secure Site Pro EV |
| ProductType for Web hosting and Web site creation services | Product type for Web hosting and Web site creation services. Permitted values are: 50 Web hosting account with Access database 51 Web hosting component - 20GB bandwidth 52 Web hosting component - 1GB Web storage 53 Web hosting component - 250MB SQL database storage 54 Web hosting component - 10 POP mailboxes 55 Web hosting setup fee 56 Web hosting bandwidth overage, at per-1GB overage protection rate 57 Web hosting bandwidth overage fee, with upgrade 58 Web hosting component - 100MB POP storage 86 Web Site Creator without Web hosting - Basic 87 Web Site Creator without Web hosting - Full 88 Web Site Creator without Web hosting - eCommerce 90 Web Site Creator added to Web hosting - Basic 91 Web Site Creator added to Web hosting - Full 92 Web Site Creator added to Web hosting |
| ProductType for email services               | Product type for email services. Permitted values are: 35 Additional storage for POP pak - renewal - 512MB/box 36 Additional storage for POP pak - 512MB per mailbox 39 POP mail 10-pak renewal 41 Registration and email forwarding by the .name Registry 42 Email forwarding by the .name Registry 44 .name registration and email forwarding renewal 45 Email forwarding by us 46 Email forwarding by us - renewal |
| ProductType for all other services | Product type for all other services. Permitted values are: 47 URL forwarding 48 URL forwarding - renewal 65 Private label annual subscription 66 Private label annual renewal 72 ID Protect \(Whois Privacy Protection\) 73 ID Protect - renewal 140 Business Listing 141 Business Listing renewal 235 goMobi 236 goMobi free trial 200 RichContent free trial 201 RichContent 215 Instant Reseller 216 Instant Reseller renewal 217 Instant Reseller upgrade                                                                                                                                                                               |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
uid=resellid&pw=resellpw&command=PE_GetRetailPricing
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
uid=resellid&pw=resellpw&command=PE_GetRetailPricing
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
uid=resellid&pw=resellpw&command=PE_GetRetailPricing
&responsetype=text
```
```html
<?xml version="1.0" ?>

  <tld>com</tld>

  <tldid>0</tldid>

  <minimumregistration>1</minimumregistration>

  <registerprice>29.95</registerprice>

  <registerenabled>True</registerenabled>

  <renewprice>29.95</renewprice>

  <renewenabled>True</renewenabled>

  <transferprice>29.95</transferprice>

  <transferenabled>True</transferenabled>

 </tld>

  <tld>net</tld>

  <tldid>1</tldid>

.

.

.

  <price>29.95</price>

  <productenabled>True</productenabled>

  <producttype>45</producttype>

 </product>

  <producttype>47</producttype>

 </pricestructure>

 <Command>PE_GETRETAILPRICING</Command>

 <ErrCount>0</ErrCount>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable>0</IsLockable>

 <IsRealTimeTLD>0</IsRealTimeTLD>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

tld1: com

tldid1: 0

registerprice1: 12.00

resellerpricereg1: 8.95

registerenabled1: true

renewprice1: 18.00

resellerpricerenew1: 8.95

renewenabled1: true

transferprice1: 22.00

resellerpricetran1: 8.95

transferenabled1: true

IDNCategory1:

tld2: net

tldid2: 1

registerprice2: 50.00

resellerpricereg2: 9.00

registerenabled2: true

renewprice2: 10.21

resellerpricerenew2: 9.00

renewenabled2: true

transferprice2: 10.21

resellerpricetran2: 9.00

transferenabled2: true

IDNCategory2:

.

.

.

Command: PE_GETRETAILPRICING

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t

Site: eNom

IsLockable:

IsRealTimeTLD:

TimeDifference: +0.00

ExecTime: 1.703

Done: true

TrackingKey: d65d3050-3f0d-4059-8f75-45a8043fbd3f

RequestDateTime: 2/4/2015 4:37:42 PM
```
.

.

.

ExecTime=1.516

TrackingKey=64db380e-f97c-437f-8550-ca6ec82a83f4

RequestDateTime=2/4/2015 4:38:33 PM
```