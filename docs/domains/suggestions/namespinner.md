NameSpinner
===========

Generate variations of a domain name that you specify, for .com, .net, .tv, and .cc TLDs.

Usage
-----

Use this command to generate variations of the domain name you specify. This command will return TLDs .com, .net, .tv, and .cc.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?command=NameSpinner&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&TLDList={Required}&SensativeContent={Optional}&MaxLength={Optional}&MaxResults={Optional}&UseHyphens={Optional}&UseNumbers={Optional}&Basic={Optional}&Related={Optional}&Similar={Optional}&Topical={Optional}&responsetype={Optional}
```
| Input Parameter | Type | Status  | Description |
| ---------------- | ------ | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | Required | NameSpinner |
| uid | string | Required | Your Account ID                                                                                                                                                  |
| pw        | string | Required | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                                             |
| TLD       | string | Required | Top-level domain name \(extension\) |
| TLDList | string | Required | List of top-level domain names \(extensions\)                                                                                                                                   |
| SensitiveContent | string | Optional | Block potentially offensive content. Default value is "False". "True" filters out offensive content. Permitted values are: - True - False |
| MaxLength | int  | Optional | Maximum length of SLD to return. Permitted values are numbers 2 to 63. Default value is: 63                                                                                                           |
| MaxResults    | int | Optional | Maximum number of suggested names to return in addition to your input. The number actually returned may be lower because of the spinning algorithm developed by the Registry. Permitted values are numbers 0 to 99. Default value is: 20 |
| UseHyphens | string | Optional | Return domain names that include hyphens. Permitted values are: - True - False Default value is "False"                                                                                                      |
| UseNumbers    | string | Optional | Return domain names that include numbers. Permitted values are: - True - False Default value is "True" |
| Basic | string | Optional | Higher values return suggestions that are built by adding prefixes, suffixes, and words to the original input. Permitted values are: - Off - Low - Medium - High Default value is: "Medium"                                                            |
| Related     | string | Optional | Higher values return domain names by interpreting the input semantically and construct suggestions with a similar meaning. "Related=High" will find terms that are synonyms of your input. Permitted values are: - Off - Low - Medium - High Default value is "High" |
| Similar | string | Optional | Higher values return suggestions that are similar to the customer’s input, but not necessarily in meaning. "Similar=High" will generate more creative terms, with a slightly looser relationship to your input, than "Related=High". Permitted values are: - Off - Low - Medium - High Default value is "Medium" |
| Topical     | string | Optional | Higher values return suggestions that reflect current topics and popular words. Permitted values are: - Off - Low - Medium - High Default value is "Medium" |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML                                                                                                                     |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send "ResponseType=HTML" or "ResponseType=XML", in your request.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                              |
| ---------------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | NameSpinner                                                              |
| SpinCount    | int | Number of SLDs returned in this response                                                |
| TLDList     | string | TLDs returned in this response                                                     |
| Domain name   | string | SLD for this set of results                                                      |
| Com       | string | Is this SLD available from the .com Registry? "Y" indicates yes; "N" indicates no.                          |
| ComScore     | int | The Registry’s score of how closely this SLD matches the SLD input you supplied in the query string. Highest possible score is 1000. |
| Net       | string | Is this SLD available from the .com Registry? "Y" indicates yes; "N" indicates no.                          |
| NetScore     | int | The Registry’s score of how closely this SLD matches the SLD input you supplied in the query string. Highest possible score is 1000. |
| Cc        | string | Is this SLD available from the .com Registry? "Y" indicates yes; "N" indicates no.                          |
| CcScore     | int | The Registry’s score of how closely this SLD matches the SLD input you supplied in the query string. Highest possible score is 1000. |
| OriginalSLD   | string | The SLD value you supplied in the input query string                                          |
| errcount     | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                    |
| errX       | string | Error messages explaining the failure. These can be presented as is back to the client.                       |
| done       | string | True value indicates this entire response has reached you successfully.                                |

Example Output
--------------

The following query retrieves 8 domains that resemble the input SLD "hand" and requests the response in format of either "text" \(default\), "xml" or "html"

```
https://resellertest.enom.com/interface.asp?command=NameSpinner&UID=YourAccountID&PW=YourApiToken&SLD=hand&TLD=com&MaxResults=8&Similar=High&ResponseType={Optional}
```
```
<interface-response>
<namespin>
<spincount>8</spincount>
<TLDList/>
<domains>
<domain name="Richandbob" com="n" comscore="778" net="y" netscore="806" tv="y" tvscore="742" cc="y" ccscore="709"/>
<domain name="Cardbyhand" com="y" comscore="799" net="y" netscore="774" tv="y" tvscore="748" cc="y" ccscore="723"/>
<domain name="Side" com="n" comscore="831" net="n" netscore="806" tv="y" tvscore="797" cc="y" ccscore="764"/>
<domain name="Glove" com="n" comscore="819" net="n" netscore="793" tv="y" tvscore="784" cc="y" ccscore="751"/>
<domain name="Handytune" com="y" comscore="783" net="y" netscore="758" tv="y" tvscore="732" cc="y" ccscore="707"/>
<domain name="Palm" com="n" comscore="806" net="n" netscore="780" tv="y" tvscore="771" cc="y" ccscore="738"/>
<domain name="Fingers" com="n" comscore="793" net="n" netscore="767" tv="y" tvscore="758" cc="y" ccscore="725"/>
<domain name="Shandawl" com="y" comscore="751" net="y" netscore="726" tv="y" tvscore="700" cc="y" ccscore="675"/>
</domains>
</namespin>
<originalsld>hand</originalsld>
<Command>NAMESPINNER</Command>
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
<ExecTime>0.750</ExecTime>
<Done>true</Done>
<debug/>
<TrackingKey>9efbb8e7-1cc3-4f2b-bbc2-cbe2517522aa</TrackingKey>
<RequestDateTime>12/9/2011 3:33:28 AM</RequestDateTime>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>spincount: </STRONG>8<br />
<STRONG>TLDList: </STRONG><br />
DomainSld1=SideOnline Domain1Tld1=com Domain1TldStatus1=n Domain1TldScore1=783 Domain1Tld2=net Domain1TldStatus2=y Domain1TldScore2=811 Domain1Tld3=tv Domain1TldStatus3=y Domain1TldScore3=748 Domain1Tld4=cc Domain1TldStatus4=y Domain1TldScore4=716 DomainSld2=Side Domain2Tld1=com Domain2TldStatus1=n Domain2TldScore1=812 Domain2Tld2=net Domain2TldStatus2=n Domain2TldScore2=787 Domain2Tld3=tv Domain2TldStatus3=y Domain2TldScore3=778 Domain2Tld4=cc Domain2TldStatus4=y Domain2TldScore4=745 DomainSld3=TheSide Domain3Tld1=com Domain3TldStatus1=n Domain3TldScore1=805 Domain3Tld2=net Domain3TldStatus2=n Domain3TldScore2=780 Domain3Tld3=tv Domain3TldStatus3=y Domain3TldScore3=772 Domain3Tld4=cc Domain3TldStatus4=y Domain3TldScore4=739 DomainSld4=MySide Domain4Tld1=com Domain4TldStatus1=n Domain4TldScore1=803 Domain4Tld2=net Domain4TldStatus2=n Domain4TldScore2=778 Domain4Tld3=tv Domain4TldStatus3=y Domain4TldScore3=769 Domain4Tld4=cc Domain4TldStatus4=y Domain4TldScore4=737 DomainSld5=Hands Domain5Tld1=com Domain5TldStatus1=n Domain5TldScore1=799 Domain5Tld2=net Domain5TldStatus2=n Domain5TldScore2=774 Domain5Tld3=tv Domain5TldStatus3=y Domain5TldScore3=765 Domain5Tld4=cc Domain5TldStatus4=y Domain5TldScore4=733 DomainSld6=TheHands Domain6Tld1=com Domain6TldStatus1=n Domain6TldScore1=793 Domain6Tld2=net Domain6TldStatus2=n Domain6TldScore2=768 Domain6Tld3=tv Domain6TldStatus3=y Domain6TldScore3=759 Domain6Tld4=cc Domain6TldStatus4=y Domain6TldScore4=727 DomainSld7=HandsOnline Domain7Tld1=com Domain7TldStatus1=n Domain7TldScore1=792 Domain7Tld2=net Domain7TldStatus2=n Domain7TldScore2=767 Domain7Tld3=tv Domain7TldStatus3=y Domain7TldScore3=758 Domain7Tld4=cc Domain7TldStatus4=y Domain7TldScore4=725 DomainSld8=MyHands Domain8Tld1=com Domain8TldStatus1=n Domain8TldScore1=790 Domain8Tld2=net Domain8TldStatus2=n Domain8TldScore2=765 Domain8Tld3=tv Domain8TldStatus3=y Domain8TldScore3=757 Domain8Tld4=cc Domain8TldStatus4=y Domain8TldScore4=724
<STRONG>count: </STRONG>8<br />
<STRONG>originalsld: </STRONG>hand<br />
<STRONG>Command: </STRONG>NAMESPINNER<br />
<STRONG>APIType: </STRONG>API.NET<br />
<STRONG>Language: </STRONG>eng<br />
<STRONG>ErrCount: </STRONG>0<br />
<STRONG>ResponseCount: </STRONG>0<br />
<STRONG>MinPeriod: </STRONG>1<br />
<STRONG>MaxPeriod: </STRONG>10<br />
<STRONG>Server: </STRONG>sjl0vwresell_t<br />
<STRONG>Site: </STRONG>eNom<br />
<STRONG>IsLockable: </STRONG>True<br />
<STRONG>IsRealTimeTLD: </STRONG>True<br />
<STRONG>TimeDifference: </STRONG>+8.00<br />
<STRONG>ExecTime: </STRONG>0.547<br />
<STRONG>Done: </STRONG>true<br />
<STRONG>TrackingKey: </STRONG>108f7fa0-15c8-4cdf-be4d-b99f4dac2ee8<br />
<STRONG>RequestDateTime: </STRONG>2/4/2015 3:48:04 PM<br />
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
spincount=8
TLDList=

