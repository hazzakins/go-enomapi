PE\GetPOPPrice
==============

Retrieve the wholesale price that this account pays for POP mail 10-paks.

Usage
-----

Use this command to retrieve the wholesale price that this account pays for POP mail 10-paks.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/ProductPricing.asp?tab=domainaddons](https://resellertest.enom.com/myaccount/ProductPricing.asp?tab=domainaddons)

In the POP3 email paks row, the value in the Your cost column is supplied by the PE\_GetPOPPrice command.

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
| Input Parameter | Status | Description                                | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                              | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 63 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Pop | Wholesale price this account pays per POP 10-pak                        |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GETPOPPRICE&uid=resellid&pw=resellpw
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETPOPPRICE&uid=resellid&pw=resellpw
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETPOPPRICE&uid=resellid&pw=resellpw
&responsetype=text
```
```
<?xml version="1.0" ?>

 <pop>18.4015</pop>

 </pricing>

 <Command>PE_GETPOPPRICE</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>8.203125E-02</ExecTime>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

Price: 9.95

Command: PE_GETPOPPRICE

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.063

Done: true

RequestDateTime: 2/4/2015 4:26:43 PM
```
;Machine is SJL0VWRESELL_T1

Server=SJL0VWRESELL_T1

ExecTime=0.047

RequestDateTime=2/4/2015 4:27:09 PM
```