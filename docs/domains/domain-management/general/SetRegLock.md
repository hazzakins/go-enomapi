SetRegLock
==========

Set the registrar lock for a domain name.

Usage
-----

Set the registrar unlock value to 0 to prevent unauthorized transfer of a domain name to another registrar.

Set the registrar unlock value to 1 to allow unrestricted transfer of a domain name.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetRegLock&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&UnlockRegistrar=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | SetRegLock |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| sld | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| tld       | string | Required | Top-level domain name \(extension\) |
| UnlockRegistrar | int  | Required | Lock option to be set. Permitted values are 1 to unlock, 0 to lock\).  |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| reg-lock | string | Current lock status. |
| RRPCodeSR | string | Registry response code for our internal use. |
| RRPText | string | Verbose description of RRPCodeSR. |
| TrackingKey | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=setreglock&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&unlockregistrar=0&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=setreglock&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&unlockregistrar=0&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=setreglock&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&unlockregistrar=0&responsetype=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <reg-lock>REGISTRAR-LOCK</reg-lock>
 <RRPCodeSR>200</RRPCodeSR>
 <RRPText>Command completed successfully</RRPText>
 <Command>SETREGLOCK</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>Dev Workstation</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+8.00</TimeDifference>
 <ExecTime>0.484</ExecTime>
 <Done>true</Done>
 <TrackingKey>d79aed2c-73e9-4503-bc38-c700a61bfe54</TrackingKey>
 <RequestDateTime>12/11/2011 10:35:56 PM</RequestDateTime>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>reg-lock: </STRONG>REGISTRAR-LOCK<br />
<STRONG>RRPCode: </STRONG>200<br />
<STRONG>RRPText: </STRONG>Command completed successfully<br />
<STRONG>Command: </STRONG>SETREGLOCK<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>0.297<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>3cacf4a9-abdb-485b-b3ca-981e3b475e7f<br />
<STRONG>RequestDateTime: </STRONG>2/5/2015 3:20:40 PM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
reg-lock=REGISTRAR-LOCK
RRPCode=200
RRPText=Command completed successfully
Command=SETREGLOCK
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
TimeDifference=+8.00
ExecTime=0.125
Done=true
TrackingKey=657eb399-f32a-41d3-9de8-a28c5612dd5a
RequestDateTime=2/5/2015 3:21:13 PM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetExtendInfo

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

SetRenew

StatusDomain

ValidatePassword