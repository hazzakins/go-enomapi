SetRenew
========

Set the auto-renew flag for a domain name

Usage
-----

Use this command with RenewFlag set to 1 to renew the domain registration automatically. Use this command with RenewFlag set to 0 to require the owner to actively renew.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=SetRenew&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&renewflag=(Required)&responsetype=(Optional)
```
| Input Parameter  | Type | Status  | Description |
| ----------------- | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command      | string | Required | SetRenew |
| uid | string | Required | Your Account ID                                                                                                      |
| pw        | string | Required | Your API Token |
| RenewFlag | string | Required | Turn the auto-renew setting for the domain on or off. Options are 1 to auto-renew, 0 otherwise                                                             |
| AutoPakRenew   | string | Optional | Turn the auto-renew setting for all POP paks in this domain on or off. Permitted values are 1 to auto- renew, 0 otherwise. To control the auto-renew settings for individual POP paks, use the SetPakRenew command. |
| EmailForwardRenew | string | Optional | Turn the auto-renew setting for email forwarding in this domain on or off. Permitted values are 1 to auto-renew, 0 otherwise.                                             |
| URLForwardRenew  | string | Optional | Turn the URL forwarding setting for POP paks in this domain on or off. Permitted values are 1 to auto-renew, 0 otherwise. |
| WPPSRenew | string | Optional | Turn the auto-renew setting for Whois Privacy Protection Service in this domain on or off. Permitted values are 1 to auto-renew, 0 otherwise.                                     |
| ResponseType   | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

This command only returns the price for original purchase of domains, not transfers or renewals. For a complete list of prices, use "PE\_GetRetailPricing" or "PE\_GetResellerPrice".

| Output Parameter | Type  | Description |
| ---------------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | string | Name of command executed |
| RenewName | int   | Current auto-renew setting \(1 = on, 0 = off\) |
| AutoPakRenew | boolean | Current auto-renew setting for POP paks \(True = on, False = off\). This return parameter only appears if the AutoPakRenew is an input parameter. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolian | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=setrenew&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&renewflag=1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=setrenew&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&renewflag=1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=setrenew&uid=resellid&pw=resellpw&sld=resellerdocs
&tld=com&renewflag=1&responsetype=text
```
```
<interface-response>
 <RenewName>True</RenewName>
 <Success>True</Success>
 <Command>SETRENEW</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>0.266</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/11/2011 11:40:17 PM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>RenewName: </STRONG>True<br />
<STRONG>Success: </STRONG>True<br />
<STRONG>Command: </STRONG>SETRENEW<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+08.00<br />
<STRONG>ExecTime: </STRONG>0.203<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/5/2015 3:23:28 PM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
RenewName=True
Success=True
Command=SETRENEW
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.125
Done=true
RequestDateTime=2/5/2015 3:23:48 PM
```
Related Commands
----------------

Extend

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetExtendInfo

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

InsertNewOrder

SetPakRenew

SetPassword

SetRegLock

StatusDomain

UpdateExpiredDomains

ValidatePassword