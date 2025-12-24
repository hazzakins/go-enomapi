GetCustomerPaymentInfo
======================

Retrieve customer payment information.

Usage
-----

Use this command to retrieve customer payment information.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/EditContact.asp](https://resellertest.enom.com/myaccount/EditContact.asp)

The Credit Card Information box contains the return values supplied by the GetCustomerPaymentInfo command.

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

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| CCName | Credit card holder’s name                                    |
| CCNumber     | Credit card number |
| CCType | Credit card type                                         |
| CCMonth     | Credit card expiration month |
| CCYear | Credit card expiration year                                   |
| CCAddress    | Credit card billing address street address |
| CCZip | Credit card billing address postal code                             |
| CCCity      | Credit card billing address city |
| CCStateProvince | Credit card billing address province                               |
| CCCountry    | Credit card billing address country |
| CCPhoneDial | Telephone country code                                      |
| CCPhone     | Telephone number with area code |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the payment information on record for an account, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetCustomerPaymentInfo&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetCustomerPaymentInfo&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetCustomerPaymentInfo&uid=resellid
&pw=resellpw&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>

 <CCName>John Doe</CCName>

 <CCNumber>************5215</CCNumber>

 <CCType>MASTERCARD</CCType>

 <CCMonth>11</CCMonth>

 <CCYear>2005</CCYear>

 <CCAddress>17462</CCAddress>

 <CCZip>98052</CCZip>

 <CCCity>Hometown</CCCity>

 <CCStateProvince>WA</CCStateProvince>

 <CCCountry>US</CCCountry>

 <CCPhoneDial>1</CCPhoneDial>

 <CCPhone>5555559999</CCPhone>

 <Command>GETCUSTOMERPAYMENTINFO</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>8.203125E-02</ExecTime>

 <Done>true</Done>

 <![CDATA [ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

CCName: JohnDoe

CCNumber: ************5215

CCType: MASTERCARD

CCMonth: 02

CCYear: 2016

CCAddress: 100 Main St.

CCZip: 99999

CCCity: ekm

CCStateProvince: WA

CCCountry: us

CCPhoneDial:

CCPhone: 1.5555559999

Command: GETCUSTOMERPAYMENTINFO

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.031

Done: true

RequestDateTime: 2/3/2015 5:03:34 PM
```
RequestDateTime=2/3/2015 5:04:00 PM
```