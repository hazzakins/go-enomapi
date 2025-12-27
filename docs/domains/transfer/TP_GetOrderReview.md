TP\GetOrderReview
=================

Retrieve information on a transfer order.

Usage
-----

Use this command to retrieve information on a transfer order— a request originated by you to transfer a name into your account.

This command differs from the TP\_GetOrderDetail command in that TP\_GetOrderDetail retrieves a larger set of information, including the status of the order and detailed contact information.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The transfer order must have originated from this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=
TP_GetOrderReview&uid=(Required)&pw=(Required)&ItemNumber=(Required)&responsetype=xml
```
| Input Parameter | Type | Status                   | Description |
| --------------- | ------ | ------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required                  | TP\_GetOrderReview |
| uid | string | Required | Your Account ID                                                          |
| pw       | string | Required                  | Your API Token |
| TransferOrderID | string | Either ItemNumber or EmptyCart is Required | Item ID number of the item to be deleted from the shopping cart. Use the GetCartContent command to retrieve the item ID numbers. |
| ResponseType  | string | Optional                  | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ------------------------ | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| TransferOrderID | boolean | ID number of the transfer order |
| OrderType | string | Type of the order |
| TransferOrderDetailID | string | Transfer order detail ID number, from our records |
| SLD | string | Second-level domain name \(for example, enom in enom.com\) |
| TLD | string | Top-level domain name \(extension\) |
| Price | int   | Price that will be charged to this account if the transfer is successful |
| Lock | string | Registrar lock setting that was specified in the transfer order |
| Renew | string | Auto-renew setting that was specifie d in the transfer order |
| DomainPassword | string | Domain password, if one was specified in the transfer order |
| UseContacts | string | Use this account’s default contacts |
| AuthInfo | string | EPP Key associated with this domain. Some TLDs require this code to authorize a transfer |
| RRProcessor | string | RRProcessor we use |
| TransferOrderDetailCount | int   | Number of domain names in this order |
| TransferTotalPrice | int   | Total charges if all domains in this order transfer successfully |
| AuthInfoStillRequired | string | Yes indicates we have not yet received the EPP key \(authorization code\) for at least one domain in this order. If you use the TP\_CreateOrder command to begin an order, then allow customers to supply authorization codes on a separate page before using TP\_SubmitOrder, you would use this value to determine whether to display the auth code page. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderReview&uid=resellid
&pw=resellpw&TransferOrderID=175623473
&OrderType=Transfer&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderReview&uid=resellid
&pw=resellpw&TransferOrderID=175623473
&OrderType=Transfer&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GetOrderReview&uid=resellid
&pw=resellpw&TransferOrderID=175623473
&OrderType=Transfer&responsetype=text
```
```html
<?xml version="1.0" ?>
<interface-response>
 <transferorderid>465681</transferorderid>
 <ordertype>Transfer</ordertype>
 <transferorderreview>
 <transferorderdetail>
  <transferorderdetailid>311389</transferorderdetailid>
  <sld>transferdomain</sld>
  <tld>com</tld>
  <price>8.95</price>
  <lock>True</lock>
  <renew>True</renew>
  <domainpassword />
  <usecontacts>1</usecontacts>
  <authinfo />
  <RRProcessor>Reseller Documents Inc.</RRProcessor>
 </transferorderdetail>
 <transferorderdetailcount>1</transferorderdetailcount>
 <transfertotalprice>8.95</transfertotalprice>
 <authinfostillrequired>No</authinfostillrequired>
 </transferorderreview>
 <Command>TP_GETORDERREVIEW</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>e</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>0.1132813</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>transferorderid: </ STRONG>175623473<br>
<STRONG>ordertype: </ STRONG>Transfer<br>
<STRONG>transferorderdetailid1: </ STRONG>77457565<br>
<STRONG>sld1: </ STRONG>00335511<br>
<STRONG>tld1: </ STRONG>net<br>
<STRONG>price1: </ STRONG>8.95<br>
<STRONG>abletolock1: </ STRONG>True<br>
<STRONG>lock1: </ STRONG>True<br>
<STRONG>renew1: </ STRONG>False<br>
<STRONG>domainpassword1:</ STRONG><br>
<STRONG>usecontacts1: </ STRONG>3<br>
<STRONG>reqauthinfo1: </ STRONG>True<br>
<STRONG>authinfo1: </ STRONG>17cb09e58862da1\<br>
<STRONG>RRProcessor: </ STRONG>eNom, Inc.<br>
<STRONG>RegistrantPartyID:</ STRONG><br>
<STRONG>RegistrantROID:</ STRONG><br>
<STRONG>PremiumDomain1: </ STRONG>False<br>
<STRONG>transferorderdetailid2: </ STRONG>77457566<br>
<STRONG>sld2: </ STRONG>accesstestdomain1<br>
<STRONG>tld2: </ STRONG>com<br>
<STRONG>price2: </ STRONG>8.95<br>
<STRONG>abletolock2: </ STRONG>True<br>
<STRONG>lock2: </ STRONG>True<br>
<STRONG>renew2: </ STRONG>False<br>
<STRONG>domainpassword2:</ STRONG><br>
<STRONG>usecontacts2: </ STRONG>3<br>
<STRONG>reqauthinfo2: </ STRONG>True<br>
<STRONG>authinfo2: </ STRONG>d47938502fa59a1$<br>
<STRONG>RRProcessor: </ STRONG>eNom, Inc.<br>
<STRONG>RegistrantPartyID:</ STRONG><br>
<STRONG>RegistrantROID:</ STRONG><br>
<STRONG>PremiumDomain2: </ STRONG>False<br>
<STRONG>transferorderdetailcount: </ STRONG>2<br>
<STRONG>transferorderdetaileucount: </ STRONG>0<br>
<STRONG>transferorderdetailcacount: </ STRONG>0<br>
<STRONG>transferorderdetaildecount: </ STRONG>0<br>
<STRONG>transferorderdetailbecount: </ STRONG>0<br>
<STRONG>transfertotalprice: </ STRONG>17.90<br>
<STRONG>authinfostillrequired: </ STRONG>No<br>
<STRONG>euinfostillrequired: </ STRONG>No<br>
<STRONG>cainfostillrequired: </ STRONG>No<br>
<STRONG>deinfostillrequired: </ STRONG>No<br>
<STRONG>beinfostillrequired: </ STRONG>No<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>Command: </ STRONG>TP_GETORDERREVIEW<br>
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
<STRONG>ExecTime: </ STRONG>0.094<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>40373bb5-bf56-4098-9b2d-9c803e4044cf<br>
<STRONG>RequestDateTime: </ STRONG>2/9/2015 12:55:22 PM<br>
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
transferorderid=175623473
ordertype=Transfer
transferorderdetailid1=77457565
sld1=00335511
tld1=net
price1=8.95
abletolock1=True
lock1=True
renew1=False
domainpassword1=
usecontacts1=3
reqauthinfo1=True
authinfo1=17cb09e58862da1\
RRProcessor=eNom, Inc.
RegistrantPartyID=
RegistrantROID=
PremiumDomain1=False
transferorderdetailid2=77457566
sld2=accesstestdomain1
tld2=com
price2=8.95
abletolock2=True
lock2=True
renew2=False
domainpassword2=
usecontacts2=3
reqauthinfo2=True
authinfo2=d47938502fa59a1$
RRProcessor=eNom, Inc.
RegistrantPartyID=
RegistrantROID=
PremiumDomain2=False
transferorderdetailcount=2
transferorderdetaileucount=0
transferorderdetailcacount=0
transferorderdetaildecount=0
transferorderdetailbecount=0
transfertotalprice=17.90
authinfostillrequired=No
euinfostillrequired=No
cainfostillrequired=No
deinfostillrequired=No
beinfostillrequired=No
ErrCount=0
ResponseCount=0
Command=TP_GETORDERREVIEW
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
ExecTime=0.203
Done=true
TrackingKey=55d16709-114a-4ffb-8478-48f9b84e8fc4
RequestDateTime=2/9/2015 12:55:48 PM
```
Related Commands
----------------

PushDomain

SynchAuthInfo

TP\_CancelOrder

TP\_CreateOrder

TP\_GetDetailsByDomain

TP\_GetOrder

TP\_GetOrderDetail

TP\_GetOrdersByDomain

TP\_GetOrderStatuses

TP\_ResendEmail

TP\_ResubmitLocked

TP\_SubmitOrder

TP\_UpdateOrderDetail

UpdatePushList