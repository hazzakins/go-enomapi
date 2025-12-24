GetAccountInfo
==============

Get account ID, password, authorization question and answer, reseller flag, and credit card agreement settings.

Usage
-----

Use this command to return a short list of account identification information and reseller status.

GetAllAccountInfo, a similar command, returns a complete list of account information.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command runs in the background of eNom’s Web site.

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

| Output Parameter | Description                                                                                                                                           |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                                                                                                                     |
| ErrCount      | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                    |
| Done        | True indicates this entire response has reached you successfully. |
| UserID | Current login ID                                                                                                                                         |
| Password      | Current account password |
| AuthQuestionType | Nature of the question used for identity verification. Permitted values are: smaiden mother’s maiden name sbirth city of birth ssocial last 4 digits of SSN shigh high school fteach favorite teacher fvspot favorite vacation spot fpet favorite pet fmovie favorite movie fbook favorite book |
| AuthQuestionAnswer | Current answer to the authorization question |
| Account | Account number                                                                                                                                          |
| Reseller      | Is this a reseller account or not |
| AcceptTerms | Returns True if the credit card agreement has been signed, False otherwise                                                                                                            |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests a limited list of account information for account resellid, and requests the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetAccountInfo&uid=resellid&pw=resellpw
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetAccountInfo&uid=resellid&pw=resellpw
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetAccountInfo&uid=resellid&pw=resellpw
&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" ?>

 <UserID>resellid</UserID>

 <Password>resellpw</Password>

 <AuthQuestionType>smaiden</AuthQuestionType>

 <AuthQuestionAnswer>Jones</AuthQuestionAnswer>

 <Account>217-no-0647</Account>

 <Reseller>True</Reseller>

 <AcceptTerms>True</AcceptTerms>

 <Command>GETACCOUNTINFO</Command>

 <ErrCount>0</ErrCount>

 <Server>ResellerTest</Server>

 <Site>enom</Site>

 <Done>true</Done>

 <![CDATA [ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

UserID: resellid

Password: resellpw

AuthQuestionType: fpet

AuthQuestionAnswer: kitty

Account: 217-no-0647

Reseller: True

AcceptTerms: True

Command: GETACCOUNTINFO

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod:

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

IsLockable:

IsRealTimeTLD:

TimeDifference: +0.00

ExecTime: 0.031

Done: true

RequestDateTime: 2/3/2015 4:06:35 PM
```
RequestDateTime=2/3/2015 4:07:01 PM
```