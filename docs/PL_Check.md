PL\Check
========

Check the status of a Premium Listed domain from a specific provider

Usage
-----

Use this command to check on the status of a specific domain at a given aftermarket provider and provide some details on the domain.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The domain must be listed for sale in the inventory of the provider

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=PL_Check&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Type | Description                      |
| --------------- | ------ | ------------------------------------------------------ |
| command     | string | GetNameSuggestions                   |
| uid       | string | Your Account ID                    |
| pw       | string | Your API Token                     |
| Sld       | string | Required - The SLD                  |
| Tld       | string | Required - The TLD                  |
| Provider    | string | Which provider to check. Currently only supports Sedo |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                      |
| ---------------- | -------- | ----------------------------------------------------- |
| domain      | string | Domain Name                      |
| price      | decimal | The price of this domain               |
| currency     | string | The currency of the price - defaults to USD     |
| status      | string | The status of the domain from the provider      |
| status\_set   | DateTime | Time and Date that the status was set         |
| listing\_type  | string | Coma separated list of Types assigned to this listing |
| provider     | string | The aftermarket provider               |

Example Output
--------------

```
;URL Interface
;Machine is SJL1VWRESELL_T1
;Encoding Type is utf-8
domain=testdomain-2212785.xyz
price=1150
currency=usd
status=for_sale
status_set=10/20/2014 4:23:25 AM
instant_transfer=true
listing_type=buy_now,instant_transfer
provider=sedo
Command=PL_CHECK
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl2vwapi01
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=1.734
Done=true
TrackingKey=db9ff16d-50ac-4944-8fe7-a5f82e126479
RequestDateTime=9/27/2016 2:52:05 PM
```
```html
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <item>
    <domain>
      <![CDATA[testdomain-2212785.xyz]]>
    </domain>
    <price>1150</price>
    <currency>usd</currency>
    <status>for_sale</status>
    <status_set>
      <![CDATA[10/20/2014 4:23:25 AM]]>
    </status_set>
    <instant_transfer>true</instant_transfer>
    <listing_type>
      <![CDATA[buy_now,instant_transfer]]>
    </listing_type>
   <provider>
      <![CDATA[sedo]]>
    </provider>
  </item>
  <Command>PL_CHECK</Command>
  <APIType>API.NET</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>SJL1VWRESELL_T1</Server>
  <Site>eNom</Site>
  <IsLockable/>
  <IsRealTimeTLD/>
  <TimeDifference>+0.00</TimeDifference>
  <ExecTime>1.750</ExecTime>
  <Done>true</Done>
  <TrackingKey>41676cea-ef8a-450a-b631-796202529d01</TrackingKey>
  <RequestDateTime>9/27/2016 2:51:12 PM</RequestDateTime>
  <debug/>
</interface-response>
```
Related commands
----------------

[PL\_Search](../docs/PL_Search.md)