PE\GetEapPricing
================

Get Early Access Program \(EAP\) pricing and phase information for a list of TLDs, including if there are pending orders for the domain for the given EAP day.

Usage
-----

Use this command to retrieve prices for a list of domains in EAP status.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This is used to show the pre-registration options for a TLD on the queue search results.

[http://resellertest.enom.com/tld-queue/pages/search-results.aspx?sld=trading&tld=cards](http://resellertest.enom.com/tld-queue/pages/search-results.aspx?sld=trading&tld=cards)

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

At least one domain must be passed in

The TLD must be a TLD sold through Donuts Inc.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                      | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                    | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML.       | 63 |
| SLDX | Required | Second-level domain name \(for example, enom in enom.com\) X = integer, 1, 2, 3, etc | 15    |
| TLDX      | Required | Top-level domain name \(extension\) X = integer, 1, 2, 3, etc             | 2 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                            |
| ---------------- | -------------------------------------------------------------------------------------------------- |
| Domain | Domain Name                                            |
| SLD       | Second level domain name. |
| TLD | Top-level domain name.                                      |
| TLDID      | Top-level-domain ID number. |
| Phase | The name of the phase                                       |
| Day       | EAP Day number for that phase. Will only be present for EAP phase items. |
| Price | The calling accounts price for the item.                             |
| RetailPrice   | Retail price. Response contains a value only if the api is called as a reseller. |
| ReleaseDate | Localized time that the phase gets released. Time is in PST.                   |
| ReleaseDate-UTC | Release Date converted to UTC |
| OrderCount | Order\(s\) exist for this SLD.TLD in the given EAP day? Will not be returned for non EAP phases. |
| Valid      | values are 0 and 1. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.     |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GetEapPricing& uid=resellid&pw=resellpw
&sld1=trading&tld1=cards&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GetEapPricing& uid=resellid&pw=resellpw
&sld1=trading&tld1=cards&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GetEapPricing& uid=resellid&pw=resellpw
&sld1=trading&tld1=cards&responsetype=text
```
<domain domain="trading.cards" sld="trading" tld="cards" tldid="488">

<phase>Pre-registration</phase>

<price>930.00</price>

<retailprice>1,208.00</retailprice>

<releasedate>2014-06-18T09:00:00</releasedate>

<releasedate-utc>2014-06-18T16:00:00</releasedate-utc>

<ordercount/>

</pricing>

<phase>EAP</phase>

<day>1</day>

<price>11,920.00</price>

<retailprice>14,946.00</retailprice>

<releasedate>2014-06-11T09:00:00</releasedate>

<releasedate-utc>2014-06-11T16:00:00</releasedate-utc>

<ordercount>0</ordercount>

<day>2</day>

<price>3,670.00</price>

<retailprice>4,634.00</retailprice>

<releasedate>2014-06-12T09:00:00</releasedate>

<releasedate-utc>2014-06-12T16:00:00</releasedate-utc>

<day>3</day>

<price>2,013.00</price>

<retailprice>2,563.00</retailprice>

<releasedate>2014-06-13T09:00:00</releasedate>

<releasedate-utc>2014-06-13T16:00:00</releasedate-utc>

<day>4</day>

<price>1,495.00</price>

<retailprice>1,944.00</retailprice>

<releasedate>2014-06-14T09:00:00</releasedate>

<releasedate-utc>2014-06-14T16:00:00</releasedate-utc>

<day>5</day>

<price>1,030.00</price>

<retailprice>1,334.00</retailprice>

<releasedate>2014-06-15T09:00:00</releasedate>

<releasedate-utc>2014-06-15T16:00:00</releasedate-utc>

<day>6</day>

<releasedate>2014-06-16T09:00:00</releasedate>

<releasedate-utc>2014-06-16T16:00:00</releasedate-utc>

<day>7</day>

<releasedate>2014-06-17T09:00:00</releasedate>

<releasedate-utc>2014-06-17T16:00:00</releasedate-utc>

</domain>

</domains>

<Command>PE_GETEAPPRICING</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod>1</MinPeriod>

<MaxPeriod>10</MaxPeriod>

<Server>SJL21WRESELLT01</Server>

<Site>eNom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>0.453</ExecTime>

<RequestDateTime>5/6/2014 11:40:48 AM</RequestDateTime>

<debug/>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

domain1: trading.cards

sld1: trading

tld1: cards

tldid1: 488

domain1phase1: GA

domain1price1: 920.00

domain1retailprice1: 5,520.00

domain1releasedate1: 2014-06-18T09:00:00

domain1releasedate-utc1: 2014-06-18T16:00:00

domain1ordercount1: 0

PricingCount1: 1

DomainCount: 1

Command: PE_GETEAPPRICING

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t

Site: eNom

TimeDifference: +0.00

ExecTime: 0.406

Done: true

TrackingKey: af5c1dbe-d063-4f51-9ce6-09dc6fe89bcb

RequestDateTime: 2/4/2015 4:24:29 PM
```
ExecTime=0.094

TrackingKey=2533dce4-fdc1-4763-b6ce-910fd2bcf31f

RequestDateTime=2/4/2015 4:25:03 PM
```