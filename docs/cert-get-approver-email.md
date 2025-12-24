CertGetApproverEmail
====================

Get list of qualified approver emails

Usage
-----

Retrieve list of qualified approver emails.

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

| Parameter  | Type | Status  | Description |
| ------------ | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Command   | string | Required | CertGetApproverEmail. |
| UID | string | Required | Your Account ID.                                                        |
| PW      | string | Required | Your API Token. |
| CertID | string | Required | ID number for this individual certificate. Retrieve this number using the [CertGetCerts](../docs/cert-get-certs.md) command. |
| Domain    | string | Required | Domain name. |
| ResponseType | string | Optional | Format of response. Permitted values: - Text \(default\) - HTML - XML                              |

Returned Parameters and Values
------------------------------

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output    | Type | Description                                           |
| ------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command    | string | Name of command executed.                                    |
| Success    | boolean | Has this cert been configured successfully?                           |
| ApproverEmail | string | Email address of approver.                                   |
| ErrorCount  | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX     | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done     | string | True value indicates this entire response has reached you successfully.             |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.asp?command=CertGetApproverEmail&uid=resellid&pw=resellpw&CertID=32969&ResponseType=XML
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<CertGetApproverEMail>
  <Success>True</Success>
  <Approver Type="Domain">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Domain">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Generic">
   <ApproverEmail>[email protected]</ApproverEmail>
  </Approver>
  <Approver Type="Manual"></Approver>
</CertGetApproverEMail>
<Command>CERTGETAPPROVEREMAIL</Command>
<APIType>API</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod></MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>SJL0VWAPI07</Server>
<Site>eNom</Site>
<IsLockable></IsLockable>
<IsRealTimeTLD></IsRealTimeTLD>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.000</ExecTime>
<Done>true</Done>
<TrackingKey>50d59809-d71f-411c-98f3-aa3def8e43ee</TrackingKey>
<RequestDateTime>12/16/2016 11:22:32 AM</RequestDateTime>
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