PurchasePOPBundle
=================

Purchase, in real time, a pak of 10 POP Email accounts for the specified domain name.

Usage
-----

Use this command to purchase, in real time, a pak of 10 POP Email accounts for the specified domain name.

We recommend the POP3 mail service for parties who want to be able to send mail from their domain name \(instead of, for example, a yahoo or hotmail sender’s address\), and for parties who want to manage multiple email accounts as a body \(rather than having everyone’s email forwarded to scattered locations\).

After you sell a POP Email bundle, you can set up users with the SetUpPOP3User command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/POPConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/POPConfig.asp?DomainNameID=152533676)

On the POP3 Mail page, the add an Email Pak button calls the PurchasePOPBundle command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- To use our credit card processing, this must be an ETP reseller account.
- The domain name must belong to this account.
- This command can be used only for purchasing POP paks for reseller accounts. Retail accounts must
- use a queue-based shopping cart process.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                         | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                       | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| Quantity    | Required | Number of 10-address paks.                 | 10 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| OrderID     | Identification number of the order |
| BundleID | Identification number of each POP3 10-pak                            |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests 1 pak of 10 POP3 mailboxes for resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=PurchasePOPBundle&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&Quantity=1&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=PurchasePOPBundle&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&Quantity=1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=PurchasePOPBundle&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&Quantity=1&responsetype=text
```
The response is as follows:

```
<interface-response>
 <orderid>157781024</orderid>
 <Bundles>
 <BundleID>30823</BundleID>
 </Bundles>
 <BundleCount>1</BundleCount>
 <Command>PURCHASEPOPBUNDLE</Command>
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
 <ExecTime>1.265</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/9/2011 4:02:14 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
orderid: 161808176
BundleID1: 1004952
BundleCount: 1
Command: PURCHASEPOPBUNDLE
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T1
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 11.848
Done: true
RequestDateTime: 2/4/2015 5:50:18 PM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

Forwarding

GetDotNameForwarding

GetForwarding

GetMailHosts

GetPOP3

GetPOPExpirations

GetPOPForwarding

ModifyPOP3

PurchaseHosting

PurchaseServices

RenewPOPBundle

SetDotNameForwarding

SetPOPForwarding

SetUpPOP3User