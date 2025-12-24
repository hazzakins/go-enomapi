GetPOPExpirations
=================

Retrieve a list of POP paks and their expiration dates, and the domains they are associated with.

Usage
-----

Use this command to retrieve a list of POP paks and their expiration dates, and the domains they are associated with.

This command allows you to synchronize your database with ours. In particular, you can use it to identify POP paks that are about to expire or have recently expired, update the expiration dates of POP paks that were renewed in a site other than yours, or retrieve the expiration date of POP paks associated with domains that were transferred or pushed into your account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?DomainNameID=152533676)

In the email settings section, the configure button returns expiration dates for POP mail paks—one component of the information retrieved by the GetPOPExpirations command.

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

| Output Parameter | Description                                                                                        |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| DomainX | Domain name with which this POP pak is associated. The response is indexed X if ResponseType=text or HTML.                                        |
| BundleIDX    | ID number of this POP pak, from our internal records. Use this ID number to renew the POP pak using the RenewPOPBundle command. The response is indexed X if ResponseType=text or HTML. |
| ExpDateX | Expiration date of this POP pak. The response is indexed X if ResponseType=text or HTML.                                                 |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                              |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                                             |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves POP pak expiration dates, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetPOPExpirations&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetPOPExpirations&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetPOPExpirations&uid=resellid
&pw=resellpw&responsetype=text
```
The response is as follows:

```
<interface-response>
 <popexpirations>
 <pop>
  <domain>1986anotherappz.us</domain>
  <bundleid>30631</bundleid>
  <expdate>10/25/2012 2:13:57 AM</expdate>
 </pop>
 <pop>
  <domain>1986anotherappz.us</domain>
  <bundleid>30630</bundleid>
  <expdate>10/25/2012 2:12:18 AM</expdate>
 </pop>
.
.
.
 <count>125</count>
 </popexpirations>
 <Command>GETPOPEXPIRATIONS</Command>
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
 <ExecTime>1.234</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>0a7ee088-86c5-4067-885c-d141144e8ae7</TrackingKey>
 <RequestDateTime>12/9/2011 2:21:48 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
domain1: 123example.net
bundleid1: 1004277
expdate1: 11/25/2015 11:09:13 AM
renew1: False
version1: 2
domain2: 45566.net
bundleid2: 1004278
expdate2: 11/25/2015 11:11:07 AM
renew2: False
version2: 2
domain3: 0592lh.com
bundleid3: 32589
expdate3: 10/12/2019 11:01:45 AM
renew3: False
version3: 2
domain4: cl-test.com
bundleid4: 32820
expdate4: 10/12/2019 11:01:45 AM
renew4: False
version4: 2
.
.
.
count: 259
success: True
Command: GETPOPEXPIRATIONS
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
ExecTime: 0.547
Done: true
TrackingKey: 15b60a1b-e6f4-49e9-8e2e-b6416d5e4d96
RequestDateTime: 2/4/2015 11:18:40 AM
```
Related Commands
----------------

Extend

GetPOP3

GetRenew

Purchase

PurchasePOPBundle

RenewPOPBundle