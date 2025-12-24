GetGlobalChangeStatus
=====================

Retrieve a list of IDs for instances in which a global edit command was run on this account.

Usage
-----

Use this command to retrieve RequestID values for global edits performed on domains in this account. Commonly, you will run this command to get a RequestID in order to run the GetGlobalChangeStatusDetail command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/globaledit.asp](https://resellertest.enom.com/domains/globaledit.asp)

The Click here to view the status and history of your previous global edits link calls the GetGlobalChangeStatus command.

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
| Input Parameter | Status | Description                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| RequestID | ID number of the request for the change described in this node                  |
| ReqDate     | Time-stamp for this request |
| ReqTypeDesc | Description of the requested change                               |
| ReqStatusDesc  | Status of the requested change |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query Retrieves the global changes that have been requested in account resellid and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetGlobalChangeStatus&UID=resellid
&PW=resellpw&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetGlobalChangeStatus&UID=resellid
&PW=resellpw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetGlobalChangeStatus&UID=resellid
&PW=resellpw&ResponseType=text
```
The response is as follows:

<RequestID>11016</RequestID>

<ReqDate>11/10/2008 12:47:03 PM</ReqDate>

<ReqTypeDesc>Custom DNS changes</ReqTypeDesc>

<ReqStatusDesc>Complete</ReqStatusDesc>

</GlobalStatusInfo>

<RequestID>10307</RequestID>

<ReqDate>3/7/2008 9:00:07 AM</ReqDate>

<RequestID>10306</RequestID>

<ReqDate>3/7/2008 8:50:35 AM</ReqDate>

<RequestID>10305</RequestID>

<ReqDate>3/7/2008 6:43:10 AM</ReqDate>

<RequestID>9908</RequestID>

<ReqDate>11/9/2006 4:01:29 PM</ReqDate>

<RequestID>8893</RequestID>

<ReqDate>2/24/2004 11:02:58 AM</ReqDate>

<ReqTypeDesc>Registrar-lock off</ReqTypeDesc>

<RequestID>8892</RequestID>

<ReqDate>2/23/2004 7:03:53 PM</ReqDate>

<RequestID>8891</RequestID>

<ReqDate>2/23/2004 6:55:23 PM</ReqDate>

<RequestID>8890</RequestID>

<ReqDate>2/23/2004 6:44:18 PM</ReqDate>

<RequestID>8889</RequestID>

<ReqDate>2/23/2004 6:41:58 PM</ReqDate>

<ReqTypeDesc>Registrar-lock on</ReqTypeDesc>

<RequestID>8343</RequestID>

<ReqDate>9/3/2003 5:29:13 PM</ReqDate>

<RequestID>8342</RequestID>

<ReqTypeDesc>Auto renew on</ReqTypeDesc>

<RequestID>8341</RequestID>

<ReqDate>9/3/2003 5:22:27 PM</ReqDate>

<RequestID>8340</RequestID>

<RequestID>8339</RequestID>

<ReqDate>9/3/2003 5:12:04 PM</ReqDate>

<ReqTypeDesc>Contact changes</ReqTypeDesc>

<RequestID>8210</RequestID>

<ReqDate>4/25/2003 3:00:03 AM</ReqDate>

<Command>GETGLOBALCHANGESTATUS</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod/>

<MaxPeriod>10</MaxPeriod>

<Server>SJL21WRESELLT01</Server>

<Site>eNom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>0.453</ExecTime>

<Done>true</Done>

<RequestDateTime>12/8/2011 4:57:39 AM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

RequestID: 14917

ReqDate: 8/16/2014 2:38:31 AM

ReqTypeDesc: Domain password changes

ReqStatusDesc: Complete

RequestID: 14916

ReqTypeDesc: Auto renew on

RequestID: 14915

ReqDate: 8/11/2014 11:55:04 PM

RequestID: 14914

RequestID: 14091

ReqDate: 1/4/2012 3:10:33 PM

RequestID: 14090

ReqDate: 1/4/2012 3:08:57 PM

ReqTypeDesc: Custom DNS changes

RequestID: 14089

ReqDate: 1/4/2012 3:07:38 PM

RequestID: 11016

ReqDate: 11/10/2008 12:47:03 PM

RequestID: 10307

ReqDate: 3/7/2008 9:00:07 AM

RequestID: 10306

ReqDate: 3/7/2008 8:50:35 AM

RequestID: 10305

ReqDate: 3/7/2008 6:43:10 AM

RequestID: 9908

ReqDate: 11/9/2006 4:01:29 PM

RequestID: 8893

ReqDate: 2/24/2004 11:02:58 AM

ReqTypeDesc: Registrar-lock off

RequestID: 8892

ReqDate: 2/23/2004 7:03:53 PM

RequestID: 8891

ReqDate: 2/23/2004 6:55:23 PM

RequestID: 8890

ReqDate: 2/23/2004 6:44:18 PM

RequestID: 8889

ReqDate: 2/23/2004 6:41:58 PM

ReqTypeDesc: Registrar-lock on

RequestID: 8343

ReqDate: 9/3/2003 5:29:13 PM

RequestID: 8342

RequestID: 8341

ReqDate: 9/3/2003 5:22:27 PM

RequestID: 8340

RequestID: 8339

ReqDate: 9/3/2003 5:12:04 PM

ReqTypeDesc: Contact changes

RequestID: 8210

ReqDate: 4/25/2003 3:00:03 AM

Command: GETGLOBALCHANGESTATUS

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.172

Done: true

RequestDateTime: 2/4/2015 10:30:59 AM
```
;Machine is SJL0VWRESELL_T

Server=SJL0VWRESELL_T

ExecTime=0.063

RequestDateTime=2/4/2015 10:31:36 AM
```