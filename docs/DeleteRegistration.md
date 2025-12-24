DeleteRegistration
==================

Delete a domain name registration.

Usage
-----

Use this command to delete a domain name registration. Fees will apply.

Availability
------------

This command is available to resellers on our DeleteRegistration whitelist. If you wish to have access to this command, contact your sales representative.

Constraints
-----------

The query must meet the following requirements:

*The login ID and password must be valid.&\#xA;* The domain name must belong to this account.

*The domain name must have been purchased less than five days ago.&\#xA;* The most common TLDs, including .com and .net, can be deleted. For other TLDs, check the resellerpages on our Web site.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=DeleteRegistration&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&EndUserIP=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetDomainExp |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| sld | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| tld       | string | Required | Top-level domain name |
| EndUserIP | string | Required | End user’s IP address. Permitted format is NNN.NNN.NNN.NNN      |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| DomainDeleted | string | Expiration date for the domain registration. |
| ErrString | string | Error string |
| ErrSource | string | Error source |
| ErrSection | string | Error section |
| RRPCode | string | Registry success code |
| RRPText | string | Text that corresponds to and explains the RRPCode value |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=deleteregistration&uid=resellid&pw=resellpw
&sld=resellerdocs3&tld=info&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=deleteregistration&uid=resellid&pw=resellpw
&sld=resellerdocs3&tld=info&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=deleteregistration&uid=resellid&pw=resellpw
&sld=resellerdocs3&tld=info&responsetype=text
```
```
<interface-response>response>
 <deletedomain>
 <domaindeleted>True</domaindeleted>
 </deletedomain>
 <ErrString/>
 <ErrSource/>
 <ErrSection>DELETEREGISTRATION</ErrSection>
 <RRPCode>200</RRPCode>
 <RRPText>Command completed successfully</RRPText>
 <Command>DELETEREGISTRATION</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>2.75</ExecTime>
 <Done>true</Done>
 <debug>
 [CDATA ]
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>domaindeleted: </STRONG>True<br />
<STRONG>ErrString:</STRONG><br />
<STRONG>ErrSource:</STRONG><br />
<STRONG>ErrSection: </STRONG>DELETEREGISTRATION<br />
<STRONG>RRPCode: </STRONG>200<br />
<STRONG>RRPText: </STRONG>Command completed successfully<br />
<STRONG>Command: </STRONG>DELETEREGISTRATION<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod:</STRONG><br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site:</STRONG><br />
<STRONG>IsLockable:</STRONG><br />
<STRONG>IsRealTimeTLD:</STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.125<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/11/2015 1:07:09 PM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
domaindeleted=True
ErrString=
ErrSource=
ErrSection=DELETEREGISTRATION
RRPCode=200
RRPText=Command completed successfully
Command=DELETEREGISTRATION
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
RequestDateTime=2/11/2015 1:09:29 PM
```
Related Commands
----------------

AddToCart

Purchase