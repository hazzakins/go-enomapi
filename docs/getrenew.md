GetRenew
========

Get the auto-renew setting for a domain name.

Usage
-----

Use this command to determine whether a domain in your account is set to renew automatically. If so, the renewal is charged to your account a month before the expiration date for the domain name. Fifteen days before that, you receive an email notifying you of the auto-renew charges.

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

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=GetRenew&uid=Your Account ID&pw=Your API Token&sld={Required}&tld={Required}&ResponseType={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | GetRenew |
| uid | string | Required | Your Account ID                                                            |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                      |
| TLD       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ----------------- | ------ | ------------------------------------------------------------------------------------------------ |
| RenewName     | int | Returns "1" if this domain is set to auto-renew, "0" otherwise.                |
| PakExist Returns | int | "1" if there are POP paks associated with this domain, "0" otherwise.              |
| AutoPakRenew   | int | Returns "1" if the POP paks for this domain are set to auto-renew, "0" otherwise.       |
| EMailFwdExists  | int | Returns "1" if this domain subscribes to email forwarding, "0" otherwise.            |
| EMailForwardRenew | int | Returns "1" if the email forwarding for this domain is set to auto-renew, "0" otherwise.    |
| URLFwdExists   | int | Returns "1" if this domain subscribes to URL forwarding, "0" otherwise.             |
| URLForwardRenew  | int | Returns "1" if the URL forwarding for this domain is set to auto-renew, "0" otherwise.     |
| Command      | string | GetRenew                                             |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | string | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the auto-renew status of "resellerdocs.com" and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=GetRenew&uid=resellid&pw=resellpw&sld=resellerdocs&tld=com&ResponseType={Optional}
```
```html
<interface-response>
<auto-renew>1</auto-renew>
<AutoPakRenew>0</AutoPakRenew>
<PakExist>1</PakExist>
<EmailForwardRenew>0</EmailForwardRenew>
<EmailFwdExists>0</EmailFwdExists>
<URLForwardRenew>0</URLForwardRenew>
<URLFwdExists>0</URLFwdExists>
<IDProtectRenew>0</IDProtectRenew>
<IDProtectExists>1</IDProtectExists>
<MobilizerRenew>0</MobilizerRenew>
<Command>GETRENEW</Command>
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
<ExecTime>0.125</ExecTime>
<Done>true</Done>
<RequestDateTime>12/9/2011 2:29:36 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>RenewName= </STRONG>1<br />
<STRONG>AutoPakRenew= </STRONG>0<br />
<STRONG>PakExist= </STRONG>1<br />
<STRONG>EmailForwardRenew= </STRONG>0<br />
<STRONG>EmailFwdExists= </STRONG>0<br />
<STRONG>URLForwardRenew= </STRONG>0<br />
<STRONG>URLFwdExists= </STRONG>0<br />
<STRONG>IDProtectRenew= </STRONG>0<br />
<STRONG>IDProtectExists= </STRONG>1<br />
<STRONG>MobilizerRenew= </STRONG>0<br />
<STRONG>Command= </STRONG>GETRENEW<br />
<STRONG>APIType= </STRONG>API.NET<br />
<STRONG>Language= </STRONG>eng<br />
<STRONG>ErrCount= </STRONG>0<br />
<STRONG>ResponseCount= </STRONG>0<br />
<STRONG>MinPeriod= </STRONG>1<br />
<STRONG>MaxPeriod= </STRONG>10<br />
<STRONG>Server= </STRONG>sjl0vwresell_t1<br />
<STRONG>Site= </STRONG>eNom<br />
<STRONG>IsLockable= </STRONG>True<br />
<STRONG>IsRealTimeTLD= </STRONG>True<br />
<STRONG>TimeDifference= </STRONG>+8.00<br />
<STRONG>ExecTime= </STRONG>0.047<br />
<STRONG>Done= </STRONG>true<br />
<STRONG>TrackingKey= </STRONG>de2469c0-ff52-4e54-9301-ebdd1db50fb8<br />
<STRONG>RequestDateTime= </STRONG>2/4/2015 12:21:47 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
RenewName=1
AutoPakRenew=0
PakExist=1
EmailForwardRenew=0
EmailFwdExists=0
URLForwardRenew=0
URLFwdExists=0
IDProtectRenew=0
IDProtectExists=1
MobilizerRenew=0
Command=GETRENEW
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.047
Done=true
TrackingKey=de2469c0-ff52-4e54-9301-ebdd1db50fb8
RequestDateTime=2/4/2015 12:21:47 PM
```
Related Commands
----------------

[Extend](../docs/extend.md)

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetDomains

GetExtendInfo

GetPasswordBit

GetPOPExpirations

[GetRegLock](../docs/getreglock.md)

GetSubAccountPassword

InsertNewOrder

RenewPOPBundle

SetPakRenew

SetPassword

SetRegLock

SetRenew

StatusDomain

UpdateExpiredDomains

ValidatePassword