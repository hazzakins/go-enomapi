GetMetaTag
==========

Retrieve the metatags for one of a domain’s host records.

Usage
-----

Use this command to retrieve the current metatags for one host record for a domain.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/HostConfig.asp?DomainNameID=152533676](https://resellertest.enom.com/domains/HostConfig.asp?DomainNameID=152533676)

After you assign to a host record a record type of URL Frame, then click save changes, then in the DNS Settings \(zone file\) section click configure, the edit link calls the GetMetaTag command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                          | Max Size |
| --------------- | ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                        | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                                                  | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| MetaTagHostID  | Required | Host record ID for this host record for this domain; our internal tracking number. You can retrieve all host IDs for a domain using the GetRegHosts command. | 8 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| TitleBar     | Metatag content for browser title bar |
| SiteDescription | Metatag content for search engine results                            |
| Keywords     | Metatag content for search engine keywords |
| DomainNameID | ID number for this domain, our internal tracking number                     |
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

The following query retrieves the HTML metatag contents for the specified host record, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getmetatag&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&MetaTagHostID=11415002
&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=getmetatag&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&MetaTagHostID=11415002
&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=getmetatag&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&MetaTagHostID=11415002
&ResponseType=text
```
The response is as follows:

```
<?xml version="1.0" ?>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <metatags>
 <titlebar>MY COOL TITLE</titlebar>
 <sitedescription>MY SITE DESCRIPTION</sitedescription>
 <keywords>MY SITE KEYWORDS</keywords>
 <DomainNameID>152533676</DomainNameID>
 </metatags>
 <Command>GETMETATAG</Command>
 <ErrCount>0</ErrCount>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable>False</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
 <debug>
 <![CDATA[ ] ]>
 </debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
DomainRRP: E
titlebar: MY COOL TITLE
sitedescription: MY SITE DESCRIPTION
keywords: MY SITE KEYWORDS
DomainNameID: 152533676
Command: GETMETATAG
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
ExecTime: 0.141
Done: true
RequestDateTime: 2/11/2015 3:02:50 PM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainRRP=E
titlebar=MY COOL TITLE
sitedescription=MY SITE DESCRIPTION
keywords=MY SITE KEYWORDS
DomainNameID=152533676
Command=GETMETATAG
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
ExecTime=0.172
Done=true
RequestDateTime=2/11/2015 3:04:36 PM
```
Related Commands
----------------

GetHosts

GetRegHosts

SetHosts

UpdateMetaTag