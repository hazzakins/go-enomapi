DeleteHostedDomain
==================

Discontinue our DNS hosting of a domain name.

Usage
-----

Use this command to cancel our DNS hosting of a domain name.

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
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                 | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                              | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                            | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| --------------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount       | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done         | True indicates this entire response has reached you successfully. |
| OldRegistrationStatus | Registration status before running this query                          |
| Status        | Success status for canceling DNS hosting on this domain |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query discontinues our DNS hosting of ExternalHostedDomain.com and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=DeleteHostedDomain&uid=resellid&pw=resellpw
&sld=ExternalHostedDomain&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteHostedDomain&uid=resellid&pw=resellpw
&sld=ExternalHostedDomain&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteHostedDomain&uid=resellid&pw=resellpw
&sld=ExternalHostedDomain&tld=com&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <DeleteHostedDomain>
 <OldRegistrationStatus>Hosted</OldRegistrationStatus>
 <Status>Success</Status>
 </DeleteHostedDomain>
 <Command>DELETEHOSTEDDOMAIN</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>Reseller3</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+03.00</TimeDifference>
 <ExecTime>0.15625</ExecTime>
 <Done>true</Done>
 <debug>
 <![CDATA [ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
OldRegistrationStatus: Hosted
Status: Success
Command: DELETEHOSTEDDOMAIN
Language: en
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: Reseller3
Site: enom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +03.00
ExecTime: 0.15625
Done: true
```
Related Commands
----------------

ExtendDomainDNS

PurchaseHosting