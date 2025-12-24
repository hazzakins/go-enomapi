PE\GetPremiumPricing
====================

Retrieve the premium pricing for a list of domains.

Usage
-----

Use this command to retrieve the premium pricing for a list of domains.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=PE_GetPremiumPricing&uid=YourAccountID&pw=YourApiToken&sldX={Required}&tldX={Required}&qnameX={Required}&producttype={Required}&responsetype={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | --------------------------------------------------------------------------------------- |
| command     | string | Required | PE\_GetPremiumPricing |
| uid | string | Required | Your Account ID                                     |
| pw       | string | Required | Your API Token |
| sldX | string | Required | Second level domain. X=1 to Domain Count                        |
| tldX      | string | Required | Top level domain X=1 to DomainCount. |
| qnameX | string | Required | Queue name                                       |
| producttype   | string | Required | Product type. Permitted values are: - Register - Renew - RGP - Extended RGP - Transfer |
| responsetype | string | Optional | Format of response. Permitted values are "Text" \(default\), "HTML", or XML      |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                             |
| ---------------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------- |
| Domain      | string | Domain name queried                                                         |
| IsPremium    | string | Is premium domain? Expected values are: - Yes - No                                         |
| IsEAP      | string | Is EAP domain? Expected values are: - Yes - No                                           |
| ICANNFees    | float | ICANN fees.                                                            |
| Price      | float | Premium domain price. 0.00 value will be returned for a non-premium domain or if the domain is not eligible in a specified queue |
| RegistrationFee | float | Registration fee                                                          |
| EAPFee      | float | EAP fee                                                               |
| Command     | string | Name of command executed                                                      |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err \(1 to ErrCount\) values                   |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.                      |
| Done       | string | "True" indicates this entire response has reached you successfully                                 |

Example Output
--------------

```
https://resellertest.enom.com/interface.asp?command=PE_GetPremiumPricing&uid=resellid&pw=resellpw&responsetype=xml&sld1=bike48n25&tld1=bike&qname1=sunrise&sld2=bike48n01&tld2=ninja&qname2=sunrise&producttype=register
```
```
https://resellertest.enom.com/interface.asp?command=PE_GetPremiumPricing&uid=resellid&pw=resellpw&responsetype=text&sld1=bike48n25&tld1=bike&qname1=sunrise&sld2=bike48n01&tld2=ninja&qname2=sunrise&producttype=register
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
 <results>
 <result>
  <Domain>bike48n25.bike</Domain>
  <IsPremium>true</IsPremium>
  <IsEAP>false</IsEAP>
  <ICANNFees/>
  <Price>230.00</Price>
  <RegistrationFee>230.00</RegistrationFee>
  <EAPFee>0.00</EAPFee>
 </result>
 <result>
  <Domain>bike48n01.ninja</Domain>
  <IsPremium>false</IsPremium>
  <IsEAP>false</IsEAP>
  <ICANNFees>0.18</ICANNFees>
  <Price>0.00</Price>
  <RegistrationFee>0.00</RegistrationFee>
  <EAPFee>0.00</EAPFee>
 </result>
 </results>
 <Command>PE_GETPREMIUMPRICING</Command>
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
 <ExecTime>0.672</ExecTime>
 <Done>true</Done>
 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
 <RequestDateTime>11/25/2013 5:31:15 PM</RequestDateTime>
 <debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Domain1=bike48n25.bike
IsPremium1=false
IsEAP1=false
ICANNFees1=0.18
Price1=
RegistrationFee1=
EAPFee1=0.00
Domain2=bike48n01.ninja
IsPremium2=false
IsEAP2=false
ICANNFees2=0.18
Price2=
RegistrationFee2=
EAPFee2=0.00
Command=PE_GETPREMIUMPRICING
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
ExecTime=0.062
Done=true
TrackingKey=70cee5d4-1a94-4029-b5a6-7c46d815658b
RequestDateTime=2/4/2015 4:29:40 PM
```
Related Commands
----------------

GetAgreementPage

Queue\_DomainPurchase

Queue\_GetDomains

Queue\_GetExtAttributes

Queue\_GetInfo

Queue\_GetOrderDetail

Queue\_GetOrders