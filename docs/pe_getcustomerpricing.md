PE\GetCustomerPricing
=====================

Get retail prices that this account charges to consumers. These are also the prices that this account’s subaccounts will be charged, unless you change them for each individual subaccount.

Usage
-----

Use this command to retrieve prices that you have changed from their default values.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/SubAccount-List.asp](https://resellertest.enom.com/resellers/SubAccount-List.asp)

If you click a Login ID of a retail subaccount, the link calls the PE\_GetCustomerPricing command. If you click a Login ID of a reseller subaccount, the link calls the PE\_GetResellerPrice command.

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
| Input Parameter | Status | Description                                | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                              | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                                                                                                                                                                                                                                                                                                                                                                         |
| ErrCount                          | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\) |
| ErrX | Error messages explaining the failure. These can be presented as-is back to the client.                                                                                                                                                                                                                                                                                                                                                                        |
| Done                            | True indicates this entire response has reached you successfully. |
| CPID | Pricing engine index, for use by eNom’s database. ProductType for domains Product type for domains. Permitted values are: 10 Domain registration 13 DNS hosting 14 DNS hosting renew 16 Domain renewal 17 Domain redemption grace period \(RGP\) 18 Domain Extended RGP \(available at our discretion, and decided by us on a name-by-name basis\) 19 transfer 41 Registration and email forwarding by the .name Registry 44 .name registration and email forwarding renewal                                                                                                                                                                            |
| ProductType for SSL certificates              | Product type for SSL certificates. Permitted values are: 211 SSL certificate - Comodo Essential 212 SSL certificate - Comodo Instant 214 SSL certificate - Comodo Essential Wildcard 213 SSL certificate - Comodo Premium Wildcard 221 SSL certificate - Comodo EV 222 SSL certificate - Comodo EV SGC 20 SSL certificate - GeoTrust QuickSSL Premium 21 SSL certificate - GeoTrust True BusinessID 24 SSL certificate - GeoTrust True BusinessID with EV 26 SSL certificate - GeoTrust QuickSSL 27 SSL certificate - GeoTrustTrueBizID Wildcard 23 SSL certificate - RapidSSL 285 SSL certificate - RapidSSL Wildcard 180 SSL certificate - VeriSign Secure Site 181 SSL certificate - VeriSign Secure Site Pro 182 SSL certificate - VeriSign Secure Site EV 183 SSL certificate - VeriSign Secure Site Pro EV |
| ProductType for Web hosting and Web site creation services | Product type for Web hosting and Web site creation services. Permitted values are: 50 Web hosting account with Access database 51 Web hosting component - 20GB bandwidth 52 Web hosting component - 1GB Web storage 53 Web hosting component - 250MB SQL database storage 54 Web hosting component - 10 POP mailboxes 55 Web hosting setup fee 56 Web hosting bandwidth overage, at per-1GB overage protection rate 57 Web hosting bandwidth overage fee, with upgrade 58 Web hosting component - 100MB POP storage 86 Web Site Creator without Web hosting - Basic 87 Web Site Creator without Web hosting - Full 88 Web Site Creator without Web hosting - eCommerce 90 Web Site Creator added to Web hosting - Basic 91 Web Site Creator added to Web hosting - Full 92 Web Site Creator added to Web hosting - eCommerce |
| ProductType for email services               | Product type for email services. Permitted values are: 35 Additional storage for POP pak - renewal - 512MB/box 36 Additional storage for POP pak - 512MB per mailbox 39 POP mail 10-pak renewal 41 Registration and email forwarding by the .name Registry 42 Email forwarding by the .name Registry 44 .name registration and email forwarding renewal 45 Email forwarding by us 46 Email forwarding by us - renewal |
| ProductType for all other services | Product type for all other services. Permitted values are: 47 URL forwarding 48 URL forwarding - renewal 65 Private label annual subscription 66 Private label annual renewal 72 ID Protect \(Whois Privacy Protection\) 73 ID Protect - renewal 140 Business Listing 141 Business Listing renewal 235 goMobi 236 goMobi free trial 190 200 RichContent free trial 201 RichContent 215 Instant Reseller 216 Instant Reseller renewal 217 Instant Reseller upgrade                                                                                                                                                                                   |
| ProductDescription                     | Product description. Values are listed under ProductType, above. |
| TLDID | Top-level-domain ID number.                                                                                                                                                                                                                                                                                                                                                                                                      |
| TLD                            | Top-level domain name. |
| RetailPrice | Retail price. Response contains a value only if price overrides the default value                                                                                                                                                                                                                                                                                                                                                                            |
| ResellerPrice                       | Reseller price. |
| RocketPrice | Registry Rocket price. Response contains a value only if price overrides the default value.                                                                                                                                                                                                                                                                                                                                                                       |
| Enabled                          | Enabled state for this TLD. Options are True or False. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GETCUSTOMERPRICING&uid=resellid
&pw=resellpw&ProductType=19&tld=org&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETCUSTOMERPRICING&uid=resellid
&pw=resellpw&ProductType=19&tld=org&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETCUSTOMERPRICING&uid=resellid
&pw=resellpw&ProductType=19&tld=org&responsetype=text
```
```
<interface-response>

 <productstructure>
 <product>
  <cpid>2490167</cpid>
  <producttype>10</producttype>
  <productdescription>Register</productdescription>
  <tldid>0</tldid>
  <tld>com</tld>
  <retailprice>15</retailprice>
  <resellerprice/>
  <rocketprice/>
  <enabled>True</enabled>
  <minimumregistration>1</minimumregistration>
 </product>
 <product>
  <cpid>2490168</cpid>
  <producttype>10</producttype>
  <productdescription>Register</productdescription>
  <tldid>0</tldid>
  <tld>com</tld>
  <retailprice>25.95</retailprice>
  <resellerprice>8.95</resellerprice>
  <rocketprice/>
  <enabled>True</enabled>
  <minimumregistration>1</minimumregistration>
 </product>
.
.
.
 </productstructure>
 <Command>PE_GETCUSTOMERPRICING</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+07.00</TimeDifference>
 <ExecTime>0.438</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 11:07:40 PM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

cpid1: 3570620

producttype1: 10

productdescription1: Register

tldid1: 425

tld1: bike

retailprice1:

resellerprice1: 23

rocketprice1:

enabled1: True

minimumregistration1: 1

cpid2: 3251440

producttype2: 10

productdescription2: Register

tldid2: 6

tld2: biz

retailprice2: 10.21

resellerprice2:

rocketprice2:

enabled2: True

minimumregistration2: 1

.
.
.
Command: PE_GETCUSTOMERPRICING

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

IsLockable: True

IsRealTimeTLD: True

TimeDifference: +07.00

ExecTime: 1.672

Done: true

RequestDateTime: 2/4/2015 4:10:17 PM
```
ExecTime=0.609

RequestDateTime=2/4/2015 4:12:41 PM
```