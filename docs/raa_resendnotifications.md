RAA\ResendNotifications
=======================

Resend Registrar Accreditation Agreement \(RAA\) notification email for a contact change.

Usage
-----

Use this command to resend RAA notification email for a contact change.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/control-panel/default.aspx](https://resellertest.enom.com/domains/control-panel/default.aspx)

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                               | Max Size |
| ---------------- | ---------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required  | Account password | 20    |
| ResponseType   | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |
| DomainName | Required\* | Domain name which contact change was made on \(or new registration\) | 272   |
| VerificationCode | Required\* | Verification code that is generated on a contact change         | ? |
| FirstName | Required\* | First name of contact change | 60    |
| LastName     | Required\* | Last name of contact change                       | 60 |
| EmailAddress | Required\* | Email address of contact change | 128   |

\*NOTE: only one \(1\) RAA parameter is required to be passed in. If more than one of the RAA parameters are present in the query string, the first value based on the order below will be processed \(other values will be ignored\):

DomainName

VerificationCode

FirstName, LastName, EmailAddress

Example: if DomainName and VerificationCode parameters are being passed, only DomainName value will be taken.

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| Command | Name of command executed                                     |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query resends RAA notification email for a contact change in various types, and requests the response in given format.

Passing DomainName
------------------

```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&domainname=testdomain.com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&domainname=testdomain.com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&domainname=testdomain.com&responsetype=text
```
Passing VerificationCode
------------------------

```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&verificationcode=1558f096-9202-4b21-81d6-9f34d207514d
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&verificationcode=1558f096-9202-4b21-81d6-9f34d207514d
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&verificationcode=1558f096-9202-4b21-81d6-9f34d207514d
&responsetype=text
```
Passing FirstName, LastName, EmailAddress
-----------------------------------------

```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&firstname=Joe&lastname=Tester
&[email protected]&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&firstname=Joe&lastname=Tester
&[email protected]&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=raa_resendnotification&uid=resellid&pw=resellpw
&firstname=Joe&lastname=Tester
&[email protected]&responsetype=text
```
Output
------

```
<?xml version="1.0" encoding="UTF-8"?>

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
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Success: true

Command: RAA_RESENDNOTIFICATION

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.094

Done: true

TrackingKey: 6b91f22e-701e-4623-b69d-9a8e5d68d2c0

RequestDateTime: 2/5/2015 12:01:22 PM
```
ExecTime=0.016

TrackingKey=4a97e800-6645-4c5d-b1b9-3d8a57d8044b

RequestDateTime=2/5/2015 12:01:48 PM
```