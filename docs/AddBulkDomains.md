AddBulkDomains
==============

Add a list of items to the shopping cart.

Usage
-----

Use this command when you want to add multiple items to the shopping cart.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain names must use a supported top-level domain.
- The number of SLDs must match the number of TLDs.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=
addbulkdomains&uid=(Required)&pw=
(Required)&producttype=register&listcount=(Required)&sld1=(Required)&tld1=(Required)&sld2=(Optional)&tld2=(Optional)&sld3=(Optional)&tld3=
(Optional)&responsetype=xml
```
| Input Parameter        | Type | Status                                 | Description |
| ------------------------------ | ------ | ---------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command            | string | Required                                | TP\_GetOrderStatuses |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                    |
| pw               | string | Required                                | Your API Token |
| ProductType | string | Required | Type of product to add. Options are - register - renew.                                                                                                                                                                                               |
| ListCount           | string | Required                                | Number of names to add |
| SLDX \(X=1 to listcount\) | string | Required | Second-level domain name, for example, *resellerdocs* in resellerdocs.com                                                                                                                                                                                      |
| TLDX \(X=1 to listcount\)   | string | Required                                | Top-level domain name, for example, *com* in resellerdocs.com |
| numyearsX \(X=1 to listcount\) | string | Optional | Number of years to renew                                                                                                                                                                                                                |
| AutoRenew           | string | Optional; default is 0                         | Auto-renew setting for all domains in this order. If AutoRenew=1, domains automatically update 30 days before expiration |
| RegLock | string | Optional; default is 1 | Registrar lock setting for all domains in this order. If RegLock=1, domain cannot be transferred to another registrar without account holder’s permission                                                                                                                                               |
| UseCart            | string | UseCart=1 is Required if UID is a retail account, otherwise Optional. | Use the shopping cart for this order. Permitted values are: - 0 - 1 If UseCart=1, this bulk list will go into the shopping cart and be processed through our queue; our system does not return an order ID until all names have been processed. If UseCart=0, our system returns an order ID and holds the funds for the order; the funds remain held until all names in the list are processed. UseCart=1 is required if UID is a retail account. |
| ResponseType | string | Optional | Format of response. Permitted values are: - Text \(default\) - HTML - XML.                                                                                                                                                                                      |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?command=addbulkdomains
&uid=resellid&pw=resellpw&producttype=register&listcount=3
&sld1=resellerdocs&tld1=com&sld2=resellerdocs2&tld2=net
&sld3=resellerdocs3&tld3=info&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?command=addbulkdomains
&uid=resellid&pw=resellpw&producttype=register&listcount=3
&sld1=resellerdocs&tld1=com&sld2=resellerdocs2&tld2=net
&sld3=resellerdocs3&tld3=info&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?command=addbulkdomains
&uid=resellid&pw=resellpw&producttype=register&listcount=3
&sld1=resellerdocs&tld1=com&sld2=resellerdocs2&tld2=net
&sld3=resellerdocs3&tld3=info&responsetype=text
```
```
<interface-response>
 <AddBulkDomains>
 <Item>
  <WscAccountOverride>False</WscAccountOverride>
  <ItemName>resellerdocs1314.com</ItemName>
  <ItemId>889487</ItemId>
  <Price>8.95</Price>
  <ICANNFees>0.20</ICANNFees>
  <CartItemID>889487</CartItemID>
  <NewDomainNameID>152933101</NewDomainNameID>
  <ItemAdded>True</ItemAdded>
  <DomainName>resellerdocs1314.com</DomainName>
 </Item>
 <Item>
  <WscAccountOverride>False</WscAccountOverride>
  <ItemName>resellerdocs465442.net</ItemName>
  <ItemId>889488</ItemId>
  <Price>8.95</Price>
  <ICANNFees>0.20</ICANNFees>
  <CartItemID>889488</CartItemID>
  <NewDomainNameID>152933102</NewDomainNameID>
  <ItemAdded>True</ItemAdded>
  <DomainName>resellerdocs465442.net</DomainName>
 </Item>
 <Item>
  <WscAccountOverride>False</WscAccountOverride>
  <ItemName>resellerdoc213131s3.info</ItemName>
  <ItemId>889489</ItemId>
  <Price>8.95</Price>
  <ICANNFees>0.20</ICANNFees>
  <CartItemID>889489</CartItemID>
  <NewDomainNameID>152933103</NewDomainNameID>
  <ItemAdded>True</ItemAdded>
  <DomainName>resellerdoc213131s3.info</DomainName>
 </Item>
 <ListCount>3</ListCount>
 <CartErrors>0</CartErrors>
 <AllItemsSuccessful>True</AllItemsSuccessful>
 <CartItems>6</CartItems>
 </AddBulkDomains>
 <Success>True</Success>
 <UseCart>True</UseCart>
 <Command>ADDBULKDOMAINS</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod/>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable/>
 <IsRealTimeTLD/>
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.219</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>845c0d1a-608a-49b4-bc5e-6694db329877</TrackingKey>
 <RequestDateTime>12/12/2011 12:59:58 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>WscAccountOverride1: </ STRONG>False<br>
<STRONG>ItemName1: </ STRONG>resellerdocs.com<br>
<STRONG>ItemId1:</ STRONG><br>
<STRONG>Price1:</ STRONG><br>
<STRONG>ICANNFees1:</ STRONG><br>
<STRONG>CartItemID1:</ STRONG><br>
<STRONG>NewDomainNameID1:</ STRONG><br>
<STRONG>ItemAdded1: </ STRONG>False<br>
<STRONG>ItemError1: </ STRONG>resellerdocs.com is already in your Shopping Cart.<br>
<STRONG>DomainName1: </ STRONG>resellerdocs.com<br>
<STRONG>WscAccountOverride2: </ STRONG>False<br>
<STRONG>ItemName2: </ STRONG>resellerdocs2.net<br>
<STRONG>ItemId2:</ STRONG><br>
<STRONG>Price2:</ STRONG><br>
<STRONG>ICANNFees2:</ STRONG><br>
<STRONG>CartItemID2:</ STRONG><br>
<STRONG>NewDomainNameID2:</ STRONG><br>
<STRONG>ItemAdded2: </ STRONG>False<br>
<STRONG>ItemError2: </ STRONG>resellerdocs2.net is already in your Shopping Cart.<br>
<STRONG>DomainName2: </ STRONG>resellerdocs2.net<br>
<STRONG>WscAccountOverride3: </ STRONG>False<br>
<STRONG>ItemName3: </ STRONG>resellerdocs3.info<br>
<STRONG>ItemId3:</ STRONG><br>
<STRONG>Price3:</ STRONG><br>
<STRONG>ICANNFees3:</ STRONG><br>
<STRONG>CartItemID3:</ STRONG><br>
<STRONG>NewDomainNameID3:</ STRONG><br>
<STRONG>ItemAdded3: </ STRONG>False<br>
<STRONG>ItemError3: </ STRONG>resellerdocs3.info is already in your Shopping Cart.<br>
<STRONG>DomainName3: </ STRONG>resellerdocs3.info<br>
<STRONG>ListCount: </ STRONG>3<br>
<STRONG>CartErrors: </ STRONG>3<br>
<STRONG>AllItemsSuccessful: </ STRONG>False<br>
<STRONG>CartItems: </ STRONG>29<br>
<STRONG>Success: </ STRONG>True<br>
<STRONG>UseCart: </ STRONG>True<br>
<STRONG>Command: </ STRONG>ADDBULKDOMAINS<br>
<STRONG>APIType: </ STRONG>API.NET<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod: </ STRONG>1<br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>sjl0vwresell_t1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.125<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>TrackingKey: </ STRONG>3cf3d460-a9f3-4c3f-81b9-7878d93987c5<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 11:29:08 AM<br>
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
WscAccountOverride1=False
ItemName1=resellerdocs.com
ItemId1=
Price1=
ICANNFees1=
CartItemID1=
NewDomainNameID1=
ItemAdded1=False
ItemError1=resellerdocs.com is already in your Shopping Cart.
DomainName1=resellerdocs.com
WscAccountOverride2=False
ItemName2=resellerdocs2.net
ItemId2=
Price2=
ICANNFees2=
CartItemID2=
NewDomainNameID2=
ItemAdded2=False
ItemError2=resellerdocs2.net is already in your Shopping Cart.
DomainName2=resellerdocs2.net
WscAccountOverride3=False
ItemName3=resellerdocs3.info
ItemId3=
Price3=
ICANNFees3=
CartItemID3=
NewDomainNameID3=
ItemAdded3=False
ItemError3=resellerdocs3.info is already in your Shopping Cart.
DomainName3=resellerdocs3.info
ListCount=3
CartErrors=3
AllItemsSuccessful=False
CartItems=29
Success=True
UseCart=True
Command=ADDBULKDOMAINS
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
ExecTime=0.109
Done=true
TrackingKey=c4096447-3143-4d38-a8d8-582b7d1a56ce
RequestDateTime=2/3/2015 11:31:38 AM
```
Related Commands
----------------

[AddToCart](../docs/addtocart.md)

[Check](../docs/check.md)

[DeleteFromCart](../docs/DeleteFromCart.md)

[GetCartContent](../docs/getcartcontent.md)

[InsertNewOrder](../docs/InsertNewOrder.md)

[Purchase](../docs/purchase.md)

[PurchasePreview](../docs/purchasepreview.md)

[UpdateCart](../docs/UpdateCart.md)