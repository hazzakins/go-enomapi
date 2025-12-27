Portal\UpdateAwardedDomains
===========================

This command will update a set list of domains awarded as reseller provisioned.

Usage
-----

Use this command to to update a single domain or a list of domains that have been awarded to your customers through the portal as reseller provisioned and imported into your system. Doing this will remove the domain from being able to be managed by the customer through the portal, and will ensure that they will do so from your system.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not utilized on enom.com as it is intended for partner integration

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- A search term must be passed in.
- At least one item to update must be passed in.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Success     | True indicates that the request was completed successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- Note: This process is not reversible. Once a domain is marked as provisioned, it will be removed from the portal and the end user will no longer be able to manage the domain through the portal in any capacity.

Example
-------

The following query updates 3 domains as "reseller provisioned" and sends the response in XML, HTML, or Text format.

```
http://resellertest.enom.com/interface.asp?
command=Portal_UpdateAwardedDomains&UID=resellid
&PW=resellpw&domainlist=123456,7890123,4567890
&ResponseType=XML
```
```
http://resellertest.enom.com/interface.asp?
command=Portal_UpdateAwardedDomains&UID=resellid
&PW=resellpw&domainlist=123456,7890123,4567890
&ResponseType=html
```
```
http://resellertest.enom.com/interface.asp?
command=Portal_UpdateAwardedDomains&UID=resellid
&PW=resellpw&domainlist=123456,7890123,4567890
&ResponseType=text
```
In the response, the presence of Success with a value of True AND the ErrCount value 0 indicate that the query was successful:

```
<interface-response>
<Success>True</Success>
<Command>Portal_UpdateAwardedDomains</Command>
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
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success: True
Command: PORTAL_UPDATEAWARDEDDOMAINS
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.328
Done: true
TrackingKey: 5c95e3d6-bc78-448e-aadb-a217e799a864
RequestDateTime: 2/4/2015 5:02:20 PM
```
Related Commands
----------------

Portal\_GetAwardedDomains

Portal\_GetDomainInfo

Portal\_GetToken