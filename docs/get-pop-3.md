GetPOP3
=======

Gets all POP3 accounts for a domain name.

Usage
-----

Use this command to display email forwarding records for a domain name.

Use this command to generate a list of the POP mail accounts belonging to a domain name. Return information includes mailbox names and BundleID numbers for the domain’s POP paks.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

In the domain control panel, email settings box, if the current service is POP mail, clicking the configure button calls the GetPOP3 command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must use eNom’s domain name servers.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                               | Max Size |
| --------------- | ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                             | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\) for the domain this service is associated with               | 63 |
| TLD | Required           | Top-level domain name \(extension\) for the domain this service is associated with | 15    |
| GetDefaultOnly | Optional | If this is set to 1, extra blank input records are returned for user input on the email forwarding form. Primarily used in XML output. | 1 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter   | Description |
| -------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| UserNameX      | POP user name for this email account \(this mailbox\). Indexed X when ResponseType=Text or HTML. |
| EmailCount | Number of email accounts on the domain name.                                                            |
| PasswordX      | POP password name for the email account. Indexed X = 1 to EmailCount if ResponseType=text or HTML. |
| QuotaX | Maximum storage capacity for each mailbox, in megabytes. Indexed X = 1 to EmailCount if ResponseType=text or HTML.                         |
| ExpDateX       | Expiration date of the POP3 account. Indexed X = 1 to EmailCount if ResponseType=text or HTML. |
| BundleIDX | ID number of this POP pak. Indexed X = 1 to EmailCount if ResponseType=text or HTML.                                        |
| AutoRenewX      | Auto-renewal setting for this POP3 account and its POP pak. 1 signifies that auto-renew is on, 0 off. Indexed X = 1 to EmailCount if ResponseType=text or HTML. |
| QuotaUsedX | Number of MB of storage space currently in use for this account. Indexed X =1 to EmailCount if ResponseType=text or HTML.                     |
| \[Pak\]BundleIDX   | ID number of this POP pak. Prefixed Pak and indexed X = 1 to Count if ResponseType=text or HTML. |
| \[Pak\]QtyPurchasedX | Total number of POP accounts that have been purchased for this domain name. Prefixed xed X = 1 to Count if ResponseType=text or HTML.               |
| \[Pak\]QtyAvailableX | Total number of POP accounts currently configured for this domain name. Prefixed xed Pak and indexed X = 1 to Count if ResponseType=text or HTML. |
| \[Pak\]ExpDateX | Expiration date of this POP pak. Prefixed Pak and indexed X = 1 to Count if ResponseType=text or HTML.                               |
| \[Pak\]AutoRenewX  | Auto-renewal setting for this POP pak. 1 signifies that auto-renew is on, 0 off. Prefixed Pak and indexed X = 1 to Count if ResponseType=text or HTML. |
| \[Pak\]QuotaX | Maximum storage capacity for each mailbox, in megabytes. Prefixed Pak and indexed X = 1 to Count if ResponseType=text or HTML.                   |
| Command       | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                   |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                                  |

Notes
-----
- The default response format is plain textThe default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter eater than 0 the transaction failed If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests a list of all POP mail accounts in resellerdocs.com, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getpop3&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getpop3&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getpop3&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&responsetype=text
```
The response is as follows:

```html
<interface-response>
 <pop>
 <username>dell</username>
 <password>78910</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:14:59 AM</expdate>
 <BundleId>26770</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>hp</username>
 <password>78910</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:14:59 AM</expdate>
 <BundleId>26770</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>lenovo</username>
 <password>1111</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:14:59 AM</expdate>
 <BundleId>26770</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>sss</username>
 <password>aaaaaaa</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:14:59 AM</expdate>
 <BundleId>26770</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>facebook</username>
 <password>78910</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:15:51 AM</expdate>
 <BundleId>26771</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>google</username>
 <password>78910</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:15:51 AM</expdate>
 <BundleId>26771</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>instragram</username>
 <password>1111</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:15:51 AM</expdate>
 <BundleId>26771</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <pop>
 <username>twitter</username>
 <password>1111</password>
 <quota>1024</quota>
 <expdate>5/3/2010 2:15:51 AM</expdate>
 <BundleId>26771</BundleId>
 <AutoRenew>0</AutoRenew>
 <admin>0</admin>
 <quotaused>0</quotaused>
 </pop>
 <EmailCount>8</EmailCount>
 <Paks>
 <Pak>
  <BundleId>26770</BundleId>
  <QtyPurchased>10</QtyPurchased>
  <QtyAvailable>6</QtyAvailable>
  <ExpDate>5/3/2010 2:14:59 AM</ExpDate>
  <AutoRenew>0</AutoRenew>
  <Quota>1024</Quota>
  <ProdType>38</ProdType>
 </Pak>
 <Pak>
  <BundleId>26771</BundleId>
  <QtyPurchased>10</QtyPurchased>
  <QtyAvailable>6</QtyAvailable>
  <ExpDate>5/3/2010 2:15:51 AM</ExpDate>
  <AutoRenew>0</AutoRenew>
  <Quota>1024</Quota>
  <ProdType>38</ProdType>
 </Pak>
 </Paks>
 <Count>2</Count>
 <AdditionalHostPop></AdditionalHostPop>
 <AdditionalHostAcctLogin/>
 <Successful>True</Successful>
 <Command>GETPOP3</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>sjl21wresellt01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+8.00</TimeDifference>
 <ExecTime>0.047</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>9940727e-36fd-4b47-92f8-732af897c07a</TrackingKey>
 <RequestDateTime>12/8/2011 5:14:56 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
Username1: neil
Password1: Pass1w0rd
Quota1: 1024
ExpDate1: 11/6/2015
BundleId1: 1003774
AutoRenew1: 0
admin1: 1
quotaused1: 0
vasitemid1: 1003774
version1: 2
EmailCount: 1
PakBundleId1: 1003774
PakQtyPurchased1: 10
PakQtyAvailable1: 9
PakExpDate1: 11/6/2015
PakAutoRenew1: 0
PakQuota1: 1024
PakProdType1: 38
PakVersion1: 2
Count: 1
AdditionalHostAcctLogin:
Successful: True
Success: True
Command: GETPOP3
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.219
Done: true
TrackingKey: 3eb23499-9163-4321-9aa4-2bfb721189db
RequestDateTime: 2/4/2015 11:16:11 AM
```
Related Commands
----------------

DeleteAllPOPPaks

DeletePOP3

DeletePOPPak

Forwarding

GetCatchAll

GetForwarding

GetMailHosts

GetPOPExpirations

GetPOPForwarding

ModifyPOP3

PurchasePOPBundle

RenewPOPBundle

SetPakRenew

SetPOPForwarding

SetUpPOP3User