PL\Search
=========

Check the status of a Premium Listed domain from a specific provider

Usage
-----

Use this command to search for a phrase which returns relevant Buy Now domains available in the aftermarket inventory

Availability
------------

All resellers have access to this command.

Constraints
-----------

The only supported aftermarket provider at this time for this command is SedoMLS

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=PL_Search&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Type | Description                                                                           |
| --------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | GetNameSuggestions                                                                       |
| uid       | string | Your Account ID                                                                         |
| pw       | string | Your API Token                                                                         |
| Keyword     | string | - \*Required\*\*: The search phrase, typically as entered by the user. For example "newyorkhouse" or "new york house" or "newyorkhouse.com".         |
| Tlds      | string | Request search only within a list of comma-delimited TLDs. Example: "com, net, org". If not provided, the search will take place across all supported TLDs. |
| Currency    | string | Specify what currency to use for the domains \(usd, gbp or eur\). If not provided, domains will be returned in US dollars.                  |
| Hyphens     | bool | Set to "false" to remove all domains in the search result. Defaults to "true".                                        |
| Offensive    | bool | Set to "true" to filter out domains with offensive or adult words. Default is "false" \(don't filter anything\)                        |
| OrderBy     | string | Set it to "price" to sort by price. Otherwise sorting is by relevance. \( Rank \)                                       |
| MaxPrice    | decimal | The maximum price for the domain.                                                               |
| MinPrice    | decimal | The minimum price for the domain                                                                |
| NumRecs     | int | Maximum number of domains to return. Defaults to 20.                                                     |
| ListingType   | string | Defines the domain types shown in your search results. Possible values: buy\_now, buy\_now\_instant, make\_offer Defaults to "buy\_now"            |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                |
| ---------------- | ------- | -------------------------------------------------------------------------- |
| domain      | string | Domain Name                                |
| price      | decimal | The price of this domain                          |
| rank       | int | Relevance ranking                             |
| listing\_type  | string | Defines the domain types shown in your search results; coma separated list |
| totalcount    | int | Number of results returned                         |
| currency     | string | The currency of the price - defaults to USD                |

Example Output
--------------

```
;URL Interface
;Machine is SJL1VWRESELL_T1
;Encoding Type is utf-8
domain1=testdomain-2212785.xyz
price1=1150
rank1=1
listing_type1=buy_now,instant_transfer
domain2=testdomain2-2212785.club
price2=1150
rank2=2
listing_type2=buy_now,instant_transfer
domain3=testdomain2-2212785.de
price3=1150
rank3=3
listing_type3=buy_now,instant_transfer
domain4=testdomain2-2212785.net
price4=1150
rank4=4
listing_type4=buy_now,instant_transfer
domain5=testdomain2-2212785.xyz
price5=1150
rank5=5
listing_type5=buy_now,instant_transfer
currency=usd
totalcount=5
Command=
APIType=API.NET
Language=
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=
Server=
Site=
IsLockable=
IsRealTimeTLD=
TimeDifference=
ExecTime=
Done=
TrackingKey=
RequestDateTime=
```
```html
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <results>
    <item>
      <domain>
        <![CDATA[testdomain-2212785.xyz]]>
      </domain>
      <price>1150</price>
      <rank>1</rank>
      <listing_type>
        <![CDATA[buy_now,instant_transfer]]>
      </listing_type>
    </item>
    <item>
      <domain>
        <![CDATA[testdomain2-2212785.club]]>
      </domain>
      <price>1150</price>
      <rank>2</rank>
      <listing_type>
        <![CDATA[buy_now,instant_transfer]]>
      </listing_type>
    </item>
    <item>
      <domain>
        <![CDATA[testdomain2-2212785.de]]>
      </domain>
      <price>1150</price>
      <rank>3</rank>
      <listing_type>
        <![CDATA[buy_now,instant_transfer]]>
      </listing_type>
    </item>
    <item>
      <domain>
        <![CDATA[testdomain2-2212785.net]]>
      </domain>
      <price>1150</price>
      <rank>4</rank>
      <listing_type>
        <![CDATA[buy_now,instant_transfer]]>
      </listing_type>
    </item>
    <item>
      <domain>
        <![CDATA[testdomain2-2212785.xyz]]>
      </domain>
      <price>1150</price>
      <rank>5</rank>
      <listing_type>
        <![CDATA[buy_now,instant_transfer]]>
      </listing_type>
    </item>
    <currency>usd</currency>
    <totalcount>5</totalcount>
  </results>
  <Command>SJL1VWRESELL_T1</Command>
  <APIType>API.NET</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>krkdt201</Server>
  <Site>eNom</Site>
  <IsLockable/>
  <IsRealTimeTLD/>
  <TimeDifference>+0.00</TimeDifference>
  <ExecTime>1.783</ExecTime>
  <Done>true</Done>
  <TrackingKey>554611c6-831e-4972-b1cb-c06dd7e33e69</TrackingKey>
  <RequestDateTime>9/27/2016 3:30:38 PM</RequestDateTime>
  <debug/>
</interface-response>
```
Related commands
----------------

[PL\_Check](../docs/PL_Check.md)