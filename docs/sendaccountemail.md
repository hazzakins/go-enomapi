SendAccountEmail
================

Email subaccount login information to the billing contact of record.

Usage
-----

Use this command to email subaccount login information to the billing contact for that subaccount.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The subaccount must be a child of this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                     | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                   | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                       | 4 |
| LoginID | Required | Login ID of subaccount | 20    |
| Account     | Required | ID number of subaccount, in NNN-aa-NNNN format. Use GetSubAccounts command to retrieve the subaccount ID number. | 11 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sends an email to the owner of the subaccount, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=SendAccountEmail&uid=resellid&pw=resellpw
&loginID=resellidsub&account=659-fs-2869&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SendAccountEmail&uid=resellid&pw=resellpw
&loginID=resellidsub&account=659-fs-2869&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SendAccountEmail&uid=resellid&pw=resellpw
&loginID=resellidsub&account=659-fs-2869&responsetype=text
```
In the response, an ErrCount value 0 confirms that the query was successful:

<Command>SENDACCOUNTEMAIL</Command>

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

<ExecTime>0.125</ExecTime>

<Done>true</Done>

<RequestDateTime>12/11/2011 10:31:43 PM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

Command: SENDACCOUNTEMAIL

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site:

TimeDifference: +0.00

ExecTime: 0.656

Done: true

RequestDateTime: 2/5/2015 2:31:32 PM
```
;Machine is SJL0VWRESELL_T1

Server=SJL0VWRESELL_T1

ExecTime=0.063

RequestDateTime=2/5/2015 2:32:54 PM
```