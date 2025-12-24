GetDomainSubServices
====================

Get current settings for domain services \(active or inactive\).

Usage
-----

Use this command to get which services are enabled to control whether to present the user the ability to get or update information related to the service. \(eg. Host or email forwarding records\). For some services, you can get more detailed information using the GetDomainServices command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=GetDomainSubServices&uid=yourloginid
&pw=yourpassword&sld=domain&tld=com
```
| Input Parameter | Status | Description                                | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------- | -------- |
| UID       | Required | Your Account ID                              | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)        | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ----------------- | ------------------------------------------------------------------------------------------------ |
| EmailForwarding | True, False or empty                                       |
| HostRecords    | True, False or empty |
| SetDomainServices | Successful or Failed                                       |
| Command      | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sets the subservices of a domain and sends the response in XML, HTML and text formats:

```
https://resellertest.enom.com/interface.asp?
command=GetDomainSubServices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainSubServices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetDomainSubServices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
```
<interface-response>
 <DomainServices>
 <EmailForwarding>True</EmailForwarding>
 <HostRecords>True</HostRecords>
 </DomainServices>
 <Command>GETDOMAINSUBSERVICES</Command>
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
 <RequestDateTime>12/8/2011 4:37:14 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
EmailForwarding:
HostRecords:
Command: GETDOMAINSUBSERVICES
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.203
Done: true
RequestDateTime: 2/3/2015 6:10:10 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
EmailForwarding=
HostRecords=
Command=GETDOMAINSUBSERVICES
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.047
Done=true
RequestDateTime=2/3/2015 6:10:36 PM
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