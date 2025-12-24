DeleteCustomerDefinedData
=========================

Delete customer-defined data records.

Usage
-----

Use this command to delete customer-defined data.

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

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query deletes the customer-defined data labeled FavoriteCuisine from account resellid, and sends the response in XML format:

```
https://resellertest.enom.com/interface.asp?
command=deletecustomerdefineddata&uid=resellid
&pw=resellpw&ObjectID=1&Type=1&Key=FavoriteCuisine
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=deletecustomerdefineddata&uid=resellid
&pw=resellpw&ObjectID=1&Type=1&Key=FavoriteCuisine
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=deletecustomerdefineddata&uid=resellid
&pw=resellpw&ObjectID=1&Type=1&Key=FavoriteCuisine
&responsetype=text
```
In the response, an ErrCount value 0 confirms that the query was successful:

```
<?xml version="1.0" ?>

 <Command>DELETECUSTOMERDEFINEDDATA</Command>

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

Command: DELETECUSTOMERDEFINEDDATA

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

RequestDateTime: 2/3/2015 3:11:56 PM
```
ExecTime=0.063

RequestDateTime=2/3/2015 3:13:29 PM
```