SetUpPOP3User
=============

Set up a POP3 mailbox for your domain.

Usage
-----

Use this command when you have purchased a POP3 email pak \(see PurchasePOPBundle\) and want to set up mailboxes for your domain.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The AuthQuestionAnswer value must be correct.
- The domain name must belong to this account.
- A POP mail pak must have been purchased for the domain and the correct BundleID supplied.
- There must be at least one mailbox available in the email pak \(BundleID\) referenced.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetUpPOP3User&uid=(Required)&PW=(Required)&SLD=(Required)&TLD=(Required)&BundleID=(Required)&UserNameX=(Required)&PasswordX=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | SetUpPOP3User |
| UID | string | Required | Your Account ID                                                                                                                                                                                                                                                                                                                        |
| PW       | string | Required | Account password |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                                                                                                                                                                                                                                                                                 |
| TLD       | string | Required | Top-level domain name \(extension\) |
| BundleID | string | Required | POP pak ID number to add the user to \(must have already been purchased\). The order confirmation for purchasing the bundle contains the Bundle ID, or you can retrieve it using the GetPOP3 command.                                                                                                                                                                                                                            |
| UserNameX    | string | Required | The user name to set up \(email address will be [\[email protected\]](../cdn-cgi/l/email-protection.html#9cc9eff9eed2fdf1f9c4dceff0f8b2e8f0f8)\). UserNameX is indexed, starting with UserName1. You can configure multiple user names in a single query string |
| IsAdminX | string | Optional | Should UserNameX have administrative privileges over all mailboxes associated with this domain name. Administrative privileges allow this user to reset the password for mailboxes, and delete mailboxes. Permitted values are: 0 Regular user \(can change the password for their own mailbox but no one else’s\) 1 Administrator \(when this user logs on to Webmail and goes to the admin page, he or she can see all mailboxes for this domain name, can reset passwords for any or all mailboxes, and can delete any or all mailboxes\). IsAdminX is indexed, starting with IsAdmin1. You can configure multiple user names in a single query string |
| PasswordX    | string | Required | POP3 password. PasswordX is indexed, starting with Password1. You can configure multiple user names in a single query string |
| DisplayNameX | string | Optional | User's full name, as it appears in return name in Web mail. DisplayNameX is indexed, starting with DisplayName1. You can configure multiple user names in a single query string                                                                                                                                                                                                                                       |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SetupPop3user&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&BundleID=5134
&UserName1=james&Password1=resellpw
&DisplayName1=James+Doe&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SetupPop3user&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&BundleID=5134
&UserName1=james&Password1=resellpw
&DisplayName1=James+Doe&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=SetupPop3user&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&BundleID=5134
&UserName1=james&Password1=resellpw
&DisplayName1=James+Doe&ResponseType=text
```
```
<interface-response>
 <Success>True</Success>
 <AccountsCreated>1</AccountsCreated>
 <AccountsFailed>0</AccountsFailed>
 <Command>SETUPPOP3USER</Command>
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
 <ExecTime>8.547</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>37bb5652-5250-4483-a877-368f2eb7674f</TrackingKey>
 <RequestDateTime>12/12/2011 12:57:20 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Success: </ STRONG>True<br>
<STRONG>AccountsCreated: </ STRONG>1<br>
<STRONG>AccountsFailed: </ STRONG>0<br>
<STRONG>Command: </ STRONG>SETUPPOP3USER<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>1<br>
<STRONG>Err1: </ STRONG>Invalid Item specified<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.172<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>6d72b4a4-8558-4c99-b222-3b00c0f3926d<br>
<STRONG>RequestDateTime: </ STRONG>2/5/2015 3:40:01 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Success=True
AccountsCreated=1
AccountsFailed=0
Command=SETUPPOP3USER
APIType=API.NET
Language=eng
ErrCount=1
Err1=Invalid Item specified
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.094
Done=true
TrackingKey=5715a3fd-e803-4105-a57e-75ebbb9231d1
RequestDateTime=2/5/2015 3:41:19 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

Forwarding

GetCatchAll

GetDotNameForwarding

GetForwarding

GetMailHosts

GetPOP3

GetPOPForwarding

ModifyPOP3

PurchasePOPBundle

SetDotNameForwarding