TLD\DeleteWatchlist
===================

Delete a set of domains from a LIST or delete an entire LIST from the watchlist. Deleting all domains in a LIST does not delete the LIST

Usage
-----

Use this command to delete a set of domains from a LIST or delete an entire LIST from the watchlist. Deleting all domains in a LIST does not delete the LIST

List
----

You can create LIST\(s\) for organizing or managing watchlist domain names. Each LIST consists of a ListID, ListName and ListEmail.

Each account has a default LIST \(ListID=0 \), containing domains to watch. Domains can appear in multiple LISTS. The default LIST cannot be deleted and email address cannot be changed.

| ListID=0                              | ListID=1 | ListID=2           |
| ------------------------------------------------------------------- | ---------------------------- | ---------------------------- |
| - .ninja .art .education one.ceo one.ninja two.ceo two.ninja | one.ceo one.ninja one.fun | two.ceo two.ninja two.fun |

Operational/Logic Table:

| ListID     | ListName | ListEmail | Result |
| -------------- | -------- | --------- | --------------------------------------------------------------- |
| DeleteList   | - | No    | Error: ListID must be supplied |
| DeleteList | -    | Yes | Delete the LIST and all domains in it              |
| DeleteListItem | No | -     | Error: WatchlistItemID must be supplied |
| DeleteListItem | Yes   | No | Delete WatchlistItemID\(s\) in the default account \(ListID=0\) |
| DeleteListItem | Yes | Yes    | Delete WatchlistItemID\(s\) in the specified ListID |

Note:

No = input parameter is not declared

Yes = input parameter is declared and has value
- = input parameter could be declared or not declared. Does not affect the end result

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/tlds/watchlist.aspx](https://resellertest.enom.com/tlds/watchlist.aspx)

The Delete button calls the TLD\_DeleteWatchlist command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status   | Description | Max Size |
| --------------- | ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID | Required  | Account login ID | 20    |
| PW       | Required | Account password                                                                                   | 20 |
| ResponseType | Optional  | Format of response. Permitted values are Text \(default\), HTML, or XML | 4    |
| Action     | Required | Type of delete operation. Permitted values are: DeleteList DeleteListItem                                                      | 8 |
| ListID | Required  | The ID for a specified LIST in the account. See Operational/Logic table for conditional information. Use the TLD\_GetWatchlist with Action=GetListIDs command to retrieve the LIST |     |
| WatchlistItemID | Conditional | A comma-separated list of WatchlistItemID. Use the TLD\_GetWatchlist command to retrieve these IDs                                         | |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| ID | WatchlistItemID for each domain name                               |
| DomainName    | Domain name |
| Success | Success status for deleting domain name from the Watchlist                    |
| ResultMessage  | Message returned from the processing system |
| Type | Reserved for future implementation                                |
| FailCount    | Number of domain\(s\) failed to add to the watchlist |
| SuccessCount | Number of domain\(s\) successfully added to the Watchlist                    |
| TotalRecords   | Total number of domain\(s\) processed |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query deletes list of domain\(s\) or a LIST from the Watchlist, and requests the response in given format.

Delete domain\(s\)
------------------

```
https://resellertest.enom.com/interface.asp?
command=TLD_DeleteWatchlist&UID=ResellID&PW=resellpw
&ResponseType=XML&Action=DeleteListItem
&WatchlistIDList=73995&ListID=23
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_DeleteWatchlist&UID=ResellID&PW=resellpw
&ResponseType=html&Action=DeleteListItem
&WatchlistIDList=73995&ListID=23
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_DeleteWatchlist&UID=ResellID&PW=resellpw
&ResponseType=text&Action=DeleteListItem
&WatchlistIDList=73995&ListID=23
```
```
<?xml version="1.0"?>

  <ID>73995</ID>

  <DomainName>73995</DomainName>

  <Success>True</Success>

  <ResultMessage><![CDATA[ ]]></ResultMessage />

 </Item>

 <FailCount>0</FailCount>

 <SuccessCount>1</SuccessCount>

 <TotalRecords>1</TotalRecords>

 <Type>Account</Type>

 <Action>DELETELISTITEM</Action>

 <ListID>23</ListID>

 </TLD_Watchlist>

 <Command>TLD_DELETEWATCHLIST</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>eNom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>2.141</ExecTime>

 <Done>true</Done>

 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>

 <RequestDateTime>7/16/2012 3:55:55 PM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

ID: 73995

DomainName: 73995

Success: True

ResultMessage:

FailCount: 0

SuccessCount: 1

TotalRecords: 1

Type: Account

Action: DELETELISTITEM

ListID: 23

Command: TLD_DELETEWATCHLIST

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.016

Done: true

TrackingKey: 526d1c89-c2bd-423b-bb7c-0e4c94b3cdb3

RequestDateTime: 2/6/2015 11:59:00 AM
```
TrackingKey=3bfbca87-442d-40cd-829b-05f0582d76c6

RequestDateTime=2/6/2015 12:03:49 PM
```
Delete LIST
-----------

```
https://resellertest.enom.com/interface.asp?
command=TLD_DeleteWatchlist&UID=ResellID&PW=resellpw
&ResponseType=XML&Action=DeleteList
&WatchlistIDList=&ListID=23
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_DeleteWatchlist&UID=ResellID&PW=resellpw
&ResponseType=html&Action=DeleteList
&WatchlistIDList=&ListID=23
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_DeleteWatchlist&UID=ResellID&PW=resellpw
&ResponseType=text&Action=DeleteList
&WatchlistIDList=&ListID=23
```
 <ListID>4</ListID>

;Machine is SJL0VWRESELL_T

Action: DELETELIST

count: 1

Server: sjl0vwresell_t

ExecTime: 0.063

TrackingKey: 38108a18-2a78-4aa8-8fe9-c23cb278c0fe

RequestDateTime: 2/6/2015 12:02:28 PM
```
TrackingKey=d7457457-885a-4fa4-bdd5-6b3e4a40f899

RequestDateTime=2/6/2015 12:03:13 PM
```