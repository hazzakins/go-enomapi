TM\GetNotice
============

Retrieve an itemized list of Trademark Clearinghouse \(TMCH\) Claims for an SLD using a Lookup Key.

Usage
-----

Use this command to retrieve an itemized list of Trademark Clearinghouse \(TMCH\) Claims for an SLD using a Lookup Key.

tcnID \(Trademark Clearinghouse Claims Notification ID\) is valid for 24 to 48 hours.

The countdown starts immediately after [TM\_Check](../docs/domains/trademark/tm-check.md) command is successfully executed.

Availability
------------

All resellers have access to this command.

Constraints
-----------
- The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | Required | TM\_Check |
| uid | string | Required | Your Account ID                                                                                |
| pw       | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(for example, enom in enom.com\)                                                          |
| TLD       | string | Required | Top-level domain name \(extension\) |
| LookupKey | string | Required | A unique Lookup Key for a domain. Use the [TM\_Check](../docs/domains/trademark/tm-check.md) command to retrieve the value. **Please use URL Encoding to avoid incorrect data being sent.** |
| ResponseType  | string | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.* |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type    | Description |
| ------------------ | ---------- | ------------------------------------------------------------------------------------------------ |
| tcnID | int    | TMCH Claims Notification ID |
| tcnStart | string   | TMCH Claims start date \(UTC\) |
| tcnExpDate | string   | TMCH Claims expiration date \(UTC\) |
| SLD | string   | Second level domain |
| MarkName | string   | Trademark name |
| Holder Entitlement | string   | Holder Entitlement |
| Org | string   | Organization name |
| Street | string   | Address |
| City | string   | City |
| State | string   | State |
| ZipCode | int, mixed | Postal Code |
| Country | string   | Country |
| Email | string   | Email address |
| Jur Desc | string   | Jurisdiction description |
| Class Desc | string   | Class description |
| GoodsAndServices | string   | Goods and Services name |
| ErrCount | int    | The number of errors if any occurred. If greater than 0, check the Err \(1 to ErrCount\) values |
| ErrX | string   | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean  | True indicates this entire response has reached you successfully. |

Example Output
--------------

The following query checks a trademark claim for test---validate.claimsgatwo domain name from Trademark Clearinghouse \(TMCH\).

```
https://resellertest.enom.com/interface.asp?UID=ResellID&PW=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&SLD=test---validate&Command=TM_GetNotice&LookupKey=2013112100%2f4%2f1%2fd%2fQdw2tSZ9X6FSqUH&responsetype={text, XML, or HTML}
```
```
<interface-response>
 <TMNotice>
 <tcnID>83ee89dc0000000000000020012</tcnID>
 <tcnStartDate>2013-11-21T00:00:00.0Z</tcnStartDate>
 <tcnExpDate>2013-11-23T00:00:00.0Z</tcnExpDate>
 <Sld>test---validate</Sld>
 <Claims>
  <Claim>
  <MarkName>Test & Validate</MarkName>
  <Holder Entitlement="owner">
  <Org>Test Organization</Org>
  <Address>
   <Street>101 West Arques Avenue</Street>
   <City>Sunnyvale</City>
   <State>CA</State>
   <ZipCode>10023-3241</ZipCode>
   <Country>US</Country>
  </Address>
  <Email>[email protected]</Email>
  </Holder>
  <Jur Desc="United States of America" Country="US" />
  <Class Desc="Musical instruments. " Num="15" />
  <GoodsAndServices>guitar</GoodsAndServices>
 </Claim>
 <Claim>
  <MarkName>Test & Validate</MarkName>
  <Holder Entitlement="owner">
  <Org>Test Organization</Org>
  <Address>
   <Street>101 West Arques Avenue</Street>
   <City>Sunnyvale</City>
   <State>CA</State>
   <ZipCode>10023-3241</ZipCode>
   <Country>US</Country>
  </Address>
  <Email>[email protected]</Email>
  </Holder>
  <Jur Desc="United States of America" Country="US" />
  <Class Desc="Precious metals and their alloys and goods
in precious metals or coated therewith, not included
in other classes; jewellery, precious stones; horological
and chronometric instruments. " Num="14" />
  <GoodsAndServices>Musical instruments</GoodsAndServices>
 </Claim>
 </TMNotice>
 <Command>TM_GetNotice</Command>
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
 <ExecTime>0.031</ExecTime>
 <Done>true</Done>
 <TrackingKey>00000000-0000-0000-0000-000000000000</TrackingKey>
 <RequestDateTime>1/10/2014 12:43:35 PM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
tcnID=83ee89dc0000000000000020012
tcnStartDate=2013-11-21T00:00:00.0Z
tcnExpDate=2013-11-23T00:00:00.0Z
Sld=test---validate
MarkName=Test & Validate
Org=Test Organization
Street=101 West Arques Avenue
City=Sunnyvale
State=CA
ZipCode=10023-3241
Country=US
[email protected]
Jur=
Class=
GoodsAndServices=guitar
MarkName=Test & Validate
Org=Test Organization
Street=101 West Arques Avenue
City=Sunnyvale
State=CA
ZipCode=10023-3241
Country=US
[email protected]
Jur=
Class=
GoodsAndServices=Musical instruments
MarkName=Test & Validate
Org=Test Organization
Street=101 West Arques Avenue
City=Sunnyvale
State=CA
ZipCode=10023-3241
Country=US
.
.
.
Command=TM_GETNOTICE
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
ExecTime=0.016
Done=true
TrackingKey=cf6a0662-9832-4bad-b5ef-0cae51319c3c
RequestDateTime=2/6/2015 12:31:10 PM
```
Related Commands
----------------
- [TM\_Check](../docs/domains/trademark/tm-check.md)
- [TM\_UpdateCart](../docs/domains/trademark/tm-updatecart.md)