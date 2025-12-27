CancelOrder
===========

Cancel a pre-order for a specific domain name.

Usage
-----

Use this command to cancel the order of a specific pre-registration domain name.

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
- The OrderID and Domain Name must be passed in and be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
http://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid
&pw=yourpassword&paramname=paramvalue
&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| DomainName   | Required | The Domain Name to cancel               | 32 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| OrderID     | The Order ID |
| Success | Boolean indicating if the action was successfull.                        |
| Result      | Text message indicating a success, or indicating an error message. |
| DomainName | The domain name                                         |
| Description   | Text message indicating the description as appropriate based on the action, and/or errors. |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The Domain Name must be a part of the order ID passed in, and both items must belong to the reseller calling the API command.

Example
-------

The following query cancels the domain abcd1234.catering from orderID 123456789, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=CancelOrder&uid=resellid&pw=resellpw
&OrderID=123456789&DomainName=abcd1234.catering
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=CancelOrder&uid=resellid&pw=resellpw
&OrderID=123456789&DomainName=abcd1234.catering
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=CancelOrder&uid=resellid&pw=resellpw
&OrderID=123456789&DomainName=abcd1234.catering
&responsetype=text
```
In the response, the presence of spun names and the ErrCount value 0 indicate that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
<Order>
 <OrderID>123456789</OrderID>
 <Success>True</Success>
 <Result>Success</Result>
 <Domains>
 <DomainName>abcd1234.catering</DomainName>
 <Description>Refunded</Description>
 </Domains>
</Order>
<Command>CANCELORDER</Command>
<APIType>API.NET</APIType>
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
<ExecTime>0.094</ExecTime>
<Done>true</Done>
<TrackingKey>dc60f425-1e6f-44de-86d8-e15390eea4eb</TrackingKey>
<RequestDateTime>7/2/2014 4:28:17 PM</RequestDateTime>
<debug>
 <![CDATA[ ]]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
OrderID: 123456789
Success: False
Result: Order does not exist or does not belong to customer
Command: CANCELORDER
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
ExecTime: 0.172
Done: true
TrackingKey: 9b15489f-bee1-476c-b24b-7876c43e4b61
RequestDateTime: 2/3/2015 1:26:48 PM
```
Related Commands
----------------

[AddToCart](../docs/addtocart.md)

[Check](../docs/domains/availability/check.md)

[Purchase](../docs/domains/registration/purchase.md)