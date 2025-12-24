DisableServices
===============

Switch off a service in an account.

Usage
-----

Use this command to retrieve name servers for one domain name.

This command does not shut down the service or cancel it at the end of the current billing period; it merely toggles it off.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- Other constraints may apply depending on the services being disabled.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=DisableServices&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&Service=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | DisableServices |
| uid | string | Required | Your Account ID                                                      |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                |
| TLD       | string | Required | Top-level domain name \(extension\) |
| Service | String | Required | Name of the service to be disabled. Permitted values are: WPPS \(Whois Privacy Protection\) EmailForwarding URLForwarding |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| Service | string | Name of the service being disabled |
| ServiceStatus | string | New status of this service |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=DISABLESERVICES&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=wpps&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=DISABLESERVICES&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=wpps&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=DISABLESERVICES&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=wpps&ResponseType=text
```
```
<interface-response>
 <Service>wpps</Service>
 <ServiceStatus>Disabled</ServiceStatus>
 <Command>DISABLESERVICES</Command>
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
 <ExecTime>0.250</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/7/2011 5:48:34 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Service: </ STRONG>wpps<br>
<STRONG>ServiceStatus: </ STRONG>Disabled<br>
<STRONG>Command: </ STRONG>DISABLESERVICES<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+08.00<br>
<STRONG>ExecTime: </ STRONG>0.469<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 3:41:37 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Service=wpps
ServiceStatus=Disabled
Command=DISABLESERVICES
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.469
Done=true
RequestDateTime=2/3/2015 3:41:37 PM
```
Related Commands
----------------

EnableServices

GetWPPSInfo

ServiceSelect