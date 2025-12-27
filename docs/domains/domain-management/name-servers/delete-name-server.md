DeleteNameServer
================

Delete a name server at the Registry.

Usage
-----

Use this command to delete a name server at the Registry. This command does not affect the name servers listed for any individual domain.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[http://www.enom.com/domains/RegNameServer.asp](http://www.enom.com/domains/RegNameServer.asp)

In the Delete a Name Server box, the submit button calls the DeleteNameServer command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name server must be registered.
- The domain name this DNS uses must belong to this account.
- The domain name server must have no domains using it.
- All registrar locks and Registry holds must be off, and the domain must be in an active state.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                       | Max Size |
| --------------- | ----------------------------- | -------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                     | 20 |
| PW | Required           | Account password | 20    |
| NS       | Required | Use name of name server, in format dns1.ServerName.com | 60 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| RegistrarLock  | Registrar lock status |
| NSSuccess | Success of the name server reset process                             |
| RegistrarLock  | Registrar lock setting |
| RRPCode | Registry code                                          |
| RRPText     | Description of Registry code |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query deletes the name server registration for dns3.resellerdocs.com at the Registry, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=deletenameserver&uid=resellid&pw=resellpw
&ns=dns3.resellerdocs.com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=deletenameserver&uid=resellid&pw=resellpw
&ns=dns3.resellerdocs.com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=deletenameserver&uid=resellid&pw=resellpw
&ns=dns3.resellerdocs.com&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <RegisterNameserver>
 <RegistrarLock>ACTIVE</RegistrarLock>
 <NsSuccess>1</NsSuccess>
 <RegistrarLock>REGISTRAR-LOCK</RegistrarLock>
 </RegisterNameserver>
 <RRPCode>200</RRPCode>
 <RRPText>Command completed successfully</RRPText>
 <Command>DELETENAMESERVER</Command>
 <ErrCount>0</ErrCount>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
 <debug>
 <![CDATA [ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RegistrarLock: ACTIVE
NsSuccess: 1
RegistrarLock: REGISTRAR-LOCK
Command: DELETENAMESERVER
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site:
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 1.016
Done: true
RequestDateTime: 2/11/2015 12:58:51 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
RegistrarLock=ACTIVE
NsSuccess=0
RegistrarLock=REGISTRAR-LOCK
Command=DELETENAMESERVER
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
ExecTime=0.359
Done=true
RequestDateTime=2/11/2015 1:00:39 PM
```
Related Commands
----------------

CheckNSStatus

GetDNS

GetDNSStatus

ModifyNS

ModifyNSHosting

RegisterNameServer

UpdateNameServer