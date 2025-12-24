PE\SetPricing
=============

Set prices you charge subaccounts for our products and services. You can use this command to set any number of prices.

Usage
-----

Use this command to set the prices you charge subaccounts for our products and services. You can use this command to set any number of prices.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/renewalpricing.asp](https://resellertest.enom.com/myaccount/renewalpricing.asp)

On the my enom > settings page, scroll to the Default Account/Sub-Account Pricing section and click change. You can use PE\_SetPricing to set the price of any or all products.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter                               | Status | Description                                                                                                                                                                                                                                                                                                                                                                                          | Max Size |
| --------------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID                                     | Required | Account login ID                                                                                                                                                                                                                                                                                                                                                                                        | 20 |
| PW | Required                                | Account password | 20    |
| PartyID                                   | If setting subaccount prices, either PartyID or LoginID is Required | Party ID of the subaccount to set pricing for                                                                                                                                                                                                                                                                                                                                                                         | 40 |
| LoginID | If setting subaccount prices, either PartyID or LoginID is Required  | Login ID of the subaccount to set pricing for | 20    |
| UpdateResellerPrice                             | Optional; default is false | If you want to update more than one product price in a reseller subaccount, supply UpdateResellerPrice=true                                                                                                                                                                                                                                                                                                                                          | 5 |
| TLDX X=1, 2, 3, . . . | Required for ProductType 10, 13, 14, 16, 19              | Top-level domain name \(extension\) to set pricing for | 15    |
| YearsX X=1, 2, 3, . . .                         | Required | Year bracket to set prices for. For example, Years=2 sets the annual price for the time bracket from 2 to 4 years; Years=5 sets the annual price for 5 to 9 years. Permitted values are 1, 2, 5, and 10. You can use only one Years parameter per query.                                                                                                                                                                                                                                                                  | 2 |
| ProductTypeX X=1, 2, 3, . . . Domains | A ProductType is Required                       | Product type for domains. Permitted values are: 10 Domain registration 13 DNS hosting 14 DNS hosting renew 16 Domain renewal 17 Domain redemption grace period \(RGP\) 18 Domain Extended RGP \(available at our discretion, and decided by us on a name-by-name basis\) 19 transfer 41 Registration and email forwarding by the .name Registry 44 .name registration and email forwarding renewal | 63    |
| ProductTypeX X=1, 2, 3, . . . SSL                    | Required Product type for SSL certificates. | Permitted values are: 211 SSL certificate - Comodo Essential 212 SSL certificate - Comodo Instant 214 SSL certificate - Comodo Essential Wildcard 213 SSL certificate - Comodo Premium Wildcard 221 SSL certificate - Comodo EV 222 SSL certificate - Comodo EV SGC 20 SSL certificate - GeoTrust QuickSSL Premium 21 SSL certificate - GeoTrust True BusinessID 24 SSL certificate - GeoTrust True BusinessID with EV 26 SSL certificate - GeoTrust QuickSSL 27 SSL certificate - GeoTrustTrueBizID Wildcard 23 SSL certificate - RapidSSL 285 SSL certificate - RapidSSL Wildcard 180 SSL certificate - VeriSign Secure Site 181 SSL certificate - VeriSign Secure Site Pro 182 SSL certificate - VeriSign Secure Site EV 183 SSL certificate - VeriSign Secure Site Pro EV | 63 |
| ProductTypeX X=1, 2, 3, . . . Web hosting and Web site creation services | Required Product type for Web hosting and Web site creation services. | Permitted values are: 50 Web hosting account with Access database 51 Web hosting component - 20GB bandwidth 52 Web hosting component - 1GB Web storage 53 Web hosting component - 250MB SQL database storage 54 Web hosting component - 10 POP mailboxes 55 Web hosting setup fee 56 Web hosting bandwidth overage, at per-1GB overage protection rate 57 Web hosting bandwidth overage fee, with upgrade 58 Web hosting component - 100MB POP storage 86 Web Site Creator without Web hosting - Basic 87 Web Site Creator without Web hosting - Full 88 Web Site Creator without Web hosting - eCommerce 90 Web Site Creator added to Web hosting - Basic 91 Web Site Creator added to Web hosting - Full 92 Web Site Creator added to Web hosting - eCommerce | 63    |
| ProductTypeX X=1, 2, 3, . . . Email services               | Required Product type for email services. | Permitted values are: 35 Additional storage for POP pak - renewal - 512MB/box 36 Additional storage for POP pak - 512MB per mailbox 39 POP mail 10-pak renewal 41 Registration and email forwarding by the .name Registry 42 Email forwarding by the .name Registry 44 .name registration and email forwarding renewal 45 Email forwarding by us 46 Email forwarding by us - renewal                                                                                                                                                                                                  | 63 |
| ProductTypeX X=1, 2, 3, . . . Miscellaneous services | Required Product type for all other services.             | Permitted values are: 47 URL forwarding 48 URL forwarding - renewal 65 Private label annual subscription 66 Private label annual renewal 72 ID Protect \(Whois Privacy Protection\) 73 ID Protect - renewal 140 Business Listing 141 Business Listing renewal 235 goMobi 236 goMobi free trial 200 RichContent free trial 201 RichContent 215 Instant Reseller 216 Instant Reseller renewal 217 Instant Reseller upgrade | 63    |
| EnabledX X=1, 2, 3, . . .                        | Optional | Enable or disable the product type. Set =1 to enable or =0 to disable.                                                                                                                                                                                                                                                                                                                                                            | 1 |
| PriceX X=1, 2, 3, . . . | Optional; use when setting retail prices for this UID         | Set the product price. Use format DD.cc | 5000.00 |
| ResellerPrice                                | Optional; use when setting one wholesale price for one subaccount | Reseller price for this product type. If you use this parameter, you can reset the reseller price for only one product per query. If you use this parameter, use parameters ProductType and Enabled rather than ProductTypeX and EnabledX.                                                                                                                                                                                                                                                                         | 5000.00 |
| RetailPrice | Optional; use when setting one retail price for this UID        | Retail price for this product type. If you use this parameter, you can reset the retail price for only one product per query. If you use this parameter, use parameters ProductType and Enabled rather than ProductTypeX and EnabledX. | 5000.00 |
| RocketPrice                                 | Optional | Registry Rocket price for this product type                                                                                                                                                                                                                                                                                                                                                                          | 5000.00 |
| ResellerKey | Required if setting a Registry Rocket price              | Registry Rocket key | 40    |
| ResponseType                                | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                                                                                                                                                                                                                                                                                                                                                     | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- Either LoginID or PartyID is required. If LoginID is passed in then it will override the PartyID if it is also
- passed in.
- At least one Price must be supplied.

Example
-------

The following query requests that the retail price for registering .org names be set at $44.00 and for renewing .org names $45.00, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_SetPricing&uid=resellid&pw=resellpw
&ProductType1=10&TLD1=org&Enabled1=1&Price1=44
&ProductType2=16&TLD2=org&Enabled2=1&Price2=45
&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=PE_SetPricing&uid=resellid&pw=resellpw
&ProductType1=10&TLD1=org&Enabled1=1&Price1=44
&ProductType2=16&TLD2=org&Enabled2=1&Price2=45
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_SetPricing&uid=resellid&pw=resellpw
&ProductType1=10&TLD1=org&Enabled1=1&Price1=44
&ProductType2=16&TLD2=org&Enabled2=1&Price2=45
&responsetype=text
```
In the response, the Status Successful and Errcount 0 confirm that the query successfully reset 2 prices:

```
<?xml version="1.0" ?>
<interface-response>
<Count>2</Count>
<Status>Successful</Status>
<Command>PE_SETPRICING</Command>
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
Count: 2
Status: Successful
Command: PE_SETPRICING
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
ExecTime: 1.063
Done: true
TrackingKey: 19b305ec-5b59-4696-8ead-a0c684312050
RequestDateTime: 2/4/2015 4:44:47 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Count=2
Status=Successful
Command=PE_SETPRICING
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
TrackingKey=d0fd143d-5877-4bc2-b766-230d824ed2d7
RequestDateTime=2/4/2015 4:45:08 PM
```
Related Commands
----------------

AuthorizeTLD

CommissionAccount

GetBalance

GetTLDList

GetWebHostingAll

PE\_GetCustomerPricing

PE\_GetDomainPricing

PE\_GetPOPPrice

PE\_GetProductPrice

PE\_GetResellerPrice

PE\_GetRetailPrice

PE\_GetRocketPrice

PE\_GetTLDID

RemoveTLD

SetResellerServicesPricing

SetResellerTLDPricing

UpdateAccountPricing