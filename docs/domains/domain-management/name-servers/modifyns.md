ModifyNS
========

Modify name servers for a domain name.

Usage
-----

Use this command if you want to change the name servers that a domain is currently pointing to.

ModifyNS updates DNS records at the Registry and in the registrar’s database. "ModifyNSHosting", a similar command, redirects to another set of name servers without updating the Registry records.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The query must pass name servers that are registered at TLD Registry, where supported.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=ModifyNS&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&usedns=default&responsetype={xml, HTML, or Text}
```
```
https://resellertest.enom.com/interface.asp?command=ModifyNS&uid=YourAccountID&pw=YourApiToken&ns1=ns1.example.com&ns2=ns2.example.com&responsetype={xml, HTML, or Text}
```
> ### Notes
>
>
>
> To set name servers to use eNom’s DNS, set the "UseDNS=default" parameter and do not pass "NSX=\{YourNameServer\}" name servers.
>
>
>
> To setname servers to your own name servers, set "NSX=\{YourNameServer\}" and do not pass "UseDNS=default". You can set up to 12 of your own name servers.

| Input Parameter       | Type | Status                | Description |
| ---------------------------- | ------ | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command           | string | Required               | ModifyNS |
| uid | string | Required | Your Account ID                                                                                                                                   |
| pw              | string | Required               | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                              |
| TLD             | string | Required               | Top-level domain name \(extension\) |
| usedns | string | Required \(Either UseDNS or NSX\) | Use this parameter to switch to our name servers. Permitted value is "UseDNS=Default"                                                                                                |
| NSX X=1 \(To maximum of 12\) | string | Required \(Either UseDNS or NSX\) - | Use this parameter to switch to custom name servers or no name servers. Maximum of 12 name servers can be set. To designate custom name servers, supply the use names, for example, "NS1=ns1.name-services.com" To designate no name servers, supply an empty parameter: "NS1=" |
| ResponseType | string | .Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML                                                                                                      |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| command     | string | Name of command executed                                     |
| errcount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| errX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| done       | string | True value indicates this entire response has reached you successfully.             |

Example Output
--------------

XML Format example:

The following query changes the name servers for "resellerdocs.com" to "ns1.name-services.com", "ns2.name-services.com" and requests the response in either XML, HTML, or Text format:

```
https://resellertest.enom.com/interface.asp?command=modifyns&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ns1=ns1.name-services.com&ns2=ns2.name-ervices.com&responsetype=xml
```
```
<?xml version="1.0" ?>0" ?>
<interface-response>
 <RRPCode>200</RRPCode>
 <RRPText>Command completed successfully</RRPText>
 <Command>MODIFYNS</Command>
 <ErrCount>0</ErrCount>
 <Server>ResellerTest</Server>
 <Site>enom</Site>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML>
<BODY>
<STRONG>RRPCode: </STRONG>200<BR />
<STRONG>RRPText: </STRONG>Command completed successfully<BR />
<STRONG>Command: </STRONG>MODIFYNS<BR />
<STRONG>APIType: </STRONG>API.NET<BR />
<STRONG>Language: </STRONG>eng<BR />
<STRONG>ErrCount: </STRONG>0<BR />
<STRONG>ResponseCount: </STRONG>0<BR />
<STRONG>MinPeriod: </STRONG>1<BR />
<STRONG>MaxPeriod: </STRONG>10<BR />
<STRONG>Server: </STRONG>sjl0vwresell_t<BR />
<STRONG>Site: </STRONG>eNom<BR />
<STRONG>IsLockable: </STRONG>True<BR />
<STRONG>IsRealTimeTLD: </STRONG>True<BR />
<STRONG>TimeDifference: </STRONG>+8.00<BR />
<STRONG>ExecTime: </STRONG>0.578<BR />
<STRONG>Done: </STRONG>true<BR />
<STRONG>TrackingKey: </STRONG>7dd10e2a-fc83-4bf0-9b5b-174f5a4bf02b<BR />
<STRONG>RequestDateTime: </STRONG>2/11/2015 5:09:30 PM<BR />
</BODY>
</HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
RRPCode=200
RRPText=Command completed successfully
Command=MODIFYNS
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
ExecTime=0.703
Done=true
TrackingKey=ea90f097-1d90-4a11-9a5a-ec0712838748
RequestDateTime=2/11/2015 5:18:24 PM
```
Related Commands
----------------

[CheckNSStatus](../docs/domains/domain-management/name-servers/check-ns-status.md)

[DeleteNameServer](../docs/domains/domain-management/name-servers/delete-name-server.md)

[GetDNS](../docs/domains/domain-management/name-servers/GetDNS.md)

[GetDNSStatus](../docs/domains/domain-management/name-servers/get-dns-status.md)

[ModifyNSHosting](../docs/domains/domain-management/name-servers/modify-ns-hosting.md)

[RegisterNameServer](../docs/domains/domain-management/name-servers/RegisterNameServer.md)

[UpdateNameServer](../docs/domains/domain-management/name-servers/UpdateNameServer.md)