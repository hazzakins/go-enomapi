SubAccountDomains
=================

List the domains in a subaccount.

Usage
-----

Use this command to list the domains in a subaccount.

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
SubAccountDomains&uid=(Required)&pw=(Required)&Account=(Required)&Tab=(Required)&responsetype=xml
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------- |
| command     | string | Required | SubAccountDomains |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| Account | string | Required | Second-level domain name \(for example, enom in enom.com\)        |
| Tab       | string | Required | Top-level domain name \(extension\) |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ----------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command      | string | Name of command executed                                     |
| Tab        | string | Category of names returned                                    |
| MultiRRP     | string | Do names in this account include credentials of more than one registrar?             |
| Domain-List Type | string | Description of the category of names returned                          |
| DomainNameID   | int | Domain name ID number, from our records                             |
| SLD        | string | Second-level domain name \(for example, enom in enom.com\)                   |
| TLD        | string | Top-level domain name \(extension\)                               |
| NS-Status    | string | YES indicates that this domain uses our default name servers                   |
| Expiration-Date | string | Expiration date of this domain registration                           |
| Auto-Renew    | string | Is this domain set to renew automatically?                            |
| EndPosition    | string | Position in the overall list of the last domain returned with this query             |
| PreviousRecords  | string | Number of domains in the overall list before the first domain in this list            |
| NextRecords    | string | Number of domains in the overall list after the last domain in this list             |
| OrderBy      | string | Sorting parameter for names                                   |
| Result      | string | Did this query yield any names?                                 |
| StartPosition   | string | Position in the overall list of the first domain returned with this query            |
| DomainCount    | int | Total number of domains in this subaccount                            |
| TotalDomainCount | int | Total number of domains in this parent account                          |
| StartLetter    | string | First letter of the domains returned with this query                       |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SUBACCOUNTDOMAINS&uid=resellid&pw=resellpw
&Account=493-yp-5836&tab=iown&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SUBACCOUNTDOMAINS&uid=resellid&pw=resellpw
&Account=493-yp-5836&tab=iown&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SUBACCOUNTDOMAINS&uid=resellid&pw=resellpw
&Account=493-yp-5836&tab=iown&responsetype=text
```
```html
<?xml version="1.0" ?>
<interface-response>
 <GetDomains>
 <tab>iown</tab>
 <multirrp>False</multirrp>
 <domain-list type="Registered">
  <domain>
  <DomainNameID>152556719</DomainNameID>
  <sld>subresellerdocs</sld>
  <tld>com</tld>
  <ns-status>NA</ns-status>
  <expiration-date>8/2/2006</expiration-date>
  <auto-renew>Yes</auto-renew>
  </domain>
  <domain>
  <DomainNameID>152556721</DomainNameID>
  <sld>subresellerdocs2</sld>
  <tld>net</tld>
  <ns-status>NA</ns-status>
  <expiration-date>8/2/2006</expiration-date>
  <auto-renew>Yes</auto-renew>
  </domain>
  <domain>
  <DomainNameID>152556722</DomainNameID>
  <sld>subresellerdocs3</sld>
  <tld>info</tld>
  <ns-status>NA</ns-status>
  <expiration-date>8/2/2006</expiration-date>
  <auto-renew>Yes</auto-renew>
  </domain>
 </domain-list>
 <EndPosition>3</EndPosition>
 <PreviousRecords>0</PreviousRecords>
 <NextRecords>0</NextRecords>
 <OrderBy />
 <Result>True</Result>
 <StartPosition>1</StartPosition>
 <DomainCount>3</DomainCount>
 <TotalDomainCount>225</TotalDomainCount>
 <StartLetter />
 </GetDomains>
 <Command>SUBACCOUNTDOMAINS</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>e</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>0.2109375</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>tab: </ STRONG>iown<br>
<STRONG>multirrp: </ STRONG>False<br>
<STRONG>domain-list type=</ STRONG>"Registered"<br>
<STRONG>DomainNameID: </ STRONG>152556719<br>
<STRONG>sld: </ STRONG>subresellerdocs<br>
<STRONG>tld: </ STRONG>com<br>
<STRONG>ns-status: </ STRONG>NA<br>
<STRONG>expiration-date: </ STRONG>8/2/2006<br>
<STRONG>auto-renew: </ STRONG>Yes<br>
<STRONG>DomainNameID: </ STRONG>152556721<br>
<STRONG>sld: </ STRONG>subresellerdocs2<br>
<STRONG>tld: </ STRONG>net<br>
<STRONG>ns-status: </ STRONG>NA<br>
<STRONG>expiration-date: </ STRONG>8/2/2006<br>
<STRONG>auto-renew: </ STRONG>Yes<br>
<STRONG>DomainNameID: </ STRONG>152556722<br>
<STRONG>sld: </ STRONG>subresellerdocs3<br>
<STRONG>tld: </ STRONG>info<br>
<STRONG>ns-status: </ STRONG>NA<br>
<STRONG>expiration-date: </ STRONG>8/2/2006<br>
<STRONG>auto-renew: </ STRONG>Yes<br>
<STRONG>EndPosition: </ STRONG>3<br>
<STRONG>PreviousRecords: </ STRONG>0<br>
<STRONG>NextRecords: </ STRONG>0<br>
<STRONG>OrderBy:</ STRONG><br>
<STRONG>Result: </ STRONG>True<br>
<STRONG>StartPosition: </ STRONG>1<br>
<STRONG>DomainCount: </ STRONG>3<br>
<STRONG>TotalDomainCount: </ STRONG>225<br>
<STRONG>StartLetter:</ STRONG><br>
<STRONG>Command: </ STRONG>SUBACCOUNTDOMAINS<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T<br>
<STRONG>Site:</ STRONG><br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.172
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/6/2015 10:47:00 AM<br>
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
tab=iown
multirrp=False
type="Registered"
DomainNameID=152556719
sld=subresellerdocs
tld=com
ns-status=NA
expiration-date=8/2/2006
auto-renew=Yes
DomainNameID=152556721
sld=subresellerdocs2
tld=net
ns-status=NA
expiration-date=8/2/2006
auto-renew=Yes
DomainNameID=152556722
sld=subresellerdocs3
tld=info
ns-status=NA
expiration-date=8/2/2006
auto-renew=Yes
EndPosition=3
PreviousRecords=0
NextRecords=0
OrderBy=
Result=True
StartPosition=1
DomainCount=3
TotalDomainCount=225
StartLetter=
Command=SUBACCOUNTDOMAINS
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.172
Done=true
RequestDateTime=2/6/2015 10:47:00 AM
```
Related Commands
----------------

CreateAccount

CreateSubAccount

GetSubAccountDetails

GetSubAccountPassword

GetSubAccounts

GetSubaccountsDetailList