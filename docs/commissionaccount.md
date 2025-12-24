CommissionAccount
=================

Returns the current commission balance for an account.

Usage
-----

Use this command to return the current commission balances for an account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[http://resellertest.enom.com/resellers/SubAccount.asp](http://resellertest.enom.com/resellers/SubAccount.asp)

On the subaccount page, the Access link calls the *CommissionAccount* command.

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
| Input Parameter | Status | Description                                | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                              | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Required | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4 |

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
command=commissionaccount&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=commissionaccount&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=commissionaccount&uid=resellid
&pw=resellpw&responsetype=text
```
<PayInfoSaved>False</PayInfoSaved>

<Action/>

<data-errors></data-errors>

<CommissionBalance>172.41</CommissionBalance>

<AvailComBalance>172.41</AvailComBalance>

<HoldPeriod>90</HoldPeriod>

<OrgType/>

<OrgDesc/>

<TaxID/>

<PayableTo>Johnny Doety</PayableTo>

<Address1>15801 NE 24th Street</Address1>

<Address2/>

<City>Bellevue</City>

<StateProvince>WA</StateProvince>

<StateProvinceChoice>S</StateProvinceChoice>

<PostalCode>98008</PostalCode>

<Country>US</Country>

<Phone>+1.4252744500</Phone>

<Fax/>

<Updatable>1</Updatable>

</AccountingInfo>

<Command>COMMISSIONACCOUNT</Command>

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

<ExecTime>1.654</ExecTime>

<Done>true</Done>

<RequestDateTime>12/7/2011 4:53:52 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

PayInfoSaved: True

CommissionBalance: 142.42

AvailComBalance: 0.00

HoldPeriod: 90

OrgType: US-CORP

OrgDesc: US Corporation

TaxID: 456

PayableTo: Rosh Bach

Address1: 15801 NE 24th Street

City: ekm

StateProvince: WA

StateProvinceChoice: S

PostalCode: 98008

Country: IN

Phone: +91.9544048048

Updatable: 1

Command: COMMISSIONACCOUNT

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.250

Done: true

RequestDateTime: 2/3/2015 2:45:40 PM
```
ExecTime=0.063

RequestDateTime=2/3/2015 2:46:09 PM
```