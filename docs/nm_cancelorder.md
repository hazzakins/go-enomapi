NM\CancelOrder
==============

Cancel a Premium Domains order that was created with the AddToCart or Purchase command with the pay-later parameter UseWireTransfer=Yes, but which was never completed.

Usage
-----

Use this command to cancel a Premium Domains order that was created with the AddToCart or Purchase command with the deferred payment parameter UseWireTransfer=Yes, but which was never completed.

If you do not redeem, extend, or cancel an order within 10 days, the order cancels automatically.

> ### This command cancels an entire order. We do not offer a means of splitting an order that includes multiple items.

Availability
------------

Premium Domains can only be sold by our direct ETP resellers.

Constraints
-----------
- The login ID and API Token must be valid.
- The Premium Domain order ID must be associated with this account.
- The Premium Domains order must have been submitted using the UseWireTransfer=Yes parameter and must still be in "Payment Pending" status.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------- |
| command     | string | Required | NM\_CancelOrder |
| uid | string | Required | Your Account ID                               |
| pw       | string | Required | Your API Token |
| OrderID | int  | Required | Identification number for the Premium Domains order to be cancelled     |
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

The following query uses the account balance to pay for a Premium Domains order that was originally submitted with our deferred payment option, UseWireTransfer=Yes

```
https://resellertest.enom.com/interface.asp?command=nm_cancelorder&uid=resellid&pw=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&orderid=162040096&responsetype={text, XML, or HTML}
```
```
<interface-response>
<Success>True</Success>
<Command>NM_CANCELORDER</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>1.422</ExecTime>
<Done>true</Done>
<TrackingKey>25750bcd-4323-4721-b3b8-3ebb29bb2afc</TrackingKey>
<RequestDateTime>7/4/2016 6:58:28 AM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL1VWRESELL_T
;Encoding Type is utf-8
Success=True
Command=NM_CANCELORDER
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl1vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.016
Done=true
TrackingKey=2169bd44-b450-4a63-b810-7a5e12a31501
RequestDateTime=7/3/2016 1:44:54 PM
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>Success: </STRONG>True<BR /><STRONG>Command: </STRONG>NM_CANCELORDER<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.016<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>93e391d9-24ab-4a67-aecd-d7103c33bd4a<BR /><STRONG>RequestDateTime: </STRONG>7/3/2016 1:45:45 PM<BR /></BODY></HTML>
```
Related Commands
----------------
- [NM\_Search](../docs/nm-search.md)
- [NM\_ProcessOrder](../docs/nm-processorder.md)
- [NM\_ExtendOrder](../docs/nm-extendorder.md)