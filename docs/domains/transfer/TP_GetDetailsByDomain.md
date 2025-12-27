TP\GetDetailsByDomain
=====================

Get transfer order information for a domain using sld.tld.

Usage
-----

Use this command to retrieve status information on one domain name that is in the process of transferring.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must be in a transfer order that belongs to this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=TP_GetDetailsByDomain&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | TP\_GetDetailsByDomain |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

TransferOrderDetail StatusID is the status of this domain in the transfer process:

0=Transfer request created - awaiting fax

5=Transferred successfully

9=Awaiting auto verification of transfer request \(no longer used due to GDPR\)

10=Unable to retrieve current domain contacts from UWhois \(no longer used due to GDPR\)

11=Auto verification of transfer request initiated \(no longer used due to GDPR\)

12=Awaiting for auto transfer string validation

13=Domain awaiting transfer initiation

14=Domain transfer initiated and awaiting approval

15=Canceled - cannot obtain domain contacts from UWhois

16=Canceled - domain contacts did not respond to verification e-mail

17=Canceled - domain contacts did not approve transfer of domain

18=Canceled - domain validation string is invalid

19=Canceled - Whois information provided does not match current registrant

20=Canceled - Domain is currently not registered and cannot be transferred

21=Canceled - Domain is already registered in account and cannot be transferred

22=Canceled - Domain is locked at current registrar

23=Canceled - Transfer already initiated for this domain

24=Canceled - Unable to transfer due to unknown error

25=Canceled - Transfer rejected by losing registrar

26=Canceled - Transfer authorization fax not received

27=Canceled by customer

28=Fax received - awaiting registrant verification

29=Awaiting manual fax verification

30=Canceled - Domain name is invalid or is Invalid for Transfers

35=Transfer request not yet submitted.

100=Pending consent - The transfer order is waiting for GDPR consent to be given.

101=Canceled - Registrant denied - GDPR consent was refused, and so the order was cancelled.

Example

| Output Parameter      | Type | Description                                                                               |
| -------------------------- | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command          | string | Name of command executed                                                                        |
| OrderCount         | int | Number of transfer orders that have been submitted for this domain.                                                  |
| OrderIDX X=1 to OrderCount | int | Transfer order detail number. Indexed X when ResponseType=Text or HTML.                                                |
| OrderDateX         | string | The date the order was submitted. Indexed X when ResponseType=Text or HTML.                                              |
| StatusIDX         | int | TransferOrderDetail status ID -- a number that indicates the status of this domain in the transfer process. See Notes. Indexed X when ResponseType=Text or HTML.  |
| StatusDescX        | string | TransferOrderDetail status description—a text description of the status of this domain in the transfer process. See Notes. Indexed X when ResponseType=Text or HTML. |
| ErrCount          | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                    |
| ErrX            | string | Error messages explaining the failure. These can be presented as is back to the client.                                        |
| Done            | boolian | "True" indicates this entire response has reached you successfully.                                                  |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GetDetailsByDomain&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetDetailsByDomain&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetDetailsByDomain&uid=resellid&pw=resellpw
&sld=resellerdocs2&tld=net&responsetype=text
```
```
<interface-response>
 <TransferOrder>
 <orderid>175581269</orderid>
 <orderdate>9/27/2010 2:26:09 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175581275</orderid>
 <orderdate>9/27/2010 2:46:50 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175581278</orderid>
 <orderdate>9/27/2010 3:43:14 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175581281</orderid>
 <orderdate>9/27/2010 3:57:16 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175581284</orderid>
 <orderdate>9/27/2010 3:57:54 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175581341</orderid>
 <orderdate>9/27/2010 9:47:10 PM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175585289</orderid>
 <orderdate>11/3/2010 7:22:14 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175585292</orderid>
 <orderdate>11/3/2010 7:26:00 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175587221</orderid>
 <orderdate>11/19/2010 9:47:07 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175587224</orderid>
 <orderdate>11/19/2010 9:47:40 AM</orderdate>
 <statusid>32</statusid>
 <statusdesc>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key</statusdesc>
 </TransferOrder>
 <TransferOrder>
 <orderid>175598255</orderid>
 <orderdate>4/5/2011 9:23:37 AM</orderdate>
 <statusid>15</statusid>
 <statusdesc>Canceled - cannot obtain domain contacts from UWhois</statusdesc>
 </TransferOrder>
.
.
.
 <ordercount>38</ordercount>
 <Command>TP_GETDETAILSBYDOMAIN</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+8.00</TimeDifference>
 <ExecTime>0.141</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>c3f152ab-e92c-4de5-8193-1898f91a6bf9</TrackingKey>
 <RequestDateTime>12/12/2011 1:59:03 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>orderid1: </STRONG>175581269<br />
<STRONG>orderdate1: </STRONG>9/27/2010 2:26:09 AM<br />
<STRONG>statusid1: </STRONG>32<br />
<STRONG>statusdesc1: </STRONG>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key<br />
<STRONG>orderid2: </STRONG>175581275<br />
<STRONG>orderdate2: </STRONG>9/27/2010 2:46:50 AM<br />
<STRONG>statusid2: </STRONG>32<br />
<STRONG>statusdesc2: </STRONG>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key<br />
<STRONG>orderid3: </STRONG>175581278<br />
<STRONG>orderdate3: </STRONG>9/27/2010 3:43:14 AM<br />
<STRONG>statusid3: </STRONG>32<br />
<STRONG>statusdesc3: </STRONG>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key<br />
<STRONG>orderid4: </STRONG>175581281<br />
<STRONG>orderdate4: </STRONG>9/27/2010 3:57:16 AM<br />
<STRONG>statusid4: </STRONG>32<br />
<STRONG>statusdesc4: </STRONG>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key<br />
<STRONG>orderid5: </STRONG>175581284<br />
<STRONG>orderdate5: </STRONG>9/27/2010 3:57:54 AM<br />
<STRONG>statusid5: </STRONG>32<br />
<STRONG>statusdesc5: </STRONG>Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key<br />
.
.
.
<STRONG>ordercount: </STRONG>84<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>Command: </STRONG>TP_GETDETAILSBYDOMAIN<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>0.078<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>e8445307-2184-4a2d-9896-b2a3bb86367c<br />
<STRONG>RequestDateTime: </STRONG>2/9/2015 12:29:02 PM<br />
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
orderid1=175581269
orderdate1=9/27/2010 2:26:09 AM
statusid1=32
statusdesc1=Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key
orderid2=175581275
orderdate2=9/27/2010 2:46:50 AM
statusid2=32
statusdesc2=Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key
orderid3=175581278
orderdate3=9/27/2010 3:43:14 AM
statusid3=32
statusdesc3=Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key
orderid4=175581281
orderdate4=9/27/2010 3:57:16 AM
statusid4=32
statusdesc4=Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key
orderid5=175581284
orderdate5=9/27/2010 3:57:54 AM
statusid5=32
statusdesc5=Canceled - Invalid EPP/authorization key - Please contact current registrar to obtain correct key
.
.
.
ordercount=84
ErrCount=0
ResponseCount=0
Command=TP_GETDETAILSBYDOMAIN
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.047
Done=true
TrackingKey=88bd3ba0-c27a-498c-81bf-34f617df3c35
RequestDateTime=2/9/2015 12:30:46 PM
```
Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

TP\_CreateOrder

TP\_GetOrder

TP\_GetOrderDetail

TP\_GetOrderReview

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList