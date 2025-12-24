EnableServices
==============

Switch on a service in a domain name account.

Usage
-----

Use this command to switch on a service in a domain name account.

This command does not give this account access to the service or subscribe to the service. Rather, it toggles on a service that the account has already subscribed to but has toggled off.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

\[[https://resellertest.enom.com/domains/ServiceSelection.asp](https://resellertest.enom.com/domains/ServiceSelection.asp)?

DomainNameID=152533676&service=wpps\]\([https://resellertest.enom.com/domains/ServiceSelection.asp](https://resellertest.enom.com/domains/ServiceSelection.asp)?

DomainNameID=152533676&service=wpps\)

When the protected option button is selected, the save changes button calls the EnableServices command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- Other constraints may apply depending on the services being enabled.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                    | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Your Account ID                                                                  | 20 |
| PW | Required | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                           | 63 |
| TLD | Required | Top-level domain name \(extension\) | 15    |
| Service     | Required | Name of the service to be enabled. Permitted values are: WPPS \(ID Protect, our Whois privacy protection service\) EmailForwarding URLForwarding | 15 |
| ResponseType | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Service     | Name of the service being enabled |
| ServiceStatus | New status of this service                                    |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query enables the ID Protect Whois privacy protection service on resellerdocs.com and sends the response in XML, HTML and text formats:

```
https://resellertest.enom.com/interface.asp?
command=ENABLESERVICES&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=wpps&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=ENABLESERVICES&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=wpps&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=ENABLESERVICES&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=wpps&responsetype=text
```
```
<interface-response>
 <Service>wpps</Service>
 <ServiceStatus>Enabled</ServiceStatus>
 <Command>ENABLESERVICES</Command>
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
 <ExecTime>0.156</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/7/2011 5:50:30 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Service: wpps
ServiceStatus: Enabled
Command: ENABLESERVICES
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
ExecTime: 0.094
Done: true
RequestDateTime: 2/3/2015 3:47:06 PM
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