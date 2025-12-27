UpdateMetaTag
=============

Add or update the HTML metatags for one of a domain’s host records.

Usage
-----

Use this command to raise the probability that a search engine will find this domain.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/metatags.asp?HostID=11415002&DomainNameID=152533676](https://resellertest.enom.com/domains/metatags.asp?HostID=11415002&DomainNameID=152533676)

On the meta tags page, the save changes button calls the UpdateMetaTag command.

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
| Input Parameter | Status | Description                                                | Max Size |
| --------------- | -------------------------------------- | --------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                             | 20 |
| PW | Required                | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                       | 63 |
| TLD | Required                | Top-level domain name \(extension\) | 15    |
| MetaTagHostID  | Required | Metatag host ID, our internal tracking number. You can retrieve host IDs using the GetRegHosts command. | 8 |
| TitleBar | Recommended to optimize search ranking | Content to display in the browser title bar. Replace spaces with \+. | 100   |
| SiteDescription | Recommended to optimize search ranking | Content to display in search engine results. Replace spaces with \+.                  | 250 |
| Keywords | Recommended to optimize search ranking | List of keywords for search engines. Separate keywords with \+. | 250   |
| ResponseType  | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML.                           | 4 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| TitleBar | New title to display in browser title bar                            |
| SiteDescription | New description to display in search engine results |
| Keywords | New search-engine keywords for this host ID                           |
| DomainNameID   | 9-digit domain name ID, our internal tracking number |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sets values for the metatags, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=UpdateMetaTag&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&MetaTagHostID=11415002
&TitleBar=MY+COOL+TITLE&SiteDescription=MY+SITE+DESCRIPTION
&Keywords=MY+SITE+KEYWORDS&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateMetaTag&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&MetaTagHostID=11415002
&TitleBar=MY+COOL+TITLE&SiteDescription=MY+SITE+DESCRIPTION
&Keywords=MY+SITE+KEYWORDS&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=UpdateMetaTag&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&MetaTagHostID=11415002
&TitleBar=MY+COOL+TITLE&SiteDescription=MY+SITE+DESCRIPTION
&Keywords=MY+SITE+KEYWORDS&ResponseType=text
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
 <Command>UPDATEMETATAG</Command>
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
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
DomainRRP: E
titlebar: MY COOL TITLE
sitedescription: MY SITE DESCRIPTION
keywords: MY SITE KEYWORDS
DomainNameID: 152533676
Command: UPDATEMETATAG
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site:
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.156
Done: true
RequestDateTime: 2/9/2015 2:15:28 PM
```
Related Commands
----------------

GetHosts

GetMetaTag

GetRegHosts

SetHosts