Portal\GetToken
===============

Get an access token which allows a user to auto login to a tld portal.

Usage
-----

Use this command to get an access token which allow users to auto login to a tld portal.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The Portal User ID must be unique.
- The Email address has to be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| PortalUserID  | Required | Unique user login ID on the reseller site       | 200 |
| Email | Required           | User email address | 300   |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Token | Access token which allows a user to auto login to a tld portal.                 |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves list of TLDs currently offered for this account, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PORTAL_GETTOKEN&UID=ResellID&PW=resellpw
&PortalUserID=uniqueUserLoginID&[email protected]
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=PORTAL_GETTOKEN&UID=ResellID&PW=resellpw
&PortalUserID=uniqueUserLoginID&[email protected]
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=PORTAL_GETTOKEN&UID=ResellID&PW=resellpw
&PortalUserID=uniqueUserLoginID&[email protected]
&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0"?>
<interface-response>
<token>E8888BD8D1D9DDD98DC4D8D1DADEC4DD8C8D8BC4888ADA8FC48D8CDD8CD8DEDC</token>
<Command>PORTAL_GETTOKEN</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>blvdt229</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.065</ExecTime>
<Done>true</Done>
<TrackingKey>4cca5db4-a074-4d8d-baa2-f3b1317f97ff</TrackingKey>
<RequestDateTime>9/26/2012 9:12:12 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
token: FACC99CFCECCCC98CBD6C3CDC2CFD6CFCFCEC2D6C39DC9C9D6C2CBCC9FCDCB9998C3CC9FCD878E889E89BB9E839A968B979ED598949687C9D4CFD4C9CBCACEDBCEC1CEC2C1C8C3DBABB687CA
Command: PORTAL_GETTOKEN
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.031
Done: true
TrackingKey: 79db56f9-02cd-4314-abc6-43e6eb8aabbf
RequestDateTime: 2/4/2015 4:59:38 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
token=81B5E0B6B7B5B5E1B2AFBAB4BBB6AFB6B6B7BBAFBAE4B0B0AFBBB2B5E6B4B2E0E1BAB5E6B4FEF7F1E7F0C2E7FAE3EFF2EEE7ACE1EDEFFEB0ADB6ADB0B2B3B7A2B4B8B2B2B8B2B1A2D2CFFEB3
Command=PORTAL_GETTOKEN
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
TrackingKey=b553260b-2aaa-485c-8797-047295d35662
RequestDateTime=2/4/2015 5:00:03 PM
```
Related Commands
----------------

Portal\_GetAwardedDomains

Portal\_GetDomainInfo

Portal\_UpdateAwardedDomains