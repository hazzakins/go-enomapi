NM\ExtendOrder
==============

Extend the payment period for a Premium Domain to 10 days from today.

Usage
-----

Use this command to extend the payment deferral period for a Premium Domain. This command can be used for Premium Domains purchased using either the AddToCart or Purchase command with the UseWireTransfer=Yes parameter.

This command extends the payment period to 10 days from today.

> ### A Premium Domain order that was submitted with payment deferral can be extended once. If you do not pay for, extend, or cancel the order, it will automatically cancel 10 days after being placed.

Availability
------------

Premium Domains can only be sold by our direct ETP resellers.

Constraints
-----------
- The login ID and API Token must be valid.
- The Premium Domain order ID must be associated with this account.
- The Premium Domains order must have been submitted with UseWireTransfer=Yes, and must be in "payment pending" status.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------- |
| command     | string | Required | NM\_ExtendOrder |
| uid | string | Required | Your Account ID                               |
| pw       | string | Required | Your API Token |
| OrderID | int  | Required | Identification number of the Premium Domains order you want to extend.   |
| ResponseType  | string | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.* |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Success | string | Success status of this query string |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | True indicates this entire response has reached you successfully. |

Example Output
--------------

The following query extends a Premium Domains order that was submitted with UseWireTransfer=Yes.

```
https://resellertest.enom.com/interface.asp?command=nm_extendorder&uid=resellid&pw=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&orderid=162040096&responsetype={text, XML, or HTML}
```
```
<interface-response>
<Success>True</Success>
<Command>NM_EXTENDORDER</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t1</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.016</ExecTime>
<Done>true</Done>
<TrackingKey>2181198f-7774-4868-bafa-1952c5b1e65d</TrackingKey>
<RequestDateTime>7/3/2016 1:42:23 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success=True
Command=NM_EXTENDORDER
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
ExecTime=0.016
Done=true
TrackingKey=09f7e075-2448-4cda-803b-d9948e93ef95
RequestDateTime=2/11/2015 5:41:17 PM
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>Success: </STRONG>True<BR /><STRONG>Command: </STRONG>NM_EXTENDORDER<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.016<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>93e391d9-24ab-4a67-aecd-d7103c33bd4a<BR /><STRONG>RequestDateTime: </STRONG>7/3/2016 1:45:45 PM<BR /></BODY></HTML>
```
Related Commands
----------------
- [NM\_GetSearchCategories](../docs/nm-getsearchcategories.md)
- [NM\_GetPremiumDomainSettings](../docs/nm-getpremiumdomainsettings.md)
- [NM\_ProcessOrder](../docs/nm-processorder.md)