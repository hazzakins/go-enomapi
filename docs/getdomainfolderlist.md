GetDomainFolderList
===================

Retrieve a list of folders that a domain is in, or retrieve a list of all folders in an account.

Usage
-----

Use this command to find out which folders a domain is in, or to retrieve a list of all folders in an account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder](https://resellertest.enom.com/domains/Domain-Manager.aspx?tab=folder)

The folders link calls the GetDomainFolderList command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                                                             | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                                                                                           | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                               | 4 |
| SLD | Optional | SLD \(second-level domain name\) of the domain for which you want to retrieve folder information. Use this parameter if you want information for one domain. Without SLD and TLD, the response lists all folders in the account. | 63    |
| TLD       | Optional | TLD \(top-level domain name\) of the domain for which you want to retrieve folder information. Use this parameter if you want information for one domain. Without SLD and TLD, the response lists all folders in the account.  | 15 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                              |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                       |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                       |
| Done       | True indicates this entire response has reached you successfully. |
| Result | Result of this query. Possible results are: 1 Successful                                       |
| ID        | Folder identification number, assigned by us |
| Name | Folder name, assigned by you                                                     |
| Type       | Folder type. Possible values are: 0 Standard folder; not Magic 1 Magic |
| DomainCount | Number of domains in this folder                                                   |
| Status      | Status of this folder. Possible results are: 0 Standard 2 Magic, in sync 3 Magic, in process of synchronization 4 Magic, out of sync |
| Count | Total number of folders in this account                                                |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests a list of folders that contain the specified domain, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetDomainFolderList&uid=resellid
&pw=resellpw&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainFolderList&uid=resellid
&pw=resellpw&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainFolderList&uid=resellid
&pw=resellpw&sld=resellerdocs&tld=com&responsetype=text
```
In the response, a Result value 1 and an ErrCount value 0 indicate that the query was successful:

```
<?xml version="1.0" ?>

 <Result>1</Result>

 </GetResult>

  <ID>3100</ID>

  <Name>ResellFolder</Name>

  <Type>1</Type>

  <DomainCount>3</DomainCount>

  <Status>4</Status>

 </item>

  <ID>974</ID>

  <Name>a3223</Name>

  <Type>0</Type>

  <DomainCount>44</DomainCount>

  <Status>0</Status>

 </folders>

 <count>2</count>

 <Command>GETDOMAINFOLDERLIST</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>1</MaxPeriod>

 <Server>RESELLER1-STG</Server>

 <Site>enom</Site>

 <IsLockable>False</IsLockable>

 <IsRealTimeTLD>True</IsRealTimeTLD>

 <TimeDifference>+03.00</TimeDifference>

 <ExecTime>0.109</ExecTime>

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

count: 0

Command: GETDOMAINFOLDERLIST

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

ExecTime: 0.078

Done: true

TrackingKey: 4f718ce6-71ee-4ded-8f78-a66c6be4019f

RequestDateTime: 2/3/2015 5:43:05 PM
```
ExecTime=0.016

TrackingKey=bb09a1ee-2217-40fc-915c-cbb4ffd62d02

RequestDateTime=2/3/2015 5:43:54 PM
```