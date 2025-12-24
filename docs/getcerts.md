GetCerts
========

Retrieve a list of the SSL certificates that we offer.

Usage
-----

Use this command to retrieve a list of the types of SSL certificates that we offer.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/secure/PurchaseCerts.asp?](https://resellertest.enom.com/secure/PurchaseCerts.asp?)

The table that lists the certs we offer could be populated by the GetCerts command. Our page uses GetCerts to retrieve prices and hard codes the cert names, but GetCerts does return cert names as well.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                      |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.               |
| Done       | True indicates this entire response has reached you successfully. |
| Reseller | Is this a reseller account.                                              |
| ProdCode     | Product text ID |
| ProdType | Product number ID                                                   |
| ProdDesc     | Product description |
| QtyDesc | Units in which this product is sold                                          |
| RetailPrice   | Retail selling price charged to customers |
| ResellerPrice | Wholesale price this account pays                                           |
| Enabled     | Is this product enabled for selling. |
| Years | Quantities available. For example, Years=3 indicates that a customer can buy a 3-year registration of this product. |
| CertCount    | Number of SSL certificate products listed in this return |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves a list of the SSL certificate products we offer, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetCerts&UID=resellid&PW=resellpw
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetCerts&UID=resellid&PW=resellpw
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetCerts&UID=resellid&PW=resellpw
&ResponseType=text
```
In the response, a list of certificates and an ErrCount value 0 confirm that the query was successful:

<Reseller>True</Reseller>

<ProdCode>certificate-rapidssl-rapidssl</ProdCode>

<ProdType>23</ProdType>

<ProdDesc>SSL Certificate - RapidSSL</ProdDesc>

<QtyDesc>Year(s)</QtyDesc>

<RetailPrice>$19.95</RetailPrice>

<ResellerPrice>$9.95</ResellerPrice>

<Enabled>True</Enabled>

<Years>1</Years>

<Years>2</Years>

<Years>3</Years>

<Years>4</Years>

<Years>5</Years>

</Cert1>

<ProdCode>certificate-geotrust-quickssl</ProdCode>

<ProdType>26</ProdType>

<ProdDesc>SSL Certificate - GeoTrust QuickSSL</ProdDesc>

<RetailPrice>$125.00</RetailPrice>

<ResellerPrice>$69.00</ResellerPrice>

</Cert2>

.

.

.

</Certs>

<CertCount>16</CertCount>

<Command>GETCERTS</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod/>

<MaxPeriod>10</MaxPeriod>

<Server>sjl21wresellt01</Server>

<Site>eNom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>0.500</ExecTime>

<Done>true</Done>

<debug/>

<TrackingKey>c0a2fd36-f8f8-451b-9eb4-8b5b406b69c5</TrackingKey>

<RequestDateTime>12/8/2011 3:47:44 AM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Reseller: True

ProdCode: certificate-rapidssl-rapidssl

ProdType: 23

ProdDesc: SSL Certificate - RapidSSL

QtyDesc: Year(s)

RetailPrice: $15.95

ResellerPrice: $9.95

Enabled: True

Years: 1

Years: 2

Years: 3

Years: 4

ProdCode: certificate-geotrust-quickssl

ProdType: 26

ProdDesc: SSL Certificate - GeoTrust QuickSSL

RetailPrice: $89.00

ResellerPrice: $69.00

ProdCode: certificate-geotrust-quickssl-premium

.

.

.

CertCount: 24

Command: GETCERTS

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t1

Site: eNom

TimeDifference: +0.00

ExecTime: 2.234

Done: true

TrackingKey: c4a89eec-0a16-44a3-a942-0659d8092633

RequestDateTime: 2/3/2015 4:49:58 PM
```
;Machine is SJL0VWRESELL_T

.

.

.

Server=sjl0vwresell_t

ExecTime=2.250

TrackingKey=ea219140-8b5a-48f1-b3b8-985dded3877c

RequestDateTime=2/3/2015 4:51:07 PM
```