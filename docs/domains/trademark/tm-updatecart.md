TM\UpdateCart
=============

Acknowledge and record the date time for a Trademark Clearinghouse \(TMCH\) Claims ID for a domain in the shopping cart.

Usage
-----

Use this command to acknowledge and record the date time for a Trademark Clearinghouse \(TMCH\) Claims ID a domain in the shopping cart.

Availability
------------

All resellers have access to this command.

Constraints
-----------
- The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | TM\_UpdateCart |
| uid | string | Required | Your Account ID                                                      |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                               |
| tcnID      | string | Required | Trademark Claims notification ID. Use the [TM\_GetNotice](../docs/domains/trademark/tm-getnotice.md) command to retrieve the value |
| tcnExpDate | string | Required | Trademark Claims expiration date \(UTC\) *Formatted as* *2014-01-10T08:00:00.0Z*                    |
| tcnAcceptDate  | string | Required | The date and time when the registrant acknowledged the Trademark Claims Notice *Formatted as* *2014-01-10T08:00:00.0Z* |
| ResponseType | string | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.*                       |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Success     | string | Update Status                                          |
| Command     | string | Name of command executed                                     |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolean | True indicates this entire response has reached you successfully.                |

Example Output
--------------

The following query acknowledges and records the date time for a Trademark Clearinghouse \(TMCH\) Claims ID for a domain in the shopping cart.

```
https://resellertest.enom.com/interface.asp?UID=ResellID&PW=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&Command=TM_UpdateCart&SLD=test---validate&tcnID=bbbf6a2f0000000000000031553&tcnAcceptDate=2014-01-10T08:00:00.0Z&tcnExpDate=2014-01-10T08:00:00.0Z&responsetype={text, XML, or HTML}
```
```
<interface-response>
 <Success>True</Success>
 <Command>TM_UpdateCart</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.031</ExecTime>
 <Done>true</Done>
 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
 <RequestDateTime>1/10/2014 2:01:35 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success=True
Command=TM_NOTIFY
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
RequestDateTime=2/6/2015 1:14:39 PM
```
```
;URL Interface<br>
;Machine is SJL0VWAPI05<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>Success:</STRONG><STRONG>True</STRONG><BR /><STRONG>Command: </STRONG>TM_CHECK<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl0vwapi05<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG>True<BR /><STRONG>IsRealTimeTLD: </STRONG>True<BR /><STRONG>TimeDifference: </STRONG>+8.00<BR /><STRONG>ExecTime: </STRONG>0.030<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>5e7ef859-36bd-4e67-bf1f-8809db691f81<BR /><STRONG>RequestDateTime: </STRONG>7/4/2016 7:40:16 AM<BR /></BODY></HTML>
```
Related Commands
----------------
- [TM\_GetNotice](../docs/domains/trademark/tm-getnotice.md)
- [TM\_Check](../docs/domains/trademark/tm-check.md)