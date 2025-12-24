PE\GetRocketPrice
=================

Get the pricing and enabled state for a Registry Rocket key.

Usage
-----

Use this command to get the price for one product, for one top-level domain. For example, you can use this command to retrieve the price for renewing a .org name. This command is most useful to resellers who offer a restricted set of top-level domains.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/CCProcessingSignUp.asp](https://resellertest.enom.com/resellers/CCProcessingSignUp.asp)

This command is not implemented on enom.com. However, on the registry rocket page,

the create link button returns similar information.

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
&paramname=paramvalue &nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                           | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                         | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML             | 4 |
| TLD | Required | Top-level domain \(extension, e.g. com\) for which you want information. | 15    |
| ProductType   | Required | Product type. Options are: 10 register 13 DNS hosting 14 DNS hosting renew 16 renew 19 transfer | 3 |
| ResellerKey | Required | Unique key created for each Registry Rocket site | 40    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Product types are: 10=register, 13=hosting, 14=hosting renew, 16=renew, 19=transfer.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GetRocketPrice&uid=resellid&pw=resellpw&tld=org
&ProductType=10&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GetRocketPrice&uid=resellid&pw=resellpw&tld=org
&ProductType=10&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GetRocketPrice&uid=resellid&pw=resellpw&tld=org
&ProductType=10&responsetype=text
```
```
<?xml version="1.0" ?>

 <price>29.95</price>

 <productenabled>True</productenabled>

 <years>1</years>

 </productprice>

 <Command>PE_GETROCKETPRICE</Command>

 <ErrCount>0</ErrCount>

 <Server>ResellerTest</Server>

 <Site>eNom</Site>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

price: 29.95

productenabled: True

years: 1

Command: PE_GETROCKETPRICE

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site:

IsLockable: True

IsRealTimeTLD: True

TimeDifference: +07.00

ExecTime: 0.152

Done: true

RequestDateTime: 2/12/2015 11:45:51 AM
```
;Machine is SJL0VWRESELL_T1

Server=SJL0VWRESELL_T1

ExecTime=0.063

RequestDateTime=2/12/2015 11:47:05 AM
```