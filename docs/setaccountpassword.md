SetAccountPassword
==================

Set the password for an account based off a reset password token

Usage
-----

Use this command to update password for an account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The Token value must be correct.
- The Token has not expired.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetAccountPassword&uid=(Required)&pw=(Requred)&Token=(Required)&Password=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------------ |
| command     | string | Required | SetAccountPassword |
| uid | string | Required | Your Account ID                                |
| pw       | string | Required | Your API Token |
| token | string | Required | Reset Password Token found in reset password email sent from GetAccounPassword |
| password    | string | Required | New Password for that account. |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML   |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| Success     | boolean | "True" means password was updated. "False" means the password was not updated.         |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolean | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetAccountPassword&nAnswuid=resellid&pw=resellpw
&LoginID=ichiro&AuthQuestionAnswer=Jones&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetAccountPassword&nAnswuid=resellid&pw=resellpw
&LoginID=ichiro&AuthQuestionAnswer=Jones&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetAccountPassword&nAnswuid=resellid&pw=resellpw
&LoginID=ichiro&AuthQuestionAnswer=Jones&ResponseType=text
```
```
<interface-response>
 <Success>True</Sucess>
 <Command>GETACCOUNTPASSWORD</Command>
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
 <ExecTime>0.438</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/7/2011 6:16:27 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>ResultCode: </ STRONG>1<br>
<STRONG>ErrorMessage: </ STRONG>Success<br>
<STRONG>GotAccountInfo: </ STRONG>true<br>
<STRONG>Command: </ STRONG>GETACCOUNTPASSWORD<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T1<br>
<STRONG>Site:</ STRONG><br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.016<br>
<STRONG>Done: </ STRONG>true<br>
RequestDateTime: </ STRONG>2/11/2015 1:34:12 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
ResultCode=1
ErrorMessage=Success
GotAccountInfo=true
Command=GETACCOUNTPASSWORD
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
RequestDateTime=2/11/2015 1:35:31 PM
```
Related Commands
----------------

CheckLogin

CreateAccount

CreateSubAccount

GetAccountInfo

GetAccountValidation

GetAllAccountInfo

GetOrderDetail

GetOrderList

GetReport

GetSubAccountDetails

GetSubAccounts

GetTransHistory

SendAccountEmail

UpdateAccountInfo

UpdateCusPreferences