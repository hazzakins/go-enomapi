SetPakRenew
===========

Set the auto-renew behavior for a POP email pak.

Usage
-----

Use this command to set auto-renew behavior for an individual POP pak.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/POPConfig.asp?DomainNameId=152533676](https://resellertest.enom.com/domains/POPConfig.asp?DomainNameId=152533676)

At the bottom of the section for each POP pak, the Attempt to auto-renew check box calls the SetPakRenew command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The POP pak bundle ID must belong to this domain.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                          | Max Size |
| --------------- | -------- | ---------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                        | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                  | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| BundleID    | Required | Numerical ID for the POP pak \(retrieve the bundle ID using the GetPOP3 command\)       | 4 |
| AutoPakRenew | Required | Auto-renew setting to be applied to this POP pak. Permitted values are 1 \(on\) or 0 \(off\) | 1    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML.           | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| DomainRRP | Domain type                                           |
| PakUpdated    | Bundle ID of the POP pak for which the auto-renew value is being reset |
| AutoPakRenew | Auto-renew setting for the pak. Values are 1 \(on\) or 0 \(off\)                |
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

The following query turns the POP pak auto-renew setting for bundle 5105 on, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SETPAKRENEW&uid=resellid&pw=resellpw
&DomainNameID=152533676&bundleid=5105
&AutoPakRenew=1&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=SETPAKRENEW&uid=resellid&pw=resellpw
&DomainNameID=152533676&bundleid=5105
&AutoPakRenew=1&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=SETPAKRENEW&uid=resellid&pw=resellpw
&DomainNameID=152533676&bundleid=5105
&AutoPakRenew=1&ResponseType=text
```
In the response, the error count of 0 and absence of error messages confirm that the query was successful:

```
<interface-response>
 <PakUpdated>5105</PakUpdated>
 <AutoPakRenew>1</AutoPakRenew>
 <Command>SETPAKRENEW</Command>
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
 <ExecTime>0.156</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/11/2011 11:33:42 PM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success: True
Command: SETPAKRENEW
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
ExecTime: 0.125
Done: true
TrackingKey: 10a9754d-525d-4026-95ab-c6ba51416831
RequestDateTime: 2/5/2015 3:10:07 PM
```
Related Commands
----------------

GetPOP3

GetRenew

SetRenew

UpdateCusPreferences