PE\GetProductPrice
==================

Retrieve the wholesale cost that this account pays for a product.

Usage
-----

Use this command to retrieve wholesale pricing information for a single product.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/RenewalPricing.asp?tab=1](https://resellertest.enom.com/myaccount/RenewalPricing.asp?tab=1)

Each value in the Your Cost column could be retrieved using one call to the PE\_GetProductPrice command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command =nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter                                             | Status | Description                                                                                                                                                                                                                                                                                                                                                                                                                | Max Size |
| ------------------------------------------------------------------------------------------------------- | ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID                                                   | Required | Account login ID                                                                                                                                                                                                                                                                                                                                                                                                             | 20 |
| PW | Required           | Account password | 20    |
| ResponseType                                              | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                                                                                                                                                                                                                                                                                                                                 | 4 |
| ProductType for domains \(other ProductTypes follow\) | A ProductType is Required  | Product type for domains. Permitted values are: 10 Domain registration 13 DNS hosting 14 DNS hosting renew 16 Domain renewal 17 Domain redemption grace period \(RGP\) 18 Domain Extended RGP \(available at our discretion, and decided by us on a name-by-name basis\) 19 transfer 41 Registration and email forwarding by the .name Registry 44 .name registration and email forwarding renewal | 63    |
| ProductType for SSL certificates \(see adjacent rows for other ProductTypes\)              | Required | Product type for SSL certificates. Permitted values are: 211 SSL certificate - Comodo Essential 212 SSL certificate - Comodo Instant 214 SSL certificate - Comodo Essential Wildcard 213 SSL certificate - Comodo Premium Wildcard 221 SSL certificate - Comodo EV 222 SSL certificate - Comodo EV SGC 20 SSL certificate - GeoTrust QuickSSL Premium 21 SSL certificate - GeoTrust True BusinessID 24 SSL certificate - GeoTrust True BusinessID with EV 26 SSL certificate - GeoTrust QuickSSL 27 SSL certificate - GeoTrustTrueBizID Wildcard 23 SSL certificate - RapidSSL 285 SSL certificate - RapidSSL Wildcard 180 SSL certificate - VeriSign Secure Site 181 SSL certificate - VeriSign Secure Site Pro 182 SSL certificate - VeriSign Secure Site EV 183 SSL certificate - VeriSign Secure Site Pro EV    | 63 |
| TLD | Required           | ProductType is 10, 16, or 19 Top-level domain name \(extension\) | 15    |
| Years                                                  | Optional; default value is 1 | Year bracket for which to retrieve prices. Permitted values are 1, 2, 5, and 10. For example, Years=2 retrieves prices for registrations of 2 to 4 years.                                                                                                                                                                                                                                                                                                                                       | 20 |
| ProductType for Web hosting and Web site creation services \(see adjacent rows for other ProductTypes\) | Required           | Product type for Web hosting and Web site creation services. Permitted values are: 50 Web hosting account with Access database 51 Web hosting component - 20GB bandwidth 52 Web hosting component - 1GB Web storage 53 Web hosting component - 250MB SQL database storage 54 Web hosting component - 10 POP mailboxes 55 Web hosting setup fee 56 Web hosting bandwidth overage, at per-1GB overage protection rate 57 Web hosting bandwidth overage fee, with upgrade 58 Web hosting component - 100MB POP storage 86 Web Site Creator without Web hosting - Basic 87 Web Site Creator without Web hosting - Full 88 Web Site Creator without Web hosting - eCommerce 90 Web Site Creator added to Web hosting - Basic 91 Web Site Creator added to Web hosting - Full 92 Web Site Creator added to Web hosting - eCommerce | 63    |
| ProductType for email services \(see adjacent rows for other ProductTypes\)               | Required | Product type for email services. Permitted values are: 35 Additional storage for POP pak - renewal - 512MB/box 36 Additional storage for POP pak - 512MB per mailbox 39 POP mail 10-pak renewal 41 Registration and email forwarding by the .name Registry 42 Email forwarding by the .name Registry 44 .name registration and email forwarding renewal 45 Email forwarding by us 46 Email forwarding by us - renewal                                                                                                                                                                                                      | 63 |
| ProductType for all other services \(see previous rows for other ProductTypes\) | Required           | Product type for all other services. Permitted values are: 47 URL forwarding 48 URL forwarding - renewal 65 Private label annual subscription 66 Private label annual renewal 72 ID Protect \(Whois Privacy Protection\) 73 ID Protect - renewal 140 Business Listing 141 Business Listing renewal 235 goMobi 236 goMobi free trial 200 RichContent free trial 201 RichContent 215 Instant Reseller 216 Instant Reseller renewal 217 Instant Reseller upgrade | 63    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Price      | Wholesale price that the UID pays for this product |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| ProductEnabled | Enabled status of this product                                  |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GETPRODUCTPRICE&uid=resellid&pw=resellpw
&ProductType=10&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETPRODUCTPRICE&uid=resellid&pw=resellpw
&ProductType=10&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETPRODUCTPRICE&uid=resellid&pw=resellpw
&ProductType=10&tld=com&responsetype=text
```
```
<?xml version="1.0" ?>

 <price>8.95</price>

 <productenabled>True</productenabled>

 </productprice>

 <Command>PE_GETPRODUCTPRICE</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable>True</IsLockable>

 <IsRealTimeTLD>True</IsRealTimeTLD>

 <TimeDifference>+03.00</TimeDifference>

 <ExecTime>0.1015625</ExecTime>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

price: 8.95

productenabled: True

Command: PE_GETPRODUCTPRICE

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

IsLockable: True

IsRealTimeTLD: True

TimeDifference: +08.00

ExecTime: 0.063

Done: true

RequestDateTime: 2/4/2015 4:31:15 PM
```
;Machine is SJL0VWRESELL_T1

Server=SJL0VWRESELL_T1

ExecTime=0.047

RequestDateTime=2/4/2015 4:31:39 PM
```