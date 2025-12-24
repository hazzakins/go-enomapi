GetIPResolver
=============

Get IP Resolver settings.

Usage
-----

Use this command to retrieve the NameMyComputer settings for a domain name.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

On the domain control panel page, the GetIPResolver command provides the content of the name my computer box.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- Other constraints may apply depending on the services being enabled.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------- | -------- |
| UID       | Required | Your Account ID                              | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)        | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| IP-Resolver-ID | Currently 1240 for ICQ.                                     |
| Host-Name    | Host name for the service. |
| Client-Service | Currently only service type 1 for ICQ.                             |
| Client-User-ID | Client ID \(ICQ \#\). |
| ID | Option ID for the client service.                                |
| Name       | Currently returns only ICQ. |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests IP resolver information for the domain resellerdocs.com and sends the response in XML, HTML and text formats:

```
https://resellertest.enom.com/interface.asp?
command=getipresolver&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getipresolver&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getipresolver&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
```
<interface-response>
 <GetIPResolver>
 <domainname sld="resellerdocs" tld="com" id="152533676">resellerdocs.com</domainname>
 <configuration>
  <ipresolver>
  <ip-resolver-id>1243</ip-resolver-id>
  <host-name>home</host-name>
  <client-service>1</client-service>
  <client-user-id>12345678</client-user-id>
  </ipresolver>
 </configuration>
 <options>
  <clients>
  <client>
   <id>1</id>
   <name>ICQ</name>
   <short-name>icq</short-name>
   <description>ICQ</description>
  </client>
  </clients>
 </options>
 </GetIPResolver>
 <Command>GETIPRESOLVER</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>0.203</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/8/2011 5:04:05 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
ip-resolver-id: 1243
host-name: home
client-service: 1
client-user-id: 12345678
id: 1
name: ICQ
short-name: icq
Command: GETIPRESOLVER
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.125
Done: true
RequestDateTime: 2/4/2015 10:55:01 AM
```
Related Commands
----------------

AddToCart

DisableServices

GetWPPSInfo

PurchaseServices

ServiceSelect

SetRenew

SetResellerServicesPricing

UpdateAccountPricing