GetDNS
======

Retrieve the name server settings for a domain name.

Usage
-----

Use this command to retrieve name servers for one domain name.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The account identified in the UID parameter must be a reseller account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetDNS&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetDNS |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| NSCount     | string | First Name                                            |
| DNSX       | string | Last Name                                            |
| UseDNS      | string | Job Title                                            |
| NSStatus     | string | Street Address Line 1                                      |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetDNS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetDNS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDNS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
```
<interface-response>
 <dns>dns1.name-services.com</dns>
 <dns>dns2.name-services.com</dns>
 <dns>dns3.name-services.com</dns>
 <dns>dns4.name-services.com</dns>
 <dns>dns5.name-services.com</dns>
 <UseDNS>default</UseDNS>
 <HostsNumLimit>15</HostsNumLimit>
 <DNSRegistrySynced>True</DNSRegistrySynced>
 <RRPCodeGDNS>200</RRPCodeGDNS>
 <RRPText>Command completed successfully</RRPText>
 <Command>GETDNS</Command>
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
 <ExecTime>0.594</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/8/2011 3:58:52 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>DNS1: </ STRONG>dns1.name-services.com<br>
<STRONG>DNS2: </ STRONG>dns2.name-services.com<br>
<STRONG>DNS3: </ STRONG>dns3.name-services.com<br>
<STRONG>DNS4: </ STRONG>dns4.name-services.com<br>
<STRONG>DNS5: </ STRONG>dns5.name-services.com<br>
<STRONG>NSCount: </ STRONG>5<br>
<STRONG>UseDNS: </ STRONG>default<br>
<STRONG>HostsNumLimit: </ STRONG>15<br>
<STRONG>NameserverRegistrySynced: </ STRONG>True<br>
<STRONG>RRPCodeGDNS: </ STRONG>200<br>
<STRONG>RRPText: </ STRONG>Command completed successfully<br>
<STRONG>Command: </ STRONG>GETDNS<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable: </ STRONG>True<br>
<STRONG>IsRealTimeTLD: </ STRONG>True<br>
<STRONG>TimeDifference: </ STRONG>+08.00<br>
<STRONG>ExecTime: </ STRONG>0.188<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 5:08:35 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DNS1=dns1.name-services.com
DNS2=dns2.name-services.com
DNS3=dns3.name-services.com
DNS4=dns4.name-services.com
DNS5=dns5.name-services.com
NSCount=5
UseDNS=default
HostsNumLimit=15
NameserverRegistrySynced=True
RRPCodeGDNS=200
RRPText=Command completed successfully
Command=GETDNS
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
ExecTime=0.234
Done=true
RequestDateTime=2/3/2015 5:09:02 PM
```
Related Commands
----------------

CheckNSStatus

DeleteNameServer

GetDNSStatus

ModifyNS

ModifyNSHosting

RegisterNameServer

UpdateNameServer