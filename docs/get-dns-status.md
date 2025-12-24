GetDNSStatus
============

Get the nameserver status for a domain name.

Usage
-----

Use this command to determine what DNS this domain uses. The query returns values only if the domain uses eNom servers.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.    | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                            |
| ---------------- | ------------------------------------------------------------------------------------------------- |
| UseDNS | Which name servers a domain is using. Default indicates our name servers.            |
| HostsNumLimit  | Maximum number of host records permitted for this domain, of types A, AAAA, CNAME, URL, and FRAME |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.     |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the DNS status of resellerdocs.com -- that is, whether the domain uses its registrar’s servers, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetDNSStatus&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDNSStatus&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDNSStatus&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
The response is as follows:

```
<interface-response>
 <UseDNS>default</UseDNS>
 <NSStatus>Yes</NSStatus>
 <HostsNumLimit>15</HostsNumLimit>
 <Command>GETDNSSTATUS</Command>
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
 <ExecTime>0.141</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/8/2011 4:03:39 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
UseDNS: default
NSStatus: Yes
HostsNumLimit: 15
Command: GETDNSSTATUS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.063
Done: true
RequestDateTime: 2/3/2015 5:11:16 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
UseDNS=default
NSStatus=Yes
HostsNumLimit=15
Command=GETDNSSTATUS
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
ExecTime=0.063
Done=true
RequestDateTime=2/3/2015 5:11:41 PM
```
Related Commands
----------------

CheckNSStatus

DeleteNameServer

GetDNS

ModifyNS

ModifyNSHosting

RegisterNameServer

UpdateNameServer