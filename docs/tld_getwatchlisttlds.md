TLD\GetWatchlistTlds
====================

Retrieve a list of Tlds currently in a Watchlist LIST in this account.

Usage
-----

Use this command to retrieve a list of Tlds currently in a Watchlist LIST in this account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/tlds/watchlist.aspx](https://resellertest.enom.com/tlds/watchlist.aspx)

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
| Input Parameter  | Status | Description                                                                           | Max Size |
| ------------------ | ------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID        | Required | Account login ID                                                                        | 20 |
| PW | Required            | Account password | 20    |
| ResponseType    | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                            | 4 |
| ApplicationFeeMin | Optional            | Filter based on minimum Application fee |     |
| ApplicationFeeMax | Optional | Filter based on Maximum Application fee                                                             | |
| DateStart | Optional            | Filter based on Date |     |
| DateEnd      | Optional | Filter based on Date                                                                      | |
| ListID | Optional            | Return only those Tlds belonging to a specific LIST. Not declaring this parameter will result in the tlds returned from the default LIST \(ListID=0\). |     |
| SortBy       | Optional default is DomainName | Sorting parameter. Permitted values are: WatchID DomainName SLD TLD TLDCategory                                        | 20 |
| SortByDirection | Optional default is Asc    | Sorting order. Permitted values are: Asc Desc | 4    |
| StartPosition   | Optional default is 1 | Return results beginning with this position in the sorted list For example, StartPosition=26&PagingPageSize=25 returns results 26 through 50 in the sorted list | |
| PagingPageSize | Optional default is 25     | Number of results to return in this response |     |
| SLD        | Optional | Filter by SLD                                                                          | 275 |
| TLD | Optional            | Filter by TLD | 15    |
| Category      | Optional | Filter by Category                                                                       | 500 |
| IsOpen | Optional            | Permitted values True / False or 0 / 1 |     |
| RegistrationFeeMin | Optional | Filter based on minimum Registration fee                                                            | |
| RegistrationFeeMax | Optional            | Filter based on Maximum Registration fee |     |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| TLD       | TLD value |
| TLDCategoryDesc | TLD Category description                                     |
| TLDCategory   | TLD Primary Category |
| DomainCount | Displays Domain count                                      |
| SortBy      | Sort By value |
| SortByDirection | Start Position value                                       |
| StartPosition  | Records starting position |
| PagingPageSize | Paging size for records to display                                |
| TotalRecords   | Total records for this type of request |
| SubStatusDesc | Sub-Status description for this specific domain                         |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves list of domains currently added to the Watchlist in this account, and requests the response in given format.

Get LIST
--------

```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlisttlds&uid=ResellID&pw=resellpw
&responsetype=xml&pagingpagesize=2&tld=&ListID=
&Action=GetListID
```
```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlisttlds&uid=ResellID&pw=resellpw
&responsetype=html&pagingpagesize=2&tld=&ListID=
&Action=GetListID
```
```
https://resellertest.enom.com/interface.asp?
command=tld_getwatchlisttlds&uid=ResellID&pw=resellpw
&responsetype=text&pagingpagesize=2&tld=&ListID=
&Action=GetListID
```
```
<?xml version="1.0" encoding="utf-8" ?>

  <![CDATA[ AGENCY ]]>

  </TLD>

  <![CDATA[ Awaiting Evaluation ]]>

  </TLDCategoryDesc>

  <![CDATA[ 1 ]]>

  </DomainCount>

 </Item>

  <![CDATA[ NINJA ]]>

  <![CDATA[ PHOTOS ]]>

  <![CDATA[ ACTOR ]]>

  <![CDATA[ 0 ]]>

  <![CDATA[ AIRFORCE ]]>

 <SortBy>none</SortBy>

 <SortByDirection>desc</SortByDirection>

 <StartPosition>1</StartPosition>

 <PagingPageSize>5</PagingPageSize>

 <TotalRecords>5</TotalRecords>

 </TLD_WatchlistTlds>

 <Command>TLD_GETWATCHLISTTLDS</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>10</MaxPeriod>

 <Server>sjl21wresell01</Server>

 <Site>eNom</Site>

 <IsLockable/>

 <IsRealTimeTLD/>

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.047</ExecTime>

 <Done>true</Done>

 <TrackingKey>89271891-f58d-4920-8b0d-ef0cef6b8b6a</TrackingKey>

 <RequestDateTime>3/11/2013 12:17:43 AM</RequestDateTime>

 <debug/>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

TLD: DELIVERY

TLDCategory: Services

TLDStatusID: 2

TLDStatusDesc: Evaluation Passed

TLDStatusAdditional: All applications for TLD have passed ICANN’s evaluation.

DateStart: 12/2/2014 8:00:00 AM

DateEnd: 2/11/2015 6:00:00 AM

Phase: Pre-registration

DomainCount: 2

TLD: energy

TLDCategory: IDN

TLDStatusID: 4

TLDStatusDesc: General Availability

TLDStatusAdditional: This TLD is now in General Availability, do a domain search and register one today!

SortBy: none

FilterByTLDStatusDesc: none

FilterByPhase: none

SortByDirection: desc

StartPosition: 1

PagingPageSize: 2

TotalRecords: 644

Command: TLD_GETWATCHLISTTLDS

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t

Site: eNom

TimeDifference: +0.00

ExecTime: 0.250

Done: true

TrackingKey: 28c8d57f-843f-49bc-8979-4ad20b4da558

RequestDateTime: 2/6/2015 12:18:19 PM
```
;Machine is SJL0VWRESELL_T1

Server=sjl0vwresell_t1

ExecTime=0.078

TrackingKey=996e815a-eaf9-4303-8a23-4da8c4fc85f0

RequestDateTime=2/6/2015 12:18:47 PM
```