DomainSld1=SideOnline Domain1Tld1=com Domain1TldStatus1=n Domain1TldScore1=783 Domain1Tld2=net Domain1TldStatus2=y Domain1TldScore2=811 Domain1Tld3=tv Domain1TldStatus3=y Domain1TldScore3=748 Domain1Tld4=cc Domain1TldStatus4=y Domain1TldScore4=716

DomainSld2=Side Domain2Tld1=com Domain2TldStatus1=n Domain2TldScore1=812 Domain2Tld2=net Domain2TldStatus2=n Domain2TldScore2=787 Domain2Tld3=tv Domain2TldStatus3=y Domain2TldScore3=778 Domain2Tld4=cc Domain2TldStatus4=y Domain2TldScore4=745

DomainSld3=TheSide Domain3Tld1=com Domain3TldStatus1=n Domain3TldScore1=805 Domain3Tld2=net Domain3TldStatus2=n Domain3TldScore2=780 Domain3Tld3=tv Domain3TldStatus3=y Domain3TldScore3=772 Domain3Tld4=cc Domain3TldStatus4=y Domain3TldScore4=739

DomainSld4=MySide Domain4Tld1=com Domain4TldStatus1=n Domain4TldScore1=803 Domain4Tld2=net Domain4TldStatus2=n Domain4TldScore2=778 Domain4Tld3=tv Domain4TldStatus3=y Domain4TldScore3=769 Domain4Tld4=cc Domain4TldStatus4=y Domain4TldScore4=737

