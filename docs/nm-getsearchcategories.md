NM\GetSearchCategories
======================

Retrieve Premium Domain search categories and subcategories.

Usage
-----

Use this command to retrieve search categories and/or subcategories for Premium Domains.

You can use this command to populate the Category and Subcategory values in the [NM\_Search](../docs/nm-search.md) command.These would typically be rendered as Category and Subcategory menus in a user interface.

Availability
------------

Premium Domains can only be sold by our direct ETP resellers.

Constraints
-----------

The login ID and API Token must be valid.

Input Parameters
----------------

| Input Parameter | Type | Status                     | Description |
| --------------- | ------ | ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required                    | NM\_GetPremiumDomainSettings |
| uid | string | Required | Your Account ID                                                                                                            |
| pw       | string | Required                    | Your API Token |
| Display | string | Optional; default is All | Which part of the search category taxonomy to display. *Permitted values are:* - All: - Retrieve all categories and sub categories - Categories: - Retrieve categories - Subcategories - Retrieve the subcategories for one category |
| CategoryID   | int | Optional **Required if Display= Subcategories** | Which category to display subcategories for. *Permitted values are category ID numbers, which you can retrieve by running this command with Display=Categories.* |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| ID | int   | Identification number of this category or subcategory |
| ParentID | int   | Identification number of this subcategory’s parent category |
| Name | string | Name of this category or subcategory |
| Command | string | Name of command executed |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | boolean | True indicates this entire response has reached you successfully. |

Example Output
--------------

The following query retrieves all premium domain categories without retrieving subcategories.

```
https://resellertest.enom.com/interface.asp?command=NM_GetSearchCategories&display=categories&uid=resellid&pw=ANLOYHTJB2J6YBZC6PWP2HNJ6TJS7XDEMDWMGSFE&responsetype={text, XML, or HTML}
```
```
<interface-response>
<category id="1" name="Business"/>
<category id="30" name="Careers"/>
<category id="44" name="Computers"/>
<category id="58" name="Education"/>
<category id="69" name="Family Life"/>
<category id="85" name="Financial"/>
<category id="105" name="Health"/>
<category id="141" name="Home"/>
<category id="169" name="Recreation"/>
<category id="190" name="Reference"/>
<category id="199" name="Region"/>
<category id="213" name="Special Events"/>
<category id="231" name="Sports"/>
<category id="250" name="Shopping"/>
<category id="278" name="Society"/>
<category id="300" name="Travel"/>
<Command>NM_GETSEARCHCATEGORIES</Command>
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
<ExecTime>0.313</ExecTime>
<Done>true</Done>
<TrackingKey>48baee94-d8c4-4ba1-9585-1280065cdce1</TrackingKey>
<RequestDateTime>7/3/2016 2:02:10 PM</RequestDateTime>
<debug/>
</interface-response>
```
```
;URL Interface
;Machine is SJL1VWRESELL_T
;Encoding Type is utf-8
categoryid1=1
categoryname1=Business
categoryid2=30
categoryname2=Careers
categoryid3=44
categoryname3=Computers
categoryid4=58
categoryname4=Education
categoryid5=69
categoryname5=Family Life
categoryid6=85
categoryname6=Financial
categoryid7=105
categoryname7=Health
categoryid8=141
categoryname8=Home
categoryid9=169
categoryname9=Recreation
categoryid10=190
categoryname10=Reference
categoryid11=199
categoryname11=Region
categoryid12=213
categoryname12=Special Events
categoryid13=231
categoryname13=Sports
categoryid14=250
categoryname14=Shopping
categoryid15=278
categoryname15=Society
categoryid16=300
categoryname16=Travel
Command=NM_GETSEARCHCATEGORIES
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
ExecTime=0.016
Done=true
TrackingKey=db6a17bc-48b3-4322-8e07-c9e44abb5c3f
RequestDateTime=7/3/2016 2:03:14 PM
```
```
;URL Interface<br>
;Machine is SJL1VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY><STRONG>categoryid1: </STRONG>1<BR /><STRONG>categoryname1: </STRONG>Business<BR /><STRONG>categoryid2: </STRONG>30<BR /><STRONG>categoryname2: </STRONG>Careers<BR /><STRONG>categoryid3: </STRONG>44<BR /><STRONG>categoryname3: </STRONG>Computers<BR /><STRONG>categoryid4: </STRONG>58<BR /><STRONG>categoryname4: </STRONG>Education<BR /><STRONG>categoryid5: </STRONG>69<BR /><STRONG>categoryname5: </STRONG>Family Life<BR /><STRONG>categoryid6: </STRONG>85<BR /><STRONG>categoryname6: </STRONG>Financial<BR /><STRONG>categoryid7: </STRONG>105<BR /><STRONG>categoryname7: </STRONG>Health<BR /><STRONG>categoryid8: </STRONG>141<BR /><STRONG>categoryname8: </STRONG>Home<BR /><STRONG>categoryid9: </STRONG>169<BR /><STRONG>categoryname9: </STRONG>Recreation<BR /><STRONG>categoryid10: </STRONG>190<BR /><STRONG>categoryname10: </STRONG>Reference<BR /><STRONG>categoryid11: </STRONG>199<BR /><STRONG>categoryname11: </STRONG>Region<BR /><STRONG>categoryid12: </STRONG>213<BR /><STRONG>categoryname12: </STRONG>Special Events<BR /><STRONG>categoryid13: </STRONG>231<BR /><STRONG>categoryname13: </STRONG>Sports<BR /><STRONG>categoryid14: </STRONG>250<BR /><STRONG>categoryname14: </STRONG>Shopping<BR /><STRONG>categoryid15: </STRONG>278<BR /><STRONG>categoryname15: </STRONG>Society<BR /><STRONG>categoryid16: </STRONG>300<BR /><STRONG>categoryname16: </STRONG>Travel<BR /><STRONG>Command: </STRONG>NM_GETSEARCHCATEGORIES<BR /><STRONG>APIType: </STRONG>API.NET<BR /><STRONG>Language: </STRONG>eng<BR /><STRONG>ErrCount: </STRONG>0<BR /><STRONG>ResponseCount: </STRONG>0<BR /><STRONG>MinPeriod: </STRONG>1<BR /><STRONG>MaxPeriod: </STRONG>10<BR /><STRONG>Server: </STRONG>sjl1vwresell_t1<BR /><STRONG>Site: </STRONG>eNom<BR /><STRONG>IsLockable: </STRONG><BR /><STRONG>IsRealTimeTLD: </STRONG><BR /><STRONG>TimeDifference: </STRONG>+0.00<BR /><STRONG>ExecTime: </STRONG>0.000<BR /><STRONG>Done: </STRONG>true<BR /><STRONG>TrackingKey: </STRONG>d414b52f-4888-421e-9a95-abec8827e672<BR /><STRONG>RequestDateTime: </STRONG>7/3/2016 2:03:22 PM<BR /></BODY></HTML>
```
Related Commands
----------------
- [NM\_Search](../docs/nm-search.md)
- [NM\_ProcessOrder](../docs/nm-processorder.md)
- [NM\_SetPremiumDomainSettings](../docs/nm-setpremiumdomainsettings.md)