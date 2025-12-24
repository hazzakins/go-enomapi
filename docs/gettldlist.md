GetTLDList
==========

Retrieve a list of the TLDs that you offer.

Usage
-----

Use this command to retrieve a list of the TLDs that you offer. If you have not specifically authorized TLDs then this command returns all TLDs offered by us.

Use the AuthorizeTLD command to authorize TLDs and the RemoveTLD command to remove TLDs from authorized status for your account.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/Settings.asp](https://resellertest.enom.com/myaccount/Settings.asp)

Clicking the TLD pricing tab shows a table that includes all TLDs offered by this account \(although the contents of this particular table are retrieved by a different command than GetTLDList\).

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                               | Max Size |
| --------------- | -------- | ------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                             | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                            |
| ---------------- | ------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.     |
| Done       | True indicates this entire response has reached you successfully. |
| TLDX | Top-level domain name \(extension\). If ResponseType=text or html, each entry is indexed X.  |
| TLDCount     | Number of TLDs listed in this response \(that is, the number of TLDs authorized by this account\) |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves a list of TLDs authorized for this account, and sends the response in XML format:

```
https://resellertest.enom.com/interface.asp?
command=gettldlist&uid=resellid&pw=resellpw
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=gettldlist&uid=resellid&pw=resellpw
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=gettldlist&uid=resellid&pw=resellpw
&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>

  <tld>com</tld>

 </tld>

  <tld>net</tld>

  <tld>org</tld>

  <tld>us</tld>

  <tld>info</tld>

  <tld>biz</tld>

  <tld>co.uk</tld>

  <tld>org.uk</tld>

  <tld>de</tld>

 <tldcount>9</tldcount>

 </tldlist>

 <Command>GETTLDLIST</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site />

 <IsLockable />

 <IsRealTimeTLD />

 <ExecTime>0.140625</ExecTime>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

tld1: am

tld2: asia

tld3: biz

tld4: cc

tld5: cm

tld6: co

tld7: co.uk

tld8: com

tld9: com.au

tld10: fr

tld11: in

tld12: info

tld13: me

tld14: mobi

tld15: net

tld16: nu

tld17: org

tld18: pl

tld19: tel

tld20: tv

tld21: us.com

tld22:

tldcount: 22

Command: GETTLDLIST

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.078

Done: true

RequestDateTime: 2/4/2015 12:53:51 PM
```
ExecTime=0.094

RequestDateTime=2/4/2015 12:54:16 PM
```