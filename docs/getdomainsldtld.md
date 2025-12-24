GetDomainSldTld
===============

Retrieve the domain name \(SLD and TLD\).

Usage
-----

Use this command when you have the domain name ID and want the domain name \(SLD and TLD\).

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
https://resellertest.enom.com/interface.asp?command=GetDomainSLDTLD&uid=Your Account ID&pw=Your API Token&DomainNameID={Required}&ResponseType={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetAllAccountInfo |
| uid | string | Required | Your Account ID                                                            |
| pw       | string | Required | Your API Token |
| DomainNameID | int  | Required | Domain name ID number. This number can be retrieved using either "GetDomainInfo" or "GetDomainNameID".                |
| ResponseType  | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| DomainRRP | string | Registrar identifier |
| SLD | string | Second-level domain name \(for example, enom in enom.com\) |
| TLD | string | Top-level domain name \(extension\) |
| DomainNameID | int  | Domain name ID |
| Command | string | GetDomainSLDTLD |
| ErrCount | int  | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | string | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the SLD and TLD for domain name ID 152533676 and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=GetDomainSLDTLD&uid=resellid&pw=resellpw&DomainNameID=152533676&ResponseType={Optional}
```
```
<interface-response>
<SLD>resellerdocs</SLD>
<TLD>com</TLD>
<DomainNameID>152533676</DomainNameID>
<Command>GETDOMAINSLDTLD</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+08.00</TimeDifference>
<ExecTime>0.125</ExecTime>
<Done>true</Done>
<RequestDateTime>12/12/2011 4:53:15 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>SLD: </STRONG>resellerdocs<br />
<STRONG>TLD: </STRONG>com<br />
<STRONG>DomainNameID: </STRONG>152533676<br />
<STRONG>Command: </STRONG>GETDOMAINSLDTLD<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+08.00<br />
<STRONG>ExecTime: </STRONG>0.047<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/3/2015 6:03:19 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
SLD=resellerdocs
TLD=com
DomainNameID=152533676
Command=GETDOMAINSLDTLD
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.078
Done=true
RequestDateTime=2/3/2015 6:03:48 PM
```
Related Commands
----------------

GetDomainInfo

GetDomainNameID