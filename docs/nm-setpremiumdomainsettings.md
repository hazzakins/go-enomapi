NM\SetPremiumDomainSettings
===========================

Set the threshold at which you want to require payment for Premium Domains by wire transfer to us.

Usage
-----

Use this command to set the threshold price above which you want to always pay your Premium Domains order by wire transfer. You can use this setting as a means of ensuring that orders don’t fail due to an insufficient account balance.

You also have the option of deferring payment at lower price levels, using the UseWireTransfer parameter in the Purchase or InsertNewOrder commands.

The account-level deferred payment threshold that you set with the NM\_SetPremiumDomainSettings command applies in the following ways:
- If you use the Purchase command to purchase domain names, the threshold applies to the price of the single Premium Domain in the query string.
- If you use AddToCart and InsertNewOrder to purchase domain names, the threshold applies to the sum of the Premium Domains, plus the services attached to them, in an order.

> ### There are a few considerations within our system to be aware of, to guide you in choosing a threshold for deferred payment:
>
>
>
> - We place a limit on orders that use our merchant services to process credit cards. If you normally use our merchant services and want to offer your customers Premium Domains over that limit in value, you must establish a system to pay us, and to charge your customers, that is outside our merchant services environment. The lower the wire transfer threshold you choose, the more likely you are to need this system.
> - When you’re ready to submit your deferred payment, use the [NM\_ProcessOrder](../docs/nm-processorder.md) command to complete the order using your account balance, or contact your sales representative for wire transfer instructions.

Availability
------------

Premium Domains can only be sold by our direct ETP resellers.

Constraints
-----------

The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter    | Type | Status          | Description |
| --------------------- | ------- | ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command        | string | Required         | NM\_SetPremiumDomainSettings |
| uid | string | Required | Your Account ID                                                                                      |
| pw          | string | Required         | Your API Token |
| WireTransferThreshold | decimal | Optional; default is 5000 | Threshold at or above which you want to always pay for your Premium Domain order by wire transfer. See Usage section for more information. *Permitted values are in DD or DD.cc format* |
| ResponseType     | string | Optional         | Format of response. *Permitted values are Text \(default\), HTML, or XML.* |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Status | string | Success status of this query string |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | True indicates this entire response has reached you successfully. |

Example Output
--------------

The following query sets the threshold at which you require payment by wire to us to $5000.

```
https://resellertest.enom.com/interface.asp?command=NM_setPremiumDomainSettings&uid=resellid&pw=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&WireTransferThreshold=5000&responsetype={text, XML, or HTML}
```
```
<interface-response>
<Status>Successful</Status>
<Command>NM_SETPREMIUMDOMAINSETTINGS</Command>
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
<ExecTime>0.156</ExecTime>
<Done>true</Done>
<TrackingKey>42251f93-2563-4409-a2ad-4d41b03eae50</TrackingKey>
<RequestDateTime>7/3/2016 1:16:47 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL1VWRESELL_T
;Encoding Type is utf-8
Status=Successful
Command=NM_SETPREMIUMDOMAINSETTINGS
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
ExecTime=0.031
Done=true
TrackingKey=fa45f918-960c-4729-9d7d-3747b8f65c56
RequestDateTime=7/3/2016 1:22:13 PM
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>Status: </STRONG>Successful<BR /><STRONG>Command: </STRONG>NM_SETPREMIUMDOMAINSETTINGS<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t1<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.016<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>24ef245a-5718-4dc1-beae-92ab5123047f<BR /><STRONG>RequestDateTime: </STRONG>7/3/2016 1:22:23 PM<BR /></BODY></HTML>
```
Related Commands
----------------
- [NM\_GetPremiumDomainSettings](../docs/nm-getpremiumdomainsettings.md)
- [NM\_Search](../docs/nm-search.md)
- [NM\_ProcessOrder](../docs/nm-processorder.md)