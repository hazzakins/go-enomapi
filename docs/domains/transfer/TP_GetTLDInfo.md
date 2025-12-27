TP\GetTLDInfo
=============

Retrieve a list of information required for transferring domains, for each TLD.

Usage
-----

Use this command to list the information that each Registry requires when domains are transferred.

A typical use for this command is administrative. When your customers submit transfer requests, you can use this command to guide you in what is required for processing the orders.

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
https://resellertest.enom.com/interface.asp?command=TP_GetTLDInfo&uid=(Required)&pw=(Required)&QueryFlag=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetBalance |
| uid | string | Required | Your Account ID                                                                                                                |
| pw       | string | Required | Your API Token |
| QueryFlag | string | Required | Which subset of TLDs to return. Permitted values are: 0 All TLDs 1 Transferable TLDs 2 Transferable via auto-verification 3 Transferable via fax 4 Authorization key required for transfer 5 Lockable TLDs 6 Both lockable and transferable |
| ResponseType  | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| TLDIDX | int   | ID number for this TLD, used by our database. Indexed X when ResponseType=Text or HTML. |
| TLDX | string | Top-level domain name \(extension\). Indexed X when ResponseType=Text or HTML. |
| ProtocolX | string | Registry protocol. Indexed X when ResponseType=Text or HTML. |
| AbleToLockX | string | Registry permits locking \(user ability to prevent domain from being transfered without authorization\). Indexed X when ResponseType=Text or HTML. |
| RealTimeX | string | Registry processes domain registrations in real time. Indexed X when ResponseType=Text or HTML. |
| TransferableX | string | Is this TLD transferable. Indexed X when ResponseType=Text or HTML. |
| HasAuthInfoX | string | Does this TLD require an authorization code for transfers? Indexed X when ResponseType=Text or HTML. |
| TransByAutoVeriX | string | Does this Registry permit transfers by auto-verification. Indexed X when ResponseType=Text or HTML. |
| TransByFaxX | string | Does this Registry permit transfers by fax. Indexed X when ResponseType=Text or HTML. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=TP_GETTLDINFO&uid=resellid&pw=resellpw
&QueryFlag=6&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GETTLDINFO&uid=resellid&pw=resellpw
&QueryFlag=6&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=TP_GETTLDINFO&uid=resellid&pw=resellpw
&QueryFlag=6&responsetype=text
```
```
<interface-response>
 <tldtable>
 <tld>
  <TLDID>0</TLDID>
  <TLD>com</TLD>
  <Protocol>RRP</Protocol>
  <AbleToLock>True</AbleToLock>
  <RealTime>True</RealTime>
  <Transferable>True</Transferable>
  <HasAuthInfo>True</HasAuthInfo>
  <TransByAutoVeri>True</TransByAutoVeri>
  <TransByFax>True</TransByFax>
 </tld>
 <tld>
  <TLDID>1</TLDID>
  <TLD>net</TLD>
  <Protocol>RRP</Protocol>
  <AbleToLock>True</AbleToLock>
  <RealTime>True</RealTime>
  <Transferable>True</Transferable>
  <HasAuthInfo>True</HasAuthInfo>
  <TransByAutoVeri>True</TransByAutoVeri>
  <TransByFax>True</TransByFax>
 </tld>
 <tld>
  <TLDID>2</TLDID>
  <TLD>org</TLD>
  <Protocol>EPP</Protocol>
  <AbleToLock>True</AbleToLock>
  <RealTime>True</RealTime>
  <Transferable>True</Transferable>
  <HasAuthInfo>True</HasAuthInfo>
  <TransByAutoVeri>True</TransByAutoVeri>
  <TransByFax>True</TransByFax>
 </tld>
 <tld>
  <TLDID>3</TLDID>
  <TLD>cc</TLD>
  <Protocol>RRP</Protocol>
  <AbleToLock>True</AbleToLock>
  <RealTime>True</RealTime>
  <Transferable>True</Transferable>
  <HasAuthInfo>True</HasAuthInfo>
  <TransByAutoVeri>True</TransByAutoVeri>
  <TransByFax>True</TransByFax>
 </tld>
.
.
.
 </tldtable>
<tldcount>64</tldcount>
 <Command>TP_GETTLDINFO</Command>
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
 <ExecTime>0.875</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 3:35:09 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>TLDID1: </STRONG>0<br />
<STRONG>TLD1: </STRONG>com<br />
<STRONG>Protocol1: </STRONG>RRP<br />
<STRONG>AbleToLock1: </STRONG>True<br />
<STRONG>RealTime1: </STRONG>True<br />
<STRONG>Transferable1: </STRONG>True<br />
<STRONG>HasAuthInfo1: </STRONG>True<br />
<STRONG>HasExtAttributes1: </STRONG>False<br />
<STRONG>TransByAutoVeri1: </STRONG>True<br />
<STRONG>TransByFax1: </STRONG>True<br />
<STRONG>TransPeriod1: </STRONG>1<br />
<STRONG>TransRequiresNewContact1: </STRONG>False<br />
<STRONG>HasPremiumNames1: </STRONG>False<br />
<STRONG>TLDID2: </STRONG>1<br />
<STRONG>TLD2: </STRONG>net<br />
<STRONG>Protocol2: </STRONG>RRP<br />
<STRONG>AbleToLock2: </STRONG>True<br />
<STRONG>RealTime2: </STRONG>True<br />
<STRONG>Transferable2: </STRONG>True<br />
<STRONG>HasAuthInfo2: </STRONG>True<br />
<STRONG>HasExtAttributes2: </STRONG>False<br />
<STRONG>TransByAutoVeri2: </STRONG>True<br />
<STRONG>TransByFax2: </STRONG>True<br />
<STRONG>TransPeriod2: </STRONG>1<br />
<STRONG>TransRequiresNewContact2: </STRONG>False<br />
<STRONG>HasPremiumNames2: </STRONG>False<br />
.
.
.
<STRONG>tldcount: </STRONG>1473<br />
<STRONG>Command: </STRONG>TP_GETTLDINFO<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod:</STRONG><br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable:</STRONG><br />
<STRONG>IsRealTimeTLD:</STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>16.375<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/9/2015 1:03:43 PM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
TLDID1=0
TLD1=com
Protocol1=RRP
AbleToLock1=True
RealTime1=True
Transferable1=True
HasAuthInfo1=True
HasExtAttributes1=False
TransByAutoVeri1=True
TransByFax1=True
TransPeriod1=1
TransRequiresNewContact1=False
HasPremiumNames1=False
TLDID2=1
TLD2=net
Protocol2=RRP
AbleToLock2=True
RealTime2=True
Transferable2=True
HasAuthInfo2=True
HasExtAttributes2=False
TransByAutoVeri2=True
TransByFax2=True
TransPeriod2=1
TransRequiresNewContact2=False
HasPremiumNames2=False
.
.
.
tldcount=1473
Command=TP_GETTLDINFO
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
ExecTime=7.402
Done=true
RequestDateTime=2/9/2015 1:04:29 PM
```
Related Commands
----------------

AddToCart

TP\_CancelOrder

TP\_CreateOrder

TP\_GetOrderStatuses

TP\_SubmitOrder

TP\_UpdateOrderDetailer