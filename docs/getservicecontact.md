GetServiceContact
=================

Retrieve the customer service contact information for a domain name account.

Usage
-----

Use this command to retrieve customer service contact information for an account. If there is a PDQ subscription associated with this account, this command can return the PDQ contact information. Otherwise, this command returns the Billing contact information.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/pdq/RE\_Default.asp?maintab=overview](https://resellertest.enom.com/pdq/RE_Default.asp?maintab=overview)

Clicking the my site tab calls the GetServiceContact command.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The domain name must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                                | Max Size |
| --------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                              | 20 |
| PW | Required | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                                  | 4 |
| Service | Optional | Use Service=PDQ to return the contact information for this account’s PDQ subscription. If the Service parameter is omitted, the Billing contact information is returned. | 20    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |
| Organization   | Name of contact’s organization |
| Address1 | Contact address, line 1                                     |
| Address2     | Contact address, line 2 |
| City | Contact’s city                                          |
| StateProvince  | Contact’s state or province |
| PostalCode | Contact’s postal code                                      |
| Country     | Contact’s country |
| Phone | Contact’s phone number                                      |
| Fax       | Contact’s fax number |
| EmailAddress | Contact’s email address                                     |
| URL       | URL of PDQ site, if any |
| ResellerKey | Reseller key of PDQ site, if any                                 |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the customer service contact information for resellid and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GETSERVICECONTACT&uid=resellid
&pw=resellpw&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GETSERVICECONTACT&uid=resellid
&pw=resellpw&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GETSERVICECONTACT&uid=resellid
&pw=resellpw&responsetype=text
```
In the response, the presence of contact information, and the ErrCount value 0, indicate that the query was successful:

```
<?xml version="1.0" ?>

 <Organization>Reseller Documents Inc.</Organization>

 <Address1>111 Main St.</Address1>

 <Address2 />

 <City>Hometown</City>

 <StateProvince>WA</StateProvince>

 <PostalCode>99999</PostalCode>

 <Country>US</Country>

 <Phone>+1.5555555555</Phone>

 <Fax>+1.5555555556</Fax>

 <EmailAddress>[email protected]</EmailAddress>

 <URL />

 <ResellerKey />

 </ServiceContact>

 <Command>GETSERVICECONTACT</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod>1</MinPeriod>

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable>True</IsLockable>

 <IsRealTimeTLD>True</IsRealTimeTLD>

 <ExecTime>8.203125E-02</ExecTime>

 <Done>true</Done>

 <![CDATA[ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Organization: Extraordinary Sales

Address1: 15801 NE 24th Street

City: ekm

StateProvince: WA

PostalCode: 98008

Country: US

Phone: +91.9544048048

Fax:

EmailAddress: [email protected]

Command: GETSERVICECONTACT

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MinPeriod:

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

IsLockable:

IsRealTimeTLD:

TimeDifference: +0.00

ExecTime: 0.078

Done: true

RequestDateTime: 2/4/2015 12:29:34 PM
```
[email protected]

RequestDateTime=2/4/2015 12:29:55 PM
```