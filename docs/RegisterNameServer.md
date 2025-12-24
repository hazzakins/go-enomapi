RegisterNameServer
==================

Register a domain name server.

Usage
-----

Use this command to register one of your own servers as a domain name server.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The server you register must use a domain name that is in your account.
- eNom must be licensed with the Registry in which you want to register a name server; you can only register name servers for TLDs that we support.
- The registrar lock for the domain must be set to off while you register a name server. You can switch the registrar lock back on \(and we recommend that you do so\) once the name server is registered.
- Name servers for .us names must be located in the United States.
- .co.uk and .org.uk names must have at least two name servers.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=RegisterNameServer&uid=(Required)&pw=(Required)&add=(Required)&NSName=(Required)&IP=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | TLD\_Overview |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| add | string | Required | Authority to add a name server. Set "Add=true" to authorize.      |
| NSName     | string | Required | Name server to register. Use format "dns1.NameServerName.com" |
| IP | string | Required | IP address of the name server.                      |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests that the computer with IP address "127.0.0.1", known as "dns1.resellerdocs.com", be registered as a name server. The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?command=registernameserver&uid=resellid&pw=resellpw&add=true&nsname=dns1.resellerdocs.com&ip=127.0.0.1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?command=registernameserver&uid=resellid&pw=resellpw&add=true&nsname=dns1.resellerdocs.com&ip=127.0.0.1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?command=registernameserver&uid=resellid&pw=resellpw&add=true&nsname=dns1.resellerdocs.com&ip=127.0.0.1&responsetype=text
```
```
<?xml version="1.0" ?>
<interface-response>
<RegisterNameserver>
<NS>dns1.resellerdocs.com</NS>
<IP>127.0.0.1</IP>
<reg-lock>1</reg-lock>
<RegistrarLock>ACTIVE</RegistrarLock>
<NsSuccess>1</NsSuccess>
<RegistrarLock>REGISTRAR-LOCK</RegistrarLock>
</RegisterNameserver>
<RRPCode>200</RRPCode>
<RRPText>Command completed successfully</RRPText>
<Command>REGISTERNAMESERVER</Command>
<ErrCount>0</ErrCount>
<Server>TestServer</Server>
<Site>enom</Site>
<Done>true</Done>
<debug>
<![CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>NS: </STRONG>dns1.resellerdocs.com<br />
<STRONG>IP: </STRONG>127.0.0.1<br />
<STRONG>RRPCode: </STRONG>200<br />
<STRONG>RRPText: </STRONG>Command completed successfully<br />
<STRONG>Command: </STRONG>REGISTERNAMESERVER<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T1<br />
<STRONG>Site: </STRONG> <br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+08.00<br />
<STRONG>ExecTime: </STRONG>0.766<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/5/2015 2:09:55 PM<br />
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
NS: dns1.resellerdocs.com
IP: 127.0.0.1
RRPCode: 200
RRPText: Command completed successfully
Command: REGISTERNAMESERVER
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site:
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.766
Done: true
RequestDateTime: 2/5/2015 2:09:55 PM
```
Related Commands
----------------

CheckNSStatus

DeleteNameServer

GetDNS

GetDNSStatus

[ModifyNS](../docs/modifyns.md)

ModifyNSHosting

UpdateNameServer