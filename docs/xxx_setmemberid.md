XXX\SetMemberId
===============

Sets the Member Authorization Token for a XXX domain name.

Usage
-----

If you need to replace a domain memberID with a new one, you need to remove the existing one assigned to the domain, and then add \(set\) the new memberID.

Availability
------------

All resellers have access to this command.

Implemenatation on eNom.com
----------------------------

Log on to resellertest.enom.com with Login ID **resellid**, password **resellpw**.

[https://resellertest.enom.com/domains/control-panel/default.asp?DomainNameID=152932175](https://resellertest.enom.com/domains/control-panel/default.asp?DomainNameID=152932175)

Go to General Settings.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The member authorization token must belong to domain name.

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
| SLD       | Required | Second-level domain name \(for example, *enom* in enom.com\)       | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| MEMBERID    | Required | Member Authorization Token for domain name.                | 16 |
| ResponseType | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\). |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.             |
| DomainNameID   | Domain Name ID. |
| MemberID | Member Authorization Token                                 |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send **ResponseType=HTML** or **ResponseType=XML** in your request.

Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter **Err\(ErrCount\)** can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

```
https://resellertest.enom.com/interface.asp?
Command=XXX_SETMEMBERID&UID=resellid
&PW=resellpw&ResponseType=XML&sld=resellertest
&tld=xxx&memberid=IXS1YFK0R185PZVA
```
```
https://resellertest.enom.com/interface.asp?
Command=XXX_SETMEMBERID&UID=resellid
&PW=resellpw&ResponseType=html&sld=resellertest
&tld=xxx&memberid=IXS1YFK0R185PZVA
```
```
https://resellertest.enom.com/interface.asp?
Command=XXX_SETMEMBERID&UID=resellid
&PW=resellpw&ResponseType=text&sld=resellertest
&tld=xxx&memberid=IXS1YFK0R185PZVA
```
<DomainNameID>152932175</DomainNameID>

<DomainName>resellertest.xxx</DomainName>

<MemberID>IXS1YFK0R185PZVA</MemberID>

<Success>True</Success>

<Command>XXX_SETMEMBERID</Command>

<APIType>API.NET</APIType>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod>1</MinPeriod>

<MaxPeriod>10</MaxPeriod>

<Server>sjl0vwresell_t1</Server>

<Site>eNom</Site>

<IsLockable>True</IsLockable>

<IsRealTimeTLD>True</IsRealTimeTLD>

<TimeDifference>+3.00</TimeDifference>

<ExecTime>0.516</ExecTime>

<Done>true</Done>

<TrackingKey>54003f60-3947-48fd-9df5-f188078a56d7</TrackingKey>

<RequestDateTime>2/11/2015 10:48:04 AM</RequestDateTime>

<debug/>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

DomainNameID: 152932175

DomainName: resellertest.xxx

MemberID: IXS1YFK0R185PZVA

Success: True

Command: XXX_SETMEMBERID

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t1

Site: eNom

IsLockable: True

IsRealTimeTLD: True

TimeDifference: +3.00

ExecTime: 0.766

Done: true

TrackingKey: 4fe33395-643c-4daf-ac91-a0968a378f3a

RequestDateTime: 2/11/2015 10:46:20 AM
```
;Machine is SJL0VWRESELL_T

Server=sjl0vwresell_t

ExecTime=0.531

TrackingKey=4759c363-fe30-44f0-8d07-0d833b04b6af

RequestDateTime=2/11/2015 10:46:48 AM
```