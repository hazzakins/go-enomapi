XXX\GetMemberId
===============

Get the Member Authorization Token for a XXX domain name.

Usage
-----

If you need to replace a domain memberID with a new one, you need to remove the existing one assigned to the domain, and then add \(set\) the new memberID.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)       | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\),HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                        |
| ---------------- | ------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                  |
| ErrCount     | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\) |
| ErrX | Error messages explaining the failure. These can be presented as-is back to the client. |
| Done       | True indicates this entire response has reached you successfully. |
| DomainNameID | Domain Name ID.                                      |
| MemberID     | Member Authorization Token |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
Command=XXX_GETMEMBERID&UID=resellid&PW=resellpw
&ResponseType=XML&sld=resellertest&tld=xxx
```
```
https://resellertest.enom.com/interface.asp?
Command=XXX_GETMEMBERID&UID=resellid&PW=resellpw
&ResponseType=html&sld=resellertest&tld=xxx
```
```
https://resellertest.enom.com/interface.asp?
Command=XXX_GETMEMBERID&UID=resellid&PW=resellpw
&ResponseType=text&sld=resellertest&tld=xxx
```
```
<interface-response>
<DomainNameID>152932175</DomainNameID>
<MemberID>IXS1YFK0R185PZVA</MemberID>
<Success>True</Success>
<Command>XXX_GETMEMBERID</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t1</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+3.00</TimeDifference>
<ExecTime>0.625</ExecTime>
<Done>true</Done>
<TrackingKey>9fa90090-5b5a-4cb2-bcf7-4b813433d414</TrackingKey>
<RequestDateTime>10/9/2017 2:50:37 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>DomainNameID: </STRONG>152932175<BR /><STRONG>MemberID: </STRONG>IXS1YFK0R185PZVA<BR /><STRONG>Success: </STRONG>True<BR /><STRONG>Command: </STRONG>XXX_GETMEMBERID<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG>True<BR /><STRONG>IsRealTimeTLD: </STRONG>True<BR /><STRONG>TimeDifference: </STRONG>+3.00<BR /><STRONG>ExecTime: </STRONG>0.609<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>1e011818-b2e6-4ced-a567-5bd1c16f9064<BR /><STRONG>RequestDateTime: </STRONG>10/9/2017 2:51:25 PM<BR /></BODY></HTML>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

DomainNameID=152932175

MemberID=IXS1YFK0R185PZVA

Success=True

Command=XXX_GETMEMBERID

APIType=API.NET

Language=eng

ErrCount=0

ResponseCount=0

MinPeriod=1

MaxPeriod=10

Server=sjl0vwresell_t1

Site=eNom

IsLockable=True

IsRealTimeTLD=True

TimeDifference=+3.00

ExecTime=0.578

Done=true

TrackingKey=1dd09f6a-cd6d-40d7-adcf-5c1e725f93af

RequestDateTime=2/11/2015 10:36:07 AM
```