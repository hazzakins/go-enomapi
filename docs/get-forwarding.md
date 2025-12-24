GetForwarding
=============

Get email forwarding records for a domain name.

Usage
-----

Use this command to display email forwarding records for a domain name.

A similar command, GetPOPForwarding, displays both email forwarding and POP records for a single email address.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?domainnameid=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?domainnameid=152533676)

On the domain control panel, if the e-mail settings service is set to e-mail forwarding, the configure button calls the GetForwarding command.

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
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                               | Max Size |
| --------------- | ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                             | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with               | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| GetDefaultOnly | Optional | If this is set to 1, extra blank input records are returned for user input on the email forwarding form. Primarily used in XML output. | 1 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| UserNameX    | Alias name to forward. Indexed X when ResponseType=Text or HTML. |
| Alias | Alias name to forward. Present when ResponseType=XML Indexed X when ResponseType=Text or HTML.                   |
| ForwardToX    | Email address to forward the alias to. |
| Forward-To | Email address to forward the alias to. Present when ResponseType=XML                                 |
| StatusX     | Enabled status of email forwarding: 0 Email forwarding enabled 1 Email forwarding disabled Indexed X when ResponseType=Text or HTML. |
| EmailCount | Number of email forwarding records currently configured                                        |
| MaxEmail     | Total number of email forwarding records permitted for this domain name |
| EmailForwarding | True                                                                 |
| HostRecords   | True |
| Command | Name of command executed                                                       |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                       |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests e-mail forwarding information for resellerdocs.com and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getforwarding&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getforwarding&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getforwarding&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows:

```html
<interface-response>
 <eforward>
 <alias>john.doe</alias>
 <forward-to>[email protected]</forward-to>
 <status>0</status>
 <rejectreason/>
 <mailid>3370493</mailid>
 </eforward>
 <EmailCount>1</EmailCount>
 <MaxEmail>100</MaxEmail>
 <DomainServices>
 <EmailForwarding>True</EmailForwarding>
 <HostRecords>True</HostRecords>
 </DomainServices>
 <Command>GETFORWARDING</Command>
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
 <RequestDateTime>12/8/2011 4:56:11 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Username1: jane.doe
ForwardTo1: [email protected]
Status1: 0
RejectReason1:
MailID1: 4322979
EmailCount: 1
MaxEmail: 100
EmailForwarding:
HostRecords:
Command: GETFORWARDING
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
ExecTime: 0.109
Done: true
TrackingKey: 1fed5526-70ac-4347-a86b-33141c5b51df
RequestDateTime: 2/4/2015 10:28:35 AM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

Forwarding

GetCatchAll

GetDotNameForwarding

GetMailHosts

GetPOP3

ModifyPOP3

PurchasePOPBundle

SetDotNameForwarding

SetUpPOP3User