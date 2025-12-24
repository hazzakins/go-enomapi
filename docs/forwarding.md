Forwarding
==========

Set email forwarding addresses.

Usage
-----

Use this command to specify forwarding addresses for email sent to the domain.

A similar command, SetPOPForwarding, allows you to create, change, or delete mailboxes for either email forwarding or POP mail.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/MailConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/MailConfig.asp?DomainNameID=152533676)

On the e-mail settings page,the save changes button calls the Forwarding command.

Constraints
-----------

The query must meet the following requirements:
- The domain must be using eNom’s domain name server.
- The domain must exist in the account cited in the query.
- The query must include all addresses that you want to keep. Forwarding addresses not included in thequery are deleted from the account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                   | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with                   | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| AddressX    | Required | Address \(mailbox user name\) to be forwarded from this domain                                         | 30 |
| ForwardToX | Required           | Where this address will be forwarded to, in format [\[email protected\]](../cdn-cgi/l/email-protection.html#64311701162a050901241708004a100800) | 130   |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                              | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- Example: [\[email protected\]](../cdn-cgi/l/email-protection.html#8df9e5e4fecce9e9ffe8fefecde0f4e9e2e0ece4e3a3eee2e0) to be forwarded to [\[email protected\].](../cdn-cgi/l/email-protection.html#503d35103d29383f3d357e333f3d7e) Pass as Address1=thisAddress and ForwardTo1=[\[email protected\]](../cdn-cgi/l/email-protection.html#d4b9b194b9adbcbbb9b1fab7bbb9)

Example
-------

The following query gives instructions that email addressed to [\[email protected\]](../cdn-cgi/l/email-protection.html#f09a919e95de949f95b0829583959c9c9582949f9383de939f9d) and [\[email protected\]](../cdn-cgi/l/email-protection.html#d4bebbbcbafab0bbb194a6b1a7b1b8b8b1a6b0bbb7a7fab7bbb9) be forwarded to [\[email protected\]](../cdn-cgi/l/email-protection.html#a9c3c8c7cc87cdc6cce9ccd1c8c4d9c5cc87cac6c4) and [\[email protected\]](../cdn-cgi/l/email-protection.html#8de7e2e5e3a3e9e2e8cde8f5ece0fde1e8a3eee2e0) and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=Forwarding&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&Address1=jane.doe
&[email protected]&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=Forwarding&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&Address1=jane.doe
&[email protected]&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=Forwarding&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&Address1=jane.doe
&[email protected]&ResponseType=text
```
The response is as follows:

```
<interface-response>
 <EmailForwardingEnabled>True</EmailForwardingEnabled>
 <MailCount>0</MailCount>
 <Command>FORWARDING</Command>
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
 <ExecTime>8.453</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/7/2011 6:14:10 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
EmailForwardingEnabled: True
MailCount: 1
Command: FORWARDING
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.766
Done: true
TrackingKey: 613be6b5-0c4d-4f67-9619-dfe9f8d6f7c2
RequestDateTime: 2/3/2015 4:04:29 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

GetCatchAll

GetDotNameForwarding

GetForwarding

GetMailHosts

GetPOP3

PurchasePOPBundle

SetCatchAll

SetDotNameForwarding

SetPOPForwarding

SetUpPOP3User