UpdateExpiredDomains
====================

Reactivate an expired domain in real time.

Usage
-----

Use this command to reactivate a domain after it has expired, while it is in "Expired" status. This command reactivates a domain in real time.

> ### For domains that are NOT expired.
>
>
>
> Before a domain expires, use the "Extend" command to renew it. We recommend renewing at least a week before a domain expires.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The domain must be in Expired status. The registrar sets the duration of the grace period between expiration and deletion from the registrar’s database. To retrieve a list of expired domains, use the "GetDomains" command with parameter "Tab=ExpiredDomains".
- Customer must have sufficient funds. The charge for reactivation is the same as a one-year renewal.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?Command=UpdateExpiredDomains&uid=Your Account ID&pw=Your API Token&DomainName={Required}&NumYear={Required}&responsetpe={Optional}
```
| Input Parameter    | Type | Status                       | Description |
| --------------------- | ------ | -------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command        | string | Required                      | UpdateExpiredDomains |
| uid | string | Required | Your Account ID                                                                                                                          |
| pw          | string | Required                      | Your API Token |
| DomainName | string | Required | Domain name \(for example, example.com\)                                                                                                             |
| NumYears       | int | Required                      | Number of years to add to the domain registration. Limit is ten years from the day that this command is executed. |
| CustomerSuppliedPrice | float | Required if the "DomainName" is a Premium Domain. | Use the "Check" command to determine if this domain is a Premium domain. If the flag "IsPremiumName=true", then price listed in the "RenewalPrice=" flag, from the same "Check" command, will need to be included here under the "CustomerSuppliedPrice=" flag. |
| ResponseType     | string | Optional                      | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| Status | string | "Success" status of this query. "True" indicates that the update was successful. |
| OrderID | string | Order identification number |
| Command | string | UpdateExpiredDomains |
| ErrCount | int  | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | string | "True" indicates this entire response has reached you successfully. |

Example Output
--------------

The following query requests the reactivation of expired domain "resellerdocs.com" and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?Command=UpdateExpiredDomains&uid=Your Account ID&pw=Your API Token&DomainName=resellerdocs.com&responsetpe={Optional}
```
```
<?xml version="1.0" ?>
<interface-response>
<ReactivateDomainName>
<Status>True</Status>
<OrderID>157614452</OrderID>
</ReactivateDomainName>
<Command>UPDATEEXPIREDDOMAINS</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+08.00</TimeDifference>
<ExecTime>0.938</ExecTime>
<Done>true</Done>
<debug>
<![CDATA[ ]]>
</debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>Status: </STRONG>True<br />
<STRONG>OrderID: </STRONG>157614452<br />
<STRONG>Command: </STRONG>UPDATEEXPIREDDOMAINS<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG> <br />
<STRONG>IsRealTimeTLD: </STRONG> <br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.000<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>0e7e519a-3f3d-4e4c-8584-94d30dcdbb0f<br />
<STRONG>RequestDateTime: </STRONG>2/9/2015 2:06:52 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Status=True
OrderID=157614452
Command=UPDATEEXPIREDDOMAINS
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
ExecTime=0.000
Done=true
TrackingKey=4c9f196e-3eb8-479e-aae7-006827a9fbed
RequestDateTime=2/9/2015 2:08:40 PM
```
Related Commands
----------------

Extend

Extend\_RGP

GetDomainExp

GetExtendInfo

[GetRenew](../docs/domains/renewal/getrenew.md)

InsertNewOrder

SetRenew