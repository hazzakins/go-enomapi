SetResellerServicesPricing
==========================

Set the prices you charge your reseller for services.

Usage
-----

Use this command to set prices you charge a subaccount for services.

This command differs from the GetSubAccountDetails command in that the SetResellerServicesPricing command allows you to set pricing of any combination of services, SetResellerServicesPricing does not require you to simultaneously convert an account to reseller status, and SetResellerServicesPricing does not require you to set all service prices in a single query.

To cover all the functionality of the GetSubAccountDetails command without its restrictions, use SetResellerTLDPricing, and SetResellerServicesPricing.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=107-vm-2729](https://resellertest.enom.com/resellers/SubAccount-Manage.asp?Account=107-vm-2729)

The save changes button calls the SetResellerServicesPricing command.

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
| Input Parameter | Status | Description                                             | Max Size |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                          | 20 |
| PW | Required                                                                  | Account password | 20    |
| SubUID     | Either SubUID or Account is Required | Login ID of the subaccount to set pricing for.                           | 20 |
| Account | Account ID of the subaccount to set pricing for, in NNN-aa-NNNN format. To retrieve the subaccount ID, use the GetSubAccounts command. | Top-level domain name \(extension\) of name in RGP status | 11    |
| DotNameBundle  | Optional | Price you are charging this subaccount for a name-and-email .name bundle, in DD.cc format    | 5 |
| DNSHosting | Optional                                                                  | Price you are charging this subaccount for one year ofDNS hosting, in DD.cc format | 5    |
| WPPS      | Optional | Price your are charging this subaccount for a one-yearsubscription to ID Protect, in DD.cc format | 5 |
| RichContent | Optional                                                                  | Price you are charging this subaccount for a monthlyRichContent subscription, in DD.cc format | 5    |
| POP3      | Optional | Price you are charging this subaccount for a POP310-pak, in DD.cc format             | 5 |
| ResponseType | Optional Format of response.                                                       | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------------------------ |
| ServiceX     | Name of this service. Indexed X when ResponseType=Text or HTML. |
| PriceX | Price you set for this service. Indexed X when ResponseType=Text or HTML.                    |
| ServicesCount  | Number of services listed in this response \(all services, not just those for which you set prices in your query\) |
| Command | Name of command executed                                              |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.             |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example
-------

The following query sets prices for services-wholesale prices that resellid charges subaccount 443-up-6579, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=SETRESELLERSERVICESPRICING&uid=resellid
&pw=resellpw&account=443-up-6579&pop3=19.95
&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=SETRESELLERSERVICESPRICING&uid=resellid
&pw=resellpw&account=443-up-6579&pop3=19.95
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SETRESELLERSERVICESPRICING&uid=resellid
&pw=resellpw&account=443-up-6579&pop3=19.95
&responsetype=text
```
In the response, the list of prices that match those you set, and the ErrCount value 0, indicate that the query was successful:

```
<?xml version="1.0" ?>
<interface-response>
<services>
 <service name="pop3" price="19.95" />
 <service name="dotnamebundle" price="" />
</services>
<servicescount>2</servicescount>
<Command>SETRESELLERSERVICESPRICING</Command>
<Language>en</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>ResellerTest</Server>
<Site>enom</Site>
<IsLockable>0</IsLockable>
<IsRealTimeTLD>0</IsRealTimeTLD>
<ExecTime>0.6875</ExecTime>
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
servicename1: pop3
price: 19.95
servicename2: dotnamebundle
price:
servicescount: 2
Command: SETRESELLERSERVICESPRICING
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
ExecTime: 1.031
Done: true
TrackingKey: 2d41bf88-b052-41ba-9861-181406e55e62
RequestDateTime: 2/5/2015 3:25:48 PM
```
Related Commands
----------------

GetSubAccounts

PE\_SetPricing

SetResellerTLDPricing