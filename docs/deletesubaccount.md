DeleteSubaccount
================

Delete a subaccount.

Usage
-----

Use this command to delete a retail or reseller subaccount.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

http://enomson/resellers/subaccount-list.asp

On the subaccounts page, the Delete link calls the DeleteSubaccount command with the Action=Confirm parameter.

Click a Delete link

On the Delete Sub-account page, the delete button calls the DeleteSubaccount command with the Action=Delete parameter.

Constraints
-----------

The query must meet the following requirements:

The login ID and password must be valid.

The subaccount must belong to this account.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                                                 | Max Size |
| --------------- | ---------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                               | 20 |
| PW | Required              | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                   | 4 |
| Account | Required              | Subaccount ID number, in NNN-aa-NNNN format. You can retrieve the subaccount ID number using the GetSubAccounts command. | 11    |
| Action     | Optional; default value is Confirm | Action to take. Permitted values are Confirm Confirm which subaccount is to be deleted Delete Delete the subaccount specified in this query | 7 |

Returned Parameters and Values
------------------------------

| Output Parameter | Description                                           |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| Command | Name of command executed                                     |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done       | True indicates this entire response has reached you successfully. |
| Action | Action to be taken with regard to deleting this subaccount                    |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, sendResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameterErr\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query deletes the sub-account with account number 332-ep-2379 and sends the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=deletesubaccount&uid=resellid&pw=resellpw
&account=332-ep-2379&Action=delete&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=deletesubaccount&uid=resellid&pw=resellpw
&account=332-ep-2379&Action=delete&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=deletesubaccount&uid=resellid&pw=resellpw
&account=332-ep-2379&Action=delete&responsetype=text
```
In the response, the Action return parameter and the ErrCount value 0 confirm that the query was successful:

```
<?xml version="1.0" ?>

 <Action>DELETE</Action>

 <Command>DELETESUBACCOUNT</Command>

 <Language>en</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLERTEST</Server>

 <Site>e</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.1523438</ExecTime>

 <Done>true</Done>

 <![CDATA [ ] ]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

Action: DELETE

Command: DELETESUBACCOUNT

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site:

TimeDifference: +0.00

ExecTime: 0.391

Done: true

RequestDateTime: 2/11/2015 1:12:28 PM
```
ExecTime=0.063

RequestDateTime=2/11/2015 1:13:31 PM
```