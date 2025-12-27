TP\ResendEmail
==============

Resend the Domain Transfer Request authorization email. This is the email we send to the WhoIs contact on record at the Registry, requesting their authorization to transfer the domain name into the account specified in this query string.

Usage
-----

Use this command to resend the Domain Transfer Request email for a transfer. You would typically use this command when the owner who is losing the domain updates their email address after we have sent an initial authorization email.

This command can only be used between the time we send an initial Domain Transfer Request email, and the time we receive a response.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must be one that is transferring into this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.    | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Success | True indicates the query was successful                             |
| DomainName    | Name of the domain to be redeemed |
| OrderID | Order ID number. Store this number to simplify progress checks.                |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- If you are transferring a domain name between two of your own accounts at different registrars, you can reduce delays by updating your email address in the losing account before you submit a transfer order.
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query resends confirmation emails for the transfer of resellerdocs3.info, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_RESENDEMAIL&uid=resellid&pw=resellpw
&SLD=resellerdocs3&tld=info&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_RESENDEMAIL&uid=resellid&pw=resellpw
&SLD=resellerdocs3&tld=info&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_RESENDEMAIL&uid=resellid&pw=resellpw
&SLD=resellerdocs3&tld=info&responsetype=text
```
In the response, a Success value of True indicates that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
<Success>True</Success>
<Command>TP_RESENDEMAIL</Command>
<Language>en</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>Reseller5</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+03.00</TimeDifference>
<ExecTime>0.15625</ExecTime>
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
Success: True
Command: TP_RESENDEMAIL
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site:
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.047
Done: true
RequestDateTime: 2/9/2015 1:06:38 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success=True
Command=TP_RESENDEMAIL
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
RequestDateTime=2/9/2015 1:07:41 PM
```
Related Commands
----------------

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrderDetail

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResubmitLocked

TP\_SubmitOrder