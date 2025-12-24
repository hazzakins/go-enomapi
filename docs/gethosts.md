GetHosts
========

Get host records for a domain name.

Usage
-----

Use this command to return all host records. "GetRegHosts", a similar command, does not return mail host \(MX\) records.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must use eNom’s domain name servers.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=GetHosts&uid=Your Account ID&pw=Your API Token&SLD={Required}&TLD={Required}&ResponseType={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetHosts |
| uid | string | Required | Your Account ID                                                            |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                       |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                                                                                                           |
| ---------------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name       | string | Host name of record.                                                                                                                                      |
| Type       | string | Record Type. Expected returned values are: - A - AAAA - URL - MX - MXE - CNAME - FRAME.                                                                                                    |
| MXPref      | int | Preference level. The lower this value, the higher priority this service. Default value is: 10.                                                                                               |
| Address     | string | Address. Can be an IP for record types "A", "MXE" and "FRAME" only. Record type "AAAA" accepts an IP address in IPv6 format. Record types "URL" and "FRAME" can be a full URL, beginning with http://. Record types "MX", "URL", "FRAME" and "CNAME" can be a fully qualified domain name. |
| Command     | string | GetHosts, Name of command executed.                                                                                                                              |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                    |
| Done       | string | "True" indicates this entire response has reached you successfully.                                                                                                              |

Example Output
--------------

The following query requests all host records, including mail hosts, for the domain "resellerdocs.com" and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=GetHosts&uid=Your Account ID&pw=Your API Token&SLD=resellerdocs&TLD=com&ResponseType={Optional}
```
```html
<interface-response>
<host>
<name>@</name>
<type>A</type>
<address>66.150.5.189</address>
<hostid>18149735</hostid>
</host>
<host>
<name>photos</name>
<type>CNAME</type>
<address>photos.msn.com.</address>
<hostid>18487894</hostid>
</host>
<host>
<name>yahoo</name>
<type>URL</type>
<address>204.71.200.72</address>
<hostid>18487895</hostid>
</host>
<host>
<name>msn</name>
<type>FRAME</type>
<address>http://www.msn.com</address>
<hostid>18487896</hostid>
</host>
<host>
<name>mail</name>
<type>MXE</type>
<mxpref>10</mxpref>
<address>209.19.56.20</address>
<hostid>18492166</hostid>
</host>
<DomainServices>
<EmailForwarding>True</EmailForwarding>
<HostRecords>True</HostRecords>
</DomainServices>
<Command>GETHOSTS</Command>
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
<ExecTime>0.172</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 5:02:39 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>HostName1: </STRONG>@<br />
<STRONG>Address1: </STRONG>66.150.5.189<br />
<STRONG>RecordType1: </STRONG>A<br />
<STRONG>MXPref1: </STRONG>10<br />
<STRONG>hostid1: </STRONG>19920062<br />
<STRONG>HostName2: </STRONG>photos<br />
<STRONG>Address2: </STRONG>photos.msn.com.<br />
<STRONG>RecordType2: </STRONG>CNAME<br />
<STRONG>MXPref2: </STRONG>10<br />
<STRONG>hostid2: </STRONG>19987522<br />
<STRONG>HostName3: </STRONG>msn<br />
<STRONG>Address3: </STRONG>http://www.msn.com<br />
<STRONG>RecordType3: </STRONG>FRAME<br />
<STRONG>MXPref3: </STRONG>10<br />
<STRONG>hostid3: </STRONG>19995603<br />
<STRONG>HostName4: </STRONG>mail<br />
<STRONG>Address4: </STRONG>209.19.56.20<br />
<STRONG>RecordType4: </STRONG>MXE<br />
<STRONG>MXPref4: </STRONG>10<br />
<STRONG>hostid4: </STRONG>20019927<br />
<STRONG>HostName5: </STRONG>yahoo<br />
<STRONG>Address5: </STRONG>204.71.200.72<br />
<STRONG>RecordType5: </STRONG>URL<br />
<STRONG>MXPref5: </STRONG>10<br />
<STRONG>hostid5: </STRONG>19987523<br />
<STRONG>HostCount: </STRONG>5<br />
<STRONG>EmailForwarding: </STRONG><br />
<STRONG>HostRecords: </STRONG><br />
<STRONG>Command: </STRONG>GETHOSTS<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+08.00<br />
<STRONG>ExecTime: </STRONG>0.188<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/4/2015 10:46:50 AM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
HostName1=@
Address1=66.150.5.189
RecordType1=A
MXPref1=10
hostid1=19920062
HostName2=photos
Address2=photos.msn.com.
RecordType2=CNAME
MXPref2=10
hostid2=19987522
HostName3=msn
Address3=http://www.msn.com
RecordType3=FRAME
MXPref3=10
hostid3=19995603
HostName4=mail
Address4=209.19.56.20
RecordType4=MXE
MXPref4=10
hostid4=20019927
HostName5=yahoo
Address5=204.71.200.72
RecordType5=URL
MXPref5=10
hostid5=19987523
HostCount=5
EmailForwarding=
HostRecords=
Command=GETHOSTS
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
RequestDateTime=2/4/2015 10:47:26 AM
```
Related Commands
----------------

GetMetaTag

GetRegHosts

GetSPFHosts

[SetDNSHost](../docs/setdnshost.md)

[SetHosts](../docs/set-hosts.md)

SetSPFHosts

UpdateMetaTag