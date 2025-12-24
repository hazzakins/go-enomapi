GetAccountValidation
====================

Retrieve the authorization question and answer for a subaccount.

Usage
-----

Use this command to retrieve the authorization question and answer for a subaccount. Typically, you would use this command to validate the identity of a user before using the GetAccountPassword command to email their password to them.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The AuthQuestionAnswer value must be correct.
- The subaccount must belong to this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetAccountValidation&uid=(Required)&LoginID=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetAccountValidation |
| UID | string | Required | Your Account ID                              |
| PW       | string | Required | Account password |
| LoginID | string | Required | Login ID of the subaccount                        |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ------------------ | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command | string | Name of command executed |
| AuthQuestionType | string | Nature of the question used for identity verification. Permitted values are: smaiden mother’s maiden name sbirth city of birth ssocial last 4 digits of SSN shigh high school fteach favorite teacher fvspot favorite vacation spot fpet favorite pet fmovie favorite movie fbook favorite book |
| AuthQuestionAnswer | string | The answer to the authorization question for the subaccount |
| LoginID | string | The login ID of the subaccount |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getaccountvalidation&uid=resellid
&pw=resellpw&loginid=ichiro&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getaccountvalidation&uid=resellid
&pw=resellpw&loginid=ichiro&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=getaccountvalidation&uid=resellid
&pw=resellpw&loginid=ichiro&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <GetAcctValidation>
 <AuthQuestionType>Mother's Maiden Name</AuthQuestionType>
 <AuthQuestionAnswer>Smith</AuthQuestionAnswer>
 <LoginID>ichiro</LoginID>
 </GetAcctValidation>
 <Command>GETACCOUNTVALIDATION</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod />
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable />
 <IsRealTimeTLD />
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.0390625</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA [ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>AuthQuestionType: </ STRONG>Mother's Maiden Name<br>
<STRONG>AuthQuestionAnswer: </ STRONG>jones<br>
<STRONG>LoginID: </ STRONG>ichiro<br>
<STRONG>Command: </ STRONG>GETACCOUNTVALIDATION<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.422<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 4:10:50 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
AuthQuestionType=Mother's Maiden Name
AuthQuestionAnswer=jones
LoginID=ichiro
Command=GETACCOUNTVALIDATION
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
RequestDateTime=2/3/2015 4:11:16 PM
```
Related Commands
----------------

GetAccountPassword

GetSubAccounts

GetSubaccountsDetailList

SendAccountEmail