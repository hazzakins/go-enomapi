CertResendFulfillmentEmail
==========================

Resend fulfillment email \(final cert\) for a Symantec, Verisign or GeoTrust certificate.

Usage
-----

Use this command to resend fulfillment email \(final cert\) for a Symantec, Verisign or GeoTrust certificate.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

N/A

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The cert must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                            | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                         | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML             | 4 |
| CertID | Required | ID number for this individual certificate. Retrieve this number using the CertGetCerts command. | 8    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| Success     | Success status for retrieving the authoritative Whois registrant email |
| ApproverEmail | Email address of approver                                    |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query changes approver email for a GeoTrust Certificate, and requests the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=CertResendFulfillmentEmail&UID=resellid
&PW=resellpw&ResponseType=XML&CertID=111
```
```
https://resellertest.enom.com/interface.asp?
command=CertResendFulfillmentEmail&UID=resellid
&PW=resellpw&ResponseType=html&CertID=111
```
```
https://resellertest.enom.com/interface.asp?
command=CertResendFulfillmentEmail&UID=resellid
&PW=resellpw&ResponseType=text&CertID=111
```
In the response, parsed information and an ErrCount value 0 confirm that the query was successful:

```
<?xml version="1.0"?>

 <Success>True</Success>

 <Command>CERTRESENDFULFILLMENTEMAIL</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>10</MaxPeriod>

 <Server>blvdt112</Server>

 <Site>eNom</Site>

 <IsLockable/>

 <IsRealTimeTLD/>

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>2.967</ExecTime>

 <Done>true</Done>

 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>

 <RequestDateTime>3/1/2012 11:29:08 AM</RequestDateTime>

</interface-response>
```