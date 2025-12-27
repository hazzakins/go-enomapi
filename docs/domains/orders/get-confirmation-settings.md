GetConfirmationSettings
=======================

Retrieve the settings for email confirmations of orders.

Usage
-----

Use this command to retrieve the current settings for email confirmations of orders \(confirmations sent to you when a customer places an order\).

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/ConfirmationSettings.asp](https://resellertest.enom.com/myaccount/ConfirmationSettings.asp)

The settings on the Reseller Confirmation Email Settings page are retrieved using the GetConfirmationSettings command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                        |
| ---------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| OrderConfirmation | The setting for sending copies of order confirmation emails to you, when orders are placed on this account         |
| TransferOrderConfirmation  | The setting for sending copies of transfer request emails to you, when orders are placed on this account |
| OrderConfirmationBCC | The setting for sending copies of order confirmation emails to you, when orders are placed on a subaccount of this account |
| DomainName          | The domain name |
| TransferOrderConfirmationBCC | The setting for sending copies of transfer request emails to you, when orders are placed on a subaccount of this account  |
| EmailHead          | If set for custom email text, the header string that will be used |
| EmailTail | If set for custom email text, the tail string that will be used                              |
| Command           | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.              |
| ErrX             | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                             |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the settings for order confirmation emails, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetConfirmationSettings&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetConfirmationSettings&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetConfirmationSettings&uid=resellid
&pw=resellpw&responsetype=text
```
In the response, return values in the ConfirmationSettings node confirms that the query was successful:

```
<interface-response>
<ConfirmationSettings>
 <OrderConfirmation>True</OrderConfirmation>
 <TransferOrderConfirmation>True</TransferOrderConfirmation>
 <OrderConfirmationBCC>True</OrderConfirmationBCC>
 <TransferOrderConfirmationBCC>True</TransferOrderConfirmationBCC>
 <EmailHead>テスト</EmailHead>
 <EmailTail>テスト</EmailTail>
</ConfirmationSettings>
<Command>GETCONFIRMATIONSETTINGS</Command>
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
<ExecTime>0.344</ExecTime>
<Done>true</Done>
<RequestDateTime>12/8/2011 3:50:04 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
OrderConfirmation: True
TransferOrderConfirmation: True
OrderConfirmationBCC: True
TransferOrderConfirmationBCC: True
EmailHead: &#12486;&#12473;&#12488;
EmailTail: &#12486;&#12473;&#12488;
Command: GETCONFIRMATIONSETTINGS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod:
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.063
Done: true
RequestDateTime: 2/3/2015 4:53:08 PM
```
Related Commands
----------------

GetAccountInfo

GetCusPreferences

GetSubAccountDetails

GetTLDList

PE\_GetRetailPricing

PE\_SetPricing

UpdateCusPreferences