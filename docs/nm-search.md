NM\Search
=========

Search for available Premium Domains.

Usage
-----

Use this command to search for Premium Domains that are available for purchase.

Availability
------------

Premium Domains can only be sold by our direct ETP resellers.

Constraints
-----------

The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter | Type | Required                              | Description |
| --------------- | ------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | Required                              | NM\_Search |
| uid | string    | Required | Your Account ID                                                                                                                                                                                         |
| pw       | string | Required                              | Your API Token |
| Keyword | string    | Recommended; default is no keyword. | Word to include in search results. **Does not function if query string includes the StartsWith or EndsWith parameter. This parameter expects full words to function as intended.** For example, Keyword=Box,Office returns matches that contain Box and Office, followed by matches that contain Box or Office. *Permitted values are single strings or comma-delimited lists of strings.* |
| StartsWith   | string | Optional; default is no filter.                  | Return domain names that start with this character. *Permitted values are alphanumeric.* |
| EndsWith | string    | Optional; default is no filter. | Return domain names that end with this character. *Permitted values are alphanumeric.*                                                                                                                                                     |
| Category    | string | Optional; by default, the scope of a search is all categories.  | Topic category to search. Use the [NM\_GetSearchCategories](../docs/nm-getsearchcategories.md) command to retrieve search categories |
| Subcategory | string    | Optional; by default, the scope of a search is all subcategories. | Topic subcategory to search. Use the [NM\_GetSearchCategories](../docs/nm-getsearchcategories.md) to retrieve search subcategories                                                                                                                              |
| TLDList     | string | Optional; default is com                      | TLDs to include in results. For example, TLDList=com,net,org returns .com, .net, and .org matches to your search *Permitted format is tld or tld1,tld2,tld3, . . ..* |
| ExcludeHyphens | string    | Optional; default is No | Exclude domain names that contain hyphens? *Use ExcludeHyphens=Yes to exclude hyphens.*                                                                                                                                                     |
| MinLength    | int | Optional; default is no minimum.                 | Minimum number of characters in the SLDs returned by the search. Most registries impose a minimum of 1 letter for TLDs. *Permitted values are numeric* |
| MaxLength | int     | Optional; default is no maximum | Maximum number of characters in the SLDs returned by the search. Most registries impose a maximum of 63 characters for SLDs. *Permitted values are numeric*                                                                                                                  |
| PriceHigh    | int, decimal | Optional; default is no maximum                  | Filter return so that all results are at or below this price. We don’t impose a maximum price but you may want to, to manage risk. *Permitted values are in DD or DD.cc format* |
| PriceLow | int, decimal | Optional; default is no minimum | Filter return so that all results are at or above this price. *Permitted values are in DD or DD.cc format*                                                                                                                                           |
| Updated     | string | Optional                              | Date Added to Inventory |
| RecordsToReturn | int     | Optional; default is 50 | Number of domains to return in this response. For example, if you return domains 25 at a time and want to retrieve the third set of 25 domains, use RecordsToReturn=25&StartPosition=51 *Permitted values are integers 1 to 500.*                                                                               |
| StartPosition  | int | Optional; default is 1                       | For this set of results, start with this position in the set of 25 domains, use RecordsToReturn=25&StartPosition=51 to retrieve the third set of 25 domains, use RecordsToReturn=25&StartPosition=51 |
| ResponseType | string    | Optional | Format of response. *Permitted values are Text \(default\), HTML, or XML.*                                                                                                                                                           |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------- |
| Count      | int | Number of Premium Domains returned in this response                                       |
| PremiumDomain  | string | Premium Domain name. Additionally, this element contains an XML attribute "price" that contains the price of the domain name. |
| Command     | string | Name of command executed                                                    |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.                    |
| Done       | boolean | True indicates this entire response has reached you successfully.                               |

Example Output
--------------

The following query searches for available premium domains that match the keyword "example" and limits the amount of domains returned using the "RecordsToReturn" parameter.

```
https://resellertest.enom.com/interface.asp?command=NM_Search&uid=resellid&pw=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&keyword=example&recordstoreturn=3&responsetype={text, XML, or HTML}
```
```
<interface-response>
<domains count="3">
<premiumdomain price="1950.00">ExampleInc.com</premiumdomain>
<premiumdomain price="2869.00">ExamPlea.com</premiumdomain>
<premiumdomain price="4162.00">ExampleMail.com</premiumdomain>
</domains>
<Command>NM_SEARCH</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>sjl1vwresell_t1</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.125</ExecTime>
<Done>true</Done>
<TrackingKey>ae4436ac-6c14-451c-9c13-28b459128c03</TrackingKey>
<RequestDateTime>7/3/2016 12:19:29 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL1VWRESELL_T
;Encoding Type is utf-8
price1=1950.00
premiumdomain1=ExampleInc.com
price2=2869.00
premiumdomain2=ExamPlea.com
price3=4162.00
premiumdomain3=ExampleMail.com
premiumdomain=ExampleInc.com
premiumdomain=ExamPlea.com
premiumdomain=ExampleMail.com
Command=NM_SEARCH
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl1vwresell_t
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.125
Done=true
TrackingKey=f96b16c2-218c-4a93-9aca-bc1a2ed84860
RequestDateTime=7/3/2016 12:21:28 PM
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>price1: </STRONG>1950.00<BR /><STRONG>premiumdomain1: </STRONG>ExampleInc.com<BR /><STRONG>price2: </STRONG>2869.00<BR /><STRONG>premiumdomain2: </STRONG>ExamPlea.com<BR /><STRONG>price3: </STRONG>4162.00<BR /><STRONG>premiumdomain3: </STRONG>ExampleMail.com<BR /><STRONG>premiumdomain: </STRONG>ExampleInc.com<BR /><STRONG>premiumdomain: </STRONG>ExamPlea.com<BR /><STRONG>premiumdomain: </STRONG>ExampleMail.com<BR /><STRONG>Command: </STRONG>NM_SEARCH<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.109<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>e8584c11-7ab9-496d-8f9b-63b425a01efc<BR /><STRONG>RequestDateTime: </STRONG>7/3/2016 12:21:42 PM<BR /></BODY></HTML>
```
Related Commands
----------------
- [NM\_GetSearchCategories](../docs/nm-getsearchcategories.md)
- [NM\_GetPremiumDomainSettings](../docs/nm-getpremiumdomainsettings.md)
- [Check](http://www.enom.com/api/API%20topics/api_Check.htm)