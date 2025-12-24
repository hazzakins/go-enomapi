GetCatchAll
===========

Get the forwarding address for emails addressed to nonexistent mailboxes.

Usage
-----

Use this command to retrieve the forwarding address for nonexistent email addresses under a domain name. The domain name must be in our system, but the forwarding address does not need to be.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/POPConfig.asp?DomainNameId=152533676](https://resellertest.enom.com/domains/POPConfig.asp?DomainNameId=152533676)

The Catch-All Email box at the bottom of the page uses the GetCatchAll command.

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
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| CatchAll | Email address that mail will be sent to if the addressee mailbox does not exist         |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the Catch-All email address for this domain name and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetCatchAll&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetCatchAll&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetCatchAll&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <CatchAll>[email protected]</CatchAll>
 <Command>GETCATCHALL</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>0.125</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA[ ]]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
success: True
ForwardTo:
Version: 2
Command: GETCATCHALL
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
ExecTime: 0.156
Done: true
TrackingKey: a559e157-6bfe-4831-95b3-d8b48ee52c1b
RequestDateTime: 2/3/2015 4:47:51 PM
```
Related Commands
----------------

DeletePOP3

Forwarding

GetDotNameForwarding

GetForwarding

GetPOP3

GetPOPForwarding

SetCatchAll

SetDotNameForwarding

SetPOPForwarding

SetUpPOP3User