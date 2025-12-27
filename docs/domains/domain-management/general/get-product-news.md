GetProductNews
==============

Getting product maintenance schedule details at a time.

Usage
-----

Use this command to get product maintenance alert messages.

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

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Product | Details of the product                                      |
| Message     | Maintenance message for the product |
| Total | Number of subscriptions                                     |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query is used to get maintenance alert for the products that have been requested in account resellid, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GETPRODUCTNEWS&uid=resellid
&pw=resellpw&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GETPRODUCTNEWS&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GETPRODUCTNEWS&uid=resellid
&pw=resellpw&responsetype=text
```
The response is as follows:

```
<interface-response>
<Alerts>
 <Total>0</Total>
</Alerts>
<Command>GETPRODUCTNEWS</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl21wresell01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.078</ExecTime>
<Done>true</Done>
<TrackingKey>b1bcb817-53ce-4c7f-ba6a-dd853d1f4de0</TrackingKey>
<RequestDateTime>11/14/2012 2:55:24 AM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Total: 0
count: 1
Command: GETPRODUCTNEWS
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.047
Done: true
TrackingKey: 53baf1ff-f357-4812-9def-aeaba99dbd57
RequestDateTime: 2/4/2015 11:23:22 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Total=0
count=1
Command=GETPRODUCTNEWS
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.031
Done=true
TrackingKey=3d1d7ccc-5eb4-4d11-a65e-ccd3a2ae88e1
RequestDateTime=2/4/2015 11:23:49 AM
```
Related Commands
----------------

GetNews