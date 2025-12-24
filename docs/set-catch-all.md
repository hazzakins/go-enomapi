SetCatchAll
===========

Set a forwarding address for any emails addressed to nonexistent mailboxes.

Usage
-----

Use this command to forward email addressed to nonexistent email addresses under a domain name. The domain name must be in our system, but the forwarding address does not need to be.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/POPConfig.asp?DomainNameId=152533676](https://resellertest.enom.com/domains/POPConfig.asp?DomainNameId=152533676)

The Catch-All Email box at the bottom of the page uses the SetCatchAll command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The ForwardTo address does not need to be in our system.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                               | Max Size |
| --------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                             | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                       | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| ForwardTo    | Required | Email address to forward mail to. Use format [\[email protected\]](../cdn-cgi/l/email-protection.html#4a272b23262825320a39262e643e262e) | 130 |
| ResponseType | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| CatchAll     | Email address that mail will be sent to if the addressee mailbox does not exist |
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

The following query forwards any email addressed to nonexistent mailboxes for this domain name, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SetCatchAll&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com
&[email protected]&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SetCatchAll&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com
&[email protected]&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SetCatchAll&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com
&[email protected]&responsetype=text
```
The response is as follows:

```
<interface-response>
 <CatchAll>[email protected]</CatchAll>
 <Command>SETCATCHALL</Command>
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
 <ExecTime>0.375</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/11/2011 10:39:46 PM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
success: True
ForwardTo: [email protected]
Version: 2
Command: SETCATCHALL
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.547
Done: true
TrackingKey: 601121bb-c00d-4be5-8028-632e4c22064f
RequestDateTime: 2/5/2015 2:38:40 PM
```
Related Commands
----------------

DeletePOP3

Forwarding

GetCatchAll

GetDotNameForwarding

GetForwarding

GetPOP3

GetPOPForwarding

SetDotNameForwarding

SetPOPForwarding

SetUpPOP3User