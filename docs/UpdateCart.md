UpdateCart
==========

Update items in the shopping cart.

Usage
-----

Use this command to activate or deactivate a shopping cart item or to change its quantity.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The item must be in the customer’s shopping cart.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=UpdateCart&uid=(Required)&pw=(Required)&ItemID1=(Required)&ItemStatus1=(Required)&ItemQty1=(Required)&responsetype=(Optional)
```
| Input Parameter         | Type | Status  | Description |
| ------------------------------- | ------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command             | string | Required | UpdateCart |
| uid | string | Required | Your Account ID                                                                                     |
| pw               | string | Required | Your API Token |
| ItemIDX\(X=1 to NumberOfItems\) | string | Required | Input value is the six-digit cart item ID. Retrieve this number using the GetCartContent command.                                           |
| ItemStatusX           | string | Optional | New status of the item. Permitted values are on to make the status active, off to make it inactive. Items set to status on are subject to the next operation you perform on the cart. |
| ItemQtyX | string | Required | New quantity for the item                                                                                |
| RenewX             | string | Optional | Renewal setting for the item. Permitted values are on to auto-renew, none for an item that is not renewable. If RenewX is not supplied, item is set to not auto-renew. |
| ClearItems | string | Optional | Toggle all items in cart to status I \(Inactive\). Permitted value is Yes.                                                       |
| ActivateItems          | string | Optional | Toggle all items in cart to status A \(Active\). Permitted value is Yes. |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML                                                        |

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
https://resellertest.enom.com/interface.asp?
command=updatecart&uid=resellid&pw=resellpw&ItemID1=365690
&ItemStatus1=on&ItemQty1=1&Renew1=1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=updatecart&uid=resellid&pw=resellpw&ItemID1=365690
&ItemStatus1=on&ItemQty1=1&Renew1=1&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=updatecart&uid=resellid&pw=resellpw&ItemID1=365690
&ItemStatus1=on&ItemQty1=1&Renew1=1&ResponseType=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <Command>UPDATECART</Command>
 <ErrCount>0</ErrCount>
 <Server>Dev Workstation</Server>
 <Site>enom</Site>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>IDAction: </ STRONG>update<br>
<STRONG>Command: </ STRONG>UPDATECART<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:</ STRONG><br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.297<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/9/2015 1:57:33 PM<br>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
IDAction=update
Command=UPDATECART
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
ExecTime=0.266
Done=true
RequestDateTime=2/9/2015 1:57:54 PM
```
Related Commands
----------------

AddBulkDomains

AddToCart

DeleteFromCart

GetCartContent

InsertNewOrder

PurchasePreview