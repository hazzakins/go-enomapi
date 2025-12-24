UpdateAccountPricing
====================

Set wholesale prices for a subaccount.

Usage
-----

Use this command to set the wholesale prices that one subaccount pays you for each product.

To set retail prices for consumers who buy directly from this account, use PE\_SetPricing.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/subaccount-list.asp](https://resellertest.enom.com/resellers/subaccount-list.asp)

Click any subaccount’s LoginID to see the wholesale prices you charge them; the values can be reset with UpdateAccountPricing.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The subaccount named with LoginID or Account parameters must be a subaccount of the account named in the uid and pw parameters.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter               | Status | Description                                                                                                                                                            | Max Size |
| -------------------------------------------- | --------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID                     | Required | Account login ID                                                                                                                                                         | 20 |
| PW | Required             | Account password | 20    |
| LoginID                   | Optional | Subaccount login ID \(use GetAccountInfo subaccount login ID\)                                                                                                                                  | 20 |
| Account | Required             | Subaccount ID number \(use GetAccountInfo subaccount ID number; format is NNN-aa-NNNN\) | 11    |
| AcctType                   | Required | Type of this subaccount. Permitted values are reseller and retail. To retrieve subaccount types, use the                                                                                                            | 8 |
| TLDprice | Optional name, in DD.cc format. | For example, comprice=44.95 sets the price of registering a .com name to $US44.95. | 5000.00 |
| TLDrenew                   | Optional | Price for renewing a domain name, in DD.cc format                                                                                                                                        | 5000.00 |
| TLDtransfer | Optional             | Price for transferring a domain name and extending its expiration date by one year, in DD.cc format | 5000.00 |
| DNSHosting                  | Optional | Price for one year of DNS hosting, in DD.cc format                                                                                                                                        | 5000.00 |
| DotNameBundle | Optional             | email service from the .name Registry, in DD.cc format | 5000.00 |
| WPPS                     | Optional | Price for one year of Whois Privacy Protection Service, in DD.cc format                                                                                                                             | 5000.00 |
| Certificate-Comodo-Essential | Optional             | Essential certificate, in DD.cc format | 5000.00 |
| Certificate-Comodo-Instant         | Optional | Price for a one-year Comodo Instant certificate, in DD.cc format                                                                                                                                | 5000.00 |
| Certificate-Comodo-Essential-Wildcard | Optional             | Price for a one-year Comodo Essential Wildcard certificate, in DD.cc format | 5000.00 |
| Certificate-Comodo-Premium-Wildcard    | Optional | Price for a one-year Comodo Premium Wildcard certificate, in DD.cc format                                                                                                                            | 5000.00 |
| Certificate-Comodo-EV | Optional             | Price for a one-year Comodo EV certificate, in DD.cc format | 5000.00 |
| Certificate-Comodo-EV-SGC         | Optional | Price for a one-year Comodo EV SGC \(server-gated cryptography\) certificate, in DD.cc format                                                                                                                 | 5000.00 |
| Certificate-GeoTrust-QuickSSL | Optional             | Price for a one-year GeoTrust QuickSSL certificate, in DD.cc format | 5000.00 |
| Certificate-GeoTrust-QuickSSL-Premium   | Optional | Price for a one-year GeoTrust QuickSSL Premium SSL certificate, in DD.cc format                                                                                                                         | 5000.00 |
| Certificate-GeoTrust-TrueBizID | Optional             | Price for a one-year GeoTrust True BusinessID SSL certificate, in DD.cc format | 5000.00 |
| Certificate-GeoTrust-TrueBizID-Wildcard  | Optional | Price for a one-year GeoTrust True BusinessID Wildcard SSL certificate, in DD.cc format                                                                                                                     | 5000.00 |
| Certificate-GeoTrust-TrueBizID-EV | Optional             | Price for a one-year GeoTrust True BusinessID with EV SSL certificate, in DD.cc format | 5000.00 |
| Certificate-RapidSSL-RapidSSL       | Optional | Price for a one-year RapidSSL certificate, in DD.cc format                                                                                                                                   | 5000.00 |
| Certificate-RapidSSL-RapidSSL-Wildcard | Optional             | Price for a one-year RapidSSL Wildcard certificate, in DD.cc format | 5000.00 |
| Certificate-VeriSign-Secure-Site     | Optional | Price for a one-year VeriSign Secure Site SSL certificate, in DD.cc format                                                                                                                           | 5000.00 |
| Certificate-VeriSign-Secure-Site-Pro | Optional             | Price for a one-year VeriSign Secure Site Pro SSL certificate, in DD.cc format | 5000.00 |
| Certificate-VeriSign-Secure-Site-EV   | Optional | Price for a one-year VeriSign Secure Site EV SSL certificate, in DD.cc format                                                                                                                          | 5000.00 |
| Certificate-VeriSign-Secure-Site-Pro-EV | Optional             | Price for a one-year VeriSign Secure Site Pro EV SSL certificate, in DD.cc format | 5000.00 |
| Seal-VeriSign-Trust-Seal         | Optional | Price per month for VeriSign Trust Seal, in DD.cc format                                                                                                                                     | 5000.00 |
| WBL | Optional             | Price per year for Business Listing, in DD.cc format | 5000.00 |
| RichContent                 | Optional | Price per month for RichContent, regardless of the billing cycle the customer chooses. Use this single value if you want a single price point for all billing cycles, or use all four other parameters if you want to offer your subaccounts a price reduction for longer billing cycles. Permitted values are in DD.cc format | 5000.00 |
| RichContent-Mo1 | Optional             | Price per month for RichContent, when the customer chooses a 1-month billing cycle, in DD.cc format | 5000.00 |
| RichContent-Mo3               | Optional | Price per month for RichContent, when the customer chooses a 3- month billing cycle, in DD.cc format                                                                                                              | 5000.00 |
| RichContent-Mo6 | Optional             | Price per month for RichContent, when the customer chooses a 6-month billing cycle, in DD.cc format | 5000.00 |
| RichContent-Mo12              | Optional | Price per month for RichContent, when the customer chooses a 12-month billing cycle, in DD.cc format                                                                                                              | 5000.00 |
| POP3 | Optional             | Price for a 10-pak of POP3 mailboxes, in DD.cc format | 5000.00 |
| 45\_valueadd                 | Optional | Price of email forwarding, in DD.cc format                                                                                                                                            | 5000.00 |
| 47\_valueadd | Optional             | Price of URL forwarding, in DD.cc format | 5000.00 |
| InstantReseller               | Optional | Price for a 1-year subscription to Instant Reseller. Use this single value if you want a single price point for all purchases, or use all three yearly parameters, below, if you want to offer your subaccounts a price reduction for longer subscriptions. Permitted values are in DD.cc format               | 5000.00 |
| InstantReseller-Yr1 | Optional             | Price for a 1-year subscription to Instant Reseller, in DD.cc format | 5000.00 |
| InstantReseller-Yr2             | Optional | Price per year for a 2-year subscription to Instant Reseller, in DD.cc format                                                                                                                          | 5000.00 |
| InstantReseller-Yr3 |                  | Optional Price per year for a 3-year subscription to Instant Reseller, in DD.cc format | 5000.00 |
| WSCBasic                   | Optional | Price for WebSite Creator Basic package, in DD.cc format                                                                                                                                     | 5000.00 |
| WSCFull | Optional             | Price for WebSite Creator Full package, in DD.cc format | 5000.00 |
| WSCEcommerce                 | Optional | Price for WebSite Creator eCommerce package, in DD.cc format                                                                                                                                   | 5000.00 |
| goMobi-Mo1 | Optional             | Price per month for goMobi, when customer chooses a 1- month billing cycle, in DD.cc format | 5000.00 |
| goMobi-Mo12                 | Optional | Price per month for goMobi, when customer chooses a 12- month billing cycle, in DD.cc format                                                                                                                  | 5000.00 |
| ResponseType | Optional Format of response.   | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Account     | Subaccount ID |
| AccountType | Subaccount type                                         |
| Count      | Number of prices modified |
| Status | Success status for this query                                  |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

The following query sets prices for several products, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
uid=resellid&pw=resellpw&Account=154-dz-5567
&AcctType=reseller&command=UPDATEACCOUNTPRICING
&comprice=14&comrenew=15&comtransfer=16&pop3=29.95
&dotnamebundle=20.95&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
uid=resellid&pw=resellpw&Account=154-dz-5567
&AcctType=reseller&command=UPDATEACCOUNTPRICING
&comprice=14&comrenew=15&comtransfer=16&pop3=29.95
&dotnamebundle=20.95&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
uid=resellid&pw=resellpw&Account=154-dz-5567
&AcctType=reseller&command=UPDATEACCOUNTPRICING
&comprice=14&comrenew=15&comtransfer=16&pop3=29.95
&dotnamebundle=20.95&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
<Set>
 <Account>154-dz-5567</Account>
 <AcctType>reseller</AcctType>
 <Count>5</Count>
 <Status>Successful</Status>
</Set>
<Command>UPDATEACCOUNTPRICING</Command>
<ErrCount>0</ErrCount>
<Server>RESELLERTEST</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<Done>true</Done>
<debug>
 <![CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Account: 154-dz-5567
AcctType: reseller
Command: UPDATEACCOUNTPRICING
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.125
Done: true
RequestDateTime: 2/9/2015 1:51:06 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Account=154-dz-5567
AcctType=reseller
Command=UPDATEACCOUNTPRICING
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
RequestDateTime=2/9/2015 1:51:30 PM
```
Related Commands
----------------

PE\_GetCustomerPricing

PE\_GetDomainPricing

PE\_GetResellerPrice

PE\_GetRetailPrice

PE\_GetRetailPricing

PE\_GetRocketPrice

PE\_SetPricing

SetResellerServicesPricing

SetResellerTLDPricing

WSC\_GetPricing