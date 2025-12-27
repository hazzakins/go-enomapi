GetHomeDomainList
=================

Retrieve a list of registered and hosted domains in this account that use our domain name servers.

Usage
-----

Use this command to retrieve a list of registered and hosted domains in this account that use our domain name servers.

A typical use of this command is to retrieve domains that are eligible to use our DNS-dependent services.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/websitecreator/wscsignup.asp](https://resellertest.enom.com/websitecreator/wscsignup.asp)

For accounts with 25 or fewer domain names that use our DNS, the Select a domain text box is a dropdown menu populated by the GetHomeDomainList command.

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
| Input Parameter | Status | Description                                                                   | Max Size |
| --------------- | ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Your Account ID                                                                 | 20 |
| PW | Required         | Account password | 20    |
| StartPosition  | Optional; default is 1 | Return domain names beginning with this number. For example, if StartPosition=26&Display=25, display names 26 through 50 in the ordered list. | 5 |
| Display | Optional; default is 25 | How many domain names to include in response. Permitted values are 1 to 100. | 100   |
| OrderBy     | Optional; default is SLD | Sort the response by this criterion. Permitted values are SLD, TLD, or RegStatus.                               | 9 |
| ResponseType | Optional         | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| ID        | Our numeric ID number for this domain |
| Name | Domain name                                           |
| RegStatus    | Registration status |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves a list of domains in this account that use our DNS servers, sorted by TLD, and sends the response in XML, HTML and text formats:

```
https://resellertest.enom.com/interface.asp?
command=GETHOMEDOMAINLIST&uid=resellid&pw=resellpw
&display=25&orderby=tld&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GETHOMEDOMAINLIST&uid=resellid&pw=resellpw
&display=25&orderby=tld&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GETHOMEDOMAINLIST&uid=resellid&pw=resellpw
&display=25&orderby=tld&responsetype=text
```
```
<interface-response>
 <GetHomeDomains>
 <StartPosition/>
 <Display>25</Display>
 <OrderBy>tld</OrderBy>
 <NewStartPosition>26</NewStartPosition>
 <DomainCount>25</DomainCount>
 <TotalDomains>3274</TotalDomains>
 <Domains>
  <Domain>
  <ID>152926976</ID>
  <Name>vras.ac</Name>
  <RegStatus>Registered</RegStatus>
  </Domain>
  <Domain>
  <ID>152926784</ID>
  <Name>fert.am</Name>
  <RegStatus>Registered</RegStatus>
  </Domain>
.
.
.
 </Domains>
 </GetHomeDomains>
 <Command>GETHOMEDOMAINLIST</Command>
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
 <ExecTime>1.018</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/8/2011 5:00:45 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
StartPosition:
Display: 25
OrderBy: tld
NewStartPosition: 26
DomainCount: 25
TotalDomains: 2898
ID1: 153001272
Name1: chrisbeckett321.am
RegStatus1: Registered
ID2: 345743445
Name2: ceatest2.asia
RegStatus2: Registered
ID3: 345756665
Name3: testgfd.asia
RegStatus3: Registered
ID4: 345388814
Name4: ardomaintest.biz
RegStatus4: Registered
ID5: 345564504
.
.
.
RegStatus25: Registered
Command: GETHOMEDOMAINLIST
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 1.422
Done: true
RequestDateTime: 2/4/2015 10:35:25 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
StartPosition=
Display=25
OrderBy=tld
NewStartPosition=26
DomainCount=25
TotalDomains=2898
ID1=153001272
Name1=chrisbeckett321.am
RegStatus1=Registered
ID2=345743445
Name2=ceatest2.asia
RegStatus2=Registered
ID3=345756665
Name3=testgfd.asia
RegStatus3=Registered
ID4=345388814
Name4=ardomaintest.biz
RegStatus4=Registered
.
.
.
Command=GETHOMEDOMAINLIST
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
ExecTime=0.141
Done=true
RequestDateTime=2/4/2015 10:36:15 AM
```
Related Commands
----------------

AddToCart

DisableServices

GetWPPSInfo

PurchaseServices

ServiceSelect

SetRenew

SetResellerServicesPricing

UpdateAccountPricing