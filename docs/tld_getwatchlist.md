TLD\GetWatchlist
================

Retrieve a list of domains currently in a Watchlist LIST in this account.

Usage
-----

Use this command to retrieve a list of domains currently in a Watchlist LIST in this account.

List
----

You can create LIST\(s\) for organizing or managing watchlist domain names. Each LIST consists of a ListID, ListName and ListEmail.

Each account has a default LIST \(ListID=0 \), containing domains to watch. Domains can appear in multiple LISTS. The default LIST cannot be deleted and email address cannot be changed.

| ListID=0                              | ListID=1 | ListID=2           |
| ------------------------------------------------------------------- | ---------------------------- | ---------------------------- |
| - .ninja .art .education one.ceo one.ninja two.ceo two.ninja | one.ceo one.ninja one.fun | two.ceo two.ninja two.fun |

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/tlds/watchlist.aspx](https://resellertest.enom.com/tlds/watchlist.aspx)

The My Watchlist tab calls the TLD\_GetWatchlist command.

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
| Input Parameter | Status | Description                                                                                                                                                                                         | Max Size |
| ---------------- | ------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                                                                                                                                       | 20 |
| PW | Required            | Account password | 20    |
| ResponseType   | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                                                                                                           | 4 |
| Action | Optional default is Empty   | This parameter determines which result set will be provided. If this parameter is not declared, or has no value, the API will return a list of all of the domains in the specified LIST. Set the parameter to GetListIDs to retreive a list of all of the LISTs for this account. Check the example section see the differences. Permitted values are: Empty \(not declared\) GetListIDs |     |
| ListID      | Optional | Return only those domain\(s\) belonging to a specific LIST. Not declaring this parameter will result in the domains returned from the default LIST \(ListID=0\).                                                                                                              | |
| SortBy | Optional default is DomainName | Sorting parameter. Permitted values are: WatchID DomainName SLD TLD TLDCategory | 20    |
| SortByDirection | Optional default is Asc | Sorting order. Permitted values are: Asc Desc                                                                                                                                                                        | 4 |
| StartPosition | Optional default is 1     | Return results beginning with this position in the sorted list For example, StartPosition=26&PagingPageSize=25 returns results 26 through 50 in the sorted list |     |
| PagingPageSize  | Optional default is 25 | Number of results to return in this response                                                                                                                                                                         | |
| SLD | Optional            | Filter by SLD | 275   |
| TLD       | Optional | Filter by TLD                                                                                                                                                                                        | 15 |
| Category | Optional            | Filter by Category | 500   |
| DomainNameFilter | Optional | Filter by Domain Name. When present, this will filter the result set according to the text supplied. This parameter is similar to using the LIKE clause in SQL to generate a result set based on the input string. For example, if the string is ion, the results may include mydomain.auction, *.auction, ion-drive.ceo, ion-drive.*, etc.                    | 300 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ------------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done        | True indicates this entire response has reached you successfully. |
| ID | Watchlist Item ID for each domain name                              |
| DomainName     | Domain name |
| SLD | SLD value                                            |
| TLD         | TLD value |
| NativeIDN | Domain name in Native IDN format                                 |
| StatusID      | Status ID for this specific domain |
| StatusDesc | Status description for this specific domain                           |
| SubStatusID     | Sub-Status ID for this specific domain |
| SubStatusDesc | Sub-Status description for this specific domain                         |
| TLDCategory     | TLD Primary Category |
| TLDCategoryDesc | TLD Category description                                     |
| TLDStatusID     | Status ID for this specific TLD |
| TLDStatusDesc | Status description for this specific TLD                             |
| ListID       | List ID |
| ListName | List Name                                            |
| ListEmail      | List Email Address |
| SortBy | Sort By value                                          |
| SortByDirection   | Sort By Direction type |
| StartPosition | Start Position value                                       |
| PagingPageSize   | Paging Page Size value |
| FilterByDomainName | Filter by Domain Name value                                   |
| FilterBySLD     | Filter by SLD value |
| FilterByTLD | Filter by TLD value                                       |
| FilterByTLDCategory | Filter by TLD Category value |
| RecordsCount | Records return in one page for this request                           |
| TotalRecords    | Total records for this type of request |
| AllRecords | All records for this account                                   |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves list of domains currently added to the Watchlist in this account, and requests the response in given format.

Get LIST domains
----------------

```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlist&uid=ResellID&pw=resellpw
&responsetype=xml&pagingpagesize=2
&tld=xn--nqv7f&ListID=0&Action=
```
```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlist&uid=ResellID&pw=resellpw
&responsetype=html&pagingpagesize=2
&tld=xn--nqv7f&ListID=0&Action=
```
```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlist&uid=ResellID&pw=resellpw
&responsetype=text&pagingpagesize=2
&tld=xn--nqv7f&ListID=0&Action=
```
```
<?xml version="1.0"?>

  <ID>69018</ID>

  <DomainName><![CDATA[ blogging.xn--nqv7f ]]></DomainName>

  <SLD><![CDATA[ blogging ]]></SLD>

  <TLD><![CDATA[ xn--nqv7f]]></TLD>

  <NativeIDN><![CDATA[ blogging.机构 ]]></NativeIDN>

  <StatusID>1</StatusID>

  <StatusDesc><![CDATA[ Watching ]]></StatusDesc>

  <SubStatusID>0</SubStatusID>

  <SubStatusDesc><![CDATA[ ]]></SubStatusDesc>

  <TLDCategory><![CDATA[ IDN ]]></TLDCategory>

  <TLDCategoryDesc><![CDATA[ ]]></TLDCategoryDesc>

  <TLDStatusID>1</TLDStatusID>

  <TLDStatusDesc><![CDATA[ Awaiting Evaluation ]]> </TLDStatusDesc>

 </Item>

.

.

.

 <ListID />

 <ListName />

 <ListEmail />

 <SortBy>none</SortBy>

 <SortByDirection>desc</SortByDirection>

 <StartPosition>1</StartPosition>

 <PagingPageSize>2</PagingPageSize>

 <SLD />

 <TLD>ninja</TLD>

 <Category />

 <DomainName />

 <RecordsCount>2</RecordsCount>

 <TotalRecords>9</TotalRecords>

 <AllRecords>3023</AllRecords>

 </TLD_Watchlist>

 <Command>TLD_GETWATCHLIST</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>eNom</Site>

 <IsLockable/>

 <IsRealTimeTLD/>

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>2.967</ExecTime>

 <Done>true</Done>

 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>

 <RequestDateTime>7/23/2012 11:29:08 AM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

ID1: 169399

DomainName1: blogging.xn--nqv7f

SLD1: blogging

TLD1: xn--nqv7f

NativeIDN1: blogging.机构

StatusID1: 1

StatusDesc1: Watching

SubStatusID1: 0

Phase1: General Availability

IsOpen1: 0

SubStatusDesc1:

TLDCategory1: IDN

TLDCategoryDesc1:

TLDStatusID1: 4

TLDStatusDesc1: General Availability

TLDStatusAdditional1: This TLD is now in General Availability, do a domain search and register one today!

TLDPhaseFilter1:

DateStart1: 2014-06-04T11:00:00

DateEnd1:

ID2: 160765

DomainName2: dm1.xn--nqv7f

SLD2: dm1

TLD2: xn--nqv7f

NativeIDN2: dm1.机构

StatusID2: 1

StatusDesc2: Watching

SubStatusID2: 0

Phase2: General Availability

IsOpen2: 0

SubStatusDesc2:

TLDCategory2: IDN

TLDCategoryDesc2:

TLDStatusID2: 4

TLDStatusDesc2: General Availability

TLDStatusAdditional2: This TLD is now in General Availability, do a domain search and register one today!

TLDPhaseFilter2:

DateStart2: 2014-06-04T11:00:00

DateEnd2:

ListID: 0

SortBy: none

SortByDirection: desc

StartPosition: 1

PagingPageSize: 2

FilterByTLD: xn--nqv7f

RecordsCount: 2

TotalRecords: 2

AllRecords: 953

count: 3

Command: TLD_GETWATCHLIST

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

TimeDifference: +0.00

ExecTime: 1.484

Done: true

TrackingKey: 9ebd0e2e-340b-47e3-a88e-075dcfedc6e9

RequestDateTime: 2/6/2015 12:12:47 PM
```
;Machine is SJL0VWRESELL_T

ExecTime=0.094

TrackingKey=6a31a2fa-4ec4-4ef0-8ba5-8e4e379f6672

RequestDateTime=2/6/2015 12:13:21 PM
```
Get LIST
--------

```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlist&uid=ResellID&pw=resellpw
&responsetype=xml&pagingpagesize=2&tld=
&ListID=&Action=GetListID
```
```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlist&uid=ResellID&pw=resellpw
&responsetype=html&pagingpagesize=2&tld=
&ListID=&Action=GetListID
```
```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlist&uid=ResellID&pw=resellpw
&responsetype=text&pagingpagesize=2&tld=
&ListID=&Action=GetListID
```
  <ListID>0</ListID>

  <ListName><![CDATA[ ]]></ListName>

  <ListEmail><![CDATA[ ]]></ListEmail>

  <ListID>18</ListID>

  <ListName><![CDATA[ Customer1 ]]></ListName>

  <ListEmail><![CDATA[ [email protected] ]]></ListEmail>

 <SortByDirection>asc</SortByDirection>

 <TotalRecords>125</TotalRecords>

ID1: 169038

DomainName1: blogging.holdings

TLD1: holdings

NativeIDN1: blogging.holdings

TLDCategory1: Money and Finance

DateStart1: 2014-01-01T08:00:00

ID2: 160408

DomainName2: dm1.holdings

TLD2: holdings

NativeIDN2: dm1.holdings

TLDCategory2: Money and Finance

DateStart2: 2014-01-01T08:00:00

TotalRecords: 1299

Server: sjl0vwresell_t

ExecTime: 0.188

TrackingKey: 5801230b-9b20-4e38-ad00-440e13d50126

RequestDateTime: 2/6/2015 12:13:57 PM
```
ExecTime=0.125

TrackingKey=6b84d0f0-f60f-4d08-943d-b168b2a23131

RequestDateTime=2/6/2015 12:14:31 PM
```