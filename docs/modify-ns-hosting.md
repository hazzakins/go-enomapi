ModifyNSHosting
===============

Modify the name server settings for a domain name, in our database, without changing the settings at the Registry.

Usage
-----

ModifyNSHosting redirects to another set of name servers without updating the Registry records. ModifyNS, a similar command, updates name server records at the Registry and in the registrar’s database.

Use this command when name servers are set correctly at the Registry but incorrectly in our records.

Use this command when a domain name is registered in your account, and uses DNS Hosting in a different account. In this case, set the name servers to our name servers at the Registry, and to N/A in our database.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The name server cited in the query must exist.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                             | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                                           | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                     | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| NSX X=1 to 12  | Required | Name server—maximum of 12 can be set. Permitted value is the use name of the name server, for example, NS1=ns1.name-services.com | 60 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query designates ns1.name-services.com as the name server for resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=modifynshosting&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ns1=ns1.name-services.com
&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=modifynshosting&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ns1=ns1.name-services.com
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=modifynshosting&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ns1=ns1.name-services.com
&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <Command>MODIFYNSHOSTING</Command>
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
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command: MODIFYNSHOSTING
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +8.00
ExecTime: 0.797
Done: true
TrackingKey: 05f4a1f5-5a9b-42f6-af05-f953dff48a69
RequestDateTime: 2/4/2015 3:38:48 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Command=MODIFYNSHOSTING
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.078
Done=true
TrackingKey=f665a851-dc11-489a-8ef4-9ac74b47e3f4
RequestDateTime=2/4/2015 3:39:13 PM
```
Related Commands
----------------

CheckNSStatus

DeleteNameServer

GetDNS

GetDNSStatus

ModifyNS

RegisterNameServer

UpdateNameServer