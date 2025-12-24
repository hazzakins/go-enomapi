UpdateNotificationAmount
========================

Update the account balance amount at which you want to be notified.

Usage
-----

Use this command to specify the account balance at which you want a reminder to refill your account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/Settings.asp?tab=balance](https://resellertest.enom.com/myaccount/Settings.asp?tab=balance)

The save changes button calls the UpdateNotificationAmount command

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
| Input Parameter | Status | Description                                                        | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                     | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                         | 4 |
| NotifyAmount | Required | Points/$US amount at which you want us to send you a reminder email that this account needs refilling. Use DD.cc format | 9999.99 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query changes the notification amount to 505 points/$US and sends the response in XML format:

```
https://resellertest.enom.com/interface.asp?
command=UpdateNotificationAmount&uid=resellid&pw=resellpw
&NotifyAmount=505&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateNotificationAmount&uid=resellid&pw=resellpw
&NotifyAmount=505&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateNotificationAmount&uid=resellid&pw=resellpw
&NotifyAmount=505&responsetype=text
```
In the response, an ErrCount value 0 confirms that the query was successful:

<Update>Successful</Update>

<Command>UPDATENOTIFICATIONAMOUNT</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod/>

<MaxPeriod>10</MaxPeriod>

<Server>RESELLERTEST</Server>

<Site>enom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>0.0703125</ExecTime>

<Done>true</Done>

<![CDATA[ ] ]>

</debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

Update: Successful

Command: UPDATENOTIFICATIONAMOUNT

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.078

Done: true

RequestDateTime: 2/9/2015 2:22:52 PM
```
;Machine is SJL0VWRESELL_T1

Server=SJL0VWRESELL_T1

ExecTime=0.047

RequestDateTime=2/9/2015 2:23:14 PM
```