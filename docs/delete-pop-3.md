DeletePOP3
==========

Delete an individual POP3 user name from the database.

Usage
-----

Use this command to delete an individual POP3 email user name. Once you delete it, you can fill that vacancy with another name. This command has no effect on the POP 10-pak nor on any other user names in the account.

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
- The user name must be listed in the POP3 list for the account.
- The user name must be in a POP pak for which the subscription is current.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                     | Max Size |
| --------------- | ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                   | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with                     | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| UserNameX    | Required | POP mailbox user name to delete. Start with UserName1 and index to delete multiple user names, for example, UserName1=John.Doe&UserName2=Jane.Doe | 16 |
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
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query deletes the POP user james and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=DeletePOP3&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&UserName1=james&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=DeletePOP3&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&UserName1=james&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=DeletePOP3&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&UserName1=james&responsetype=text
```
The response is as follows:

```
<interface-response>
 <AccountsDeleted>1</AccountsDeleted>
 <success>True</success>
 <Command>DELETEPOP3</Command>
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
 <ExecTime>0.219</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>ef0e27fe-c956-4f45-a1f9-da8b6b8adea5</TrackingKey>
 <RequestDateTime>12/12/2011 3:17:09 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
success: True
Command: DELETEPOP3
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
ExecTime: 0.594
Done: true
TrackingKey: fbed2f7d-f570-4a55-82dd-f327f31684fe
RequestDateTime: 2/11/2015 1:03:15 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOPPak

Forwarding

GetCatchAll

GetForwarding

GetMailHosts

GetPOP3

ModifyPOP3

PurchasePOPBundle

SetCatchAll

SetUpPOP3User