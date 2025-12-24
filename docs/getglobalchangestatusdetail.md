GetGlobalChangeStatusDetail
===========================

Retrieve a list of the domains affected by the specified global update, and the success status for each.

Usage
-----

Use this command to check the success of global changes.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[http://www.enom.com/domains/globalaccountstatus.asp](http://www.enom.com/domains/globalaccountstatus.asp)

Clicking any of the RequestID links calls the GetGlobalChangeStatusDetail command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The ResponseID must match a global update that was performed in this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                    | Max Size |
| --------------- | -------- | ----------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                 | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                     | 4 |
| RequestID | Required | ID number for the global change to one parameter. Use the GetGlobalChangeStatus command to retrieve this value. | 6    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| DomainName    | Domain name |
| StatusDesc | Status of this change to this domain                               |
| ErrorDesc    | Description of any error encountered during this change |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the status for each domain affected by global change RequestID=8331 and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetGlobalChangeStatusDetail&UID=resellid
&PW=resellpw&RequestID=8331&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetGlobalChangeStatusDetail&UID=resellid
&PW=resellpw&RequestID=8331&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetGlobalChangeStatusDetail&UID=resellid
&PW=resellpw&RequestID=8331&ResponseType=text
```
The response is as follows:

<DomainName>enomretailcart070103.biz</DomainName>

<StatusDesc>Processed successfully</StatusDesc>

<ErrorDesc>none</ErrorDesc>

</GlobalStatusDetail>

<DomainName>enomretailcart070103.com</DomainName>

<DomainName>enomretailcart070803.biz</DomainName>

<DomainName>enomretailcart070803.com</DomainName>

<DomainName>enomretailcart071703.com</DomainName>

<DomainName>enomretailcart071703.us</DomainName>

<Command>GETGLOBALCHANGESTATUSDETAIL</Command>

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

<ExecTime>0.469</ExecTime>

<Done>true</Done>

<RequestDateTime>12/8/2011 4:59:20 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

DomainName: enomretailcart070103.biz

StatusDesc: Processed successfully

ErrorDesc: none

DomainName: enomretailcart070103.com

DomainName: enomretailcart070803.biz

DomainName: enomretailcart070803.com

DomainName: enomretailcart071703.com

DomainName: enomretailcart071703.us

Command: GETGLOBALCHANGESTATUSDETAIL

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.141

Done: true

RequestDateTime: 2/4/2015 10:33:16 AM
```
ExecTime=0.047

RequestDateTime=2/4/2015 10:33:45 AM
```