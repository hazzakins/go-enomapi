GetDomainSRVHosts
=================

Retrieve SRV host records for a domain name.

Usage
-----

Use this command to retrieve a list of all SRV records for a domain name.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/control-panel/default.asp?DomainNameID=152533676\#](https://resellertest.enom.com/domains/control-panel/default.asp?DomainNameID=152533676)

The list of SRV records is populated by the GetDomainSRVHosts command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must use our DNS servers.

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

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| HostID | Identification number for this SRV record, assigned by us                    |
| HostName     | Use name you have assigned to this SRV record |
| Protocol | Internet protocol you have assigned for this SRV record                     |
| Address     | Fully qualified domain name you have assigned for this SRV record |
| RecordType | Host record type for this record—should always be SRV                      |
| MXPref      | MXPref value. Not used for SRV records |
| Weight | Weight you have assigned to this SRV record \(for load balancing\)                |
| Priority     | Priority you have assigned to this SRV record \(for backup servers\) |
| Port | Port you have assigned to this SRV record                            |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves SRV records, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GETDOMAINSRVHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GETDOMAINSRVHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GETDOMAINSRVHOSTS&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
The response is as follows:

```
<interface-response>
 <srv-records>
 <srv>
  <HostID>18459357</HostID>
  <HostName>_imap</HostName>
  <Protocol>_TCP</Protocol>
  <Address>imap.resellerdocs.com.</Address>
  <RecordType>SRV</RecordType>
  <mxPref>0</mxPref>
  <Weight>50</Weight>
  <priority>10</priority>
  <Port>1</Port>
 </srv>
 <srv>
  <HostID>18459358</HostID>
  <HostName>_xmpp</HostName>
  <Protocol>_UDP</Protocol>
  <Address>xmpp.resellerdocs.com.</Address>
  <RecordType>SRV</RecordType>
  <mxPref>0</mxPref>
  <Weight>50</Weight>
  <priority>10</priority>
  <Port>98</Port>
 </srv>
 </srv-records>
 <Command>GETDOMAINSRVHOSTS</Command>
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
 <RequestDateTime>12/8/2011 4:31:56 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
HostID1: 19937348
HostName1: _imap
Protocol1: _TCP
Address1: imap.resellerdocs.com.
RecordType1: SRV
mxPref1: 0
Weight1: 50
priority1: 10
Port1: 1
HostID2: 19937351
HostName2: _imap
Protocol2: _TCP
Address2: imap.resellerdocs.com.
RecordType2: SRV
mxPref2: 0
Weight2: 50
priority2: 10
Port2: 1
HostID3: 19937352
HostName3: _xmpp
Protocol3: _UDP
Address3: xmpp.resellerdocs.com.
RecordType3: SRV
mxPref3: 0
Weight3: 50
priority3: 10
Port3: 98
HostID4: 19937349
HostName4: _xmpp
Protocol4: _UDP
Address4: xmpp.resellerdocs.com.
RecordType4: SRV
mxPref4: 0
Weight4: 50
priority4: 10
Port4: 98
srv-count: 4
Command: GETDOMAINSRVHOSTS
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
ExecTime: 0.141
Done: true
RequestDateTime: 2/3/2015 6:05:22 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
HostID1=19937348
HostName1=_imap
Protocol1=_TCP
Address1=imap.resellerdocs.com.
RecordType1=SRV
mxPref1=0
Weight1=50
priority1=10
Port1=1
HostID2=19937351
HostName2=_imap
Protocol2=_TCP
Address2=imap.resellerdocs.com.
RecordType2=SRV
mxPref2=0
Weight2=50
priority2=10
Port2=1
HostID3=19937352
HostName3=_xmpp
Protocol3=_UDP
Address3=xmpp.resellerdocs.com.
RecordType3=SRV
mxPref3=0
Weight3=50
priority3=10
Port3=98
HostID4=19937349
HostName4=_xmpp
Protocol4=_UDP
Address4=xmpp.resellerdocs.com.
RecordType4=SRV
mxPref4=0
Weight4=50
priority4=10
Port4=98
srv-count=4
Command=GETDOMAINSRVHOSTS
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
RequestDateTime=2/3/2015 6:05:49 PM
```
Related Commands
----------------

GetDomainInfo

GetHosts

SetDomainSRVHosts

SetHosts