GetIDNCodes
===========

Retrieve a list of language codes currently supported for a TLD.

Usage
-----

Use this command to retrieve a list of language codes currently supported for a TLD.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.comwith Login ID resellid, password resellpw.

Go to: [https://resellertest.enom.com/domains/PreConfigure.asp](https://resellertest.enom.com/domains/PreConfigure.asp)

With a configurable domain in the cart, click the configure button below the list of cart contents.

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
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| TLD       | Required | Top-level domain name \(extension\)         | 15 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Language     | Language name and code value |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves list of language codes currently supported for a TLD, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
Command=GetIDNCodes&UID=ResellID
&PW=resellpw&TLD=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
Command=GetIDNCodes&UID=ResellID
&PW=resellpw&TLD=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
Command=GetIDNCodes&UID=ResellID
&PW=resellpw&TLD=com&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
<tlds>
 <tld tld="com">
 <language code="afr" name="Afrikaans" />
 <language code="alb" name="Albanian" />
 <language code="ara" name="Arabic" />
 <language code="arg" name="Aragonese" />
........<cut> ........
 <language code="wel" name="Welsh" />
 <language code="yid" name="Yiddish" />
 </tld>
</tlds>
<Command>GETIDNCODES</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.047</ExecTime>
<Done>true</Done>
<TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
<RequestDateTime>12/12/2012 12:12:12 AM</RequestDateTime>
<debug />
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
language:
language:
language:
language:
language:
language:
language:
language:
language:
language:
language:
language:
language:
.
.
.
Command: GETIDNCODES
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +8.00
ExecTime: 0.016
Done: true
TrackingKey: 411c4bba-eca4-4901-9d80-fd479127959a
RequestDateTime: 2/4/2015 10:51:37 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
language=
language=
language=
language=
language=
language=
language=
language=
language=
language=
language=
language=
.
.
.
Command=GETIDNCODES
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.016
Done=true
TrackingKey=1516fdf4-ae23-48e3-8c10-4b11d07e5c95
RequestDateTime=2/4/2015 10:52:40 AM
```
Related Commands
----------------

Preconfigure

Purchase