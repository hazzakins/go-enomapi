RPT\GetReport
=============

Retrieve an itemized list of one type of activity in a domain name account.

Usage
-----

Use this command to retrieve a list of one of the following types of account activity in a domain name account: registrations, renewals, transfers, private label, sub-accounts, accounting, or expiring names.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The beginning date for reports must be no earlier than 6 months before today.
- The "ResponseType" must be XML. You can parse the XML response after receiving it.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=RPT_GETREPORT&uid=(Required)&pw=(Required)&ReportType=(Required)&BeginDate=(Optional)&EndDate=(Optional)&Download=(Required)&responsetype=xml
```
| Input Parameter | Type | Status                                       | Description |
| ---------------- | ------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| command     | string | Required                                      | RPT\_GetReport |
| uid | string | Required | Your Account ID                                                                                                                                                                                                                                                                                                                    |
| pw        | string | Required                                      | Your API Token |
| Version | string | Optional; default is 0 | Version of this command. Permitted values are 0 and 1. Use "Version=1" to receive the return parameters listed in the Returned parameters table below. "Version=0" and "Version=1" return parameters vary slightly in the data returned and in the tag names. We recommend that you use "Version=1",which avoids timeouts by supporting paging of results, and queuing of the query.                                                                                                                               |
| ReportType    | boolean | Required                                      | Type of report to generate. Permitted values are: 0 Registrations 1 Renewals 2 Transfers In 11 Transfers Out 3 Private Label \(Instant Reseller and Registry Rocket\) 4 Sub-Accounts 5 Accounting 6 Expiring Domains 10 Pricing Editor Updates 13 Business Listing Bulk Update 15 Business Listing Domain Update 16 Transactions 17 Order History 18 Periodic Billing History 20 Premium Registration 21 Monthly Invoices 22 Invoice History 23 Invoice Detail 24 List of Invoices 25 Top Watchlist TLDs 26 Watchlist Summary 27 Top Watchlists 28 All Watchlists 29 All Watchlist Details 30 Domain Name History 31 Domain Verification/Suspension |
| BeginDate | boolean | Optional; default is the last day of the month, two months before the current month | First date to include in the report. Must be no earlier than six months before today. Use format MM/DD/YYYY                                                                                                                                                                                                                                                                     |
| EndDate     | string | Optional; default is today’s date                          | Last date to include in the report. Use format MM/DD/YYYY |
| Start | string | Optional; default is 1 | In the overall list of returns for this report type and date range, the start position for the return to this query.                                                                                                                                                                                                                                                                 |
| RecordsToReturn | int | Optional; default is 100 if Download=False                     | Number of items to return in this report. |
| Download | boolean | Required | Do you wish to run report in real time or deliver to this account’s Billing contact email address? Permitted values are: "True" - Send the report via email "False" - Real-time and deliver in API response Reports will always be delivered by email if they extend past four months before today, or if they include more than 25,000 results.                                                                                                                                                 |
| ReportOutputType | string | Required, if "Download=True"                            | If delivering the report via email, "ReportOutputType" must be 2. |
| ResponseType | string | Required | Format of response. Permitted value is "XML".                                                                                                                                                                                                                                                                                                    |

Returned Parameters and Values
------------------------------

The default response format is XML.

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter   | Type | Description                                                     |
| --------------------- | ------- | -------------------------------------------------------------------------------------------------------------------- |
| RptString       | string | List report descriptions                                               |
| ThreeMoPast      | string | The date three months before today                                          |
| SixMoPast       | string | The date six months before today                                           |
| BeginDate       | string | The first date included in this report                                        |
| EndDate        | string | The last date included in this report                                        |
| ID          | int | Report type identification number                                          |
| Name         | string | Report name                                                     |
| ReportType      | string | Report type specified in this query string                                      |
| ReportName      | string | Name of the report type specified in this query string                                |
| QueueReportResults  | int | Should report results be queued. Permitted values are: 1 to queue results; 0 to send results immediately via HTTP. |
| ReportResults     | string | Results specific to this report type. The results, and return parameters, vary by report type.           |
| DomainName      | string | Domain name                                                     |
| OrderID        | int | Order identification number                                             |
| Commission      | float | Commission                                                      |
| SalePrice       | float | Sale price                                                      |
| DateApplied      | string | Date sale was recorded                                                |
| LoginID        | string | Login ID to which this order was delivered                                      |
| NewStartPosition   | string | In the full list of results, the start position of the next page of results                     |
| ThisStartPosition   | string | In the full list of results, the position of the first result listed on this page                  |
| ThisEndPosition    | string | In the full list of results, the position of the last result listed on this page                   |
| PreviousStartPosition | string | In the full list of results, the position of the previous page of results                      |
| TotalReturned     | int | Number of results listed on this page                                        |
| TotalRows       | int | Number of results total for this report                                       |
| ShowPrevious     | string | Should this page show a link to the previous page of results?                            |
| ShowNext       | string | Should this page show a link to the next page of results?                              |
| ShowPaging      | string | Should this page show links to previous and next results?                              |
| Version        | float | Version of this command that generated this return                                  |
| Command        | string | Name of command executed                                               |
| ErrCount       | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.           |
| ErrX         | string | Error messages explaining the failure. These can be presented as is back to the client.              |
| Done         | boolian | "True" indicates this entire response has reached you successfully.                         |

Example Output
--------------

The following query requests a list of registrations for the period 12/1/2009 to 1/1/2010, and sends the response in XML format.

```
https://resellertest.enom.com/interface.asp?command=RPT_GETREPORT&uid=resellid&pw=resellpw&ReportType=0&BeginDate=12/01/2010&EndDate=01/01/2011&responsetype=xml
```
```
<interface-response>
<rpt>
<rptString>Registrations | Renewals | Transfers In | Private Label |
Sub-Accounts | Accounting | Expiring Domains | | | | Pricing Editor Updates |
Transfers Out | | Business Listing Bulk Update | | Business Listing
Domain Update | Transactions | Order History | Periodic Billing History | |
Premium Registration | Monthly Invoices | Invoice History | Invoice Detail |
List of Invoices | Top Watchlist TLDs | Watchlist Summary | Top Watchlists |
All Watchlists | All Watchlist Details | Domain Name History |
Domain Verification/Suspension
</rptString>
<threemopast>9/12/2011</threemopast>
<sixmopast>6/12/2011</sixmopast>
<begindate>12/6/2011</begindate>
<enddate>12/12/2011</enddate>
<option>
<reporttype id="0" name="Registrations"></reporttype>
<reporttype id="1" name="Renewals"></reporttype>
<reporttype id="2" name="Transfers In"></reporttype>
<reporttype id="11" name="Transfers Out"></reporttype>
<reporttype id="3" name="Private Label"></reporttype>
<reporttype id="4" name="Sub-Accounts"></reporttype>
<reporttype id="5" name="Accounting"></reporttype>
<reporttype id="6" name="Expiring Domains"></reporttype>
<reporttype id="10" name="Pricing Editor Updates"></reporttype>
<reporttype id="13" name="Business Listing Bulk Update"></reporttype>
<reporttype id="15" name="Business Listing Domain Update"></reporttype>
<reporttype id="16" name="Transactions"></reporttype>
<reporttype id="17" name="Order History"></reporttype>
<reporttype id="18" name="Periodic Billing History"></reporttype>
<reporttype id="20" name="Premium Registration"></reporttype>
<reporttype id="21" name="Monthly Invoices"></reporttype>
<reporttype id="22" name="Invoice History"></reporttype>
<reporttype id="23" name="Invoice Detail"></reporttype>
<reporttype id="24" name="List of Invoices"></reporttype>
<reporttype id="25" name="Top Watchlist TLDs"></reporttype>
<reporttype id="26" name="Watchlist Summary"></reporttype>
<reporttype id="27" name="Top Watchlists"></reporttype>
<reporttype id="28" name="All Watchlists"></reporttype>
<reporttype id="29" name="All Watchlist Details"></reporttype>
<reporttype id="30" name="Domain Name History"></reporttype>
<reporttype id="31" name="Domain Verification/Suspension"></reporttype>
</option>
<reporttype>0</reporttype>
<reportname> Registrations </reportname>
<results>
<rpttld>
<report0-single tld="com" sum="2"/>
<report0-single tld="net" sum="1"/>
</rpttld>
<rptrawxml>
<row OrderID="157781016" SLD="resellerdocs1"
TLD="com" CreationDate="2011-12-09T03:25:03.253"
ExpDate="2012-12-09T11:25:00" PaidAmount="8.9500"
BilledToCust="0.0000"/>
<row OrderID="157781016" SLD="resellerdocs123"
TLD="com" CreationDate="2011-12-09T03:25:06.657"
ExpDate="2012-12-09T11:25:00" PaidAmount="8.9500"
BilledToCust="0.0000"/>
<row OrderID="157781021" SLD="resellerdocs12312"
TLD="net" CreationDate="2011-12-09T03:58:41.283"
ExpDate="2012-12-09T03:58:00" PaidAmount="8.9500"
BilledToCust="0.0000"/>
</rptrawxml>
<NewThisMonth>3</NewThisMonth>
<AverageRegPeriod>1.000000</AverageRegPeriod>
</results>
</rpt>
<Command>RPT_GETREPORT</Command>
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
<ExecTime>0.422</ExecTime>
<Done>true</Done>
<RequestDateTime>1/6/2014 4:16:19 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
Related Commands
----------------

[CreateAccount](../docs/createaccount.md)

CreateSubAccount

GetAccountInfo

GetAccountPassword

[GetAllAccountInfo](../docs/getallaccountinfo.md)

GetOrderDetail

GetOrderList

GetReport

GetSubAccounts

GetTransHistory