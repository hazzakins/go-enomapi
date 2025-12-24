GetSPFHosts
===========

Retrieve Sender Policy Framework \(SPF\) host records for a domain.

Usage
-----

Use this command to retrieve SPF host records for a domain.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

In the DNS settings \(zone file\) section, the test record shows the SPF host record.

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

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Name | Name of this SPF record                                     |
| Type       | Type of this SPF record |
| A-Radio | Setting for A-radio                                       |
| MX-Radio    | Setting for MX-radio |
| PTR-Radio | Setting for PTR-radio                                      |
| ALL-Radio    | Setting for ALL-radio |
| A-Records | Value for A records                                       |
| MX-Records   | Value for MX records |
| IP-Records | Value for IP records                                       |
| INC-Records   | Value for INC record |
| HostID | Host ID number                                          |
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

The following query retrieves the SPF hosts for a domain, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getspfHosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=getspfHosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getspfHosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows:

```html
<?xml version="1.0" ?>>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <host>
 <name>test</name>
 <type>TXT</type>
 <a-radio>off</a-radio>
 <mx-radio>on</mx-radio>
 <ptr-radio>on</ptr-radio>
 <all-radio>on</all-radio>
 <a-records>regular hostnames po</a-records>
 <mx-records>mx servers to</mx-records>
 <ip-records>127.0.0.1</ip-records>
 <inc-records>myisp.net mydomain.com</inc-records>
 <hostid>15195981</hostid>
 </host>
 <Command>GETSPFHOSTS</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>0.171875</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
HostName1: test
a-Radio1: off
mx-Radio1: on
ptr-Radio1: on
all-Radio1: on
a-Records1: regular hostnames po
mx-Records1: MX servers to
ip-Records1: 63.251.174.113
inc-Records1: myISP.net mydomain.com
RecordType1: TXT
MXPref1: 10
hostid1: 20021489
HostCount: 1
Command: GETSPFHOSTS
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
ExecTime: 0.094
Done: true
RequestDateTime: 2/11/2015 3:38:46 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
HostName1=test
a-Radio1=off
mx-Radio1=on
ptr-Radio1=on
all-Radio1=on
a-Records1= regular hostnames po
mx-Records1= MX servers to
ip-Records1= 63.251.174.113
inc-Records1= myISP.net mydomain.com
RecordType1=TXT
MXPref1=10
hostid1=20021489
HostCount=1
Command=GETSPFHOSTS
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
ExecTime=0.063
Done=true
RequestDateTime=2/11/2015 3:39:13 PM
```
Related Commands
----------------

GetHosts

SetHosts

SetSPFHosts