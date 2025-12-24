GetCustomerDefinedData
======================

Retrieve customer-defined data.

Usage
-----

Use this command to retrieve custom data.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

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
| Input Parameter | Status | Description                                                           | Max Size |
| --------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                        | 20 |
| PW | Required      | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                            | 4 |
| ObjectID | Required      | Object ID number, an integer assigned when this customer-defined field was first established. | 2    |
| Type      | Required | Object type. Permitted values are: 1 Data pertaining to an account 2 Data pertaining to a domain 3 Data pertaining to an order | 1 |
| SLD | Required If Type=2 | Second-level domain name \(for example, enom in enom.com\) | 63    |
| TLD       | Required If Type=2 | Top-level domain name \(extension\)                                              | 15 |
| OrderID | Required If Type=3 | Order ID, which you can retrieve using the GetDomainStatus command | 11    |
| Key       | Required | Title of this entry, or label describing this data field                                    | 50 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                                                          |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command | Name of command executed                                                                    |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                   |
| Done       | True indicates this entire response has reached you successfully. |
| Value | Content of this entry                                                                     |
| DisplayFlag   | Visibility to subaccount. Options are: 0 This entry not visible when logged on using subaccount ID 1 This entry is visible when logged on using subaccount ID |
| EnteredBy | Name of the person adding this entry                                                              |
| LastUpdatedDate | Last date on which this entry was changed |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query retrieves the value for the customer-defined data titled FavoriteCuisine for account resellid, and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=getcustomerdefineddata&uid=resellid&pw=resellpw
&ObjectID=1&Type=1&Key=FavoriteCuisine&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getcustomerdefineddata&uid=resellid&pw=resellpw
&ObjectID=1&Type=1&Key=FavoriteCuisine&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getcustomerdefineddata&uid=resellid&pw=resellpw
&ObjectID=1&Type=1&Key=FavoriteCuisine&responsetype=text
```
The response is as follows:

```
<?xml version="1.0" ?>

 <Value>Italian</Value>

 <DisplayFlag>True</DisplayFlag>

 <EnteredBy>John</EnteredBy>

 <LastUpdatedDate>7/7/2003 5:46:10 PM</LastUpdatedDate>

 </CustomerData>

 <Command>GETCUSTOMERDEFINEDDATA</Command>

 <ErrCount>0</ErrCount>

 <Server>RESELLERTEST</Server>

 <Site>enom</Site>

 <IsLockable>True</IsLockable>

 <IsRealTimeTLD>True</IsRealTimeTLD>

 <Done>true</Done>

 <![CDATA [ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Value:

DisplayFlag:

EnteredBy:

LastUpdatedDate:

Command: GETCUSTOMERDEFINEDDATA

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

ExecTime: 0.063

Done: true

RequestDateTime: 2/3/2015 5:01:33 PM
```
;Machine is SJL0VWRESELL_T

Server=SJL0VWRESELL_T

ExecTime=0.078

RequestDateTime=2/3/2015 5:02:00 PM
```