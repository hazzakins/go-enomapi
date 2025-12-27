ParseDomain
===========

Separate the domain name into its host, SLD, and TLD.

Usage
-----

Use this command to separate a domain name into its constituent parts.

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

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| PassedDomain  | Required | Full name to parse, including the third level if appropriate | 70 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Host       | Host name |
| SLD | Second-level domain name \(for example, enom in enom.com\)                   |
| TLD       | Top-level domain name \(extension\) |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query separates the domain name [www.resellerdocs.com](http://www.resellerdocs.com/) into its Host, SLD, and TLD, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=parsedomain&uid=resellid&pw=resellpw
&passeddomain=www.resellerdocs.com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=parsedomain&uid=resellid&pw=resellpw
&passeddomain=www.resellerdocs.com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=parsedomain&uid=resellid&pw=resellpw
&passeddomain=www.resellerdocs.com&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
<ParseDomain>
 <Host>www</Host>
 <SLD>resellerdocs</SLD>
 <TLD>com</TLD>
</ParseDomain>
<Command>PARSEDOMAIN</Command>
<ErrCount>0</ErrCount>
<Server>RESELLERTEST</Server>
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
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Host: www
SLD: resellerdocs
TLD: com
Command: PARSEDOMAIN
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
ExecTime: 0.047
Done: true
RequestDateTime: 2/4/2015 4:07:28 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Host=www
SLD=resellerdocs
TLD=com
Command=PARSEDOMAIN
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
RequestDateTime=2/4/2015 4:07:50 PM
```