GetCartContent
==============

Get the contents of the shopping cart.

Usage
-----

Use this command to display the contents of the shopping cart.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=GetCartContent&uid=YourAccountID&pw=YourApiToken&responsetype={Optional}
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ---------------------------------------------------------------------------- |
| command     | string | Required | NameSpinner |
| uid | string | Required | Your Account ID                               |
| pw       | string | Required | Your API Token |
| ResponseType | string | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML. |

Returned Parameters and Values
------------------------------
- The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.
- Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                                                                                                       |
| ---------------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ItemID      | int | Item number                                                                                                                                       |
| Description   | string | Nature of the transaction. Options are: - Register - Renew                                                                                                               |
| Years      | int | Number of years to register                                                                                                                               |
| NameID      | int | Name ID                                                                                                                                         |
| Name       | string | Domain name                                                                                                                                       |
| StatusID     | string | Current status of the name. Value "A" indicates item is active and will be affected by future transactions, or "I" for inactive.                                                                           |
| ItemPrice    | float | Item price                                                                                                                                       |
| Renew      | int | Auto-renew setting. Expected values are: - 1 - Auto-Renew service is "ON". - 0 - Auto-Renew service is "OFF"                                                                                   |
| NeedsConfig   | int | A "NeedsConfig" value of 1 indicates that the "Preconfigure" icon should be displayed for this cart item. This happens when the cart item is a domain name that requires preconfiguration, and the preconfiguration has not yet been done.                      |
| IsMailPak    | int | An "IsMailPak" value of 1 indicates that the POP Email icon should be displayed for this cart item. This happens when the cart item is a domain name for which POP mail is available, and the user has not added the POP Email to the shopping cart.                 |
| ShowWPPSLink   | int | A "ShowWPPSLink" value of 1 indicates that the Whois Privacy Protection Service \(WPPS, ID Protect\) icon should be displayed for this cart item. This happens when the cart item is a domain name for which WPPS is available, and the user has not added WPPS to the shopping cart. |
| QtyDesc     | int | Unit by which this item is sold                                                                                                                             |
| SubItem     | int | A "SubItem" value of 0 indicates that this item is a single line item in the shopping cart.                                                                                              |
| FreeTrial    | int | A "FreeTrial" value of 0 indicates that the user begins paying for this item as soon as it is purchased                                                                                         |
| command     | string | Contacts                                                                                                                                        |
| errcount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                            |
| errX       | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                                                |
| done       | string | True value indicates this entire response has reached you successfully.                                                                                                        |

Example Output
--------------

The following query requests the contents of the shopping cart for account resellid and requests the response in format of either "text" \(default\), "xml" or "html"

The response indicates that the shopping cart currently contains a domain to register \(resellerdocs.us\)

