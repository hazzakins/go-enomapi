Queue\DomainPurchase
====================

Purchase a domain in a specified queue. Note, these orders will be processed once the queue closes.

Usage
-----

Use this command to purchase a domain in a specified queue. Note, these orders will be processed once the queue closes.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/tld-queue/pages/manage/purchase.aspx](https://resellertest.enom.com/tld-queue/pages/manage/purchase.aspx)

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                                                                                                                                                                                                                                                                                                                                                                 | Max Size |
| --------------- | -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                                                                                                                                                                                                                                                                                                                                                              | 20 |
| PW | Required                   | Account password | 20    |
| ItemList    | Required | Collection of domain orders, grouped in domain module\(s\). Each parameters for a domain are separated by pipe symbol or "\\|". Multiple domain entry is allowed by adding comma or "," between each domain module. Use the PE\_GetPremiumPricing command to retrieve the premium domain price for CustomerSuppliedPrice. Domain module format: QName= \[phase\] \\| SLD= \[sld\] \\| TLD= \[tld\] \\| ExtAttributes \\| Signed\_Mark\_Data \\| CustomerSuppliedPrice Example: ItemList=QName=Sunrise\\|SLD=IWantToBeA\\|TLD=ninja\\|ninja\_name=Super\+Samurai\\|ninja\_number=007\\|signed\_mark\_data=Shogun, QName=Sunrise\\|SLD=Training\\|TLD=ninja\\|ninja\_name=Jim\+Smith\\|signed\_mark\_data=DkOnNp...CustomerSuppliedPrice=65.00 Note: please use URL Encoding for ListItem or the entire API string to avoid any unexpected results. | ?    |
| UseCreditCard  | Optional | If UseCreditCard=Yes, use our credit card processing services. This service is available only to resellers who have entered into a credit card processing agreement with us. When you pass credit card information with this command, you must use the secure HTTPS protocol.                                                                                                                                                                                                                                                                              | 3 |
| Amount | Required if UseCreditCard=Yes        | Amount to charge for this order. Format: DDD.CC | 8    |
| EndUserIP    | Required if using our credit card processing | End user’s IP address. This is used in fraud checking, as part of our order processing service. Format: NNN.NNN.NNN.NNN.                                                                                                                                                                                                                                                                                                                                                       | 15 |
| CCName | Required if using our credit card processing | Cardholder's name | 60    |
| CCExpMonth   | Required if using our credit card processing | Credit card expiration month. Format: MM                                                                                                                                                                                                                                                                                                                                                                                                  | 2 |
| CCNumber | Required if using our credit card processing | Customer's credit card number | 128   |
| CCExpYear    | Required if using our credit card processing | Credit card expiration year. Format: YYYY                                                                                                                                                                                                                                                                                                                                                                                                 | 4 |
| CVV2 | Required if using our credit card processing | Credit card verification code | 4    |
| CCAddress    | Required if using our credit card processing | Credit card billing address                                                                                                                                                                                                                                                                                                                                                                                                         | 60 |
| CCCity | Required if using our credit card processing | Credit card billing city | 60    |
| CCStateProvince | Required if using our credit card processing | Credit card billing state or province                                                                                                                                                                                                                                                                                                                                                                                                    | 60 |
| CCZip | Required if using our credit card processing | Credit card billing postal code | 15    |
| CCCountry    | Required if using our credit card processing | Credit card billing country \(2 letter code\) Format: XX                                                                                                                                                                                                                                                                                                                                                                                          | 60 |
| ResponseType | Optional Format of response.        | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter   | Description |
| -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| Success       | Overall item status. If there is one failure within a transaction \(multiple domains\), the entire transaction fails. |
| OrderID | Order ID for this transaction. All domains are purchased in a single order.                               |
| Amount        | Total amount for this transaction. Please notice that the amount might be modified if the price does not meet the account's minimum cost |
| TransactionNumber | Credit card transaction number                                                      |
| TransactionReference | Credit card transaction reference |
| Command | Name of command executed                                                         |
| ErrCount       | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                         |
| Done         | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query purchases a .NINJA domain name from the corresponding sunrise queue, and sends the response in XML, HTML, or Text format.

**NON-PREMIUM** - without Signed Mark Data

```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=resellpw
&Command=Queue_DomainPurchase&ItemList=qname
%3dsunrise%7csld%3dmydomain%7ctld%3dninja
%7csigned_mark_data%3d&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=resellpw
&Command=Queue_DomainPurchase&ItemList=qname
%3dsunrise%7csld%3dmydomain%7ctld%3dninja
%7csigned_mark_data%3d&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=resellpw
&Command=Queue_DomainPurchase&ItemList=qname
%3dsunrise%7csld%3dmydomain%7ctld%3dninja
%7csigned_mark_data%3d&ResponseType=text
```
**PREMIUM** - with Signed Mark Data and CustomerSuppliedPrice samples

```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=resellpw
&Command=Queue_DomainPurchase&ItemList=qname%3dsunrise
%7csld%3dbike48N12%7ctld%3dninja%7csigned_mark_data
%3dDQo8......TaW1cmU%2bPC9zbWQ6c2dEFmdGVyPg0KIClnTWFyaz4
%3d%7cCustomerSuppliedPrice%3d365.00&ResponseType=xml
```
```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=resellpw
&Command=Queue_DomainPurchase&ItemList=qname%3dsunrise
%7csld%3dbike48N12%7ctld%3dninja%7csigned_mark_data
%3dDQo8......TaW1cmU%2bPC9zbWQ6c2dEFmdGVyPg0KIClnTWFyaz4
%3d%7cCustomerSuppliedPrice%3d365.00&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
UID=resellid&PW=resellpw
&Command=Queue_DomainPurchase&ItemList=qname%3dsunrise
%7csld%3dbike48N12%7ctld%3dninja%7csigned_mark_data
%3dDQo8......TaW1cmU%2bPC9zbWQ6c2dEFmdGVyPg0KIClnTWFyaz4
%3d%7cCustomerSuppliedPrice%3d365.00&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<Queue_DomainPurchase>
 <Success>true</Success>
 <OrderID>157938961</OrderID>
 <Amount>30.00</Amount>
 <TransactionNumber>819799639717151206</TransactionNumber>
 <TransactionReference>otbfEzxtoEeWgzdhq11L/Q==</TransactionReference>
</Queue_DomainPurchase>
<Command>QUEUE_DOMAINPURCHASE</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>2.250</ExecTime>
<Done>true</Done>
<TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
<RequestDateTime>6/5/2013 1:49:08 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success: true
OrderID: 161808479
TransactionNumber:
TransactionReference:
Command: QUEUE_DOMAINPURCHASE
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
ExecTime: 22.735
Done: true
TrackingKey: 5e351eb0-22f3-488a-a8fb-e6c7aa5a292a
RequestDateTime: 2/5/2015 11:15:37 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
Success=true
OrderID=161808481
TransactionNumber=
TransactionReference=
Command=QUEUE_DOMAINPURCHASE
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
ExecTime=22.094
Done=true
TrackingKey=e8d40776-f4fa-4e7e-98ba-908ae602a4c2
RequestDateTime=2/5/2015 11:16:35 AM
```
Related Commands
----------------

GetAgreementPage

PE\_GetPremiumPricing

Queue\_GetDomains

Queue\_GetExtAttributes

Queue\_GetInfo

Queue\_GetOrderDetail

Queue\_GetOrders