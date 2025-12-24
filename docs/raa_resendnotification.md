RAA\ResendNotification
======================

Resend Registrar Accreditation Agreement \(RAA\) notification email for a contact change.

Usage
-----

Use this command to resend RAA notification email for a contact change.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=RAA_ResendNotification&uid=YourAccountID&pw=YourApiToken&DomainName={Required}&VerficationCode={Required}&FirstName={Required}&LastName={Required}&EmailAddress={Required}&responsetype={xml, HTML, or Text}
```
| Input Parameter | Type | Status   | Description |
| ---------------- | ------ | ---------- | ------------------------------------------------------------------------- |
| command     | string | Required  | ModifyNS |
| uid | string | Required | Your Account ID                              |
| pw        | string | Required  | Your API Token |
| DomainName | string | Required | Domain name which contact change was made on \(or new registration\)   |
| VerificationCode | string | Required  | Verification code that is generated on a contact change |
| FirstName | string | Required | First name of contact change                       |
| LastName     | string | Required  | Last name of contact change |
| EmailAddress | string | Required | Email address of contact change                      |
| ResponseType   | string | .Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| command | string | Name of command executed |
| errcount | int  | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| errX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| done | string | True value indicates this entire response has reached you successfully. |

Example Output
--------------

XML Format example:

The following query changes the name servers for "resellerdocs.com" to "ns1.name-services.com", "ns2.name-services.com" and requests the response in either XML, HTML, or Text format:

```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&domainname=testdomain.com&responsetype=xml
```
```
<?xml version="1.0" encoding="UTF-8"?>
<interface-response>
 <Success>true</Success>
 <Command>RAA_RESENDNOTIFICATION</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>blvdt224</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.368</ExecTime>
 <Done>true</Done>
 <TrackingKey>b708ed6c-7ba1-427d-a467-3015b4373e5b</TrackingKey>
 <RequestDateTime>12/4/2013 8:25:02 AM</RequestDateTime>
</interface-response>
```
```
URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML>
<BODY>
<STRONG>Success: </STRONG>true<BR />
<STRONG>Command: </STRONG>RAA_RESENDNOTIFICATION<BR />
<STRONG>APIType: </STRONG>API.NET<BR />
<STRONG>Language: </STRONG>eng<BR />
<STRONG>ErrCount: </STRONG>0<BR />
<STRONG>ResponseCount: </STRONG>0<BR />
<STRONG>MinPeriod: </STRONG>1<BR />
<STRONG>MaxPeriod: </STRONG>10<BR />
<STRONG>Server: </STRONG>sjl0vwresell_t1<BR />
<STRONG>Site: </STRONG>eNom<BR />
<STRONG>IsLockable:</STRONG><BR />
<STRONG>IsRealTimeTLD:</STRONG><BR />
<STRONG>TimeDifference: </STRONG>+0.00<BR />
<STRONG>ExecTime: </STRONG>0.094<BR />
<STRONG>Done: </STRONG>true<BR />
<STRONG>TrackingKey: </STRONG>6b91f22e-701e-4623-b69d-9a8e5d68d2c0<BR />
<STRONG>RequestDateTime: </STRONG>2/5/2015 12:01:22 PM<BR />
</BODY>
</HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Success=true
Command=RAA_RESENDNOTIFICATION
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.016
Done=true
TrackingKey=4a97e800-6645-4c5d-b1b9-3d8a57d8044b
RequestDateTime=2/5/2015 12:01:48 PM
```
Related Commands
----------------

RAA\_GetInfo

RPT\_GetReport