TLD\AddWatchlist
================

Add a list of domain\(s\) to an existing the watchlist or create new watchlist with these domains. An empty watchlist cannot be created.

Usage
-----

Use this command to add a list of domain\(s\) to an existing the watchlist or create new watchlist with these domains. An empty watchlist cannot be created.

List
----

You can create LIST\(s\) for organizing or managing watchlist domain names. Each LIST consist of a ListID, ListName and ListEmail.

Each account has a default LIST \(ListID=0 \), containing domains to watch. Domains can appear in multiple LISTS. The default LIST cannot be deleted and email address cannot be changed.

| ListID=0                              | ListID=1 | ListID=2           |
| ------------------------------------------------------------------- | ---------------------------- | ---------------------------- |
| - .ninja .art .education one.ceo one.ninja two.ceo two.ninja | one.ceo one.ninja one.fun | two.ceo two.ninja two.fun |

Operational/Logic Table:

| ListID | ListName | ListEmail | Result |
| ------ | -------- | --------- | ------------------------------------------------------------------------------------------------------- |
| No   | No | -     | Error: ListID or ListName are required |
| No | Yes   | No | Error: Email Address is required if ListName is sent                          |
| No   | Yes | Yes    | Create a new LIST and add domain\(s\). ListID will be automatically created and returned in the output |
| Yes | -    | No | Add domain\(s\) to specified LIST                                    |
| Yes  | - | Yes    | Add domain\(s\) to specified LIST and update email address |

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

The Watch Domains button calls the TLD\_AddWatchlist command.

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
| Input Parameter | Status          | Description | Max Size |
| --------------- | ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID | Required         | Account login ID | 20    |
| PW       | Required | Account password                                                                                                                                                                                                                                                                                                                                             | 20 |
| ResponseType | Optional         | Format of response. Permitted values are Text \(default\), HTML, or XML | 4    |
| DomainList   | Required | A comma-separated list of domains Format: blogging.ninja - .ninja blogging.\* International Domain Name \(IDN\) allows more web users to navigate the Internet in their preferred native language. Most domain names are registered in ASCII characters. The non-Latin scripts such as Chinese, Russian, Korean or other languages cannot be rendered in ASCII. IDN is fully supported in PUNY Code and URL Encoded format. To get more information and conversion tool for PUNY Code, check the sites below: - Verisign's IDN Tool: [http://mct.verisign-grs.com](http://mct.verisign-grs.com/) - RFC 3492: [http://tools.ietf.org/html/rfc3492](http://tools.ietf.org/html/rfc3492) 4096 | 4096 |
| ListID | Conditional; default is 0 | List ID. See Operational/Logic table for conditional information Use the TLD\_GetWatchlist with Action=GetListID command to retrieve LIST |     |
| ListName    | Conditional | Name of the LIST \(must be unique\)                                                                                                                                                                                                                                                                                                                                   | 50 |
| ListEmail | Conditional        | Email address for the List | 128   |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| ID        | Watchlist item ID for each domain |
| DomainName | Domain name                                           |
| Success     | Domain add status |
| ResultMessage | Domain add message                                        |
| FailCount    | Number of domain\(s\) failed to add to the Watchlist |
| SuccessCount | Number of domain\(s\) successfully added to the Watchlist                    |
| TotalRecords   | Total number of domain\(s\) processed |
| Type | Reserved for future implementation                                |
| ListID      | List ID. New number will be created whenever a new LIST is created |
| ListName | List Name                                            |
| ListEmail    | List Email Address |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query adds domain\(s\) to the watchlist, and requests the response in given format.

Add to default watchlist
------------------------

