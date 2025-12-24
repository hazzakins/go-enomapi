ModifyPOP3
==========

Modify POP account password.

Usage
-----

Use this command to change the password of one or more POP3 mail accounts in a domain.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/POPConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/POPConfig.asp?DomainNameID=152533676)

On the Domain Name Maintenance page, in the Edit POP mail accounts box, typing a new password and clicking the add to cart button to save modifications calls the ModifyPOP3 command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The POP3 service need not be active for the password change to be successful.
- The values of UserName in the query must exist for the domain.
- The value for EmailCount must match the number of UserNames and Passwords in the query.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter       | Status | Description                                                                                                                                                                                                                                                                                | Max Size |
| --------------------------- | ---------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID             | Required | Account login ID                                                                                                                                                                                                                                                                             | 20 |
| PW | Required                 | Account password | 20    |
| SLD             | Required | Second-level domain name \(for example, enom in enom.com\)                                                                                                                                                                                                                                                       | 63 |
| TLD | Required                 | Top-level domain name \(extension\) | 15    |
| EmailCount         | Required | Number of accounts on this domain to update.                                                                                                                                                                                                                                                               | 2 |
| UserNameX X=1 to EmailCount | Required                 | POP3 user name to update the password on. | 30    |
| IsAdminX X=1 to EmailCount | Either IsAdminX or PasswordX is Required | They may be used together. Should UserNameX have administrative privileges over all mailboxes associated with this domain name? Administrative privileges allow this user to reset the password for mailboxes, and delete mailboxes. Permitted values are: 0 Regular user \(can change the password for their own mailbox but no one else’s\) 1 Administrator \(when this user logs on to Webmail and goes to the admin page, he or she can see all mailboxes for this domain name, can reset passwords for any or all mailboxes, and can delete any or all mailboxes\) | 1 |
| PasswordX X=1 to EmailCount | Either IsAdminX or PasswordX is Required | They may be used together. New password for UserNameX’s mailbox. | 50    |
| ResponseType        | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                                                                                                                                                                                                                                           | 4 |

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
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests that two POP3 passwords change in resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=modifypop3&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&emailcount=2&username1=john
&password1=johnpw&username2=jane
&password2=janepw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=modifypop3&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&emailcount=2&username1=john
&password1=johnpw&username2=jane
&password2=janepw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=modifypop3&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&emailcount=2&username1=john
&password1=johnpw&username2=jane
&password2=janepw&responsetype=text
```
The response is as follows:

```
<interface-response>
 <Success>True</Success>
 <AccountsModified>2</AccountsModified>
 <AccountsFailed>0</AccountsFailed>
 <Command>MODIFYPOP3</Command>
 <APIType>API.NET</APIType>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl0vwresell_t1</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.016</ExecTime>
 <Done>true</Done>
 <TrackingKey>1a9002ec-ad7b-449d-ab41-a53a4788a80c</TrackingKey>
 <RequestDateTime>2/11/2015 5:25:11 PM</RequestDateTime>
 <debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Success: True
AccountsModified: 2
AccountsFailed: 0
Command: MODIFYPOP3
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
ExecTime: 0.016
Done: true
TrackingKey: e9ac093f-59a4-43cf-b72a-e94b8343acb0
RequestDateTime: 2/11/2015 5:28:24 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

Forwarding

GetForwarding

GetMailHosts

GetPOP3

PurchasePOPBundle

SetUpPOP3User