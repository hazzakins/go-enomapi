GetReport
=========

Get a report on a user’s account.

Usage
-----

Use this command to retrieve an overview of the domains in an account: domain names, name server status, registration status, and expiration date.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[http://resellertest.enom.com/domains/DomainManager.asp?tab=iown](http://resellertest.enom.com/domains/DomainManager.asp?tab=iown)

The GetReport command works in the background of the my domains page. Each of the four tabs on the page displays one category of domain names returned by the GetReport query.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| Done        | True indicates this entire response has reached you successfully. |
| FName | First \(use\) name of the account registrant.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| LName       | Last \(family\) name of the account registrant |
| Party | Party identification number. Format is 32 hexadecimal digits, hyphenated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| DN         | Domain name parameter. |
| DomName | Domain name attribute.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| NSStatus      | Name server status. Value yes indicates our name servers. Value no indicates not our name servers. |
| RegistrationStatus | Registration status in this account. Return values are: Deleted Domain has been deleted from registrar’s database. Expired Domain registration has expired but has not been deleted from registrar’s database \(domain is within “grace period” for reactivaton\). Expired Transfers Domain that is being transferred to another registrar in expired state. Hosted Domain that is hosted by us but registered with another registrar. Hosted Deleted Domain that was hosted by us but registered elsewhere, its registration has expired, and it has been deleted from our database. Hosted Expired Domain that is hosted by us but registered elsewhere, its registration has expired, but it is still present in our database. Hosted Transfers Domain that is hosted by us but registered elsewhere, and the registration is transferring to another registrar. Keyword Keyword, an obsolete feature. Registered Domain that is registered by us. Transferred away Hosted and Registered Domain that was registered and hosted by us, and both the registration and hosting services are transferring away from us. |
| ExpirationDate   | Date on which this domain registration expires. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests a report on account resellid, and requests the response in XML format:

```
https://resellertest.enom.com/interface.asp?
command=getreport&uid=resellid&pw=resellpw
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getreport&uid=resellid&pw=resellpw
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getreport&uid=resellid&pw=resellpw
&responsetype=text
```
In the response, the ReportDetail section lists the domains in account resellid:

```
<?xml version="1.0" ?>

 <FName>John</FName>

 <LName>Doe</LName>

 <Party>{BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}</Party>

  <dn DomName="resellerdocs.com" nsstatus="No" registrationstatus="Hosted"

ExpirationDate="07/15/03" />

  <dn DomName="resellerdocs2.net" nsstatus="Yes" registrationstatus="Registered"

ExpirationDate="06/25/04" />

  <dn DomName="resellerdocs3.info" nsstatus="Yes" registrationstatus="Registered"

 </ReportDetail>

 <Command>GETREPORT</Command>

 <ErrCount>0</ErrCount>

 <Server>ResellerTest</Server>

 <Site>enom</Site>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

FName: Rosh

LName: Bach

Party: {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}

Command: GETREPORT

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod:

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

IsLockable:

IsRealTimeTLD:

TimeDifference: +0.00

ExecTime: 0.766

Done: true

RequestDateTime: 2/4/2015 12:23:11 PM
```
ExecTime=0.234

RequestDateTime=2/4/2015 12:23:37 PM
```