NM\GetPremiumDomainSettings
===========================

Retrieve account-level settings associated with Premium Domains.

Usage
-----

Use this command to retrieve account-level settings that are specific to Premium Domains:

The minimum and maximum price of Premium Domains to be returned by Instant Reseller.

The threshold price above which you want to pay by wire transfer instead of using your account balance.

Availability
------------

Premium Domains can only be sold by our direct ETP resellers.

Constraints
-----------

The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------- |
| command     | string | Required | NM\_GetPremiumDomainSettings |
| uid | string | Required | Your API Token                                |
| pw       | string | Required | Your Account Password |
| ResponseType | string | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.* |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter   | Type | Description                                                                                           |
| --------------------- | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| MinimumPrice     | decimal | Minimum price of Premium Domains to be returned in searches in your Instant Reseller site.                                                   |
| MaximumPrice     | decimal | Maximum price of Premium Domains to be returned in searches your Instant Reseller site.                                                    |
| WireTransferThreshold | decimal | $US amount above which you want to pay by wire transfer instead of account balance or credit card. See Usage notes in [NM\_SetPremiumDomainSettings](../docs/nm-setpremiumdomainsettings.md) |
| Command        | string | Name of command executed                                                                                    |
| ErrCount       | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                |
| ErrX         | string | Error messages explaining the failure. These can be presented as is back to the client.                                                    |
| Done         | boolean | True indicates this entire response has reached you successfully.                                                               |

Example Output
--------------

The following query retrieves account-level premium domain settings.

```
https://resellertest.enom.com/interface.asp?
command=NM_GetPremiumDomainSettings&uid=resellid
&pw=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&responsetype={text, XML, or HTML}
```
```
<interface-response>
<PremiumDomainSettings>
<MinimumPrice>589.00</MinimumPrice>
<MaximumPrice>5000.00</MaximumPrice>
<WireTransferThreshold>4000.00</WireTransferThreshold>
</PremiumDomainSettings>
<Command>NM_GETPREMIUMDOMAINSETTINGS</Command>
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
<TrackingKey>275010ae-0b18-4e77-93a7-dc6692a16ae5</TrackingKey>
<RequestDateTime>7/3/2016 1:07:04 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL1VWRESELL_T
;Encoding Type is utf-8
MinimumPrice=589.00
MaximumPrice=5000.00
WireTransferThreshold=4000.00
Command=NM_GETPREMIUMDOMAINSETTINGS
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
TrackingKey=1b4a21f4-2cec-483e-bad0-7e633ab30dc3
RequestDateTime=7/3/2016 1:08:04 PM
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>MinimumPrice: </STRONG>589.00<BR /><STRONG>MaximumPrice: </STRONG>5000.00<BR /><STRONG>WireTransferThreshold: </STRONG>4000.00<BR /><STRONG>Command: </STRONG>NM_GETPREMIUMDOMAINSETTINGS<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t1<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.000<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>c2a961fa-cf6d-445b-b1f4-6364e5b7fcb7<BR /><STRONG>RequestDateTime: </STRONG>7/3/2016 1:08:16 PM<BR /></BODY></HTML>
```
Related Commands
----------------
- [NM\_Search](../docs/nm-search.md)
- [NM\_ProcessOrder](../docs/nm-processorder.md)
- [NM\_SetPremiumDomainSettings](../docs/nm-setpremiumdomainsettings.md)