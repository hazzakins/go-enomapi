XXX\RemoveMemberId
==================

Removes the Member Authorization Token for a XXX domain name.

Usage
-----

If you need to replace a domain memberID with a new one, you need to remove the existing one assigned to the domain, and then add \(set\) the new memberID.

Availability
------------

All resellers have access to this command

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID **resellid**, password **resellpw**.

[https://resellertest.enom.com/domains/control-panel/default.asp?DomainNameID=152932175](https://resellertest.enom.com/domains/control-panel/default.asp?DomainNameID=152932175)

Go to General Settings.

*XXX\_RemoveMemberId* is executed before user updating/removing Member Authorization Token value.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=GetNameSuggestions&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Status | Description                                 | Max Size |
| --------------- | -------- | ---------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                               | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)         | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| MEMBERID    | Required | Member Authorization Token for domain name.                 | 16 |
| ResponseType | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\)  |
| ErrX       | Error messages explaining the failure. These can be presented as-is back to the client. |
| Done | True indicates this entire response has reached you successfully.             |
| DomainNameID   | Domain Name ID. |
| DomainName | Domain Name.                                       |
| MemberID     | Member Authorization Token \(If this is empty, then the remove failed\). |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send **ResponseType=HTML** or **ResponseType=XML** in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

Example
-------

```
https://resellertest.enom.com/interface.asp?
Command=XXX_REMOVEMEMBERID&UID=resellid&PW=resellpw
&ResponseType=XML&sld=resellertest&tld=xxx
&memberid=IXS1YFK0R185PZVA
```
```
https://resellertest.enom.com/interface.asp?
Command=XXX_REMOVEMEMBERID&UID=resellid&PW=resellpw
&ResponseType=html&sld=resellertest&tld=xxx
&memberid=IXS1YFK0R185PZVA
```
```
https://resellertest.enom.com/interface.asp?
Command=XXX_REMOVEMEMBERID&UID=resellid&PW=resellpw
&ResponseType=text&sld=resellertest&tld=xxx
&memberid=IXS1YFK0R185PZVA
```
<MemberID>IXS1YFK0R185PZVA</MemberID>

<Command>XXX_REMOVEMEMBERID</Command>

<APIType>API.NET</APIType>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod>1</MinPeriod>

<MaxPeriod>10</MaxPeriod>

<Server>sjl0vwresell_t</Server>

<Site>eNom</Site>

<IsLockable>True</IsLockable>

<IsRealTimeTLD>True</IsRealTimeTLD>

<TimeDifference>+3.00</TimeDifference>

<ExecTime>0.391</ExecTime>

<Done>true</Done>

<TrackingKey>be1bb912-2bbb-4d82-8930-226b7f7b2acd</TrackingKey>

<RequestDateTime>2/11/2015 10:39:16 AM</RequestDateTime>

<debug/>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

MemberID: IXS1YFK0R185PZVA

Command: XXX_REMOVEMEMBERID

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

ExecTime: 0.500

Done: true

TrackingKey: 33548038-4b86-4bbe-8a77-7ec3b7e6db52

RequestDateTime: 2/11/2015 10:42:00 AM
```
ExecTime=0.484

TrackingKey=e3c2660e-2160-4f0a-a73f-5f073346bbce

RequestDateTime=2/11/2015 10:43:54 AM
```