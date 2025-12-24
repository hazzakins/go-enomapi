GetBalance
==========

Get account balances.

Usage
-----

Use this command to return the current balance on an account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/SubAccount.asp](https://resellertest.enom.com/resellers/SubAccount.asp)

The Access link calls the GetBalance command.

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
| Input Parameter | Status | Description                                 | Max Size |
| --------------- | -------- | ---------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                               | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                        |
| ---------------- | ------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                  |
| ErrCount     | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\) |
| ErrX | Error messages explaining the failure. These can be presented as-is back to the client. |
| Done       | True indicates this entire response has reached you successfully. |
| Reseller | Returns 1 if this is a reseller account, 0 otherwise                    |
| Balance     | Current account balance |
| AvailableBalance | Current available balance                                 |
| Price      | Default price for names |
| COMPrice | Price for .com                                      |
| NETPrice     | Price for .net |
| ORGPrice | Price for .org                                      |
| CCPrice     | Price for .cc |
| TVPrice | Price for .tv                                       |
| DomainCount   | Current count of domains in the account |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

This command only returns the price for original purchase of domains, not tranfers or renewals. For acomplete list of prices, use PE\_GetRetailPricing or PE\_GetResellerPrice.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetBalance&uid=resellid&pw=resellpw
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetBalance&uid=resellid&pw=resellpw
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetBalance&uid=resellid&pw=resellpw
&responsetype=text
```
<Reseller>1</Reseller>

<Balance>2,261.65</Balance>

<AvailableBalance>1,939.50</AvailableBalance>

<Price>8.95</Price>

<COMPrice>8.95</COMPrice>

<NETPrice>8.95</NETPrice>

<ORGPrice>8.95</ORGPrice>

<CCPrice>24.95</CCPrice>

<TVPrice>39.95</TVPrice>

<BZPrice>24.95</BZPrice>

<NUPrice>24.95</NUPrice>

<DomainCount>2943</DomainCount>

<Command>GETBALANCE</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod>2</MinPeriod>

<MaxPeriod>10</MaxPeriod>

<Server>SJL21WRESELLT01</Server>

<Site>eNom</Site>

<IsLockable>True</IsLockable>

<IsRealTimeTLD>True</IsRealTimeTLD>

<TimeDifference>+03.00</TimeDifference>

<ExecTime>0.312</ExecTime>

<Done>true</Done>

<RequestDateTime>12/13/2011 2:59:04 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```bash
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Reseller: 1

Balance: -1,832.23

AvailableBalance: 6,193.37

Price: 8.95

COMPrice: 8.95

NETPrice: 9.00

ORGPrice: 9.00

CCPrice: 24.95

TVPrice: 39.95

BZPrice: 24.95

NUPrice: 24.95

DomainCount: 2692

Command: GETBALANCE

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

IsLockable: False

IsRealTimeTLD: True

TimeDifference: +03.00

ExecTime: 0.156

Done: true

RequestDateTime: 2/3/2015 4:41:26 PM
```
;Machine is SJL0VWRESELL_T

Server=SJL0VWRESELL_T

ExecTime=0.141

RequestDateTime=2/3/2015 4:41:51 PM
```