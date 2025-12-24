AuthorizeTLD
============

Create or add to a list of TLDs that you offer to your resellers and retail customers.

Usage
-----

Use this command to specify the TLDs that you offer to your resellers and retail customers. This command controls which TLDs your Web site can offer for registrations, renewals, and transfers.

The first time you use this command, the list of TLDs we return to you switches from the list of all TLDs we offer, which is the default, to your own list that contains only the TLDs you authorize.

You can also use this command to add to your list of authorized TLDs. When you want to add one TLD, run AuthorizeTLD with that single TLD.

To revert from the authorized TLD mode back to the default mode \(all TLDs that we support\), or to remove TLDs from your list, use the RemoveTLD command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/Settings.asp](https://resellertest.enom.com/myaccount/Settings.asp)

Clicking the TLD Pricing tab shows either the complete set of TLDs we offer, or the TLDs you have authorized.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

You can only authorize TLDs that we offer.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid
&pw=yourpassword&paramname=paramvalue
&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                        | Max Size |
| --------------- | ------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                      | 20 |
| PW | Required               | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                          | 4 |
| DomainList | Either TLD or DomainList is Required | Comma-separated list of TLDs to add to your list of authorized TLDs. For example, if you want to add .com, .net, and .org, use DomainList=com,net,org | 100   |
| TLD       | Either TLD or DomainList is Required | One top-level domain \(TLD\) to add to your list of authorized TLDs. For example, if you want to add .biz, use TLD=biz                 | 15 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| AuthorizeTLDX | TLD to add to authorized list. Indexed X when ResponseType=Text or HTML.            |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query authorizes TLDs .com, .net, and .org for account resellid, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=authorizetld&uid=resellid&pw=resellpw
&domainlist=com,net,org,info,biz&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=authorizetld&uid=resellid&pw=resellpw
&domainlist=com,net,org,info,biz&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=authorizetld&uid=resellid&pw=resellpw
&domainlist=com,net,org,info,biz&responsetype=text
```
In the response, a Party ID and an ErrCount value of 0 confirm that the query was successful:

<authorizetld>com</authorizetld>

<authorizetld>net</authorizetld>

<authorizetld>org</authorizetld>

<authorizetld>info</authorizetld>

<authorizetld>biz</authorizetld>

</tldlist>

<Command>AUTHORIZETLD</Command>

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

<ExecTime>0.188</ExecTime>

<Done>true</Done>

<debug/>

<TrackingKey>1b2b72b9-77df-4666-a0cc-30c02aefede9</TrackingKey>

<RequestDateTime>12/7/2011 3:34:07 AM</RequestDateTime>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

authorizetld1: com

authorizetld2: net

authorizetld3: org

authorizetld4: info

authorizetld5: biz

Command: AUTHORIZETLD

APIType: API.NET

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod: 1

MaxPeriod: 10

Server: sjl0vwresell_t

Site: eNom

TimeDifference: +0.00

ExecTime: 0.203

Done: true

TrackingKey: a53be21e-f6a4-4272-bc03-90cd5f7c29a6

RequestDateTime: 2/3/2015 1:05:13 PM
```
;Machine is SJL0VWRESELL_T1

Server=sjl0vwresell_t1

ExecTime=0.047

TrackingKey=13df8fcc-0818-41f1-b281-d356be6691f1

RequestDateTime=2/3/2015 1:05:42 PM
```