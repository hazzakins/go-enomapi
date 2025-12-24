SetResellerTLDPricing
=====================

Set the prices you charge your resellers for domain names. You can use this command to set any number of prices.

Usage
-----

Use this command to set prices you charge your resellers for domain names.

We recommend that you use SetResellerTLDPricing for setting prices charged to subaccounts. The GetSubAccountDetails command also offers this functionality, but with much stricter requirements on the query string.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=107-vm-2729](https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=107-vm-2729)

The save changes button calls the SetResellerTLDPricing command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The subaccount must belong to this domain name account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                       | Max Size |
| --------------- | ------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                     | 20 |
| PW | Required               | Account password | 20    |
| SubUID     | Either SubUID or Account is Required | Login ID of the subaccount to set pricing for.                                                     | 20 |
| Account | Either SubUID or Account is Required | Account ID number of the subaccount to set pricing for, in NNN-aa-NNNN format. To retrieve the subaccount ID number, use the GetSubAccounts command. | 11    |
| TLDPrice    | Optional | Price you are charging this subaccount for registering this TLD, in DD.cc format                                    | 5 |
| TLDRenew | Optional               | Price you are charging this subaccount for renewing this TLD, in DD.cc format | 5    |
| TLDTransfer   | Optional | Price you are charging this subaccount for transferring this TLD, in DD.cc format                                    | 5 |
| DotNameBundle | Optional               | Price you are charging this subaccount for a name-and-email .name bundle, in DD.cc format | 5    |
| POP3      | Required | Price you are charging this subaccount for a POP3 10-pak, in DD.cc format                                       | 5 |
| WPPS | Required               | Price you are charging this subaccount for WhoIs Privacy Protection Service, in DD.cc format | 5    |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                                                  | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                                        |
| ----------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| price tld="TLD" id="ID" prod="Prod" | When ResponseType=XML, this multi-part parameter name is used as the XML tag. The return value is the price of the product described by the param name. |
| tldX TLD.             | Indexed X when ResponseType=Text or HTML. |
| tldIDX | TLD ID number from our database. Indexed X when ResponseType=Text or HTML.                                        |
| productIDX             | Product ID number from our database. Indexed X when ResponseType=Text or HTML. |
| priceX | Product price. Indexed X when ResponseType=Text or HTML.                                                 |
| DotNameBundle            | The price of a one-year subscription to a .name name-and-email bundle |
| NameX | Name of product. Indexed X when ResponseType=Text or HTML.                                                |
| PriceX               | The price for a one-year subscription to product. Indexed X when ResponseType=Text or HTML. |
| Command | Name of command executed                                                                  |
| ErrCount              | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                 |
| Done                | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sets prices for some TLDs and services in account resellid, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SETRESELLERTLDPRICING&uid=resellid
&pw=resellpw&account=443-up-6579&comprice=12.95
&comrenew=12.95&comtransfer=12.95&netprice=12.95
&netrenew=12.95&nettransfer=12.95&orgprice=8.95
&orgrenew=8.95&orgtransfer=8.95&wpps=8.50&pop3=19.95
&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=SETRESELLERTLDPRICING&uid=resellid
&pw=resellpw&account=443-up-6579&comprice=12.95
&comrenew=12.95&comtransfer=12.95&netprice=12.95
&netrenew=12.95&nettransfer=12.95&orgprice=8.95
&orgrenew=8.95&orgtransfer=8.95&wpps=8.50&pop3=19.95
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SETRESELLERTLDPRICING&uid=resellid
&pw=resellpw&account=443-up-6579&comprice=12.95
&comrenew=12.95&comtransfer=12.95&netprice=12.95
&netrenew=12.95&nettransfer=12.95&orgprice=8.95
&orgrenew=8.95&orgtransfer=8.95&wpps=8.50&pop3=19.95
&responsetype=text
```
In the response, the list of TLDs and services and their prices, and the ErrCount value 0, confirm that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
<tlds>
 <price tld="biz" id="6" prod="10" />
 <price tld="biz" id="6" prod="16" />
 <price tld="biz" id="6" prod="19" />
 <price tld="com" id="0" prod="10">12.95</price>
 <price tld="com" id="0" prod="16">12.95</price>
 <price tld="com" id="0" prod="19">12.95</price>
 <price tld="info" id="5" prod="10" />
 <price tld="info" id="5" prod="16" />
 <price tld="info" id="5" prod="19" />
 <price tld="net" id="1" prod="10">12.95</price>
 <price tld="net" id="1" prod="16">12.95</price>
 <price tld="net" id="1" prod="19">12.95</price>
 <price tld="org" id="2" prod="10">12.95</price>
 <price tld="org" id="2" prod="16">12.95</price>
 <price tld="org" id="2" prod="19">12.95</price>
</tlds>
<tldcount>15</tldcount>
<services>
 <service name="dotnamebundle" price="" />
</services>
<servicescount>3</servicescount>
<Command>SETRESELLERTLDPRICING</Command>
<Language>en</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>RESELLERTEST</Server>
<Site />
<IsLockable>False</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<ExecTime>1.492188</ExecTime>
<Done>true</Done>
<debug>
 <![CDATA[ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
tld: biz
id: 6
prod: 10
price:
tld: biz
id: 6
prod: 16
price:
tld: biz
id: 6
prod: 19
price:
tld: com
id: 0
prod: 10
price: 12.95
tld: com
id: 0
prod: 16
price: 12.95
tld: com
id: 0
prod: 19
price: 12.95
.
.
.
Command: SETRESELLERTLDPRICING
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 1.250
Done: true
TrackingKey: 683f858d-d0c5-4e0f-a41c-fd71e7301079
RequestDateTime: 2/12/2015 2:00:34 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
tld=biz
id=6
prod=10
price=
tld=biz
id=6
prod=16
price=
tld=biz
id=6
prod=19
price=
tld=com
id=0
prod=10
price=12.95
tld=com
id=0
prod=16
price=12.95
tld=com
id=0
prod=19
price=12.95
.
.
.
Command=SETRESELLERTLDPRICING
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
TrackingKey=8392668f-f734-42b0-b45a-605b4e3a1f87
RequestDateTime=2/12/2015 2:07:05 PM
```
Related Commands
----------------

AuthorizeTLD

GetSubAccountDetails

GetSubAccounts

GetTLDList

PE\_GetTLDID

PE\_SetPricing

RemoveTLD

SetResellerServicesPricing