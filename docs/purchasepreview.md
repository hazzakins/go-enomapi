PurchasePreview
===============

Preview a shopping cart order.

Usage
-----

Use this command to view the items that will be purchased if you check out now.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=PurchasePreview&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | ------------------------------------------------------------------------- |
| command     | string | Required | PurchasePreview |
| uid | string | Required | Your Account ID                              |
| pw       | string | Required | Your API Token |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| ItemID      | int | Shopping cart item ID, from our internal records                         |
| Description   | string | Description of shopping cart item                                |
| Years      | int | Number of time units customer wants to subscribe                         |
| NameID      | string | Domain name ID, from our internal records                            |
| Name       | string | Name the customer has requested for this item                          |
| StatusID     | string | Status ID                                            |
| ItemPrice    | int | Price for one unit of this item                                 |
| ExtPrice     | int | ItemPrice multiplied by Years \(number of time units\)                      |
| NeedsConfig   | string | Flag that marks domains requiring pre-configuration                       |
| ProductType   | string | Product type ID number, from our internal records                        |
| FreeTrial    | string | Is this product being offered as a free trial?                          |
| ParentItemID   | string | Parent item ID number, from our internal records                         |
| ICANNFees    | int | Fees charged by ICANN for this product                              |
| BasePrice    | int | Price of this product before the ICANN fee                            |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

```
https://resellertest.enom.com/interface.asp?
command=purchasepreview&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=purchasepreview&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=purchasepreview&uid=resellid
&pw=resellpw&responsetype=text
```
```
<interface-response>
 <Purchase-Summary>
 <PremDomainSalesAmount/>
 <Reseller>True</Reseller>
 <Balance>2454.8</Balance>
 <AvailableBalance>2083.85</AvailableBalance>
 <DomainCount>0</DomainCount>
 <DotComPrice>8.95</DotComPrice>
 <DotNetPrice>8.95</DotNetPrice>
 <DotOrgPrice>8.95</DotOrgPrice>
 <DotCcPrice>24.95</DotCcPrice>
 <DotTvPrice>39.95</DotTvPrice>
 <TotalPrice>0.00</TotalPrice>
 <TotalDiscount>0</TotalDiscount>
 </Purchase-Summary>
 <Command>PURCHASEPREVIEW</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>0.984</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/9/2011 4:03:28 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>PremDomainSalesAmount:</STRONG><br />
<STRONG>Reseller: </STRONG>True<br />
<STRONG>Balance: </STRONG>-1955.78<br />
<STRONG>AvailableBalance: </STRONG>6069.8225<br />
<STRONG>Command: </STRONG>PURCHASEPREVIEW<br />
<STRONG>APIType: </STRONG>API<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod:</STRONG><br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>SJL0VWRESELL_T<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable:</STRONG><br />
<STRONG>IsRealTimeTLD:</STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.281<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>RequestDateTime: </STRONG>2/5/2015 11:01:59 AM<br />
 </BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
PremDomainSalesAmount=
Reseller=True
Balance=-1955.78
AvailableBalance=6069.8225
Command=PURCHASEPREVIEW
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
ExecTime=0.063
Done=true
RequestDateTime=2/5/2015 11:02:27 AM
```
Related Commands
----------------

[AddBulkDomains](../docs/addbulkdomains.md) BulkDomains

[AddToCart](../docs/addtocart.md) oCart

DeleteFromCart

GetCartContent

InsertNewOrder

UpdateCart