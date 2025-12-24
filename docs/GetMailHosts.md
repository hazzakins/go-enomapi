GetMailHosts
============

Get mail hosts \(DNS information\) only

Usage
-----

Use this command to retrieve the mail records associated with a domain name.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must use eNom’s domain name servers.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetMailHosts&uid=(Required)&pw=(Required)&sld=(Required)&tld=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | GetBalance |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| sld | string | Required | Second-level domain name \(for example, enom in enom.com\)       |
| tld       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                                                     |
| ---------------- | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command     | string | Name of command executed                                                                              |
| HostNameX    | string | Host name. For theHost name. For the GetMailHosts command, this is mail.sld.tld                                                 |
| RecordTypeX   | string | Record type. Permitted values are MXE \(email forwarding\) or MX \(POP3 mail\).                                                  |
| AddressX     | string | IP address of eNom’s mail forwarding service.                                                                   |
| MXPrefX     | string | Preference level: the lower this value, the higher priority this service. Default value is 10.                                          |
| ServiceSelect  | string | The identification number of the email service for this domain. 1048 no email 1051 email forwarding \("easy" mail\) 1054 user configured 1105 user simplified 1114 enable POP Mail |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                          |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.                                              |
| Done       | boolian | "True" indicates this entire response has reached you successfully.                                                        |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getmailhosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getmailhosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getmailhosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
```
<interface-response>
 <mailhosts>
 <EmailForwarding>False</EmailForwarding>
 <host>
  <hostID/>
  <HostName>@ (none)</HostName>
  <RecordType>MX</RecordType>
  <Address/>
  <MXPref>10</MXPref>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <RecordType>MX</RecordType>
  <Address/>
  <MXPref>10</MXPref>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <RecordType>MX</RecordType>
  <Address/>
  <MXPref>10</MXPref>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <RecordType>MX</RecordType>
  <Address/>
  <MXPref>10</MXPref>
 </host>
 <host>
  <hostID/>
  <HostName>* (other)</HostName>
  <RecordType>MX</RecordType>
  <Address/>
  <MXPref>10</MXPref>
 </host>
 <LastNumberFromRegHostCounter>10</LastNumberFromRegHostCounter>
 <MailHostCount>15</MailHostCount>
 </mailhosts>
 <ParkingEnabled>False</ParkingEnabled>
 <ServiceSelect/>
 <Command>GETMAILHOSTS</Command>
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
 <ExecTime>0.297</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 5:19:02 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>LastNumberFromRegHostCounter: </STRONG>0<br />
<STRONG>MailHostCount: </STRONG>0<br />
<STRONG>ParkingEnabled: </STRONG>False<br />
<STRONG>ServiceSelect: </STRONG>1048<br />
<STRONG>Command: </STRONG>GETMAILHOSTS<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>0.500<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>129ba2a1-15df-4ce7-a8d8-79eed0b1822c<br />
<STRONG>RequestDateTime: </STRONG>2/4/2015 10:57:22 AM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Reseller=1
Balance=-1,832.23
AvailableBalance=6,193.37
Price=8.95
COMPrice=8.95
NETPrice=9.00
ORGPrice=9.00
CCPrice=24.95
TVPrice=39.95
BZPrice=24.95
NUPrice=24.95
DomainCount=2692
Command=GETBALANCE
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=False
IsRealTimeTLD=True
TimeDifference=+03.00
ExecTime=0.141
Done=true
RequestDateTime=2/3/2015 4:41:51 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

Forwarding

GetForwarding

GetPOP3

ModifyPOP3

PurchasePOPBundle

SetUpPOP3User