AssignToDomainFolder
====================

Assign one or more domains to one folder.

Usage
-----

Use this command to assign one or more domains to a folder. Most commonly, you’ll put domains in a standard folder to organize them; in a Magic Folder to manage settings.

You can assign domains in several ways:

Put a list of domains into a folder.

Move a list of domains from one folder to another, or from all folders where the domains currently reside to a single folder.

Copy a list of domains to a folder so that they are in multiple folders \(but a domain can only be in one Magic Folder at a time; see Notes\).

Remove a list of domains from a folder or from all folders

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The folder must belong to this login ID.
- A domain can only be in one Magic Folder at a time. See Notes.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=AssignToDomainFolder&uid=(Required)&pw=(Required)&DomainNameList=(Required)&TransferMode=(Required)&FromFolderName=(Optional)&responsetype=(Optional)
```
| Input Parameter              | Type | Status          | Description |
| ------------------------------------------ | ------ | ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command                  | string | Required         | AssignToDomainFolder |
| uid | string | Required | Your Account ID                                                                                                                                                                                                       |
| pw                     | string | Required         | Your API Token |
| DomainNameList | string | Optional; default is Asc | Comma-separated list of domain names to put into this folder. Permitted format is sld1.tld1,sld2.tld2,sld3.tld3                                                                                                                                                     |
| TransferMode                | string | Optional         | Action to execute. Permitted values: 1 Copy to ToFolderName \(and leave in all folders where it is currently placed\) 2 Move to ToFolderName \(and remove from FromFolderName or from all\) 3 Remove from FromFolderName or from all folders |
| FromFolderName *Ignored if TransferMode=1* | string | Optional | If TransferMode=2 or 3, the query operates only on domains that are currently in FromFolderName. Permitted values are: - one folder name - - \(all folders\) If TransferMode=2 and no value is provided for FromFolderName, domains are moved from all folders where they are currently associated, to ToFolderName. If TransferMode=3 and no value is provided for FromFolderName, domains are deleted from all folders. |
| ToFolderName                | string | Optional         | Assign the listed domains to this folder. Permitted value is a single folder name. |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML                                                                                                                                                                          |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

A domain can only be in one Magic Folder at a time.

If a query string uses a wildcard value when removing domains from folders, or does not specify a foldername \(TransferMode=3&FromFolderName=\*\), it will remove the domains from all folders they arecurrently in.

| Output Parameter | Type | Description                                                                                                                                                                                                                  |
| ---------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command     | string | Name of command executed                                                                                                                                                                                                            |
| ToFolderName   | string | Destination folder for domains                                                                                                                                                                                                         |
| FromFolderName  | string | Source folder for domains                                                                                                                                                                                                           |
| ID        | string | Domain name IDs. To retrieve the domain names from these IDs, use the GetDomainSLDTLD command.                                                                                                                                                                        |
| TransferMode   | string | Type of reassignment done by this query. Values are: 1 Copy to ToFolderName \(and leave in all folders where it is currently associated\) 2 Move to ToFolderName \(and remove from FromFolderName or from all folders where it is currently associated\) 3 Remove from FromFolderName                                                                             |
| Result      | string | Result of this query. Values are: 0 General failure 1 Success 2 Unable to add domains to folder 3 No such folder 4 Cannot move a domain from a folder to the same folder 5 At least one domain is not found as active in this account 6 At least one domain cannot be removed from folder because synchronization is in progress 7 At least one domain is already associated with another Magic Folder 8 Unable to remove domains from folder |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                                                                                        |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                                                                                           |
| Done       | boolian | "True" indicates this entire response has reached you successfully.                                                                                                                                                                                      |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=AssignToDomainFolder&uid=resellid
&pw=resellpw&DomainIDList=152656650,152656651
&ToFolderName=Lock&TransferMode=1&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=AssignToDomainFolder&uid=resellid
&pw=resellpw&DomainIDList=152656650,152656651
&ToFolderName=Lock&TransferMode=1&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=AssignToDomainFolder&uid=resellid
&pw=resellpw&DomainIDList=152656650,152656651
&ToFolderName=Lock&TransferMode=1&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <AssignResult>
 <ToFolderName>Lock</ToFolderName>
 <FromFolderName>*</FromFolderName>
 <DomainNames>
  <ID>152656650</ID>
  <ID>152656651</ID>
 </DomainNames>
 <TransferMode>1</TransferMode>
 <Result>1</Result>
 </AssignResult>
 <Command>ASSIGNTODOMAINFOLDER</Command>
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
 <ExecTime>0.078</ExecTime>
 <Done>true</Done>
 <debug> <![CDATA[ ]]> </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>ToFolderName: </ STRONG>Lock<br>
<STRONG>FromFolderName: </ STRONG>*<br>
<STRONG>DomainList: </ STRONG><DomainNames><ID>152656650</ID><ID>152656651</ID></DomainNames><br>
<STRONG>TransferMode: </ STRONG>1<br>
<STRONG>Result: </ STRONG>3<br>
<STRONG>ErrorMessage: </ STRONG>Folder Lock does not exist<br>
<STRONG>Command: </ STRONG>ASSIGNTODOMAINFOLDER<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>1<br>
<STRONG>Err1: </ STRONG>Folder Lock does not exist<br>
<STRONG>ResponseCount: </ STRONG>1<br>
<STRONG>ResponseNumber1: </ STRONG>321393<br>
<STRONG>ResponseString1: </ STRONG>Validation error; failed to process;<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site:</ STRONG><br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.375<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 1:01:34 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
ToFolderName=Lock
FromFolderName=*
DomainList=<DomainNames><ID>152656650</ID><ID>152656651</ID></DomainNames>
TransferMode=1
Result=3
ErrorMessage=Folder Lock does not exist
Command=ASSIGNTODOMAINFOLDER
APIType=API
Language=eng
ErrCount=1
Err1=Folder Lock does not exist
ResponseCount=1
ResponseNumber1=321393
ResponseString1=Validation error; failed to process;
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.082
Done=true
RequestDateTime=2/3/2015 1:02:25 PM
```
Related Commands
----------------

[AddDomainFolder](../docs/adddomainfolder.md)

[AdvancedDomainSearch](../docs/advanced-domain-search.md)

[DeleteDomainFolder](../docs/deletedomainfolder.md)

[GetDomainFolderDetail](../docs/getdomainfolderdetail.md)

[GetDomainFolderList](../docs/getdomainfolderlist.md)

[RemoveUnsyncedDomains](../docs/removeunsynceddomains.md)

[UpdateDomainFolder](../docs/updatedomainfolder.md)