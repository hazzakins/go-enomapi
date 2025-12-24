GetRegLock
==========

Get the registrar lock setting for a domain name.

Usage
-----

Use this command to determine whether one of your domain names is locked, that is, whether it is protected from being transferred to another registrar.

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

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=GetRegLock&uid=Your Account ID&pw=Your API Token&sld={Required}&tld={Required}&ResponseType={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetRegLock |
| uid | string | Required | Your Account ID                                                            |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                      |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| Reg-Lock    | int | Returns "1" if the domain is locked, "0" otherwise, if the response type is Text.        |
| RRPCode     | string | Response code from Registry.                                  |
| RRPText     | string | Response text from Registry.                                  |
| Command     | string | GetRegLock                                            |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | string | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the lock status of a domain name and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=GetRegLock&uid=resellid&pw=resellpw&sld=resellerdocs&tld=com&ResponseType={Optional}
```
```
<interface-response>
<reg-lock>1</reg-lock>
<registrar>eNom</registrar>
<RRPCode>200</RRPCode>
<RRPText>Command completed successfully</RRPText>
<Command>GETREGLOCK</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl21wresellt01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.563</ExecTime>
<Done>true</Done>
<debug/>
<TrackingKey>95f6daeb-3bd3-4bf3-a97f-e82be665edc2</TrackingKey>
<RequestDateTime>12/9/2011 2:28:08 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>RegLock= </STRONG>0<br />
<STRONG>registrar= </STRONG>eNom<br />
<STRONG>RRPCode= </STRONG>200<br />
<STRONG>RRPText= </STRONG>Command completed successfully<br />
<STRONG>Command= </STRONG>GETREGLOCK<br />
<STRONG>APIType= </STRONG>API.NET<br />
<STRONG>Language= </STRONG>eng<br />
<STRONG>ErrCount= </STRONG>0<br />
<STRONG>ResponseCount= </STRONG>0<br />
<STRONG>MinPeriod= </STRONG>1<br />
<STRONG>MaxPeriod= </STRONG>10<br />
<STRONG>Server= </STRONG>sjl0vwresell_t<br />
<STRONG>Site= </STRONG>eNom<br />
<STRONG>IsLockable= </STRONG>True<br />
<STRONG>IsRealTimeTLD= </STRONG>True<br />
<STRONG>TimeDifference= </STRONG>+8.00<br />
<STRONG>ExecTime= </STRONG>0.469<br />
<STRONG>Done= </STRONG>true<br />
<STRONG>TrackingKey= </STRONG>9eda41bc-cb30-4525-b25f-f3cb7b7f34ef<br />
<STRONG>RequestDateTime= </STRONG>2/4/2015 11:35:55 AM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RegLock=0
registrar=eNom
RRPCode=200
RRPText=Command completed successfully
Command=GETREGLOCK
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
ExecTime=0.469
Done=true
TrackingKey=9eda41bc-cb30-4525-b25f-f3cb7b7f34ef
RequestDateTime=2/4/2015 11:35:55 AM
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

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

StatusDomain

ValidatePassword