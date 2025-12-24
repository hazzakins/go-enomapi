DeleteAllPOPPaks
================

Delete all POP3 email paks in a domain.

Usage
-----

Use this command to delete all POP3 paks from a domain. Note that if you delete the POP paks, the only way to replace them is to buy new paks.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

All POP paks in the domain must have no user emails configured. You can use the DeletePOP3 command to remove user email accounts.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                 | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                              | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                            | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| PaksRemoved | Number of POP paks removed from the domain name                         |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- Use the DeletePOP3 command to delete any email user accounts before running DeleteAllPOPPaks.

Example
-------

The following query deletes all POP3 10-paks from resellerdocs.com and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=DeleteAllPOPPaks&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteAllPOPPaks&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteAllPOPPaks&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=text
```
The response is as follows:

```
<interface-response>
 <PaksRemoved>35</PaksRemoved>
 <Command>DELETEALLPOPPAKS</Command>
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
 <ExecTime>2.908</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/7/2011 5:07:14 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
PaksRemoved:35
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
Related Commands
----------------

DeletePOP3

DeletePOPPak

Forwarding

GetForwarding

GetMailHosts

GetPOP3

PurchasePOPBundle

SetUpPOP3User