UpdateRenewalSettings
=====================

Update the settings regarding our notifying your customers about domain renewals.

Usage
-----

Use this command to set or change the way in which we notify your customers about upcoming domain renewals.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID **resellid**, password **resellpw**.

[https://resellertest.enom.com/myaccount/RenewalSettings.asp?RenewalSetting=0](https://resellertest.enom.com/myaccount/RenewalSettings.asp?RenewalSetting=0)

The save changes button calls the UpdateRenewalSettings command.

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
| Input Parameter | Status | Description                                                                                                        | Max Size |
| --------------- | --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                                                     | 20 |
| PW | Required                                     | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                                         | 4 |
| RenewalSetting | Optional                                     | Renewal setting. Permitted values are: 0 Do not contact my customers for renewal 1 Email my customers 25 and 10 days before expiration; send them to my Web site to renew 2 Contact and charge my customers for renewals | 1    |
| RenewalBCC   | Optional | BCC me on all email correspondence sent to my customers. Permitted values are: 0 Do not BCC me 1 BCC me                                                         | 1 |
| AcceptTerms | Optional; must be accepted before you can use our credit card processing services | Do I accept eNom’s credit card processing agreement? Permitted values are: 0 I do not accept eNom’s credit card processing agreement 1 I accept eNom’s credit card processing agreement | 1    |
| URL       | Recommended when RenewalSetting =1 | URL where customers can renew their domains; use format [www.DomainName.com](http://www.DomainName.com/)                                                        | 63 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ----------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| AcceptTermsStatus | Success status of updating credit card terms setting                       |
| RenewalSetting  | Success status of updating renewal settings |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query sets preferences for notifying customers about domain renewals, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=UPDATERENEWALSETTINGS&uid=resellid&pw=resellpw
&RenewalSetting=2&RenewalBCC=1&AcceptTerms=1
&URL=www.resellerdocs.com&ResponseType=XML
```
```
https://resellertest.enom.com/interface.asp?
command=UPDATERENEWALSETTINGS&uid=resellid&pw=resellpw
&RenewalSetting=2&RenewalBCC=1&AcceptTerms=1
&URL=www.resellerdocs.com&ResponseType=html
```
```
https://resellertest.enom.com/interface.asp?
command=UPDATERENEWALSETTINGS&uid=resellid&pw=resellpw
&RenewalSetting=2&RenewalBCC=1&AcceptTerms=1
&URL=www.resellerdocs.com&ResponseType=text
```
In the response, the return values Successful and the ErrCount value 0 confirm that the query was successful:

```
<?xml version="1.0" ?>

 <AcceptTermsStatus>Successful</AcceptTermsStatus>

 <RenewalSetting>Successful</RenewalSetting>

 <Command>UPDATERENEWALSETTINGS</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>LOCALHOST</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.046875</ExecTime>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

AcceptTermsStatus: Successful

RenewalSetting: Successful

Command: UPDATERENEWALSETTINGS

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.516

Done: true

RequestDateTime: 2/9/2015 2:31:27 PM
```
;Machine is SJL0VWRESELL_T

Server=SJL0VWRESELL_T

ExecTime=0.063

RequestDateTime=2/9/2015 2:31:50 PM
```