GetExpiredDomains
=================

Retrieve a list of an account’s domains that are in Expired, Redemption Grace Period \(RGP\), and Extended RGP status.

Usage
-----

Use this command to retrieve a list of Expired, RGP, and Extended RGP names in the account specified in the query string. Typically, a reseller will use this information to determine which command to use to reactivate a domain, and what price they will charge.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The AuthQuestionAnswer value must be correct.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetExpiredDomains&uid=(Required)&PW=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetExpiredDomains |
| UID | string | Required | Your Account ID                              |
| PW       | string | Required | Login ID for the account for which you want the password. |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| DomainName    | string | Domain name                                           |
| DomainNameID   | string | Domain name ID number, from our database                             |
| Status      | string | Expiration status of this domain                                 |
| Expiration-Date | string | Expiration date of this domain                                  |
| LockStatus    | string | Registrar lock status of this domain                               |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getexpireddomains&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getexpireddomains&uid=resellid
&pw=resellpw&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=getexpireddomains&uid=resellid
&pw=resellpw&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <DomainDetail>
 <DomainName>northwestcrafts.com</DomainName>
 <DomainNameID>152139242</DomainNameID>
 <status>Extended RGP</status>
 <expiration-date>2/26/2003</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
 <DomainDetail>
 <DomainName>fabrics-wallpaper.com</DomainName>
 <DomainNameID>152134074</DomainNameID>
 <status>RGP</status>
 <expiration-date>3/1/2003 6:54:40 AM</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
 <DomainDetail>
 <DomainName>onlinebrochure.com</DomainName>
 <DomainNameID>152129772</DomainNameID>
 <status>RGP</status>
 <expiration-date>6/4/2002 6:18:36 PM</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
 <DomainDetail>
 <DomainName>eweathervane.com</DomainName>
 <DomainNameID>2062920</DomainNameID>
 <status>Expired</status>
 <expiration-date>3/8/2002 2:14:45 AM</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
 <DomainDetail>
 <DomainName>signonworldwide.com</DomainName>
 <DomainNameID>2063248</DomainNameID>
 <status>Expired</status>
 <expiration-date>3/8/2011 11:38:32 PM</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
 <DomainDetail>
 <DomainName>Youcanhelp.net</DomainName>
 <DomainNameID>152151738</DomainNameID>
 <status>Expired</status>
 <expiration-date>3/7/2011 6:44:29 PM</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
 <DomainDetail>
 <DomainName>ccpays.com</DomainName>
 <DomainNameID>2082140</DomainNameID>
 <status>Extended RGP</status>
 <expiration-date>2/26/2011 4:01:46 PM</expiration-date>
 <lockstatus>Locked</lockstatus>
 </DomainDetail>
.
.
.
 <domaincount>120</domaincount>
 <Command>GETEXPIREDDOMAINS</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod />
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable />
 <IsRealTimeTLD />
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>1.503906</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA [ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>DomainName1: </ STRONG>00000000001111.com<br>
<STRONG>DomainNameID1: </ STRONG>340691704<br>
<STRONG>status1: </ STRONG>Expired<br>
<STRONG>expiration-date1: </ STRONG>7/7/2013 8:37:00 AM<br>
<STRONG>lockstatus1: </ STRONG>Locked<br>
<STRONG>DomainName2: </ STRONG>00000startajay.name<br>
<STRONG>DomainNameID2: </ STRONG>152708848<br>
<STRONG>status2: </ STRONG>Expired<br>
<STRONG>expiration-date2: </ STRONG>2/4/2009 6:25:00 AM<br>
<STRONG>lockstatus2: </ STRONG>Not Locked<br>
<STRONG>DomainName3: </ STRONG>00000startajay.net<br>
<STRONG>DomainNameID3: </ STRONG>152708843<br>
<STRONG>status3: </ STRONG>Expired<br>
<STRONG>expiration-date3: </ STRONG>2/4/2009 6:25:00 AM<br>
<STRONG>lockstatus3: </ STRONG>Locked<br>
<STRONG>DomainName4: </ STRONG>00000startajay.tv<br>
<STRONG>DomainNameID4: </ STRONG>152708844<br>
<STRONG>status4: </ STRONG>Expired<br>
<STRONG>expiration-date4: </ STRONG>2/4/2009 6:25:00 AM<br>
<STRONG>lockstatus4: </ STRONG>Locked<br>
<STRONG>DomainName5: </ STRONG>0000ragavan.com<br>
<STRONG>DomainNameID5: </ STRONG>322157180<br>
<STRONG>status5: </ STRONG>Expired<br>
<STRONG>expiration-date5: </ STRONG>2/19/2009 7:06:26 AM<br>
<STRONG>lockstatus5: </ STRONG>Locked<br>
<STRONG>DomainName6: </ STRONG>000ajaycouk.com<br>
<STRONG>DomainNameID6: </ STRONG>152708718<br>
<STRONG>status6: </ STRONG>Expired<br>
<STRONG>expiration-date6: </ STRONG>1/30/2009 7:05:00 AM<br>
<STRONG>lockstatus6: </ STRONG>Locked<br>
.
.
.
<STRONG>domaincount: </ STRONG>7606<br>
<STRONG>Command: </ STRONG>GETEXPIREDDOMAINS<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.469<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>ef1c62d8-12ba-4dde-9482-dc5c0ce526ab<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 6:14:12 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainName1=00000000001111.com
DomainNameID1=340691704
status1=Expired
expiration-date1=7/7/2013 8:37:00 AM
lockstatus1=Locked
DomainName2=00000startajay.name
DomainNameID2=152708848
status2=Expired
expiration-date2=2/4/2009 6:25:00 AM
lockstatus2=Not Locked
DomainName3=00000startajay.net
DomainNameID3=152708843
status3=Expired
expiration-date3=2/4/2009 6:25:00 AM
lockstatus3=Locked
DomainName4=00000startajay.tv
DomainNameID4=152708844
status4=Expired
expiration-date4=2/4/2009 6:25:00 AM
lockstatus4=Locked
DomainName5=0000ragavan.com
DomainNameID5=322157180
status5=Expired
expiration-date5=2/19/2009 7:06:26 AM
lockstatus5=Locked
DomainName6=000ajaycouk.com
DomainNameID6=152708718
status6=Expired
expiration-date6=1/30/2009 7:05:00 AM
lockstatus6=Locked
.
.
.
domaincount=7606
Command=GETEXPIREDDOMAINS
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
ExecTime=0.266
Done=true
TrackingKey=e12a3134-f988-4513-9009-57b989345db6
RequestDateTime=2/3/2015 6:15:06 PM
```
Related Commands
----------------

AddToCart

Extend\_RGP

GetDomainInfo

GetDomains

PE\_SetPricing

Purchase