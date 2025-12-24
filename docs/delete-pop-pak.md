DeletePOPPak
============

Delete a POP Email pak under a domain name.

Usage
-----

Use this command to delete an entire 10-pak of POP accounts. Note that once a pak is deleted, the only way to replace it is to purchase a new one.

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
- The pak must have no user email accounts currently configured.

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
| BundleID    | Required | ID number of the POP pak to delete. Use GetPOP3 to get POP pak IDs.                    | 6 |
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
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err \(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.
- Use the DeletePOP3 command to delete any email user accounts before running DeletePOPPak.

Example
-------

The following query deletes POP bundle 5733 from resellerdocs.com and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=DeletePOPPak&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&BundleID=5733&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=DeletePOPPak&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&BundleID=5733&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=DeletePOPPak&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&BundleID=5733&ResponseType=text
```
The response is as follows:

```
<interface-response>
 <Command>DELETEPOPPAK</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+8.00</TimeDifference>
 <ExecTime>0.750</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>31063cc4-7994-4d96-b8bd-1ff3665d6039</TrackingKey>
 <RequestDateTime>12/12/2011 3:25:01 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Command: DELETEPOPPAK
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +8.00
ExecTime: 0.125
Done: true
TrackingKey: 0618aab6-876c-4f07-b791-a955a0d1fe8c
RequestDateTime: 2/3/2015 3:32:38 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Command=DELETEPOPPAK
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
ExecTime=0.031
Done=true
TrackingKey=323f2cdc-6c00-486d-b458-12401dd086f6
RequestDateTime=2/3/2015 3:33:04 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

Forwarding

GetForwarding

GetMailHosts

GetPOP3

ModifyPOP3

PurchasePOPBundle

SetUpPOP3User