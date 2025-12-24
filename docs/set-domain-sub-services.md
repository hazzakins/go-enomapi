SetDomainSubServices
====================

Update settings for domain services \(active or inactive for email forwarding host records\).

Usage
-----

Use this command to enable or disable services. \(e.g., Host or email forwarding records\).

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
command=nameofcommand&uid=yourloginid&pw=yourpassword&
paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter   | Status | Description                                                 | Max Size |
| ------------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------- | -------- |
| UID         | Required | Account login ID                                              | 20 |
| PW | Required           | Account password | 20    |
| SLD         | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| ActivateForwarding | Optional; default is 1 | Enable email forwarding. Permitted values are: 1 to enable 0 to disable                  | 1 |
| ActivateHostRecords | Optional; default is 1    | Enable URL forwarding. Permitted values are: 1 to enable 0 to disable | 1    |
| ResponseType    | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                            | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ----------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| EmailForwarding | 1, 0 or empty                                          |
| HostRecords    | 1, 0 or empty |
| SetDomainServices | Successful status of this query                                 |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sets the subservices of a domain and sends the response in XML format:

```
https://resellertest.enom.com/interface.asp?
command=setdomainsubservices&uid=resellid&pw=resellpw&
sld=resellerdocs2&tld=net&ActivateForwarding=1&
ActivateHostRecords=1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=setdomainsubservices&uid=resellid&pw=resellpw&
sld=resellerdocs2&tld=net&ActivateForwarding=1&
ActivateHostRecords=1&responsetype=text
```
```
https://resellertest.enom.com/interface.asp?
command=renewservices&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&service=WPPS
&numyears=2&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <DomainServices>
 <EmailForwarding>1</EmailForwarding>
 <HostRecords>1</HostRecords>
 <SetDomainServices>Successful</SetDomainServices>
 </DomainServices>
 <Command>SETDOMAINSUBSERVICES</Command>
 <ErrCount>0</ErrCount>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
EmailForwardingEnabled: True
MailCount: 0
Command: SETDOMAINSUBSERVICES
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
ExecTime: 0.938
Done: true
RequestDateTime: 2/5/2015 2:52:36 PM
```