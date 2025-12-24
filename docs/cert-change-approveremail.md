CertChangeApproverEmail
=======================

Change approver email

Usage
-----

Change approver email for a Symantec, Verisign, GeoTrust or Comodo certificate.

Availability
------------

All resellers have access to this command.

Contraints
----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The certificate must belong to this account.

Input parameters
----------------

| Parameter   | Type | Status  | Description |
| ------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command    | string | Required | CertChangeApproverEmail. |
| UID | string | Required | Your Account ID.                                                                                |
| PW      | string | Required | Your API Token. |
| CertID | string | Required | ID number for this individual certificate. Retrieve this number using the [CertGetCerts](../docs/cert-get-certs.md) command.                         |
| ApproverEmail | string | Required | Email addresses of qualified approvers for the domain name embedded in a CSR. Use the [CertGetApproverEmail](../docs/cert-get-approver-email.md) command to retrieve the list |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.                                                   |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output    | Type | Description                                           |
| ------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command    | string | Name of command executed.                                    |
| Success    | boolean | Success status for retrieving the authoritative Whois registrant email.             |
| ApproverEmail | string | Email address of approver.                                   |
| ErrorCount  | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX     | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done     | string | True value indicates this entire response has reached you successfully.             |
| TotalRecords | integer | Total number of record lines returned in the query.                       |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.asp?command=CertChangeApproverEmail&uid=resellid&pw=resellpw&certid=48455&[email protected]&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertChangeApproverEmail>
 <ApproverEmail><![CDATA[[email protected]]]></ApproverEmail>
 <Success>true</Success>
</CertChangeApproverEmail>
<Command>CERTCHANGEAPPROVEREMAIL</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>krkdt198</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>2.736</ExecTime>
<Done>true</Done>
<TrackingKey>5029caf8-48ad-4ff6-a201-fba9b3a88c62</TrackingKey>
<RequestDateTime>9/1/2016 9:36:10 AM</RequestDateTime>
```
Related Commands
----------------

[CertChangeApproverEmail](../docs/cert-change-approveremail.md)

[CertConfigureCert](../docs/cert-configure-cert.md)

[CertGetApproverEmail](../docs/cert-get-approver-email.md)

[CertGetCertDetail](../docs/cert-get-certdetail.md)

[CertGetCerts](../docs/cert-get-certs.md)

[CertModifyOrder](../docs/cert-modify-order.md)

[CertParseCSR](../docs/cert-parse-csr.md)

[CertPurchaseCert](../docs/cert-purchase-cert.md)

[CertResendApproverEmail](../docs/cert-resend-approveremail.md)