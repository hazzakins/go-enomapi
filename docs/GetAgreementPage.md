GetAgreementPage
================

Retrieve the up-to-date Registration Agreement.

Usage
-----

Use this command to retrieve the up-to-date verbiage for the requested agreement. You can implement this command in your own Web site so that your agreement page always contains the current language.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetAgreementPage&uid=(Required)&pw=(Required)&responsetype=(Optional)
```
| Input Parameter | Type | Status             | Description |
| --------------- | ------ | ------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required            | GetBalance |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| pw       | string | Required            | Your API Token |
| Page | string | Optional; default is Agreement | Which agreement to retrieve. Permitted values are: ----General---- Agreement General Registration Agreement Anti-Malware GeoTrust Anti-Malware Agreement DRP Uniform Domain Name Dispute Policy ID-Protect ID-Protect Agreement goMobi additional requirements for goMobi POP-Email POP E-Mail Agreement TLDPortal TLDPortal Agreement TrustSeal Symantec Safe Site \(aka Verisign Trust Seal\) Service Agreement Watchlist TLD Watchlist Agreement WBL\_Terms Business Listing Agreement ----TLD---- AC additional requirements for .ac Altdotcom additional Registry requirements for CentralNIC's TLDs \(such as .br.com and .se.net\) AM additional requirements for .am AT additional requirements for .at AU additional requirements for .au BE additional requirements for .be BIZ additional requirements for .biz BZ additional requirements for .bz CA additional requirements for .ca CC additional requirements for .cc CN additional requirements for .cn CO additional requirements for .co DE additional requirements for .de ES additional requirements for .es EU additional requirements for .eu GS additional requirements for .gs IN additional requirements for .in INFO additional requirements for .info IO additional requirements for .io IT additional requirements for .it JOBS additional requirements for .jobs JP additional requirements for .jp LA additional requirements for .la MOBI additional requirements for .mobi MS additional requirements for .ms MX additional requirements for .mx NAME additional requirements for .name NL additional requirements for .nl NU additional requirements for .nu NZ additional requirements for .nz PE additional requirements for .pe PRO additional requirements for .pro SH additional requirements for .sh TC additional requirements for .tc TLDTW additional requirements for .tw 3rd level domain TM additional requirements for .tm TV additional requirements for .tv TW additional requirements for .tw UK additional requirements for .uk US additional requirements for .us VG additional requirements for .vg WS additional requirements for .ws XXX additional requirements for .xxx |
| Language    | string | Optional; default is Eng    | Language of agreement. Permitted values are: Eng English Ger German Por Portuguese Spa Spanish |
| ResponseType | string | Optional | Format of response. Permitted values are - Text \(default\) - HTML - XML                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

For RichContent, display the user agreement found at [http://www.richcontentwidget.com/terms.aspx](http://www.richcontentwidget.com/terms.aspx)

For Comodo SSL certificates, display the user agreement that link from [http://www.comodo.com/about/comodo-agreements.php](http://www.comodo.com/about/comodo-agreements.php)

For GeoTrust SSL certificates, display the user agreement found at [http://www.geotrust.com/resources/cps/pdfs/gt\_ssl\_SA\_v.2.0.pdf](http://www.geotrust.com/resources/cps/pdfs/gt_ssl_SA_v.2.0.pdf)

For VeriSign SSL certificates, display the user agreement found at [http://www.verisign.com/repository/agreements/serverClass3Org.html](http://www.verisign.com/repository/agreements/serverClass3Org.md)

| Output Parameter | Type | Description                                           |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command     | string | Name of command executed                                     |
| Content     | int | Content of the Registration Agreement                              |
| ErrCount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | boolian | "True" indicates this entire response has reached you successfully.               |

Example Output
--------------

The following query requests the account balance for account Login Id "resellid". The query also requests the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=GetAgreementPage&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetAgreementPage&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetAgreementPage&uid=resellid
&pw=resellpw&responsetype=text
```
```html
<interface-response>
 <content>
 <p>THIS AGREEMENT HAS A PROVISION FOR ARBITRATION OF DISPUTES
BETWEEN THE PARTIES.</p>
 <p>
This Registration Agreement ("Agreement") sets forth the terms
and conditions of your use of domain name registration and
related services ("Services"). In this Agreement "you" and
"your" refer to you and the registrant listed in the WHOIS
contact information for the domain name. "We", "us" and
"our" refer to the registrars listed at the <a href=
"#registrars">bottom of this document</a>, any one of which
will be the registrar for your domain name and all of which
share common ownership, common terms and conditions, and a
shared Services infrastructure. To determine which registrar
your domain name is registered with, perform a WHOIS lookup
at <a href="http://www.uwhois.com" target="_blank">http://
www.uwhois.com</a>. You obtain the Services through your
primary service provider, with whom we have a wholesale
relationship (your "Primary Service Provider"). Your
relationship with your Primary Service Provider may be
governed by additional terms, as you and your Primary Service
Provider may agree. "We," "us" and "our" does not include your
Primary Service Provider, except when specifically mentioned
or unless your Primary Service Provider is one of us (i.e.,
if your Primary Service Provider is also one of the registrars
listed at the <a href="#registrars">bottom of this document
</a>)).
 </p>
.
.
.

 </content>
 <Command>GETAGREEMENTPAGE</Command>
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
 <ExecTime>0.141</ExecTime>
 <Done>true</Done>
 <debug/>
 <TrackingKey>b6812a69-5c50-43dd-85db-aeb3f37ff497</TrackingKey>
 <RequestDateTime>12/8/2011 3:31:10 AM</RequestDateTime>
</interface-response>
```
```html
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>page: </STRONG>agreement<br />
<STRONG>content: </STRONG><p>THIS AGREEMENT HAS A PROVISION FOR ARBITRATION OF DISPUTES BETWEEN THE PARTIES.</p> <p>This Registration Agreement ("Agreement") sets forth the terms and conditions of your use of domain name registration and related services ("Services"). In this Agreement "you" and "your" refer to you and the registrant listed in the WHOIS contact information for the domain name. "We", "us" and "our" refer to the registrars listed at the bottom of this document, any one of which will be the registrar for your domain name and all of which share common ownership, common terms and conditions, and a shared Services infrastructure. To determine which registrar your domain name is registered with, perform a WHOIS lookup at <a href="http://www.uwhois.com" target="_blank">http://www.uwhois.com</a>. You obtain the Services through your primary service provider, with whom we have a wholesale relationship (your "Primary Service Provider"). Your relationship with your Primary Service Provider may be governed by additional terms, as you and your Primary Service Provider may agree. "We," "us" and "our" does not include your Primary Service Provider, except when specifically mentioned or unless your Primary Service Provider is one of us (i.e., if your Primary Service Provider is also one of the registrars listed at the <a href="#registrars">bottom of this document</a>)). </p> <ol> <!--1--> <li>YOUR AGREEMENT:
.
.
.<br />

<STRONG>Command: </STRONG>GETAGREEMENTPAGE<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t1<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable:</STRONG><br />
<STRONG>IsRealTimeTLD:</STRONG><br />
<STRONG>TimeDifference: </STRONG>+0.00<br />
<STRONG>ExecTime: </STRONG>0.000<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>7057e4d1-bad5-4acb-b4e1-84cf38f26d0a<br />
<STRONG>RequestDateTime: </STRONG>2/3/2015 4:16:46 PM<br />
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
page=agreement
content=%0d%0a%3cp%3eTHIS+AGREEMENT+HAS+A+PROVISION+FOR+ARBITRATION+OF+DISPUTES+BETWEEN+THE+PARTIES.%3c%2fp%3e%0d%0a%0d%0a%3cp%3eThis+Registration+Agreement+(%22Agreement%22)+sets+forth+the+terms+and+conditions+of+your+use+of+domain+name+registration+and+related+services+(%22Services%22).+In+this+Agreement+%22you%22+and+%22your%22+refer+to+you+and+the+registrant+listed+in+the+WHOIS+contact+information+for+the+domain+name.+%22We%22%2c+%22us%22+and+%22our%22+refer+to+the+registrars+listed+at+the+bottom+of+this+document%2c+any+one+of+which+will+be+the+registrar+for+your+domain+name+and+all+of+which+share+common+ownership%2c+common+terms+and+conditions%2c+and+a+shared+Services+infrastructure.+To+determine+which+registrar+your+domain+name+is+registered+with%2c+perform+a+WHOIS+lookup+at+%3ca+href%3d%22http%3a%2f%2fwww.uwhois.com%22+target%3d%22_blank%22%3ehttp%3a%2f%2fwww.uwhois.com%3c%2fa%3e.+You+obtain+the+Services+through+your+primary+service+provider%2c+with+whom+we+have+a+wholesale+relationship+(your+%22Primary+Service+Provider%22).+Your+relationship+with+your+Primary+Service+Provider+may+be+governed+by+additional+terms%2c+as+you+and+your+Primary+Service+Provider+may+agree.+%22We%2c%22+%22us%22+and+%22our%22+does+not+include+your+Primary+Service+Provider%2c+except+when+specifically+mentioned+or+unless+your+Primary+Service+Provider+is+one+of+us+(i.e.%2c+if+your+Primary+Service+Provider+is+also+one+of+the+registrars+listed+at+the+%3ca+href%3d%22%23registrars%22%3ebottom+of+this+document%3c%2fa%3e)).+%3c%2fp%3e%0d%0a%0d%0a%3col%3e%0d%0a++++%3c!--1--%3e%0d%0a++++%3cli%3eYOUR+AGREEMENT%3a
.
.
.
Command=GETAGREEMENTPAGE
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t1
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.000
Done=true
TrackingKey=62a87f31-d33f-441b-95a8-217866559332
RequestDateTime=2/3/2015 4:19:27 PM
```
Related Commands