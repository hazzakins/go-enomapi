RAA\GetInfo
===========

Get Registrar Accreditation Agreement \(RAA\) information for a contact change.

Usage
-----

Use this command to get RAA information for a contact change.

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

| Output Parameter  | Description |
| ------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                  |
| Done        | True indicates this entire response has reached you successfully. |
| Domain | Domain name                                                         |
| VerificationCode  | Verification code |
| VerificationStatus | Verification status                                                     |
| ValidationType   | Validation type |
| FirstName | First name                                                          |
| LastName      | Last name |
| EmailAddress | Email address                                                        |
| IsSuspended     | Is this domain suspended? |
| IsPendingSuspension | Is this domain in pending suspension?                                            |
| SuspensionDate   | Suspension date |
| ToBeSuspendedDate | Following suspension date                                                  |
| ResultText     | Result description |
| EmailsSent | List of notification emails sent out to the domain owner\(s\). Information includes destination email adrress and sent date |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above

Example
-------

The following query gets RAA information for a contact change in various types, and requests the response in XML format.

Passing DomainName
------------------

You will only get back what information you passed in and was used for the API Call. If you send in DomainName you would only get back the domain name in the output, the other parameters if sent are cleared and not used or returned.

```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&domainname=testingdomain.com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&domainname=testingdomain.com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&domainname=testingdomain.com&responsetype=text
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Domain: resellerdocs.com

VerificationStatus: Unverified

ValidationType:

IsSuspended: False

IsPendingSuspension: False

SuspensionDate:

ToBeSuspendedDate:

ResultText:

EmailsSent:

Command: RAA_GETINFO

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

ExecTime: 0.203

Done: true

TrackingKey: 18922575-9778-4393-b20d-8152f31d45fe

RequestDateTime: 2/5/2015 11:53:34 AM
```
ResellerKey0: 20B0FF00-18C2-4425-9198-5

HostPrice0: 11.21

CCPrice0: 27.69

InfoPrice0: 11.21

BizPrice0: 7.14

TvPrice0: 43.14

WsPrice0: 16.36

BzPrice0: 27.69

NuPrice0: 27.69

UsPrice0: 11.21

PictureURL0:

ContactEmail0:

CompanyName0: RRTest

Referrer0:

NameMyPhone0: False

NameMyComputer0: False

NameMyMap0: False

SiteBuilder0: False

StyleSheetNum0: 1

success: False

Command: GETWEBHOSTINGALL

ExecTime: 0.031

TrackingKey: df057da9-adb6-4ff2-9732-bea505633b60

RequestDateTime: 2/4/2015 12:58:30 PM
```
;Machine is SJL0VWRESELL_T

ExecTime=0.141

TrackingKey=bd65d6a1-c8bf-442c-a894-e3ae8ddeb422

RequestDateTime=2/5/2015 11:54:26 AM
```
Passing VerificationCode
------------------------

You will only get back what information you passed in and was used for the API Call. If you send in VerificationCode you would only get back the verification code in the output, the other parameters if sent are cleared and not used or returned.

```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&verificationcode=1558f096-9202-4b21-81d6-9f34d207514d
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&verificationcode=1558f096-9202-4b21-81d6-9f34d207514d
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&verificationcode=1558f096-9202-4b21-81d6-9f34d207514d
&responsetype=text
```
```
<?xml version="1.0" encoding="UTF-8"?>

<interface-response>
 <Domain />
 <VerificationCode>1558f096-9202-4b21-81d6-9f34d207514d
					</VerificationCode>
 <VerificationStatus>Pending Suspension</VerificationStatus>
 <ValidationType>New Domain</ValidationType>
 <FirstName />
 <LastName />
 <EmailAddress />
 <IsSuspended>False</IsSuspended>
 <IsPendingSuspension>True</IsPendingSuspension>
 <SuspensionDate/>
 <ToBeSuspendedDate>12/18/2013 7:35:57 AM</ToBeSuspendedDate>
 <ResultText/>
 <EmailsSent>
 <Email>
  <To> [email protected]</To>
  <SendDate>12/3/2013 7:45:00 AM</SendDate>
 </Email>
 <Email>
  <To> [email protected]</To>
  <SendDate>12/4/2013 12:45:00 PM</SendDate>
 </Email>
 </EmailsSent>
 <Command>RAA_GETINFO</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>1.031</ExecTime>
 <Done>true</Done>
 <TrackingKey>9769b21b-5e3f-42fb-88d2-e470f4b4fa5f</TrackingKey>
 <RequestDateTime>12/4/2013 7:57:44 AM</RequestDateTime>
</interface-response>
```
VerificationCode: 1558f096-9202-4b21-81d6-9f34d207514d

VerificationStatus: Invalid Verification code

Server: sjl0vwresell_t

ExecTime: 0.047

TrackingKey: f79ab093-6a43-4d26-8a5a-98e350487be2

RequestDateTime: 2/5/2015 11:55:18 AM
```
ExecTime=0.016

TrackingKey=641b1832-e858-4621-b3be-b9510d18388d

RequestDateTime=2/5/2015 11:55:53 AM
```
Passing FirstName, LastName, EmailAddress
-----------------------------------------

You will only get back what information you passed in and was used for the API Call. For example if you send in FirstName, LastName, EmailAddress you would only get back the firstname, lastname, emailaddress in the output, the other parameters if sent are cleared and not used or returned.

```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&firstname=Joe&lastname=Tester
&[email protected]&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&firstname=Joe&lastname=Tester
&[email protected]&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=raa_getinfo&uid=resellid&pw=resellpw
&firstname=Joe&lastname=Tester
&[email protected]&responsetype=text
```
<interface-response>
<Domain />
<VerificationCode />
<VerificationStatus>Pending Suspension</VerificationStatus>
<ValidationType>New Domain</ValidationType>
<FirstName>Joe</FirstName>
<LastName>Tester</LastName>
<EmailAddress>[email protected]</EmailAddress>
<IsSuspended>False</IsSuspended>
<IsPendingSuspension>True</IsPendingSuspension>
<SuspensionDate/>
<ToBeSuspendedDate>12/18/2013 7:35:57 AM</ToBeSuspendedDate>
<ResultText/>
<EmailsSent>
<Email>
<To> [email protected]</To>
<SendDate>12/3/2013 7:45:00 AM</SendDate>
</Email>
<Email>
<To> [email protected]</To>
<SendDate>12/4/2013 12:45:00 PM</SendDate>
</Email>
</EmailsSent>
<Command>RAA_GETINFO</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>1.031</ExecTime>
<Done>true</Done>
<TrackingKey>9769b21b-5e3f-42fb-88d2-e470f4b4fa5f</TrackingKey>
<RequestDateTime>12/4/2013 7:57:44 AM</RequestDateTime>
</interface-response>
```
FirstName: Joe

LastName: Tester

EmailAddress: [email protected]

ExecTime: 44.188

TrackingKey: 0b8210c7-7542-44ab-9f7f-0b06215f91b3

RequestDateTime: 2/5/2015 11:57:24 AM
```
[email protected]

ExecTime=44.907

TrackingKey=51ffd83d-37f8-4968-96ee-ff6ecef24c3a

RequestDateTime=2/5/2015 11:58:37 AM
```