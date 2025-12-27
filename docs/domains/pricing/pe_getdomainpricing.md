PE\GetDomainPricing
===================

Get the retail pricing that this account charges for registrations, renewals, and transfers, by top-level domain.

Usage
-----

Use this command when you want a list of retail prices for one top-level domain, for registrations, renewals, and transfers. To get a list of retail prices for all products offered by this account, use PE\_GetRetailPricing.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=850-tn-1053](https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=850-tn-1053)

PE\_GetDomainPricing is not implemented on enom.com. The sub-account configuration page displays similar information, but for all top-level domains.

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
| Input Parameter | Status | Description                                        | Max Size |
| --------------- | ---------------------- | ----------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                     | 20 |
| PW | Required        | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML.        | 4 |
| UseQtyEngine | Optional; default is 0 | Return prices for multiple-year registrations. Permitted values are 0 and 1. | 1    |
| Years      | Optional; default is 1 | Number of years for multiple-year registrations. Permitted values are 1, 2, 5, and 10. | 2 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                        |
| ---------------- | ------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                  |
| ErrCount     | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\) |
| ErrX | Error messages explaining the failure. These can be presented as-is back to the client. |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GETDOMAINPRICING&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETDOMAINPRICING&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETDOMAINPRICING&uid=resellid
&pw=resellpw&responsetype=text
```
<DisplayPriceIncrease>0</DisplayPriceIncrease>

<tld>com</tld>

<tldid>0</tldid>

<registerprice>8.95</registerprice>

<resellerpricereg>8.95</resellerpricereg>

<registerenabled>True</registerenabled>

<renewprice>8.95</renewprice>

<resellerpricerenew>8.95</resellerpricerenew>

<renewenabled>True</renewenabled>

<transferprice>8.95</transferprice>

<resellerpricetran>8.95</resellerpricetran>

<transferenabled>True</transferenabled>

</product>

<tld>net</tld>

<tldid>1</tldid>

.

.

.

<count>96</count>

</pricestructure>

<Command>PE_GETDOMAINPRICING</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod/>

<MaxPeriod>10</MaxPeriod>

<Server>SJL21WRESELLT01</Server>

<Site>eNom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>3.470</ExecTime>

<Done>true</Done>

<RequestDateTime>12/9/2011 3:40:47 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

DisplayPriceIncrease: 0

tld1: com

tldid1: 0

registerprice1: 8.95

resellerpricereg1: 8.95

registerenabled1: True

renewprice1: 8.95

resellerpricerenew1: 8.95

renewenabled1: True

transferprice1: 8.95

resellerpricetran1: 8.95

transferenabled1: True

rgpprice1: 160

resellerpricergp1: 160

rgpenabled1: True

tld2: net

tldid2: 1

registerprice2: 9

resellerpricereg2: 9

registerenabled2: True

renewprice2: 9

resellerpricerenew2: 9

renewenabled2: True

transferprice2: 9

resellerpricetran2: 9

transferenabled2: True

rgpprice2: 160

resellerpricergp2: 160

rgpenabled2: True

.

.

.

count: 358

Command: PE_GETDOMAINPRICING

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 58.301

Done: true

RequestDateTime: 2/4/2015 4:21:08 PM
```
.

.

.

ExecTime=29.758

RequestDateTime=2/4/2015 4:22:33 PM
```