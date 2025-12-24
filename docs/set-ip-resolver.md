SetIPResolver
=============

Update IP Resolver settings.

Usage
-----

Use this command to set information for the NameMyComputer service.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/IPResolverConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/IPResolverConfig.asp?DomainNameID=152533676)

On the name my computer page, the save changes button calls the SetIPResolver command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword&
paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                 | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                              | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| IPResolverID  | Optional | IP resolver ID                                               | 60 |
| OrigHostName | Optional           | Original host name | 60    |
| HostName    | Required | Name for your host                                             | 60 |
| ClientUserID | Required           | ICQ\# or UserID | 60    |
| ServiceID    | Required | City to build the map for                                          | 60 |
| OfflineURL | Optional           | Offline URL | 78    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                            | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| ip-resolver-id | Current ID if one exists.                                    |
| host-name    | Name for your host. |
| client-user-id | Current ICQ\# or UserID.                                    |
| service-id   | Current service ID \("1"\). |
| SetIPResolver | Success of this query. Response is Successful or Failed.                    |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sets information for the Name My Computer service and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SetIPResolver&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ClientUserID=12345678
&HostName=home&ServiceID=1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SetIPResolver&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ClientUserID=12345678
&HostName=home&ServiceID=1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SetIPResolver&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ClientUserID=12345678
&HostName=home&ServiceID=1&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <SetIPResolver>
 <domainname sld="resellerdocs" tld="com" id="152533676">resellerdocs.com</domainname>
 <IPResolver>
  <ip-resolver-id />
  <orig-host-name />
  <host-name>home</host-name>
  <client-user-id>
  <![CDATA[ 12345678 ] ]>
  </client-user-id>
  <service-id>1</service-id>
  <offine-url>
  <![CDATA[ ] ]>
  </offine-url>
  <data-errors />
 </IPResolver>
 <SetIPResolver>Successful</SetIPResolver>
 </SetIPResolver>
 <Command>SETIPRESOLVER</Command>
 <ErrCount>0</ErrCount>
 <Server>Dev Workstation</Server>
 <Site>eNom</Site>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```bash
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
domainname: resellerdocs.com
ip-resolver-id: -1
orig-host-name:
host-name: home
client-user-id: 12345678
service-id: 1
offline-url:
data-errors:
SetIPResolver: Successful
Command: SETIPRESOLVER
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
ExecTime: 0.547
Done: true
TrackingKey: fbdaddec-4371-49dc-b877-88d2f2773773
RequestDateTime: 2/5/2015 3:07:11 PM
```
Related commands
----------------

GetIPResolver

ServiceSelect