```
https://resellertest.enom.com/interface.asp?
command=TLD_Addwatchlist&UID=ResellID&PW=resellpw
&ResponseType=html
&DomainList=blogging.ninja,*.ninja,blogging.*
&ListID=&ListName=&ListEmail=
```
```
https://resellertest.enom.com/interface.asp?
command=TEL_UpdatePrivacy&uid=resellid&pw=resellpw
&domainname=resellerdocs.tel&telwhoistype=natural_person
&telpublishwhois=no&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_Addwatchlist&UID=ResellID&PW=resellpw
&ResponseType=text
&DomainList=blogging.ninja,*.ninja,blogging.*
&ListID=&ListName=&ListEmail=
```
```
<?xml version="1.0"?>

  <ID>74670</ID>

  <DomainName><![CDATA[ blogging.ninja ]]></DomainName>

  <Success>True</Success>

 </Item>

  <ID>74671</ID>

  <DomainName><![CDATA[ domain.ninja ]]></DomainName>

.....

 <FailCount>0</FailCount>

 <SuccessCount>3</SuccessCount>

 <TotalRecords>3</TotalRecords>

 <Type>Account</Type>

 <ListID>0</ListID>

 <ListName />

 <ListEmail />

 </TLD_Watchlist>

 <Command>TLD_ADDWATCHLIST</Command>

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

 <ExecTime>0.594</ExecTime>

 <Done>true</Done>

 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>

 <RequestDateTime>7/16/2012 2:52:25 PM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

ID: 74670

DomainName: blogging.ninja

Success: True

ResultMessage:

ID: 74671

.

.

.

FailCount: 0

SuccessCount: 3

TotalRecords: 3

Type: Account

ListID: 0

Command: TLD_ADDWATCHLIST

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t

Site: eNom

TimeDifference: +0.00

ExecTime: 0.016

Done: true

TrackingKey: 27ce2251-2fca-4c48-a02a-8ec07cfe6ecd

RequestDateTime: 2/6/2015 11:38:04 AM
```
.

.

.

TrackingKey=13fad540-a864-491d-8362-e70648ab62d3

RequestDateTime=2/6/2015 11:49:56 AM
```
Create a new LIST and add domains
---------------------------------

```
https://resellertest.enom.com/interface.asp?
command=TLD_AddWatchlist&UID=ResellID&PW=resellpw
&ResponseType=XML
&domainlist=blogging.ninja,*.ninja,blogging.*
&ListID=&ListName=Customer1&[email protected]
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_AddWatchlist&UID=ResellID&PW=resellpw
&ResponseType=html
&domainlist=blogging.ninja,*.ninja,blogging.*
&ListID=&ListName=Customer1&[email protected]
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_AddWatchlist&UID=ResellID&PW=resellpw
&ResponseType=text
&domainlist=blogging.ninja,*.ninja,blogging.*
&ListID=&ListName=Customer1&[email protected]
```
  <ID>74681</ID>

  <ResultMessage><![CDATA[ ]]> </ResultMessage>

  <ID>74682</ID>

.....

 <ListID>33</ListID>

 <ListName>customer1</ListName>

 <ListEmail>[email protected]</ListEmail>

ID: 74681

ID: 74682

DomainName: domain.ninja

.

.

.

ListID: 33

ListName: customer1

ListEmail: [email protected]

.

.

.

[email protected]

Add to a specific LIST
----------------------

```
https://resellertest.enom.com/interface.asp?
command=TLD_AddWatchlist&UID=ResellID&PW=resellpw
&ResponseType=XML
&domainlist=blogging.ninja,*.ninja,blogging.*
&ListID=16&ListName=&ListEmail=
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_AddWatchlist&UID=ResellID&PW=resellpw
&ResponseType=html
&domainlist=blogging.ninja,*.ninja,blogging.*
&ListID=16&ListName=&ListEmail=
```
```
https://resellertest.enom.com/interface.asp?
command=TLD_AddWatchlist&UID=ResellID&PW=resellpw
&ResponseType=text
&domainlist=blogging.ninja,*.ninja,blogging.*
&ListID=16&ListName=&ListEmail=
```
 <ID>74698</ID>

 <ID>74699</ID>

.....

ID: 74698

ID: 74699

.

.

.

.

.

.