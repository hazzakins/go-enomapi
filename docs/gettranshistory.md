GetTransHistory
===============

Return a list of up to 25 transactions in an account.

Usage
-----

Use this command to retrieve an overview of transactions.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/myaccount/](https://resellertest.enom.com/myaccount/)

When the my enom submenu is open on the left side of the Web page, clicking the transactions link calls the GetTransHistory command.

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
| Input Parameter | Status | Description                                                                            | Max Size |
| --------------- | ------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| UID       | Required | Account login ID                                                                         | 20 |
| PW | Required                        | Account password | 20    |
| ResponseType  | Optional | Format of response. Permitted values are Text \(default\), HTML, or XML                                             | 4 |
| StartPosition | Optional; default is 1                 | Return sets of 25 records in reverse chronological order. For example, StartPosition=26 returns the 26th through 50th most recent transactions. Defaults to 1. | 4    |
| StartDate    | Optional; if omitted, six months of orders are returned | Beginning date of transactions to return. Must be no more than six months before EndDate. Permitted format is MM/DD/YYYY                    | 10 |
| EndDate | Optional; default is today’s date            | End date of transactions to return. Must be no more than six months after StartDate. Permitted format is MM/DD/YYYY | 10    |

Returned Parameters and Values
------------------------------

| Output Parameter   | Description |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command       | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                                                                                                                                                                                                                                     |
| ErrX         | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                                                                                                                                                                                                                                                                                                    |
| AdjustedBeginDate  | Start date for this response |
| AdjustedEndDate | End date for this response                                                                                                                                                                                                                                                                                                                        |
| FName        | First \(given\) name |
| LName | Last \(family\) name                                                                                                                                                                                                                                                                                                                           |
| PartyID       | Party identification number. Format is 32 hexadecimal characters, hyphenated, in curly braces |
| CurrentStartPosition | Start position of this response, in the reverse chronological list of transactions                                                                                                                                                                                                                                                                                            |
| RecordCount     | Number of transactions in the transaction history |
| NewStartPosition | Start position of the next response, in the reverse chronological list of transactions                                                                                                                                                                                                                                                                                          |
| OldStartPosition   | Start position of the previous response, in the reverse chronological list of transactions |
| Trans-Date | Time stamp indicates when the transaction was submitted to our accounting system                                                                                                                                                                                                                                                                                             |
| Amount        | Amount of the transaction, $US |
| TransDescription | Word description of this transaction                                                                                                                                                                                                                                                                                                                   |
| TransType      | Identification number of this transaction type. Values are: 1 Order 2 Debit Credit Card 3 Credit Credit Card 4 Refill Reseller Account 5 Check Refill 7 Credit Reseller Balance 6 Debit Reseller Balance 8 Transfer Order 9 Reverse Refi ll 10 Refund and Reverse Commissions 11 Transfer from commission account 12 Commission account cashout 16 Sales Commission 17 CCTransaction Fee 20 Refund and Reverse Commissions with Charge Back 21 Charge Back 22 Reverse Commission 23 Reverse Commission with Charge Back 25 Preregistration Fee 26 Activation Fee 28 Refill Service Convenience Charge 30 PDQ Annual Fee 32 Refill Reseller Account \(No Service Charge\) |
| LinkValue | Used in XML to build links to individual orders.                                                                                                                                                                                                                                                                                                             |
| OrderID       | Order identification number, a nine-digit number |
| OrderProcessFlag | True indicates the order has been successfully processed.                                                                                                                                                                                                                                                                                                        |
| Balance       | Balance in the account following this transaction |
| ComBalance | Commission balance following this transaction.                                                                                                                                                                                                                                                                                                              |
| TransStatus     | Status of the transaction. Options are: 1 Ready for billing \(Processing in UI\) 2 Transaction failed \(Failed in UI\) 3 Transaction successful \(Successful in UI\) 6 Void \(Voided in UI\) |
| Language | Language                                                                                                                                                                                                                                                                                                                                 |
| ResponseCount    | Response count |
| MinPeriod | Minimum registration period, for domain name purchases                                                                                                                                                                                                                                                                                                          |
| MaxPeriod      | Maximum registration period, for domain name purchases |
| Server | Name of server                                                                                                                                                                                                                                                                                                                              |
| Site         | Name of site |
| IsLockable | Is this a lockable domain name.                                                                                                                                                                                                                                                                                                                     |
| IsRealTimeTLD    | Is this a TLD that supports real-time purchase. |
| TimeDifference | Time difference between end user and our server                                                                                                                                                                                                                                                                                                              |
| ExecTime       | Time to execute this API query |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the transaction history for account resellid, and requests the response in given format:

```
https://resellertest.enom.com/interface.asp?
command=GetTransHistory&uid=resellid&pw=resellpw
&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=GetTransHistory&uid=resellid&pw=resellpw
&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=GetTransHistory&uid=resellid&pw=resellpw
&responsetype=text
```
By default, the response gives a summary of the 25 most recent transactions:

```
<?xml version="1.0" ?>

 <AdjustedBeginDate />

 <AdjustedEndDate />

 <FName>John</FName>

 <LName>Doe</LName>

 <PartyID>{BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}</PartyID>

 <CurrentStartPosition>1</CurrentStartPosition>

 <RecordCount>396</RecordCount>

 <NewStartPosition>26</NewStartPosition>

 <OldStartPosition>0</OldStartPosition>

  <Trans-Date>7/18/2007 1:25:11 PM</Trans-Date>

  <Amount>-$4.40</Amount>

  <TransDescription>Order</TransDescription>

  <TransType>1</TransType>

   OrderDetail.asp?OrderID=157385793&OrderProcess=Successful

  </LinkValue>

  <OrderID>157385793</OrderID>

  <OrderProcessFlag>True</OrderProcessFlag>

  <Balance>$7,028.39</Balance>

  <ComBalance />

  <TransStatus>Successful</TransStatus>

  </Transaction>

  <Trans-Date>7/16/2007 2:05:01 AM</Trans-Date>

  <Amount>-$8.95</Amount>

   OrderDetail.asp?OrderID=157385407&OrderProcess=Successful

  <OrderID>157385407</OrderID>

  <Balance>$7,032.79</Balance>

.

.

.

 </TransHistory>

 <Command>GETTRANSHISTORY</Command>

 <Language>eng</Language>

 <ErrCount>0</ErrCount>

 <ResponseCount>0</ResponseCount>

 <MinPeriod />

 <MaxPeriod>10</MaxPeriod>

 <Server>RESELLER1-STG</Server>

 <Site>enom</Site>

 <IsLockable />

 <IsRealTimeTLD />

 <TimeDifference>+0.00</TimeDifference>

 <ExecTime>0.750</ExecTime>

 <Done>true</Done>

 <![CDATA[ ]]>

 </debug>

</interface-response>
```
```
;URL Interface

;Machine is SJL0VWRESELL_T

;Encoding Type is utf-8

FName: Rosh

LName: Bach

PartyID: {BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}

CurrentStartPosition: 1

RecordCount: 1362

NewStartPosition: 26

OldStartPosition: 0

Trans-Date1: 2/4/2015 12:20:50 PM

Amount1: -$8.95

TransDescription1: Order

TransType1: 1

LinkValue1: OrderDetail.asp?OrderID=161808041&OrderProcess=Successful

OrderID1: 161808041

OrderProcessFlag1: True

Balance1: -$1,873.13

ComBalance1:

TransStatus1: Successful

.

.

.

Command: GETTRANSHISTORY

APIType: API

Language: eng

ErrCount: 0

ResponseCount: 0

MaxPeriod: 10

Server: SJL0VWRESELL_T

Site: eNom

TimeDifference: +0.00

ExecTime: 0.906

Done: true

RequestDateTime: 2/4/2015 12:55:51 PM
```
.

.

.

ExecTime=0.125

RequestDateTime=2/4/2015 12:56:46 PM
```