GetAccountPassword
==================

Generate a reset password link email for an account.

Usage
-----

Use this command to have reset password link e-mailed to the contact for that account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The AuthQuestionAnswer value must be correct.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetAccountPassword&uid=(Required)&LoginID=(Required)&AuthQuestionAnswer=(Required)&responsetype=(Optional)
```
| Input Parameter  | Type | Status  | Description |
| ------------------ | ------ | -------- | ------------------------------------------------------------------------- |
| command      | string | Required | GetAccountPassword |
| uid | string | Required | Your Account ID                              |
| pw         | string | Required | Your API Token |
| LoginID | string | Required | Login ID for the account for which you want the password.        |
| AuthQuestionAnswer | string | Required | Answer to the authorization question |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| GotAccountInfo  | boolean | True/False                                            |
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
 <GetSubAcctLogin>
 <ResultCode>1</ResultCode>
 <ErrorMessage>Success</ErrorMessage>
 <GotAccountInfo>true</GotAccountInfo>
 </GetSubAcctLogin>
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

SetAccountPassword

UpdateAccountInfo

UpdateCusPreferences