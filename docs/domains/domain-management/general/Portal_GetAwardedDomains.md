Portal\GetAwardedDomains
========================

Retrieve a list of an account’s domains that are in Expired, Redemption Grace Period \(RGP\), and Extended RGP status.

Usage
-----

This command will get a list of all domains that have been awarded to the resellers watchlist users. Only domains that have not yet been marked as provisioned will be returned by this command.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The AuthQuestionAnswer value must be correct.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=Portal_GetAwardedDomains&uid=(Required)&PW=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status          | Description |
| --------------- | ------ | ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required         | Portal\_GetAwardedDomains |
| UID | string | Required | Your Account ID                                                                                                                           |
| PW       | string | Required         | Account password |
| RecordCount | int  | Optional; default is 25 | The maximum number of records to return. maximum records is capped at 100                                                                                             |
| RecordOption  | int | Optional; default is 0  | Determine the type of domains you will get back; Permitted values are: 0 - Only return domains that a Reseller has not provisioned 1 - Only return domains that a Reseller has already provisioned 2 - Return all data, even ones that have already been awarded |
| StartPosition | int  | Optional; default is 1 | Optionally set which value to start at. This allows you to optionally page through the results.                                                                                  |
| Tld       | string | Optional; default is NULL | Optionally allows you to filter the results only to domains with the specified TLD |
| DomainNameID | string | Optional; default is NULL | Optionally allows you to filter the results only to the single domain with the specified Domain Name ID.                                                                              |
| ResponseType  | string | Optional         | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

Note: NewStartPosition is calculated by total results, start position and records to return. It is advised that if you are paging through the results that you do the calculation yourself with the values that make sense for your application.

Note: ForeignLoginID is the "LoginID" that you pass along to the PORTAL\_GETTOKEN API Command

| Output Parameter | Type  | Description |
| ----------------------- | ------- | ------------------------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| DomainCount | string | Number of domains returned in this response |
| TotalDomainCount | string | Total number of domains available to be returned |
| NextStartPosition | string | Next starting position to use for those paging through the results. |
| DomainName | string | The Domain Name |
| EmailAddress | string | The Email Address of the portal user |
| ExpirationDate | string | The domains expiration date. |
| RegisterDate | string | The domains registration date. |
| ForeignLoginID | string | The login ID on your system that registered the domain. |
| PortalDomainId | string | The Portal Domain Name ID of the domain. |
| RegisterPrice | string | The users registration price. This is the price that they paid. |
| RenewPrice | string | The users registration price. This is the price that they would pay to renew the domain at your current pricing. |
| RegistrationPeriod | string | The number of years the domain is registered for |
| ResellerProvisioned | boolean | Boolean value indicating if you have marked this domain as provisioned into your system. |
| ResellerProvisionedDate | string | If the domain is marked as provisioned, this is the date that it was done so. |
| RegistrationStatus | string | Registration status of the domain. Example: Registered, Expired, etc. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
http://resellertest.enom.com/interface.asp?
command=Portal_GetAwardedDomains&UID=resellid
&PW=resellpw&RecordCount=1&RecordOption=1&responsetype=xml
```
```
http://resellertest.enom.com/interface.asp?
command=Portal_GetAwardedDomains&UID=resellid
&PW=resellpw&RecordCount=1&RecordOption=1&ResponseType=html
```
```
http://resellertest.enom.com/interface.asp?
command=Portal_GetAwardedDomains&UID=resellid
&PW=resellpw&RecordCount=1&RecordOption=1&ResponseType=text
```
```
<interface-response>
 <Domains>
 <Domain>
  <DomainName>mynewdomain.ninja</DomainName>
  <EmailAddress>[email protected]</EmailAddress>
  <ExpirationDate>5/28/2019</ExpirationDate>
  <RegisterDate>5/28/2014</RegisterDate>
  <ForeignLoginId>Fred Penner</ForeignLoginId>
  <PortalDomainId>123456789</PortalDomainId>
  <RegisterPrice>19.99</RegisterPrice>
  <RenewPrice>19.99</RenewPrice>
  <RegistrationPeriod>5</RegistrationPeriod>
  <ResellerProvisioned>1</ResellerProvisioned>
  <ResellerProvisionedDate>5/29/2014</ResellerProvisionedDate>
  <IsPremium>0</IsPremium>
  <RegistrationStatus>Registered</RegistrationStatus>
 </Domain>
 <DomainCount>1</DomainCount>
 <TotalDomainCount>10</TotalDomainCount>
 <NextStartPosition>2</NextStartPosition>
 </Domains>
 <Success>True</Success>
 <Command>Portal_GetAwardedDomains</Command>
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
 <ExecTime>0.750</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>9efbb8e7-1cc3-4f2b-bbc2-cbe2517522aa</TrackingKey>
 <RequestDateTime>5/30/2014 2:51:56 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>DomainCount=</ STRONG>0<br>
<STRONG>TotalDomainCount=</ STRONG>0<br>
<STRONG>NextStartPosition=</ STRONG>0<br>
<STRONG>Success=</ STRONG>True<br>
<STRONG>Command=</ STRONG>PORTAL_GETAWARDEDDOMAINS<br>
<STRONG>APIType=</ STRONG>API.NET<br>
<STRONG>Language=</ STRONG>eng<br>
<STRONG>ErrCount=</ STRONG>0<br>
<STRONG>ResponseCount=</ STRONG>0<br>
<STRONG>MinPeriod=</ STRONG>1<br>
<STRONG>MaxPeriod=</ STRONG>10<br>
<STRONG>Server=</ STRONG>sjl0vwresell_t<br>
<STRONG>Site=</ STRONG>eNom<br>
<STRONG>IsLockable=</ STRONG><br>
<STRONG>IsRealTimeTLD=</ STRONG><br>
<STRONG>TimeDifference=</ STRONG>+0.00<br>
<STRONG>ExecTime=</ STRONG>0.172<br>
<STRONG>Done=</ STRONG>true<br>
<STRONG>TrackingKey=</ STRONG>c916b392-b9ed-47b4-8dd5-e1b1fa888971<br>
<STRONG>RequestDateTime=</ STRONG>2/4/2015 4:47:11 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainCount=0
TotalDomainCount=0
NextStartPosition=0
Success=True
Command=PORTAL_GETAWARDEDDOMAINS
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
TrackingKey=c715fe2e-136c-426e-a4a3-149dc9393806
RequestDateTime=2/4/2015 4:47:48 PM
```
Related Commands
----------------

Portal\_GetToken

Portal\_UpdateAwardedDomains