CheckNSStatus
=============

Retrieve information about a name server.

Usage
-----

Use this command when you want status information about a name server, or want to determine whether a name server exists, or is outside our credential.
- If the name server exists and is registered using our credential, this command returns status information.
- If a name server does not exist, this command informs you of that fact.
- If a name server exists but is outside our credential, this command informs you of that fact.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- To retrieve status information, the name server must exist and must be registered under our credential.
- The query must call the name server in a valid format \(by its use name, not its IP address\).

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                              | Max Size |
| --------------- | ----------------------------- | ---------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                            | 20 |
| PW | Required           | Account password | 20    |
| CheckNSName   | Required | Name server use name to check. Use format dnsX.NameServerName.com. | 63 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| NSCheckSuccess  | Was this name server check successful. 1 indicates yes; 0 indicates no. |
| Other parameters | Other parameters will be returned, and will vary from one name server to another         |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter Check the return parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the status of name server ns1.name-services.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=checknsstatus&e.asp?command=checknsstatus
&uid=resellid&pw=resellpw&checknsname=ns1.name-services.com
&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=checknsstatus&e.asp?command=checknsstatus
&uid=resellid&pw=resellpw&checknsname=ns1.name-services.com
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=checknsstatus&e.asp?command=checknsstatus
&uid=resellid&pw=resellpw&checknsname=ns1.name-services.com
&responsetype=text
```
The response is as follows:

```html
<interface-response>
 <NsCheckSuccess>1</NsCheckSuccess>
 <CheckNsStatus>
 <name>ns2.budip.com</name>
 <attrib-id>5693446_HOST_CNE-VRSN</attrib-id>
 <ipaddress>12.123.225.3</ipaddress>
 <attrib-upid>rccbatch</attrib-upid>
 <attrib-clid>024</attrib-clid>
 <attrib-crid>rccbatch</attrib-crid>
 <status>
  <status>ok</status>
  <status>linked</status>
 </status>
 <attrib-update>2005-02-02T17:38:04.0000Z</attrib-update>
 <attrib-crdate>2005-02-02T17:38:04.0000Z</attrib-crdate>
 </CheckNsStatus>
 <RRPCode>200</RRPCode>
 <RRPText>Command completed successfully</RRPText>
 <Command>CHECKNSSTATUS</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod/>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELL01</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>1.609</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/7/2011 4:40:34 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
NsCheckSuccess: 1
Item: Value
Whoisname: ns1.budip.com
Whoisattrib-id: 5693445_HOST_CNE-VRSN
Whoisipaddress: 12.123.225.2
Whoisattrib-upid: enom1
Whoisattrib-clid: 024
Whoisattrib-crid: rccbatch
WhoisstatusCount: 2
Whoisstatus: ok
Whoisstatus: linked
Whoisattrib-update: 2011-03-02T20:24:16Z
Whoisserver: SJL21WLGDEV05
Whoisattrib-crdate: 2005-02-02T17:37:51Z
RRPCode: 200
RRPText: Command completed successfully
Command: CHECKNSSTATUS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.328
Done: true
RequestDateTime: 2/3/2015 2:42:57 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
NsCheckSuccess=1
Item=Value
Whoisname=ns1.budip.com
Whoisattrib-id=5693445_HOST_CNE-VRSN
Whoisipaddress=12.123.225.2
Whoisattrib-upid=enom1
Whoisattrib-clid=024
Whoisattrib-crid=rccbatch
WhoisstatusCount=2
Whoisstatus=ok
Whoisstatus=linked
Whoisattrib-update=2011-03-02T20:24:16Z
Whoisserver=SJL21WLGDEV06
Whoisattrib-crdate=2005-02-02T17:37:51Z
RRPCode=200
RRPText=Command completed successfully
Command=CHECKNSSTATUS
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.219
Done=true
RequestDateTime=2/3/2015 2:43:29 PM
```
Related Commands
----------------

DeleteNameServer

GetDNS

GetDNSStatus

ModifyNS

ModifyNSHosting

RegisterNameServer

UpdateNameServer