```
<interface-response>
<GetCartContent>
<ContainsDomain>true</ContainsDomain>
<item>
<ItemID>889018</ItemID>
<Description>Email Pak - 10 email boxes</Description>
<ProdType>38</ProdType>
<Years>1</Years>
<NameID>152533676</NameID>
<Name>resellerdocs.com</Name>
<StatusID>A</StatusID>
<ItemPrice>9.95</ItemPrice>
<ItemTotalPrice>9.9500</ItemTotalPrice>
<ItemUnitPrice>9.9500</ItemUnitPrice>
<Renew>0</Renew>
<cd>
<IsMailPak>1</IsMailPak>
<ShowWPPSLink>0</ShowWPPSLink>
<ShowWBLLink>0</ShowWBLLink>
<ShowMobilizerLink>0</ShowMobilizerLink>
<ShowMailLink>0</ShowMailLink>
<p>
<QtyDesc>Pak(s)</QtyDesc>
<SubItem>0</SubItem>
<FreeTrial>0</FreeTrial>
</p>
</cd>
</item>
<item>
<ItemID>889049</ItemID>
<Description>Register</Description>
<ProdType>10</ProdType>
<Years>1</Years>
<NameID>152932786</NameID>
<Name>resellerdocs1.com</Name>
<StatusID>A</StatusID>
<ItemPrice>8.95</ItemPrice>
<ItemTotalPrice>8.9500</ItemTotalPrice>
<ItemUnitPrice>8.9500</ItemUnitPrice>
<Renew>0</Renew>
<cd>
<IsMailPak>0</IsMailPak>
<ShowWPPSLink>1</ShowWPPSLink>
<ShowWBLLink>1</ShowWBLLink>
<ShowMobilizerLink>0</ShowMobilizerLink>
<ShowMailLink>1</ShowMailLink>
<p>
<QtyDesc>Year(s)</QtyDesc>
<SubItem>0</SubItem>
<FreeTrial>0</FreeTrial>
<ICANNFees>0.20</ICANNFees>
<BaseItemPrice>8.75</BaseItemPrice>
</p>
</cd>
</item>
<item>
<ItemID>889066</ItemID>
<Description>Register</Description>
<ProdType>10</ProdType>
<Years>1</Years>
<NameID>152932796</NameID>
<Name>resellerdocs123.com</Name>
<StatusID>A</StatusID>
<ItemPrice>8.95</ItemPrice>
<ItemTotalPrice>8.9500</ItemTotalPrice>
<ItemUnitPrice>8.9500</ItemUnitPrice>
<Renew>0</Renew>
<cd>
<IsMailPak>0</IsMailPak>
<ShowWPPSLink>1</ShowWPPSLink>
<ShowWBLLink>1</ShowWBLLink>
<ShowMobilizerLink>0</ShowMobilizerLink>
<ShowMailLink>1</ShowMailLink>
<p>
<QtyDesc>Year(s)</QtyDesc>
<SubItem>0</SubItem>
<FreeTrial>0</FreeTrial>
<ICANNFees>0.20</ICANNFees>
<BaseItemPrice>8.75</BaseItemPrice>
</p>
</cd>
</item>
</GetCartContent>
<GroupCounts>
<Group ParentItemID="889049" GroupCount="1"/>
<Group ParentItemID="889066" GroupCount="1"/>
<Group ParentItemID="889018" GroupCount="1"/>
</GroupCounts>
<CartItems>3</CartItems>
<Command>GETCARTCONTENT</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod/>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.843</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 3:43:27 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>ItemID1: </STRONG>988418<br />
<STRONG>Description1: </STRONG>SSL Certificate - RapidSSL<br />
<STRONG>ProdType1: </STRONG>23<br />
<STRONG>Years1: </STRONG>2<br />
<STRONG>NameID1: </STRONG> <br />
<STRONG>Name1: </STRONG> <br />
<STRONG>StatusID1: </STRONG>I<br />
<STRONG>ItemPrice1: </STRONG>9.95<br />
<STRONG>Renew1: </STRONG> <br />
<STRONG>ParentItemID1: </STRONG> <br />
<STRONG>ICANNFees1: </STRONG> <br />
<STRONG>BaseItemPrice1: </STRONG> <br />
<STRONG>ItemID2: </STRONG>988419<br />
<STRONG>Description2: </STRONG>GeoTrust Anti-Malware Scan<br />
<STRONG>ProdType2: </STRONG>177<br />
<STRONG>Years2: </STRONG>1<br />
<STRONG>NameID2: </STRONG> <br />
<STRONG>Name2: </STRONG> <br />
<STRONG>StatusID2: </STRONG>I<br />
<STRONG>ItemPrice2: </STRONG>15.00<br />
<STRONG>Renew2: </STRONG> <br />
<STRONG>ParentItemID2: </STRONG> <br />
<STRONG>ICANNFees2: </STRONG> <br />
<STRONG>BaseItemPrice2: </STRONG> <br />
<STRONG>ItemID3: </STRONG>987632<br />
<STRONG>Description3: </STRONG>RichContent Free Trial<br />
<STRONG>ProdType3: </STRONG>200<br />
<STRONG>Years3: </STRONG>1<br />
<STRONG>NameID3: </STRONG> <br />
<STRONG>Name3: </STRONG>[email protected]<br />
<STRONG>StatusID3: </STRONG>I<br />
<STRONG>ItemPrice3: </STRONG>0.00<br />
<STRONG>Renew3: </STRONG> <br />
<STRONG>ParentItemID3: </STRONG> <br />
<STRONG>ICANNFees3: </STRONG>0<br />
<STRONG>BaseItemPrice3: </STRONG>0<br />
<STRONG>ItemID4: </STRONG>988111<br />
<STRONG>Description4: </STRONG>Register<br />
.
.
.
<STRONG>GroupCount: </STRONG>34<br />
<STRONG>CartItems: </STRONG>34<br />
<STRONG>Command: </STRONG>GETCARTCONTENT<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG> <br />
<STRONG>IsRealTimeTLD: </STRONG> <br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.594<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/3/2015 4:43:53 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
ItemID1=988418
Description1=SSL Certificate - RapidSSL
ProdType1=23
Years1=2
NameID1=
Name1=
StatusID1=I
ItemPrice1=9.95
Renew1=
ParentItemID1=
ICANNFees1=
BaseItemPrice1=
ItemID2=988419
Description2=GeoTrust Anti-Malware Scan
ProdType2=177
Years2=1
NameID2=
Name2=
StatusID2=I
ItemPrice2=15.00
Renew2=
ParentItemID2=
ICANNFees2=
BaseItemPrice2=
ItemID3=987632
Description3=RichContent Free Trial
ProdType3=200
Years3=1
NameID3=
Name3=; [email protected]
StatusID3=I
ItemPrice3=0.00
Renew3=
ParentItemID3=
ICANNFees3=0
BaseItemPrice3=0
.
.
.
GroupCount=34
CartItems=34
Command=GETCARTCONTENT
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.391
Done=true
RequestDateTime=2/3/2015 4:45:21 PM
```
Related Commands
----------------

AddBulkDomains

AddToCart

DeleteFromCart

InsertNewOrder

PurchasePreview

UpdateCart