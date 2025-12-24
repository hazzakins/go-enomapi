SetDotNameForwarding
====================

Set the email forwarding address for a .name domain.

Usage
-----

Use this command if you want to use the .name Registry’s email forwarding feature \(you can also use eNom’s email forwarding with .name domains\).

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

On the domain names menu, click my domains, and then click a .name domain.

On the domain control panel of the .name domain, scroll down to Email Settings.

Click the .name email button.

On the Manage this domain's ".NAME email forward" settings page, the save changes button calls the SetDotNameForwarding command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must have a .name TLD.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                        | Max Size |
| --------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                     | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                               | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| ForwardTo    | Required | Email address to forward to, in the format [\[email protected\]](../cdn-cgi/l/email-protection.html#2a6f674b43466b46434b596a6e45474b4344644b474f045e464e) | 120 |
| ResponseType | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| ForwardTo    | Email address to which email will be forwarded |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTMLor ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

John Doe registers domain john.doe.name, which automatically includes the email address [\[email protected\].](../cdn-cgi/l/email-protection.html#7f151017113f1b101a51111e121a51) The following query forwards email addressed to [\[email protected\]](../cdn-cgi/l/email-protection.html#167c797e78567279733878777b73) to [\[email protected\]](../cdn-cgi/l/email-protection.html#89e3e6e1e7a7ede6ecc9fbecfaece5e5ecfbede6eafaa7eae6e4), and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=setdotnameforwarding&uid=resellid&pw=resellpw
&sld=john.doe&tld=name&[email protected]
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=setdotnameforwarding&uid=resellid&pw=resellpw
&sld=john.doe&tld=name&[email protected]
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=setdotnameforwarding&uid=resellid&pw=resellpw
&sld=john.doe&tld=name&[email protected]
&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <dotnameforwarding>
 <ForwardTo>[email protected]</ForwardTo>
 </dotnameforwarding>
 <Command>SETDOTNAMEFORWARDING</Command>
 <ErrCount>0</ErrCount>
 <Server>Reseller3</Server>
 <Site>enom</Site>
 <IsLockable>False</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
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
DomainRRP: E
ForwardTo: [email protected]
Command: SETDOTNAMEFORWARDING
ErrCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +7.00
ExecTime: 0.016
Done: true
TrackingKey: a8d9633b-8f9c-493c-8c3f-ce90cd64def1
RequestDateTime: 2/5/2015 2:55:52 PM
```
Related Commands
----------------

Forwarding

GetCatchAll

GetDotNameForwarding

GetForwarding

PurchasePOPBundle

SetUpPOP3User