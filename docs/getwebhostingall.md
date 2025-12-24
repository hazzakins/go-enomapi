GetWebHostingAll
================

Retrieve reseller keys for Registry Rocket accounts.

Usage
-----

Use this command to retrieve reseller keys for Registry Rocket accounts.

Do not use this command to retrieve prices, because the prices returned here may be obsolete. Instead, use one of the PE\_Get\* pricing commands.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/Settings.asp](https://resellertest.enom.com/myaccount/Settings.asp)

In the Registry Rocket Settings box at the bottom of the page, the ResellerKey line posts the return value from the GetWebHostingAll command.

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

| Output Parameter | Description                                                          |
| ---------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                   |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                   |
| Done       | True indicates this entire response has reached you successfully. |
| ResellerKey | Registry Rocket reseller key                                                 |
| HostPrice    | Price when this Registry Rocket account was created; likely to be obsolete. For current pricing, use the PE\_Get\*commands. |
| CCPrice | Price when this Registry Rocket account was created; likely to be obsolete. For current pricing, use the PE\_Get\*commands. |
| InfoPrice    | Price when this Registry Rocket account was created; likely to be obsolete. For current pricing, use the PE\_Get\*commands. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves a list of the reseller keys for account resellid, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetWebhostingAll&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetWebhostingAll&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetWebhostingAll&uid=resellid
&pw=resellpw&responsetype=text
```
In the response, the results for ResellerKey and the ErrCount of 0 confirm that the query was successful:

<ResellerKey>20B0FF00-18C2-4425-9198-5</ResellerKey>

<HostPrice>11.21</HostPrice>

<CCPrice>27.69</CCPrice>

<InfoPrice>11.21</InfoPrice>

<BizPrice>7.14</BizPrice>

<TvPrice>43.14</TvPrice>

<WsPrice>16.36</WsPrice>

<BzPrice>27.69</BzPrice>

<NuPrice>27.69</NuPrice>

<UsPrice>11.21</UsPrice>

<PictureURL/>

<ContactEmail/>

<CompanyName>RRTest</CompanyName>

<Referrer/>

<NameMyPhone>False</NameMyPhone>

<NameMyComputer>False</NameMyComputer>

<NameMyMap>False</NameMyMap>

<SiteBuilder>False</SiteBuilder>

<StyleSheetNum>1</StyleSheetNum>

</RocketLink>

<success>False</success>

<Command>GETWEBHOSTINGALL</Command>

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

<ExecTime>0.047</ExecTime>

<Done>true</Done>

<debug/>

<TrackingKey>75373637-e0d2-4b8a-b56d-55765f16d96a</TrackingKey>

<RequestDateTime>12/9/2011 3:08:37 AM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

ResellerKey0: 20B0FF00-18C2-4425-9198-5

HostPrice0: 11.21

CCPrice0: 27.69

InfoPrice0: 11.21

BizPrice0: 7.14

TvPrice0: 43.14

WsPrice0: 16.36

BzPrice0: 27.69

NuPrice0: 27.69

UsPrice0: 11.21

PictureURL0:

ContactEmail0:

CompanyName0: RRTest

Referrer0:

NameMyPhone0: False

NameMyComputer0: False

NameMyMap0: False

SiteBuilder0: False

StyleSheetNum0: 1

success: False

Command: GETWEBHOSTINGALL

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.031

Done: true

TrackingKey: df057da9-adb6-4ff2-9732-bea505633b60

RequestDateTime: 2/4/2015 12:58:30 PM
```
ExecTime=0.016

TrackingKey=4c4d9556-cf53-4b36-8ba4-339a35acabc3

RequestDateTime=2/4/2015 12:58:52 PM
```