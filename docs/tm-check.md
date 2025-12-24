TM\Check
========

Check a trademark claim for a domain name from Trademark Clearinghouse \(TMCH\).

Usage
-----

Use this command to check a trademark claim for a domain name from Trademark Clearinghouse \(TMCH\).

The Lookup Key returns a value when 1 \(one\) or more claim\(s\) exist for the specified domain name.

> ### The key value is valid for up to 24 hours. The countdown starts immediately after this command \(TM\Check\) is successfully executed.

Availability
------------

All resellers have access to this command.

Constraints
-----------
- The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------- |
| command     | string | Required | TM\_Check |
| uid | string | Required | Your Account ID                               |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)         |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.* |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                            |
| ---------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| LookupKey    | string | Lookup Key value. - \*Valid for up to 24 hours.\*\* The countdown starts immediately after this command is successfully executed |
| Command     | string | Name of command executed                                                      |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                  |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.                     |
| Done       | boolean | True indicates this entire response has reached you successfully.                                 |

Example Output
--------------

The following query checks a trademark claim for test---validate.claimsgatwo domain name from Trademark Clearinghouse \(TMCH\).

```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&Command=TM_Check&SLD=test---validate
&tld=claimsgatwo&responsetype={text, XML, or HTML}
```
```
<interface-response>
 <LookupKey>2014011000/6/e/2/bi1ZaGOpzse22CALy</LookupKey>
 <Command>TM_CHECK</Command>
 <APIType>API.NET</APIType>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.422</ExecTime>
 <Done>true</Done>
 <TrackingKey>96f0dd31-9692-4405-b960-aa120a815a4a</TrackingKey>
 <RequestDateTime>1/10/2014 9:00:00 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
TM_Check=
Command=TM_CHECK
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
TimeDifference=+0.00
ExecTime=0.016
Done=true
TrackingKey=2bbe75a6-d6bb-4d7b-b143-c2545ab5d20c
RequestDateTime=2/6/2015 12:27:25 PM
```
```
;URL Interface<br>
;Machine is SJL0VWAPI05<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>TM_Check: </STRONG><BR /><STRONG>Command: </STRONG>TM_CHECK<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl0vwapi05<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG>True<BR /><STRONG>IsRealTimeTLD: </STRONG>True<BR /><STRONG>TimeDifference: </STRONG>+8.00<BR /><STRONG>ExecTime: </STRONG>0.030<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>5e7ef859-36bd-4e67-bf1f-8809db691f81<BR /><STRONG>RequestDateTime: </STRONG>7/4/2016 7:40:16 AM<BR /></BODY></HTML>
```
Related Commands
----------------
- [TM\_GetNotice](../docs/tm-getnotice.md)
- [TM\_UpdateCart](../docs/tm-updatecart.md)