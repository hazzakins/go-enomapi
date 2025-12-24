Extend\RGP
==========

Renew domain names that are in RGP status, in near-real time. This command charges not only the usual renewal fee, but also the higher redemption fee.

Usage
-----

Use this command to extend names that are in RGP status. This command does not yet accommodate names in Extended RGP status. Use the GetExpiredDomains command to check a name’s status.

RGP redemptions are performed at the Registry and generally take at least 3 hours to complete. When you run this command, we recommend that you store the Order ID so that you can check progress conveniently using the GetOrderDetail command.

This command will result in both a redemption and a renewal charge for the domain name. We recommend that you check the RGP redemption price before running this command, as it can be quite high.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- This command is available only to ETP reseller accounts.
- The login ID and password must be valid.
- The domain name must belong to this account.
- This command does not support our credit card processing; it deducts from the reseller balance.
- This command applies only to names in RGP status, not Extended RGP or domains in Expired status. Todetermine the status of a name, use the GetExpiredDomains command.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                  | Max Size |
| --------------- | ----------------------------- | -------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) of name in Redemption Grace Period \(RGP\) status | 63 |
| TLD | Required           | Top-level domain name \(extension\) of name in RGP status | 15    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                             | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Extension | Success status of the query                                   |
| DomainName    | Name of the domain to be redeemed |
| OrderID | Order ID number. Store this number to simplify progress checks.                |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- We recommend that you check the redemption price for RGP names before running this command, as theprice can be quite high. Use the PE\_GetResellerPrice command, and look at product types 16 \(renew\) and 17 \(RGP\).
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query redeems and renews a domain name that is in RGP status, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=Extend_RGP&uid=resellid&pw=resellpw
&sld=VeryExpiredName&tld=com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=Extend_RGP&uid=resellid&pw=resellpw
&sld=VeryExpiredName&tld=com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=Extend_RGP&uid=resellid&pw=resellpw
&sld=VeryExpiredName&tld=com&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
<DomainRRP>E</DomainRRP>
<Extension>successful</Extension>
<DomainName>VeryExpiredName.com</DomainName>
<OrderID>156251141</OrderID>
<Command>EXTEND_RGP</Command>
<Language>en</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site>enom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+03.00</TimeDifference>
<ExecTime>0.2109375</ExecTime>
<Done>true</Done>
<debug>
 <![CDATA [ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
DomainRRP: E
Extension: successful
DomainName: VeryExpiredName.com
OrderID: 156251141
Command: EXTEND_RGP
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site:
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.031
Done: true
RequestDateTime: 2/11/2015 1:24:57 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainRRP=E
Extension=successful
DomainName=VeryExpiredName.com
OrderID=156251141
Command=EXTEND_RGP
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.047
Done=true
RequestDateTime=2/11/2015 1:26:19 PM
```
Related Commands
----------------

AddToCart

Extend

ExtendDomainDNS

GetExpiredDomains

GetOrderDetail

GetRenew

PE\_GetResellerPrice

PE\_SetPricing

Purchase

PurchaseServices

SetRenew

UpdateExpiredDomains