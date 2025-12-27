PE\GetTLDID
===========

Retrieve the ID number for a TLD.

Usage
-----

Use this command to retrieve the ID number for a TLD.

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
| TLD       | Required | Top-level domain name \(extension\)         | 15 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| TLDID      | ID number for the TLD specified in the query string. |
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

The following query requests the ID number for the .org TLD, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PE_GETTLDID&uid=resellid&pw=resellpw
&tld=org&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETTLDID&uid=resellid&pw=resellpw
&tld=org&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PE_GETTLDID&uid=resellid&pw=resellpw
&tld=org&responsetype=text
```
In the response, a value for the TLDID parameter confirms that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
<productid>
 <tldid>2</tldid>
</productid>
<Command>PE_GETTLDID</Command>
<Language>en</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+07.00</TimeDifference>
<ExecTime>0.1015625</ExecTime>
<Done>true</Done>
<debug>
 <![CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
tldid: 2
Command: PE_GETTLDID
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +07.00
ExecTime: 0.094
Done: true
RequestDateTime: 2/4/2015 4:42:48 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
tldid=2
Command=PE_GETTLDID
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+07.00
ExecTime=0.047
Done=true
RequestDateTime=2/4/2015 4:43:11 PM
```
Related Commands
----------------

PE\_SetPricing

SetResellerTLDPricing