DeleteDomainFolder
==================

Delete a folder.

Usage
-----

Use this command to delete one domain name folder.

This command does not delete domains; only the folder.

When you delete a folder, the domains that were in it keep the settings assigned by the folder until you change them.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder](https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder)

In the list of folders, the delete link calls the DeleteDomainFolder command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The folder must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |
| FolderName | Required | Name of the folder you want to delete | 125   |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| FolderName    | Name of folder |
| Result | Success of this query. Values are: 0 Failure 1 Success 2 Folder not found in this account    |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query deletes a domain folder, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=DeleteDOMAINFOLDER&uid=resellid&pw=resellpw
&FolderName=Favorites%202&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteDOMAINFOLDER&uid=resellid&pw=resellpw
&FolderName=Favorites%202&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteDOMAINFOLDER&uid=resellid&pw=resellpw
&FolderName=Favorites%202&ResponseType=text
```
The response is as follows:

<FolderName>Favorites 2</FolderName>

<Result>1</Result>

</DeleteResult>

<Command>DELETEDOMAINFOLDER</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod/>

<MaxPeriod>10</MaxPeriod>

<Server>sjl21wresellt01</Server>

<Site>eNom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>0.469</ExecTime>

<Done>true</Done>

<debug/>

<TrackingKey>50ab0431-8c68-4d63-b47b-a7b27bc56b54</TrackingKey>

<RequestDateTime>12/7/2011 5:25:53 AM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

FolderName: Favorites2

Result: 1

Command: DELETEDOMAINFOLDER

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.047

Done: true

TrackingKey: cadc29c6-4242-4df3-87ae-922050db543c

RequestDateTime: 2/3/2015 3:15:51 PM
```