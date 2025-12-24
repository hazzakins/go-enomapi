AddDomainFolder
===============

Create a folder for organizing or managing your domain names.

Usage
-----

Use this command to create either a Standard folder that simply lets you organize your domains, or a Magic folder that allows you to apply the same settings to all the domains in the folder.

This command creates folders; another command, UpdateDomainFolder, configures their settings. Settings that can be managed via a Magic folder include the following, in any combination:

Auto-renew

Registrar lock

Access password

DNS servers

Host records

Contact information \(any combination of Registrant, Administrative, Auxiliary Billing, Technical\)

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder&start=1](https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder&start=1)

In the Create a new folder section, the GO button calls the AddDomainFolder command.

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
| Input Parameter  | Status | Description                                                                                            | Max Size |
| ----------------- | ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID        | Required | Account login ID                                                                                          | 20 |
| PW | Required        | Account password | 20    |
| ResponseType   | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                              | 4 |
| FolderName | Required        | Folder name. Permitted characters are letters, numbers, hyphens, underscores, and spaces; name must begin with a letter or number. | 125   |
| FolderDescription | Optional | Description of folder or contents                                                                                 | 256 |
| FolderType | Optional; default is 0 | Type of folder to create. Permitted values: 0 Standard folder \(for organizing and sorting\) 1 Magic folder \(a folder that allows you to manage settings of domains\) | 1    |
| CopyFolderName  | Optional | Copy a folder into this folder, as well as leaving it in the folder list. Permitted values are names of existing folders in this account.                            | 125 |
| CopyDomainName | Optional        | Copy one domain that is in another folder, and put it also in this folder. A domain can only be in one Magic folder at a time. The domain must be in this account. Permitted format is sld.tld | 70    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ----------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command      | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                           |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                          |
| FolderID     | Identification number assigned by our database |
| FolderName | Folder name, as assigned by you                                                            |
| FolderType    | Folder type. 0 indicates Standard folder, 1 indicates Magic folder. |
| FolderDescription | Folder description, as typed by you                                                          |
| Result      | Result of running this query string. Values are: 0 Folder was not created 1 Folder was created successfully 2 A folder with this name already exists |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query creates a new folder and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=AddDomainFolder&uid=resellid&pw=resellpw
&FolderName=Favorites&FolderDescription=FolderForMyFavoriteNames
&FolderType=1&CopyFolderName=Subfavorites
&CopyDomainName=resellerdocs.com&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=AddDomainFolder&uid=resellid&pw=resellpw
&FolderName=Favorites&FolderDescription=FolderForMyFavoriteNames
&FolderType=1&CopyFolderName=Subfavorites
&CopyDomainName=resellerdocs.com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=AddDomainFolder&uid=resellid&pw=resellpw
&FolderName=Favorites&FolderDescription=FolderForMyFavoriteNames
&FolderType=1&CopyFolderName=Subfavorites
&CopyDomainName=resellerdocs.com&responsetype=text
```
In the response, the presence of your input values and an ErrCount value 0 indicate that the query was successful:

<WscAccountOverride>False</WscAccountOverride>

<ItemName>resellerdocs1.com</ItemName>

<ItemId/>

<Price>8.95</Price>

<ICANNFees>0.20</ICANNFees>

<CartItemID>889049</CartItemID>

<NewDomainNameID>152932786</NewDomainNameID>

<ItemAdded>True</ItemAdded>

<Command>ADDTOCART</Command>

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

<ExecTime>0.688</ExecTime>

<Done>true</Done>

<RequestDateTime>12/6/2011 11:51:55 PM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

FolderID: 27161

FolderName: Favorites

FolderType: 1

FolderDescription: FolderForMyFavoriteNames

Result: 1

Command: ADDDOMAINFOLDER

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

ExecTime: 0.066

Done: true

RequestDateTime: 2/3/2015 12:04:28 PM
```
FolderID=27162

ExecTime=0.137

RequestDateTime=2/3/2015 12:05:49 PM
```