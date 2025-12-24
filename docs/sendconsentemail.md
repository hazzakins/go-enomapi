SendConsentEmail
================

Send GDPR consent email for a domain.

Usage
-----

Use this command to send consent email for a single domain name.

Availability
------------

All resellers have access to this command.

Constraints
-----------

Only domains owned by reseller or their sub retail accounts.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SendConsentEmail&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Type | Description                                |
| --------------- | ------ | -------------------------------------------------------------------------- |
| command     | string | **SendConsentEmail**                            |
| uid       | string | Your Account ID                              |
| pw       | string | Your API Token                               |
| domainname   | string | Domain Name                                |
| ResponseType  | string | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Success     | boolean | True if email was successfully sent. Otherwise false.                     |
| Command     | | Name of command executed                                     |
| ErrCount     | | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | | True indicates this entire response has reached you successfully.                |

Response Params
---------------

```
<?xml version="1.0" encoding="UTF-8"?>
<interface-response>
  <Success>True</Success>
  <Command>SENDCONSENTEMAIL</Command>
  <APIType>API.NET</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>resellertest</Server>
  <Site>eNom</Site>
  <IsLockable />
  <IsRealTimeTLD />
  <TimeDifference>+0.00</TimeDifference>
  <ExecTime>0.475</ExecTime>
  <Done>true</Done>
  <TrackingKey>ababd976-a29a-4fd8-a309-430dc031eeb2</TrackingKey>
  <RequestDateTime>5/17/2018 10:43:23 AM</RequestDateTime>
  <debug />
</interface-response>
```
Related Commands