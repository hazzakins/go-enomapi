ServiceSelect
=============

Enable and disable services for a domain. You can also use this command to retrieve the current settings for a service.

Usage
-----

Use this command to enable or disable email \(MX or MXE\), DNS, ID Protect services for a domain.

Use this command to get the current settings for a service.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- One query can select only one service. You must repeat the query for each service you want to select.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=serviceselect&uid=YourAccountID&pw=YourApiToken&sld={Required}&tld={Required}&newoptionid={Optional}&service={Required}&update={Required}&responsetype={Optional}
```
| Input Parameter | Type | Status                                                           | Description |
| --------------- | ------ | -------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required                                                          | ServiceSelect |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                                                                                                                       |
| pw       | string | Required                                                          | Your API Token |
| Service | string | Required; if "Service" is used without "NewOptionID" and "Update", the response gives the current setting "Service" type. | Permitted values are: - EmailSet \(email services\) - WPPS \(ID Protect Whois Privacy Protection\)                                                                                                                                                                                                                                                                             |
| NewOptionID   | int | Optional; if you use "NewOptionID" you must also use "Update" flag for the service.                    | Permitted values are: - **Service=DNSServer** \(which domain name servers\): - \*1006\*\* - use our name servers - \*1012 \*\*- use user-specified domain servers, including none - **Service=EmailSet** \(Email services\): - \*1048 \*\*- no email - \*1051 \*\*- email forwarding \(to a POP or WebMail address\) - \*1054 \*\*- user \(mail server name required\) - \*1105 \*\*- user simplified \(mail server’s IP address required\) - \*1114\*\* - POP3/WebMail plus email forwarding - **Service=WPPS** \(ID Protect Whois Privacy Protection\): - \*1120 \*\*- WhoIs information is masked - \*1123\*\* - WhoIs information is viewable |
| Update | sting | Required if you use "NewOptionID" | "True" updates the "NewOptionID" value for the service                                                                                                                                                                                                                                                                                                   |
| SLD       | string | Required for Service=EmailSet and for Service=WPPS                                     | Second-level domain name \(for example, enom in enom.com\) |
| TLD | string | Required for Service=EmailSet and for Service=WPPS | Top-level domain name \(extension\)                                                                                                                                                                                                                                                                                                            |
| HostAccount   | string | Required if Service=WSB and NewOptionID=1066                                        | Name of the Web hosting account to enable or disable |
| ResponseType | string | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML.                                                                                                                                                                                                                                                                                        |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| current-svc   | string | Current service option setting                                  |
| command     | string | ServiceSelect                                          |
| errcount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| errX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| done       | string | True value indicates this entire response has reached you successfully.             |

Example Output
--------------

The following query requests the contents of the shopping cart for account resellid and requests the response in format of either "text" \(default\), "xml" or "html"

The response indicates that email services have changed from option 1048 \(no email\) to 1114 \(POP3/WebMail plus email forwarding\):

```
https://resellertest.enom.com/interface.asp?command=serviceselect&uid=YourAccountID&pw=YourApiToken&sld=resellerdocs&tld=com&newoptionid=1114&service=emailset&update=True&responsetype={Optional}
```
```
<interface-response>
<ParkingEnabled>False</ParkingEnabled>
<ServiceSelect>
<domainname sld="resellerdocs" tld="com" domainnameid="152533676" resellerdocs.com</domainname>
<service name="emailset">
<option name="1048"/>
<option name="1051"/>
<option name="1054"/>
<option name="1105"/>
<option name="1114"/>
</service>
<current-svc>1114</current-svc>
</ServiceSelect>
<Command>SERVICESELECT</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl21wresellt01</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>0.500</ExecTime>
<Done>true</Done>
<debug/>
<TrackingKey>a0987545-ce87-4455-9965-559cbaef1d1e</TrackingKey>
<RequestDateTime>12/11/2011 10:37:40 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>domainname: </STRONG>resellerdocs.com<br />
<STRONG>option: </STRONG><br />
<STRONG>option: </STRONG><br />
<STRONG>option: </STRONG><br />
<STRONG>option: </STRONG><br />
<STRONG>option: </STRONG><br />
<STRONG>current-svc: </STRONG>1114<br />
<STRONG>Command: </STRONG>SERVICESELECT<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>0.250<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>d0e23231-fa8e-490a-957d-b80c986526f4<br />
<STRONG>RequestDateTime: </STRONG>2/5/2015 2:35:30 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
domainname=resellerdocs.com
option=
option=
option=
option=
option=
current-svc=1114
Command=SERVICESELECT
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.047
Done=true
TrackingKey=72d0a9c3-46b1-4bef-9f94-3825d9567b41
RequestDateTime=2/5/2015 2:36:54 PM
```
Related Commands
----------------

GetIPResolver

GetWPPSInfo

ModifyNS

PE\_SetPricing

SetHosts

SetIPResolver

SetResellerServicesPricing