UpdateNameServer
================

Change the IP address of a name server in the Registry’s records.

Usage
-----

Use this command when the IP address of a name server changes.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name of the name server must belong to this account.
- Enom must be the authoritative registrar \(in the Registry’s records\) for the root domain.
- The registrar lock for the name server domain must be set to off while you update name servers. Once The name server IP has been updated, you can \(and we recommend that you do\) switch the registrar lock back on.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=UpdateNameServer&uid=(Required)&pw=(Required)&OldIP=(Required)&NewIP=(Required)&NS=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | ValidatePassword |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| OldIP | string | Required | Old IP address of the name server in the Registry’s records        |
| NewIP      | string | Required | New IP address of the name server in the Registry’s records |
| NS | string | Required | Use name of the name server, in the Registry’s records          |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| NSSuccess | string | 1 indicates that the name server IP update is successful |
| RRPCode | int   | Registry processing status code |
| RRPText | string | Text explanation of registry code |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=updatenameserver&uid=resellid&pw=resellpw
&OldIP=127.0.0.1&NewIP=127.0.0.2
&NS=dns1.resellerdocs.com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=validatepassword&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&domainpassword=resellerdocs2pw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=validatepassword&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&domainpassword=resellerdocs2pw&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <RegisterNameserver>
 <NsSuccess>1</NsSuccess>
 </RegisterNameserver>
 <RRPCode>200</RRPCode>
 <RRPText>Command completed successfully</RRPText>
 <Command>UPDATENAMESERVER</Command>
 <ErrCount>0</ErrCount>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
 <STRONG>NsSuccess: </ STRONG>0<br>
 <STRONG>RRPCode: </ STRONG>200<br>
 <STRONG>RRPText: </ STRONG>Command completed successfully<br>
 <STRONG>Command: </ STRONG>UPDATENAMESERVER<br>
 <STRONG>APIType: </ STRONG>API<br>
 <STRONG>Language: </ STRONG>eng<br>
 <STRONG>ErrCount: </ STRONG>0<br>
 <STRONG>ResponseCount: </ STRONG>0<br>
 <STRONG>MinPeriod: </ STRONG>1<br>
 <STRONG>MaxPeriod: </ STRONG>10<br>
 <STRONG>Server: </ STRONG>SJL0VWRESELL_T1<br>
 <STRONG>Site:</ STRONG><br>
 <STRONG>IsLockable: </ STRONG>True<br>
 <STRONG>IsRealTimeTLD: </ STRONG>True<br>
 <STRONG>TimeDifference: </ STRONG>+08.00<br>
 <STRONG>ExecTime: </ STRONG>1.719<br>
 <STRONG>Done: </ STRONG>true<br>
 <STRONG>RequestDateTime: </ STRONG>2/9/2015 2:19:58 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
NsSuccess=0
RRPCode=200
RRPText=Command completed successfully
Command=UPDATENAMESERVER
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
ExecTime=0.816
Done=true
RequestDateTime=2/9/2015 2:20:56 PM
```
Related Commands
----------------

CheckNSStatus

DeleteNameServer

GetDNS

GetDNSStatus

ModifyNS

ModifyNSHosting

RegisterNameServer