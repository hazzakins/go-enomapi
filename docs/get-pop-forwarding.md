GetPOPForwarding
================

Retrieve the email forwarding record—either POP or our email forwarding—for a specified mailbox name.

Usage
-----

Use this command to retrieve the email forwarding record for a specified mailbox user name.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

In the Email Settings section, the configure button retrieves forwarding records in a manner similar to the GetPOPForwarding command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The user name must be an existing email address for this domain name.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                            | Max Size |
| --------------- | ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                          | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                    | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| UserName    | Required | Mailbox name \(for example, john in [\[email protected\]](../cdn-cgi/l/email-protection.html#bed4d1d6d0fedbc6dfd3ced2db90ddd1d3)\) | 50 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| DomainName    | Domain SLD and TLD |
| DomainNameID | Numeric domain name ID, from our database                            |
| UserName     | Mailbox name |
| ForwardTo | Email forwarding address                                     |
| Active      | Is this forwarding address active. 1 indicates yes. |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the email forwarding record for [\[email protected\]](../cdn-cgi/l/email-protection.html#216b404f440f654e4461534452444d4d4453454e42520f424e4c) and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GETPOPFORWARDING&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&username=Jane.Doe
&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GETPOPFORWARDING&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&username=Jane.Doe
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GETPOPFORWARDING&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&username=Jane.Doe
&ResponseType=text
```
The response is as follows:

```html
<interface-response>
 <popforwarding>
 <domainname>resellerdocs.com</domainname>
 <domainnameid>152533676</domainnameid>
 <username>Jane.Doe</username>
 <forwardto/>
 <active/>
 </popforwarding>
 <Count>0</Count>
 <Command>GETPOPFORWARDING</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+8.00</TimeDifference>
 <ExecTime>1.063</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>0a5b34f2-b363-41f0-8686-4f64c8e32f30</TrackingKey>
 <RequestDateTime>12/9/2011 2:24:10 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
domainname: resellerdocs.com
domainnameid: 152533676
username: Jane.Doe
forwardto:
active:
Success: True
Command: GETPOPFORWARDING
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
ExecTime: 0.250
Done: true
TrackingKey: dae94654-2a03-4bf2-bb11-9ffb562a09e1
RequestDateTime: 2/11/2015 3:14:21 PM
```
Related Commands
----------------

DeletePOP3

GetCatchAll

GetPOP3

PurchasePOPBundle

SetPOPForwarding

SetUpPOP3User