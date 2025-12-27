SynchAuthInfo
=============

Synchronize a domain’s EPP key with the Registry, and optionally email it to the registrant.

Usage
-----

Use this command to create a new EPP key, synchronize the key with the Registry, and optionally email it to the registrant.

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
https://resellertest.enom.com/interface.asp?command=
SynchAuthInfo&uid=(Required)&pw=(Required)&SLD=(Required)&TLD=(Required)&EmailEPP=(Required)&RunSynchAutoInfo=(Required)&responsetype=xml
```
| Input Parameter | Type | Status  | Description |
| ---------------- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required | TP\_GetOrderReview |
| uid | string | Required | Your Account ID                                                                   |
| pw        | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                            |
| TLD       | string | Required | Top-level domain name \(extension\) |
| EmailEPP | string | Required | Should we mail the EPP key to the registrant of this domain name. Permitted values are True and False                       |
| RunSynchAutoInfo | string | Required | Should we synchronize the EPP key for this domain and make sure our records are identical to the Registry’s. Permitted values are True and False. |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML.                                     |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                               |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------------- |
| Command     | string | Name of command executed                                        |
| InfoSynched   | boolean | Returns True if EPP key has been synchronized at Registry. This parameter displays when EmailEPP=False |
| EPPEmailMessage | string | Returns confirmation that EPP key is emailed to registrant. This parameter displays when EmailEPP=True |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.    |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.        |
| Done       | boolian | "True" indicates this entire response has reached you successfully.                  |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SynchAuthInfo&UID=resellid&pw=resellpw
&sld=resellerdocs&tld=com&EmailEPP=True
&RunSynchAutoInfo=True&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=SynchAuthInfo&UID=resellid&pw=resellpw
&sld=resellerdocs&tld=com&EmailEPP=True
&RunSynchAutoInfo=True&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SynchAuthInfo&UID=resellid&pw=resellpw
&sld=resellerdocs&tld=com&EmailEPP=True
&RunSynchAutoInfo=True&responsetype=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <EPPEmailMessage>Email has been sent.</EPPEmailMessage>
 <Command>SYNCHAUTHINFO</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLER1-STG</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>1.715</ExecTime>
 <Done>true</Done>
 <debug><![CDATA[ ]]></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>transferorderid: </ STRONG>175623473<br>
<STRONG>EPPEmailMessage: </ STRONG>Email has been sent.
<STRONG>Command: </ STRONG>SYNCHAUTHINFO
<STRONG>APIType: </ STRONG>API
<STRONG>Language: </ STRONG>eng
<STRONG>ErrCount: </ STRONG>0
<STRONG>ResponseCount: </ STRONG>0
<STRONG>MinPeriod: </ STRONG>1
<STRONG>MaxPeriod: </ STRONG>10
<STRONG>Server: </ STRONG>SJL0VWRESELL_T1
<STRONG>Site: </ STRONG>eNom
<STRONG>IsLockable: </ STRONG>True
<STRONG>IsRealTimeTLD: </ STRONG>True
<STRONG>TimeDifference: </ STRONG>+08.00
<STRONG>ExecTime: </ STRONG>1.406
<STRONG>Done: </ STRONG>true
<STRONG>RequestDateTime: </ STRONG>2/6/2015 10:59:41 AM
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
EPPEmailMessage=Email has been sent.
Command=SYNCHAUTHINFO
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
ExecTime=0.969
Done=true
RequestDateTime=2/6/2015 11:00:07 AM
```
Related Commands
----------------

TP\_CreateOrder