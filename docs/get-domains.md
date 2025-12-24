GetDomains
==========

Get a single page of domain names with the ability to load the next or previous page of names.

Usage
-----

Use this command to list a single page of domains in an account. GetAllDomains is our older command for large accounts. GetAllDomains, a similar command, returns the complete list of domain names in an account but can time out for accounts with more than 200 domain names. AdvancedDomainSearch is the most up-to-date command for accounts of all sizes.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/](https://resellertest.enom.com/myaccount/)

In the my enom section, the x domain names link calls the GetDomains command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                                                                                                                             | Max Size |
| --------------- | -------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                                                                                                                           | 20 |
| PW | Required                      | Account password | 20    |
| Tab       | Optional; default is IOwn | The type of domains to return. Permitted values are: IOwn current names in this account Sub\_IOwn names in retail subaccounts WatchList names in this account’s watchlist IHost DNS hosting names, this account ExpiringNames names nearing expiration ExpiredDomains expired but able to renew RGP RGP and Extended RGP names Promotion names on promotional basis | 10 |
| DaysToExpired | Optional with Tab=ExpiringNames          | Return names that expire with in this number of days, whether they are set to auto-renew or not | 4    |
| RegStatus    | Optional with Tab=Sub\_IOwn; default is Registered | The type of domains to return for a subaccount. Permitted values are Registered and Expired.                                                                                                                                    | 10 |
| Display | Optional; default is 25              | Number of domains to return in one response. Permitted values are 0 to 100. | 4    |
| Start      | Optional; default is 1 | Return names that start with this number in the sorted list. For example, Display=25&Start=26 returns the 26th through 50th names from a numero-alphabetically sorted list.                                                                                            | 4 |
| OrderBy | Optional                      | The order to return the results. Permitted values are SLD, TLD, DNS, and ExpirationDate. | 15    |
| StartLetter   | Optional | Return names that start with this letter                                                                                                                                                               | 1 |
| MultiLang | Optional                      | If MultiLang=On, display SLD in native character set in UI. | 1    |
| Domain     | Optional | Return names that match this name. Use format SLD.TLD                                                                                                                                                       | 60 |
| ExtFormat | Optional                      | Returns XML tags without hyphens, to more rigorously adhere to XML standards and allow more trouble-free performance with automated parsers. Use ExtFormat=1 to return tags without hyphens. | 1    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                                                                                                                                        | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                      |
| ----------------- | --------------------------------------------------------------------------------------------------------------------- |
| DomainNameIDX | Domain name ID. Indexed X when ResponseType=Text or HTML.                              |
| SLDX       | Second-level domain name \(for example, enom in enom.com\). Indexed X when ResponseType=Text or HTML. |
| TLDX | Top-level domain name \(extension\). Indexed X when ResponseType=Text or HTML.                   |
| NS-StatusX    | Name server status. YES indicates this domain uses our name servers. Indexed X when ResponseType=Text or HTML. |
| Expiration-DateX | Expiration date of the domain registration. Indexed X when ResponseType=Text or HTML.                |
| Auto-RenewX   | Auto-renew setting. Return values are Yes or No. Indexed X when ResponseType= Text or HTML. |
| WPPSStatusX | WhoIs privacy protection setting. Return values are Enabled or Disabled. Indexed X when ResponseType=Text or HTML. |
| WPPSExpDate    | WhoIs privacy protection setting expiration date. |
| RRProcessor | RR processor. Indexed X when ResponseType=Text or HTML.                               |
| Command      | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.           |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                          |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests a list of domains and basic information about them for account resellid, and sends the response in XML, HTML, or Text format. This command returns 25 domains at a time. Because it does not specify which domain to start with, the response by default starts with the first domain \(sorting by domain name: numbers first, then letters\):

```
https://resellertest.enom.com/interface.asp?
command=GetDomains&uid=resellid&pw=resellpw
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomains&uid=resellid&pw=resellpw
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomains&uid=resellid&pw=resellpw
&ResponseType=text
```
The response lists the first 25 domains in account resellid, starting with resellerdocs.com and resellerdocs3.info. It also provides some summary information about the total contents of the account:

```html
<interface-response>
<GetDomains>
 <tab/>
 <multirrp>False</multirrp>
 <domain-list type="Registered">
 <domain>
  <DomainNameID>340691704</DomainNameID>
  <sld>00000000001111</sld>
  <tld>com</tld>
  <ns-status>YES</ns-status>
  <expiration-date>7/7/2012</expiration-date>
  <auto-renew>No</auto-renew>
  <wppsstatus>disabled</wppsstatus>
  <wppsexpdate>3/23/2010 12:21:15 PM</wppsexpdate>
  <RRProcessor>E</RRProcessor>
 </domain>
 <domain>
  <DomainNameID>152708845</DomainNameID>
  <sld>00000startajay</sld>
  <tld>info</tld>
  <ns-status>NA</ns-status>
  <expiration-date>2/4/2011</expiration-date>
  <auto-renew>No</auto-renew>
  <wppsstatus>disabled</wppsstatus>
  <wppsexpdate>3/23/2010 12:21:15 PM</wppsexpdate>
  <RRProcessor>E</RRProcessor>
 </domain>
.
.
.
 </domain-list>
 <EndPosition>25</EndPosition>
 <PreviousRecords>0</PreviousRecords>
 <NextRecords>26</NextRecords>
 <OrderBy/>
 <Result>True</Result>
 <StartPosition>1</StartPosition>
 <DomainCount>3084</DomainCount>
 <TotalDomainCount>3084</TotalDomainCount>
 <StartLetter/>
</GetDomains>
<Command>GETDOMAINS</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod/>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>1.828</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 4:27:04 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
tab:
multirrp: False
DomainNameID1: 345245402
sld1: 00000101010101fourth
tld1: com
ns-status1: NA
expiration-date1: 8/7/2015
auto-renew1: No
wppsstatus1: enabled
wppsexpdate1: 7/28/2015 4:46:59 AM
RRProcessor1: E
DomainNameID2: 345730105
sld2: 00000bbb
tld2: com
ns-status2: NA
expiration-date2: 8/13/2015
auto-renew2: No
wppsstatus2: enabled
wppsexpdate2: 10/2/2015 1:08:04 AM
RRProcessor2: E
DomainNameID3: 152708789
sld3: 000ajaycouk
tld3: us
ns-status3: NA
expiration-date3: 1/30/2012
auto-renew3: No
wppsstatus3: n/a
RRProcessor3: E
.
.
.
count: 25
EndPosition: 25
PreviousRecords: 0
NextRecords: 26
OrderBy:
StartPosition: 1
DomainCount: 2692
TotalDomainCount: 2692
StartLetter:
Command: GETDOMAINS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 1.359
Done: true
RequestDateTime: 2/3/2015 5:52:47 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
tab=
multirrp=False
DomainNameID1=345245402
sld1=00000101010101fourth
tld1=com
ns-status1=NA
expiration-date1=8/7/2015
auto-renew1=No
wppsstatus1=enabled
wppsexpdate1=7/28/2015 4:46:59 AM
RRProcessor1=E
DomainNameID2=345730105
sld2=00000bbb
tld2=com
ns-status2=NA
expiration-date2=8/13/2015
auto-renew2=No
wppsstatus2=enabled
wppsexpdate2=10/2/2015 1:08:04 AM
RRProcessor2=E
DomainNameID3=152708789
sld3=000ajaycouk
tld3=us
ns-status3=NA
expiration-date3=1/30/2012
auto-renew3=No
wppsstatus3=n/a
RRProcessor3=E
.
.
.
count=25
EndPosition=25
PreviousRecords=0
NextRecords=26
OrderBy=
StartPosition=1
DomainCount=2692
TotalDomainCount=2692
StartLetter=
Command=GETDOMAINS
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.188
Done=true
RequestDateTime=2/3/2015 5:53:54 PM
```
Related Commands
----------------

GetAllDomains

GetDomainCount

GetDomainExp

GetDomainInfo

GetExpiredDomains

GetExtendInfo

GetHomeDomainList

GetPasswordBit

GetRegistrationStatus

GetRegLock

GetRenew

GetSubAccountPassword

SetPassword

SetRegLock

SetRenew

StatusDomain

ValidatePassword