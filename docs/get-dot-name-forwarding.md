GetDotNameForwarding
====================

Retrieve the current address for email forwarding by the .name Registry.

Usage
-----

Use this command to retrieve the email forwarding address used by the .name Registry for forwarding email addressed to this domain.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

On the domain names menu, click my domains, and then click a .name domain.

In the email settings box of the domain control panel, the .name email button calls the GetDotNameForwarding command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must be a .name domain.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                 | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                              | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                            | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| DomainNameID | Domain name ID, eNom’s internal accounting number                        |
| SLD       | Second-level domain name \(for example, enom in enom.com\) |
| TLD | Top-level domain. For this command, value is name.                      |
| Status      | Status |
| Address | Email address to which email addressed to this domain is forwarded.               |
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

The following query retrieves the email forwarding address for john.doe.name and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getdotnameforwarding&uid=resellid&pw=resellpw
&sld=john.doe&tld=name&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getdotnameforwarding&uid=resellid&pw=resellpw
&sld=john.doe&tld=name&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getdotnameforwarding&uid=resellid&pw=resellpw
&sld=john.doe&tld=name&responsetype=text
```
The response is as follows:

```html
<?xml version="1.0" ?>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <dotnameforwarding>
 <domainname domainnameid="157171163" sld="john.doe" tld="name" />
 <status>1</status>
 <address>[email protected]</address>
 </dotnameforwarding>
 <Command>GETDOTNAMEFORWARDING</Command>
 <ErrCount>0</ErrCount>
 <Server>Reseller5</Server>
 <Site>enom</Site>
 <IsLockable>False</IsLockable>
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
DomainRRP: E
domainnameid: "157171163"
sld: "john.doe"
tld: "name"
status: 1
address: [email protected]
Command: GETDOTNAMEFORWARDING
ErrCount: 0
Server: Reseller5
Site: enom
IsLockable: False
IsRealTimeTLD: True
Done: true
```
Related Commands
----------------

Forwarding

GetCatchAll

GetForwarding

GetPOP3

SetCatchAll

SetDotNameForwarding