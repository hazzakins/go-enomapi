GetSubAccounts
==============

List subaccounts.

Usage
-----

Use this command to list subaccounts for an account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetDomainServices&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetSubAccounts |
| uid | string | Required | Your Account ID                                                                |
| pw       | string | Required | Your API Token |
| ListBy | string | Optional | Sorting parameter. Options are LName, EMailAddress, LoginID.                                         |
| StartLetter   | string | Optional | First letter of the ListBy value |
| StartPosition | string | Optional | First return value in the list generated with the ListBy and StartLetter parameters. Maximum number of subaccounts returned per query is 25. |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| PartyID | string | Party identification number expressed as 32 hexadecimal characters, hyphenated |
| LoginID | string | Login ID of the subaccount |
| DomainCount | string | Number of domains in this subaccount |
| Account | string | Subaccount ID number in NNN-aa-NNNN format |
| Reseller | string | Reseller=0 indicates that this subaccount is a retail account; Reseller=1 indicates this subaccount is a reseller account |
| Count | int   | Number of subaccounts in this account that fit the ListBy and StartLetter criteria |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=GetDomainServices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainServices&uid=resellid&pw=resellpw
&sld=resellerdocs&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainServices&uid=resellid&pw=resellpw
&sld=resellerdocs&ResponseType=text
```
```
<interface-response>
 <SubAccounts>
 <SubAccount>
  <Account>012-gf-7319</Account>
  <DomainCount>0</DomainCount>
  <PartyID>c3f7c836-f3d6-42ab-b628-519bf957786c</PartyID>
  <LoginID>23</LoginID>
  <Reseller>0</Reseller>
  <FName>john</FName>
  <LName>doe</LName>
  <EmailAddress>[email protected]</EmailAddress>
 </SubAccount>
 <SubAccount>
  <Account>946-bw-5211</Account>
  <DomainCount>0</DomainCount>
  <PartyID>80585e17-7a95-4f35-a8cd-c2261b7f2f02</PartyID>
  <LoginID>23142123412341</LoginID>
  <Reseller>0</Reseller>
  <FName>John</FName>
  <LName>Doe</LName>
  <EmailAddress>[email protected]</EmailAddress>
 </SubAccount>
.
.
.
 <EndPosition>31</EndPosition>
 <PreviousRecords>-18</PreviousRecords>
 <NextRecords>32</NextRecords>
 <Count>88</Count>
 <RecordsDisplayed>25</RecordsDisplayed>
 <StartLetter>D</StartLetter>
 <StartPosition>7</StartPosition>
 </SubAccounts>
 <success>True</success>
 <Command>GETSUBACCOUNTS</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod/>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>1.141</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>5de6f4fa-69b3-4107-a34e-a1f6bbeb58f4</TrackingKey>
 <RequestDateTime>12/12/2011 5:42:54 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Account1: </ STRONG>012-gf-7319<br>
<STRONG>DomainCount1: </ STRONG>0<br>
<STRONG>PartyID1: </ STRONG>c3f7c836-f3d6-42ab-b628-519bf957786c<br>
<STRONG>LoginID1: </ STRONG>23<br>
<STRONG>Reseller1: </ STRONG>0<br>
<STRONG>FName1: </ STRONG>john<br>
<STRONG>LName1: </ STRONG>doe<br>
<STRONG>EmailAddress1: </ STRONG>[email protected]<br>
<STRONG>Account2: </ STRONG>946-bw-5211<br>
<STRONG>DomainCount2: </ STRONG>0<br>
<STRONG>PartyID2: </ STRONG>80585e17-7a95-4f35-a8cd-c2261b7f2f02<br>
<STRONG>LoginID2: </ STRONG>23142123412341<br>
<STRONG>Reseller2: </ STRONG>0<br>
<STRONG>FName2: </ STRONG>John<br>
<STRONG>LName2:</ STRONG> Doe<br>
<STRONG>EmailAddress2: </ STRONG>[email protected]<br>
<STRONG>Account3: </ STRONG>842-az-5523<br>
<STRONG>DomainCount3: </ STRONG>0<br>
<STRONG>PartyID3: </ STRONG>193830d8-482f-485f-b281-5e7f5ae3174c<br>
<STRONG>LoginID3: </ STRONG>23442524325<br>
<STRONG>Reseller3: </ STRONG>0<br>
<STRONG>FName3: </ STRONG>John<br>
<STRONG>LName3: </ STRONG>Doe<br>
<STRONG>EmailAddress3: </ STRONG>[email protected]<br>
.
.
.
<STRONG>EndPosition: </ STRONG>31<br>
<STRONG>PreviousRecords: </ STRONG>-18<br>
<STRONG>NextRecords: </ STRONG>32<br>
<STRONG>Count: </ STRONG>98<br>
<STRONG>RecordsDisplayed: </ STRONG>25<br>
<STRONG>StartLetter: </ STRONG>D<br>
<STRONG>StartPosition: </ STRONG>7<br>
<STRONG>success: </ STRONG>True<br>
<STRONG>Command: </ STRONG>GETSUBACCOUNTS<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>1.047<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>ade17f8d-b3f8-455e-b614-2bc4851aac89<br>
<STRONG>RequestDateTime: </ STRONG>2/4/2015 12:43:20 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Account1=012-gf-7319
DomainCount1=0
PartyID1=c3f7c836-f3d6-42ab-b628-519bf957786c
LoginID1=23
Reseller1=0
FName1=john
LName1=doe
[email protected]
Account2=946-bw-5211
DomainCount2=0
PartyID2=80585e17-7a95-4f35-a8cd-c2261b7f2f02
LoginID2=23142123412341
Reseller2=0
FName2=John
LName2=Doe
[email protected]
Account3=842-az-5523
DomainCount3=0
PartyID3=193830d8-482f-485f-b281-5e7f5ae3174c
LoginID3=23442524325
Reseller3=0
FName3=John
LName3=Doe
[email protected]
.
.
.
EndPosition=31
PreviousRecords=-18
NextRecords=32
Count=98
RecordsDisplayed=25
StartLetter=D
StartPosition=7
success=True
Command=GETSUBACCOUNTS
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.156
Done=true
TrackingKey=24406ff8-b632-4c30-99ab-df13a49d91b3
RequestDateTime=2/4/2015 12:44:15 PM
```
Related Commands
----------------

CheckLogin

CreateAccount

CreateSubAccount

DeleteContact

DeleteSubaccount

GetAccountInfo

GetAccountPassword

GetAllAccountInfo

GetCustomerPaymentInfo

GetOrderDetail

GetOrderList

GetReport

GetSubAccountDetails

GetSubaccountsDetailList

GetTransHistory

GetWebHostingAll

SendAccountEmail

SetResellerServicesPricing

SetResellerTLDPricing

SubAccountDomains

UpdateAccountInfo

UpdateCusPreferences