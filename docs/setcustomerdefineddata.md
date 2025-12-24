SetCustomerDefinedData
======================

Enter data fields defined by you, and the corresponding data.

Usage
-----

Use this command to create a storage place for data fields that do not currently exist.

Also use this command to put data into those custom data fields. This command allows you to add data one field at a time.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com

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
command=nameofcommand&uid=yourloginid&pw=yourpassword&
paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                             | Max Size |
| --------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                                           | 20 |
| PW | Required      | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                               | 4 |
| ObjectID | Required      | Object ID number, an integer assigned by you. | 2    |
| Type      | Required | Object type. Options are: 1 Data pertaining to an account 2 Data pertaining to a domain 3 Data pertaining to an order        | 1 |
| SLD | Required if Type=2 | Second-level domain name \(for example, enom in enom.com\) | 63    |
| TLD       | Required if Type=2 | Top-level domain name \(extension\)                                                 | 15 |
| OrderID | Required if Type=3 | Order ID, which you can retrieve using GetDomainStatus. | 11    |
| Key       | Required | Title of this entry, or label describing this data field                                       | 50 |
| Value | Required      | Content of this entry | 50    |
| DisplayFlag   | Required | Visibility to subaccount. Options are: 0 Not visible when logged on with subaccount ID 1 Visible when logged on using subaccount ID | 1 |
| EnteredBy | Recommended    | Name of the person adding this entry | 50    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query creates a new data field labeled FavoriteCuisine and assigns a value Italian for this account.:

```
https://resellertest.enom.com/interface.asp?
command=SetCustomerDefinedData&uid=resellid&pw=resellpw
&ObjectID=1&Type=1&Key=FavoriteCuisine&Value=Italian&
DisplayFlag=0&EnteredBy=John&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=SetCustomerDefinedData&uid=resellid&pw=resellpw
&ObjectID=1&Type=1&Key=FavoriteCuisine&Value=Italian&
DisplayFlag=0&EnteredBy=John&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=SetCustomerDefinedData&uid=resellid&pw=resellpw
&ObjectID=1&Type=1&Key=FavoriteCuisine&Value=Italian&
DisplayFlag=0&EnteredBy=John&responsetype=text
```
In the response, the value 0 for ErrCount confirms that the query executed successfully:

<Command>SETCUSTOMERDEFINEDDATA</Command>

<Language>eng</Language>

<ErrCount>0</ErrCount>

<ResponseCount>0</ResponseCount>

<MinPeriod/>

<MaxPeriod>10</MaxPeriod>

<Server>SJL21WRESELLT01</Server>

<Site>eNom</Site>

<IsLockable/>

<IsRealTimeTLD/>

<TimeDifference>+0.00</TimeDifference>

<ExecTime>0.313</ExecTime>

<Done>true</Done>

<RequestDateTime>12/11/2011 10:44:27 PM</RequestDateTime>

<debug></debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T1

;Encoding Type is utf-8

Command: SETCUSTOMERDEFINEDDATA

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T1

Site: eNom

TimeDifference: +0.00

ExecTime: 0.266

Done: true

RequestDateTime: 2/5/2015 2:40:51 PM
```
ExecTime=0.078

RequestDateTime=2/5/2015 2:41:14 PM
```