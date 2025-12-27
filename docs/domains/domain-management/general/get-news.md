GetNews
=======

Getting registry maintenance schedule details.

Usage
-----

Use this command to get registry maintenance alert messages.

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
| Message     | Registry message for the product |
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

The following query is used to get registry maintenance alerts that have been requested in account resellid, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetNews&uid=resellid&pw=resellpw&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=GetNews&uid=resellid&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetNews&uid=resellid&pw=resellpw&responsetype=text
```
The response is as follows:

```
<interface-response>
<Alerts>
 <Alert>
 <Product>
  <![CDATA[ cn ]]>
 </Product>
 <Message>
  <![CDATA[
CNNIC, the registry that manages .CN registrations has made a temporary policy
decision to stop all new .CN registrations effective Jan 06, 2010. Renewals will
continue to work as expected. Meanwhile all attempts to register new .CN domains
will fail or be returned as unavailable by eNom. No date has been set
for when the new registrations will be re-opened.
That is all the information we have at this time but we are working directly
with the registry to get this resolved as soon as possible.
]]>
 </Message>
 </Alert>
 <Total>1</Total>
</Alerts>
<Command>GETNEWS</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl21wresellt01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.016</ExecTime>
<Done>true</Done>
<TrackingKey>ceee4a3f-e1ea-45c0-83fd-62cb5faf38df</TrackingKey>
<RequestDateTime>11/14/2012 3:19:41 AM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Product: cn
Message: CNNIC, the registry that manages .CN registrations has made a temporary policy
decision to stop all new .CN registrations effective Jan 06, 2010. Renewals will
continue to work as expected. Meanwhile all attempts to register new .CN domains
will fail or be returned as unavailable by eNom. No date has been set
for when the new registrations will be re-opened.
That is all the information we have at this time but we are working directly
with the registry to get this resolved as soon as possible.
Total: 1
count: 1
Command: GETNEWS
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.031
Done: true
TrackingKey: cf85fda7-fa29-4207-921c-c444446734ef
RequestDateTime: 2/11/2015 2:54:47 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Product: cn
Message: CNNIC, the registry that manages .CN registrations has made a temporary policy
decision to stop all new .CN registrations effective Jan 06, 2010. Renewals will
continue to work as expected. Meanwhile all attempts to register new .CN domains
will fail or be returned as unavailable by eNom. No date has been set
for when the new registrations will be re-opened.
That is all the information we have at this time but we are working directly
with the registry to get this resolved as soon as possible.
Total=1
count=1
Command=GETNEWS
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
ExecTime=0.016
Done=true
TrackingKey=d541854f-a212-41f9-b4b3-db027d315e4a
RequestDateTime=2/11/2015 2:56:31 PM
```
Related Commands
----------------

GetProductNews