DomainSld5=Hands Domain5Tld1=com Domain5TldStatus1=n Domain5TldScore1=799 Domain5Tld2=net Domain5TldStatus2=n Domain5TldScore2=774 Domain5Tld3=tv Domain5TldStatus3=y Domain5TldScore3=765 Domain5Tld4=cc Domain5TldStatus4=y Domain5TldScore4=733

DomainSld6=TheHands Domain6Tld1=com Domain6TldStatus1=n Domain6TldScore1=793 Domain6Tld2=net Domain6TldStatus2=n Domain6TldScore2=768 Domain6Tld3=tv Domain6TldStatus3=y Domain6TldScore3=759 Domain6Tld4=cc Domain6TldStatus4=y Domain6TldScore4=727

DomainSld7=HandsOnline Domain7Tld1=com Domain7TldStatus1=n Domain7TldScore1=792 Domain7Tld2=net Domain7TldStatus2=n Domain7TldScore2=767 Domain7Tld3=tv Domain7TldStatus3=y Domain7TldScore3=758 Domain7Tld4=cc Domain7TldStatus4=y Domain7TldScore4=725

DomainSld8=MyHands Domain8Tld1=com Domain8TldStatus1=n Domain8TldScore1=790 Domain8Tld2=net Domain8TldStatus2=n Domain8TldScore2=765 Domain8Tld3=tv Domain8TldStatus3=y Domain8TldScore3=757 Domain8Tld4=cc Domain8TldStatus4=y Domain8TldScore4=724

count=8
originalsld=hand
Command=NAMESPINNER
APIType=API.NET
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=sjl0vwresell_t
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+8.00
ExecTime=0.125
Done=true
TrackingKey=eac484f9-f6d1-4761-bdb7-f29ae171b6c4
RequestDateTime=2/4/2015 3:48:33 PM
```
Related Commands
----------------

AddToCart

Check

Purchase