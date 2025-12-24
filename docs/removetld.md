RemoveTLD
=========

Remove TLDs from your list of authorized TLDs that you offer to your resellers and retail customers. Or, revert to our default list, which includes all TLDs that we support.

Usage
-----

Use this command to remove TLDs from your custom list of authorized TLDs that you offer to your resellers and retail customers. This command controls the list of TLDs we send you for registrations, renewals, and transfers.

This command does not function unless you have previously used the AuthorizeTLD command to authorize TLDs. It will not remove TLDs if your account is using our default list.

If you run RemoveTLD with one TLD, our system starts with your current list, removes the single TLD, and leaves the rest. If you remove all TLDs from your list of authorized TLDs, your account will revert to our default list \(all TLDs that we support\).

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/RenewalPricing.asp?tab=1](https://resellertest.enom.com/myaccount/RenewalPricing.asp?tab=1)

If you remove TLDs from your authorized list, refreshing this page will reflect the change.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

You can only remove TLDs that you added using the AuthorizeTLD command.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                            | Max Size |
| --------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                                                          | 20 |
| PW | Required               | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                              | 4 |
| TLD | Either TLD or DomainList is Required | One top-level domain \(TLD\) to remove from your list of authorized TLDs. For example, if you want to remove .biz, use TLD=biz | 15    |
| DomainList   | Either TLD or DomainList is Required | Comma-separated list of TLDs to remove from your list of authorized TLDs. For example, if you want to remove .com, .net, and .org, use DomainList=com,net,org | 100 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                      |
| ---------------- | --------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                               |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.               |
| Done       | True indicates this entire response has reached you successfully. |
| DeleteTLDX | List of TLDs to remove from the list of TLDs authorized for this account. Indexed X when ResponseType=Text or HTML. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query removes TLDs .us and .ca from resellid’s list of authorized TLDs, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=removetld&uid=resellid&pw=resellpw
&domainlist=us,ca&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=removetld&uid=resellid&pw=resellpw
&domainlist=us,ca&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=removetld&uid=resellid&pw=resellpw
&domainlist=us,ca&responsetype=text
```
In the response, the list of TLDs and the ErrCount value of 0 confirm that the query was successful:

```
<?xml version="1.0" ?>

 <deletetld>us</deletetld>

 <deletetld>ca</deletetld>

 </tldlist>

 <Command>REMOVETLD</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <ExecTime>1.191406</ExecTime>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

deletetld1: us

deletetld2: ca

Command: REMOVETLD

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.109

Done: true

RequestDateTime: 2/5/2015 2:13:12 PM
```
ExecTime=0.047

RequestDateTime=2/5/2015 2:13:35 PM
```