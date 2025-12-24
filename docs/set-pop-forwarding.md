SetPOPForwarding
================

Create, change, or delete an email forwarding address for a domain name.

Usage
-----

Use this command to set auto-renew behavior for an individual POP pak.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/MailConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/MailConfig.asp?DomainNameID=152533676)

The save changes button calls the SetPOPForwarding command.

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
| Input Parameter | Status | Description                                                                                                      | Max Size |
| --------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                                                   | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                                                             | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| UserName    | Required | The user name for this mailbox \(email address will be [\[email protected\]](../cdn-cgi/l/email-protection.html#e8bd9b8d9aa689858da89b848cc69c848c)\).                                | 30 |
| EMail | Required | Domain for this mailbox. Must be the same values as SLD and TLD above. Use format SLD.TLD \(email address will be [\[email protected\]](../cdn-cgi/l/email-protection.html#1b4e687e69557a767e5b68777f356f777f)\). | 78    |
| ForwardTo    | Required | Email address to forward to. Use format [\[email protected\]](../cdn-cgi/l/email-protection.html#47012835302635232e29200623233522343407223f262a372b226924282a)                            | 130 |
| Enable | Required | Add this email forwarding. Permitted values are: 0 Delete this email forwarding address 1 Add or change this email forwarding address | 1    |
| StatusFlags   | Optional | Deliver to forwarding address. Permitted values are: no value Deliver to original mailbox only 0 Deliver to both original mailbox and forwarding address 1 Deliver only to the forwarding address          | 1 |
| ResponseType | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Success     | 1 indicates the query was successful |
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

The following query creates an email forwarding record: Email addressed to [\[email protected\]](../cdn-cgi/l/email-protection.html#bccfddd0d9cffcced9cfd9d0d0d9ced8d3dfcf92dfd3d1) will be forwarded to [\[email protected\]](../cdn-cgi/l/email-protection.html#d6bfb8b0b996a4b3a5b3babab3a4b2b9b5a5f8b5b9bb), and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SETPOPFORWARDING&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&username=sales
&Email=resellerdocs.com&[email protected]
&Enable=1&StatusFlags=1&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SETPOPFORWARDING&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&username=sales
&Email=resellerdocs.com&[email protected]
&Enable=1&StatusFlags=1&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=SETPOPFORWARDING&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&username=sales
&Email=resellerdocs.com&[email protected]
&Enable=1&StatusFlags=1&ResponseType=text
```
In the response, the Success value 1 indicates that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <success>1</success>
 <Command>SETPOPFORWARDING</Command>
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
 <ExecTime>0.1328125</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
success: 1
Command: SETPOPFORWARDING
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
ExecTime: 1.063
Done: true
TrackingKey: 9f4ee08c-d5c7-4a2e-8f97-afeab5cbaaaa
RequestDateTime: 2/5/2015 3:17:18 PM
```
Related Commands
----------------

DeletePOP3

GetCatchAll

GetPOP3

GetPOPForwarding

PurchasePOPBundle

SetUpPOP3User