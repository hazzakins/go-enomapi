SetDNSHost
==========

Dynamically updates the IP address of the host address in our name server records.

Usage
-----

Use this command to update the IP address of a domain Host Record that does not have a static IP address. It is recommended to use this command in a secure mode \(i.e. "https;//" instead of "http://"\)

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain name must have a password.

> ### Important Note\!
>
>
>
> The Zone parameter must be set as an IP address.
>
> This command ONLY works for Zone/Host records of type "A".

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetDNSHost&uid={YourAccountID}&pw={YourApiToken}&zone={required value}&domainpassword={required value}&address={optional value}&responsetype={xml}
```
| Input Parameter | Type | Description                                                                                                               |
| --------------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | SetDNSHost                                                                                                               |
| uid       | string | Your Account ID                                                                                                             |
| pw       | string | Your API Token                                                                                                             |
| zone      | string | Required - The host and domain name that you want to update in the DNS. For example, [www.resellerdocs.com](http://www.resellerdocs.com/)                                              |
| domainpassword | string | Required - Password for managing the domain. A domain that uses the SetDNSHost command must have a password.                                                             |
| address     | string | Optional - The IP address to set the DNS record to. If omitted, the IP you are coming from \(as our server sees it\) is used. For example, if you are connecting to our server through a proxy, the proxy server's IP will be used. |
| responsetype  | string | Optional - Format of response. Permitted values are Text \(default\), HTML, or XML.                                                                         |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                            |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------- |
| command     | string | Name of command executed                                     |
| ErrCount     | int | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\) values. |
| err\{X\}     | string | Error messages explaining the failure. These can be presented as-is back to the client.    |
| done       | boolean | True value indicates this entire response has reached you successfully.             |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, set the "ResponseType=" flag to either "HTML" or "XML" in your request.

Check the return parameter "ErrCount". If it is greater than 0, the transaction has failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example of Output
-----------------

The following query uses a secure server to set the host record for "www" on the domain name "resellerdocs.com".

Please note that using this specific address ip of 127.0.0.1 will prevent the "www" host record from resolving in a web browser.

```
https://resellertest.enom.com/interface.asp?command=SetDNSHost&uid=resellertest&pw=resellertest&zone=www.resellerdocs.com&address=127.0.0.1&DomainPassword=tester&responsetype={text, XML, or HTML}
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
IP=127.0.0.1
Command=SETDNSHOST
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.016
Done=true
TrackingKey=f0322f7d-3948-4a76-a5f7-50dd36dd491c
RequestDateTime=2/12/2015 1:47:18 PM
```
```
<?xml version="1.0" ?>
<interface-response>
 <IP>127.0.0.1</IP>
 <Command>SETDNSHOST</Command>
 <ErrCount>0</ErrCount>
 <Server>Reseller5</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWAPI08<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>IP: </STRONG>127.0.0.1<BR /><STRONG>Command: </STRONG>SETDNSHOST<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl0vwresell_t<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.063<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>002c4aca-e979-4b1e-ba21-4b9e2c788bc6<BR /><STRONG>RequestDateTime: </STRONG>2/12/2015 1:46:23 PM<BR /></BODY></HTML>
```
Related Commands
----------------

GetHosts

GetRegHosts

SetHosts