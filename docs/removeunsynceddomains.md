RemoveUnsyncedDomains
=====================

Remove out-of-sync domains from a specified Magic Folder, or from all Magic Folders in the account.

Usage
-----

Use this command to remove out-of-sync domains from a specified Magic Folder, or from all Magic Folders in the account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/folder-control-panel/default.asp?FolderName=AAAaa](https://resellertest.enom.com/domains/folder-control-panel/default.asp?FolderName=AAAaa)

The Remove Names button calls the RemoveUnsyncedDomains command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                             | Max Size |
| --------------- | ------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                                           | 20 |
| PW | Required                  | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                               | 4 |
| FolderName | Optional; default is all folders in account | Folder from which to remove unsynced domains. If no FolderName is supplied, all unsynced domains in all Magic Folders are removed. | 125   |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| Result      | Result of processing this query. Possible return values are: 1 Success 2 Failure |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query removes out-of-sync domains from one folder, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=RemoveUnsyncedDomains&uid=resellid
&pw=resellpw&FolderName=Master-Magic&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=RemoveUnsyncedDomains&uid=resellid
&pw=resellpw&FolderName=Master-Magic&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=RemoveUnsyncedDomains&uid=resellid
&pw=resellpw&FolderName=Master-Magic&responsetype=text
```
In the response, a Result value 1 indicates that the query was successful:

```
<?xml version="1.0" ?>

 <Result>1</Result>

 </RemoveResult>

 <Command>REMOVEUNSYNCEDDOMAINS</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLER1-STG</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.125</ExecTime>

 <Done>true</Done>

 <![CDATA[ ]]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Result: 1

Command: REMOVEUNSYNCEDDOMAINS

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site:

TimeDifference: +0.00

ExecTime: 0.188

Done: true

RequestDateTime: 2/5/2015 2:15:08 PM
```
;Machine is SJL0VWRESELL_T

Server=SJL0VWRESELL_T

ExecTime=0.156

RequestDateTime=2/5/2015 2:16:49 PM
```