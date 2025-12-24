DeleteFromCart
==============

Delete an item or all items from the shopping cart.

Usage
-----

Use this command when you want to remove one or all items from the shopping cart.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- One or all items in the shopping cart must be deleted.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=
DeleteFromCart&uid=(Required)&pw=(Required)&ItemNumber=(Required)&responsetype=xml
```
| Input Parameter | Type | Status                      | Description |
| --------------- | ------ | ------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required                     | DeleteFromCart |
| uid | string | Required | Your Account ID                                                          |
| pw       | string | Required                     | Your API Token |
| ItemNumber | string | Either ItemNumber or EmptyCart is Required | Item ID number of the item to be deleted from the shopping cart. Use the GetCartContent command to retrieve the item ID numbers. |
| EmptyCart    | string | Either Either ItemNumber or EmptyCart is Required | Use EmptyCart=On to completely empty the shopping cart. |
| ResponseType | string | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML.                             |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| ItemDeleted   | boolean | True if deleted.                                        |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=DeleteFromCart&uid=resellid&pw=resellpw
&ItemNumber=365485&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteFromCart&uid=resellid&pw=resellpw
&ItemNumber=365485&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=DeleteFromCart&uid=resellid&pw=resellpw
&ItemNumber=365485&responsetype=text
```
```
<?xml version="1.0" ?>
<interface-response>
 <ItemDeleted>True</ItemDeleted>
 <IDAction>update</IDAction>
 <Command>DELETEFROMCART</Command>
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
 <ExecTime>0.109</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/12/2011 2:52:44 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T1<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>ItemDeleted: </ STRONG>False<br>
<STRONG>IDAction: </ STRONG>update<br>
<STRONG>Command: </ STRONG>DELETEFROMCART<br>
<STRONG>APIType: </ STRONG>API<br>
<STRONG>Language: </ STRONG>eng<br>
<STRONG>ErrCount: </ STRONG>0<br>
<STRONG>ResponseCount: </ STRONG>0<br>
<STRONG>MinPeriod:</ STRONG><br>
<STRONG>MaxPeriod: </ STRONG>10<br>
<STRONG>Server: </ STRONG>SJL0VWRESELL_T1<br>
<STRONG>Site: </ STRONG>eNom<br>
<STRONG>IsLockable:</ STRONG><br>
<STRONG>IsRealTimeTLD:<br>
<STRONG>TimeDifference: </ STRONG>+0.00<br>
<STRONG>ExecTime: </ STRONG>0.047<br>
<STRONG>Done: </ STRONG>true<br>
<STRONG>RequestDateTime: </ STRONG>2/3/2015 3:20:44 PM<br>
 </HTML></BODY>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
ItemDeleted=False
IDAction=update
Command=DELETEFROMCART
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.047
Done=true
RequestDateTime=2/3/2015 3:21:08 PM
```
Related Commands
----------------

AddBulkDomains

AddToCart

GetCartContent

InsertNewOrder

PurchasePreview

UpdateCart