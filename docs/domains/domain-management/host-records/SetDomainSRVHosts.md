SetDomainSRVHosts
=================

Create or edit SRV host records for a domain name.

Usage
-----

Use this command to create or edit SRV \(service\) records for a domain name.

This command deletes all previously existing SRV records for the domain name. You must include all SRV records that you want to end up with, in the query string.

SRV records are required for some newer Internet protocols such as SIP and XMPP. Some other protocols support, but do not require, SRV records.

This command allows a user to supply the components of an SRV record that are discretionary, then adds TTL, Class, and formatting to produce a complete SRV record.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain name must use our name servers

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetDomainSRVHosts&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&ServiceX=(Required)&ProtocolX=(Required)&PriorityX=(Required)&WeightX(Required)&PortX=(Required)&TargetX=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status      | Description |
| --------------- | ------ | ----------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required     | SetDomainSRVHosts |
| uid | string | Required | Your Account ID                                                                                                                |
| pw       | string | Required     | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                                                                         |
| TLD       | string | Required     | Top-level domain name \(extension\) |
| HostIDX | string | Optional Required | ID number of an existing host record, if it exists and you want to edit it. Use the GetDomainSRVHosts command to retrieve SRV record ID numbers.                                              |
| ServiceX    | int | Required     | Service type of this record. Required format is 1 to 14 characters in length; includes lower-case letters, digits, and hyphens; and begins and ends with a letter or digit |
| ProtocolX | string | Required | Transport protocol for this record, such as TCP or UDP                                                                                            |
| PriorityX    | int | Required     | Priority for this record. Lowest priority values are used first, working toward higher priority values when lower values are unavailable. Use this value to designate backup service. Permitted values are numbers from 0 to 65535 |
| WeightX | int  | Required | Proportion of time to use this record. Records of equal priority are added and the total normalized to 100%. Use this parameter to specify load balancing among records of the same priority. Permitted values are numbers from 0 to 65535 |
| PortX      | int | Required     | Port to use for this service. Permitted format is numbers from 0 to 65535 |
| TargetX | string | Required | Fully qualified domain name for this SRV record. Permitted format is SRVRecordName.SLD.TLD. \(include the trailing period in a fully qualified domain name\)                                       |
| ResponseType  | string | Optional     | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| Balance | int   | Current account balance |
| AvailableBalance | int   | Current available balance |
| Price | int   | Default price for names |
| COMPrice | int   | Price for .com |
| NETPrice | int   | Price for .net |
| ORGPrice | int   | Price for .org |
| CCPrice | int   | Price for .cc |
| TVPrice | int   | Price for .tv |
| DomainCount | int   | Current count of domains in the account |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SETDOMAINSRVHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&HostID1=17217418&Service1=imap
&Protocol1=TCP&Priority1=10&Weight1=50&Port1=1
&Target1=imap.resellerdocs.com.&HostID2=&Service2=xmpp
&Protocol2=UDP&Priority2=10&Weight2=50&Port2=98
&Target2=xmpp.resellerdocs.com.&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=SETDOMAINSRVHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&HostID1=17217418&Service1=imap
&Protocol1=TCP&Priority1=10&Weight1=50&Port1=1
&Target1=imap.resellerdocs.com.&HostID2=&Service2=xmpp
&Protocol2=UDP&Priority2=10&Weight2=50&Port2=98
&Target2=xmpp.resellerdocs.com.&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=SETDOMAINSRVHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&HostID1=17217418&Service1=imap
&Protocol1=TCP&Priority1=10&Weight1=50&Port1=1
&Target1=imap.resellerdocs.com.&HostID2=&Service2=xmpp
&Protocol2=UDP&Priority2=10&Weight2=50&Port2=98
&Target2=xmpp.resellerdocs.com.&ResponseType=text
```
```
<interface-response>
 <InsertedCount>0</InsertedCount>
 <UpdatedCount>0</UpdatedCount>
 <DeletedCount>0</DeletedCount>
 <Success>true</Success>
 <Command>SETDOMAINSRVHOSTS</Command>
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
 <ExecTime>0.188</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/11/2011 11:15:59 PM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>InsertedCount: </ STRONG>0<br>
<STRONG>UpdatedCount: </ STRONG>0<br>
<STRONG>DeletedCount: </ STRONG>2<br>
<STRONG>Success: </ STRONG>True<br>
<STRONG>Command: </ STRONG>SETDOMAINSRVHOSTS<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.578<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>b223faa1-bec4-451a-9a1b-0f353617fb19<br>
<STRONG>RequestDateTime: </ STRONG>2/5/2015 2:48:52 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
InsertedCount=0
UpdatedCount=0
DeletedCount=2
Success=True
Command=SETDOMAINSRVHOSTS
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
TrackingKey=361c182a-70ea-4d3a-a9f6-0c54f166a582
RequestDateTime=2/5/2015 2:49:25 PM
```
Related Commands
----------------

GetDomainInfo

GetDomainSRVHosts

GetHosts

SetHosts