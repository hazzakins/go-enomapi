SetSPFHosts
===========

Create or update Sender Policy Framework \(SPF\) host records for a domain name.

Usage
-----

Use this command to create or update SPF host records for a domain name.

The SetSPFHosts command creates or updates one record per query. To set more than one SPF host record per query, use the SetHosts command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/HostConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/HostConfig.asp?DomainNameID=152533676)

The host records setup section behaves similarly to the SetSPFHosts command.

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
| Input Parameter | Status | Description                                                                 | Max Size |
| --------------- | ----------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                              | 20 |
| PW | Required                 | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                        | 63 |
| TLD | Required                 | Top-level domain name \(extension\) | 15    |
| HostID     | Required when updating an SPF host record | Host ID number from our database. You can retrieve this value using the GetSPFHosts command.                       | 10 |
| HostName | Required                 | Host name, assigned by you | 63    |
| RecordType   | Required | For SPF records, the required value is txt                                                 | 5 |
| Address | Required                 | The SPF record. For help writing this record, you can go to [http://spf.pobox.com/wizard.html](http://spf.pobox.com/wizard.md) | 255   |
| MXPref     | Optional | Host record preference for setting the SPF record. The lower the number, the higher the priority. If not specified, default value is 10. | 5 |
| ResponseType | Optional Format of response.       | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTMLor ResponseType=XMLin your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

The following query sets SPF host records for domain resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SETSPFHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&HostName=test
&Address=v%3dspf1%20mx%20ptr%20a:regular%20a:hostnames
%20a:po%20mx:MX%20mx:servers%20mx:to%20ip4:63.251.174.113
%20include:myISP.net%20include:mydomain.com%20-all
&RecordType=txt&mxPref=&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=SETSPFHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&HostName=test
&Address=v%3dspf1%20mx%20ptr%20a:regular%20a:hostnames
%20a:po%20mx:MX%20mx:servers%20mx:to%20ip4:63.251.174.113
%20include:myISP.net%20include:mydomain.com%20-all
&RecordType=txt&mxPref=&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SETSPFHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&HostName=test
&Address=v%3dspf1%20mx%20ptr%20a:regular%20a:hostnames
%20a:po%20mx:MX%20mx:servers%20mx:to%20ip4:63.251.174.113
%20include:myISP.net%20include:mydomain.com%20-all
&RecordType=txt&mxPref=&responsetype=text
```
The response is as follows:

```
<interface-response>
 <Command>SETSPFHOSTS</Command>
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
 <ExecTime>0.485</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 12:54:35 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command: SETSPFHOSTS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.188
Done: true
RequestDateTime: 2/5/2015 3:36:38 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Command=SETSPFHOSTS
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.094
Done=true
RequestDateTime=2/5/2015 3:37:04 PM
```
Related Commands
----------------

GetHosts

SetHosts